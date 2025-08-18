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
extern void* dword_587000_155144;
void* sub_425770(void* a1);
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
nox_list_item_t* nox_common_list_getNext_425940(nox_list_item_t* list);
int sub_4BDB30(int a1);
int sub_4BD300(uint32_t* a1, int a2);
int sub_4BDB40(int a2);
int sub_486350(void* a1, int a2);
int sub_486520(void* a2);
uint32_t* sub_4BD280(int a1, int a2);
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
uint32_t* sub_4BD340(int a1, int a2, int a3, int a4);
uint32_t* sub_4BD2E0(uint32_t** a1);
uint32_t* sub_4BD470(uint32_t** a1, int a2);
int* sub_452810(int a1, char a2);
int sub_4BD940(struct312* a1);
int sub_4BD9B0(struct312* a2);
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

int sub_452770(struct312* a1);
int sub_4526D0(int a1);
int sub_4526F0(struct312* a1);
int sub_4873C0(struct264* a3);
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
nox_list_item_t* nox_common_list_getNextSafe_4258A0(nox_list_item_t* list);
int sub_4BDA80(struct312* a1);
void* sub_486320(void* a1, int a2);
void* sub_425770(void* a1);
int* sub_4870E0(int* a1);

void* sub_4864A0(void* a3);
nox_list_item_t* nox_common_list_getNext_425940(nox_list_item_t* list);
void* sub_425770(void* a1);
int sub_4BD8C0(struct312* a1);
int sub_4BD940(struct312* a1);
int sub_4BD9B0(struct312* a2);

int sub_425960(int a1);
int sub_487D60(int a1);
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/audio"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

var (
	AudioModule *audio.AudioModule
)

func initExterns() *audio.AudioExterns {
	return &audio.AudioExterns{
		Dword_5d4594_1193336: 0,

		Dword_587000_126996:          (*uint32)(&C.dword_587000_126996),
		Dword_5d4594_1045420:         (*uint32)(&C.dword_5d4594_1045420),
		Dword_5d4594_1045424:         (*uint32)(&C.dword_5d4594_1045424),
		Dword_5d4594_1045428:         (**audio.Struct264)(unsafe.Pointer(&C.dword_5d4594_1045428)),
		Dword_5d4594_1045432:         (*uint32)(&C.dword_5d4594_1045432),
		Dword_5d4594_1045436:         (*uint32)(&C.dword_5d4594_1045436),
		Ptr_TimerGroup_587000_127004: (**timer.TimerGroup)(unsafe.Pointer(&C.dword_587000_127004)),
		Dword_587000_155144:          (**audio.Struct587000_155144)(unsafe.Pointer(&C.dword_587000_155144)),
		Dword_587000_127004:          (**timer.TimerGroup)(unsafe.Pointer(&C.dword_587000_127004)),
		Dword_5d4594_805984:          (**audio.Struct264)(unsafe.Pointer(&C.dword_5d4594_805984)),

		Ptr_TimerGroup_5d4594_1193340: memmap.PtrT[*timer.TimerGroup](0x5D4594, 1193340),
		TimerGroup_5d4594_1045228:     memmap.PtrT[timer.TimerGroup](0x5D4594, 1045228),
		ListHeads_5d4594_839892:       memmap.PtrT[[6][10]audio.ListElement[audio.Struct200Field28, *audio.Struct200Field28]](0x5D4594, 839892),
		ListHead_5d4594_840612:        memmap.PtrT[audio.ListElement[audio.Struct576, *audio.Struct576]](0x5D4594, 840612),
		Struct200Arr_5d4594_840628:    memmap.PtrT[[1023]audio.Struct200](0x5D4594, 840628),
		Ptr_uint32_5d4594_1045444:     memmap.PtrUint32(0x5D4594, 1045444),

		Sub_4873C0_ptr: unsafe.Pointer(C.sub_4873C0),
		Sub_4BD8C0_ptr: unsafe.Pointer(C.sub_4BD8C0),
		Sub_4BD940_ptr: unsafe.Pointer(C.sub_4BD940),
		Sub_4BD9B0_ptr: unsafe.Pointer(C.sub_4BD9B0),
		Sub_452770_ptr: unsafe.Pointer(C.sub_452770),
		Sub_4526F0_ptr: unsafe.Pointer(C.sub_4526F0),
		Sub_4526D0_ptr: unsafe.Pointer(C.sub_4526D0),
	}
}

