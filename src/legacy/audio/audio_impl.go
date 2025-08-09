package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

// Extracted struct definitions from defs.h
type Struct200 struct {
	field_0    unsafe.Pointer
	field_1    uint32
	field_2    uint32
	field_3    uint32
	field_4    timer.Timer
	field_12   uint32
	field_13   uint32
	field_14   uint32
	field_15   uint32
	field_16   uint32
	field_17   uint32
	field_18   uint32
	field_19   uint32
	field_20   uint32
	sndName_21 unsafe.Pointer // pointer to string
	field_22   ListItem
	field_25   uint32
	field_26   uint32
	field_27   int32
	field_28   Struct200Field28
	field_31   uint32
	field_32   [32]uint16
	field_48   uint32
	field_49   uint32
}

type Struct200Field28 struct {
	ListItem
}

func (s *Struct200Field28) getList() *ListItem {
	return &s.ListItem
}

func (s *Struct200Field28) getStruct() *Struct200 {
	return (*Struct200)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - 28*4))
}

var _ = [1]struct{}{}[200-unsafe.Sizeof(Struct200{})]
var _ = [1]struct{}{}[unsafe.Sizeof(Struct200{})-200]

type Struct576 struct {
	ListItem
	field_3       ListItem
	field_6       uint8
	field_6_1     uint8
	field_6_2     uint8
	field_6_3     uint8
	field_7       uint32
	field_8       uint32
	field_9       *Struct200
	field_10      [32]unsafe.Pointer
	field_42      uint32
	field_43      uint32
	field_44      *Struct312
	field_45      uint32
	timerGroup_46 timer.TimerGroup
	field_70      uint32
	field_71      uint32
	field_72      uint64
	field_74      uint32
	field_75      uint32
	field_76      [32]uint32
	field_108     uint32
	field_109     uint32
	field_110     [32]uint32
	field_142     uint32
	field_143     uint32
}

func (s *Struct576) getList() *ListItem {
	return &s.ListItem
}

var _ = [1]struct{}{}[576-unsafe.Sizeof(Struct576{})]

func (m *AudioModule) Sub_451850(a2p *Struct264, a3p unsafe.Pointer) int32 {
	var (
		a3     int32 = int32(uintptr(a3p))
		result int32
	)
	v4 := m.externs.Struct200Arr_5d4594_840628
	for i := int32(0); i < 1023; i++ {
		m.sub_451920(&v4[i])
		v4[i].sndName_21 = unsafe.Pointer(m.nox_xxx_getSndName_40AF80(int(i)))
	}
	*m.externs.Dword_5d4594_1045420 = uint32(a3)
	*m.externs.Dword_5d4594_1045428 = a2p
	if a3 != 0 {
		*m.externs.Dword_5d4594_1045424 = uint32(uintptr(unsafe.Pointer(m.sub_4BD340(int(a3), 0x100000, 200, 0x2000))))
		*m.externs.Dword_5d4594_1045436 = uint32(uintptr(unsafe.Pointer(m.sub_4BD280(200, 576))))
	}
	if *m.externs.Dword_5d4594_1045424 == 0 || *m.externs.Dword_5d4594_1045420 == 0 || *m.externs.Dword_5d4594_1045428 == nil || *m.externs.Dword_5d4594_1045436 == 0 {
		return 0
	}
	m.nox_common_list_clear_425760(unsafe.Pointer(m.externs.ListHead_5d4594_840612))
	m.externs.TimerGroup_5d4594_1045228.Init()
	result = 1
	(*m.externs.Dword_5d4594_1045428).field_46 = m.externs.TimerGroup_5d4594_1045228
	*m.externs.Dword_5d4594_1045432 = 1
	return result
}

func bool2int32(v bool) int32 {
	if v {
		return 1
	}
	return 0
}

func (m *AudioModule) sub_451920(a2p *Struct200) int32 {
	a2p.field_0 = nil
	a2p.field_1 = 0
	a2p.field_2 = 0
	a2p.field_14 = 0
	a2p.field_15 = 0
	a2p.field_19 = 0
	a2p.field_20 = 0
	a2p.field_12 = 1
	a2p.field_48 = 0
	a2p.field_18 = 0
	a2p.field_17 = 0
	a2p.field_25 = 0
	a2p.field_26 = 0
	a2p.field_16 = 600
	return bool2int32(a2p.field_4.Init(0x4000))
}

func (m *AudioModule) sub_452010() int32 {
	heads := m.externs.ListHeads_5d4594_839892
	for v1 := 0; v1 < 6; v1++ {
		for v2 := 0; v2 < 10; v2++ {
			heads[v1][v2].Clear()
		}
	}
	return int32(func() uint32 {
		p_ := memmap.PtrUint32(0x5D4594, 1045444)
		*p_++
		return *p_
	}())
}

