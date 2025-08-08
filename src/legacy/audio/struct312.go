package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct312 struct {
	ListItem
	field_3       int32
	timerGroup_4  timer.TimerGroup
	field_28      uint32
	field_29      uint32
	field_30      uint32
	field_31      uint32
	field_32      uint32
	field_33      *Struct264
	field_34      uint32
	field_35      uint32
	field_36      uint32
	field_37      uint32
	field_38      uint32
	field_39      uint32
	field_40      uint32
	field_41      uint32
	field_42      uint32
	field_43      unsafe.Pointer
	timerGroup_44 timer.TimerGroup
	field_68      uint32
	field_69      unsafe.Pointer // sub_4BD8C0_ptr
	field_70      unsafe.Pointer // sub_4BD940_ptr
	field_71      unsafe.Pointer // sub_4BD9B0_ptr
	field_72      uint32
	field_73      [5]uint32
}

var _ = [1]struct{}{}[312-unsafe.Sizeof(Struct312{})]

func (s *Struct312) getList() *ListItem {
	return &s.ListItem
}

type Struct312Module struct {
	moduleName string
	externs    *AudioExterns

	// External variables

	// External functions
	sub_4864A0                     func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_getNext_425940 func(unsafe.Pointer) unsafe.Pointer
	sub_425770                     func(*ListItem) unsafe.Pointer
	sub_4BD8C0_ptr                 unsafe.Pointer
	sub_4BD940_ptr                 unsafe.Pointer
	sub_4BD9B0_ptr                 unsafe.Pointer
	ptr_uint32_5d4594_1193340      *uint32
}

func NewStruct312Module(
	externs *AudioExterns,
	moduleName string,
	sub_4864A0 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_getNext_425940 func(unsafe.Pointer) unsafe.Pointer,
	sub_425770 func(*ListItem) unsafe.Pointer,
	sub_4BD8C0_ptr unsafe.Pointer,
	sub_4BD940_ptr unsafe.Pointer,
	sub_4BD9B0_ptr unsafe.Pointer,

) *Struct312Module {
	return &Struct312Module{
		moduleName:                     moduleName,
		externs:                        externs,
		sub_4864A0:                     sub_4864A0,
		nox_common_list_getNext_425940: nox_common_list_getNext_425940,
		sub_425770:                     sub_425770,
		sub_4BD8C0_ptr:                 sub_4BD8C0_ptr,
		sub_4BD940_ptr:                 sub_4BD940_ptr,
		sub_4BD9B0_ptr:                 sub_4BD9B0_ptr,
		ptr_uint32_5d4594_1193340:      memmap.PtrUint32(0x5D4594, 1193340),
	}
}
