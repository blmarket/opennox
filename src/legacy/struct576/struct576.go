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
	Field4       uint32             // 4
	Field5       uint32             // 5
	Field6       uint8              // 6
	Field6_1     uint8              // 6 - stores a1 parameter
	Field6_2     uint8              // 6 - stores a2 parameter
	Field6_3     uint8              // 6 - stores a3 parameter
	Field7       uint32             // 7
	Field8       uint32             // 8
	Field9       *uint32            // 9 - stores a1 parameter
	Field10      [32]unsafe.Pointer // 10-41
	Field42      int32              // 42 - contains length of field_10 array
	Field43      int32              // 43
	Field44      uint32             // 44 - looks like a pointer to some struct
	Field45      uint32             // 45
	TimerGroup46 timer.TimerGroup   // 46
	Field70      uint32             // 70 - stores incrementing counter
	Field71      uint32             // 71
	Field72      uint64             // 72 - 64 bit
	Field74      uint32             // 74 - looks like a pointer to some struct
	Field75      uint32             // 75
	Field76      [32]int32          // 76-107
	Field108     int32              // 108
	Field109     int32              // 109
	Field110     [32]uint32         // 110-141
	Field142     uint32             // 142
	Field143     uint32             // 143
}

var _ = [1]struct{}{}[576-unsafe.Sizeof(Struct576{})] // Ensure Struct576 is 576 bytes

// Module represents the struct576 module state and operations
type Module struct {
	dword_5d4594_1045424              *unsafe.Pointer
	dword_5d4594_1045436              *unsafe.Pointer
	dword_587000_126996               *uint32
	dword_5d4594_1045432              *uint32
	get_list_at_840612                func() *ListItem
	nox_common_randomIntMinMax_415FF0 func(min, max int, file unsafe.Pointer, line int) int
	sub_4BD470                        func(a1 unsafe.Pointer, a2 int32) unsafe.Pointer
	sub_4BD650                        func(a1 unsafe.Pointer)
	sub_4BD660                        func(a1 unsafe.Pointer)
	sub_4BD710                        func(a1 unsafe.Pointer) int32
}