func (m *AudioModule) sub_452190(a1_ *Struct200) {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	m.nox_common_list_remove_425920(unsafe.Pointer((**uint32)(unsafe.Pointer(&a1_.field_28))))
}

func (m *AudioModule) sub_4521A0(a1 int32) *Struct200 {
	heads := m.externs.ListHeads_5d4594_839892
	if a1 > 0 {
		for v1 := int32(0); v1 < a1; v1++ {
			v2 := heads[v1]
			for v3 := 0; v3 < 10; v3++ {
				v4 := v2[v3]
				v5 := v4.First()
				if v5 != nil {
					return v5.getStruct()
				}
			}
		}
	}
	return nil
}

func (m *AudioModule) sub_4521F0() int32 {
	var (
		result int32
		v1     *ListItem
		v2     *ListItem
	)
	result = int32(*m.externs.Dword_5d4594_1045432)
	if *m.externs.Dword_5d4594_1045432 != 0 {
		v1 = m.externs.ListHead_5d4594_840612.next
		if unsafe.Pointer(m.externs.ListHead_5d4594_840612.next) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
			for {
				v2 = v1.next
				m.Sub_4523D0((*Struct576)(unsafe.Pointer(v1)))
				result = m.sub_451FE0((*Struct576)(unsafe.Pointer(v1)))
				v1 = v2
				if unsafe.Pointer(v2) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
					break
				}
			}
		}
	}
	return result
}

func (m *AudioModule) sub_452230() {
	if *m.externs.Dword_5d4594_1045432 == 0 {
		return
	}
	if unsafe.Pointer(m.externs.ListHead_5d4594_840612.next) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
		res := m.externs.ListHead_5d4594_840612.First()
		for {
			v1 := res.next
			if res.field_6&1 != 0 {
				m.sub_451FE0(res)
			}
			if unsafe.Pointer(v1) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
				break
			}
			res = (*Struct576)(unsafe.Pointer(v1))
		}
	}
}

func (m *AudioModule) Nox_xxx_draw_452270(a1 int32) *Struct200 {
	if *m.externs.Dword_5d4594_1045432 != 0 && a1 >= 0 && a1 < 1023 {
		return &m.externs.Struct200Arr_5d4594_840628[a1]
	}
	return nil
}

func (m *AudioModule) Sub_4526F0(a1p *Struct312) int32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
	var (
		v1 *uint32
		v2 int32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 152)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*6)) &= 0xFFFFFFFD
	v2 = 4
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*7)) != 4 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*74)) != 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*142)) != 0 {
			v2 = 1
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*71)) = 0
		}
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*71)) != 0 {
			m.sub_452690((*Struct576)(unsafe.Pointer(v1)), int64(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*71))), v2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*71)) = 0
			return 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*7)) = uint32(v2)
	}
	return 0
}

func (m *AudioModule) sub_451CA0(a1_ *Struct576) int32 {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	_ = a1
	var v1 int32
	var v3 int32
	var v4 *uint32
	v1 = int32(a1_.field_42)
	a1_.field_108 = uint32(v1)
	if v1 == 0 {
		return 0
	}
	v3 = 0
	if v1 > 0 {
		v4 = &a1_.field_76[0]
		for {
			*v4 = uint32(func() int32 {
				p_ := &v3
				x := *p_
				*p_++
				return x
			}())
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1))
			if uint32(v3) >= a1_.field_108 {
				break
			}
		}
	}
	a1_.field_43 = 4294967295
	return m.sub_451CF0((*uint32)(unsafe.Pointer(a1_)))
}

func (m *AudioModule) sub_451F30(a1_ *Struct576, a2 int32) int32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int32
	var result int32
	*(*uint32)(unsafe.Pointer(&a1_.field_10[a1_.field_42])) = uint32(uintptr(unsafe.Pointer(m.sub_4BD470(unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.externs.Dword_5d4594_1045424))), int(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1_.field_9)) + uint32(a2*2) + 128))))))))
	v2 = int32(a1_.field_42)
	result = int32(*(*uint32)(unsafe.Pointer(&a1_.field_10[v2])))
	if result != 0 {
		m.sub_4BD650(int32(*(*uint32)(unsafe.Pointer(&a1_.field_10[v2]))))
		result = int32(a1_.field_42 + 1)
		a1_.field_42 = uint32(result)
	}
	return result
}

