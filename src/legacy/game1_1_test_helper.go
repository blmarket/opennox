package legacy

/*
#include <stdint.h>
#include <stdlib.h>

short nox_float2int16(float a1);
short nox_float2int16_abs(float a1);
float nox_double2float(double a1);
int nox_double2int(double a1);

typedef struct nox_list_item_t nox_list_item_t;
void nox_common_list_clear_425760(nox_list_item_t* list);
void* sub_425770(void* a1);
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
nox_list_item_t* nox_common_list_getNextSafe_4258A0(nox_list_item_t* list);
int* sub_425A60(int* a1);
int* sub_425BE0(int* a1);
int* sub_425BC0(int a1);
int sub_425960(int a1);
void nox_common_list_remove_425920(void* a1);
*/
import "C"
import "unsafe"

// Pure numeric-conversion helpers from GAME1_1.c.

func C_nox_float2int16(a1 float32) int16     { return int16(C.nox_float2int16(C.float(a1))) }
func C_nox_float2int16_abs(a1 float32) int16 { return int16(C.nox_float2int16_abs(C.float(a1))) }
func C_nox_double2float(a1 float64) float32  { return float32(C.nox_double2float(C.double(a1))) }
func C_nox_double2int(a1 float64) int        { return int(C.nox_double2int(C.double(a1))) }

// List helpers from GAME1_1.c — pure on caller-supplied buffers via raw offsets.

// C_nox_common_list_clear_425760 allocates a list header, clears it, and returns field values.
func C_nox_common_list_clear_425760() (base, f0, f1, f2 uint32) {
	buf := C.malloc(12)
	defer C.free(buf)
	base = uint32(uintptr(buf))
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(buf))
	f0 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 0))
	f1 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 4))
	f2 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 8))
	return
}

// C_sub_425770 initializes a node-like struct and returns field values and the returned pointer.
func C_sub_425770() (base, f0, f1, f2 uint32, ret uint32) {
	buf := C.malloc(12)
	defer C.free(buf)
	base = uint32(uintptr(buf))
	r := C.sub_425770(buf)
	ret = uint32(uintptr(r))
	f0 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 0))
	f1 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 4))
	f2 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 8))
	return
}

// C_nox_common_list_getNextSafe_4258A0_nil tests nil input.
func C_nox_common_list_getNextSafe_4258A0_nil() int {
	return int(uintptr(unsafe.Pointer(C.nox_common_list_getNextSafe_4258A0(nil))))
}

// C_nox_common_list_getNextSafe_4258A0_empty sets up an empty cleared list; expects 0.
func C_nox_common_list_getNextSafe_4258A0_empty() (base, ret int) {
	buf := C.malloc(12)
	defer C.free(buf)
	base = int(uintptr(buf))
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(buf))
	r := C.nox_common_list_getNextSafe_4258A0((*C.nox_list_item_t)(buf))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}

// C_nox_common_list_getNextSafe_4258A0_next sets up list->field0=node where node.field2 != node; expects node address.
func C_nox_common_list_getNextSafe_4258A0_next(nodeF2Self bool) (base, nodeAddr, ret int) {
	list := C.malloc(12)
	node := C.malloc(12)
	defer C.free(list)
	defer C.free(node)
	base = int(uintptr(list))
	nodeAddr = int(uintptr(node))
	*(*uint32)(unsafe.Pointer(uintptr(list) + 0)) = uint32(nodeAddr)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 4)) = uint32(base) // prev not used
	*(*uint32)(unsafe.Pointer(uintptr(list) + 8)) = uint32(base)
	if nodeF2Self {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = uint32(nodeAddr)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(node) + 0)) = 0
	*(*uint32)(unsafe.Pointer(uintptr(node) + 4)) = 0
	r := C.nox_common_list_getNextSafe_4258A0((*C.nox_list_item_t)(list))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}

// C_nox_common_list_getFirstSafe_425890 mirrors getNextSafe behavior.
func C_nox_common_list_getFirstSafe_425890_nil() int {
	return int(uintptr(unsafe.Pointer(C.nox_common_list_getFirstSafe_425890(nil))))
}
func C_nox_common_list_getFirstSafe_425890_empty() (base, ret int) {
	buf := C.malloc(12)
	defer C.free(buf)
	base = int(uintptr(buf))
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(buf))
	r := C.nox_common_list_getFirstSafe_425890((*C.nox_list_item_t)(buf))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}
func C_nox_common_list_getFirstSafe_425890_next(nodeF2Self bool) (base, nodeAddr, ret int) {
	list := C.malloc(12)
	node := C.malloc(12)
	defer C.free(list)
	defer C.free(node)
	base = int(uintptr(list))
	nodeAddr = int(uintptr(node))
	*(*uint32)(unsafe.Pointer(uintptr(list) + 0)) = uint32(nodeAddr)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 4)) = uint32(base)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 8)) = uint32(base)
	if nodeF2Self {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = uint32(nodeAddr)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = 0
	}
	r := C.nox_common_list_getFirstSafe_425890((*C.nox_list_item_t)(list))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}

