package struct576

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type ListItem struct {
	next *ListItem
	prev *ListItem
	head *ListItem
}

var _ = [1]struct{}{}[12-unsafe.Sizeof(ListItem{})] // Ensure listItem is 12 bytes

type Struct576 struct {
	Next         *Struct576 // 0
	Prev         *Struct576 // 1
	Head         unsafe.Pointer
	Field3       uint32
	Field4       uint32           // 4
	Field5       uint32           // 5
	Field6       uint8            // 6
	Field6_1     uint8            // 6 - stores a1 parameter
	Field6_2     uint8            // 6 - stores a2 parameter
	Field6_3     uint8            // 6 - stores a3 parameter
	Field7       uint32           // 7
	Field8       uint32           // 8
	Field9       *uint32          // 9 - stores a1 parameter
	Field10      [32]uint32       // 10-41
	Field42      uint32           // 42 - contains length of field_10 array
	Field43      uint32           // 43
	Field44      uint32           // 44 - looks like a pointer to some struct
	Field45      uint32           // 45
	TimerGroup46 timer.TimerGroup // 46
	Field70      uint32           // 70 - stores incrementing counter
	Field71      uint32           // 71
	Field72      uint64           // 72 - 64 bit
	Field74      uint32           // 74 - looks like a pointer to some struct
	Field75      uint32           // 75
	Field76      [32]uint32       // 76-107
	Field108     uint32           // 108
	Field109     uint32           // 109
	Field110     [32]uint32       // 110-141
	Field142     uint32           // 142
	Field143     uint32           // 143
}

var _ = [1]struct{}{}[576-unsafe.Sizeof(Struct576{})] // Ensure Struct576 is 576 bytes

// Module represents the struct576 module state and operations
type Module struct {
	dword_5d4594_1045424 unsafe.Pointer
	dword_5d4594_1045436 *uint32
	dword_587000_126996  *uint32
	dword_5d4594_1045432 *uint32
	get_list_at_840612   func() *ListItem
}

// NewModule creates a new struct576 module instance
func NewModule(
	dword_5d4594_1045424 unsafe.Pointer,
	dword_5d4594_1045436 *uint32,
	dword_587000_126996 *uint32,
	dword_5d4594_1045432 *uint32,
	get_list_at_840612 func() *ListItem,
) *Module {
	return &Module{
		dword_5d4594_1045424: dword_5d4594_1045424,
		dword_5d4594_1045436: dword_5d4594_1045436,
		dword_587000_126996:  dword_587000_126996,
		dword_5d4594_1045432: dword_5d4594_1045432,
		get_list_at_840612:   get_list_at_840612,
	}
}

func (m *Module) Sub_452F10(a1p *Struct576, a2 int32) uint32 {
	v2 := a2
	if a2 <= 100 {
		if a2 < 0 {
			v2 = 0
		}
	} else {
		v2 = 100
	}

	// Access field_9 + 20 bytes offset
	// field_9 is a pointer to uint32, so we need to get the value at offset 20/4 = 5 uint32s
	field9Ptr := uintptr(unsafe.Pointer(a1p.Field9))
	valueAtOffset := *(*uint32)(unsafe.Pointer(field9Ptr + 20))

	return uint32((163 * v2 * int32(valueAtOffset>>16)) >> 14)
}

func (m *Module) Sub_452EE0(a1 *Struct576, a2 int32) {
	v2 := m.Sub_452F10(a1, a2)
	a1.TimerGroup46.Timers[0].SetRaw(v2)
	a1.TimerGroup46.Timers[0].Update()
}

func (m *Module) Sub_452F50(a1p *Struct576, a2 int32) {
	v2 := m.Sub_452F10(a1p, a2)
	a1p.TimerGroup46.Timers[0].SetInterp(v2)
}

// //----- (00452F50) --------------------------------------------------------
// int sub_452F50(struct576* a1p, int a2) {
// 	int v2; // eax
// 	v2 = sub_452F10(a1p, a2);
// 	return sub_486350(&a1p->timerGroup_46.field_0, v2);
// }

func (m *Module) Sub_452F80(a1 *Struct576, a2 int32) {
	v2 := m.sub_452FA0(a2)
	a1.TimerGroup46.Timers[2].SetRaw(uint32(v2))
}

// //----- (00452F80) --------------------------------------------------------
// uint32_t* sub_452F80(struct576* a1, int a2) {
// 	int v2; // eax

// 	v2 = sub_452FA0(a2);
// 	return sub_486320(&a1->timerGroup_46.field_16, v2);
// }

func (m *Module) sub_452FA0(a1 int32) int32 {
	v1 := a1
	if a1 <= 50 {
		if a1 < -50 {
			v1 = -50
		}
	} else {
		v1 = 50
	}
	return (v1*8192)/50 + 8192
}

// //----- (00452FA0) --------------------------------------------------------
// int sub_452FA0(int a1) {
// 	int v1; // eax

// 	v1 = a1;
// 	if (a1 <= 50) {
// 		if (a1 < -50) {
// 			v1 = -50;
// 		}
// 	} else {
// 		v1 = 50;
// 	}
// 	return (v1 * 8192) / 50 + 8192;
// }

func (m *Module) Sub_452FE0(a1p *Struct576, a2 int32) {
	v2 := m.sub_452FA0(a2)
	a1p.TimerGroup46.Timers[2].SetInterp(uint32(v2))
}

// //----- (00452FE0) --------------------------------------------------------
// int sub_452FE0(struct576* a1p, int a2) {
// 	int v2; // eax

// 	v2 = sub_452FA0(a2);
// 	return sub_486350(&a1p->timerGroup_46.field_16, v2);
// }
