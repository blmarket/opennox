package legacy

/*
#include "client__gui__gamewin__gamewin.h"
#include "client__gui__guimsg.h"
#include "client__gui__window.h"
#include "client__io__win95__focus.h"
#include <stdlib.h>

extern obj_5D4594_754088_t* ptr_5D4594_754088;
extern int ptr_5D4594_754088_cnt;
extern obj_5D4594_754088_t* ptr_5D4594_754092;
extern int ptr_5D4594_754092_cnt;
*/
import "C"

import "unsafe"

func newCTestWindow() unsafe.Pointer {
	return C.calloc(1, C.size_t(C.sizeof_nox_window))
}

func freeCTestWindow(p unsafe.Pointer) {
	C.free(p)
}

func cTestWindowSetRect(p unsafe.Pointer, x, y, w, h int) {
	win := (*C.nox_window)(p)
	win.off_x = C.int(x)
	win.off_y = C.int(y)
	win.width = C.int(w)
	win.height = C.int(h)
	win.end_x = C.int(x + w)
	win.end_y = C.int(y + h)
}

func cTestWindowSetParent(p, parent unsafe.Pointer) {
	(*C.nox_window)(p).parent = (*C.nox_window)(parent)
}

func cTestWindowFlags(p unsafe.Pointer) uint32 {
	return uint32((*C.nox_window)(p).flags)
}

func cTestWindowDrawDataWin(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer((*C.nox_window)(p).draw_data.win)
}

func cNoxGUIGetWindowOffs(p unsafe.Pointer) (int, int, int) {
	x, y := C.uint32_t(0xdeadbeef), C.uint32_t(0xdeadbeef)
	r := C.nox_gui_getWindowOffs_46AA20((*C.nox_window)(p), &x, &y)
	return int(r), int(x), int(y)
}

func cNoxClientWndGetPosition(p unsafe.Pointer) (int32, uint32, uint32) {
	x, y := C.uint32_t(0xdeadbeef), C.uint32_t(0xdeadbeef)
	r := C.nox_client_wndGetPosition_46AA60((*C.nox_window)(p), &x, &y)
	return int32(r), uint32(x), uint32(y)
}

func cNoxWindowGetSize(p unsafe.Pointer) (int32, int32, int32) {
	var w, h C.int32_t
	r := C.nox_window_get_size((*C.nox_window)(p), &w, &h)
	return int32(r), int32(w), int32(h)
}

func cNoxWndPointInWnd(p unsafe.Pointer, x, y int) bool {
	return bool(C.nox_xxx_wndPointInWnd_46AAB0((*C.uint32_t)(p), C.int32_t(x), C.int32_t(y)))
}

func cSub46ACE0EmptyRange(p unsafe.Pointer) {
	C.sub_46ACE0((*C.uint)(p), 2, 1, 1)
}

func cSub46AD20EmptyRange(p unsafe.Pointer) {
	C.sub_46AD20((*C.uint)(p), 2, 1, 1)
}

func cSub46AB20(p unsafe.Pointer, x, y int) int {
	return int(C.sub_46AB20((*C.uint)(p), C.int(x), C.int(y)))
}

func cNoxWnd46ABB0(p unsafe.Pointer, hidden int) int {
	return int(C.nox_xxx_wnd_46ABB0((*C.nox_window)(p), C.int(hidden)))
}

func cNoxWnd46AD60(p unsafe.Pointer, flags int) int {
	return int(C.nox_xxx_wnd_46AD60(C.int(uintptr(p)), C.int(flags)))
}

func cNoxWndClearFlag46AD80(p unsafe.Pointer, flags int) int {
	return int(C.nox_xxx_wndClearFlag_46AD80(C.int(uintptr(p)), C.int(flags)))
}

func cNoxWndGetFlags46ADA0(p unsafe.Pointer) int {
	return int(C.nox_xxx_wndGetFlags_46ADA0(C.int(uintptr(p))))
}

func cNoxWindowIsChild(parent, child unsafe.Pointer) int {
	return int(C.nox_window_is_child((*C.nox_window)(parent), (*C.nox_window)(child)))
}

func cNoxWnd46B280(p, parent unsafe.Pointer) int {
	return int(C.nox_xxx_wnd_46B280(C.int(uintptr(p)), C.int(uintptr(parent))))
}

func cNoxPrintCenteredNil() {
	C.nox_xxx_printCentered_445490(nil)
}

func cNoxClientPickupNil() {
	C.nox_xxx_clientPickup_46C140(nil)
}

func cFocusReset() {
	C.free(unsafe.Pointer(C.ptr_5D4594_754088))
	C.ptr_5D4594_754088 = nil
	C.ptr_5D4594_754088_cnt = 0
	C.free(unsafe.Pointer(C.ptr_5D4594_754092))
	C.ptr_5D4594_754092 = nil
	C.ptr_5D4594_754092_cnt = 0
}

func cFocusPrimaryCount() int {
	return int(C.ptr_5D4594_754088_cnt)
}

func cFocusSecondaryCount() int {
	return int(C.ptr_5D4594_754092_cnt)
}

func cFocusPrimary(i int) *C.obj_5D4594_754088_t {
	return &(*[1 << 20]C.obj_5D4594_754088_t)(unsafe.Pointer(C.ptr_5D4594_754088))[i]
}

func cFocusSecondary(i int) *C.obj_5D4594_754088_t {
	return &(*[1 << 20]C.obj_5D4594_754088_t)(unsafe.Pointer(C.ptr_5D4594_754092))[i]
}

func cFocusPrimaryField(i int) int {
	return int(cFocusPrimary(i).field_4)
}

func cFocusSecondaryField(i int) int {
	return int(cFocusSecondary(i).field_4)
}

func cFocusPrimaryName(i int) string {
	return C.GoString((*C.char)(unsafe.Pointer(&cFocusPrimary(i).name[0])))
}

func cFocusSecondaryName(i int) string {
	return C.GoString((*C.char)(unsafe.Pointer(&cFocusSecondary(i).name[0])))
}
