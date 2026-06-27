// Command ccoverreport merges GCC coverage data from cgo build directories.
package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
)

type gcovReport struct {
	WorkingDirectory string     `json:"current_working_directory"`
	Files            []gcovFile `json:"files"`
}

type gcovFile struct {
	Name      string         `json:"file"`
	Functions []gcovFunction `json:"functions"`
	Lines     []gcovLine     `json:"lines"`
}

type gcovFunction struct {
	Name           string `json:"name"`
	StartLine      int    `json:"start_line"`
	ExecutionCount int64  `json:"execution_count"`
}

type gcovLine struct {
	Number   int          `json:"line_number"`
	Count    int64        `json:"count"`
	Branches []gcovBranch `json:"branches"`
}

type gcovBranch struct {
	Count int64 `json:"count"`
}

type item struct {
	Count int64
}

type fileCoverage struct {
	Lines     map[int]*item
	Functions map[string]*item
	Branches  map[string]*item
}

type totals struct {
	Covered int `json:"covered"`
	Total   int `json:"total"`
}

type fileResult struct {
	File               string   `json:"file"`
	Lines              totals   `json:"lines"`
	Functions          totals   `json:"functions"`
	Branches           totals   `json:"branches"`
	UncoveredLines     []int    `json:"uncovered_lines,omitempty"`
	UncoveredFunctions []string `json:"uncovered_functions,omitempty"`
}

type result struct {
	Lines     totals       `json:"lines"`
	Functions totals       `json:"functions"`
	Branches  totals       `json:"branches"`
	Files     []fileResult `json:"files"`
}

func main() {
	rootFlag := flag.String("root", "legacy", "legacy source directory")
	workFlag := flag.String("work", "", "Go build work directory containing .gcno files")
	outFlag := flag.String("out", "", "optional detailed JSON report")
	failUnder := flag.Float64("fail-under", 0, "fail when line coverage is below this percentage")
	flag.Parse()

	if *workFlag == "" {
		fatal(errors.New("-work is required"))
	}
	root, err := filepath.Abs(*rootFlag)
	if err != nil {
		fatal(err)
	}
	work, err := filepath.Abs(*workFlag)
	if err != nil {
		fatal(err)
	}

	profiles, err := findFiles(work, ".gcno")
	if err != nil {
		fatal(err)
	}
	if len(profiles) == 0 {
		fatal(fmt.Errorf("no .gcno files found under %s", work))
	}

	coverage, err := readProfiles(root, profiles)
	if err != nil {
		fatal(err)
	}
	expected, err := findFiles(root, ".c")
	if err != nil {
		fatal(err)
	}
	for _, name := range expected {
		rel, err := filepath.Rel(root, name)
		if err != nil {
			fatal(err)
		}
		rel = filepath.ToSlash(rel)
		if _, ok := coverage[rel]; !ok {
			// GCC omits translation units that contain data declarations but no
			// executable lines (legacy/cgo_blobs.c is one such unit).
			coverage[rel] = &fileCoverage{
				Lines:     make(map[int]*item),
				Functions: make(map[string]*item),
				Branches:  make(map[string]*item),
			}
		}
	}

	res := buildResult(coverage)
	printTotals(res, len(expected))
	if *outFlag != "" {
		data, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fatal(err)
		}
		data = append(data, '\n')
		if err := os.WriteFile(*outFlag, data, 0o644); err != nil {
			fatal(err)
		}
	}
	percent := percentage(res.Lines)
	if percent+1e-9 < *failUnder {
		fatal(fmt.Errorf("legacy C line coverage %.2f%% is below %.2f%%", percent, *failUnder))
	}
}

func findFiles(root, ext string) ([]string, error) {
	var out []string
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() && filepath.Ext(entry.Name()) == ext {
			out = append(out, path)
		}
		return nil
	})
	sort.Strings(out)
	return out, err
}

