package legacy

/*
#include <stdint.h>
#include <stdlib.h>

struct nox_drawable;
int sub_44E8D0();
int nox_xxx_quickbarButtonBookDraw_45EF30();
int sub_45EF40();
int nox_xxx_quickbar_45F8D0(int a1, int a2, int a3, int a4);
int sub_454000(int a1, int a2);
struct nox_drawable* sub_45A010(struct nox_drawable* dr);
int sub_45F500(int a1, int a2);
int sub_4526D0(int a1);
int sub_459DB0(struct nox_drawable* dr);
int nox_xxx_spriteSetActiveMB_45A990_drawable(int a1);
*/
import "C"
import "unsafe"

// C_sub_44E8D0 wraps sub_44E8D0 which returns constant 1.
func C_sub_44E8D0() int {
	return int(C.sub_44E8D0())
}

// C_nox_xxx_quickbarButtonBookDraw_45EF30 wraps constant 1 return.
func C_nox_xxx_quickbarButtonBookDraw_45EF30() int {
	return int(C.nox_xxx_quickbarButtonBookDraw_45EF30())
}

// C_sub_45EF40 wraps constant 0 return.
func C_sub_45EF40() int {
	return int(C.sub_45EF40())
}

// C_nox_xxx_quickbar_45F8D0 wraps pure logic on a2.
func C_nox_xxx_quickbar_45F8D0(a1, a2, a3, a4 int) int {
	return int(C.nox_xxx_quickbar_45F8D0(C.int(a1), C.int(a2), C.int(a3), C.int(a4)))
}

// C_sub_454000 tests bitmask lookup in caller buffer.
func C_sub_454000(a2 int, setBit bool) int {
	buf := C.malloc(1024)
	defer C.free(buf)
	// zero buffer
	for i := 0; i < 1024; i++ {
		*(*byte)(unsafe.Pointer(uintptr(buf) + uintptr(i))) = 0
	}
	idx := (a2 / 32) & 0xFF
	off := idx * 4
	if setBit {
		bit := uint32(1 << (a2 % 32))
		*(*uint32)(unsafe.Pointer(uintptr(buf) + uintptr(off))) = bit
	}
	return int(C.sub_454000(C.int(uintptr(buf)), C.int(a2)))
}

// C_sub_45A010 returns dr->field_104 (offset 416).
func C_sub_45A010() (ret uintptr, expected uintptr) {
	dr := C.malloc(512)
	defer C.free(dr)
	target := C.malloc(16)
	defer C.free(target)
	*(*uintptr)(unsafe.Pointer(uintptr(dr) + 416)) = uintptr(target)
	ret = uintptr(unsafe.Pointer(C.sub_45A010((*C.struct_nox_drawable)(dr))))
	expected = uintptr(target)
	return
}

// C_sub_45F500 reads bit 1 from *(uint32_t*)(*(uint32_t*)(a2+4*a1+232)+36)
func C_sub_45F500(a1 int, bitSet bool) int {
	a2 := C.malloc(1024)
	defer C.free(a2)
	inner := C.malloc(64)
	defer C.free(inner)
	// zero
	for i := 0; i < 64; i++ {
		*(*byte)(unsafe.Pointer(uintptr(inner) + uintptr(i))) = 0
	}
	if bitSet {
		*(*uint32)(unsafe.Pointer(uintptr(inner) + 36)) = 2 // bit1 set
	}
	off := 4*a1 + 232
	*(*uint32)(unsafe.Pointer(uintptr(a2) + uintptr(off))) = uint32(uintptr(inner))
	return int(C.sub_45F500(C.int(a1), C.int(uintptr(a2))))
}

// C_sub_4526D0 writes 4 to *(uint32_t*)(*(uint32_t*)(a1+152)+28)
func C_sub_4526D0() (ret int, written uint32) {
	a1 := C.malloc(256)
	defer C.free(a1)
	inner := C.malloc(64)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 152)) = uint32(uintptr(inner))
	*(*uint32)(unsafe.Pointer(uintptr(inner) + 28)) = 0
	ret = int(C.sub_4526D0(C.int(uintptr(a1))))
	written = *(*uint32)(unsafe.Pointer(uintptr(inner) + 28))
	return
}

// C_sub_459DB0 checks flags at +112 and +116
func C_sub_459DB0(flag112 uint32, flag116 uint8) int {
	dr := C.malloc(512)
	defer C.free(dr)
	*(*uint32)(unsafe.Pointer(uintptr(dr) + 112)) = flag112
	*(*uint8)(unsafe.Pointer(uintptr(dr) + 116)) = flag116
	return int(C.sub_459DB0((*C.struct_nox_drawable)(dr)))
}

// C_nox_xxx_spriteSetActiveMB_45A990_drawable sets bit 4 at offset 120 and returns a1
func C_nox_xxx_spriteSetActiveMB_45A990_drawable() (ret int, flagsAfter uint32) {
	buf := C.malloc(256)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 120)) = 0
	ret = int(C.nox_xxx_spriteSetActiveMB_45A990_drawable(C.int(uintptr(buf))))
	flagsAfter = *(*uint32)(unsafe.Pointer(uintptr(buf) + 120))
	return
}