// NewModule creates a new struct576 module instance
func NewModule(
	dword_5d4594_1045424 *unsafe.Pointer,
	dword_5d4594_1045436 *unsafe.Pointer,
	dword_587000_126996 *uint32,
	dword_5d4594_1045432 *uint32,
	get_list_at_840612 func() *ListItem,
	nox_common_randomIntMinMax_415FF0 func(min, max int, file unsafe.Pointer, line int) int,
	sub_4BD470 func(a1 unsafe.Pointer, a2 int32) unsafe.Pointer,
	sub_4BD650 func(a1 unsafe.Pointer),
	sub_4BD660 func(a1 unsafe.Pointer),
	sub_4BD710 func(a1 unsafe.Pointer) int32,
) *Module {
	return &Module{
		dword_5d4594_1045424:              dword_5d4594_1045424,
		dword_5d4594_1045436:              dword_5d4594_1045436,
		dword_587000_126996:               dword_587000_126996,
		dword_5d4594_1045432:              dword_5d4594_1045432,
		get_list_at_840612:                get_list_at_840612,
		nox_common_randomIntMinMax_415FF0: nox_common_randomIntMinMax_415FF0,
		sub_4BD470:                        sub_4BD470,
		sub_4BD650:                        sub_4BD650,
		sub_4BD660:                        sub_4BD660,
		sub_4BD710:                        sub_4BD710,
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

func (m *Module) Sub_451F30(a1p *Struct576, a2 int32) int32 {
	// Calculate offset into field_9: field_9 + 2 * a2 + 128
	// field_9 is *uint32, so we work with 16-bit values (int16)
	field9Ptr := uintptr(unsafe.Pointer(a1p.Field9))
	offsetPtr := unsafe.Pointer(field9Ptr + uintptr(2*a2+128))
	value := *(*int16)(offsetPtr)

	// Call sub_4BD470 and store result in field_10[field_42]
	a1p.Field10[a1p.Field42] = m.sub_4BD470(*m.dword_5d4594_1045424, int32(value))
	if a1p.Field10[a1p.Field42] == nil {
		return 0
	}

	m.sub_4BD650(a1p.Field10[a1p.Field42])
	a1p.Field42 += 1
	return a1p.Field42
}

// //----- (00451F30) --------------------------------------------------------
// int sub_451F30(struct576* a1p, int a2) {
//     int v2;     // edx
//     int result; // eax

//     a1p->field_10[a1p->field_42] = sub_4BD470(dword_5d4594_1045424, *(short*)((uint32_t)a1p->field_9 + 2 * a2 + 128));
//     v2 = a1p->field_42;
//     result = a1p->field_10[v2];
//     if (result) {
//         sub_4BD650(a1p->field_10[v2]);
//         result = a1p->field_42 + 1;
//         a1p->field_42 = result;
//     }
//     return result;
// }

func (m *Module) Sub_451F90(a1p *Struct576) {
	if a1p.Field42 <= 0 {
		a1p.Field42 = 0
	} else {
		for i := int32(0); i < a1p.Field42; i++ {
			m.sub_4BD660(a1p.Field10[i])
			a1p.Field10[i] = nil
		}
		a1p.Field42 = 0
	}
}

// //----- (00451F90) --------------------------------------------------------
// int sub_451F90(struct576* a1p) {
// 	int v1;     // edi
// 	int result; // eax
// 	int* v3;    // esi

// 	v1 = 0;
// 	result = a1p->field_42;
// 	if (result <= 0) {
// 		a1p->field_42 = 0;
// 	} else {
// 		v3 = a1p->field_10;
// 		do {
// 			sub_4BD660(v3[0]);
// 			v3[0] = 0;
// 			result = a1p->field_42;
// 			++v1;
// 			++v3;
// 		} while (v1 < result);
// 		a1p->field_42 = 0;
// 	}
// 	return result;
// }

func (m *Module) Sub_451CA0(a1p *Struct576) int32 {
	v1 := a1p.Field42
	a1p.Field108 = v1
	if v1 == 0 {
		return 0
	}

	if v1 > 0 {
		for i := int32(0); i < a1p.Field108; i++ {
			a1p.Field76[i] = i
		}
	}

	a1p.Field43 = -1
	return m.Sub_451CF0(a1p)
}

// //----- (00451CA0) --------------------------------------------------------
// int sub_451CA0(struct576* a1p) {
// 	int v1;       // ecx
// 	int v3;       // eax
// 	uint32_t* v4; // ecx

// 	v1 = a1p->field_42;
// 	a1p->field_108 = v1;
// 	if (!v1) {
// 		return 0;
// 	}
// 	v3 = 0;
// 	if (v1 > 0) {
// 		v4 = a1p->field_76;
// 		do {
// 			*v4 = v3++;
// 			++v4;
// 		} while (v3 < a1p->field_108);
// 	}
// 	a1p->field_43 = -1;
// 	return sub_451CF0(a1p);
// }

func (m *Module) Sub_451CF0(a1p *Struct576) int32 {
	v1 := uintptr(unsafe.Pointer(a1p.Field9))
	result := a1p.Field108
	v3 := *(*uint32)(unsafe.Pointer(v1 + 4))

	if a1p.Field108 != 0 {
		if v3&2 != 0 {
			// Random selection mode
			v5 := m.nox_common_randomIntMinMax_415FF0(0, int(result-1), nil, 376)
			v6 := a1p.Field108 - 1
			a1p.Field43 = a1p.Field76[v5]
			v7 := int32(v5)

			// Shift remaining elements in field_76 array
			if v7 < v6 {
				for i := v7; i < a1p.Field108-1; i++ {
					a1p.Field76[i] = a1p.Field76[i+1]
				}
			}
		} else {
			// Sequential mode
			a1p.Field43++
		}

		v9 := a1p.Field43
		a1p.Field108--
		result = m.sub_4BD710(a1p.Field10[v9])
	} else if v3&1 != 0 {
		// Loop mode
		v1_60 := *(*uint32)(unsafe.Pointer(v1 + 60))
		if v1_60 != 0 {
			v4 := a1p.Field109 + 1
			a1p.Field109 = v4
			if v4 >= int32(v1_60) {
				result = 0
			} else {
				result = m.Sub_451CA0(a1p)
			}
		} else {
			result = m.Sub_451CA0(a1p)
		}
	}

	return result
}

// //----- (00451CF0) --------------------------------------------------------
// int sub_451CF0(struct576* a1p) {
// 	int v1;       // ecx
// 	int result;   // eax
// 	int v3;       // edx
// 	int v4;       // edi
// 	int v5;       // eax
// 	int v6;       // edi
// 	int v7;       // ecx
// 	uint32_t* v8; // eax
// 	int v9;       // eax

// 	v1 = (int)a1p->field_9;
// 	result = a1p->field_108;
// 	v3 = *(uint32_t*)(v1 + 4);
// 	if (result) {
// 		if (v3 & 2) {
// 			v5 = nox_common_randomIntMinMax_415FF0(0, result - 1, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 376);
// 			v6 = a1p->field_108 - 1;
// 			a1p->field_43 = a1p->field_76[v5];
// 			v7 = v5;
// 			if (v5 < v6) {
// 				v8 = &a1p->field_76[v5];
// 				do {
// 					++v7;
// 					*v8 = v8[1];
// 					++v8;
// 				} while (v7 < a1p->field_108 - 1);
// 			}
// 		} else {
// 			++a1p->field_43;
// 		}
// 		v9 = a1p->field_43;
// 		--a1p->field_108;
// 		result = sub_4BD710(a1p->field_10[v9]);
// 	} else if (v3 & 1) {
// 		if (*(uint32_t*)(v1 + 60) && (v4 = a1p->field_109 + 1, a1p->field_109 = v4, v4 >= *(int*)(v1 + 60))) {
// 			result = 0;
// 		} else {
// 			result = sub_451CA0(a1p);
// 		}
// 	}
// 	return result;
// }

func (m *Module) Sub_451DC0(a1p *Struct576) {
	v1 := a1p.Field9
	result := a1p.Field42

	// Access v1[1] - field_9 is *uint32, so v1[1] is at offset 4 bytes
	field9Ptr := uintptr(unsafe.Pointer(v1))
	v3 := *(*uint32)(unsafe.Pointer(field9Ptr + 4))

	if result != 0 {
		// Access v1[17] - at offset 17*4 = 68 bytes
		v1_17 := *(*uint32)(unsafe.Pointer(field9Ptr + 68))
		if v1_17 < 0x21 {
			return
		}
		m.Sub_451F90(a1p)
	}

	if v3&4 != 0 {
		// Access v1[17] again
		v1_17 := *(*uint32)(unsafe.Pointer(field9Ptr + 68))
		if v1_17 >= 0x21 {
			v5 := m.Sub_451E80(a1p)
			result = m.Sub_451F30(a1p, v5)
		} else {
			// Access v1[48] - at offset 48*4 = 192 bytes
			v1_48 := *(*uint32)(unsafe.Pointer(field9Ptr + 192))
			result = int32(v1_48)
			for i := int32(0); i < result; i++ {
				m.Sub_451F30(a1p, i)
				// Re-read v1[48] in case it changed
				result = int32(*(*uint32)(unsafe.Pointer(field9Ptr + 192)))
			}
		}
	} else if v3&2 != 0 {
		// Access v1[48] for random range
		v1_48 := *(*uint32)(unsafe.Pointer(field9Ptr + 192))
		v6 := m.nox_common_randomIntMinMax_415FF0(0, int(v1_48-1), nil, 536)
		result = m.Sub_451F30(a1p, int32(v6))
	} else {
		result = m.Sub_451F30(a1p, 0)
	}
}

// //----- (00451DC0) --------------------------------------------------------
// int sub_451DC0(struct576* a1p) {
// 	uint32_t* v1; // esi
// 	int result;   // eax
// 	int v3;       // ebx
// 	int i;        // edi
// 	int v5;       // eax
// 	int v6;       // eax

// 	v1 = a1p->field_9;
// 	result = a1p->field_42;
// 	v3 = v1[1];
// 	if (result) {
// 		if (v1[17] < 0x21u) {
// 			return result;
// 		}
// 		sub_451F90(a1p);
// 	}
// 	if (v3 & 4) {
// 		if (v1[17] >= 0x21u) {
// 			v5 = sub_451E80(a1p);
// 			result = sub_451F30(a1p, v5);
// 		} else {
// 			result = v1[48];
// 			for (i = 0; i < result; ++i) {
// 				sub_451F30(a1p, i);
// 				result = v1[48];
// 			}
// 		}
// 	} else if (v3 & 2) {
// 		v6 = nox_common_randomIntMinMax_415FF0(0, v1[48] - 1, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 536);
// 		result = sub_451F30(a1p, v6);
// 	} else {
// 		result = sub_451F30(a1p, 0);
// 	}
// 	return result;
// }

func (m *Module) Sub_451E80(a1p *Struct576) int32 {
	// v1 = *(uint32_t*)(a1 + 36) - offset 36 corresponds to Field9
	v1 := a1p.Field9
	field9Ptr := uintptr(unsafe.Pointer(v1))

	// v2 = *(uint32_t*)(v1 + 4) - access v1[1]
	v2 := *(*uint32)(unsafe.Pointer(field9Ptr + 4))

	// if (*(int*)(a1 + 568) <= 0) - offset 568 corresponds to Field142 (568/4 = 142)
	if a1p.Field142 <= 0 {
		// v3 = *(uint32_t*)(v1 + 192) - access v1[48] (192/4 = 48)
		v3 := *(*uint32)(unsafe.Pointer(field9Ptr + 192))
		v4 := uint32(0)

		// *(uint32_t*)(a1 + 568) = v3
		a1p.Field142 = v3

		if v3 > 0 {
			// v5 = a1 + 440 - offset 440 corresponds to Field110 (440/4 = 110)
			for v4 < v3 {
				// v6 = v3 - v4++ - 1
				v6 := v3 - v4 - 1
				// *(uint32_t*)(v5 - 4) = v6 - store in Field110[v4]
				a1p.Field110[v4] = v6
				v4++
				// v3 = *(uint32_t*)(a1 + 568) - reload Field142
				v3 = a1p.Field142
			}
		}
	}

	// v7 = *(uint32_t*)(a1 + 568) - 1
	v7 := a1p.Field142 - 1
	// *(uint32_t*)(a1 + 568) = v7
	a1p.Field142 = v7

	if (v2 & 2) == 0 {
		// return *(uint32_t*)(a1 + 4 * v7 + 440) - return Field110[v7]
		return int32(a1p.Field110[v7])
	}

	// Random selection mode
	v8 := m.nox_common_randomIntMinMax_415FF0(0, int(v7), nil, 431)
	// v9 = *(uint32_t*)(a1 + 4 * v8 + 440) - get Field110[v8]
	v9 := a1p.Field110[v8]
	v10 := int32(v8)

	// Shift remaining elements in Field110 array
	if v8 < int(a1p.Field142) {
		for v10 < int32(a1p.Field142) {
			a1p.Field110[v10] = a1p.Field110[v10+1]
			v10++
		}
	}

	return int32(v9)
}

// //----- (00451E80) --------------------------------------------------------
// int sub_451E80(struct576* a1p) {
// 	int a1 = a1p;
// 	int v1;        // eax
// 	int v2;        // ebx
// 	int v3;        // eax
// 	int v4;        // ecx
// 	int v5;        // edx
// 	int v6;        // eax
// 	int v7;        // edx
// 	int v8;        // eax
// 	int v9;        // edi
// 	int v10;       // ecx
// 	uint32_t* v11; // eax

// 	v1 = *(uint32_t*)(a1 + 36);
// 	v2 = *(uint32_t*)(v1 + 4);
// 	if (*(int*)(a1 + 568) <= 0) {
// 		v3 = *(uint32_t*)(v1 + 192);
// 		v4 = 0;
// 		*(uint32_t*)(a1 + 568) = v3;
// 		if (v3 > 0) {
// 			v5 = a1 + 440;
// 			do {
// 				v5 += 4;
// 				v6 = v3 - v4++ - 1;
// 				*(uint32_t*)(v5 - 4) = v6;
// 				v3 = *(uint32_t*)(a1 + 568);
// 			} while (v4 < v3);
// 		}
// 	}
// 	v7 = *(uint32_t*)(a1 + 568) - 1;
// 	*(uint32_t*)(a1 + 568) = v7;
// 	if (!(v2 & 2)) {
// 		return *(uint32_t*)(a1 + 4 * v7 + 440);
// 	}
// 	v8 = nox_common_randomIntMinMax_415FF0(0, v7, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 431);
// 	v9 = *(uint32_t*)(a1 + 4 * v8 + 440);
// 	v10 = v8;
// 	if (v8 < *(int*)(a1 + 568)) {
// 		v11 = (uint32_t*)(a1 + 4 * v8 + 440);
// 		do {
// 			++v10;
// 			*v11 = v11[1];
// 			++v11;
// 		} while (v10 < *(int*)(a1 + 568));
// 	}
// 	return v9;
// }
