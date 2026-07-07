//go:build !windows

package legacy

/*
#include "windows_compat.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import (
	"path/filepath"
	"sync"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/ifs"
)

type findFileHandle struct {
	matches []string
	idx     int
}

var (
	findFilesMu sync.Mutex
	findFiles   = make(map[unsafe.Pointer]*findFileHandle)
)

func fillFindData(data *C.WIN32_FIND_DATAA, path string) {
	// Get filename from path
	filename := path
	if i := len(path) - 1; i >= 0 {
		for i >= 0 && path[i] != '/' {
			i--
		}
		if i >= 0 {
			filename = path[i+1:]
		}
	}

	// Stat the path
	var attr uint32 = C.FILE_ATTRIBUTE_NORMAL
	var size int64 = 0
	if info, err := ifs.Stat(path); err == nil {
		if info.IsDir() {
			attr = C.FILE_ATTRIBUTE_DIRECTORY
		}
		size = info.Size()
	} else {
		// Try os.Stat as fallback (for paths not in ifs)
		// ifs.Stat should work for most cases
	}

	// Memset to zero
	C.memset(unsafe.Pointer(data), 0, C.sizeof_WIN32_FIND_DATAA)

	data.dwFileAttributes = C.uint32_t(attr)
	data.nFileSizeHigh = C.uint32_t(uint64(size) >> 32)
	data.nFileSizeLow = C.uint32_t(uint64(size) & 0xffffffff)

	// Copy filename into cFileName (max 260)
	cname := (*[C.MAX_PATH]C.char)(unsafe.Pointer(&data.cFileName[0]))
	for i := 0; i < int(C.MAX_PATH)-1 && i < len(filename); i++ {
		cname[i] = C.char(filename[i])
	}
	// Null terminator already zeroed by memset
}

//export compatFindFirstFileA
func compatFindFirstFileA(lpFileName *C.char, lpFindFileData *C.WIN32_FIND_DATAA) C.HANDLE {
	if lpFileName == nil || lpFindFileData == nil {
		return C.HANDLE(unsafe.Pointer(uintptr(0xffffffff)))
	}

	// Convert C string to Go string and normalize
	path := GoString(lpFileName)
	path = ifs.Normalize(path)

	// If pattern ends with ".*", remove it (C code does this before glob)
	if len(path) >= 2 && path[len(path)-2] == '.' && path[len(path)-1] == '*' {
		path = path[:len(path)-2]
	}

	// Use filepath.Glob to find matches
	matches, err := filepath.Glob(path)
	if err != nil || len(matches) == 0 {
		return C.HANDLE(unsafe.Pointer(uintptr(0xffffffff)))
	}

	// Create handle
	h := &findFileHandle{
		matches: matches,
		idx:     0,
	}

	// Allocate a dummy pointer as handle using C.malloc
	ptr := C.malloc(1)
	if ptr == nil {
		return C.HANDLE(unsafe.Pointer(uintptr(0xffffffff)))
	}

	findFilesMu.Lock()
	findFiles[ptr] = h
	findFilesMu.Unlock()

	// Fill find data for first match
	fillFindData(lpFindFileData, matches[0])
	h.idx = 1

	return C.HANDLE(ptr)
}

//export compatFindNextFileA
func compatFindNextFileA(hFindFile C.HANDLE, lpFindFileData *C.WIN32_FIND_DATAA) C.int {
	if hFindFile == nil || lpFindFileData == nil {
		return 0
	}

	ptr := unsafe.Pointer(hFindFile)
	findFilesMu.Lock()
	h, ok := findFiles[ptr]
	findFilesMu.Unlock()

	if !ok {
		return 0
	}

	if h.idx >= len(h.matches) {
		return 0
	}

	fillFindData(lpFindFileData, h.matches[h.idx])
	h.idx++

	return 1
}

//export compatFindClose
func compatFindClose(hFindFile C.HANDLE) C.int {
	if hFindFile == nil {
		return 0
	}

	ptr := unsafe.Pointer(hFindFile)
	findFilesMu.Lock()
	_, ok := findFiles[ptr]
	if ok {
		delete(findFiles, ptr)
	}
	findFilesMu.Unlock()

	if ok {
		C.free(ptr)
	}

	return 1
}
