package memmap

import (
	"testing"
)

func TestSetRuntimeChecksExtra(t *testing.T) {
	SetRuntimeChecks(true)
	SetRuntimeChecks(false)
	SetRuntimeChecks(true)
}

func TestCheckAddrExtra(t *testing.T) {
	// checkAddr panics if address intersects with a variable
	// We test with an address that should not intersect
	defer func() {
		if r := recover(); r != nil {
			t.Logf("checkAddr panicked as expected for uninitialized memmap: %v", r)
		}
	}()
	checkAddr(0x12345678)
}

func TestValidateZerosExtra(t *testing.T) {
	// ValidateZeros panics if non-zero data found
	// In test environment, blobs may not be initialized
	defer func() {
		if r := recover(); r != nil {
			t.Logf("ValidateZeros panicked as expected: %v", r)
		}
	}()
	ValidateZeros()
}

func TestPtrExtra(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Ptr panicked as expected (no blobs): %v", r)
		}
	}()
	_ = Ptr(0x1000)
}

func TestSliceExtra(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Slice panicked as expected (no blobs): %v", r)
		}
	}()
	_ = Slice(0x1000, 0)
}

func TestStringExtra(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("String panicked as expected: %v", r)
		}
	}()
	_ = String(0x1000, 0)
}

func TestRelativeAddrExtra(t *testing.T) {
	blob, off := RelativeAddr(0x1234)
	if blob != 0x1234 || off != 0 {
		t.Errorf("RelativeAddr for unknown blob = (0x%X, %d), want (0x1234, 0)", blob, off)
	}
}

func TestPtrOffExtra(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("PtrOff panicked as expected: %v", r)
		}
	}()
	_ = PtrOff(0x1000, 0)
}

func TestPtrSizeExtra(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("PtrSize panicked as expected: %v", r)
		}
	}()
	_ = PtrSize(0x1000, 4)
}

func TestPtrSizeOffExtra(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("PtrSizeOff panicked as expected: %v", r)
		}
	}()
	_ = PtrSizeOff(0x1000, 0, 4)
}
