package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPercentage(t *testing.T) {
	if percentage(totals{}) != 100 {
		t.Fatalf("expected 100 for empty")
	}
	if percentage(totals{Covered: 1, Total: 2}) != 50 {
		t.Fatalf("expected 50")
	}
}

func TestAddTotals(t *testing.T) {
	var dst totals
	addTotals(&dst, totals{Covered: 2, Total: 5})
	addTotals(&dst, totals{Covered: 1, Total: 3})
	if dst.Covered != 3 || dst.Total != 8 {
		t.Fatalf("got %+v", dst)
	}
}

func TestMergeCount(t *testing.T) {
	m := make(map[string]*item)
	mergeCount(m, "a", 1)
	mergeCount(m, "a", 2)
	if m["a"].Count != 3 {
		t.Fatalf("expected 3 got %d", m["a"].Count)
	}
}

func TestBuildResult(t *testing.T) {
	cov := map[string]*fileCoverage{
		"a.c": {
			Lines:     map[int]*item{1: {Count: 1}, 2: {Count: 0}},
			Functions: map[string]*item{"f1": {Count: 1}, "f2": {Count: 0}},
			Branches:  map[string]*item{"b1": {Count: 1}, "b2": {Count: 0}},
		},
	}
	res := buildResult(cov)
	if res.Lines.Total != 2 || res.Lines.Covered != 1 {
		t.Fatalf("lines %+v", res.Lines)
	}
	if res.Functions.Total != 2 || res.Functions.Covered != 1 {
		t.Fatalf("funcs %+v", res.Functions)
	}
	if res.Branches.Total != 2 || res.Branches.Covered != 1 {
		t.Fatalf("branches %+v", res.Branches)
	}
	if len(res.Files) != 1 || res.Files[0].File != "a.c" {
		t.Fatalf("files %+v", res.Files)
	}
	if len(res.Files[0].UncoveredLines) != 1 || res.Files[0].UncoveredLines[0] != 2 {
		t.Fatalf("uncovered lines %+v", res.Files[0].UncoveredLines)
	}
	if len(res.Files[0].UncoveredFunctions) != 1 {
		t.Fatalf("uncovered funcs %+v", res.Files[0].UncoveredFunctions)
	}
}

func TestPrintTotals(t *testing.T) {
	// should not panic
	printTotals(result{Lines: totals{1, 2}, Functions: totals{1, 1}, Branches: totals{0, 1}}, 1)
}

func TestFindFiles(t *testing.T) {
	dir := t.TempDir()
	// create files
	mustWrite(t, dir+"/a.c", "a")
	mustWrite(t, dir+"/b.txt", "b")
	mustWrite(t, dir+"/sub/c.c", "c")
	out, err := findFiles(dir, ".c")
	if err != nil {
		t.Fatalf("err %v", err)
	}
	if len(out) != 2 {
		t.Fatalf("expected 2 got %d %+v", len(out), out)
	}
}

func mustWrite(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}
