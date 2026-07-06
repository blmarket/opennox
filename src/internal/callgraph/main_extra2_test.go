package main

import (
	"testing"
)

func TestRunExtra(t *testing.T) {
	// run should fail with invalid root
	err := run("/nonexistent")
	if err == nil {
		t.Error("run with invalid root should fail")
	}

	// run should fail without cxgo installed or invalid structure
	err = run("/tmp")
	if err == nil {
		t.Error("run with /tmp should fail")
	}
}
