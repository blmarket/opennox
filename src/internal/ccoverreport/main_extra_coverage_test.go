package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPercentage_Extra(t *testing.T) {
	// Test with zero total (should return 100)
	p := percentage(totals{Covered: 0, Total: 0})
	require.Equal(t, 100.0, p)

	// Test with some coverage
	p = percentage(totals{Covered: 50, Total: 100})
	require.Equal(t, 50.0, p)

	// Test with full coverage
	p = percentage(totals{Covered: 100, Total: 100})
	require.Equal(t, 100.0, p)
}

func TestAddTotals_Extra(t *testing.T) {
	var dst totals
	src := totals{Covered: 10, Total: 20}
	addTotals(&dst, src)
	require.Equal(t, 10, dst.Covered)
	require.Equal(t, 20, dst.Total)

	// Add again
	addTotals(&dst, totals{Covered: 5, Total: 10})
	require.Equal(t, 15, dst.Covered)
	require.Equal(t, 30, dst.Total)
}

func TestBuildResult_Extra(t *testing.T) {
	// Test with empty coverage
	coverage := make(map[string]*fileCoverage)
	res := buildResult(coverage)
	require.Empty(t, res.Files)
	require.Equal(t, 0, res.Lines.Total)

	// Test with some data
	coverage["test.c"] = &fileCoverage{
		Lines: map[int]*item{
			1: {Count: 1},
			2: {Count: 0},
			3: {Count: 5},
		},
		Functions: map[string]*item{
			"10:func1": {Count: 1},
			"20:func2": {Count: 0},
		},
		Branches: map[string]*item{
			"1:0": {Count: 1},
			"1:1": {Count: 0},
		},
	}
	res = buildResult(coverage)
	require.Len(t, res.Files, 1)
	require.Equal(t, "test.c", res.Files[0].File)
	require.Equal(t, 3, res.Files[0].Lines.Total)
	require.Equal(t, 2, res.Files[0].Lines.Covered)
	require.Equal(t, []int{2}, res.Files[0].UncoveredLines)
	require.Equal(t, 2, res.Files[0].Functions.Total)
	require.Equal(t, 1, res.Files[0].Functions.Covered)
	require.Equal(t, 2, res.Files[0].Branches.Total)
	require.Equal(t, 1, res.Files[0].Branches.Covered)
}

func TestMergeCount_Extra(t *testing.T) {
	dst := make(map[string]*item)

	mergeCount(dst, "key1", 5)
	require.Equal(t, int64(5), dst["key1"].Count)

	// Merge again
	mergeCount(dst, "key1", 3)
	require.Equal(t, int64(8), dst["key1"].Count)

	// New key
	mergeCount(dst, "key2", 10)
	require.Equal(t, int64(10), dst["key2"].Count)
}

func TestFatal(t *testing.T) {
	// Test fatal by running it in a subprocess (it calls os.Exit)
	if os.Getenv("TEST_FATAL") == "1" {
		fatal(os.ErrInvalid)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestFatal")
	cmd.Env = append(os.Environ(), "TEST_FATAL=1")
	err := cmd.Run()
	require.Error(t, err)
	// fatal should exit with code 1
	if exitErr, ok := err.(*exec.ExitError); ok {
		require.Equal(t, 1, exitErr.ExitCode())
	}
}

func TestPrintTotals_Extra(t *testing.T) {
	res := result{
		Lines:     totals{Covered: 80, Total: 100},
		Functions: totals{Covered: 10, Total: 20},
		Branches:  totals{Covered: 5, Total: 10},
	}
	// Should not panic
	printTotals(res, 5)
}

func TestFindFiles_Extra(t *testing.T) {
	tmp := t.TempDir()

	// Create some test files
	err := os.WriteFile(tmp+"/test.c", []byte("int main() {}"), 0644)
	require.NoError(t, err)
	err = os.WriteFile(tmp+"/test.h", []byte("#pragma once"), 0644)
	require.NoError(t, err)

	// Find .c files
	files, err := findFiles(tmp, ".c")
	require.NoError(t, err)
	require.Len(t, files, 1)

	// Find .h files
	files, err = findFiles(tmp, ".h")
	require.NoError(t, err)
	require.Len(t, files, 1)

	// Find non-existent extension
	files, err = findFiles(tmp, ".xyz")
	require.NoError(t, err)
	require.Empty(t, files)
}

func TestFindFilesError(t *testing.T) {
	// Test with non-existent directory
	_, err := findFiles("/non/existent/path", ".c")
	require.Error(t, err)
}
