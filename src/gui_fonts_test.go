package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestNoxXxxFontLoadMB(t *testing.T) {
	// Test that nox_xxx_fontLoadMB_43F1C0 sets values to 0
	// Set some non-zero values first
	v1 := [5]uint32{1, 2, 3, 4, 5}
	ptr := memmap.PtrOff(0x5D4594, 816464)
	for i, v := range v1 {
		*memmap.PtrUint32(0x5D4594, uintptr(816464+i*4)) = v
	}

	c := &Client{}
	c.nox_xxx_fontLoadMB_43F1C0()

	// Verify values are now 0
	for i := 0; i < 5; i++ {
		val := *memmap.PtrUint32(0x5D4594, uintptr(816464+i*4))
		if val != 0 {
			t.Fatalf("expected 0 at index %d, got %d", i, val)
		}
	}

	// Verify the dword is set to 0
	_ = ptr
}
