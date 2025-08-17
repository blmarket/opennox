package audio

import (
	unsafe "unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

// Extracted struct definitions from defs.h
type Struct200 struct {
	field_0    unsafe.Pointer
	field_1    uint8
	field_1_1  uint8
	field_1_2  uint8
	field_1_3  uint8
	field_2    uint32
	field_3    uint32
	field_4    timer.Timer
	field_12   uint32
	field_13   uint32
	field_14   uint32
	field_15   int32
	field_16   uint32
	field_17   uint32
	field_18   uint32
	field_19   uint32
	field_20   uint32
	sndName_21 unsafe.Pointer // pointer to string
	field_22   ListElement[Struct576Field3, *Struct576Field3]
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
	ListElement[Struct200Field28, *Struct200Field28]
}

func (s *Struct200Field28) GetStruct200() *Struct200 {
	return (*Struct200)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - 28*4))
}

var _ = [1]struct{}{}[200-unsafe.Sizeof(Struct200{})]
var _ = [1]struct{}{}[unsafe.Sizeof(Struct200{})-200]

type Struct576Field3 struct {
	ListElement[Struct576Field3, *Struct576Field3]
}

func (s *Struct576Field3) GetStruct576() *Struct576 {
	return (*Struct576)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - 3*4))
}

type Struct576 struct {
	ListElement[Struct576, *Struct576]
	field_3       Struct576Field3
	field_6       uint8
	field_6_1     uint8
	field_6_2     uint8
	field_6_3     uint8
	field_7       uint32
	field_8       uint32
	field_9       *Struct200
	field_10      [32]unsafe.Pointer
	field_42      uint32
	field_43      int32
	field_44      *Struct312
	field_45      uint32
	timerGroup_46 timer.TimerGroup
	field_70      uint32
	field_71      uint32
	field_72      uint64
	field_74      uint32
	field_75      uint32
	field_76      [32]int32
	field_108     uint32
	field_109     int32
	field_110     [32]uint32
	field_142     uint32
	field_143     uint32
}

var _ = [1]struct{}{}[576-unsafe.Sizeof(Struct576{})]

func (m *AudioModule) Sub_451850(a2p *Struct264, a3p unsafe.Pointer) int32 {
	var (
		result int32
	)
	v4 := m.Externs.Struct200Arr_5d4594_840628
	for i := int32(0); i < 1023; i++ {
		m.sub_451920(&v4[i])
		v4[i].sndName_21 = m.nox_xxx_getSndName_40AF80(int(i))
	}
	*m.Externs.Dword_5d4594_1045420 = uint32(int32(uintptr(a3p)))
	*m.Externs.Dword_5d4594_1045428 = a2p
	if int32(uintptr(a3p)) != 0 {
		*m.Externs.Dword_5d4594_1045424 = uint32(uintptr(m.sub_4BD340(int(int32(uintptr(a3p))), 0x100000, 200, 0x2000)))
		*m.Externs.Dword_5d4594_1045436 = uint32(uintptr(m.sub_4BD280(200, 576)))
	}
	if *m.Externs.Dword_5d4594_1045424 == 0 || *m.Externs.Dword_5d4594_1045420 == 0 || *m.Externs.Dword_5d4594_1045428 == nil || *m.Externs.Dword_5d4594_1045436 == 0 {
		return 0
	}
	m.nox_common_list_clear_425760(unsafe.Pointer(m.Externs.ListHead_5d4594_840612))
	m.Externs.TimerGroup_5d4594_1045228.Init()
	result = 1
	(*m.Externs.Dword_5d4594_1045428).field_46 = m.Externs.TimerGroup_5d4594_1045228
	*m.Externs.Dword_5d4594_1045432 = 1
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
	heads := m.Externs.ListHeads_5d4594_839892
	for v1 := 0; v1 < 6; v1++ {
		for v2 := 0; v2 < 10; v2++ {
			heads[v1][v2].Clear_425760()
		}
	}
	(*m.Externs.Ptr_uint32_5d4594_1045444)++
	return int32(*m.Externs.Ptr_uint32_5d4594_1045444)
}

func (m *AudioModule) sub_452190(a1_ *Struct200) {
	m.nox_common_list_remove_425920(unsafe.Pointer(&a1_.field_28))
}

func (m *AudioModule) sub_4521A0(a1 int32) *Struct200 {
	heads := m.Externs.ListHeads_5d4594_839892
	if a1 > 0 {
		for v1 := int32(0); v1 < a1; v1++ {
			v2 := heads[v1]
			for v3 := 0; v3 < 10; v3++ {
				v4 := v2[v3]
				v5 := v4.FirstSafe_4258A0()
				if v5 != nil {
					return v5.PromoteUnsafe().GetStruct200()
				}
			}
		}
	}
	return nil
}