func readProfiles(root string, profiles []string) (map[string]*fileCoverage, error) {
	type profileResult struct {
		report gcovReport
		err    error
	}
	jobs := make(chan string)
	results := make(chan profileResult)
	workers := runtime.GOMAXPROCS(0)
	if workers > 8 {
		workers = 8
	}
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for profile := range jobs {
				data, err := exec.Command("gcov", "--json-format", "--stdout", "--branch-probabilities", profile).Output()
				var report gcovReport
				if err == nil {
					err = json.Unmarshal(data, &report)
				}
				results <- profileResult{report: report, err: err}
			}
		}()
	}
	go func() {
		for _, profile := range profiles {
			jobs <- profile
		}
		close(jobs)
		wg.Wait()
		close(results)
	}()

	merged := make(map[string]*fileCoverage)
	for got := range results {
		if got.err != nil {
			return nil, got.err
		}
		for _, file := range got.report.Files {
			name := file.Name
			if !filepath.IsAbs(name) {
				name = filepath.Join(got.report.WorkingDirectory, name)
			}
			name, err := filepath.Abs(name)
			if err != nil {
				return nil, err
			}
			rel, err := filepath.Rel(root, name)
			if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) || filepath.Ext(name) != ".c" {
				continue
			}
			rel = filepath.ToSlash(rel)
			dst := merged[rel]
			if dst == nil {
				dst = &fileCoverage{
					Lines:     make(map[int]*item),
					Functions: make(map[string]*item),
					Branches:  make(map[string]*item),
				}
				merged[rel] = dst
			}
			for _, line := range file.Lines {
				mergeCount(dst.Lines, line.Number, line.Count)
				for i, branch := range line.Branches {
					mergeCount(dst.Branches, fmt.Sprintf("%d:%d", line.Number, i), branch.Count)
				}
			}
			for _, fn := range file.Functions {
				mergeCount(dst.Functions, fmt.Sprintf("%d:%s", fn.StartLine, fn.Name), fn.ExecutionCount)
			}
		}
	}
	return merged, nil
}

func mergeCount[K comparable](dst map[K]*item, key K, count int64) {
	cur := dst[key]
	if cur == nil {
		cur = new(item)
		dst[key] = cur
	}
	cur.Count += count
}

func buildResult(coverage map[string]*fileCoverage) result {
	names := make([]string, 0, len(coverage))
	for name := range coverage {
		names = append(names, name)
	}
	sort.Strings(names)
	var out result
	for _, name := range names {
		cov := coverage[name]
		file := fileResult{File: name}
		lineNumbers := make([]int, 0, len(cov.Lines))
		for line := range cov.Lines {
			lineNumbers = append(lineNumbers, line)
		}
		sort.Ints(lineNumbers)
		for _, line := range lineNumbers {
			file.Lines.Total++
			if cov.Lines[line].Count > 0 {
				file.Lines.Covered++
			} else {
				file.UncoveredLines = append(file.UncoveredLines, line)
			}
		}
		functionNames := make([]string, 0, len(cov.Functions))
		for fn := range cov.Functions {
			functionNames = append(functionNames, fn)
		}
		sort.Strings(functionNames)
		for _, fn := range functionNames {
			file.Functions.Total++
			if cov.Functions[fn].Count > 0 {
				file.Functions.Covered++
			} else {
				file.UncoveredFunctions = append(file.UncoveredFunctions, fn)
			}
		}
		for _, branch := range cov.Branches {
			file.Branches.Total++
			if branch.Count > 0 {
				file.Branches.Covered++
			}
		}
		addTotals(&out.Lines, file.Lines)
		addTotals(&out.Functions, file.Functions)
		addTotals(&out.Branches, file.Branches)
		out.Files = append(out.Files, file)
	}
	return out
}

func addTotals(dst *totals, src totals) {
	dst.Covered += src.Covered
	dst.Total += src.Total
}

func percentage(value totals) float64 {
	if value.Total == 0 {
		return 100
	}
	return 100 * float64(value.Covered) / float64(value.Total)
}

func printTotals(res result, files int) {
	fmt.Printf("legacy C files:     %d\n", files)
	fmt.Printf("line coverage:      %d/%d (%.2f%%)\n", res.Lines.Covered, res.Lines.Total, percentage(res.Lines))
	fmt.Printf("function coverage:  %d/%d (%.2f%%)\n", res.Functions.Covered, res.Functions.Total, percentage(res.Functions))
	fmt.Printf("branch coverage:    %d/%d (%.2f%%)\n", res.Branches.Covered, res.Branches.Total, percentage(res.Branches))
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, "ccoverreport:", err)
	os.Exit(1)
}
