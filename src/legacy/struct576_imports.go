package legacy

/*
#include <stdint.h>
#include "defs.h"

extern struct28* dword_5d4594_1045424;
extern uint32_t* dword_5d4594_1045436;
extern uint32_t dword_587000_126996;
extern uint32_t dword_5d4594_1045432;

uint32_t* sub_4BD470(uint32_t** a1, int a2);
void sub_4BD650(int a1);
int sub_4BD660(int a1);
int sub_4BD710(int a1);
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/struct576"
)

var (
	Struct576Module *struct576.Module
)

func initStruct576() {
	var (
		dword_5d4594_1045424 *unsafe.Pointer = (*unsafe.Pointer)(unsafe.Pointer(&C.dword_5d4594_1045424))
		dword_5d4594_1045436 *unsafe.Pointer = (*unsafe.Pointer)(unsafe.Pointer(&C.dword_5d4594_1045436))
		dword_587000_126996  *uint32         = (*uint32)(&C.dword_587000_126996)
		dword_5d4594_1045432 *uint32         = (*uint32)(&C.dword_5d4594_1045432)
	)
	Struct576Module = struct576.NewModule(
		dword_5d4594_1045424,
		dword_5d4594_1045436,
		dword_587000_126996,
		dword_5d4594_1045432,
		func() *struct576.ListItem {
			return memmap.PtrT[struct576.ListItem](0x5D4594, 840612)
		},
		func(min, max int, file unsafe.Pointer, line int) int {
			return nox_common_randomIntMinMax_415FF0(
				min, max,
				(*C.char)(file), line,
			)
		},
		func(a1 unsafe.Pointer, a2 int32) unsafe.Pointer {
			return unsafe.Pointer(C.sub_4BD470((**C.uint32_t)(a1), C.int(a2)))
		},
		func(a1 unsafe.Pointer) {
			C.sub_4BD650((C.int)(uintptr(a1)))
		},
		func(a1 unsafe.Pointer) {
			C.sub_4BD660((C.int)(uintptr(a1)))
		},
		func(a1 unsafe.Pointer) int32 {
			return int32(C.sub_4BD710((C.int)(uintptr(a1))))
		},
	)
}

//export sub_452F10
func sub_452F10(a1p *C.struct576, a2 int32) uint32 {
	return Struct576Module.Sub_452F10((*struct576.Struct576)(unsafe.Pointer(a1p)), a2)
}

//export sub_452EE0
func sub_452EE0(a1p *C.struct576, a2 int32) {
	Struct576Module.Sub_452EE0((*struct576.Struct576)(unsafe.Pointer(a1p)), a2)
}

//export sub_452F50
func sub_452F50(a1p *C.struct576, a2 int32) {
	Struct576Module.Sub_452F50((*struct576.Struct576)(unsafe.Pointer(a1p)), a2)
}

//export sub_452F80
func sub_452F80(a1 *C.struct576, a2 int32) {
	Struct576Module.Sub_452F80((*struct576.Struct576)(unsafe.Pointer(a1)), a2)
}

//export sub_452FE0
func sub_452FE0(a1 *C.struct576, a2 int32) {
	Struct576Module.Sub_452FE0((*struct576.Struct576)(unsafe.Pointer(a1)), a2)
}

//export sub_451F30
func sub_451F30(a1p *C.struct576, a2 int32) {
	Struct576Module.Sub_451F30((*struct576.Struct576)(unsafe.Pointer(a1p)), a2)
}

//export sub_451F90
func sub_451F90(a1p *C.struct576) {
	Struct576Module.Sub_451F90((*struct576.Struct576)(unsafe.Pointer(a1p)))
}

//export sub_451CA0
func sub_451CA0(a1p *C.struct576) int32 {
	return Struct576Module.Sub_451CA0((*struct576.Struct576)(unsafe.Pointer(a1p)))
}

//export sub_451CF0
func sub_451CF0(a1p *C.struct576) int32 {
	return Struct576Module.Sub_451CF0((*struct576.Struct576)(unsafe.Pointer(a1p)))
}

//export sub_451E80
func sub_451E80(a1p *C.struct576) int32 {
	return Struct576Module.Sub_451E80((*struct576.Struct576)(unsafe.Pointer(a1p)))
}

//export sub_451DC0
func sub_451DC0(a1p *C.struct576) {
	Struct576Module.Sub_451DC0((*struct576.Struct576)(unsafe.Pointer(a1p)))
}
