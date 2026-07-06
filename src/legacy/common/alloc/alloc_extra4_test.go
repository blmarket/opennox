package alloc

import (
	"testing"
)

func TestStringFunctions2(t *testing.T) {
	// Test FreeStrings (should not panic)
	FreeStrings()

	// Test InternCString with empty string
	s := InternCString("")
	_ = s

	// Test GoString with nil
	result := GoString(nil)
	if result != "" {
		t.Error("GoString nil should return empty")
	}
}

func TestMemlogFunctions2(t *testing.T) {
	// Test EnableMemlog
	EnableMemlog(false)
	EnableMemlog(true)
	EnableMemlog(false)
}