// C_sub_425A60 is a thin wrapper around getNextSafe with int* signature.
func C_sub_425A60_nil() int {
	return int(uintptr(unsafe.Pointer(C.sub_425A60(nil))))
}
func C_sub_425A60_next(nodeF2Self bool) (base, nodeAddr, ret int) {
	list := C.malloc(12)
	node := C.malloc(12)
	defer C.free(list)
	defer C.free(node)
	base = int(uintptr(list))
	nodeAddr = int(uintptr(node))
	*(*uint32)(unsafe.Pointer(uintptr(list) + 0)) = uint32(nodeAddr)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 4)) = uint32(base)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 8)) = uint32(base)
	if nodeF2Self {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = uint32(nodeAddr)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = 0
	}
	r := C.sub_425A60((*C.int)(list))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}

// C_sub_425BE0 identical to sub_425A60.
func C_sub_425BE0_nil() int {
	return int(uintptr(unsafe.Pointer(C.sub_425BE0(nil))))
}
func C_sub_425BE0_next(nodeF2Self bool) (base, nodeAddr, ret int) {
	list := C.malloc(12)
	node := C.malloc(12)
	defer C.free(list)
	defer C.free(node)
	base = int(uintptr(list))
	nodeAddr = int(uintptr(node))
	*(*uint32)(unsafe.Pointer(uintptr(list) + 0)) = uint32(nodeAddr)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 4)) = uint32(base)
	*(*uint32)(unsafe.Pointer(uintptr(list) + 8)) = uint32(base)
	if nodeF2Self {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = uint32(nodeAddr)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = 0
	}
	r := C.sub_425BE0((*C.int)(list))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}

// C_sub_425BC0 expects a struct where list header is at offset +40.
func C_sub_425BC0_empty() (base, ret int) {
	buf := C.malloc(64)
	defer C.free(buf)
	base = int(uintptr(buf))
	listPtr := unsafe.Pointer(uintptr(buf) + 40)
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(listPtr))
	r := C.sub_425BC0(C.int(base))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}
func C_sub_425BC0_next(nodeF2Self bool) (base, nodeAddr, ret int) {
	buf := C.malloc(64)
	node := C.malloc(12)
	defer C.free(buf)
	defer C.free(node)
	base = int(uintptr(buf))
	nodeAddr = int(uintptr(node))
	listPtr := uintptr(buf) + 40
	*(*uint32)(unsafe.Pointer(listPtr + 0)) = uint32(nodeAddr)
	*(*uint32)(unsafe.Pointer(listPtr + 4)) = uint32(listPtr)
	*(*uint32)(unsafe.Pointer(listPtr + 8)) = uint32(listPtr)
	if nodeF2Self {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = uint32(nodeAddr)
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(node) + 8)) = 0
	}
	r := C.sub_425BC0(C.int(base))
	ret = int(uintptr(unsafe.Pointer(r)))
	return
}

// C_sub_425960 reads *(a1+4) as pointer p, then checks *(p+8) != p.
func C_sub_425960(equal bool) (ret int) {
	a1 := C.malloc(16)
	p := C.malloc(16)
	defer C.free(a1)
	defer C.free(p)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 4)) = uint32(uintptr(p))
	if equal {
		*(*uint32)(unsafe.Pointer(uintptr(p) + 8)) = uint32(uintptr(p))
	} else {
		*(*uint32)(unsafe.Pointer(uintptr(p) + 8)) = 0
	}
	ret = int(C.sub_425960(C.int(uintptr(a1))))
	return
}

// C_nox_common_list_remove_425920 removes a node from doubly linked list.
func C_nox_common_list_remove_425920() (node, prev, next, nodeF0, nodeF1, prevNext, nextPrev uint32) {
	prevBuf := C.malloc(12)
	nodeBuf := C.malloc(12)
	nextBuf := C.malloc(12)
	defer C.free(prevBuf)
	defer C.free(nextBuf)
	defer C.free(nodeBuf)
	prev = uint32(uintptr(prevBuf))
	node = uint32(uintptr(nodeBuf))
	next = uint32(uintptr(nextBuf))
	// prev <-> node <-> next circular initially
	*(*uint32)(unsafe.Pointer(uintptr(prevBuf) + 0)) = node
	*(*uint32)(unsafe.Pointer(uintptr(prevBuf) + 4)) = 0 // not used
	*(*uint32)(unsafe.Pointer(uintptr(nodeBuf) + 0)) = next
	*(*uint32)(unsafe.Pointer(uintptr(nodeBuf) + 4)) = prev
	*(*uint32)(unsafe.Pointer(uintptr(nextBuf) + 4)) = node
	*(*uint32)(unsafe.Pointer(uintptr(nextBuf) + 0)) = 0
	// set field_2 not needed for remove
	C.nox_common_list_remove_425920(nodeBuf)
	nodeF0 = *(*uint32)(unsafe.Pointer(uintptr(nodeBuf) + 0))
	nodeF1 = *(*uint32)(unsafe.Pointer(uintptr(nodeBuf) + 4))
	prevNext = *(*uint32)(unsafe.Pointer(uintptr(prevBuf) + 0))
	nextPrev = *(*uint32)(unsafe.Pointer(uintptr(nextBuf) + 4))
	return
}
