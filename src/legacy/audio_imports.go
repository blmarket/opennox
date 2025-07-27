package legacy

/*
#include "defs.h"

extern uint32_t dword_587000_126996;
extern uint32_t dword_5d4594_1045420;
extern uint32_t dword_5d4594_1045424;
extern uint32_t dword_5d4594_1045428;
extern uint32_t dword_5d4594_1045432;
extern uint32_t dword_5d4594_1045436;
extern void* dword_587000_127004;
void* sub_425770(void* a1);
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
int sub_4BDB30(int a1);
int sub_4BD300(uint32_t* a1, int a2);
void sub_4BDB90(uint32_t* a1, uint32_t* a2);
int sub_4BDB40(int a2);
void* sub_4864A0(void* a3);
int sub_486350(void* a1, int a2);
int sub_486520(void* a2);
uint32_t* sub_4BD280(int a1, int a2);
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
uint32_t* sub_4BD340(int a1, int a2, int a3, int a4);
uint32_t* sub_4BD2E0(uint32_t** a1);
uint32_t* sub_4BD470(uint32_t** a1, int a2);
int* sub_452810(int a1, char a2);
void sub_4BD2D0(void* lpMem);
int sub_4BDA80(int a1);
void nox_common_list_clear_425760(nox_list_item_t* list);
void* sub_486320(void* a1, int a2);
char* nox_xxx_getSndName_40AF80(int a1);
void nox_common_list_remove_425920(void* a1);
int sub_4862E0(void* a3, int a4);
int nox_common_randomIntMinMax_415FF0(int min, int max, const char* file, int line);
int sub_4863B0(void* a2);
void sub_4BD3C0(void* lpMem);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

var (
	AudioModule *audio.AudioModule
)

func initAudio() {
	var (
		dword_587000_126996 *uint32 = (*uint32)(&C.dword_587000_126996)
		dword_5d4594_1045420 *uint32 = (*uint32)(&C.dword_5d4594_1045420)
		dword_5d4594_1045424 *uint32 = (*uint32)(&C.dword_5d4594_1045424)
		dword_5d4594_1045428 *uint32 = (*uint32)(&C.dword_5d4594_1045428)
		dword_5d4594_1045432 *uint32 = (*uint32)(&C.dword_5d4594_1045432)
		dword_5d4594_1045436 *uint32 = (*uint32)(&C.dword_5d4594_1045436)
		dword_587000_127004 unsafe.Pointer = (unsafe.Pointer)(unsafe.Pointer(&C.dword_587000_127004))
	)
	AudioModule = audio.NewAudioModule(
		"audio",
		dword_587000_126996,
		dword_5d4594_1045420,
		dword_5d4594_1045424,
		dword_5d4594_1045428,
		dword_5d4594_1045432,
		dword_5d4594_1045436,
		dword_587000_127004,
		sub_425770,
		nox_common_list_append_4258E0,
		sub_4BDB30,
		sub_4BD300,
		sub_4BDB90,
		sub_4BDB40,
		sub_4864A0,
		sub_486350,
		sub_486520,
		sub_4BD280,
		nox_common_list_getFirstSafe_425890,
		sub_4BD340,
		sub_4BD2E0,
		sub_4BD470,
		sub_452810,
		sub_4BD2D0,
		sub_4BDA80,
		nox_common_list_clear_425760,
		sub_486320,
		nox_xxx_getSndName_40AF80,
		nox_common_list_remove_425920,
		sub_4862E0,
		nox_common_randomIntMinMax_415FF0,
		sub_4863B0,
		sub_4BD3C0,
	)
}

func sub_425770(a1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_425770((unsafe.Pointer)(a1)))
}

func nox_common_list_append_4258E0(list unsafe.Pointer, cur unsafe.Pointer) {
	C.nox_common_list_append_4258E0((*C.nox_list_item_t)(list), (*C.nox_list_item_t)(cur))
}

func sub_4BDB30(a1 int) int {
	return int(C.sub_4BDB30(C.int(a1)))
}

func sub_4BD300(a1 unsafe.Pointer, a2 int) int {
	return int(C.sub_4BD300((*C.uint32_t)(a1), C.int(a2)))
}

func sub_4BDB90(a1 unsafe.Pointer, a2 unsafe.Pointer) {
	C.sub_4BDB90((*C.uint32_t)(a1), (*C.uint32_t)(a2))
}

func sub_4BDB40(a2 int) int {
	return int(C.sub_4BDB40(C.int(a2)))
}

func sub_4864A0(a3 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4864A0((unsafe.Pointer)(a3)))
}

func sub_486350(a1 unsafe.Pointer, a2 int) int {
	return int(C.sub_486350((unsafe.Pointer)(a1), C.int(a2)))
}

func sub_486520(a2 unsafe.Pointer) int {
	return int(C.sub_486520((unsafe.Pointer)(a2)))
}

func sub_4BD280(a1 int, a2 int) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4BD280(C.int(a1), C.int(a2)))
}

func nox_common_list_getFirstSafe_425890(list unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.nox_common_list_getFirstSafe_425890((*C.nox_list_item_t)(list)))
}

func sub_4BD340(a1 int, a2 int, a3 int, a4 int) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4BD340(C.int(a1), C.int(a2), C.int(a3), C.int(a4)))
}

func sub_4BD2E0(a1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4BD2E0((**C.uint32_t)(a1)))
}

func sub_4BD470(a1 unsafe.Pointer, a2 int) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4BD470((**C.uint32_t)(a1), C.int(a2)))
}

func sub_452810(a1 int, a2 byte) unsafe.Pointer {
	return unsafe.Pointer(C.sub_452810(C.int(a1), C.char(a2)))
}

func sub_4BD2D0(lpMem unsafe.Pointer) {
	C.sub_4BD2D0((unsafe.Pointer)(lpMem))
}

func sub_4BDA80(a1 int) int {
	return int(C.sub_4BDA80(C.int(a1)))
}

func nox_common_list_clear_425760(list unsafe.Pointer) {
	C.nox_common_list_clear_425760((*C.nox_list_item_t)(list))
}

func sub_486320(a1 unsafe.Pointer, a2 int) unsafe.Pointer {
	return unsafe.Pointer(C.sub_486320((unsafe.Pointer)(a1), C.int(a2)))
}

func nox_xxx_getSndName_40AF80(a1 int) unsafe.Pointer {
	return unsafe.Pointer(C.nox_xxx_getSndName_40AF80(C.int(a1)))
}

func nox_common_list_remove_425920(a1 unsafe.Pointer) {
	C.nox_common_list_remove_425920((unsafe.Pointer)(a1))
}

func sub_4862E0(a3 unsafe.Pointer, a4 int) int {
	return int(C.sub_4862E0((unsafe.Pointer)(a3), C.int(a4)))
}

func nox_common_randomIntMinMax_415FF0(min int, max int, file unsafe.Pointer, line int) int {
	return int(C.nox_common_randomIntMinMax_415FF0(C.int(min), C.int(max), (*C.char)(file), C.int(line)))
}

func sub_4863B0(a2 unsafe.Pointer) int {
	return int(C.sub_4863B0((unsafe.Pointer)(a2)))
}

func sub_4BD3C0(lpMem unsafe.Pointer) {
	C.sub_4BD3C0((unsafe.Pointer)(lpMem))
}