func (m *AudioModule) sub_4521F0() int32 {
	var (
		result int32
	)
	result = int32(*m.Externs.Dword_5d4594_1045432)
	if *m.Externs.Dword_5d4594_1045432 != 0 {
		v1 := m.Externs.ListHead_5d4594_840612.Next()
		if m.Externs.ListHead_5d4594_840612.Next() != m.Externs.ListHead_5d4594_840612 {
			for {
				v2 := v1.next
				m.Sub_4523D0(v1.PromoteUnsafe())
				result = m.sub_451FE0(v1.PromoteUnsafe())
				v1 = v2
				if v2 == m.Externs.ListHead_5d4594_840612 {
					break
				}
			}
		}
	}
	return result
}

func (m *AudioModule) sub_452230() {
	if *m.Externs.Dword_5d4594_1045432 == 0 {
		return
	}
	if unsafe.Pointer(m.Externs.ListHead_5d4594_840612.next) != unsafe.Pointer(m.Externs.ListHead_5d4594_840612) {
		res := m.Externs.ListHead_5d4594_840612.Next().PromoteUnsafe()
		for {
			v1 := res.next
			if res.field_6&1 != 0 {
				m.sub_451FE0(res)
			}
			if v1 == m.Externs.ListHead_5d4594_840612 {
				break
			}
			res = v1.PromoteUnsafe()
		}
	}
}

func (m *AudioModule) Nox_xxx_draw_452270(a1 int32) *Struct200 {
	if *m.Externs.Dword_5d4594_1045432 != 0 && a1 >= 0 && a1 < 1023 {
		return &m.Externs.Struct200Arr_5d4594_840628[a1]
	}
	return nil
}

func (m *AudioModule) Sub_4526F0(a1p *Struct312) int32 {
	var (
		v2 int32
	)
	v1p := a1p.field_38
	v1p.field_6 &= 0xFD
	v2 = 4
	if v1p.field_7 != 4 {
		if v1p.field_74 != 0 || v1p.field_142 != 0 {
			v2 = 1
		} else {
			v1p.field_71 = 0
		}
		if v1p.field_71 != 0 {
			m.sub_452690(v1p, int64(v1p.field_71), v2)
			v1p.field_71 = 0
			return 0
		}
		v1p.field_7 = uint32(v2)
	}
	return 0
}

func (m *AudioModule) sub_451CA0(a1p *Struct576) int32 {
	var v1 int32
	var v3 int32
	v1 = int32(a1p.field_42)
	a1p.field_108 = uint32(v1)
	if v1 == 0 {
		return 0
	}

	v3 = 0
	for v3 = 0; v3 < v1; v3++ {
		a1p.field_76[v3] = v3
	}
	a1p.field_43 = -1
	return m.sub_451CF0(a1p)
}

func (m *AudioModule) sub_451F30(a1p *Struct576, a2 int32) int32 {
	var v2 int32
	var result int32
	a1p.field_10[a1p.field_42] = m.sub_4BD470(
		unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045424))),
		int(a1p.field_9.field_32[a2]),
	)
	v2 = int32(a1p.field_42)
	result = int32(*(*uint32)(unsafe.Pointer(&a1p.field_10[v2])))
	if result != 0 {
		m.sub_4BD650(int32(*(*uint32)(unsafe.Pointer(&a1p.field_10[v2]))))
		result = int32(a1p.field_42 + 1)
		a1p.field_42 = uint32(result)
	}
	return result
}

func (m *AudioModule) sub_451F90(a1p *Struct576) int32 {
	result := int32(a1p.field_42)
	if result <= 0 {
		a1p.field_42 = 0
	} else {
		for i := int32(0); i < result; i++ {
			m.sub_4BD660(a1p.field_10[i])
			a1p.field_10[i] = nil
		}
		a1p.field_42 = 0
	}
	return result
}

func (m *AudioModule) sub_451FE0(a1_ *Struct576) int32 {
	m.nox_common_list_remove_425920(unsafe.Pointer(a1_))
	a1_.field_70 = 0
	return int32(m.sub_4BD300(unsafe.Pointer(*(**uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436))), int(uintptr(unsafe.Pointer(a1_)))))
}

