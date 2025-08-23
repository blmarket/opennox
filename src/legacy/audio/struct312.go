package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct312 struct {
	ListElement[Struct312]
	field_3       int32
	timerGroup_4  timer.TimerGroup
	field_28      *timer.TimerGroup
	field_29      *timer.TimerGroup
	field_30      int32
	field_31      uint32
	field_32      int32
	field_33      *Struct264
	field_34      uint32
	field_35      unsafe.Pointer // sub_452770_ptr
	field_36      unsafe.Pointer // sub_4526F0_ptr
	field_37      unsafe.Pointer // sub_4526D0_ptr
	field_38      *Struct576
	field_39      uint32
	field_40      uint32
	field_41      uint32
	field_42      uint32
	field_43      *Struct264Field64
	timerGroup_44 timer.TimerGroup
	field_68      uint32
	field_69      unsafe.Pointer // sub_4BD8C0_ptr
	field_70      unsafe.Pointer // sub_4BD940_ptr
	field_71      unsafe.Pointer // sub_4BD9B0_ptr
	field_72      unsafe.Pointer
	field_73      *UnknownListElement
	field_74      uint32
	field_75      uint32
	field_76      uint32
	field_77      uint32
}

var _ = [1]struct{}{}[312-unsafe.Sizeof(Struct312{})]
