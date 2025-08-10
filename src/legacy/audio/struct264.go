package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct264 struct {
	field_0       ListElement[Struct264, *Struct264]
	field_3       uint32
	field_4       uint32
	field_5       unsafe.Pointer
	field_6       uint32
	field_7       uint32
	field_8       uint32
	field_9       uint32
	field_10      uint32
	field_11      uint32
	field_12      uint32
	field_13      uint32
	field_14      uint32
	field_15      [7]uint32
	TimerGroup_22 timer.TimerGroup
	field_46      *timer.TimerGroup
	field_47      uint32
	field_48      int32
	field_49      int32
	field_50      ListElement[Struct312, *Struct312]
	field_53      int32
	field_54      unsafe.Pointer // sub_4873C0_ptr, takes *Struct264
	field_55      uint32
	field_56      uint32
	field_57      uint32
	field_58      uint32
	field_59      uint32
	field_60      uint32
	field_61      uint32
	field_62      uint32
	field_63      uint32
	field_64      unsafe.Pointer
	field_65      uint32
}

var _ = [1]struct{}{}[264-unsafe.Sizeof(Struct264{})]

type Struct264Module struct {
	moduleName string
	externs    *AudioExterns

	// External functions
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer
	sub_487100                          func(*unsafe.Pointer) unsafe.Pointer
	nox_common_list_remove_425920       func(unsafe.Pointer)
	nullsub_10                          func(uint32)
	sub_4BDA60                          func(unsafe.Pointer)
	nox_common_list_clear_425760        func(unsafe.Pointer)
	nox_common_list_getNextSafe_4258A0  func(unsafe.Pointer) unsafe.Pointer
	sub_4BDA80                          func(*Struct312) int32
	sub_486320                          func(unsafe.Pointer, int) unsafe.Pointer
	sub_425770                          func(unsafe.Pointer) unsafe.Pointer
	sub_4870E0                          func(*unsafe.Pointer) unsafe.Pointer
	sub_4BD720                          func(*Struct264) *Struct312
}

func NewStruct264Module(
	externs *AudioExterns,
	moduleName string,
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer,
	sub_487100 func(*unsafe.Pointer) unsafe.Pointer,
	nox_common_list_remove_425920 func(unsafe.Pointer),
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
	nullsub_10 func(uint32),
	sub_4BDA60, nox_common_list_clear_425760 func(unsafe.Pointer),
	nox_common_list_getNextSafe_4258A0 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BDA80 func(*Struct312) int32,
	sub_486320 func(unsafe.Pointer, int) unsafe.Pointer,
	sub_425770 func(unsafe.Pointer) unsafe.Pointer,
	sub_4870E0 func(*unsafe.Pointer) unsafe.Pointer,
	sub_4BD720 func(*Struct264) *Struct312,
) *Struct264Module {
	return &Struct264Module{
		moduleName:                          moduleName,
		externs:                             externs,
		nox_common_list_getFirstSafe_425890: nox_common_list_getFirstSafe_425890,
		sub_487100:                          sub_487100,
		nox_common_list_remove_425920:       nox_common_list_remove_425920,
		nullsub_10:                          nullsub_10,
		sub_4BDA60:                          sub_4BDA60,
		nox_common_list_clear_425760:        nox_common_list_clear_425760,
		nox_common_list_getNextSafe_4258A0:  nox_common_list_getNextSafe_4258A0,
		sub_4BDA80:                          sub_4BDA80,
		sub_486320:                          sub_486320,
		sub_425770:                          sub_425770,
		sub_4870E0:                          sub_4870E0,
		sub_4BD720:                          sub_4BD720,
	}
}
