package legacy

/*
#include "server__script__file.h"
#include "common/fs/nox_fs.h"
#include <stdio.h>
#include <stdlib.h>

int nox_script_readInt_542B70(FILE* f);
double nox_script_readFloat_542B90(FILE* f);
int nox_script_writeInt_542BB0(int v, FILE* f);
int nox_script_writeFloat_542BD0(float v, FILE* f);

int test_nox_script_readInt(FILE* f) {
	return nox_script_readInt_542B70(f);
}

double test_nox_script_readFloat(FILE* f) {
	return nox_script_readFloat_542B90(f);
}

int test_nox_script_writeInt(int v, FILE* f) {
	return nox_script_writeInt_542BB0(v, f);
}

int test_nox_script_writeFloat(float v, FILE* f) {
	return nox_script_writeFloat_542BD0(v, f);
}
*/
import "C"
import "unsafe"

func C_nox_script_readInt(f unsafe.Pointer) int {
	return int(C.test_nox_script_readInt((*C.FILE)(f)))
}

func C_nox_script_readFloat(f unsafe.Pointer) float64 {
	return float64(C.test_nox_script_readFloat((*C.FILE)(f)))
}

func C_nox_script_writeInt(v int, f unsafe.Pointer) int {
	return int(C.test_nox_script_writeInt(C.int(v), (*C.FILE)(f)))
}

func C_nox_script_writeFloat(v float32, f unsafe.Pointer) int {
	return int(C.test_nox_script_writeFloat(C.float(v), (*C.FILE)(f)))
}

func C_nox_fs_open_read(path string) unsafe.Pointer {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return unsafe.Pointer(C.nox_fs_open(cPath))
}

func C_nox_fs_open_write(path string) unsafe.Pointer {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	return unsafe.Pointer(C.nox_fs_create(cPath))
}