func (m *AudioModule) sub_451F90(a1_ *Struct576) int32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v1 int32
	var result int32
	var v3 *int32
	v1 = 0
	result = int32(a1_.field_42)
	if result <= 0 {
		a1_.field_42 = 0
	} else {
		v3 = (*int32)(unsafe.Pointer(&a1_.field_10[0]))
		for {
			m.sub_4BD660(*v3)
			*v3 = 0
			result = int32(a1_.field_42)
			v1++
			v3 = (*int32)(unsafe.Add(unsafe.Pointer(v3), 4*1))
			if v1 >= result {
				break
			}
		}
		a1_.field_42 = 0
	}
	return result
}

func (m *AudioModule) sub_451FE0(a1_ *Struct576) int32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	m.nox_common_list_remove_425920(unsafe.Pointer(a1_))
	a1_.field_70 = 0
	return int32(m.sub_4BD300(unsafe.Pointer(*(**uint32)(unsafe.Pointer(m.externs.Dword_5d4594_1045436))), int(uintptr(unsafe.Pointer(a1_)))))
}

func (m *AudioModule) sub_452120(a1p *Struct576) *int32 {
	var v1 int32
	var result *int32
	var v3 *int32
	var v4 *uint8
	var v5 *uint8
	v1 = 0
	result = (*int32)(unsafe.Pointer(m.sub_4521A0(int32(a1p.field_75 + a1p.field_9.field_12))))
	v3 = result
	if result != nil {
		m.sub_452190((*Struct200)(unsafe.Pointer(result)))
		v4 = *(**uint8)(unsafe.Pointer(m.externs.ListHead_5d4594_840612))
		if unsafe.Pointer(*(**uint8)(unsafe.Pointer(m.externs.ListHead_5d4594_840612))) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
			for {
				v5 = *(**uint8)(unsafe.Pointer(v4))
				if *((**int32)(unsafe.Add(unsafe.Pointer((**int32)(unsafe.Pointer(v4))), unsafe.Sizeof((*int32)(nil))*9))) == v3 {
					m.Sub_4523D0((*Struct576)(unsafe.Pointer(v4)))
					m.sub_451FE0((*Struct576)(unsafe.Pointer(v4)))
					v1 = 1
				}
				v4 = v5
				if unsafe.Pointer(v5) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
					break
				}
			}
		}
		result = (*int32)(unsafe.Pointer(uintptr(v1)))
	}
	return result
}

func (m *AudioModule) Sub_4523D0(a1p_ *Struct576) int32 {
	var (
		a1p    unsafe.Pointer = unsafe.Pointer(a1p_)
		a1     *uint32        = (*uint32)(a1p)
		result int32          = 0
	)
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*6)) & 1) == 0 {
		m.sub_452410((*Struct576)(unsafe.Pointer(a1)))
		m.sub_451F90((*Struct576)(unsafe.Pointer(a1)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*7)) = 4
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*70)) = 0
		result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*6)))
		*((*uint8)(unsafe.Pointer(&result))) = uint8(int8(result | 1))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*6)) = uint32(result)
	}
	return result
}

func (m *AudioModule) sub_452410(a1p *Struct576) {
	var result *Struct312
	result = a1p.field_44
	if result != nil && a1p == result.field_38 {
		if int32(a1p.field_6)&2 != 0 {
			m.sub_4BDA80(a1p.field_44)
		}
		m.Sub_4BDB30(a1p.field_44)
		a1p.field_44.field_38 = nil
		a1p.field_44.field_37 = nil
		a1p.field_44.field_35 = nil
		a1p.field_44.field_36 = nil
		a1p.field_44.field_28 = nil
		a1p.field_44 = nil
	}
}

func (m *AudioModule) sub_452490(a1p *Struct576) int32 {
	// var v1 int32
	var v3 int32
	var v4 int32
	v1 := a1p.field_44
	if a1p != v1.field_38 {
		return 0
	}
	v3 = int32(a1p.field_74)
	m.sub_4BDB90(v1, unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a1p.field_74)))))
	a1p.field_7 = 3
	v4 = int32(a1p.field_6)
	*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4 | 2))
	a1p.field_6 = uint8(int8(v4))
	a1p.field_74 = 0
	if m.sub_4BDB40(a1p.field_44) == 0 {
		return 1
	}
	a1p.field_7 = 1
	a1p.field_74 = uint32(v3)
	a1p.field_6 &= uint8(0xFD)
	return 0
}

