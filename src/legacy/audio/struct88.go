package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct88 struct {
	field_0 ListElement[Struct88, *Struct88]
	field_3 *Struct587000_94032
	field_4 uint32
	field_5 uint32
	field_6 [16]*Struct264
}

var _ = [1]struct{}{}[88-unsafe.Sizeof(Struct88{})]

type Struct587000_94032 struct {
	field_0 unsafe.Pointer
	field_1 unsafe.Pointer
	field_2 unsafe.Pointer
	field_3 uint32
	field_4 unsafe.Pointer
	field_5 unsafe.Pointer
	field_6 unsafe.Pointer
	field_7 unsafe.Pointer
	field_8 unsafe.Pointer
	field_9 *Struct264Field64
}

type UnknownListElement struct {
	ListElement[UnknownListElement, *UnknownListElement]
}

type Struct587000_155144 struct {
	field_0      ListElement[Struct88, *Struct88]
	field_3      ListElement[Struct264, *Struct264]
	field_6      uint32
	field_7      uint32 // unknown
	timerGroup_8 timer.TimerGroup
}
