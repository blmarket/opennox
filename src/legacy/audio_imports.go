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
int sub_4BDA80(struct312* a1);
void nox_common_list_clear_425760(nox_list_item_t* list);
void* sub_486320(void* a1, int a2);
char* nox_xxx_getSndName_40AF80(int a1);
void nox_common_list_remove_425920(void* a1);
int sub_4862E0(void* a3, int a4);
int nox_common_randomIntMinMax_415FF0(int min, int max, char* file, int line);
int sub_4863B0(void* a2);
void sub_4BD3C0(void* lpMem);

int sub_452770(uint32_t* a1);
int sub_4526D0(int a1);
int sub_4526F0(int a1);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/audio"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

var (
	AudioModule  *audio.AudioModule
	AudioExterns *audio.AudioExterns
)

func initAudio(externs *audio.AudioExterns) {
	var (
		dword_587000_126996  *uint32           = (*uint32)(&C.dword_587000_126996)
		dword_5d4594_1045420 *uint32           = (*uint32)(&C.dword_5d4594_1045420)
		dword_5d4594_1045424 *uint32           = (*uint32)(&C.dword_5d4594_1045424)
		dword_5d4594_1045428 **audio.Struct264 = (**audio.Struct264)(unsafe.Pointer(&C.dword_5d4594_1045428))
		dword_5d4594_1045432 *uint32           = (*uint32)(&C.dword_5d4594_1045432)
		dword_5d4594_1045436 *uint32           = (*uint32)(&C.dword_5d4594_1045436)
		dword_587000_127004  unsafe.Pointer    = unsafe.Pointer(&C.dword_587000_127004)
	)
	AudioModule = audio.NewAudioModule(
		externs,
		"audio",
		PlatformTicks,
		unsafe.Pointer(C.sub_452770),
		unsafe.Pointer(C.sub_4526F0),
		unsafe.Pointer(C.sub_4526D0),
		dword_587000_126996,
		dword_5d4594_1045420,
		dword_5d4594_1045424,
		dword_5d4594_1045428,
		dword_5d4594_1045432,
		dword_5d4594_1045436,
		dword_587000_127004,
		memmap.PtrT[timer.TimerGroup](0x5D4594, 1045228),
		memmap.PtrT[[6][10]audio.ListHead[audio.Struct200Field28, *audio.Struct200Field28]](0x5D4594, 839892),
		memmap.PtrT[audio.ListHead[audio.Struct576, *audio.Struct576]](0x5D4594, 840612),
		sub_425770,
		nox_common_list_append_4258E0,
		sub_4BDB30,
		sub_4BD300,
		func(a1 unsafe.Pointer, a2 unsafe.Pointer) {
			Struct312Module.Sub_4BDB90((*uint32)(a1), (*uint32)(a2))
		},
		sub_4BDB40,
		sub_4864A0,
		sub_486350,
		sub_486520,
		sub_4BD280,
		nox_common_list_getFirstSafe_425890,
		sub_4BD340,
		sub_4BD2E0,
		sub_4BD470,
		func(a1 int, a2 byte) unsafe.Pointer {
			return unsafe.Pointer(C.sub_452810(C.int(a1), C.char(a2)))
		},
		sub_4BD2D0,
		func(a1 int) int {
			return int(sub_4BDA80((*C.struct312)(unsafe.Pointer(uintptr(a1)))))
		},
		nox_common_list_clear_425760,
		sub_486320,
		func(id int) unsafe.Pointer {
			return unsafe.Pointer(nox_xxx_getSndName_40AF80(id))
		},
		nox_common_list_remove_425920,
		func(a1 unsafe.Pointer, a2 int) int {
			return bool2int((*timer.Timer)(a1).Init(int32(a2)))
		},
		func(min, max int, file unsafe.Pointer, line int) int {
			return nox_common_randomIntMinMax_415FF0(min, max, (*C.char)(file), line)
		},
		sub_4863B0,
		sub_4BD3C0,
	)

	initStruct88(externs)
	initPhase6(externs)
}

func sub_425770(a1 unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_425770((unsafe.Pointer)(a1)))
}

