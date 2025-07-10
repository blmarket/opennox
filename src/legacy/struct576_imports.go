package legacy

/*
#include <stdint.h>
#include "defs.h"

extern struct28* dword_5d4594_1045424;
extern uint32_t* dword_5d4594_1045436;
extern uint32_t dword_587000_126996;
extern uint32_t dword_5d4594_1045432;
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
		dword_5d4594_1045424 unsafe.Pointer = unsafe.Pointer(&C.dword_5d4594_1045424)
		dword_5d4594_1045436 *uint32        = (*uint32)(unsafe.Pointer(&C.dword_5d4594_1045436))
		dword_587000_126996  *uint32        = (*uint32)(&C.dword_587000_126996)
		dword_5d4594_1045432 *uint32        = (*uint32)(&C.dword_5d4594_1045432)
	)
	Struct576Module = struct576.NewModule(
		dword_5d4594_1045424,
		dword_5d4594_1045436,
		dword_587000_126996,
		dword_5d4594_1045432,
		func() *struct576.ListItem {
			return memmap.PtrT[struct576.ListItem](0x5D4594, 840612)
		},
	)
}

//export sub_452F10
func sub_452F10(a1p *C.struct576, a2 int32) uint32 {
	return Struct576Module.Sub_452F10((*struct576.Struct576)(unsafe.Pointer(a1p)), a2)
}
