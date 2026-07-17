package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

int sub_486640(void* a1, int a2);
int sub_4866D0(uint32_t* a1, int a2);
unsigned int sub_486A10(int a1, void* a2);
int sub_486AA0(uint32_t* a1, int a2, uint32_t* a3);
signed int sub_486DB0(int a1, char* a2, signed int a3);
void* sub_486E00(int a1);
int* sub_487100(int** a1);
int* sub_4877D0(int a1, int* a2);
int* sub_4877F0(int** a1);
int* sub_487810(int a1, int a2);
int sub_487910(int a1, int a2);
int* sub_487970(int a1, int a2);
void sub_487C30(uint32_t* a1);
int sub_487C80(int a1);
int sub_487D00(uint32_t* a1);
uint32_t* sub_487D30(uint32_t* a1, int a2, int a3);
int sub_487D60(int a1);
uint16_t* sub_480250(uint8_t* a1, uint16_t* a2);
int sub_487590(int a1, const void* a2);
void sub_487090(uint32_t** a1);
void sub_481410();
int sub_484450(int a1, int a2);
extern uint32_t nox_xxx_waypointCounterMB_587000_154948;
extern uint32_t dword_5d4594_3798804;
*/
import "C"
import "unsafe"

type game22SafeStateResult struct {
	emptySearch      uint32
	recordStride     uintptr
	normalFormat     int
	normalChannels   uint32
	normalMode       uint32
	specialFormat    int
	specialInput     uint32
	specialMode      uint32
	emptyRead        int
	closeEmpty       uintptr
	nextEmpty        uintptr
	firstEmpty       uintptr
	selectEmpty      uintptr
	removeEmpty      int
	releaseEmpty     uintptr
	listInitialized  bool
	bufferSize       int
	quarterSize      int
	configuredStride uint32
	configuredWidth  uint32
	clearedSize      uint32
}

func C_game22SafeStatePaths() (res game22SafeStateResult) {
	search := C.calloc(1, 8)
	key := C.CString("missing")
	defer C.free(search)
	defer C.free(unsafe.Pointer(key))
	res.emptySearch = uint32(C.sub_486A10(C.int(uintptr(search)), unsafe.Pointer(key)))

	table := C.calloc(1, 8)
	record := C.calloc(1, 36)
	output := C.calloc(7, 4)
	defer C.free(table)
	defer C.free(record)
	defer C.free(output)
	*(*unsafe.Pointer)(table) = record
	res.recordStride = uintptr(uint32(C.sub_4866D0((*C.uint32_t)(table), 0)))
	words := (*[9]uint32)(record)
	out := (*[7]uint32)(output)
	words[6], words[7], words[8] = 13, 5, 17
	res.normalFormat = int(C.sub_486AA0((*C.uint32_t)(table), 0, (*C.uint32_t)(output)))
	res.normalChannels, res.normalMode = out[3], out[4]
	words[7] = 8
	res.specialFormat = int(C.sub_486AA0((*C.uint32_t)(table), 0, (*C.uint32_t)(output)))
	res.specialInput, res.specialMode = out[1], out[4]

	reader := C.calloc(1, 288)
	dst := C.calloc(1, 8)
	defer C.free(reader)
	defer C.free(dst)
	res.emptyRead = int(C.sub_486DB0(C.int(uintptr(reader)), (*C.char)(dst), 8))
	res.closeEmpty = uintptr(unsafe.Pointer(C.sub_486E00(C.int(uintptr(reader)))))

	cursor := C.calloc(1, 4)
	defer C.free(cursor)
	res.nextEmpty = uintptr(unsafe.Pointer(C.sub_487100((**C.int)(cursor))))

	parent := C.calloc(1, 300)
	defer C.free(parent)
	C.sub_487C30((*C.uint32_t)(unsafe.Pointer(uintptr(parent) + 192)))
	res.firstEmpty = uintptr(unsafe.Pointer(C.sub_4877D0(C.int(uintptr(parent)), (*C.int)(cursor))))
	res.selectEmpty = uintptr(unsafe.Pointer(C.sub_487810(C.int(uintptr(parent)), -1)))
	res.removeEmpty = int(C.sub_487910(C.int(uintptr(parent)), -1))
	res.releaseEmpty = uintptr(unsafe.Pointer(C.sub_487970(C.int(uintptr(parent)), -1)))

	list := C.calloc(7, 4)
	defer C.free(list)
	listWords := (*[7]uint32)(list)
	for i := range listWords {
		listWords[i] = uint32(i + 1)
	}
	C.sub_487C30((*C.uint32_t)(list))
	res.listInitialized = listWords[0] == 0 && listWords[1] == 0 &&
		listWords[5] == 0 && listWords[6] == 0

	size := C.calloc(6, 4)
	defer C.free(size)
	sizeWords := (*[6]uint32)(size)
	sizeWords[2], sizeWords[3], sizeWords[4] = 2, 3, 4
	res.bufferSize = int(C.sub_487D00((*C.uint32_t)(size)))
	sizeWords[1] = 1
	res.quarterSize = int(C.sub_487D00((*C.uint32_t)(size)))
	C.sub_487D30((*C.uint32_t)(size), 7, 9)
	res.configuredStride, res.configuredWidth = sizeWords[3], sizeWords[4]
	C.sub_487D60(C.int(uintptr(size)))
	res.clearedSize = sizeWords[5]
	return res
}

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

type game22LightBufferResult struct {
	ret          int
	first        uint16
	second       uint32
	middle       uint32
	penultimate  uint16
	last         uint16
	untouchedGap uint16
}

func C_game22InitializeLightBuffer(value uint16) (out game22LightBufferResult) {
	const stride = 100
	const rows = 46
	oldStride := C.dword_5d4594_3798804
	C.dword_5d4594_3798804 = stride
	defer func() { C.dword_5d4594_3798804 = oldStride }()

	buf := C.calloc(1, stride*rows+64)
	defer C.free(buf)
	base := uintptr(buf)
	out.ret = int(C.sub_484450(C.int(value), C.int(base)))
	out.first = *(*uint16)(unsafe.Pointer(base + 46))
	out.second = *(*uint32)(unsafe.Pointer(base + stride + 44))
	out.middle = *(*uint32)(unsafe.Pointer(base + 24*stride + 44))
	out.penultimate = *(*uint16)(unsafe.Pointer(base + 43*stride + 48))
	out.last = *(*uint16)(unsafe.Pointer(base + 45*stride + 46))
	out.untouchedGap = *(*uint16)(unsafe.Pointer(base + 45*stride + 42))
	return out
}
