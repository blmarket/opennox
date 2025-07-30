package struct264

import (
	"unsafe"
)

type Struct264 struct {
	field_0 [66]uint32
}

type Struct264Module struct {
	moduleName string

	// External variables
	sub_4873C0_ptr       unsafe.Pointer
	dword_5d4594_1045428 *uint32
	dword_587000_127004  *unsafe.Pointer
	dword_587000_155144  *unsafe.Pointer
	dword_5d4594_805984  *unsafe.Pointer

	// External functions
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer
	sub_487100                          func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_remove_425920       func(unsafe.Pointer)
	nox_common_list_append_4258E0       func(unsafe.Pointer, unsafe.Pointer)
	nullsub_10                          func(uint32)
	sub_4BDA60                          func(unsafe.Pointer)
	nox_common_list_clear_425760        func(unsafe.Pointer)
	sub_4864A0                          func(unsafe.Pointer) unsafe.Pointer
	sub_4873C0                          func(int) int
	nox_common_list_getNextSafe_4258A0  func(unsafe.Pointer) unsafe.Pointer
	sub_4BDA80                          func(int) int
	sub_486320                          func(unsafe.Pointer, int) unsafe.Pointer
	sub_425770                          func(unsafe.Pointer) unsafe.Pointer
	sub_4870E0                          func(unsafe.Pointer) unsafe.Pointer
}

func NewStruct264Module(
	moduleName string,
	sub_4873C0_ptr unsafe.Pointer,
	dword_5d4594_1045428 *uint32,
	dword_587000_127004 *unsafe.Pointer,
	dword_587000_155144 *unsafe.Pointer,
	dword_5d4594_805984 *unsafe.Pointer,
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer,
	sub_487100 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_remove_425920 func(unsafe.Pointer),
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
	nullsub_10 func(uint32),
	sub_4BDA60 func(unsafe.Pointer),
	nox_common_list_clear_425760 func(unsafe.Pointer),
	sub_4864A0 func(unsafe.Pointer) unsafe.Pointer,
	sub_4873C0 func(int) int,
	nox_common_list_getNextSafe_4258A0 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BDA80 func(int) int,
	sub_486320 func(unsafe.Pointer, int) unsafe.Pointer,
	sub_425770 func(unsafe.Pointer) unsafe.Pointer,
	sub_4870E0 func(unsafe.Pointer) unsafe.Pointer,
) *Struct264Module {
	return &Struct264Module{
		moduleName:                          moduleName,
		sub_4873C0_ptr:                      sub_4873C0_ptr,
		dword_5d4594_1045428:                dword_5d4594_1045428,
		dword_587000_127004:                 dword_587000_127004,
		dword_587000_155144:                 dword_587000_155144,
		dword_5d4594_805984:                 dword_5d4594_805984,
		nox_common_list_getFirstSafe_425890: nox_common_list_getFirstSafe_425890,
		sub_487100:                          sub_487100,
		nox_common_list_remove_425920:       nox_common_list_remove_425920,
		nox_common_list_append_4258E0:       nox_common_list_append_4258E0,
		nullsub_10:                          nullsub_10,
		sub_4BDA60:                          sub_4BDA60,
		nox_common_list_clear_425760:        nox_common_list_clear_425760,
		sub_4864A0:                          sub_4864A0,
		sub_4873C0:                          sub_4873C0,
		nox_common_list_getNextSafe_4258A0:  nox_common_list_getNextSafe_4258A0,
		sub_4BDA80:                          sub_4BDA80,
		sub_486320:                          sub_486320,
		sub_425770:                          sub_425770,
		sub_4870E0:                          sub_4870E0,
	}
}
