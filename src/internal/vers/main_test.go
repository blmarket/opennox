package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	// Test with no args
	err := run([]string{})
	if err == nil {
		t.Error("run with no args should return error")
	}

	// Test with unsupported command
	err = run([]string{"unsupported"})
	if err == nil {
		t.Error("run with unsupported command should return error")
	}

	// Test with valid commands (may fail if not in git repo, but shouldn't panic)
	_ = run([]string{"commit"})
	_ = run([]string{"version"})
	_ = run([]string{"full"})
	_ = run([]string{"sha"})
	_ = run([]string{"vers"})
}
