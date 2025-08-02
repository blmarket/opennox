package legacy

/*
#include "defs.h"

void* sub_4864A0(void* a3);
nox_list_item_t* nox_common_list_getNext_425940(nox_list_item_t* list);
void* sub_425770(void* a1);
int sub_4BD8C0(int a1);
int sub_4BD940(int a1);
int sub_4BD9B0(uint32_t* a2);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

var (
	Struct312Module *audio.Struct312Module
)

func initStruct312() {
	Struct312Module = audio.NewStruct312Module(
		"struct312",
		sub_4864A0,
		nox_common_list_getNext_425940,
		func(a1 *audio.Nox_list_item_t) unsafe.Pointer {
			return unsafe.Pointer(C.sub_425770((unsafe.Pointer)(a1)))
		},
		unsafe.Pointer(C.sub_4BD8C0),
		unsafe.Pointer(C.sub_4BD940),
		unsafe.Pointer(C.sub_4BD9B0),
	)
}

func nox_common_list_getNext_425940(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getNext_425940((*C.nox_list_item_t)(list)))
}

//export sub_4BDB90
func sub_4BDB90(a1 *uint32, a2 *uint32) {
	Struct312Module.Sub_4BDB90(a1, a2)
}

//export sub_4BD720
func sub_4BD720(a1 int32) *uint32 {
	return Struct312Module.Sub_4BD720(a1)
}

//export sub_4BD7A0
func sub_4BD7A0(lpMem unsafe.Pointer) {
	Struct312Module.Sub_4BD7A0(lpMem)
}

//export sub_4BD8C0
func sub_4BD8C0(a1 int32) int32 {
	return Struct312Module.Sub_4BD8C0(a1)
}

//export sub_4BD940
func sub_4BD940(a1 int32) int32 {
	return Struct312Module.Sub_4BD940(a1)
}

//export sub_4BD9B0
func sub_4BD9B0(a2 *uint32) int32 {
	return Struct312Module.Sub_4BD9B0(a2)
}
