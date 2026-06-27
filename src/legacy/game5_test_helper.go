package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int nox_xxx_strikeBomber_549BB0();
int sub_5456B0(int a1);
int sub_5456C0(int a1);
int sub_546410(int a1);
int sub_546420(int a1);
int sub_547EE0(int a1, unsigned char a2);
*/
import "C"
import "unsafe"

func C_nox_xxx_strikeBomber_549BB0() int {
	return int(C.nox_xxx_strikeBomber_549BB0())
}

// helper to build outer->inner structure for sub_534750 family.
// outer at +748 points to inner; inner at +1440 holds flags uint32.
func alloc534() (outer, inner uintptr) {
	outerBuf := C.malloc(2048)
	innerBuf := C.malloc(2048)
	outer = uintptr(outerBuf)
	inner = uintptr(innerBuf)
	*(*uint32)(unsafe.Pointer(outer + 748)) = uint32(inner)
	return outer, inner
}
func free534(outer, inner uintptr) {
	C.free(unsafe.Pointer(outer))
	C.free(unsafe.Pointer(inner))
}

func C_sub_5456B0(initial uint32) (ret uint32, after uint32) {
	outer, inner := alloc534()
	defer free534(outer, inner)
	*(*uint32)(unsafe.Pointer(inner + 1440)) = initial
	ret = uint32(C.sub_5456B0(C.int(outer)))
	after = *(*uint32)(unsafe.Pointer(inner + 1440))
	return
}
func C_sub_5456C0(initial uint32) (ret uint32, after uint32) {
	outer, inner := alloc534()
	defer free534(outer, inner)
	*(*uint32)(unsafe.Pointer(inner + 1440)) = initial
	ret = uint32(C.sub_5456C0(C.int(outer)))
	after = *(*uint32)(unsafe.Pointer(inner + 1440))
	return
}
func C_sub_546410(initial uint32) (ret uint32, after uint32) {
	outer, inner := alloc534()
	defer free534(outer, inner)
	*(*uint32)(unsafe.Pointer(inner + 1440)) = initial
	ret = uint32(C.sub_546410(C.int(outer)))
	after = *(*uint32)(unsafe.Pointer(inner + 1440))
	return
}
func C_sub_546420(initial uint32) (ret uint32, after uint32) {
	outer, inner := alloc534()
	defer free534(outer, inner)
	*(*uint32)(unsafe.Pointer(inner + 1440)) = initial
	ret = uint32(C.sub_546420(C.int(outer)))
	after = *(*uint32)(unsafe.Pointer(inner + 1440))
	return
}

// sub_547EE0 reads byte at +480 &1 and byte at +477 & a2
func C_sub_547EE0(flag480 uint8, flag477 uint8, a2 uint8) int {
	buf := C.malloc(1024)
	defer C.free(buf)
	*(*uint8)(unsafe.Pointer(uintptr(buf) + 480)) = flag480
	*(*uint8)(unsafe.Pointer(uintptr(buf) + 477)) = flag477
	return int(C.sub_547EE0(C.int(uintptr(buf)), C.uchar(a2)))
}
