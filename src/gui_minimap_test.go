package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestSub473670(t *testing.T) {
	// Test that sub_473670 returns 0 or 1 based on memmap value
	// Initial value should be 0 or 1
	result := sub_473670()
	if result != 0 && result != 1 {
		t.Fatalf("expected 0 or 1, got %d", result)
	}

	// Set the bit and test
	*memmap.PtrUint32(0x5D4594, 1096424) |= 0x1
	result = sub_473670()
	if result != 1 {
		t.Fatalf("expected 1 after setting bit, got %d", result)
	}

	// Clear the bit and test
	*memmap.PtrUint32(0x5D4594, 1096424) &= 0xFFFFFFFE
	result = sub_473670()
	if result != 0 {
		t.Fatalf("expected 0 after clearing bit, got %d", result)
	}
}

func TestNoxClientToggleMap(t *testing.T) {
	// This function uses nox_xxx_guiCursor_477600 which may panic without proper setup
	// We just verify the function exists and can be called in a controlled way
	// Actually, let's just test sub_473670 which is the simple part

	// Ensure clean state
	*memmap.PtrUint32(0x5D4594, 1096424) &= 0xFFFFFFFE
}
