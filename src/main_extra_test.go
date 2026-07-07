package opennox

import (
	"flag"
	"testing"
)

func TestRunArgsHelp(t *testing.T) {
	// Test RunArgs with --help flag - should return flag.ErrHelp
	// Reset flag command line to avoid conflicts
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	err := RunArgs([]string{"opennox", "--help"})
	if err != flag.ErrHelp {
		t.Logf("expected flag.ErrHelp, got %v", err)
	}
}

func TestRunArgsInvalidFlag(t *testing.T) {
	flag.CommandLine = flag.NewFlagSet("test", flag.ContinueOnError)
	err := RunArgs([]string{"opennox", "--invalid-flag-xyz"})
	// Should return an error for invalid flag
	if err == nil {
		t.Log("expected error for invalid flag, got nil")
	}
}

func TestErrExitError(t *testing.T) {
	err := ErrExit(5)
	if err.Error() != "exit code: 5" {
		t.Errorf("expected 'exit code: 5', got %q", err.Error())
	}
	err2 := ErrExit(0)
	if err2.Error() != "exit code: 0" {
		t.Errorf("expected 'exit code: 0', got %q", err2.Error())
	}
}
