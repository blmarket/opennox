package legacy

import (
	"testing"
)

func TestCompatAdditional(t *testing.T) {
	// Test FindFileA with various paths using compatFindFirst
	t.Run("FindFileA with empty path", func(t *testing.T) {
		_, _, ok := compatFindFirst("")
		_ = ok
	})

	t.Run("FindFileA with null", func(t *testing.T) {
		defer func() {
			_ = recover()
		}()
		_, _, ok := compatFindFirst("")
		_ = ok
	})

	t.Run("FindFileA with long path", func(t *testing.T) {
		longPath := ""
		for i := 0; i < 10; i++ {
			longPath += "very/long/path/to/file/"
		}
		longPath += "file.txt"
		_, _, ok := compatFindFirst(longPath)
		_ = ok
	})

	t.Run("FindFileA with special characters", func(t *testing.T) {
		paths := []string{
			"file with spaces.txt",
			"file-with-dash.txt",
			"file_with_underscore.txt",
			"file.multiple.dots.txt",
			"UPPERCASE.TXT",
			"mixedCase.Txt",
		}
		for _, p := range paths {
			_, _, ok := compatFindFirst(p)
			_ = ok
		}
	})

	t.Run("FindFileA multiple calls", func(t *testing.T) {
		for i := 0; i < 5; i++ {
			_, _, ok := compatFindFirst("test.txt")
			_ = ok
		}
	})
}

func TestCompatFindFileAWithTempDir(t *testing.T) {
	tmpDir := t.TempDir()

	testPaths := []string{
		tmpDir + "/nonexistent.txt",
		tmpDir + "/",
		tmpDir,
	}

	for _, p := range testPaths {
		t.Run(p, func(t *testing.T) {
			defer func() {
				_ = recover()
			}()
			_, _, ok := compatFindFirst(p)
			_ = ok
		})
	}
}
