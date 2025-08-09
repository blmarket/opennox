package legacy

/*
#include "defs.h"

void* sub_4864A0(void* a3);
nox_list_item_t* nox_common_list_getNext_425940(nox_list_item_t* list);
void* sub_425770(void* a1);
int sub_4BD8C0(int a1);
int sub_4BD940(struct312* a1);
int sub_4BD9B0(struct312* a2);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

var (
	Struct312Module *audio.Struct312Module
)

func initStruct312(externs *audio.AudioExterns) {
	Struct312Module = audio.NewStruct312Module(externs, "struct312", nox_common_list_getNext_425940, func(a1 *audio.ListItem) unsafe.Pointer {
		return unsafe.Pointer(C.sub_425770((unsafe.Pointer)(a1)))
	})
}

func nox_common_list_getNext_425940(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getNext_425940((*C.nox_list_item_t)(list)))
}

//export sub_4BD7A0
func sub_4BD7A0(lpMem unsafe.Pointer) {
	AudioModule.Sub_4BD7A0((*audio.Struct312)(lpMem))
}

//export sub_4BD8C0
func sub_4BD8C0(a1 int32) int32 {
	return AudioModule.Sub_4BD8C0(a1)
}

//export sub_4BD940
func sub_4BD940(a1 *C.struct312) int32 {
	return AudioModule.Sub_4BD940((*audio.Struct312)(unsafe.Pointer(a1)))
}

//export sub_4BD9B0
func sub_4BD9B0(a2 *C.struct312) int32 {
	return AudioModule.Sub_4BD9B0((*audio.Struct312)(unsafe.Pointer(a2)))
}