func (m *AudioModule) sub_452510(a3_ *Struct576) {
	var a3 int32 = int32(uintptr(unsafe.Pointer(a3_)))
	_ = a3
	var v1 int32
	var v2 int32
	if *m.externs.Dword_587000_126996 == 0 {
		a3_.field_7 = 4
	}
	for {
		v1 = int32(a3_.field_7)
		if v1 == 0 {
			break
		}
		v2 = v1 - 2
		if v2 != 0 {
			var v3 uint32 = uint32(v2 - 2)
			if v3 == 0 {
				m.Sub_4523D0(a3_)
			}
			return
		}
		if uint64(m.nox_platform_get_ticks()) <= a3_.field_72 {
			return
		}
		a3_.field_7 = a3_.field_8
	}
	if m.sub_452580(a3_) == 0 {
		m.Sub_4523D0(a3_)
	}
}

func (m *AudioModule) sub_452690(a3_ *Struct576, a4 int64, a5 int32) int64 {
	var result int64
	a3_.field_8 = uint32(a5)
	result = a4 + int64(m.nox_platform_get_ticks())
	a3_.field_72 = uint64(result)
	a3_.field_7 = 2
	return result
}

func (m *AudioModule) Nox_xxx_draw_452300(a1p *Struct200) *uint32 {
	if *m.externs.Dword_5d4594_1045432 == 0 {
		return nil
	}
	if *m.externs.Dword_587000_126996 == 0 {
		return nil
	}
	if a1p.field_0 == nil {
		return nil
	}
	var v1p *Struct576 = (*Struct576)(unsafe.Pointer(m.sub_4BD2E0(unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.externs.Dword_5d4594_1045436))))))
	// var v1 *uint32 = (*uint32)(unsafe.Pointer(v1p))
	if v1p == nil {
		m.sub_452230()
		v1p = (*Struct576)(m.sub_4BD2E0(unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.externs.Dword_5d4594_1045436)))))
		if v1p == nil {
			return nil
		}
	}
	alloc.Memset(unsafe.Pointer(v1p), 0, 0x240)
	v1p.field_9 = a1p
	m.sub_425770(unsafe.Pointer(v1p))
	v1p.field_7 = 0
	v1p.field_75 = 0
	v1p.field_142 = 0
	v1p.field_108 = 0
	v1p.field_42 = 0
	v1p.timerGroup_46.Init()
	m.nox_common_list_append_4258E0((unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(m.externs.ListHead_5d4594_840612)))))), (unsafe.Pointer(v1p)))
	v1p.field_70 = func() uint32 {
		p_ := memmap.PtrUint32(0x587000, 127000)
		x := *p_
		*p_++
		return x
	}()
	return (*uint32)(unsafe.Pointer(v1p))
}

func (m *AudioModule) Sub_452E90(a1 *uint32, a2_ *Struct576) int32 {
	var a2 int32 = int32(uintptr(unsafe.Pointer(a2_)))
	_ = a2
	var result int32
	result = int32(uintptr(unsafe.Pointer(a2_)))
	*a1 = uint32(uintptr(unsafe.Pointer(a2_)))
	if a2_ != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1)) = a2_.field_70
		result = int32(*(*uint32)(unsafe.Pointer(&a2_.field_9)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)) = uint32(result)
	}
	return result
}

func (m *AudioModule) Sub_452EE0(a1_ *Struct576, a2 int32) int32 {
	var v2 int32 = int32(m.sub_452F10(a1_, a2))
	m.sub_486320(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1_.timerGroup_46))), int(v2))
	return int32(m.sub_4863B0(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1_.timerGroup_46)))))
}

func (m *AudioModule) sub_452F10(a1_ *Struct576, a2 int32) uint32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int32
	v2 = a2
	if a2 <= 100 {
		if a2 < 0 {
			v2 = 0
		}
	} else {
		v2 = 100
	}
	return (uint32(v2*163) * (*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1_.field_9)) + 20))) >> 16)) >> 14
}

func (m *AudioModule) Sub_452F50(a1_ *Struct576, a2 int32) int32 {
	var v2 int32 = int32(m.sub_452F10(a1_, a2))
	return int32(m.sub_486350(unsafe.Pointer(&a1_.timerGroup_46), int(v2)))
}

func (m *AudioModule) Sub_452F80(a1_ *Struct576, a2 int32) *uint32 {
	var v2 int32 = m.sub_452FA0(a2)
	return (*uint32)(m.sub_486320(unsafe.Pointer(&a1_.timerGroup_46.Timers[2]), int(v2)))
}

