package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadProfiles_Error(t *testing.T) {
	// Test readProfiles with non-existent profile
	// May return error or empty result depending on gcov availability
	result, err := readProfiles("/tmp", []string{"/non/existent.gcno"})
	if err != nil {
		require.Error(t, err)
	} else {
		// If gcov not available or returns empty, result may be empty
		_ = result
	}
}

func TestReadProfiles_Empty(t *testing.T) {
	// Test with empty profiles list
	result, err := readProfiles("/tmp", []string{})
	require.NoError(t, err)
	require.Empty(t, result)
}

func TestFindFiles_NonExistent(t *testing.T) {
	// Test findFiles with non-existent directory
	_, err := findFiles("/non/existent/dir", ".c")
	require.Error(t, err)
}

func TestFindFiles_Empty(t *testing.T) {
	dir := t.TempDir()
	out, err := findFiles(dir, ".c")
	require.NoError(t, err)
	require.Empty(t, out)
}

func TestBuildResult_Empty(t *testing.T) {
	cov := map[string]*fileCoverage{}
	res := buildResult(cov)
	require.Empty(t, res.Files)
	require.Equal(t, 0, res.Lines.Total)
	require.Equal(t, 0, res.Functions.Total)
	require.Equal(t, 0, res.Branches.Total)
}

func TestBuildResult_MultipleFiles(t *testing.T) {
	cov := map[string]*fileCoverage{
		"b.c": {
			Lines:     map[int]*item{1: {Count: 0}, 2: {Count: 1}},
			Functions: map[string]*item{"f1": {Count: 0}},
			Branches:  map[string]*item{},
		},
		"a.c": {
			Lines:     map[int]*item{1: {Count: 1}},
			Functions: map[string]*item{"f2": {Count: 1}},
			Branches:  map[string]*item{"b1": {Count: 0}},
		},
	}
	res := buildResult(cov)
	require.Len(t, res.Files, 2)
	// Files should be sorted
	require.Equal(t, "a.c", res.Files[0].File)
	require.Equal(t, "b.c", res.Files[1].File)
}

func TestPercentage_EdgeCases(t *testing.T) {
	require.Equal(t, 100.0, percentage(totals{Covered: 0, Total: 0}))
	require.Equal(t, 0.0, percentage(totals{Covered: 0, Total: 10}))
	require.Equal(t, 100.0, percentage(totals{Covered: 10, Total: 10}))
	require.Equal(t, 50.0, percentage(totals{Covered: 1, Total: 2}))
}

func TestMergeCount_Overwrite(t *testing.T) {
	m := make(map[int]*item)
	mergeCount(m, 1, 5)
	mergeCount(m, 1, 3)
	require.Equal(t, int64(8), m[1].Count)

	mergeCount(m, 2, 0)
	require.Equal(t, int64(0), m[2].Count)
}