func nox_common_list_append_4258E0(list unsafe.Pointer, cur unsafe.Pointer) {
	C.nox_common_list_append_4258E0((*C.nox_list_item_t)(list), (*C.nox_list_item_t)(cur))
}

func sub_4BDB30(a1 int) int {
	return int(Phase6Module.Sub_4BDB30(int32(a1)))
}

func sub_4BD300(a1 unsafe.Pointer, a2 int) int {
	return int(C.sub_4BD300((*C.uint32_t)(a1), C.int(a2)))
}

func sub_4BDB40(a2 int) int {
	return int(Phase6Module.Sub_4BDB40(int32(a2)))
}

func sub_4BD280(a1 int, a2 int) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4BD280(C.int(a1), C.int(a2)))
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

func sub_4BD2D0(lpMem unsafe.Pointer) {
	C.sub_4BD2D0((unsafe.Pointer)(lpMem))
}

func sub_4BD3C0(lpMem unsafe.Pointer) {
	C.sub_4BD3C0((unsafe.Pointer)(lpMem))
}

//export sub_452770
func sub_452770(a1 *C.uint32_t) C.int {
	return C.int(AudioModule.Sub_452770((*uint32)(a1)))
}

//export sub_4526F0
func sub_4526F0(a1 C.int) C.int {
	return C.int(AudioModule.Sub_4526F0(int32(a1)))
}

//export sub_4526D0
func sub_4526D0(a1 C.int) C.int {
	return C.int(AudioModule.Sub_4526D0(int32(a1)))
}

//export nox_xxx_clientPlaySoundSpecial_452D80
func nox_xxx_clientPlaySoundSpecial_452D80(a1, a2 C.int) {
	AudioModule.Nox_xxx_clientPlaySoundSpecial_452D80(int32(a1), int32(a2))
}

//export sub_4519C0
func sub_4519C0() {
	AudioModule.Sub_4519C0()
}

//export sub_4523D0
func sub_4523D0(a1 *C.struct576) int32 {
	return AudioModule.Sub_4523D0((*audio.Struct576)(unsafe.Pointer(a1)))
}

//export sub_451970
func sub_451970() {
	AudioModule.Sub_451970()
}

//export sub_452FE0
func sub_452FE0(a1 *C.struct576, a2 int32) int32 {
	return AudioModule.Sub_452FE0((*audio.Struct576)(unsafe.Pointer(a1)), a2)
}

//export sub_452F50
func sub_452F50(a1 *C.struct576, a2 int32) int32 {
	return AudioModule.Sub_452F50((*audio.Struct576)(unsafe.Pointer(a1)), a2)
}

//export nox_xxx_draw_452300
func nox_xxx_draw_452300(a1 *C.struct200) *uint32 {
	return AudioModule.Nox_xxx_draw_452300((*audio.Struct200)(unsafe.Pointer(a1)))
}

//export nox_xxx_draw_452270
func nox_xxx_draw_452270(a1 int32) *C.char {
	return (*C.char)(unsafe.Pointer(AudioModule.Nox_xxx_draw_452270(a1)))
}

//export sub_452EE0
func sub_452EE0(a1 *C.struct576, a2 int32) int32 {
	return AudioModule.Sub_452EE0((*audio.Struct576)(unsafe.Pointer(a1)), a2)
}

//export sub_452F80
func sub_452F80(a1 *C.struct576, a2 int32) *uint32 {
	return AudioModule.Sub_452F80((*audio.Struct576)(unsafe.Pointer(a1)), a2)
}

//export sub_452E90
func sub_452E90(a1 *uint32, a2 *C.struct576) int32 {
	return AudioModule.Sub_452E90(a1, (*audio.Struct576)(unsafe.Pointer(a2)))
}

//export sub_452DC0
func sub_452DC0(a1 int32, a2 int32, a3 int32) {
	AudioModule.Sub_452DC0(a1, a2, a3)
}

//export sub_452E10
func sub_452E10(a1 int32, a2 int32, a3 int32) {
	AudioModule.Sub_452E10(a1, a2, a3)
}