func (m *AudioModule) sub_451CF0(a1 *uint32) int32 {
	var (
		v1     int32
		result int32
		v3     int32
		v4     int32
		v5     int32
		v6     int32
		v7     int32
		v8     *uint32
		v9     int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*9)))
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*108)))
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
	if result != 0 {
		if v3&2 != 0 {
			v5 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(result)-1, unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 376))
			v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*108)) - 1)
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*43)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*uintptr(v5+76)))
			v7 = v5
			if v5 < v6 {
				v8 = (*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*uintptr(v5+76)))
				for {
					v7++
					*v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
					v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
					if uint32(v7) >= *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*108))-1 {
						break
					}
				}
			}
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*43))++
		}
		v9 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*43)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*108))--
		result = m.sub_4BD710(int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*uintptr(v9+10)))))
	} else if v3&1 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 60))) != 0 && (func() bool {
			v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*109)) + 1)
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*109)) = uint32(v4)
			return v4 >= *(*int32)(unsafe.Pointer(uintptr(v1 + 60)))
		}()) {
			result = 0
		} else {
			result = m.sub_451CA0((*Struct576)(unsafe.Pointer(a1)))
		}
	}
	return result
}

func (m *AudioModule) sub_451DC0(a1 int32) int32 {
	var (
		v1     *uint32
		result int32
		v3     int32
		i      int32
		v5     int32
		v6     int32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 36)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))))
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
	if result != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*17)) < 0x21 {
			return result
		}
		m.sub_451F90((*Struct576)(unsafe.Pointer(uintptr(a1))))
	}
	if v3&4 != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*17)) >= 0x21 {
			v5 = m.sub_451E80(a1)
			result = m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(a1))), v5)
		} else {
			result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*48)))
			for i = 0; i < result; i++ {
				m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(a1))), i)
				result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*48)))
			}
		}
	} else if v3&2 != 0 {
		v6 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*48))-1), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 536))
		result = m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(a1))), v6)
	} else {
		result = m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(a1))), 0)
	}
	return result
}

func (m *AudioModule) sub_452580(a1_ *Struct576) int32 {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	_ = a1
	var v1 int32
	var res *Struct312
	var v3 int32
	var v4 int32
	var v5 int32
	var ret int32 = 0
	v1 = int32(uintptr(unsafe.Pointer(a1_.field_9)))
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 192))) == 0 {
		return 0
	}
	v3 = int32(a1_.field_75)
	a1_.field_109 = 0
	res = (*Struct312)(unsafe.Pointer(m.sub_452810(int(int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 48)))+uint32(v3))), 0)))
	a1_.field_44 = res
	if res != nil {
		v4 = int32(m.nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 76)))), int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 80)))), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1482))
		a1_.field_44.timerGroup_4.Timers[1].SetRaw(uint32(v4 + 100))
		m.sub_4BDB20(a1_.field_44)
		a1_.field_44.field_38 = a1_
		a1_.field_44.field_35 = m.externs.Sub_452770_ptr
		a1_.field_44.field_36 = m.externs.Sub_4526F0_ptr
		a1_.field_44.field_37 = m.externs.Sub_4526D0_ptr
		a1_.field_7 = 1
		a1_.field_44.field_28 = &a1_.timerGroup_46
		if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 4))))&8 != 0 {
			v5 = int32(m.nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 68)))), int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 72)))), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1497))
			if v5 > 33 {
				m.sub_452690(a1_, int64(v5), 1)
			}
		}
		ret = 1
	}
	return ret
}

func (m *AudioModule) sub_451E80(a1 int32) int32 {
	var (
		v1  int32
		v2  int32
		v3  int32
		v4  int32
		v5  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 *uint32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
	if *(*int32)(unsafe.Pointer(uintptr(a1 + 568))) <= 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 192))))
		v4 = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))) = uint32(v3)
		if v3 > 0 {
			v5 = a1 + 440
			for {
				v5 += 4
				v6 = v3 - func() int32 {
					p_ := &v4
					x := *p_
					*p_++
					return x
				}() - 1
				*(*uint32)(unsafe.Pointer(uintptr(v5 - 4))) = uint32(v6)
				v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))))
				if v4 >= v3 {
					break
				}
			}
		}
	}
	v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))) = uint32(v7)
	if (v2 & 2) == 0 {
		return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + v7*4 + 440))))
	}
	v8 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(v7), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 431))
	v9 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + v8*4 + 440))))
	v10 = v8
	if v8 < *(*int32)(unsafe.Pointer(uintptr(a1 + 568))) {
		v11 = (*uint32)(unsafe.Pointer(uintptr(a1 + v8*4 + 440)))
		for {
			v10++
			*v11 = *(*uint32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
			v11 = (*uint32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
			if v10 >= *(*int32)(unsafe.Pointer(uintptr(a1 + 568))) {
				break
			}
		}
	}
	return v9
}

func (m *AudioModule) Sub_452770(a1 *Struct312) int32 {
	var (
		v1 *uint32
		v2 *uint32
		v4 int32
		v5 uint32
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*38)))))
	v2 = (*uint32)(unsafe.Pointer(uintptr(m.sub_451CF0((*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*38)))))))))
	if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*9)) + 72))) < 0x21 {
		m.sub_4BDB90(a1, unsafe.Pointer(v2))
		return 0
	}
	m.sub_4BDB90(a1, nil)
	v4 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*9)))
	if (int32(*(*uint8)(unsafe.Pointer(uintptr(v4 + 4))))&8) == 0 || v2 != nil || *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*142)) != 0 {
		v5 = uint32(m.nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(uintptr(v4 + 68)))), int(*(*uint32)(unsafe.Pointer(uintptr(v4 + 72)))), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 706))
		if v5 < 0x21 {
			m.sub_4BDB90(a1, unsafe.Pointer(v2))
			return 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*71)) = v5
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*74)) = uint32(uintptr(unsafe.Pointer(v2)))
	}
	return 0
}

