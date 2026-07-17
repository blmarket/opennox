package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include "client__draw__parse__parse.h"
*/
import "C"

import "unsafe"

func C_getAnimationKindID(name string) int {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return int(C.get_animation_kind_id_44B4C0(cname))
}

func C_loadVectorAnimatedHeader(f unsafe.Pointer) (ret int, frames, delay uint16, kind int) {
	anim := C.calloc(1, 64)
	defer C.free(anim)
	ret = int(C.nox_xxx_loadVectorAnimated_44B8B0(C.int(uintptr(anim)), (*C.nox_memfile)(f)))
	frames = *(*uint16)(unsafe.Pointer(uintptr(anim) + 40))
	delay = *(*uint16)(unsafe.Pointer(uintptr(anim) + 42))
	kind = *(*int)(unsafe.Pointer(uintptr(anim) + 44))
	return
}

func C_loadVectorAnimatedEmptyDirections() int {
	anim := C.calloc(1, 64)
	defer C.free(anim)
	ret := int(C.nox_xxx_loadVectorAnimated_44BC50(C.int(uintptr(anim)), nil))
	for _, off := range []uintptr{4, 8, 12, 16, 24, 28, 32, 36} {
		ptr := *(*uint32)(unsafe.Pointer(uintptr(anim) + off))
		if ptr != 0 {
			C.free(unsafe.Pointer(uintptr(ptr)))
		}
	}
	return ret
}

func C_loadStaticRandomEmpty(f unsafe.Pointer) (ok bool, size uint32, count uint8) {
	attr := C.calloc(1, 256)
	defer C.free(attr)
	data := C.nox_xxx_spriteLoadStaticRandomData_44C000((*C.char)(attr), (*C.nox_memfile)(f))
	if data == nil {
		return false, 0, 0
	}
	defer C.free(data)
	size = *(*uint32)(data)
	count = *(*uint8)(unsafe.Pointer(uintptr(data) + 8))
	items := *(*uint32)(unsafe.Pointer(uintptr(data) + 4))
	if items != 0 {
		C.free(unsafe.Pointer(uintptr(items)))
	}
	return true, size, count
}