func initAudio(externs *audio.AudioExterns) {
	AudioModule = audio.NewAudioModule(
		externs,
		"audio",
		PlatformTicks,
		func(a1 unsafe.Pointer, a2 int) int {
			return int(C.sub_4BD300((*C.uint32_t)(a1), C.int(a2)))
		},
		func(a1 int, a2 int) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD280(C.int(a1), C.int(a2)))
		},
		func(a1 int, a2 int, a3 int, a4 int) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD340(C.int(a1), C.int(a2), C.int(a3), C.int(a4)))
		},
		func(a1 unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD2E0((**C.uint32_t)(a1)))
		},
		func(a1 unsafe.Pointer, a2 int) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD470((**C.uint32_t)(a1), C.int(a2)))
		},
		func(lpMem unsafe.Pointer) {
			C.sub_4BD2D0(lpMem)
		},
		func(id int) unsafe.Pointer {
			return unsafe.Pointer(nox_xxx_getSndName_40AF80(id))
		},
		func(min, max int, file unsafe.Pointer, line int) int {
			return nox_common_randomIntMinMax_415FF0(min, max, (*C.char)(file), line)
		},
		func(lpMem unsafe.Pointer) {
			C.sub_4BD3C0(lpMem)
		},
	)
}

//export sub_452770
func sub_452770(a1 *C.struct312) C.int {
	return C.int(AudioModule.Sub_452770((*audio.Struct312)(unsafe.Pointer(a1))))
}

//export sub_4526F0
func sub_4526F0(a1 *C.struct312) C.int {
	return C.int(AudioModule.Sub_4526F0((*audio.Struct312)(unsafe.Pointer(a1))))
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
func nox_xxx_draw_452300(a1 *C.struct200) *C.struct576 {
	return (*C.struct576)(unsafe.Pointer(AudioModule.Nox_xxx_draw_452300((*audio.Struct200)(unsafe.Pointer(a1)))))
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
func sub_452F80(a1 *C.struct576, a2 int32) {
	AudioModule.Sub_452F80((*audio.Struct576)(unsafe.Pointer(a1)), a2)
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

//export sub_4BD7A0
func sub_4BD7A0(lpMem unsafe.Pointer) {
	AudioModule.Sub_4BD7A0((*audio.Struct312)(lpMem))
}

//export sub_4BD8C0
func sub_4BD8C0(a1 *C.struct312) int32 {
	return AudioModule.Sub_4BD8C0((*audio.Struct312)(unsafe.Pointer(a1)))
}

//export sub_4BD940
func sub_4BD940(a1 *C.struct312) int32 {
	return AudioModule.Sub_4BD940((*audio.Struct312)(unsafe.Pointer(a1)))
}

//export sub_4BD9B0
func sub_4BD9B0(a2 *C.struct312) int32 {
	return AudioModule.Sub_4BD9B0((*audio.Struct312)(unsafe.Pointer(a2)))
}

//export sub_4873C0
func sub_4873C0(a3 *C.struct264) int32 {
	return AudioModule.Sub_4873C0((*audio.Struct264)(unsafe.Pointer(a3)))
}

//export sub_4BDA60
func sub_4BDA60(lpMem_ *C.struct312) {
	AudioModule.Sub_4BDA60((*audio.Struct312)(unsafe.Pointer(lpMem_)))
}

//export sub_4BDA80
func sub_4BDA80(a1_ *C.struct312) int32 {
	return AudioModule.Sub_4BDA80((*audio.Struct312)(unsafe.Pointer(a1_)))
}

//export sub_425960
func sub_425960(a1 C.int) C.int {
	return C.int(AudioModule.Sub_425960(int32(a1)))
}

//export sub_487D60
func sub_487D60(a1 C.int) C.int {
	return C.int(AudioModule.Sub_487D60(int32(a1)))
}
