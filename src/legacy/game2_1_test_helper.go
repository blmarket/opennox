package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_464B40(int a1, int a2);
int nox_xxx_XorEaxEaxSub_464BA0();
int nox_xxx_inventoryWndProc_464BB0(int a1, int a2);
int nox_xxx_movEax1Sub_4661C0();
int sub_46F060();
int nox_xxx_Proc_46F070();
int sub_46AEC0(int a1, int a2);
int sub_4739E0(uint32_t* a1, void* a2, void* a3);
int sub_473A10(uint32_t* a1, void* a2, uint32_t* a3);
void* sub_46AF40(void* a1p);
*/
import "C"
import "unsafe"

func C_sub_464B40(a1, a2 int) int {
	return int(C.sub_464B40(C.int(a1), C.int(a2)))
}

func C_nox_xxx_XorEaxEaxSub_464BA0() int {
	return int(C.nox_xxx_XorEaxEaxSub_464BA0())
}

func C_nox_xxx_inventoryWndProc_464BB0(a1, a2 int) int {
	return int(C.nox_xxx_inventoryWndProc_464BB0(C.int(a1), C.int(a2)))
}

func C_nox_xxx_movEax1Sub_4661C0() int {
	return int(C.nox_xxx_movEax1Sub_4661C0())
}

func C_sub_46F060() int {
	return int(C.sub_46F060())
}

func C_nox_xxx_Proc_46F070() int {
	return int(C.nox_xxx_Proc_46F070())
}

// C_sub_46AEC0 writes 0 or -2 and sets *(a1+92)=a2 when a1 non-zero
func C_sub_46AEC0(a1NonZero bool, a2 uint32) (ret int, written uint32) {
	var buf unsafe.Pointer
	if a1NonZero {
		buf = C.malloc(128)
		defer C.free(buf)
		*(*uint32)(unsafe.Pointer(uintptr(buf) + 92)) = 0
		ret = int(C.sub_46AEC0(C.int(uintptr(buf)), C.int(a2)))
		written = *(*uint32)(unsafe.Pointer(uintptr(buf) + 92))
	} else {
		ret = int(C.sub_46AEC0(C.int(0), C.int(a2)))
		written = 0
	}
	return
}

// C_sub_4739E0 computes a3 = a2 + *a1 - a1[4] etc, pure arithmetic on caller buffers
func C_sub_4739E0(a1_0, a1_1, a1_4, a1_5, a2_0, a2_4 uint32) (ret uint32, out0, out4 uint32) {
	a1 := C.malloc(24)
	a2 := C.malloc(8)
	a3 := C.malloc(8)
	defer C.free(a1)
	defer C.free(a2)
	defer C.free(a3)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 0)) = a1_0
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 4)) = a1_1
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 16)) = a1_4
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 20)) = a1_5
	*(*uint32)(unsafe.Pointer(uintptr(a2) + 0)) = a2_0
	*(*uint32)(unsafe.Pointer(uintptr(a2) + 4)) = a2_4
	ret = uint32(C.sub_4739E0((*C.uint32_t)(a1), a2, a3))
	out0 = *(*uint32)(unsafe.Pointer(uintptr(a3) + 0))
	out4 = *(*uint32)(unsafe.Pointer(uintptr(a3) + 4))
	return
}

// C_sub_473A10 similar but output to uint32* a3
func C_sub_473A10(a1_0, a1_1, a1_4, a1_5, a2_0, a2_4 uint32) (ret uint32, out0, out1 uint32) {
	a1 := C.malloc(24)
	a2 := C.malloc(8)
	a3 := C.malloc(8)
	defer C.free(a1)
	defer C.free(a2)
	defer C.free(a3)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 0)) = a1_0
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 4)) = a1_1
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 16)) = a1_4
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 20)) = a1_5
	*(*uint32)(unsafe.Pointer(uintptr(a2) + 0)) = a2_0
	*(*uint32)(unsafe.Pointer(uintptr(a2) + 4)) = a2_4
	ret = uint32(C.sub_473A10((*C.uint32_t)(a1), a2, (*C.uint32_t)(a3)))
	out0 = *(*uint32)(unsafe.Pointer(uintptr(a3) + 0))
	out1 = *(*uint32)(unsafe.Pointer(uintptr(a3) + 4))
	return
}

// C_sub_46AF40 returns *(a1+236) or 0 if null
func C_sub_46AF40(withBuf bool, val uint32) uintptr {
	if !withBuf {
		return uintptr(unsafe.Pointer(C.sub_46AF40(nil)))
	}
	buf := C.malloc(256)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 236)) = val
	return uintptr(unsafe.Pointer(C.sub_46AF40(buf)))
}
