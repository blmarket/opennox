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
void sub_4BDA60(struct312* lpMem);
void nox_common_list_clear_425760(nox_list_item_t* list);
void* sub_4864A0(void* a3);
nox_list_item_t* nox_common_list_getNextSafe_4258A0(nox_list_item_t* list);
int sub_4BDA80(struct312* a1);
void* sub_486320(void* a1, int a2);
void* sub_425770(void* a1);
int* sub_4870E0(int* a1);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

func nox_common_list_getFirstSafe_425890(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getFirstSafe_425890((*C.nox_list_item_t)(list)))
}

func nox_common_list_remove_425920(a1 unsafe.Pointer) {
	C.nox_common_list_remove_425920((unsafe.Pointer)(a1))
}

func nullsub_10(a1 uint32) {
	C.nullsub_10(C.uint32_t(a1))
}

func nox_common_list_clear_425760(list unsafe.Pointer) {
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(list))
}

func nox_common_list_getNextSafe_4258A0(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getNextSafe_4258A0((*C.nox_list_item_t)(list)))
}

//export sub_452810
func sub_452810(a1 int32, a2 C.char) *int32 {
	return (*int32)(unsafe.Pointer(AudioModule.Sub_452810(a1, int8(a2))))
}

//export sub_431270
func sub_431270() {
	AudioModule.Sub_431270()
}

//export sub_487680
func sub_487680(lpMem_ *C.struct264) {
	AudioModule.Sub_487680((*audio.Struct264)(unsafe.Pointer(lpMem_)))
}

//export sub_431290
func sub_431290() {
	AudioModule.Sub_431290()
}
