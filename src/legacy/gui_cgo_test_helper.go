package legacy

/*
#include "client__gui__window.h"

int nox_window_set_all_funcs(nox_window* win, int (*a2)(int, int, int, int), int (*draw)(nox_window*, void*), void* a4);
int nox_xxx_wndSetProc_46B2C0(int a1, int (*a2)(int, int, int, int));
int nox_xxx_wndSetWindowProc_46B300(int a1, int (*a2)(int, int, int, int));
int nox_xxx_wndSetDrawFn_46B340(int a1, int (*a2)(int, int));

static int test_nox_window_set_all_funcs(void* win) {
	return nox_window_set_all_funcs((nox_window*)win, NULL, NULL, NULL);
}

static int test_nox_xxx_wndSetProc_46B2C0(int win) {
	return nox_xxx_wndSetProc_46B2C0(win, NULL);
}

static int test_nox_xxx_wndSetWindowProc_46B300(int win) {
	return nox_xxx_wndSetWindowProc_46B300(win, NULL);
}

static int test_nox_xxx_wndSetDrawFn_46B340(int win) {
	return nox_xxx_wndSetDrawFn_46B340(win, NULL);
}
*/
import "C"
import "unsafe"

func C_nox_window_set_all_funcs(win unsafe.Pointer) int {
	return int(C.test_nox_window_set_all_funcs(win))
}

func C_nox_xxx_wndSetProc_46B2C0(win int) int {
	return int(C.test_nox_xxx_wndSetProc_46B2C0(C.int(win)))
}

func C_nox_xxx_wndSetWindowProc_46B300(win int) int {
	return int(C.test_nox_xxx_wndSetWindowProc_46B300(C.int(win)))
}

func C_nox_xxx_wndSetDrawFn_46B340(win int) int {
	return int(C.test_nox_xxx_wndSetDrawFn_46B340(C.int(win)))
}
