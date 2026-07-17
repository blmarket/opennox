package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_43AF40();
uint16_t* sub_42ADA0(int a1, int a2, short a3, unsigned int* a4);
uint16_t* sub_42B810(short* a1, unsigned int* a2);
extern uint32_t nox_game_createOrJoin_815048;
*/
import "C"
import "unsafe"

func C_sub_43AF40() int {
	return int(C.sub_43AF40())
}

func C_game1_2_setCreateOrJoin(v uint32) {
	C.nox_game_createOrJoin_815048 = C.uint32_t(v)
}

func C_game1_2_getCreateOrJoin() uint32 {
	return uint32(C.nox_game_createOrJoin_815048)
}

func C_game12SerializeRules() (lengths [4]uint32, headers [4]uint16) {
	rules := C.calloc(1, 640)
	defer C.free(rules)
	*(*int16)(unsafe.Pointer(uintptr(rules) + 6)) = 1
	copy((*[256]byte)(unsafe.Pointer(uintptr(rules) + 96))[:], "synthetic-map")
	copy((*[256]byte)(unsafe.Pointer(uintptr(rules) + 352))[:], "synthetic-game")

	name := C.CString("player-one")
	defer C.free(unsafe.Pointer(name))
	blob := C.calloc(1, 2)
	defer C.free(blob)
	*(*uint16)(blob) = 0x1234

	ptrs := []struct {
		offset uintptr
		value  unsafe.Pointer
	}{
		{608, C.calloc(1, 4)},
		{612, C.calloc(1, 4)},
		{616, C.calloc(1, 4)},
		{620, C.calloc(1, 1)},
		{624, C.calloc(1, 1)},
		{628, C.calloc(1, 4)},
		{632, C.calloc(1, 1)},
	}
	for _, p := range ptrs {
		defer C.free(p.value)
		*(*uint32)(unsafe.Pointer(uintptr(rules) + p.offset)) = uint32(uintptr(p.value))
	}
	*(*uint32)(ptrs[0].value) = uint32(uintptr(unsafe.Pointer(name)))
	*(*uint32)(ptrs[1].value) = 11
	*(*uint32)(ptrs[2].value) = 22
	*(*uint8)(ptrs[3].value) = 3
	*(*uint8)(ptrs[4].value) = 4
	*(*uint32)(ptrs[5].value) = 55
	*(*uint8)(ptrs[6].value) = 6
	*(*uint32)(unsafe.Pointer(uintptr(rules) + 636)) = uint32(uintptr(blob))

	for mode := 0; mode < 3; mode++ {
		var n C.uint
		out := C.sub_42ADA0(C.int(uintptr(rules)), C.int(mode), 1, &n)
		lengths[mode] = uint32(n)
		if out != nil {
			headers[mode] = uint16(*out)
			C.free(unsafe.Pointer(out))
		}
	}

	score := C.calloc(1, 580)
	defer C.free(score)
	*(*int16)(score) = 1
	copy((*[256]byte)(unsafe.Pointer(uintptr(score) + 24))[:], "score-map")
	copy((*[256]byte)(unsafe.Pointer(uintptr(score) + 280))[:], "score-game")
	scoreArrays := []struct {
		offset uintptr
		size   uintptr
	}{
		{536, 4}, {540, 4}, {544, 4}, {548, 4}, {552, 4},
		{556, 4}, {560, 4}, {564, 4}, {568, 4}, {572, 4}, {576, 1},
	}
	for i, p := range scoreArrays {
		mem := C.calloc(1, C.size_t(p.size))
		defer C.free(mem)
		if p.offset == 536 {
			*(*uint32)(mem) = uint32(uintptr(unsafe.Pointer(name)))
		} else if p.size == 1 {
			*(*uint8)(mem) = uint8(i)
		} else {
			*(*uint32)(mem) = uint32(100 + i)
		}
		*(*uint32)(unsafe.Pointer(uintptr(score) + p.offset)) = uint32(uintptr(mem))
	}
	var n C.uint
	out := C.sub_42B810((*C.short)(score), &n)
	lengths[3] = uint32(n)
	if out != nil {
		headers[3] = uint16(*out)
		C.free(unsafe.Pointer(out))
	}
	return lengths, headers
}
