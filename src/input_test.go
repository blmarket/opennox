package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestInputInitMouse(t *testing.T) {
	// inputInitMouse sets a memmap value to 1
	// It should not panic
	inputInitMouse()

	// Verify the value was set
	val := memmap.Uint32(0x5D4594, 1193108)
	if val != 1 {
		t.Errorf("inputInitMouse should set memmap value to 1, got %d", val)
	}
}

func TestNoxMouseSelectOpt(t *testing.T) {
	// Verify the mouse select options are defined correctly
	expected := []string{"Left", "Right", "Middle", "Wheel"}
	if len(noxMouseSelectOpt) != len(expected) {
		t.Errorf("noxMouseSelectOpt length = %d, want %d", len(noxMouseSelectOpt), len(expected))
	}
	for i, v := range expected {
		if noxMouseSelectOpt[i] != v {
			t.Errorf("noxMouseSelectOpt[%d] = %q, want %q", i, noxMouseSelectOpt[i], v)
		}
	}
}
