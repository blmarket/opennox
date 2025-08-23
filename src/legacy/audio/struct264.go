package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

// Size unknown
type Struct264Field64 struct {
	field_0 unsafe.Pointer
	field_1 unsafe.Pointer // function pointer takes *Struct312 as an argument
	field_2 unsafe.Pointer // Function pointer takes *Struct312 as an argument
	field_3 unsafe.Pointer
	field_4 unsafe.Pointer
	field_5 unsafe.Pointer
	field_6 unsafe.Pointer
	field_7 unsafe.Pointer
	field_8 unsafe.Pointer // function pointer takes *Struct312 as an argument
}

type Struct264 struct {
	field_0       ListElement[Struct264]
	field_3       uint32
	field_4       uint32
	field_5       *Struct88
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
	field_50      ListElement[Struct312]
	field_53      int32
	field_54      unsafe.Pointer // sub_4873C0_ptr, takes *Struct264
	field_55      uint32
	field_56      int64
	field_58      int64
	field_60      int64
	field_62      int64
	field_64      *Struct264Field64
	field_65      uint32
}

var _ = [1]struct{}{}[264-unsafe.Sizeof(Struct264{})]
