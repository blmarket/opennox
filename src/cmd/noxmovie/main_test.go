package main

import "testing"

func TestRun(t *testing.T) {
	// Test run with a non-existent file, should return an error
	err := run("/non/existent/file.vqa")
	if err == nil {
		t.Error("expected error for non-existent file, got nil")
	}
}

func TestRunEmpty(t *testing.T) {
	// Test run with empty filename, should return an error
	err := run("")
	if err == nil {
		t.Error("expected error for empty filename, got nil")
	}
}
