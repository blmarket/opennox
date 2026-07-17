package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_50E7A0(uint32_t* a1, int a2);
int nox_xxx_servSendShopItems_50F280(int a1, int a2);
uint32_t* nox_xxx_tradeSetPlayer_50F370(uint32_t* a1, int a2);
int sub_50FD60(uint32_t* a1, int a2);
void sub_510320(int a1, int a2);
int sub_510540(int a1);
int sub_5105D0(int a1);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func C_sub_50E7A0Missing() (empty, missing int) {
	shop := C.calloc(1, 64)
	node := C.calloc(1, 16)
	defer C.free(shop)
	defer C.free(node)

	empty = int(C.sub_50E7A0((*C.uint32_t)(shop), 10))
	*(*uint32)(node) = 10
	*(*uint32)(unsafe.Pointer(uintptr(shop) + 5*4)) = uint32(uintptr(node))
	missing = int(C.sub_50E7A0((*C.uint32_t)(shop), 20))
	return
}

func C_nox_xxx_servSendShopItemsEmpty() bool {
	shop := C.calloc(1, 64)
	defer C.free(shop)
	return int(C.nox_xxx_servSendShopItems_50F280(0, C.int(uintptr(shop)))) == int(uintptr(shop))
}

func C_nox_xxx_tradeSetPlayerNonPlayer() (same bool, state uint32) {
	session := C.calloc(1, 64)
	object := C.calloc(1, 16)
	defer C.free(session)
	defer C.free(object)

	ret := C.nox_xxx_tradeSetPlayer_50F370((*C.uint32_t)(session), C.int(uintptr(object)))
	return unsafe.Pointer(ret) == session, *(*uint32)(session)
}

func tradeTestSetObjectID(ptr unsafe.Pointer, id uint16) {
	*(*uint16)(unsafe.Pointer(uintptr(ptr) + 4)) = id
}

func C_sub_50FD60Capacity() (empty, same, overflow int) {
	empty = int(C.sub_50FD60(nil, 0))

	oneNode := C.calloc(1, 16)
	oneObject := C.calloc(1, 16)
	target := C.calloc(1, 16)
	defer C.free(oneNode)
	defer C.free(oneObject)
	defer C.free(target)
	tradeTestSetObjectID(oneObject, 1)
	tradeTestSetObjectID(target, 1)
	*(*uint32)(oneNode) = uint32(uintptr(oneObject))
	same = int(C.sub_50FD60((*C.uint32_t)(oneNode), C.int(uintptr(target))))

	nodes := C.calloc(4, 16)
	objects := C.calloc(4, 16)
	defer C.free(nodes)
	defer C.free(objects)
	for i := uintptr(0); i < 4; i++ {
		node := unsafe.Pointer(uintptr(nodes) + 16*i)
		object := unsafe.Pointer(uintptr(objects) + 16*i)
		tradeTestSetObjectID(object, uint16(i+1))
		*(*uint32)(node) = uint32(uintptr(object))
		if i+1 < 4 {
			*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = uint32(uintptr(nodes) + 16*(i+1))
		}
	}
	tradeTestSetObjectID(target, 99)
	overflow = int(C.sub_50FD60((*C.uint32_t)(nodes), C.int(uintptr(target))))
	return
}

func C_sub_510320Guards() {
	C.sub_510320(0, 0)
	C.sub_510320(1, 0)
	C.sub_510320(0, 1)
}

func C_sub_510540NonQuest() int {
	return int(C.sub_510540(0))
}

func C_sub_5105D0Cached() [4]int {
	const base = uintptr(0x5D4594)
	slots := [3]*uint32{
		memmap.PtrUint32(base, 2386536),
		memmap.PtrUint32(base, 2386540),
		memmap.PtrUint32(base, 2386544),
	}
	saved := [3]uint32{*slots[0], *slots[1], *slots[2]}
	*slots[0], *slots[1], *slots[2] = 11, 12, 13
	defer func() {
		*slots[0], *slots[1], *slots[2] = saved[0], saved[1], saved[2]
	}()

	object := C.calloc(1, 16)
	defer C.free(object)
	var out [4]int
	for i, id := range []uint16{11, 12, 13, 99} {
		tradeTestSetObjectID(object, id)
		out[i] = int(C.sub_5105D0(C.int(uintptr(object))))
	}
	return out
}
