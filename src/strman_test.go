package opennox

import (
	"testing"
)

func TestStrmanReadFile(t *testing.T) {
	// Reset state
	Nox_strman_free_410020()
	defer Nox_strman_free_410020()

	// Test with non-existent file
	err := StrmanReadFile("/non/existent/path.str")
	if err == nil {
		t.Error("StrmanReadFile with non-existent file should return error")
	}

	// Test that second call after failure still tries (strManDone is false on error)
	err = StrmanReadFile("/non/existent/path2.str")
	if err == nil {
		t.Error("StrmanReadFile with non-existent file should return error")
	}
}

func TestNoxStrmanFree(t *testing.T) {
	// Ensure no panic
	Nox_strman_free_410020()
	Nox_strman_free_410020() // Call twice to ensure idempotency
}