func (m *AudioModule) sub_452050(a1p *Struct576) {
	var v1 *Struct200
	var v2 int32
	var v3 uint32
	v1 = a1p.field_9
	v2 = int32(v1.field_12 + a1p.field_75)
	v3 = (a1p.timerGroup_46.Timers[0].Current >> 16) / 0x666
	heads := m.externs.ListHeads_5d4594_839892
	if v1.field_26 == *memmap.PtrUint32(0x5D4594, 1045444) {
		result := v1.field_27
		if v2 <= result {
			if v2 == result && v3 > v1.field_31 {
				v1.field_31 = v3
				v7 := (unsafe.Pointer((&v1.field_28)))
				m.nox_common_list_remove_425920(unsafe.Pointer(v7))
				m.nox_common_list_append_4258E0((unsafe.Pointer(&heads[v2][v3])), (unsafe.Pointer(v7)))
			}
		} else {
			v1.field_27 = v2
			v1.field_31 = v3
			v6 := (unsafe.Pointer((&v1.field_28)))
			m.nox_common_list_remove_425920(unsafe.Pointer(v6))
			m.nox_common_list_append_4258E0((unsafe.Pointer(&heads[v2][v3])), (unsafe.Pointer(v6)))
		}
	} else {
		v1.field_26 = *memmap.PtrUint32(0x5D4594, 1045444)
		v1.field_27 = v2
		v1.field_31 = v3
		v8 := unsafe.Pointer(&v1.field_28)
		m.sub_425770(unsafe.Pointer(v8))
		m.nox_common_list_append_4258E0((unsafe.Pointer(&heads[v2][v3])), (unsafe.Pointer(v8)))
	}
}

func (m *AudioModule) sub_451BE0(a1_ *Struct576) int32 {
	var v1 *Struct576
	var v3 uint32
	var v5 int32
	var v6 int32
	var v7 *uint32
	var result int32
	var v9 int32
	var v10 *uint32
	v1 = a1_
	v2p := a1_.field_9
	v3 = a1_.timerGroup_46.Timers[0].Current >> 16
	v4 := v2p.field_22.next
	if v4 != &v2p.field_22 {
		for {
			v5 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*44)) >> 16) - v3)
			if v5 < 0 {
				v5 = int32(v3 - (*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*44)) >> 16))
			}
			if uint32(v5) >= (*(*uint32)(unsafe.Pointer(&v2p.field_4.Current))>>16)/10 {
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*44))>>16 < v3 {
					break
				}
			} else {
				v6 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*4)))
				if int32(*(*uint8)(unsafe.Pointer(&v2p.field_1)))&0x10 != 0 {
					if v6 != 0 {
						break
					}
				} else if v6 == 0 {
					break
				}
			}
			v4 = v4.next
			if v4 == &v2p.field_22 {
				break
			}
		}
		v1 = a1_
	}
	v7 = (*uint32)(unsafe.Pointer(&v1.field_3))
	m.sub_425770(unsafe.Pointer(&v1.field_3))
	m.nox_common_list_append_4258E0((unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v4)))))), (unsafe.Pointer(v7)))
	result = int32(*(*uint32)(unsafe.Pointer(&v2p.field_14)))
	v9 = int32(*(*uint32)(unsafe.Pointer(&v2p.field_13)) + 1)
	*(*uint32)(unsafe.Pointer(&v2p.field_13)) = uint32(v9)
	if result != 0 {
		if v9 > result {
			v10 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&v2p.field_22.prev)) - 12)))
			m.nox_common_list_remove_425920(unsafe.Pointer(*(***uint32)(unsafe.Pointer(&v2p.field_22.prev))))
			m.Sub_4523D0((*Struct576)(unsafe.Pointer(v10)))
			result = int32(*(*uint32)(unsafe.Pointer(&v2p.field_13)) - 1)
			*(*uint32)(unsafe.Pointer(&v2p.field_13)) = uint32(result)
		}
	}
	return result
}

