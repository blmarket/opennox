package legacy

/*
#include "defs.h"

extern void* dword_587000_122852;
int sub_451920(uint32_t* a2);
struct28* sub_4BD340(int a1, int a2, int a3, int a4);
uint32_t* sub_4BD280(int a1, int a2);
void nox_common_list_clear_425760(nox_list_item_t* list);
void* sub_4864A0(void* a3);
extern uint32_t dword_5d4594_1045420;
extern uint32_t dword_5d4594_1045428;
extern struct28* dword_5d4594_1045424;
extern uint32_t dword_5d4594_1045436;
extern uint32_t dword_5d4594_1045432;
*/
import "C"

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/tmp2"
)

var (
	Tmp *tmp2.Tmp
)

func initTmp() {
	var ptr = memmap.PtrT[[1023][200]uint8](0x5d4594, 840628)
	Tmp = tmp2.NewTmp(
		ptr,
		func(i int32) unsafe.Pointer {
			return unsafe.Pointer(nox_xxx_getSndName_40AF80(int(i)))
		},
		func(p1 *uint32) int32 {
			return int32(C.sub_451920((*C.uint32_t)(unsafe.Pointer(p1))))
		},
		func(a1, a2, a3, a4 int32) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD340(C.int(a1), C.int(a2), C.int(a3), C.int(a4)))
		},
		func(a1, a2 int32) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD280(C.int(a1), C.int(a2)))
		},
		func(list unsafe.Pointer) {
			C.nox_common_list_clear_425760((*C.nox_list_item_t)(list))
		},
		func(a3 unsafe.Pointer) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4864A0(a3))
		},
		(*uint32)(&C.dword_5d4594_1045420),
		(*uint32)(&C.dword_5d4594_1045428),
		(*unsafe.Pointer)(unsafe.Pointer(&C.dword_5d4594_1045424)),
		(*unsafe.Pointer)(unsafe.Pointer(&C.dword_5d4594_1045436)),
		(*uint32)(&C.dword_5d4594_1045432),
	)
}

//export nox_xxx_draw_452270
func nox_xxx_draw_452270(a1 int) *C.struct200 {
	return (*C.struct200)(unsafe.Pointer(Tmp.Get(a1)))
}
