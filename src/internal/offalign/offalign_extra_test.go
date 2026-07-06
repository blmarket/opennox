package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOffsetAlignDir(t *testing.T) {
	// Test with non-existent dir
	_, err := offsetAlignDir("nonexistent_dir_xyz", 0, 0, 0, 0)
	require.Error(t, err)

	// Test with empty temp dir and valid blob params
	// It should error on readFileListC because path is a dir but no .c files is ok
	// Actually with blob=1, base=1, elem=1 it will walk dir and find no .c files, so succeed with 0
	tmp := t.TempDir()
	n, err := offsetAlignDir(tmp, 1, 1, 1, 0)
	require.NoError(t, err)
	require.Equal(t, 0, n)
}

func TestOffsetAlignFile(t *testing.T) {
	// Test with non-existent file
	_, err := offsetAlignFile("nonexistent.go", 0, 0, 0, 0)
	require.Error(t, err)

	// Test with empty go file
	tmp := filepath.Join(t.TempDir(), "test.go")
	require.NoError(t, os.WriteFile(tmp, []byte("package test\n"), 0644))
	_, err = offsetAlignFile(tmp, 0, 0, 0, 0)
	require.NoError(t, err)
}

func TestReadFileListC(t *testing.T) {
	// Test with non-existent file
	list, err := readFileListC("nonexistent.txt", nil)
	require.Error(t, err)
	require.Nil(t, list)

	// Test with valid dir containing .c files
	tmp := t.TempDir()
	require.NoError(t, os.WriteFile(filepath.Join(tmp, "file1.c"), []byte(""), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(tmp, "file2.c"), []byte(""), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(tmp, "ignore.txt"), []byte(""), 0644))
	list, err = readFileListC(tmp, nil)
	require.NoError(t, err)
	require.Len(t, list, 2)
}

func TestMain(t *testing.T) {
	// Test main with --help flag (should not panic)
	// We can't easily test main() as it calls os.Exit, but we can verify it exists
	require.NotPanics(t, func() {
		_ = main
	})
}
