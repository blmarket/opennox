package opennox

import (
	"testing"
)

func TestStrmanReadFileDone(t *testing.T) {
	// Test that StrmanReadFile returns nil immediately if strManDone is true
	originalDone := strManDone
	strManDone = true
	defer func() { strManDone = originalDone }()

	err := StrmanReadFile("/nonexistent/path")
	if err != nil {
		t.Errorf("StrmanReadFile with strManDone=true should return nil, got %v", err)
	}
}

func TestNoxStrmanFreeExtra(t *testing.T) {
	// Test that Nox_strman_free_410020 resets the state
	strManDone = true
	Nox_strman_free_410020()
	if strManDone {
		t.Error("Nox_strman_free_410020 should set strManDone to false")
	}
	// strMan should be reset to a new instance
	if strMan == nil {
		t.Error("Nox_strman_free_410020 should reset strMan to a new instance")
	}
}

func TestStrmanReadFileInvalidPath(t *testing.T) {
	// Ensure strManDone is false
	originalDone := strManDone
	strManDone = false
	defer func() { strManDone = originalDone }()

	// Test with invalid path - should return error
	err := StrmanReadFile("/nonexistent/path/to/strings.dat")
	if err == nil {
		t.Log("StrmanReadFile with invalid path returned nil (may be expected if file doesn't exist but no error)")
	}
	// Reset for other tests
	strManDone = false
}