func (m *AudioModule) sub_452120(a1p *Struct576) bool {
	var ret bool = false
	v3 := m.sub_4521A0(int32(a1p.field_75 + a1p.field_9.field_12))
	if v3 == nil {
		return false
	}
	m.sub_452190((*Struct200)(unsafe.Pointer(v3)))
	v4 := m.Externs.ListHead_5d4594_840612.next.PromoteUnsafe()
	if unsafe.Pointer(*(**uint8)(unsafe.Pointer(m.Externs.ListHead_5d4594_840612))) != unsafe.Pointer(m.Externs.ListHead_5d4594_840612) {
		for {
			v5 := v4.next
			if v4.field_9 == v3 {
				m.Sub_4523D0((*Struct576)(unsafe.Pointer(v4)))
				m.sub_451FE0((*Struct576)(unsafe.Pointer(v4)))
				ret = true
			}
			v4 = v5.PromoteUnsafe()
			if unsafe.Pointer(v5) == unsafe.Pointer(m.Externs.ListHead_5d4594_840612) {
				break
			}
		}
	}
	return ret
}

func (m *AudioModule) Sub_4523D0(a1 *Struct576) int32 {
	var (
		result int32 = 0
	)
	if (*(*uint32)(unsafe.Pointer(&a1.field_6)) & 1) == 0 {
		m.sub_452410((*Struct576)(unsafe.Pointer(a1)))
		m.sub_451F90((*Struct576)(unsafe.Pointer(a1)))
		*(*uint32)(unsafe.Pointer(&a1.field_7)) = 4
		*(*uint32)(unsafe.Pointer(&a1.field_70)) = 0
		result = int32(*(*uint32)(unsafe.Pointer(&a1.field_6)))
		*((*uint8)(unsafe.Pointer(&result))) = uint8(int8(result | 1))
		*(*uint32)(unsafe.Pointer(&a1.field_6)) = uint32(result)
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
	m.Sub_4BDB90(v1, unsafe.Pointer((*uint32)(unsafe.Pointer(uintptr(a1p.field_74)))))
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
	var v1 int32
	var v2 int32
	if *m.Externs.Dword_587000_126996 == 0 {
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

func (m *AudioModule) sub_452690(a3p *Struct576, a4 int64, a5 int32) int64 {
	var result int64
	a3p.field_8 = uint32(a5)
	result = a4 + int64(m.nox_platform_get_ticks())
	a3p.field_72 = uint64(result)
	a3p.field_7 = 2
	return result
}

func (m *AudioModule) Nox_xxx_draw_452300(a1p *Struct200) *Struct576 {
	if *m.Externs.Dword_5d4594_1045432 == 0 {
		return nil
	}
	if *m.Externs.Dword_587000_126996 == 0 {
		return nil
	}
	if a1p.field_0 == nil {
		return nil
	}
	var v1p *Struct576 = (*Struct576)(unsafe.Pointer(m.sub_4BD2E0(unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436))))))
	if v1p == nil {
		m.sub_452230()
		v1p = (*Struct576)(m.sub_4BD2E0(unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)))))
		if v1p == nil {
			return nil
		}
	}
	alloc.Memset(unsafe.Pointer(v1p), 0, 0x240)
	v1p.field_9 = a1p
	v1p.Init_425770()
	v1p.field_7 = 0
	v1p.field_75 = 0
	v1p.field_142 = 0
	v1p.field_108 = 0
	v1p.field_42 = 0
	v1p.timerGroup_46.Init()
	m.Externs.ListHead_5d4594_840612.Append_4258E0(&v1p.ListElement)
	v1p.field_70 = func() uint32 {
		p_ := memmap.PtrUint32(0x587000, 127000)
		x := *p_
		*p_++
		return x
	}()
	return v1p
}

func (m *AudioModule) Sub_452E90(a1 *uint32, a2_ *Struct576) int32 {
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

func (m *AudioModule) Sub_452EE0(a1p *Struct576, a2 int32) int32 {
	a1p.timerGroup_46.Timers[0].SetRaw(uint32(int32(m.sub_452F10(a1p, a2))))
	return bool2int32(a1p.timerGroup_46.Timers[0].Update())
}

func (m *AudioModule) sub_452F10(a1_ *Struct576, a2 int32) uint32 {
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

func (m *AudioModule) Sub_452F50(a1p *Struct576, a2 int32) int32 {
	a1p.timerGroup_46.Timers[0].SetInterp(uint32(int32(m.sub_452F10(a1p, a2))))
	return 0
}

func (m *AudioModule) Sub_452F80(a1p *Struct576, a2 int32) {
	var v2 int32 = m.sub_452FA0(a2)
	a1p.timerGroup_46.Timers[2].SetRaw(uint32(v2))
}

func (m *AudioModule) sub_451CF0(a1 *Struct576) int32 {
	var (
		v1     *Struct200
		result int32
		// v3     int32
		v5 int32
		v6 int32
		v7 int32
		v8 *uint32
		v9 int32
	)
	v1 = a1.field_9
	result = int32(a1.field_108)
	v3 := v1.field_1
	if result != 0 {
		if v3&2 != 0 {
			v5 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(result)-1, unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 376))
			v6 = int32(a1.field_108 - 1)
			a1.field_43 = a1.field_76[v5]
			v7 = v5
			if v5 < v6 {
				v8 = (*uint32)(unsafe.Pointer((&a1.field_76[v5])))
				for {
					v7++
					*v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
					v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
					if uint32(v7) >= *(*uint32)(unsafe.Pointer(&a1.field_108))-1 {
						break
					}
				}
			}
		} else {
			*(*uint32)(unsafe.Pointer(&a1.field_43))++
		}
		v9 = int32(*(*uint32)(unsafe.Pointer(&a1.field_43)))
		*(*uint32)(unsafe.Pointer(&a1.field_108))--
		result = m.sub_4BD710(int32(*(*uint32)(unsafe.Pointer(&a1.field_10[v9]))))
	} else if v3&1 != 0 {
		if v1.field_15 != 0 && (func() bool {
			v4 := a1.field_109 + 1
			a1.field_109 = v4
			return v4 >= v1.field_15
		}()) {
			result = 0
		} else {
			result = m.sub_451CA0((*Struct576)(unsafe.Pointer(a1)))
		}
	}
	return result
}

