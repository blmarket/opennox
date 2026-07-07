package main

import (
	"testing"
)

func TestMainPackage(t *testing.T) {
	// Test that main package can be imported and main function exists
	// We don't call main() because it calls os.Exit
}

func TestRun(t *testing.T) {
	// Test run function - it calls RunAudioTest which may fail without audio device
	// We just verify it doesn't panic
	code := run([]string{"opennox-audio", "-h"})
	_ = code
}
