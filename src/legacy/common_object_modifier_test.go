package legacy

import (
	"testing"
)

func TestCommonObjectModifierNullsubs(t *testing.T) {
	// Execute all empty subs to cover those lines
	C_nullsubs()
}

func TestCommonObjectModifierSub413480(t *testing.T) {
	// Call with non-matching char (e.g. 0xff) to cover loop and not found path
	res := C_sub_413480(0xff)
	if res != nil {
		t.Errorf("Expected nil for non-matching modifier char, got %p", res)
	}
}
