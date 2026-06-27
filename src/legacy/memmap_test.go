package legacy

import (
	"testing"
)

func TestMemmapGetPtrSize(t *testing.T) {
	bases := []struct {
		base uintptr
		size int
	}{
		{0x581450, 23472},
		{0x587000, 316820},
		{0x5D4594, 2598284},
		{0x973CE0, 568},
		{0x973F18, 44881},
		{0x85B3FC, 1029636},
		{0x852978, 40},
		{0x973A20, 704},
	}

	for _, tc := range bases {
		// Test valid start pointer
		ptr := C_mem_getPtrSize(tc.base, 0, uintptr(tc.size))
		if ptr == nil {
			t.Errorf("C_mem_getPtrSize failed for base %x with size %d", tc.base, tc.size)
		}

		// Test generic mem_getPtr
		ptr2 := C_mem_getPtr(tc.base, 0)
		if ptr2 == nil {
			t.Errorf("C_mem_getPtr failed for base %x", tc.base)
		}
	}
}

func TestMemmapTypedGetters(t *testing.T) {
	base := uintptr(0x581450)

	if C_mem_getU8Ptr(base, 0) == nil {
		t.Error("C_mem_getU8Ptr failed")
	}
	if C_mem_getI8Ptr(base, 0) == nil {
		t.Error("C_mem_getI8Ptr failed")
	}
	if C_mem_getU16Ptr(base, 0) == nil {
		t.Error("C_mem_getU16Ptr failed")
	}
	if C_mem_getI16Ptr(base, 0) == nil {
		t.Error("C_mem_getI16Ptr failed")
	}
	if C_mem_getU32Ptr(base, 0) == nil {
		t.Error("C_mem_getU32Ptr failed")
	}
	if C_mem_getI32Ptr(base, 0) == nil {
		t.Error("C_mem_getI32Ptr failed")
	}
	if C_mem_getU64Ptr(base, 0) == nil {
		t.Error("C_mem_getU64Ptr failed")
	}
	if C_mem_getI64Ptr(base, 0) == nil {
		t.Error("C_mem_getI64Ptr failed")
	}
	if C_mem_getFloatPtr(base, 0) == nil {
		t.Error("C_mem_getFloatPtr failed")
	}
	if C_mem_getDoublePtr(base, 0) == nil {
		t.Error("C_mem_getDoublePtr failed")
	}
}
