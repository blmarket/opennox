package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_451CF0(uint32_t* a1);
int sub_451DC0(int a1);
int sub_451E80(int a1);
int sub_452580(uint32_t* a1);
int sub_452770(uint32_t* a1);
*/
import "C"

import "unsafe"

func putAudioEventU32(ptr unsafe.Pointer, off uintptr, val uint32) {
	*(*uint32)(unsafe.Pointer(uintptr(ptr) + off)) = val
}

// C_sub_451CF0Empty covers the exhausted-repeat path. The configured repeat
// limit is reached before the function needs to create another sound buffer.
func C_sub_451CF0Empty() (ret int, repeats uint32) {
	event := C.calloc(1, 576)
	config := C.calloc(1, 256)
	defer C.free(event)
	defer C.free(config)

	putAudioEventU32(event, 9*4, uint32(uintptr(config)))
	putAudioEventU32(config, 4, 1)
	putAudioEventU32(config, 60, 1)
	ret = int(C.sub_451CF0((*C.uint32_t)(event)))
	repeats = *(*uint32)(unsafe.Pointer(uintptr(event) + 109*4))
	return
}

// C_sub_451CF0Queued consumes one queued entry. sub_4BD710 is a pure pointer
// adjustment, so the sentinel is intentionally not a dereferenceable pointer.
func C_sub_451CF0Queued() (ret int, remaining uint32, index int32) {
	event := C.calloc(1, 576)
	config := C.calloc(1, 256)
	defer C.free(event)
	defer C.free(config)

	putAudioEventU32(event, 9*4, uint32(uintptr(config)))
	putAudioEventU32(event, 108*4, 1)
	putAudioEventU32(event, 43*4, ^uint32(0))
	putAudioEventU32(event, 10*4, 1000)
	ret = int(C.sub_451CF0((*C.uint32_t)(event)))
	remaining = *(*uint32)(unsafe.Pointer(uintptr(event) + 108*4))
	index = *(*int32)(unsafe.Pointer(uintptr(event) + 43*4))
	return
}

func C_sub_451DC0Active() int {
	event := C.calloc(1, 576)
	config := C.calloc(1, 256)
	defer C.free(event)
	defer C.free(config)

	putAudioEventU32(event, 36, uint32(uintptr(config)))
	putAudioEventU32(event, 168, 77)
	return int(C.sub_451DC0(C.int(uintptr(event))))
}

func C_sub_451E80Sequence() [3]int {
	event := C.calloc(1, 576)
	config := C.calloc(1, 256)
	defer C.free(event)
	defer C.free(config)

	putAudioEventU32(event, 36, uint32(uintptr(config)))
	putAudioEventU32(config, 192, 3)
	return [3]int{
		int(C.sub_451E80(C.int(uintptr(event)))),
		int(C.sub_451E80(C.int(uintptr(event)))),
		int(C.sub_451E80(C.int(uintptr(event)))),
	}
}

func C_sub_452580Empty() int {
	event := C.calloc(1, 576)
	config := C.calloc(1, 256)
	defer C.free(event)
	defer C.free(config)

	putAudioEventU32(event, 9*4, uint32(uintptr(config)))
	return int(C.sub_452580((*C.uint32_t)(event)))
}

func C_sub_452770NoBuffer() (ret int, attached uint32) {
	player := C.calloc(1, 384)
	event := C.calloc(1, 576)
	config := C.calloc(1, 256)
	defer C.free(player)
	defer C.free(event)
	defer C.free(config)

	putAudioEventU32(player, 38*4, uint32(uintptr(event)))
	putAudioEventU32(event, 9*4, uint32(uintptr(config)))
	ret = int(C.sub_452770((*C.uint32_t)(player)))
	attached = *(*uint32)(unsafe.Pointer(uintptr(player) + 72*4))
	return
}