func (m *AudioModule) Sub_451970() {
	m.sub_4521F0()
	m.sub_452230()
	if *m.externs.Dword_5d4594_1045424 != 0 {
		m.sub_4BD3C0(*(*unsafe.Pointer)(unsafe.Pointer(m.externs.Dword_5d4594_1045424)))
		*m.externs.Dword_5d4594_1045424 = 0
	}
	if *m.externs.Dword_5d4594_1045436 != 0 {
		m.sub_4BD2D0(*(*unsafe.Pointer)(unsafe.Pointer(m.externs.Dword_5d4594_1045436)))
		*m.externs.Dword_5d4594_1045436 = 0
	}
	*m.externs.Dword_5d4594_1045432 = 0
}

func (m *AudioModule) Sub_4519C0() {
	var (
		result int32
		v1     int32
		v2     int32
		v3     int32
		v4     int32
		v5     *uint8
		v6     *uint8
		v7     *uint8
		v8     int32
		v9     int32
		v10    int32
	)
	result = int32(*m.externs.Dword_5d4594_1045432)
	if *m.externs.Dword_5d4594_1045432 == 0 {
		return
	}
	result = int32(*memmap.PtrUint32(0x5D4594, 1045448))
	if *memmap.PtrUint32(0x5D4594, 1045448) != 0 {
		return
	}
	*memmap.PtrUint32(0x5D4594, 1045448) = 1
	(*m.externs.Ptr_TimerGroup_587000_127004).Update()
	v1 = int32(*memmap.PtrUint32(0x5D4594, 840612))
	*memmap.PtrUint32(0x5D4594, 1045440)++
	if unsafe.Pointer(*(**uint8)(unsafe.Pointer(m.externs.ListHead_5d4594_840612))) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
		for {
			v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
			if *(*uint32)(unsafe.Pointer(uintptr(v2 + 100))) != *memmap.PtrUint32(0x5D4594, 1045440) {
				m.nox_common_list_clear_425760((unsafe.Pointer(uintptr(v2 + 88))))
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 52))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 100))) = *memmap.PtrUint32(0x5D4594, 1045440)
			}
			m.sub_486520(unsafe.Pointer(uintptr(v1 + 184)))
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 28))) != 4 {
				m.sub_451BE0((*Struct576)(unsafe.Pointer(uintptr(v1))))
			}
			v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1))))
			if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
				break
			}
		}
		v1 = int32(*memmap.PtrUint32(0x5D4594, 840612))
		if unsafe.Pointer(*(**uint8)(unsafe.Pointer(m.externs.ListHead_5d4594_840612))) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
			for {
				m.sub_452510((*Struct576)(unsafe.Pointer(uintptr(v1))))
				v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1))))
				if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
					break
				}
			}
			v1 = int32(*memmap.PtrUint32(0x5D4594, 840612))
		}
	}
	v3 = 0
	m.sub_452010()
	if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
		for {
			v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 176))))
			v5 = *(**uint8)(unsafe.Pointer(uintptr(v1)))
			if v4 == 0 || uint32(v1) != *(*uint32)(unsafe.Pointer(uintptr(v4 + 152))) {
				m.Sub_4523D0((*Struct576)(unsafe.Pointer(uintptr(v1))))
			}
			if int32(*(*uint8)(unsafe.Pointer(uintptr(v1 + 24))))&1 != 0 {
				m.sub_451FE0((*Struct576)(unsafe.Pointer(uintptr(v1))))
			} else {
				v3 += int32(((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 20))) >> 16) * 33) >> 14)
				m.sub_452050((*Struct576)(unsafe.Pointer(uintptr(v1))))
			}
			v1 = int32(uintptr(unsafe.Pointer(v5)))
			if unsafe.Pointer(v5) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
				break
			}
		}
	}
	if v3 <= 100 {
		m.externs.TimerGroup_5d4594_1045228.Timers[0].SetInterp(0x4000)
	} else {
		m.externs.TimerGroup_5d4594_1045228.Timers[0].SetInterp(0x4000 * 100 / uint32(v3))
	}
	m.externs.TimerGroup_5d4594_1045228.Update()
	v6 = *(**uint8)(unsafe.Pointer(m.externs.ListHead_5d4594_840612))
	if unsafe.Pointer(*(**uint8)(unsafe.Pointer(m.externs.ListHead_5d4594_840612))) != unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
		for {
			v7 = *(**uint8)(unsafe.Pointer(v6))
			result = int32(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), 4*7))))
			if result == 1 {
				m.sub_451DC0(int32(uintptr(unsafe.Pointer(v6))))
				v8 = m.sub_451CA0((*Struct576)(unsafe.Pointer(v6)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), 4*74))) = uint32(v8)
				if v8 == 0 {
					for {
						if m.sub_452120((*Struct576)(unsafe.Pointer(v6))) == nil {
							break
						}
						v7 = *(**uint8)(unsafe.Pointer(v6))
						m.sub_451DC0(int32(uintptr(unsafe.Pointer(v6))))
						v9 = m.sub_451CA0((*Struct576)(unsafe.Pointer(v6)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), 4*74))) = uint32(v9)
						if v9 != 0 {
							break
						}
					}
				}
				v10 = m.sub_451CA0((*Struct576)(unsafe.Pointer(v6)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), 4*74))) = uint32(v10)
				if v10 == 0 || (func() int32 {
					result = m.sub_452490((*Struct576)(unsafe.Pointer(v6)))
					return result
				}()) == 0 {
					m.Sub_4523D0((*Struct576)(unsafe.Pointer(v6)))
					result = m.sub_451FE0((*Struct576)(unsafe.Pointer(v6)))
				}
			}
			v6 = v7
			if unsafe.Pointer(v7) == unsafe.Pointer(m.externs.ListHead_5d4594_840612) {
				break
			}
		}
	}
	*memmap.PtrUint32(0x5D4594, 1045448) = 0
}

