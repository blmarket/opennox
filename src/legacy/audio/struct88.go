package audio

import (
	"unsafe"
)

type Struct88 struct {
	field_0 ListItem
	field_3 unsafe.Pointer
	field_4 uint32
	field_5 [17]uint32
}

var _ = [1]struct{}{}[88-unsafe.Sizeof(Struct88{})]

type Struct587000_155144 struct {
	field_0 ListItem
	field_3 ListItem
	field_6 uint32
}

type Struct88Module struct {
	moduleName string

	// External variables
	dword_587000_155144 **Struct587000_155144

	// External functions
	sub_487680                          func(*Struct264)
	nox_common_list_getNextSafe_4258A0  func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_remove_425920       func(unsafe.Pointer)
	sub_425770                          func(*ListItem) unsafe.Pointer
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_append_4258E0       func(unsafe.Pointer, unsafe.Pointer)
}

func NewStruct88Module(
	moduleName string,
	dword_587000_155144 **Struct587000_155144,
	sub_487680 func(*Struct264),
	nox_common_list_getNextSafe_4258A0 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_remove_425920 func(unsafe.Pointer),
	sub_425770 func(*ListItem) unsafe.Pointer,
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
) *Struct88Module {
	return &Struct88Module{
		moduleName:                          moduleName,
		dword_587000_155144:                 dword_587000_155144,
		sub_487680:                          sub_487680,
		nox_common_list_getNextSafe_4258A0:  nox_common_list_getNextSafe_4258A0,
		nox_common_list_remove_425920:       nox_common_list_remove_425920,
		sub_425770:                          sub_425770,
		nox_common_list_getFirstSafe_425890: nox_common_list_getFirstSafe_425890,
		nox_common_list_append_4258E0:       nox_common_list_append_4258E0,
	}
}
