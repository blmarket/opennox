package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_4BD680(int a1);
int sub_4BD710(int a1);
int sub_4BD300(uint32_t* a1, int a2);
int sub_4BDB20(int a1);
int sub_4BDB30(int a1);
int nox_xxx_updDrawMonsterGen_4BC920();
int sub_4BDDA0();
int sub_4BE320();
int sub_4C24A0();
int sub_4C2BD0();
int sub_4C2BE0();
int nox_xxx_updDrawUndeadKiller_4CCCF0();
*/
import "C"
import "unsafe"

// These wrappers back the GAME3_1.c helpers with C.malloc'd buffers so the GC
// never relocates memory the C code dereferences via raw offsets. The legacy
// 32-bit ABI passes object pointers as plain ints.

// C_sub_4BD680 reads the uint32 field at offset +12 of a struct.
func C_sub_4BD680(valAt12 uint32) int {
	buf := C.malloc(64)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 12)) = valAt12
	return int(C.sub_4BD680(C.int(uintptr(buf))))
}

// C_sub_4BD710 returns the input pointer advanced by 24. It reports both the
// base and the result so the caller can assert the +24 offset.
func C_sub_4BD710() (base, ret int) {
	buf := C.malloc(64)
	defer C.free(buf)
	base = int(uintptr(buf))
	return base, int(C.sub_4BD710(C.int(base)))
}

// C_sub_4BD300 pushes node (at a2-4) onto the free list whose head is *a1.
// Returns the function result, the head value after the call, and the value
// written into the node's next-field, plus the original head value.
func C_sub_4BD300() (ret int, headAfter, nodeNext, oldHead uint32) {
	head := C.malloc(8)
	node := C.malloc(64)
	defer C.free(head)
	defer C.free(node)
	oldHead = 0xDEADBEEF
	*(*uint32)(unsafe.Pointer(head)) = oldHead
	a2 := int(uintptr(node)) + 32 // a2-4 lands safely inside the node buffer
	ret = int(C.sub_4BD300((*C.uint32_t)(head), C.int(a2)))
	headAfter = *(*uint32)(unsafe.Pointer(head))
	nodeNext = *(*uint32)(unsafe.Pointer(uintptr(node) + 32 - 4))
	return
}

// C_sub_4BDB20 sets bit 0x10 in the uint32 flags at offset +124.
func C_sub_4BDB20(initial uint32) (ret int, flagsAfter uint32) {
	buf := C.malloc(256)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 124)) = initial
	ret = int(C.sub_4BDB20(C.int(uintptr(buf))))
	flagsAfter = *(*uint32)(unsafe.Pointer(uintptr(buf) + 124))
	return
}

// C_sub_4BDB30 clears bit 0x10 in the uint32 flags at offset +124.
func C_sub_4BDB30(initial uint32) (ret int, flagsAfter uint32) {
	buf := C.malloc(256)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 124)) = initial
	ret = int(C.sub_4BDB30(C.int(uintptr(buf))))
	flagsAfter = *(*uint32)(unsafe.Pointer(uintptr(buf) + 124))
	return
}

func C_nox_xxx_updDrawMonsterGen_4BC920() int   { return int(C.nox_xxx_updDrawMonsterGen_4BC920()) }
func C_sub_4BDDA0() int                         { return int(C.sub_4BDDA0()) }
func C_sub_4BE320() int                         { return int(C.sub_4BE320()) }
func C_sub_4C24A0() int                         { return int(C.sub_4C24A0()) }
func C_sub_4C2BD0() int                         { return int(C.sub_4C2BD0()) }
func C_sub_4C2BE0() int                         { return int(C.sub_4C2BE0()) }
func C_nox_xxx_updDrawUndeadKiller_4CCCF0() int { return int(C.nox_xxx_updDrawUndeadKiller_4CCCF0()) }
