package legacy

/*
#include "defs.h"

extern uint32_t dword_5d4594_1045428;
extern void* dword_587000_127004;
extern void* dword_587000_155144;
extern void* dword_5d4594_805984;
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
int* sub_487100(int** a1);
void nox_common_list_remove_425920(void* a1);
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
void nullsub_10(uint32_t a1);
void sub_4BDA60(void* lpMem);
void nox_common_list_clear_425760(nox_list_item_t* list);
void* sub_4864A0(void* a3);
int sub_4873C0(int a3);
nox_list_item_t* nox_common_list_getNextSafe_4258A0(nox_list_item_t* list);
int sub_4BDA80(int a1);
void* sub_486320(void* a1, int a2);
void* sub_425770(void* a1);
int* sub_4870E0(int* a1);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

var (
	Struct264Module *audio.Struct264Module
)

func initStruct264() {
	var (
		sub_4873C0_ptr                       = unsafe.Pointer(C.sub_4873C0)
		dword_5d4594_1045428 *uint32         = (*uint32)(&C.dword_5d4594_1045428)
		dword_587000_127004  *unsafe.Pointer = (&C.dword_587000_127004)
		dword_587000_155144  *unsafe.Pointer = (&C.dword_587000_155144)
		dword_5d4594_805984  *unsafe.Pointer = (&C.dword_5d4594_805984)
	)
	Struct264Module = audio.NewStruct264Module(
		"struct264",
		sub_4873C0_ptr,
		dword_5d4594_1045428,
		dword_587000_127004,
		dword_587000_155144,
		dword_5d4594_805984,
		nox_common_list_getFirstSafe_425890,
		sub_487100,
		nox_common_list_remove_425920,
		nox_common_list_append_4258E0,
		nullsub_10,
		sub_4BDA60,
		nox_common_list_clear_425760,
		sub_4864A0,
		sub_4873C0,
		nox_common_list_getNextSafe_4258A0,
		sub_4BDA80,
		sub_486320,
		sub_425770,
		sub_4870E0,
		func(a1 *audio.Struct264) *audio.Struct312 {
			return Struct312Module.Sub_4BD720(a1)
		},
	)
}

func nox_common_list_getFirstSafe_425890(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getFirstSafe_425890((*C.nox_list_item_t)(list)))
}

func sub_487100(a1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_487100((**C.int)(a1)))
}

func nox_common_list_remove_425920(a1 unsafe.Pointer) {
	C.nox_common_list_remove_425920((unsafe.Pointer)(a1))
}

func nullsub_10(a1 uint32) {
	C.nullsub_10(C.uint32_t(a1))
}

func sub_4BDA60(lpMem unsafe.Pointer) {
	C.sub_4BDA60((unsafe.Pointer)(lpMem))
}

func nox_common_list_clear_425760(list unsafe.Pointer) {
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(list))
}

func sub_4873C0(a3 int) int {
	return int(C.sub_4873C0(C.int(a3)))
}

func nox_common_list_getNextSafe_4258A0(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getNextSafe_4258A0((*C.nox_list_item_t)(list)))
}

func sub_4BDA80(a1 int) int {
	return int(C.sub_4BDA80(C.int(a1)))
}

func sub_4870E0(a1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4870E0((*C.int)(a1)))
}

//export sub_452810
func sub_452810(a1 int32, a2 C.char) *int32 {
	return Struct264Module.Sub_452810(a1, int8(a2))
}

//export sub_431270
func sub_431270() {
	Struct264Module.Sub_431270()
}

//export sub_487680
func sub_487680(lpMem_ *C.struct264) {
	Struct264Module.Sub_487680((*audio.Struct264)(unsafe.Pointer(lpMem_)))
}

//export sub_487150
func sub_487150(a1 C.int, a2 unsafe.Pointer) *C.struct264 {
	return (*C.struct264)(unsafe.Pointer(Struct264Module.Sub_487150(int32(a1), a2)))
}

//export sub_431290
func sub_431290() {
	Struct264Module.Sub_431290()
}
