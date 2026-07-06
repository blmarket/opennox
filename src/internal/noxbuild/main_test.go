package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCgoCompiler(t *testing.T) {
	cc, cxx, err := cgoCompiler("gcc", "linux")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cc != "gcc" {
		t.Fatalf("expected gcc, got %s", cc)
	}
	if cxx != "" {
		t.Fatalf("expected empty cxx")
	}

	cc, cxx, err = cgoCompiler("zig", "linux")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cc == "" || cxx == "" {
		t.Fatalf("expected zig cc/cxx")
	}

	_, _, err = cgoCompiler("zig", "windows")
	if err == nil {
		t.Fatalf("expected error for unsupported OS")
	}
}

func TestIsDir(t *testing.T) {
	dir := t.TempDir()
	if !isDir(dir) {
		t.Fatalf("expected dir to exist")
	}
	fpath := filepath.Join(dir, "file.txt")
	os.WriteFile(fpath, []byte("x"), 0644)
	if isDir(fpath) {
		t.Fatalf("expected file not to be dir")
	}
	if isDir(filepath.Join(dir, "nonexistent")) {
		t.Fatalf("expected nonexistent not to be dir")
	}
}

func TestBuildTargetInvalid(t *testing.T) {
	err := buildTarget("invalid-target-xyz")
	if err == nil {
		t.Fatalf("expected error for invalid target")
	}
}
