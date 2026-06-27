package legacy

/*
#include <stdarg.h>
#include <stdlib.h>

void nox_xxx_networkLog_printf_413D30(char* fmt, ...);

static void test_nox_xxx_networkLog_printf_s(const char* fmt, const char* s) {
	nox_xxx_networkLog_printf_413D30((char*)fmt, s);
}

static void test_nox_xxx_networkLog_printf_d(const char* fmt, int val) {
	nox_xxx_networkLog_printf_413D30((char*)fmt, val);
}
*/
import "C"
import "unsafe"

func C_nox_xxx_networkLog_printf_s(fmt string, s string) {
	cFmt := C.CString(fmt)
	cS := C.CString(s)
	defer C.free(unsafe.Pointer(cFmt))
	defer C.free(unsafe.Pointer(cS))
	C.test_nox_xxx_networkLog_printf_s(cFmt, cS)
}

func C_nox_xxx_networkLog_printf_d(fmt string, val int) {
	cFmt := C.CString(fmt)
	defer C.free(unsafe.Pointer(cFmt))
	C.test_nox_xxx_networkLog_printf_d(cFmt, C.int(val))
}
