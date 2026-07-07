//go:build !windows

package legacy

/*
#include "windows_compat.h"
#include <stdlib.h>

static int test_compat_find_first(const char* path, WIN32_FIND_DATAA* data, HANDLE* out) {
	HANDLE h = FindFirstFileA(path, data);
	if (h == (HANDLE)-1) {
		return 0;
	}
	*out = h;
	return 1;
}

static int test_compat_find_next(HANDLE h, WIN32_FIND_DATAA* data) {
	return FindNextFileA(h, data);
}
*/
import "C"
import "unsafe"

type compatFindData struct {
	FileName   string
	Attributes uint32
	Size       uint64
}

const compatFileAttributeDirectory = uint32(C.FILE_ATTRIBUTE_DIRECTORY)

func compatFindFirst(pattern string) (unsafe.Pointer, compatFindData, bool) {
	cPattern := C.CString(pattern)
	defer C.free(unsafe.Pointer(cPattern))
	var data C.WIN32_FIND_DATAA
	var h C.HANDLE
	if C.test_compat_find_first(cPattern, &data, &h) == 0 {
		return nil, compatFindData{}, false
	}
	return unsafe.Pointer(h), compatFindDataFromC(&data), true
}

func compatFindNext(h unsafe.Pointer) (compatFindData, bool) {
	var data C.WIN32_FIND_DATAA
	if C.test_compat_find_next((C.HANDLE)(h), &data) == 0 {
		return compatFindData{}, false
	}
	return compatFindDataFromC(&data), true
}

func closeCompatFind(h unsafe.Pointer) bool {
	return C.FindClose((C.HANDLE)(h)) != 0
}

func compatFindDataFromC(data *C.WIN32_FIND_DATAA) compatFindData {
	return compatFindData{
		FileName:   C.GoString((*C.char)(unsafe.Pointer(&data.cFileName[0]))),
		Attributes: uint32(data.dwFileAttributes),
		Size:       uint64(data.nFileSizeHigh)<<32 | uint64(data.nFileSizeLow),
	}
}
