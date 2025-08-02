package struct312

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

type Struct312Module struct {
	moduleName string

	// External variables

	// External functions
	sub_4864A0                     func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_getNext_425940 func(unsafe.Pointer) unsafe.Pointer
	sub_425770                     func(unsafe.Pointer) unsafe.Pointer
	sub_4BD8C0_ptr                 unsafe.Pointer
	sub_4BD940_ptr                 unsafe.Pointer
	sub_4BD9B0_ptr                 unsafe.Pointer
	ptr_uint32_5d4594_1193340      *uint32
}

func NewStruct312Module(
	moduleName string,
	sub_4864A0 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_getNext_425940 func(unsafe.Pointer) unsafe.Pointer,
	sub_425770 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BD8C0_ptr unsafe.Pointer,
	sub_4BD940_ptr unsafe.Pointer,
	sub_4BD9B0_ptr unsafe.Pointer,

) *Struct312Module {
	return &Struct312Module{
		moduleName:                     moduleName,
		sub_4864A0:                     sub_4864A0,
		nox_common_list_getNext_425940: nox_common_list_getNext_425940,
		sub_425770:                     sub_425770,
		sub_4BD8C0_ptr:                 sub_4BD8C0_ptr,
		sub_4BD940_ptr:                 sub_4BD940_ptr,
		sub_4BD9B0_ptr:                 sub_4BD9B0_ptr,
		ptr_uint32_5d4594_1193340:      memmap.PtrUint32(0x5D4594, 1193340),
	}
}