func (m *AudioModule) sub_451DC0(a1p *Struct576) int32 {
	var (
		v1     *uint32
		result int32
		v3     int32
		i      int32
		v5     int32
		v6     int32
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(&a1p.field_9)))))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(&a1p.field_42))))))
	v3 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*1)))
	if result != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*17)) < 0x21 {
			return result
		}
		m.sub_451F90((*Struct576)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(a1p))))))
	}
	if v3&4 != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*17)) >= 0x21 {
			v5 = m.sub_451E80(a1p)
			result = m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(a1p))))), v5)
		} else {
			result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*48)))
			for i = 0; i < result; i++ {
				m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(a1p))))), i)
				result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*48)))
			}
		}
	} else if v3&2 != 0 {
		v6 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*48))-1), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 536))
		result = m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(a1p))))), v6)
	} else {
		result = m.sub_451F30((*Struct576)(unsafe.Pointer(uintptr(uintptr(unsafe.Pointer(a1p))))), 0)
	}
	return result
}

func (m *AudioModule) sub_452580(a1 *Struct576) int32 {
	var res *Struct312
	var v3 int32
	var v4 int32
	var v5 int32
	var ret int32 = 0
	v1p := a1.field_9
	if *(*uint32)(unsafe.Pointer(&v1p.field_48)) == 0 {
		return 0
	}
	v3 = int32(a1.field_75)
	a1.field_109 = 0
	res = (*Struct312)(unsafe.Pointer(m.sub_452810(int(int32(*(*uint32)(unsafe.Pointer(&v1p.field_12))+uint32(v3))), 0)))
	a1.field_44 = res
	if res != nil {
		v4 = int32(m.nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(&v1p.field_19))), int(*(*uint32)(unsafe.Pointer(&v1p.field_20))), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1482))
		a1.field_44.timerGroup_4.Timers[1].SetRaw(uint32(v4 + 100))
		m.sub_4BDB20(a1.field_44)
		a1.field_44.field_38 = a1
		a1.field_44.field_35 = m.Externs.Sub_452770_ptr
		a1.field_44.field_36 = m.Externs.Sub_4526F0_ptr
		a1.field_44.field_37 = m.Externs.Sub_4526D0_ptr
		a1.field_7 = 1
		a1.field_44.field_28 = &a1.timerGroup_46
		if int32(*(*uint8)(unsafe.Pointer(&v1p.field_1)))&8 != 0 {
			v5 = int32(m.nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(&v1p.field_17))), int(*(*uint32)(unsafe.Pointer(&v1p.field_18))), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1497))
			if v5 > 33 {
				m.sub_452690(a1, int64(v5), 1)
			}
		}
		ret = 1
	}
	return ret
}

