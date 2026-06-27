package legacy

import (
	"bytes"
	"testing"
	"unsafe"
)

func TestNoxCharToOemW(t *testing.T) {
	src := []uint16{'T', 'e', 's', 't', '1', '2', '3', 0}
	dst := make([]byte, 20)

	ret := C_nox_CharToOemW(&src[0], &dst[0])
	if ret <= 0 {
		t.Errorf("Expected positive return code from nox_CharToOemW, got %d", ret)
	}

	expected := []byte("Test123\x00")
	if !bytes.HasPrefix(dst, expected[:7]) {
		t.Errorf("Expected destination prefix %q, got %q", expected[:7], dst[:7])
	}
}

func TestGetPlayerClassFromObjPtr(t *testing.T) {
	if unsafe.Sizeof(uintptr(0)) != 4 {
		t.Skip("legacy GameEx ABI passes object pointers through 32-bit int parameters")
	}

	// Allocate main unit buffer
	unitBuf := AllocBuffer(800)
	defer FreeBuffer(unitBuf)

	buf1 := AllocBuffer(300)
	defer FreeBuffer(buf1)

	buf2 := AllocBuffer(2300)
	defer FreeBuffer(buf2)

	// unitBuf + 748 points to buf1 (data_update)
	*(*uintptr)(unsafe.Pointer(uintptr(unitBuf) + 748)) = uintptr(buf1)

	// buf1 + 276 points to buf2
	*(*uintptr)(unsafe.Pointer(uintptr(buf1) + 276)) = uintptr(buf2)

	// Set class byte at buf2 + 2251 to 5
	*(*byte)(unsafe.Pointer(uintptr(buf2) + 2251)) = 5

	class := C_getPlayerClassFromObjPtr(int(uintptr(unitBuf)))
	if class != 5 {
		t.Errorf("Expected player class to be 5, got %d", class)
	}
}

func TestGetFlagValueFromFlagIndex(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 2},
		{2, 4},
		{3, 8},
		{4, 16},
		{5, 32},
		{-3, 0},
	}

	for _, tc := range tests {
		res := C_getFlagValueFromFlagIndex(tc.input)
		if res != tc.expected {
			t.Errorf("getFlagValueFromFlagIndex(%d): expected %d, got %d", tc.input, tc.expected, res)
		}
	}
}

func TestPlayerDropATrap(t *testing.T) {
	// Test nil playerObj (0) - early return 0
	ret := C_playerDropATrap(0)
	if ret != 0 {
		t.Errorf("Expected 0 for nil playerObj, got %d", ret)
	}
}

func TestPlayerInfoStructParsersSentinel(t *testing.T) {
	if got := C_playerInfoStructParser0Sentinel(); got != 0 {
		t.Fatalf("playerInfoStructParser_0(-2) = %d, want 0", got)
	}
	if got := C_playerInfoStructParser1Sentinel(); got != 0 {
		t.Fatalf("playerInfoStructParser_1(-2) = %d, want 0", got)
	}
}

func TestMixMouseKeyboardWeaponRollGuards(t *testing.T) {
	if unsafe.Sizeof(uintptr(0)) != 4 {
		t.Skip("legacy GameEx ABI passes object pointers through 32-bit int parameters")
	}

	unit := AllocBuffer(800)
	defer FreeBuffer(unit)
	update := AllocBuffer(300)
	defer FreeBuffer(update)
	player := AllocBuffer(3700)
	defer FreeBuffer(player)

	*(*uintptr)(unsafe.Pointer(uintptr(unit) + 748)) = uintptr(update)
	*(*uintptr)(unsafe.Pointer(uintptr(update) + 276)) = uintptr(player)

	*(*byte)(unsafe.Pointer(uintptr(player) + 3680)) = 1
	if got := C_mix_MouseKeyboardWeaponRoll(unit, true); got != 0 {
		t.Fatalf("mix_MouseKeyboardWeaponRoll observer guard = %d, want 0", got)
	}

	*(*byte)(unsafe.Pointer(uintptr(player) + 3680)) = 0
	*(*byte)(unsafe.Pointer(uintptr(update) + 88)) = 1
	if got := C_mix_MouseKeyboardWeaponRoll(unit, false); got != 0 {
		t.Fatalf("mix_MouseKeyboardWeaponRoll update guard = %d, want 0", got)
	}
}
