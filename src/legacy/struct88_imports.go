package legacy

/*
#include "defs.h"

extern void* dword_587000_155144;
void nox_common_list_remove_425920(void* a1);
void* sub_425770(void* a1);
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
nox_list_item_t* nox_common_list_getNextSafe_4258A0(nox_list_item_t* list);
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

var (
	Struct88Module *audio.Struct88Module
)

func initStruct88(externs *audio.AudioExterns) {
	var (
		dword_587000_155144 **audio.Struct587000_155144 = (**audio.Struct587000_155144)(unsafe.Pointer(&C.dword_587000_155144))
	)
	var _ **audio.Struct587000_155144 = dword_587000_155144
	Struct88Module = audio.NewStruct88Module(
		externs,
		"struct88",
		Struct264Module.Sub_487680,
		nox_common_list_getNextSafe_4258A0,
		nox_common_list_remove_425920,
		nox_common_list_getFirstSafe_425890,
		nox_common_list_append_4258E0,
	)
}

//export sub_4870A0
func sub_4870A0() {
	Struct88Module.Sub_4870A0()
}

//export sub_4875F0
func sub_4875F0() int32 {
	return Struct88Module.Sub_4875F0()
}

//export sub_486FE0
func sub_486FE0(a1 int) *C.struct88 {
	return (*C.struct88)(unsafe.Pointer(Struct88Module.Sub_486FE0(unsafe.Pointer(uintptr(a1)))))
}

//export sub_487050
func sub_487050(a1 *C.struct88) {
	Struct88Module.Sub_487050((*audio.Struct88)(unsafe.Pointer(a1)))
}
