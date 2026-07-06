package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	// Test run with empty string (should return error about missing ./src)
	err := run("")
	if err == nil {
		t.Error("run with empty root should return error")
	}

	// Test run with invalid directory
	err = run("/nonexistent/path/xyz")
	if err == nil {
		t.Error("run with invalid path should return error")
	}
}

func TestMainFunc(t *testing.T) {
	// main() calls run(os.Args) and os.Exit on error
	// We can't easily test os.Exit, but we can verify main exists
	// and doesn't panic when called with invalid args in a subprocess
	// For now, just verify the function exists
	_ = main
}
