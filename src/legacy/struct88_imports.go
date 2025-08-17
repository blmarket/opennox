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

// TODO: No need to use C binding.
//
//export sub_4870A0
func sub_4870A0() {
	AudioModule.Sub_4870A0()
}

//export sub_4875F0
func sub_4875F0() int32 {
	return AudioModule.Sub_4875F0()
}

//export sub_486FE0
func sub_486FE0(a1 int) *C.struct88 {
	return (*C.struct88)(unsafe.Pointer(AudioModule.Sub_486FE0((*audio.Struct587000_94032)(unsafe.Pointer(uintptr(a1))))))
}

//export sub_487050
func sub_487050(a1 *C.struct88) {
	AudioModule.Sub_487050((*audio.Struct88)(unsafe.Pointer(a1)))
}