func (m *AudioModule) sub_4BDB20(a1p *Struct312) {
	a1p.field_31 |= 0x10
}

func (m *AudioModule) sub_4BD710(a1 int32) int32 {
	return a1 + 24
}

func (m *AudioModule) Sub_4526D0(a1 int32) int32 {
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 152))) + 28))) = 4
	return 0
}

func (m *AudioModule) Sub_452FE0(a1_ *Struct576, a2 int32) int32 {
	var v2 int32 = m.sub_452FA0(a2)
	return int32(m.sub_486350(unsafe.Pointer(&a1_.timerGroup_46.Timers[0]), int(v2)))
}

func (m *AudioModule) sub_452FA0(a1 int32) int32 {
	var v1 int32
	v1 = a1
	if a1 <= 50 {
		if a1 < -50 {
			v1 = -50
		}
	} else {
		v1 = 50
	}
	return (v1*8192)/50 + 8192
}

func (m *AudioModule) sub_4BD650(a1 int32) int32 {
	var result int32 = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))++
	return result
}

func (m *AudioModule) sub_4BD660(a1 int32) int32 {
	var result int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = 0
	}
	return result
}

func (m *AudioModule) Nox_xxx_clientPlaySoundSpecial_452D80(a1 int32, a2 int32) {
	var (
		result *uint32
		v3     *uint32
	)
	result = (*uint32)(unsafe.Pointer(m.Nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = m.Nox_xxx_draw_452300((*Struct200)(unsafe.Pointer(result)))
	v3 = result
	if result == nil {
		return
	}
	m.Sub_452EE0((*Struct576)(unsafe.Pointer(result)), a2)
	m.sub_452510((*Struct576)(unsafe.Pointer(v3)))
}

func (m *AudioModule) Sub_452DC0(a1 int32, a2 int32, a3 int32) {
	var (
		result *uint32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(m.Nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = m.Nox_xxx_draw_452300((*Struct200)(unsafe.Pointer(result)))
	v4 = result
	if result == nil {
		return
	}
	m.Sub_452EE0((*Struct576)(unsafe.Pointer(result)), a2)
	m.Sub_452F80((*Struct576)(unsafe.Pointer(v4)), a3)
	m.sub_452510((*Struct576)(unsafe.Pointer(v4)))
}

func (m *AudioModule) Sub_452E10(a1 int32, a2 int32, a3 int32) {
	var (
		result *uint32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(m.Nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = m.Nox_xxx_draw_452300((*Struct200)(unsafe.Pointer(result)))
	v4 = result
	if result == nil {
		return
	}
	m.Sub_452EE0((*Struct576)(unsafe.Pointer(result)), a2)
	m.Sub_452F80((*Struct576)(unsafe.Pointer(v4)), a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*75)) = 2
	m.sub_452510((*Struct576)(unsafe.Pointer(v4)))
}

func (m *AudioModule) Sub_4BDB30(a1 *Struct312) {
	a1.field_31 &= 0xFFFFFFEF
}