func (m *AudioModule) sub_451E80(a1p *Struct576) int32 {
	var (
		v2 int32
		v3 int32
		// v4  int32
		v6  int32
		v7  int32
		v8  int32
		v9  int32
		v10 int32
		v11 *uint32
	)
	v1p := a1p.field_9
	v2 = int32(*(*uint32)(unsafe.Pointer(&v1p.field_1)))
	if *(*int32)(unsafe.Pointer(&a1p.field_142)) <= 0 {
		v3 = int32(*(*uint32)(unsafe.Pointer(&v1p.field_48)))
		*(*uint32)(unsafe.Pointer(&a1p.field_142)) = uint32(v3)
		if v3 > 0 {
			for i := int32(0); i < v3; i++ {
				v6 = v3 - i - 1
				a1p.field_110[i] = uint32(v6)
				v3 = int32(*(*uint32)(unsafe.Pointer(&a1p.field_142)))
			}
		}
	}
	v7 = int32(*(*uint32)(unsafe.Pointer(&a1p.field_142)) - 1)
	*(*uint32)(unsafe.Pointer(&a1p.field_142)) = uint32(v7)
	if (v2 & 2) == 0 {
		return int32(*(*uint32)(unsafe.Pointer(&a1p.field_110[v7])))
	}
	v8 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(v7), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 431))
	v9 = int32(*(*uint32)(unsafe.Pointer(&a1p.field_110[v8])))
	v10 = v8
	if v8 < *(*int32)(unsafe.Pointer(&a1p.field_142)) {
		v11 = (*uint32)(unsafe.Pointer(&a1p.field_110[v8]))
		for {
			v10++
			*v11 = *(*uint32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
			v11 = (*uint32)(unsafe.Add(unsafe.Pointer(v11), 4*1))
			if v10 >= *(*int32)(unsafe.Pointer(&a1p.field_142)) {
				break
			}
		}
	}
	return v9
}

func (m *AudioModule) Sub_452770(a1 *Struct312) int32 {
	var (
		v1 *Struct576
		v2 *uint32
		v5 uint32
	)
	v1 = a1.field_38
	v2 = (*uint32)(unsafe.Pointer(uintptr(m.sub_451CF0(a1.field_38))))
	if v1.field_9.field_18 < 0x21 {
		m.Sub_4BDB90(a1, unsafe.Pointer(v2))
		return 0
	}
	m.Sub_4BDB90(a1, nil)
	v4 := v1.field_9
	if (v4.field_1&8) == 0 || v2 != nil || *(*uint32)(unsafe.Pointer(&v1.field_142)) != 0 {
		v5 = uint32(m.nox_common_randomIntMinMax_415FF0(int(v4.field_17), int(v4.field_18), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 706))
		if v5 < 0x21 {
			m.Sub_4BDB90(a1, unsafe.Pointer(v2))
			return 0
		}
		*(*uint32)(unsafe.Pointer(&v1.field_71)) = v5
		*(*uint32)(unsafe.Pointer(&v1.field_74)) = uint32(uintptr(unsafe.Pointer(v2)))
	}
	return 0
}

func (m *AudioModule) sub_452050(a1p *Struct576) {
	var v2 int32
	var v3 uint32
	v1 := a1p.field_9
	v2 = int32(v1.field_12 + a1p.field_75)
	v3 = (a1p.timerGroup_46.Timers[0].Current >> 16) / 0x666
	heads := m.Externs.ListHeads_5d4594_839892
	if v1.field_26 == *m.Externs.Ptr_uint32_5d4594_1045444 {
		result := v1.field_27
		if v2 <= result {
			if v2 == result && v3 > v1.field_31 {
				v1.field_31 = v3
				v7 := &v1.field_28
				m.nox_common_list_remove_425920(unsafe.Pointer(v7))
				heads[v2][v3].Append_4258E0(&v7.ListElement)
			}
		} else {
			v1.field_27 = v2
			v1.field_31 = v3
			v6 := &v1.field_28
			m.nox_common_list_remove_425920(unsafe.Pointer(v6))
			heads[v2][v3].Append_4258E0(&v6.ListElement)
		}
	} else {
		v1.field_26 = *m.Externs.Ptr_uint32_5d4594_1045444
		v1.field_27 = v2
		v1.field_31 = v3
		v8 := &v1.field_28
		v8.Init_425770()
		heads[v2][v3].Append_4258E0(&v8.ListElement)
	}
}

