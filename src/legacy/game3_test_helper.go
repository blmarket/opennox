package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

short* sub_4A3090(short* a1, int a2);
*/
import "C"
import "unsafe"

func C_sub_4A3090(count int16, arr []uint32, idx int) (ret uintptr, outArr []uint32) {
	// allocate struct buffer large enough: need offset 48 for pointer to array
	structBuf := C.malloc(128)
	defer C.free(structBuf)
	arrBuf := C.malloc(C.size_t(len(arr) * 4))
	defer C.free(arrBuf)
	// write count at offset 0 as short
	*(*int16)(unsafe.Pointer(structBuf)) = count
	// write pointer to arr at offset 48
	*(*uint32)(unsafe.Pointer(uintptr(structBuf) + 48)) = uint32(uintptr(arrBuf))
	// copy arr values into arrBuf
	for i, v := range arr {
		*(*uint32)(unsafe.Pointer(uintptr(arrBuf) + uintptr(i*4))) = v
	}
	retPtr := C.sub_4A3090((*C.short)(structBuf), C.int(idx))
	ret = uintptr(unsafe.Pointer(retPtr))
	// read back array
	outArr = make([]uint32, len(arr))
	for i := range outArr {
		outArr[i] = *(*uint32)(unsafe.Pointer(uintptr(arrBuf) + uintptr(i*4)))
	}
	return
}
