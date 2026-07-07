package blobs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSetPathAndPath(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	SetPath("/tmp/testpath")
	if got := Path(); got != "/tmp/testpath" {
		t.Errorf("Path() = %q, want %q", got, "/tmp/testpath")
	}

	if got := Path("a", "b", "c"); got != filepath.Join("/tmp/testpath", "a", "b", "c") {
		t.Errorf("Path(a,b,c) = %q, want %q", got, filepath.Join("/tmp/testpath", "a", "b", "c"))
	}

	SetPath(".")
	if got := Path(); got != "." {
		t.Errorf("Path() = %q, want %q", got, ".")
	}
}

func TestReadFile(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	tmpDir := t.TempDir()
	SetPath(tmpDir)

	// Create a test file
	testContent := []byte("test content")
	testFile := filepath.Join(tmpDir, "test.txt")
	if err := os.WriteFile(testFile, testContent, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	data, err := ReadFile("test.txt")
	if err != nil {
		t.Fatalf("ReadFile failed: %v", err)
	}
	if string(data) != string(testContent) {
		t.Errorf("ReadFile content = %q, want %q", string(data), string(testContent))
	}

	// Test non-existent file
	_, err = ReadFile("nonexistent.txt")
	if err == nil {
		t.Error("ReadFile should fail for non-existent file")
	}
}

func TestEachFile(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	tmpDir := t.TempDir()
	SetPath(tmpDir)

	// Create test files
	files := []string{"test1.c", "test2.go", "test3.txt", "vardefs.c", "memmap.go"}
	for _, f := range files {
		if err := os.WriteFile(filepath.Join(tmpDir, f), []byte("content"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Create a subdirectory (should be skipped)
	subDir := filepath.Join(tmpDir, "subdir")
	if err := os.Mkdir(subDir, 0755); err != nil {
		t.Fatalf("Failed to create subdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(subDir, "test.c"), []byte("content"), 0644); err != nil {
		t.Fatalf("Failed to create test file in subdir: %v", err)
	}

	var found []string
	err := EachFile(func(path string) error {
		found = append(found, path)
		return nil
	})
	if err != nil {
		t.Fatalf("EachFile failed: %v", err)
	}

	// Should find test1.c and test2.go, but not test3.txt (wrong ext), vardefs.c (excluded), memmap.go (excluded)
	// Note: memmap.go is excluded because it's in the exclude list as memmapGo1
	expected := map[string]bool{"test1.c": true, "test2.go": true}
	if len(found) != len(expected) {
		t.Errorf("EachFile found %d files, want %d: %v", len(found), len(expected), found)
	}
	for _, f := range found {
		if !expected[f] {
			t.Errorf("EachFile found unexpected file: %q", f)
		}
	}
}

func TestEachFileError(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	// Set to non-existent directory
	SetPath("/nonexistent/path/that/does/not/exist")

	err := EachFile(func(path string) error {
		return nil
	})
	if err == nil {
		t.Error("EachFile should fail for non-existent directory")
	}
}
