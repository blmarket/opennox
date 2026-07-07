package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunNoArgs(t *testing.T) {
	err := run([]string{})
	if err == nil {
		t.Error("run with no args should return an error")
	}
}

func TestRunInvalidCommand(t *testing.T) {
	err := run([]string{"invalid_command_xyz"})
	if err == nil {
		t.Error("run with invalid command should return an error")
	}
}

func TestRunCommit(t *testing.T) {
	origNoNewLine := *fNoNewLine
	origOut := *fOut
	defer func() {
		*fNoNewLine = origNoNewLine
		*fOut = origOut
	}()

	*fNoNewLine = true
	*fOut = "-"

	err := run([]string{"commit"})
	if err != nil {
		t.Errorf("run with 'commit' command should not return an error, got %v", err)
	}

	err = run([]string{"sha"})
	if err != nil {
		t.Errorf("run with 'sha' command should not return an error, got %v", err)
	}
}

func TestRunVersion(t *testing.T) {
	origNoNewLine := *fNoNewLine
	origOut := *fOut
	defer func() {
		*fNoNewLine = origNoNewLine
		*fOut = origOut
	}()

	*fNoNewLine = true
	*fOut = "-"

	err := run([]string{"version"})
	if err != nil {
		t.Errorf("run with 'version' command should not return an error, got %v", err)
	}

	err = run([]string{"vers"})
	if err != nil {
		t.Errorf("run with 'vers' command should not return an error, got %v", err)
	}
}

func TestRunFull(t *testing.T) {
	origNoNewLine := *fNoNewLine
	origOut := *fOut
	defer func() {
		*fNoNewLine = origNoNewLine
		*fOut = origOut
	}()

	*fNoNewLine = true
	*fOut = "-"

	err := run([]string{"full"})
	if err != nil {
		t.Errorf("run with 'full' command should not return an error, got %v", err)
	}
}

func TestRunOutputFile(t *testing.T) {
	tmpDir := t.TempDir()
	outFile := filepath.Join(tmpDir, "output.txt")

	origNoNewLine := *fNoNewLine
	origOut := *fOut
	defer func() {
		*fNoNewLine = origNoNewLine
		*fOut = origOut
	}()

	*fNoNewLine = false
	*fOut = outFile

	err := run([]string{"version"})
	if err != nil {
		t.Fatalf("run with output file should not return an error, got %v", err)
	}

	// Verify file was created and has content
	data, err := os.ReadFile(outFile)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	if len(data) == 0 {
		t.Error("Output file should not be empty")
	}
}

func TestRunOutputFileDash(t *testing.T) {
	origNoNewLine := *fNoNewLine
	origOut := *fOut
	defer func() {
		*fNoNewLine = origNoNewLine
		*fOut = origOut
	}()

	*fNoNewLine = true
	*fOut = "-"

	// Using "-" as output should write to stdout (not a file)
	err := run([]string{"version"})
	if err != nil {
		t.Errorf("run with '-' output should not return an error, got %v", err)
	}
}
