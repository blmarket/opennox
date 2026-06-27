package legacy

import (
	"testing"
)

func TestGuiCgoWindowFuncs(t *testing.T) {
	// Test nil window pointer
	ret1 := C_nox_window_set_all_funcs(nil)
	if ret1 != -2 {
		t.Errorf("Expected -2 for nil window set_all_funcs, got %d", ret1)
	}

	ret2 := C_nox_xxx_wndSetProc_46B2C0(0)
	if ret2 != -2 {
		t.Errorf("Expected -2 for nil window wndSetProc, got %d", ret2)
	}

	ret3 := C_nox_xxx_wndSetWindowProc_46B300(0)
	if ret3 != -2 {
		t.Errorf("Expected -2 for nil window wndSetWindowProc, got %d", ret3)
	}

	ret4 := C_nox_xxx_wndSetDrawFn_46B340(0)
	if ret4 != -2 {
		t.Errorf("Expected -2 for nil window wndSetDrawFn, got %d", ret4)
	}
}