func (m *AudioModule) sub_451BE0(a1 *Struct576) int32 {
	var v1 *Struct576
	var v3 uint32
	var v5 int32
	var v6 int32
	var result int32
	var v9 int32
	v1 = a1
	v2p := a1.field_9
	v3 = a1.timerGroup_46.Timers[0].Current >> 16
	v4 := v2p.field_22.next
	if v4 != &v2p.field_22 {
		for {
			v4p := v4.PromoteUnsafe().GetStruct576()
			v5 = int32((v4p.timerGroup_46.Timers[0].Current >> 16) - v3)
			if v5 < 0 {
				v5 = int32(v3 - (v4p.timerGroup_46.Timers[0].Current >> 16))
			}
			if uint32(v5) >= (v2p.field_4.Current>>16)/10 {
				if v4p.timerGroup_46.Timers[0].Current>>16 < v3 {
					break
				}
			} else {
				v6 = int32(*(*uint32)(&v4p.field_7))
				if int32(v2p.field_1)&0x10 != 0 {
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
		v1 = a1
	}
	v7 := &v1.field_3.ListElement
	v1.field_3.Init_425770()
	v4.Append_4258E0(v7)
	result = int32(v2p.field_14)
	v9 = int32(v2p.field_13 + 1)
	v2p.field_13 = uint32(v9)
	if result != 0 {
		if v9 > result {
			v10 := v2p.field_22.prev.PromoteUnsafe().GetStruct576()
			v2p.field_22.prev.Remove_425920()
			m.Sub_4523D0(v10)
			result = int32(v2p.field_13 - 1)
			v2p.field_13 = uint32(result)
		}
	}
	return result
}

func (m *AudioModule) Sub_451970() {
	m.sub_4521F0()
	m.sub_452230()
	if *m.Externs.Dword_5d4594_1045424 != 0 {
		m.sub_4BD3C0(*(*unsafe.Pointer)(unsafe.Pointer(m.Externs.Dword_5d4594_1045424)))
		*m.Externs.Dword_5d4594_1045424 = 0
	}
	if *m.Externs.Dword_5d4594_1045436 != 0 {
		m.sub_4BD2D0(*(*unsafe.Pointer)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)))
		*m.Externs.Dword_5d4594_1045436 = 0
	}
	*m.Externs.Dword_5d4594_1045432 = 0
}

func (m *AudioModule) Sub_4519C0() {
	var (
		result int32
		v3     int32
		// v7     *uint8
		v8  int32
		v9  int32
		v10 int32
	)
	result = int32(*m.Externs.Dword_5d4594_1045432)
	if *m.Externs.Dword_5d4594_1045432 == 0 {
		return
	}
	result = int32(*memmap.PtrUint32(0x5D4594, 1045448))
	if *memmap.PtrUint32(0x5D4594, 1045448) != 0 {
		return
	}
	*memmap.PtrUint32(0x5D4594, 1045448) = 1
	(*m.Externs.Ptr_TimerGroup_587000_127004).Update()
	// v1 = int32(*memmap.PtrUint32(0x5D4594, 840612))
	v1p := m.Externs.ListHead_5d4594_840612.next
	*memmap.PtrUint32(0x5D4594, 1045440)++
	if m.Externs.ListHead_5d4594_840612.next != m.Externs.ListHead_5d4594_840612 {
		for {
			v1pp := v1p.PromoteUnsafe()
			v2p := v1pp.field_9
			if v2p.field_25 != *memmap.PtrUint32(0x5D4594, 1045440) {
				v2p.field_22.Clear_425760()
				v2p.field_13 = 0
				v2p.field_25 = *memmap.PtrUint32(0x5D4594, 1045440)
			}

			v1pp.timerGroup_46.Update()
			if v1pp.field_7 != 4 {
				m.sub_451BE0(v1pp)
			}
			v1p = v1p.next
			if v1p == m.Externs.ListHead_5d4594_840612 {
				break
			}
		}
		v1p = m.Externs.ListHead_5d4594_840612.next
		if v1p != m.Externs.ListHead_5d4594_840612 {
			for {
				m.sub_452510(v1p.PromoteUnsafe())
				v1p = v1p.next
				if v1p == m.Externs.ListHead_5d4594_840612 {
					break
				}
			}
			v1p = m.Externs.ListHead_5d4594_840612.next
		}
	}
	v3 = 0
	m.sub_452010()
	if v1p != m.Externs.ListHead_5d4594_840612 {
		for {
			v1pp := v1p.PromoteUnsafe()
			v4 := v1pp.field_44
			// v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 176))))
			v5 := v1p.next
			if v4 == nil || v1pp != v4.field_38 {
				m.Sub_4523D0(v1pp)
			}
			if v1pp.field_6&1 != 0 {
				m.sub_451FE0(v1pp)
			} else {
				v3 += int32(((v1pp.field_9.field_4.Current >> 16) * 33) >> 14)
				m.sub_452050(v1pp)
			}
			v1p = v5
			if v5 == m.Externs.ListHead_5d4594_840612 {
				break
			}
		}
	}
	if v3 <= 100 {
		m.Externs.TimerGroup_5d4594_1045228.Timers[0].SetInterp(0x4000)
	} else {
		m.Externs.TimerGroup_5d4594_1045228.Timers[0].SetInterp(0x4000 * 100 / uint32(v3))
	}
	m.Externs.TimerGroup_5d4594_1045228.Update()

	v6x := m.Externs.ListHead_5d4594_840612.next
	v6 := v6x.PromoteUnsafe()
	if v6x != m.Externs.ListHead_5d4594_840612 {
		for {
			v7 := v6.next
			result = int32(*((*uint32)(unsafe.Pointer(&v6.field_7))))
			if result == 1 {
				m.sub_451DC0(v6)
				v8 = m.sub_451CA0(v6)
				*((*uint32)(unsafe.Pointer(&v6.field_74))) = uint32(v8)
				if v8 == 0 {
					for {
						if m.sub_452120(v6) == false {
							break
						}
						v7 = v6.next
						m.sub_451DC0(v6)
						v9 = m.sub_451CA0(v6)
						*((*uint32)(unsafe.Pointer(&v6.field_74))) = uint32(v9)
						if v9 != 0 {
							break
						}
					}
				}
				v10 = m.sub_451CA0(v6)
				*((*uint32)(unsafe.Pointer(&v6.field_74))) = uint32(v10)
				if v10 == 0 || (func() int32 {
					result = m.sub_452490(v6)
					return result
				}()) == 0 {
					m.Sub_4523D0(v6)
					result = m.sub_451FE0(v6)
				}
			}
			v6 = v7.PromoteUnsafe()
			if v7 == m.Externs.ListHead_5d4594_840612 {
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

func (m *AudioModule) Sub_452FE0(a1p *Struct576, a2 int32) int32 {
	a1p.timerGroup_46.Timers[0].SetInterp(uint32(m.sub_452FA0(a2)))
	return 0
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

func (m *AudioModule) sub_4BD660(a1p unsafe.Pointer) int32 {
	var result int32 = int32(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a1p)) + 12))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a1p)) + 12))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a1p)) + 12))) = 0
	}
	return result
}

