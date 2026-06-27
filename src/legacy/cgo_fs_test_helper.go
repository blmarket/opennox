package legacy

/*
#include "common/fs/nox_fs.h"
#include <stdio.h>
#include <stdlib.h>

int test_nox_fs_fprintf_s(FILE* f, const char* format, const char* s) {
	return nox_fs_fprintf(f, format, s);
}

int test_nox_fs_fprintf_d(FILE* f, const char* format, int val) {
	return nox_fs_fprintf(f, format, val);
}
*/
import "C"
import "unsafe"

func C_nox_fs_create_text(path string) unsafe.Pointer {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return unsafe.Pointer(C.nox_fs_create_text(cPath))
}

func C_nox_fs_close(f unsafe.Pointer) {
	C.nox_fs_close((*C.FILE)(f))
}

func C_nox_fs_fprintf_s(f unsafe.Pointer, format string, s string) int {
	cFormat := C.CString(format)
	cS := C.CString(s)
	defer C.free(unsafe.Pointer(cFormat))
	defer C.free(unsafe.Pointer(cS))
	return int(C.test_nox_fs_fprintf_s((*C.FILE)(f), cFormat, cS))
}

func C_nox_fs_fprintf_d(f unsafe.Pointer, format string, val int) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_fs_fprintf_d((*C.FILE)(f), cFormat, C.int(val)))
}
