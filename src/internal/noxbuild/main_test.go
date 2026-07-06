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

func TestBuildInvalid(t *testing.T) {
	// Test build with invalid target
	err := build([]string{"invalid-target-xyz"})
	if err == nil {
		t.Fatalf("expected error for invalid target")
	}
}

func TestBuildEmpty(t *testing.T) {
	// Test build with empty targets (uses defTargets) - should fail quickly due to invalid setup
	// We just verify it doesn't panic
	_ = build([]string{})
}

func TestGoBuildInvalid(t *testing.T) {
	// Test goBuild with invalid target - should return error
	err := goBuild("nonexistent/cmd", "testbin", &buildOpts{
		CGO: false,
	})
	if err == nil {
		t.Fatalf("expected error for invalid go build")
	}
}

func TestGoBuildWithTags(t *testing.T) {
	// Test goBuild with tags and safe mode
	*fSafe = true
	defer func() { *fSafe = false }()
	*fOut = t.TempDir()
	*fSrc = t.TempDir()

	err := goBuild("nonexistent", "testbin", &buildOpts{
		CGO:  false,
		Tags: []string{"testtag"},
	})
	if err == nil {
		t.Fatalf("expected error for invalid go build")
	}
}

func TestCgoCompilerOther(t *testing.T) {
	cc, cxx, err := cgoCompiler("clang", "linux")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cc != "clang" {
		t.Fatalf("expected clang, got %s", cc)
	}
	if cxx != "" {
		t.Fatalf("expected empty cxx for non-zig")
	}

	_, _, err = cgoCompiler("zig", "darwin")
	if err == nil {
		t.Fatalf("expected error for unsupported OS darwin")
	}
}

func TestDo(t *testing.T) {
	// Test do with a simple command that should succeed
	err := do("true")
	if err != nil {
		t.Fatalf("expected true to succeed: %v", err)
	}

	// Test do with a command that should fail
	err = do("false")
	if err == nil {
		t.Fatalf("expected false to fail")
	}

	// Test do with invalid command
	err = do("nonexistent-command-xyz")
	if err == nil {
		t.Fatalf("expected error for nonexistent command")
	}
}

func TestDoEnvs(t *testing.T) {
	dir := t.TempDir()
	// Test doEnvs with a simple command
	err := doEnvs(dir, []string{"TEST_VAR=value"}, "true")
	if err != nil {
		t.Fatalf("expected true to succeed: %v", err)
	}

	// Test doEnvs with invalid command
	err = doEnvs(dir, nil, "nonexistent-command-xyz")
	if err == nil {
		t.Fatalf("expected error for nonexistent command")
	}
}

func TestBuildTargetValid(t *testing.T) {
	// Test buildTarget with valid targets but invalid setup - should fail at goBuild stage
	*fOut = t.TempDir()
	*fSrc = t.TempDir()

	targets := []string{
		BinServer, "server",
		BinOpenNox, "client",
		BinOpenNoxHD, "client-hd", "hd",
		BinOpenNoxDebug, "client-debug",
		BinOpenNoxHDDebug, "client-hd-debug",
	}
	for _, target := range targets {
		err := buildTarget(target)
		if err == nil {
			t.Fatalf("expected error for target %s with invalid setup", target)
		}
	}
}
