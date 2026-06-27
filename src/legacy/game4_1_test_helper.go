package legacy

/*
#include <stdint.h>
#include <stdlib.h>

void nox_xxx_utilNormalizeVector_509F20(float* a1);
int nox_xxx_mobActionGet_50A020(int a1);
int nox_xxx_monsterActionIsCondition_50A010(int a1);
int sub_509FF0(int a1);
*/
import "C"
import "unsafe"

func C_nox_xxx_utilNormalizeVector_509F20(x, y float32) (float32, float32) {
	buf := C.malloc(8)
	defer C.free(buf)
	*(*float32)(unsafe.Pointer(buf)) = x
	*(*float32)(unsafe.Pointer(uintptr(buf) + 4)) = y
	C.nox_xxx_utilNormalizeVector_509F20((*C.float)(buf))
	return *(*float32)(unsafe.Pointer(buf)), *(*float32)(unsafe.Pointer(uintptr(buf) + 4))
}

func C_nox_xxx_mobActionGet_50A020() int {
	// build outer struct with +748 pointing to inner, inner +544 byte = index, inner +24*index used
	outer := C.malloc(2048)
	inner := C.malloc(4096)
	defer C.free(outer)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(outer) + 748)) = uint32(uintptr(inner))
	// set byte at inner+544 to 0, so index =0 => offset 0 => we need inner+0 to hold return value? Actually function returns *(uint32*)(inner +24*(... +23))
	// Let's set *(char*)(inner+544)=1 => then index =1+23=24 => offset 24*24=576
	// Simpler: set byte to 0 => index 23 => offset 552
	*(*uint8)(unsafe.Pointer(uintptr(inner) + 544)) = 0
	// put test value at inner+552
	*(*uint32)(unsafe.Pointer(uintptr(inner) + 552)) = 0x12345678
	return int(C.nox_xxx_mobActionGet_50A020(C.int(uintptr(outer))))
}

func C_nox_xxx_monsterActionIsCondition_50A010(a1 int) int {
	return int(C.nox_xxx_monsterActionIsCondition_50A010(C.int(a1)))
}

func C_sub_509FF0(flag byte) (ret int, out uint32) {
	buf := C.malloc(8)
	inner := C.malloc(32)
	defer C.free(buf)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(buf)) = uint32(uintptr(inner))
	*(*uint8)(unsafe.Pointer(uintptr(inner) + 16)) = flag
	// function expects a1 points to uint32 location holding pointer to inner.
	// If flag 0x20 set, it clears *buf to 0.
	ret = int(C.sub_509FF0(C.int(uintptr(buf))))
	out = *(*uint32)(unsafe.Pointer(buf))
	return
}