func (m *AudioModule) Nox_xxx_clientPlaySoundSpecial_452D80(a1 int32, a2 int32) {
	var (
		result *uint32
	)
	result = (*uint32)(unsafe.Pointer(m.Nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	res := m.Nox_xxx_draw_452300((*Struct200)(unsafe.Pointer(result)))
	if res == nil {
		return
	}
	m.Sub_452EE0(res, a2)
	m.sub_452510(res)
}

func (m *AudioModule) Sub_452DC0(a1 int32, a2 int32, a3 int32) {
	res := m.Nox_xxx_draw_452270(a1)
	if res == nil {
		return
	}
	res2 := m.Nox_xxx_draw_452300(res)
	v4 := res2
	if res2 == nil {
		return
	}
	m.Sub_452EE0(res2, a2)
	m.Sub_452F80(v4, a3)
	m.sub_452510(v4)
}

func (m *AudioModule) Sub_452E10(a1 int32, a2 int32, a3 int32) {
	res1 := m.Nox_xxx_draw_452270(a1)
	if res1 == nil {
		return
	}
	res2 := m.Nox_xxx_draw_452300(res1)
	v4 := res2
	if res2 == nil {
		return
	}
	m.Sub_452EE0(res2, a2)
	m.Sub_452F80(v4, a3)
	v4.field_75 = 2
	m.sub_452510(v4)
}

func (m *AudioModule) Sub_4BDB30(a1 *Struct312) {
	a1.field_31 &= 0xFFFFFFEF
}

func (m *AudioModule) Sub_4BD720(a1p *Struct264) *Struct312 {
	var v1 *uint32
	v1pp, _ := alloc.Calloc(1, 0x138)
	v1p := (*Struct312)(v1pp)
	alloc.Memset(unsafe.Pointer(v1p), 0, 0x138)
	v1p.Init_425770()
	m.sub_4BDC00(int32(uintptr(unsafe.Pointer(&v1p.field_30))))
	v1p.timerGroup_44.Init()
	m.sub_4BD7C0(v1p)
	v1p.field_33 = a1p
	v1p.field_43 = a1p.field_64
	if ccall.CallIntPtr(a1p.field_64.field_1, unsafe.Pointer(v1p)) == 0 {
		return v1p
	}
	if v1 != nil {
		m.Sub_4BD7A0(v1p)
	}
	return nil
}

func (m *AudioModule) Sub_4BD7A0(a1 *Struct312) {
	ccall.CallVoidPtr(a1.field_43.field_2, unsafe.Pointer(a1))
	alloc.FreePtr(unsafe.Pointer(a1))
}

func (m *AudioModule) sub_4BD7C0(a1p *Struct312) {
	a1p.field_69 = m.Externs.Sub_4BD8C0_ptr
	a1p.field_70 = m.Externs.Sub_4BD940_ptr
	a1p.field_71 = m.Externs.Sub_4BD9B0_ptr
	a1p.field_34 = 0
	a1p.field_35 = nil
	a1p.field_36 = nil
	a1p.field_38 = nil
	a1p.field_3 = 1
	m.sub_4BDC00(int32(uintptr(unsafe.Pointer(&a1p.field_30))))
	a1p.field_30 = 0
	a1p.field_29 = *m.Externs.Ptr_TimerGroup_5d4594_1193340
	a1p.field_28 = nil
	a1p.timerGroup_4.Init()
	a1p.field_72 = nil
}

func (m *AudioModule) Sub_4BD8C0(a1 int32) int32 {
	var (
		v1     func(int32) int32
		result int32
		v3     int32
		v4     int32
	)
	v1 = *(*func(int32) int32)(unsafe.Pointer(uintptr(a1 + 136)))
	if v1 != nil {
		result = v1(a1)
		if result != 0 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 300))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 304))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 296))) = 0
			return result
		}
	} else {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 292))) != 0 {
			v3 = int32(uintptr(unsafe.Pointer(m.nox_common_list_getNext_425940((unsafe.Pointer(*(**int32)(unsafe.Pointer(uintptr(a1 + 292)))))))))
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 292))) = uint32(v3)
			if v3 != 0 {
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 296))) = *(*uint32)(unsafe.Pointer(uintptr(v3 + 12)))
				v4 = int32(*(*uint32)(unsafe.Pointer(uintptr(v3 + 16))))
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 300))) = uint32(v4)
				*(*uint32)(unsafe.Pointer(uintptr(a1 + 304))) = uint32(v4)
				return 0
			}
		}
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 300))) = 0
	}
	return 0
}

