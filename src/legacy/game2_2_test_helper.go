package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

int sub_486640(void* a1, int a2);
int sub_487C80(int a1);
uint16_t* sub_480250(uint8_t* a1, uint16_t* a2);
int sub_487590(int a1, const void* a2);
void sub_487090(uint32_t** a1);
void sub_481410();
extern uint32_t nox_xxx_waypointCounterMB_587000_154948;
*/
import "C"
import "unsafe"

// C_sub_486640 wraps sub_486640 returning a2 * (*(uint32_t*)(a1+36)>>16) / 100.
func C_sub_486640(vAt36 uint32, a2 int) int {
	buf := C.malloc(64)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 36)) = vAt36
	return int(C.sub_486640(buf, C.int(a2)))
}

// C_sub_487C80_empty wraps sub_487C80 with an empty list node at offset+8 returning 0.
func C_sub_487C80_empty() int {
	container := C.malloc(64)
	defer C.free(container)
	*(*uintptr)(unsafe.Pointer(uintptr(container) + 8)) = 0
	*(*uintptr)(unsafe.Pointer(uintptr(container) + 12)) = 0
	*(*uintptr)(unsafe.Pointer(uintptr(container) + 16)) = 0
	return int(C.sub_487C80(C.int(uintptr(container))))
}

// C_sub_487C80_withNext wraps sub_487C80 where list node points to next, returns 1 if non-zero.
func C_sub_487C80_withNext() int {
	container := C.malloc(64)
	defer C.free(container)
	next := C.malloc(12)
	defer C.free(next)
	*(*uintptr)(unsafe.Pointer(uintptr(container) + 8)) = uintptr(next)
	*(*uintptr)(unsafe.Pointer(uintptr(container) + 12)) = 0
	*(*uintptr)(unsafe.Pointer(uintptr(container) + 16)) = 0
	*(*uintptr)(unsafe.Pointer(next)) = 0
	*(*uintptr)(unsafe.Pointer(uintptr(next) + 4)) = 0
	*(*uintptr)(unsafe.Pointer(uintptr(next) + 8)) = 0
	res := int(C.sub_487C80(C.int(uintptr(container))))
	if res != 0 {
		return 1
	}
	return 0
}

// C_sub_480250 wraps sub_480250 computing packed uint16 from 3 bytes.
func C_sub_480250(b0, b1, b2 byte) uint16 {
	a1 := C.malloc(3)
	defer C.free(a1)
	a2 := C.malloc(2)
	defer C.free(a2)
	*(*byte)(unsafe.Pointer(uintptr(a1) + 0)) = b0
	*(*byte)(unsafe.Pointer(uintptr(a1) + 1)) = b1
	*(*byte)(unsafe.Pointer(uintptr(a1) + 2)) = b2
	C.sub_480250((*C.uint8_t)(a1), (*C.uint16_t)(a2))
	return *(*uint16)(unsafe.Pointer(a2))
}

// C_sub_487590 wraps sub_487590 copying 28 bytes to offset 60 and returning a1.
func C_sub_487590() (ret int, copied [28]byte) {
	buf := C.malloc(128)
	src := C.malloc(28)
	defer C.free(buf)
	defer C.free(src)
	for i := 0; i < 28; i++ {
		*(*byte)(unsafe.Pointer(uintptr(src) + uintptr(i))) = byte(i + 1)
	}
	ret = int(C.sub_487590(C.int(uintptr(buf)), src))
	for i := 0; i < 28; i++ {
		copied[i] = *(*byte)(unsafe.Pointer(uintptr(buf) + 60 + uintptr(i)))
	}
	return
}

// C_sub_487090 wraps sub_487090 list remove on a self-linked node, returns 1 on success.
func C_sub_487090() int {
	node := C.malloc(12)
	defer C.free(node)
	*(*uintptr)(unsafe.Pointer(node)) = uintptr(node)
	*(*uintptr)(unsafe.Pointer(uintptr(node) + 4)) = uintptr(node)
	*(*uintptr)(unsafe.Pointer(uintptr(node) + 8)) = 0
	C.sub_487090((**C.uint32_t)(unsafe.Pointer(node)))
	return 1
}

func C_sub_481410() {
	C.sub_481410()
}

func C_game2_2_getWaypointCounter() uint32 {
	return uint32(C.nox_xxx_waypointCounterMB_587000_154948)
}

func C_game2_2_setWaypointCounter(v uint32) {
	C.nox_xxx_waypointCounterMB_587000_154948 = C.uint32_t(v)
}
