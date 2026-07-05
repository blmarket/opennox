package legacy

import (
	"testing"
	"unsafe"
)

func TestCGuiWindowGeometryAndFlags(t *testing.T) {
	if unsafe.Sizeof(uintptr(0)) != 4 {
		t.Skip("legacy C window APIs pass pointers through 32-bit int parameters")
	}

	parent := newCTestWindow()
	defer freeCTestWindow(parent)
	child := newCTestWindow()
	defer freeCTestWindow(child)
	grandchild := newCTestWindow()
	defer freeCTestWindow(grandchild)

	cTestWindowSetRect(parent, 5, 7, 100, 50)
	cTestWindowSetRect(child, 11, 13, 20, 10)
	cTestWindowSetRect(grandchild, 2, 3, 4, 5)
	cTestWindowSetParent(child, parent)
	cTestWindowSetParent(grandchild, child)

	if r, x, y := cNoxGUIGetWindowOffs(nil); r != -2 || x != 0 || y != 0 {
		t.Fatalf("nil offsets = (%d, %d, %d), want (-2, 0, 0)", r, x, y)
	}
	if r, x, y := cNoxGUIGetWindowOffs(child); r != 0 || x != 11 || y != 13 {
		t.Fatalf("child offsets = (%d, %d, %d), want (0, 11, 13)", r, x, y)
	}

	if r, x, y := cNoxClientWndGetPosition(nil); r != -2 {
		t.Fatalf("nil position result = %d, want -2 with outputs %d,%d", r, x, y)
	}
	if r, x, y := cNoxClientWndGetPosition(grandchild); r != 0 || x != 18 || y != 23 {
		t.Fatalf("grandchild position = (%d, %d, %d), want (0, 18, 23)", r, x, y)
	}

	if r, w, h := cNoxWindowGetSize(nil); r != -2 || w != 0 || h != 0 {
		t.Fatalf("nil size = (%d, %d, %d), want (-2, 0, 0)", r, w, h)
	}
	if r, w, h := cNoxWindowGetSize(child); r != 0 || w != 20 || h != 10 {
		t.Fatalf("child size = (%d, %d, %d), want (0, 20, 10)", r, w, h)
	}

	if !cNoxWndPointInWnd(child, 16, 20) {
		t.Fatal("point inside child window was reported outside")
	}
	if cNoxWndPointInWnd(child, 40, 20) {
		t.Fatal("point outside child window was reported inside")
	}

	if got := cNoxWindowIsChild(parent, grandchild); got != 1 {
		t.Fatalf("grandchild should be child of parent, got %d", got)
	}
	if got := cNoxWindowIsChild(grandchild, parent); got != 0 {
		t.Fatalf("parent should not be child of grandchild, got %d", got)
	}
	if got := cNoxWindowIsChild(nil, child); got != 0 {
		t.Fatalf("nil parent child check got %d, want 0", got)
	}
	if got := cNoxWindowIsChild(parent, nil); got != 0 {
		t.Fatalf("nil child check got %d, want 0", got)
	}

	if got := cNoxWnd46ABB0(nil, 1); got != -2 {
		t.Fatalf("nox_xxx_wnd_46ABB0(nil) = %d, want -2", got)
	}
	if got := cNoxWnd46ABB0(child, 1); got != 0 || cTestWindowFlags(child)&8 == 0 {
		t.Fatalf("hide child result = %d flags=%#x, want hidden bit", got, cTestWindowFlags(child))
	}
	if got := cNoxWnd46ABB0(child, 0); got != 0 || cTestWindowFlags(child)&8 != 0 {
		t.Fatalf("show child result = %d flags=%#x, want hidden bit clear", got, cTestWindowFlags(child))
	}

	if got := cNoxWnd46AD60(nil, 0x20); got != -2 {
		t.Fatalf("nox_xxx_wnd_46AD60(nil) = %d, want -2", got)
	}
	if got := cNoxWnd46AD60(child, 0x20); got != 0 || cTestWindowFlags(child)&0x20 == 0 {
		t.Fatalf("set flag result=%#x flags=%#x, want flag set", got, cTestWindowFlags(child))
	}
	if got := cNoxWndClearFlag46AD80(nil, 0x20); got != -2 {
		t.Fatalf("nox_xxx_wndClearFlag_46AD80(nil) = %d, want -2", got)
	}
	if got := cNoxWndClearFlag46AD80(child, 0x20); got != 0x20 || cTestWindowFlags(child)&0x20 != 0 {
		t.Fatalf("clear flag result=%#x flags=%#x, want old flag and cleared value", got, cTestWindowFlags(child))
	}
	if got := cNoxWndGetFlags46ADA0(nil); got != -2 {
		t.Fatalf("nox_xxx_wndGetFlags_46ADA0(nil) = %d, want -2", got)
	}
	if got := cNoxWndGetFlags46ADA0(child); got != int(cTestWindowFlags(child)) {
		t.Fatalf("nox_xxx_wndGetFlags_46ADA0(child) = %#x, want %#x", got, cTestWindowFlags(child))
	}

	if got := cNoxWnd46B280(nil, parent); got != -2 {
		t.Fatalf("nox_xxx_wnd_46B280(nil) = %d, want -2", got)
	}
	if got := cNoxWnd46B280(child, nil); got != 0 || cTestWindowDrawDataWin(child) != child {
		t.Fatalf("nox_xxx_wnd_46B280 self result=%d draw_data.win=%p child=%p", got, cTestWindowDrawDataWin(child), child)
	}
	if got := cNoxWnd46B280(child, parent); got != 0 || cTestWindowDrawDataWin(child) != parent {
		t.Fatalf("nox_xxx_wnd_46B280 parent result=%d draw_data.win=%p parent=%p", got, cTestWindowDrawDataWin(child), parent)
	}

	cSub46ACE0EmptyRange(parent)
	cSub46AD20EmptyRange(parent)
	if got := cSub46AB20(nil, 1, 2); got != -2 {
		t.Fatalf("sub_46AB20(nil) = %d, want -2", got)
	}
}

func TestCFocusCallbackRegistration(t *testing.T) {
	cFocusReset()
	t.Cleanup(cFocusReset)

	Sub_42EBB0(1, nil, 7, "FocusA")
	Sub_42EBB0(2, nil, 9, "FocusB")
	Sub_42EBB0(99, nil, 11, "Ignored")

	if got := cFocusPrimaryCount(); got != 1 {
		t.Fatalf("primary focus callback count = %d, want 1", got)
	}
	if got := cFocusSecondaryCount(); got != 1 {
		t.Fatalf("secondary focus callback count = %d, want 1", got)
	}
	if got := cFocusPrimaryField(0); got != 7 {
		t.Fatalf("primary callback field = %d, want 7", got)
	}
	if got := cFocusSecondaryField(0); got != 9 {
		t.Fatalf("secondary callback field = %d, want 9", got)
	}
	if got := cFocusPrimaryName(0); got != "FocusA" {
		t.Fatalf("primary callback name = %q, want FocusA", got)
	}
	if got := cFocusSecondaryName(0); got != "FocusB" {
		t.Fatalf("secondary callback name = %q, want FocusB", got)
	}
}

func TestCGuardedGuiEntryPoints(t *testing.T) {
	cNoxPrintCenteredNil()

	setMemU32ForTest(t, 0x5D4594, 1064928, 1)
	cNoxClientPickupNil()
}