func (m *AudioModule) Sub_4BD940(a1p *Struct312) int32 {
	if *(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))) + 128))) != 0 {
		if *(*int32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))) + 128))) != -1 {
			*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))) + 128)))--
		}
		m.Sub_4BDB90(a1p, a1p.field_72)
	} else {
		m.Sub_4BDB90(a1p, nil)
	}
	v1 := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))) + 140)))
	if v1 != nil {
		ccall.CallVoidPtr(v1, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))))))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))) + 288))) != 0 {
		ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))) + 172))) + 36))), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))))))
	}
	return 0
}

func (m *AudioModule) Sub_4BD9B0(a2p *Struct312) int32 {
	var (
		result int32
	)
	a2p.field_72 = nil
	a2p.field_31 &= 0xFA
	a2p.field_32 = 0
	a2p.timerGroup_4.Init()
	v2 := a2p.field_36
	if v2 != nil {
		result = int32(ccall.CallIntPtr(v2, unsafe.Pointer((*uint32)(unsafe.Pointer(a2p)))))
	} else {
		result = 0
	}
	return result
}

func (m *AudioModule) sub_4BDC00(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = 0
	return result
}

func (m *AudioModule) Sub_4BDB90(a1p *Struct312, a2p unsafe.Pointer) {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1p))
	var a2 *uint32 = (*uint32)(a2p)
	var (
		v2 int32
		v3 int32
		v4 *uint32
		v5 int32
	)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*72)) = uint32(uintptr(unsafe.Pointer(a2)))
	if a2 != nil {
		v2 = m.sub_487C80(int32(uintptr(unsafe.Pointer(a2))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*73)) = uint32(v2)
		if v2 != 0 {
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*74)) = *(*uint32)(unsafe.Pointer(uintptr(v2 + 12)))
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(v2 + 16))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*75)) = uint32(v3)
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*76)) = uint32(v3)
			*a2 = 0
		} else {
			v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*72)))))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*74)) = *v4
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1)))
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*75)) = uint32(v5)
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*76)) = uint32(v5)
		}
	}
}

func (m *AudioModule) sub_487C80(a1 int32) int32 {
	return int32(uintptr(unsafe.Pointer(m.nox_common_list_getNext_425940((unsafe.Pointer(uintptr(a1 + 8)))))))
}

func (m *AudioModule) Sub_486F30() int {
	(*m.Externs.Dword_587000_155144).field_0.Clear_425760()
	(*m.Externs.Dword_587000_155144).field_3.Clear_425760()
	(*m.Externs.Dword_587000_155144).field_6 = 0
	(*m.Externs.Ptr_TimerGroup_5d4594_1193340) = &(*m.Externs.Dword_587000_155144).timerGroup_8
	(*m.Externs.Dword_587000_155144).timerGroup_8.Init()
	m.Externs.Dword_5d4594_1193336 = 1
	return 0
}

func (m *AudioModule) Sub_486EF0() {
	if m.Externs.Dword_5d4594_1193336 != 0 {
		if (*m.Externs.Dword_587000_155144).field_6 == 0 {
			v1 := (*m.Externs.Dword_587000_155144).field_3.next
			for it := &((*m.Externs.Dword_587000_155144).field_3); v1 != it; v1 = v1.next {
				v1p := v1.PromoteUnsafe()
				if (v1p.field_3 & 2) == 0 {
					ccall.CallVoidPtr(v1p.field_54, unsafe.Pointer(v1p))
				}
			}
		}
	}
}
