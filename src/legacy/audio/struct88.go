package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct88 struct {
	field_0 ListElement[Struct88, *Struct88]
	field_3 unsafe.Pointer
	field_4 uint32
	field_5 [17]uint32
}

var _ = [1]struct{}{}[88-unsafe.Sizeof(Struct88{})]

type UnknownListElement struct {
	ListElement[UnknownListElement, *UnknownListElement]
}

type Struct587000_155144 struct {
	field_0      ListElement[UnknownListElement, *UnknownListElement]
	field_3      ListElement[UnknownListElement, *UnknownListElement]
	field_6      uint32
	field_7      uint32 // unknown
	timerGroup_8 timer.TimerGroup
}

type Struct88Module struct {
	moduleName string
	externs    *AudioExterns

	// External functions
	sub_487680                          func(*Struct264)
	nox_common_list_getNextSafe_4258A0  func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_remove_425920       func(unsafe.Pointer)
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_append_4258E0       func(unsafe.Pointer, unsafe.Pointer)
}

func NewStruct88Module(
	externs *AudioExterns,
	moduleName string,
	sub_487680 func(*Struct264),
	nox_common_list_getNextSafe_4258A0 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_remove_425920 func(unsafe.Pointer),
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
) *Struct88Module {
	return &Struct88Module{
		moduleName:                          moduleName,
		externs:                             externs,
		sub_487680:                          sub_487680,
		nox_common_list_getNextSafe_4258A0:  nox_common_list_getNextSafe_4258A0,
		nox_common_list_remove_425920:       nox_common_list_remove_425920,
		nox_common_list_getFirstSafe_425890: nox_common_list_getFirstSafe_425890,
		nox_common_list_append_4258E0:       nox_common_list_append_4258E0,
	}
}
