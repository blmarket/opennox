package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestBuildTargetInvalidExtra(t *testing.T) {
	err := buildTarget("invalid_target_xyz")
	if err == nil {
		t.Error("buildTarget with invalid target should return an error")
	}
}

func TestBuildTargetEmpty(t *testing.T) {
	err := buildTarget("")
	if err == nil {
		t.Error("buildTarget with empty target should return an error")
	}
}

func TestCgoCompilerExtra(t *testing.T) {
	tests := []struct {
		name    string
		goos    string
		wantErr bool
	}{
		{"gcc", "linux", false},
		{"clang", "linux", false},
		{"zig", "linux", false},
		{"zig", "windows", true}, // zig only supports linux in the map
		{"zig", "darwin", true},
		{"zig", "invalid", true},
	}

	for _, tt := range tests {
		cc, cxx, err := cgoCompiler(tt.name, tt.goos)
		if tt.wantErr {
			if err == nil {
				t.Errorf("cgoCompiler(%q, %q) should return an error", tt.name, tt.goos)
			}
		} else {
			if err != nil {
				t.Errorf("cgoCompiler(%q, %q) should not return an error, got %v", tt.name, tt.goos, err)
			}
			if tt.name == "zig" {
				if cc == "" || cxx == "" {
					t.Errorf("cgoCompiler(zig) should return non-empty cc and cxx")
				}
			} else {
				if cc != tt.name {
					t.Errorf("cgoCompiler(%q) cc = %q, want %q", tt.name, cc, tt.name)
				}
				if cxx != "" {
					t.Errorf("cgoCompiler(%q) cxx should be empty for non-zig, got %q", tt.name, cxx)
				}
			}
		}
	}
}

func TestIsDirExtra(t *testing.T) {
	// Test with existing directory
	tmpDir := t.TempDir()
	if !isDir(tmpDir) {
		t.Errorf("isDir(%q) should return true for existing directory", tmpDir)
	}

	// Test with non-existent path
	if isDir("/non/existent/path/xyz") {
		t.Error("isDir with non-existent path should return false")
	}

	// Test with a file (not a directory)
	tmpFile := filepath.Join(tmpDir, "testfile")
	if err := os.WriteFile(tmpFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	if isDir(tmpFile) {
		t.Errorf("isDir(%q) should return false for a file", tmpFile)
	}
}

func TestBuildInvalidTargetExtra(t *testing.T) {
	origOS := *fOS
	defer func() { *fOS = origOS }()
	*fOS = "linux"

	err := build([]string{"invalid_target"})
	if err == nil {
		t.Error("build with invalid target should return an error")
	}
}

func TestBuildEmptyArgsExtra(t *testing.T) {
	// build with empty args uses defTargets, which will fail because we don't have
	// the actual source files, but it should not panic
	origOS := *fOS
	defer func() { *fOS = origOS }()
	*fOS = "linux"

	// This will fail during actual build, but should not panic
	err := build([]string{})
	// We expect an error because the source files don't exist in test environment
	if err == nil {
		t.Log("build with empty args unexpectedly succeeded (may be ok in some environments)")
	}
}
