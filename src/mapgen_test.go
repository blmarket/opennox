package opennox

import (
	"testing"
	"unsafe"
)

func TestNoxXxxMapGenMakeInfo(t *testing.T) {
	// Allocate a buffer for the map info structure
	// The function writes to various offsets, so we need a large enough buffer
	buf := make([]byte, 2048)
	p := unsafe.Pointer(&buf[0])

	// Should not panic
	nox_xxx_mapGenMakeInfo_4D5DB0(p)

	// Verify some of the written values
	// At offset 0: "Generated Map" string
	// At offset 1392: uint32 value 3
	val := *(*uint32)(unsafe.Add(p, 1392))
	if val != 3 {
		t.Errorf("nox_xxx_mapGenMakeInfo should set uint32 at offset 1392 to 3, got %d", val)
	}
}

func TestNoxXxxMapGenMakeInfoStrings(t *testing.T) {
	buf := make([]byte, 2048)
	p := unsafe.Pointer(&buf[0])

	nox_xxx_mapGenMakeInfo_4D5DB0(p)

	// Check that strings were written (not empty)
	// At offset 0: "Generated Map"
	s := string(buf[0:13])
	if s != "Generated Map" {
		t.Errorf("Expected 'Generated Map' at offset 0, got %q", s)
	}
}
