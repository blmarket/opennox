package legacy

import (
	"testing"
	"unsafe"
)

func TestObjectLookLoadersMarkEmptyTablesLoaded(t *testing.T) {
	setMemU32ForTest(t, 0x5D4594, 371256, 0)
	setMemU32ForTest(t, 0x587000, 35500, 0)
	if got := C_nox_xxx_loadLook_415D50(); got != nil {
		t.Fatalf("nox_xxx_loadLook_415D50 returned %p for empty table", got)
	}
	if got := getMemU32ForTest(0x5D4594, 371256); got != 1 {
		t.Fatalf("armor look loaded flag = %d, want 1", got)
	}

	setMemU32ForTest(t, 0x5D4594, 371248, 0)
	setMemU32ForTest(t, 0x587000, 33396, 0)
	if got := C_nox_xxx_loadModifyers_4158C0(); got != nil {
		t.Fatalf("nox_xxx_loadModifyers_4158C0 returned %p for empty table", got)
	}
	if got := getMemU32ForTest(0x5D4594, 371248); got != 1 {
		t.Fatalf("weapon look loaded flag = %d, want 1", got)
	}
}

func setMemU32ForTest(t testing.TB, base, off uintptr, val uint32) {
	t.Helper()
	ptr := (*uint32)(C_mem_getU32Ptr(base, off))
	old := *ptr
	*ptr = val
	t.Cleanup(func() { *ptr = old })
}

func getMemU32ForTest(base, off uintptr) uint32 {
	return *(*uint32)(unsafe.Pointer(C_mem_getU32Ptr(base, off)))
}
