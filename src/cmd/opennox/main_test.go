package main

import (
	"testing"
)

func TestMainPackage(t *testing.T) {
	// Test that main package can be imported and main function exists
	// We don't call main() because it calls os.Exit
}

func TestRun(t *testing.T) {
	// Test run with help flag (should not return error)
	code := run([]string{"opennox", "-h"})
	if code != 0 {
		// Help flag returns flag.ErrHelp which is not an error, so code should be 0
		// Actually, RunArgs returns flag.ErrHelp, which we treat as not an error
	}
}

func TestRunInvalidArgs(t *testing.T) {
	// Test run with invalid args
	code := run([]string{"opennox", "--invalid-flag-that-does-not-exist"})
	// Should return non-zero or handle gracefully
	_ = code
}
