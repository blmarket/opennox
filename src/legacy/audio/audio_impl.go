package audio

import (
	"log"
	"math"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct24[T any] struct {
	ListElement[Struct24[T]]
	field_3 *[0x2000]byte
	field_4 uint32
	field_5 *Struct84Field6
	body_6  T
}

var _ = [1]struct{}{}[2024-unsafe.Sizeof(Struct24[[2000]byte]{})]

type Struct84Field6 struct {
	field_0 *[0x2000]byte
	field_1 uint32
	field_2 ListElement[Struct24[[0x2000]byte]]
	field_5 *Struct84Field14
	field_6 uint32
}

type Struct84Field14 struct {
	field_0 uint32
	field_1 uint32
	field_2 uint32
	field_3 uint32
	field_4 uint32
	field_5 uint32
	field_6 uint32
}

type Struct84 struct {
	ListElement[Struct84]
	field_3  uint32
	field_4  uint32
	field_5  uint32
	field_6  Struct84Field6
	field_13 *Struct28[[0x2000]byte]
	field_14 Struct84Field14
}

var _ = [1]struct{}{}[84-unsafe.Sizeof(Struct84{})]

type Struct28[T any] struct {
	field_0 *AudioStructXxx
	field_1 *FreeList[Struct24[T]]
	field_2 *FreeList[Struct84]
	field_3 ListElement[Struct84]
	field_6 uint32
}

var _ = [1]struct{}{}[28-unsafe.Sizeof(Struct28[byte]{})]
var _ = [1]struct{}{}[unsafe.Sizeof(Struct28[byte]{})-28]

type FreeListItem[T any] struct {
	next *FreeListItem[T]
	item T
}

type FreeList[T any] struct {
	first *FreeListItem[T]
	item0 FreeListItem[T]
	// There are more FreeListItems, just can't use Go slice...
}

func (l *FreeList[T]) Release_4BD300(a2 *T) *FreeListItem[T] {
	result := (*FreeListItem[T])(unsafe.Add(unsafe.Pointer(a2), -4))
	result.next = l.first
	l.first = result
	return result
}

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
	field_22   ListElement[Struct576Field3]
	field_25   uint32
	field_26   uint32
	field_27   int32
	field_28   Struct200Field28
	field_31   uint32
	field_32   [32]uint16
	field_48   int32
	field_49   uint32
}

type Struct200Field28 struct {
	ListElement[Struct200Field28]
}

func (s *Struct200Field28) GetStruct200() *Struct200 {
	return (*Struct200)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - 28*4))
}

var _ = [1]struct{}{}[200-unsafe.Sizeof(Struct200{})]
var _ = [1]struct{}{}[unsafe.Sizeof(Struct200{})-200]

type Struct576Field3 struct {
	ListElement[Struct576Field3]
}

func (s *Struct576Field3) GetStruct576() *Struct576 {
	return (*Struct576)(unsafe.Pointer(uintptr(unsafe.Pointer(s)) - 3*4))
}

type Struct576 struct {
	ListElement[Struct576]
	field_3       Struct576Field3
	field_6       uint8
	field_6_1     uint8
	field_6_2     uint8
	field_6_3     uint8
	field_7       uint32
	field_8       uint32
	field_9       *Struct200
	field_10      [32]*Struct84
	field_42      uint32
	field_43      int32
	field_44      *Struct312
	field_45      uint32
	timerGroup_46 timer.TimerGroup
	field_70      uint32
	field_71      uint32
	field_72      uint64
	field_74      *Struct84Field6
	field_75      uint32
	field_76      [32]int32
	field_108     uint32
	field_109     int32
	field_110     [32]uint32
	field_142     uint32
	field_143     uint32
}

var _ = [1]struct{}{}[576-unsafe.Sizeof(Struct576{})]

func (m *AudioModule) Sub_451850(a2p *Struct264, a3p *AudioStructXxx) int32 {
	var (
		result int32
	)
	v4 := m.Externs.Struct200Arr_5d4594_840628
	for i := int32(0); i < 1023; i++ {
		m.sub_451920(&v4[i])
		v4[i].sndName_21 = m.nox_xxx_getSndName_40AF80(int(i))
	}
	*m.Externs.Dword_5d4594_1045420 = a3p
	*m.Externs.Dword_5d4594_1045428 = a2p
	if a3p != nil {
		*m.Externs.Dword_5d4594_1045424 = sub_4BD340[[0x2000]byte](m, a3p, 0x100000, 200)
		*m.Externs.Dword_5d4594_1045436 = createFreeList_4BD280[Struct576](200)
	}
	if *m.Externs.Dword_5d4594_1045424 == nil || *m.Externs.Dword_5d4594_1045420 == nil || *m.Externs.Dword_5d4594_1045428 == nil || *m.Externs.Dword_5d4594_1045436 == nil {
		return 0
	}
	m.Externs.ListHead_5d4594_840612.Clear_425760()
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
	*m.Externs.Ptr_uint32_5d4594_1045444++
	return int32(*m.Externs.Ptr_uint32_5d4594_1045444)
}

func (m *AudioModule) sub_452190(a1 *Struct200) {
	a1.field_28.Remove_425920()
}

func (m *AudioModule) sub_4521A0(a1 int32) *Struct200 {
	heads := m.Externs.ListHeads_5d4594_839892
	if a1 > 0 {
		for i := int32(0); i < a1; i++ {
			for j := 0; j < 10; j++ {
				v5 := heads[i][j].FirstSafe_4258A0()
				if v5 != nil {
					return v5.PromoteUnsafe().GetStruct200()
				}
			}
		}
	}
	return nil
}

func (m *AudioModule) sub_4521F0() {
	if *m.Externs.Dword_5d4594_1045432 != 0 {
		v1 := m.Externs.ListHead_5d4594_840612.Next()
		if m.Externs.ListHead_5d4594_840612.Next() != m.Externs.ListHead_5d4594_840612 {
			for {
				v2 := v1.next
				m.Sub_4523D0(v1.PromoteUnsafe())
				m.sub_451FE0(v1.PromoteUnsafe())
				v1 = v2
				if v2 == m.Externs.ListHead_5d4594_840612 {
					break
				}
			}
		}
	}
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
		if v1p.field_74 != nil || v1p.field_142 != 0 {
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

func (m *AudioModule) sub_451CA0(a1p *Struct576) *Struct84Field6 {
	var v1 int32
	var v3 int32
	v1 = int32(a1p.field_42)
	a1p.field_108 = uint32(v1)
	if v1 == 0 {
		return nil
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
	a1p.field_10[a1p.field_42] = m.Sub_4BD470(
		*m.Externs.Dword_5d4594_1045424,
		int32(a1p.field_9.field_32[a2]),
	)
	v2 = int32(a1p.field_42)
	result = int32(*(*uint32)(unsafe.Pointer(&a1p.field_10[v2])))
	if result != 0 {
		m.sub_4BD650(*(**Struct84)(unsafe.Pointer(&a1p.field_10[v2])))
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

func (m *AudioModule) sub_451FE0(a1 *Struct576) {
	a1.Remove_425920()
	a1.field_70 = 0
	(*m.Externs.Dword_5d4594_1045436).Release_4BD300(a1)
}

func (m *AudioModule) sub_452120(a1p *Struct576) bool {
	var ret = false
	v3 := m.sub_4521A0(int32(a1p.field_75 + a1p.field_9.field_12))
	if v3 == nil {
		return false
	}
	m.sub_452190(v3)
	v4 := m.Externs.ListHead_5d4594_840612.next.PromoteUnsafe()
	if unsafe.Pointer(*(**uint8)(unsafe.Pointer(m.Externs.ListHead_5d4594_840612))) != unsafe.Pointer(m.Externs.ListHead_5d4594_840612) {
		for {
			v5 := v4.next
			if v4.field_9 == v3 {
				m.Sub_4523D0(v4)
				m.sub_451FE0(v4)
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
		m.sub_452410(a1)
		m.sub_451F90(a1)
		a1.field_7 = 4
		a1.field_70 = 0
		result = int32(*(*uint32)(unsafe.Pointer(&a1.field_6)))
		*((*uint8)(unsafe.Pointer(&result))) = uint8(int8(result | 1))
		*(*uint32)(unsafe.Pointer(&a1.field_6)) = uint32(result)
	}
	return result
}

func (m *AudioModule) sub_452410(a1p *Struct576) {
	var result = a1p.field_44
	if result != nil && a1p == result.field_38 {
		if int32(a1p.field_6)&2 != 0 {
			m.Sub_4BDA80(a1p.field_44)
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
	//var v3 unsafe.Pointer
	var v4 int32
	v1 := a1p.field_44
	if a1p != v1.field_38 {
		return 0
	}
	v3 := a1p.field_74
	m.Sub_4BDB90(v1, a1p.field_74)
	a1p.field_7 = 3
	v4 = int32(a1p.field_6)
	*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4 | 2))
	a1p.field_6 = uint8(int8(v4))
	a1p.field_74 = nil
	if m.Sub_4BDB40(a1p.field_44) == 0 {
		return 1
	}
	a1p.field_7 = 1
	a1p.field_74 = v3
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
			var v3 = uint32(v2 - 2)
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
	var v1p = (*Struct576)(unsafe.Pointer(sub_4BD2E0(*(m.Externs.Dword_5d4594_1045436))))
	if v1p == nil {
		m.sub_452230()
		v1p = (*Struct576)(unsafe.Pointer(sub_4BD2E0(*(m.Externs.Dword_5d4594_1045436))))
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
	v1p.field_70 = *m.Externs.Ptr_uint32_587000_127000
	*m.Externs.Ptr_uint32_587000_127000 += 1
	return v1p
}

func (m *AudioModule) Sub_452E90(a1 *uint32, a2 *Struct576) int32 {
	var result int32
	result = int32(uintptr(unsafe.Pointer(a2)))
	*a1 = uint32(uintptr(unsafe.Pointer(a2)))
	if a2 != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1)) = a2.field_70
		result = int32(*(*uint32)(unsafe.Pointer(&a2.field_9)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)) = uint32(result)
	}
	return result
}

func (m *AudioModule) Sub_452EE0(a1p *Struct576, a2 int32) int32 {
	a1p.timerGroup_46.Timers[0].SetRaw(uint32(int32(m.sub_452F10(a1p, a2))))
	return bool2int32(a1p.timerGroup_46.Timers[0].Update())
}

func (m *AudioModule) sub_452F10(a1 *Struct576, a2 int32) uint32 {
	var v2 int32
	v2 = a2
	if a2 <= 100 {
		if a2 < 0 {
			v2 = 0
		}
	} else {
		v2 = 100
	}
	v1 := a1.field_9
	return (uint32(v2*163) * (v1.field_4.Current >> 16)) >> 14
}

func (m *AudioModule) Sub_452F50(a1p *Struct576, a2 int32) int32 {
	a1p.timerGroup_46.Timers[0].SetInterp(uint32(int32(m.sub_452F10(a1p, a2))))
	return 0
}

func (m *AudioModule) Sub_452F80(a1p *Struct576, a2 int32) {
	var v2 = m.sub_452FA0(a2)
	a1p.timerGroup_46.Timers[2].SetRaw(uint32(v2))
}

func (m *AudioModule) sub_451CF0(a1 *Struct576) *Struct84Field6 {
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
				v8 = (*uint32)(unsafe.Pointer(&a1.field_76[v5]))
				for {
					v7++
					*v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
					v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v8), 4*1))
					if uint32(v7) >= a1.field_108-1 {
						break
					}
				}
			}
		} else {
			*(*uint32)(unsafe.Pointer(&a1.field_43))++
		}
		v9 = int32(*(*uint32)(unsafe.Pointer(&a1.field_43)))
		a1.field_108--
		return m.sub_4BD710(a1.field_10[v9])
	}

	if v3&1 != 0 {
		if v1.field_15 != 0 && (func() bool {
			v4 := a1.field_109 + 1
			a1.field_109 = v4
			return v4 >= v1.field_15
		}()) {
			return nil
		} else {
			return m.sub_451CA0(a1)
		}
	}

	return nil
}

func (m *AudioModule) sub_451DC0(a1p *Struct576) int32 {
	var (
		v1     *Struct200
		result int32
		v3     int32
		i      int32
		v5     int32
		v6     int32
	)
	v1 = a1p.field_9
	result = int32(a1p.field_42)
	v3 = int32(*(*uint32)(unsafe.Pointer(&v1.field_1)))
	if result != 0 {
		if v1.field_17 < 0x21 {
			return result
		}
		m.sub_451F90(a1p)
	}
	if v3&4 != 0 {
		if v1.field_17 >= 0x21 {
			v5 = m.sub_451E80(a1p)
			result = m.sub_451F30(a1p, v5)
		} else {
			result = int32(v1.field_48)
			for i = 0; i < result; i++ {
				m.sub_451F30(a1p, i)
				result = int32(v1.field_48)
			}
		}
	} else if v3&2 != 0 {
		v6 = int32(m.nox_common_randomIntMinMax_415FF0(0, int(v1.field_48-1), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 536))
		result = m.sub_451F30(a1p, v6)
	} else {
		result = m.sub_451F30(a1p, 0)
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
	if v1p.field_48 == 0 {
		return 0
	}
	v3 = int32(a1.field_75)
	a1.field_109 = 0
	res = (*Struct312)(m.Sub_452810(int32(v1p.field_12+uint32(v3)), 0))
	a1.field_44 = res
	if res != nil {
		v4 = int32(m.nox_common_randomIntMinMax_415FF0(int(v1p.field_19), int(v1p.field_20), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1482))
		a1.field_44.timerGroup_4.Timers[1].SetRaw(uint32(v4 + 100))
		m.sub_4BDB20(a1.field_44)
		a1.field_44.field_38 = a1
		a1.field_44.field_35 = m.Externs.Sub_452770_ptr
		a1.field_44.field_36 = m.Externs.Sub_4526F0_ptr
		a1.field_44.field_37 = m.Externs.Sub_4526D0_ptr
		a1.field_7 = 1
		a1.field_44.field_28 = &a1.timerGroup_46
		if int32(v1p.field_1)&8 != 0 {
			v5 = int32(m.nox_common_randomIntMinMax_415FF0(int(v1p.field_17), int(v1p.field_18), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1497))
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
		v5 uint32
	)
	v1 = a1.field_38
	v2 := m.sub_451CF0(a1.field_38)
	if v1.field_9.field_18 < 0x21 {
		m.Sub_4BDB90(a1, v2)
		return 0
	}
	m.Sub_4BDB90(a1, nil)
	v4 := v1.field_9
	if (v4.field_1&8) == 0 || v2 != nil || v1.field_142 != 0 {
		v5 = uint32(m.nox_common_randomIntMinMax_415FF0(int(v4.field_17), int(v4.field_18), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 706))
		if v5 < 0x21 {
			m.Sub_4BDB90(a1, v2)
			return 0
		}
		v1.field_71 = v5
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
				v7.Remove_425920()
				heads[v2][v3].Append_4258E0(&v7.ListElement)
			}
		} else {
			v1.field_27 = v2
			v1.field_31 = v3
			v6 := &v1.field_28
			v6.Remove_425920()
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

//goland:noinspection ALL
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
	for v4 := v2p.field_22.next; v4 != &v2p.field_22; v4 = v4.next {
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
			v6 = int32(v4p.field_7)
			if int32(v2p.field_1)&0x10 != 0 {
				if v6 != 0 {
					break
				}
			} else if v6 == 0 {
				break
			}
		}
	}

	v7 := &v1.field_3.ListElement
	v1.field_3.Init_425770()
	v2p.field_22.Append_4258E0(v7)
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
	if *m.Externs.Dword_5d4594_1045424 != nil {
		Sub_4BD3C0(m, *m.Externs.Dword_5d4594_1045424)
		*m.Externs.Dword_5d4594_1045424 = nil
	}
	if *m.Externs.Dword_5d4594_1045436 != nil {
		m.Sub_4BD2D0(*(*unsafe.Pointer)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)))
		*m.Externs.Dword_5d4594_1045436 = nil
	}
	*m.Externs.Dword_5d4594_1045432 = 0
}

func (m *AudioModule) Sub_4519C0() {
	var (
		v3 int32
	)
	if *m.Externs.Dword_5d4594_1045432 == 0 {
		return
	}
	if *m.Externs.Ptr_uint32_5d4594_1045448 != 0 {
		return
	}
	*m.Externs.Ptr_uint32_5d4594_1045448 = 1
	(*m.Externs.Ptr_TimerGroup_587000_127004).Update()
	v1p := m.Externs.ListHead_5d4594_840612.next
	*m.Externs.Ptr_uint32_5d4594_1045440++
	if m.Externs.ListHead_5d4594_840612.next != m.Externs.ListHead_5d4594_840612 {
		for {
			v1pp := v1p.PromoteUnsafe()
			v2p := v1pp.field_9
			if v2p.field_25 != *m.Externs.Ptr_uint32_5d4594_1045440 {
				v2p.field_22.Clear_425760()
				v2p.field_13 = 0
				v2p.field_25 = *m.Externs.Ptr_uint32_5d4594_1045440
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
			if int32(v6.field_7) == 1 {
				m.sub_451DC0(v6)
				v8 := m.sub_451CA0(v6)
				v6.field_74 = v8
				if v8 == nil {
					for m.sub_452120(v6) {
						v7 = v6.next
						m.sub_451DC0(v6)
						v9 := m.sub_451CA0(v6)
						v6.field_74 = v9
						if v9 != nil {
							break
						}
					}
				}
				v10 := m.sub_451CA0(v6)
				v6.field_74 = v10
				if v10 == nil || m.sub_452490(v6) == 0 {
					m.Sub_4523D0(v6)
					m.sub_451FE0(v6)
				}
			}
			v6 = v7.PromoteUnsafe()
			if v7 == m.Externs.ListHead_5d4594_840612 {
				break
			}
		}
	}
	*m.Externs.Ptr_uint32_5d4594_1045448 = 0
}

func (m *AudioModule) sub_4BDB20(a1p *Struct312) {
	a1p.field_30.field_1 |= 0x10
}

func (m *AudioModule) sub_4BD710(a1 *Struct84) *Struct84Field6 {
	return &a1.field_6
}

func (m *AudioModule) Sub_4526D0(a1p *Struct312) int32 {
	a1p.field_38.field_7 = 4
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

func (m *AudioModule) sub_4BD650(a1p *Struct84) {
	*(*uint32)(unsafe.Pointer(&a1p.field_3))++
}

func (m *AudioModule) sub_4BD660(a1p *Struct84) int32 {
	var result = int32(a1p.field_3) - 1
	a1p.field_3 = uint32(result)
	if result < 0 {
		a1p.field_3 = 0
	}
	return result
}

func (m *AudioModule) Nox_xxx_clientPlaySoundSpecial_452D80(a1 int32, a2 int32) {
	var result = (*uint32)(unsafe.Pointer(m.Nox_xxx_draw_452270(a1)))
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
	if res2 == nil {
		return
	}
	m.Sub_452EE0(res2, a2)
	m.Sub_452F80(res2, a3)
	res2.field_75 = 2
	m.sub_452510(res2)
}

func (m *AudioModule) Sub_4BDB30(a1 *Struct312) {
	a1.field_30.field_1 &= 0xFFFFFFEF
}

func (m *AudioModule) Sub_4BD720(a1p *Struct264) *Struct312 {
	v1pp, _ := alloc.Calloc(1, 0x138)
	v1p := (*Struct312)(v1pp)
	alloc.Memset(unsafe.Pointer(v1p), 0, 0x138)
	v1p.Init_425770()
	m.sub_4BDC00(&v1p.field_30)
	v1p.timerGroup_44.Init()
	m.sub_4BD7C0(v1p)
	v1p.field_33 = a1p
	v1p.field_43 = a1p.field_64
	if ccall.CallIntPtr(a1p.field_64.field_1, unsafe.Pointer(v1p)) == 0 {
		return v1p
	}
	m.Sub_4BD7A0(v1p)
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
	m.sub_4BDC00(&a1p.field_30)
	a1p.field_30.field_0 = 0
	a1p.field_29 = *m.Externs.Ptr_TimerGroup_5d4594_1193340
	a1p.field_28 = nil
	a1p.timerGroup_4.Init()
	a1p.field_72 = nil
}

func (m *AudioModule) Sub_4BD8C0(a1p *Struct312) int32 {
	var (
		result int32
		v4     int32
	)
	v1 := *(*unsafe.Pointer)(unsafe.Pointer(&a1p.field_34))
	if v1 != nil {
		result = int32(ccall.CallIntPtr(v1, unsafe.Pointer(a1p)))
		if result != 0 {
			a1p.field_75 = 0
			a1p.field_76 = 0
			a1p.field_74 = nil
			return result
		}
	} else {
		if a1p.field_73 != nil {
			v3 := a1p.field_73.NextSafe_425940()
			a1p.field_73 = v3
			if v3 != nil {
				a1p.field_74 = v3.field_3
				v4 = int32(*(*uint32)(unsafe.Pointer(unsafe.Pointer(&v3.field_4))))
				a1p.field_75 = uint32(v4)
				a1p.field_76 = uint32(v4)
				return 0
			}
		}
		a1p.field_75 = 0
	}
	return 0
}

func (m *AudioModule) Sub_4BD940(a1p *Struct312) int32 {
	if a1p.field_30.field_2 != 0 {
		if a1p.field_30.field_2 != -1 {
			a1p.field_30.field_2--
		}
		m.Sub_4BDB90(a1p, a1p.field_72)
	} else {
		m.Sub_4BDB90(a1p, nil)
	}
	v1 := a1p.field_35
	if v1 != nil {
		ccall.CallVoidPtr(v1, unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))))))
	}
	if *(*uint32)(unsafe.Pointer(&a1p.field_72)) != 0 {
		ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1p.field_43)) + 36))), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(a1p))))))
	}
	return 0
}

func (m *AudioModule) Sub_4BD9B0(a2p *Struct312) int32 {
	var (
		result int32
	)
	a2p.field_72 = nil
	a2p.field_30.field_1 &= 0xFA
	a2p.field_30.field_2 = 0
	a2p.timerGroup_4.Init()
	v2 := a2p.field_36
	if v2 != nil {
		result = int32(ccall.CallIntPtr(v2, unsafe.Pointer(a2p)))
	} else {
		result = 0
	}
	return result
}

func (m *AudioModule) sub_4BDC00(a1p *Struct312Field30) {
	a1p.field_2 = 0
	a1p.field_1 = 0
}

func (m *AudioModule) Sub_4BDB90(a1p *Struct312, a2p *Struct84Field6) {
	var (
		v3 int32
		v5 int32
	)
	a1p.field_72 = a2p
	if a2p != nil {
		v2 := m.sub_487C80(a2p)
		a1p.field_73 = v2
		if v2 != nil {
			a1p.field_74 = v2.field_3
			v3 = int32(v2.field_4)
			a1p.field_75 = uint32(v3)
			a1p.field_76 = uint32(v3)
			a2p.field_0 = nil
		} else {
			v4 := a1p.field_72
			a1p.field_74 = v4.field_0
			v5 = int32(v4.field_1)
			a1p.field_75 = uint32(v5)
			a1p.field_76 = uint32(v5)
		}
	}
}

func (m *AudioModule) sub_487C80(a1 *Struct84Field6) *Struct24[[0x2000]byte] {
	return a1.field_2.NextSafe_425940()
}

func (m *AudioModule) Sub_486F30() int {
	(*m.Externs.Dword_587000_155144).field_0.Clear_425760()
	(*m.Externs.Dword_587000_155144).field_3.Clear_425760()
	(*m.Externs.Dword_587000_155144).field_6 = 0
	*m.Externs.Ptr_TimerGroup_5d4594_1193340 = &(*m.Externs.Dword_587000_155144).timerGroup_8
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

func (m *AudioModule) Sub_4BDA80(a1 *Struct312) {
	if int32(a1.field_30.field_1)&5 != 0 {
		ccall.CallVoidPtr(a1.field_43.field_4, unsafe.Pointer(a1))
	}
	result2 := a1.field_37
	if result2 != nil {
		ccall.CallIntPtr(result2, unsafe.Pointer(a1))
	}
	*(*uint32)(unsafe.Pointer(&a1.field_72)) = 0
}

func (m *AudioModule) sub_486E90(a1 *Struct312) int32 {
	var (
		result int32
	)
	v1p := a1.field_33
	a1.Remove_425920()
	v1p.field_48--
	v1p.field_53++
	a1.Remove_425920()
	result = v1p.field_53 - 1
	v1p.field_53 = result
	if result < 0 {
		v1p.field_53 = 0
	}
	return result
}

func (m *AudioModule) Sub_4BDA60(a1 *Struct312) {
	m.Sub_4BDA80(a1)
	m.sub_486E90(a1)
	m.Sub_4BD7A0(a1)
}

func (m *AudioModule) Sub_4873C0(a3p *Struct264) int32 {
	var (
		v3  int64
		v11 *timer.TimerGroup
		v17 bool
	)
	if a3p.field_53 != 0 {
		return -2146304000
	}
	v3 = int64(m.nox_platform_get_ticks())

	diff1 := v3 - a3p.field_62
	if diff1 >= a3p.field_56 {
		a3p.field_58 = diff1
		if a3p.field_60 > a3p.field_56*10 {
			a3p.field_60 = 0
		}
		if diff1 > a3p.field_60 {
			a3p.field_60 = diff1
		}

		v11 = &a3p.TimerGroup_22
		a3p.TimerGroup_22.Update()

		if a3p.field_46 != nil {
			a3p.field_46.Update()
			v17 = a3p.field_46.IsUpdated()
		}
		if a3p.field_46 == nil || !v17 {
			v17 = a3p.TimerGroup_22.IsUpdated()
		}
		a3p.field_62 = v3

		for v12 := a3p.field_50.next; v12 != &a3p.field_50; v12 = v12.next {
			v12p := v12.PromoteUnsafe()
			if (v12p.field_30.field_1&1) != 0 && v12p.field_72 != nil {
				v12p.timerGroup_4.Update()
				if v17 || v12p.timerGroup_4.IsUpdated() || (v12p.field_29 != nil && v12p.field_29.IsUpdated()) || (v12p.field_28 != nil && v12p.field_28.IsUpdated()) {
					m.sub_4BD840(v12p)
					ccall.CallVoidPtr(v12p.field_43.field_8, unsafe.Pointer(v12p))
				}
			}
		}
		v14 := a3p.field_46
		if v14 != nil {
			v14.ClearUpdated()
		}
		v11.ClearUpdated()
	}
	return 0
}

func (m *AudioModule) sub_4BD840(a3p *Struct312) {
	var v1p = a3p.field_33

	a3p.timerGroup_44.Init()
	a3p.timerGroup_44.Mix(&a3p.timerGroup_4)
	a3p.timerGroup_4.ClearUpdated()
	if a3p.field_28 != nil {
		a3p.timerGroup_44.Mix(a3p.field_28)
		a3p.field_28.ClearUpdated()
	}
	if a3p.field_29 != nil {
		a3p.timerGroup_44.Mix(a3p.field_29)
	}
	a3p.timerGroup_44.Mix(&v1p.TimerGroup_22)
	if v1p.field_46 != nil {
		a3p.timerGroup_44.Mix(v1p.field_46)
	}
}

func (m *AudioModule) Sub_4BDB40(a2p *Struct312) int32 {
	var result int32
	if int32(*(*uint8)(unsafe.Pointer(&a2p.field_30.field_1)))&5 != 0 {
		return -2146500608
	}
	if a2p.field_72 == nil {
		return -2147024896
	}
	a2p.timerGroup_4.Update()
	m.sub_4BD840(a2p)
	result = int32(ccall.CallIntPtr(a2p.field_43.field_3, unsafe.Pointer(a2p)))
	if result == 0 {
		a2p.field_30.field_1 |= 1
	}
	return result
}

func (m *AudioModule) sub_4871C0(a1p *Struct88, a2 int32, a3 *[7]uint32) *Struct264 {
	var v3 = a1p.field_3
	v4p, _ := alloc.Calloc(1, 0x108)
	v4m := (*Struct264)(v4p)
	alloc.Memset(unsafe.Pointer(v4m), 0, 0x108)
	v4m.field_0.Init_425770()
	v4m.field_6 = uint32(a2)
	v4m.field_5 = a1p
	v4m.field_4 = 0
	a1p.field_4++
	a1p.field_6[a2] = v4m
	v4m.field_64 = v3.field_9
	v4m.field_50.Clear_425760()
	v4m.TimerGroup_22.Init()
	v4m.field_53 = 0
	v4m.field_56 = 33
	v4m.field_60 = 0
	v4m.field_58 = 0
	v4m.field_62 = 0
	v4m.field_54 = m.Externs.Sub_4873C0_ptr
	if a3 != nil {
		m.sub_487590(v4m, a3)
	}
	if ccall.CallIntPtr(v3.field_7, unsafe.Pointer(v4m)) == 0 {
		return v4m
	}
	if v4m != nil {
		m.sub_4872C0(v4m)
	}
	return nil
}

func (m *AudioModule) Sub_487150(aa1 int32, a2 *[7]uint32) *Struct264 {
	var (
		v2 int32
		v6 int32
	)
	v2 = aa1
	if aa1 == -1 {
		v2 = 0
	}
	var a1 *Struct88
	m.sub_487360(v2, &a1, &v6)
	if a1 == nil {
		return nil
	}
	v3m := a1.field_6[v6]
	if v3m == nil {
		v4m := m.sub_4871C0(a1, v6, a2)
		v3m = v4m
		if v4m == nil {
			return nil
		}
		v4m.field_47 = uint32(v2)
		m.sub_487310(v4m)
	}
	v3m.field_4 += 1
	return v3m
}

func (m *AudioModule) sub_487310(a1_ *Struct264) int32 {
	var (
		result int32
	)
	(*m.Externs.Dword_587000_155144).field_6++
	(*m.Externs.Dword_587000_155144).field_3.Append_4258E0(&a1_.field_0)
	result = int32((*m.Externs.Dword_587000_155144).field_6 - 1)
	(*m.Externs.Dword_587000_155144).field_6 = uint32(result)
	if result < 0 {
		(*m.Externs.Dword_587000_155144).field_6 = 0
	}
	return result
}

func (m *AudioModule) Sub_431270() {
	if *m.Externs.Dword_5d4594_805984 != nil {
		m.Sub_487680(*m.Externs.Dword_5d4594_805984)
		*m.Externs.Dword_5d4594_805984 = nil
	}
}

func (m *AudioModule) Sub_487680(lpMem_ *Struct264) {
	var lpMem = unsafe.Pointer(lpMem_)
	m.sub_4876A0((*Struct264)(unsafe.Pointer((**uint32)(lpMem))))
	m.sub_4872C0((*Struct264)(lpMem))
}

func (m *AudioModule) Sub_431290() {
	if *m.Externs.Dword_5d4594_805984 != nil {
		m.sub_487970(*m.Externs.Dword_5d4594_805984, -1)
	}
}

func (m *AudioModule) sub_487970(a1p *Struct264, a2 int32) {
	var (
		a1x *Struct312
		v3  *Struct312
		v4  int32
	)
	v3 = m.sub_4877D0(a1p, &a1x)
	if v3 != nil {
		v4 = a2
		for {
			v5 := m.sub_4877F0(&a1x)
			if v4 == -1 || v3.field_3 == v4 {
				m.Sub_4BDA80(v3)
			}
			v3 = v5
			if v5 == nil {
				break
			}
		}
	}
}

func (m *AudioModule) sub_487590(a1_ *Struct264, a2 *[7]uint32) int32 {
	var (
		a1     = int32(uintptr(unsafe.Pointer(a1_)))
		result int32
	)
	result = a1
	alloc.Memcpy(unsafe.Pointer(&a1_.field_15), unsafe.Pointer(a2), 0x1C)
	return result
}

func (m *AudioModule) sub_4872C0(a1p *Struct264) {
	var (
		v1 *Struct88
		v2 int32
	)
	m.sub_487910(a1p, -1)
	ccall.CallVoidPtr(a1p.field_5.field_3.field_8, unsafe.Pointer(a1p))
	a1p.field_5.field_6[a1p.field_6] = nil
	v1 = a1p.field_5
	v2 = int32(v1.field_4) - 1
	v1.field_4 = uint32(v2)
	if v2 < 0 {
		v1.field_4 = 0
	}
	alloc.Free(a1p)
}

func (m *AudioModule) sub_4876A0(a1 *Struct264) {
	var (
		result int32
	)
	(*m.Externs.Dword_587000_155144).field_6++
	a1.field_0.Remove_425920()
	result = int32((*m.Externs.Dword_587000_155144).field_6 - 1)
	(*m.Externs.Dword_587000_155144).field_6 = uint32(result)
	if result < 0 {
		(*m.Externs.Dword_587000_155144).field_6 = 0
	}
}

func (m *AudioModule) sub_487910(a1p *Struct264, a2 int32) int32 {
	var (
		// v2 *int32
		a1x *Struct312
		v3  int32
		// v4 *int32
	)
	v2 := m.sub_4877D0(a1p, &a1x)
	if v2 == nil {
		return 0
	}
	v3 = a2
	for {
		v4 := m.sub_4877F0(&a1x)
		if v3 == -1 || v2.field_3 == v3 {
			m.Sub_4BDA60(v2)
		}
		v2 = v4
		if v4 == nil {
			break
		}
	}
	return 0
}

func (m *AudioModule) Sub_452810(a1 int32, a2 int8) *Struct312 {
	var (
		v2 *Struct312
		v3 *Struct312
	)
	v2 = nil
	if *m.Externs.Dword_5d4594_1045428 != nil {
		v3 = m.sub_487810(*m.Externs.Dword_5d4594_1045428, 1)
		v2 = v3
		if v3 != nil {
			if v3.field_30.field_1&0x15 != 0 && v3.field_30.field_0 > a1 {
				return nil
			}
			m.Sub_4BDA80(v3)
			v2.field_29 = *m.Externs.Dword_587000_127004
			v2.field_30.field_0 = a1
			if int32(a2)&1 != 0 {
				v2.field_30.field_2 = -1
			} else {
				v2.field_30.field_2 = 0
			}
			v2.timerGroup_4.Timers[0].SetRaw(0x4000)
		}
	}
	return v2
}

func (m *AudioModule) sub_487810(a1p *Struct264, a2 int32) *Struct312 {
	var (
		v2     uint32
		v3     int32
		v4     *Struct312
		result *Struct312
		v6     int32
		v7     uint32
		v8     int32
		v9     *Struct312
		a1x    *Struct312
	)
	v2 = 4294967295
	if a2 == -1 {
		a2 = 1
	}
	v3 = 127
	v4 = nil
	v8 = 127
	v9 = nil
	for result = m.sub_4877D0(a1p, &a1x); result != nil; result = m.sub_4877F0(&a1x) {
		if result.field_3 == a2 {
			if (*(*int32)(unsafe.Pointer(&result.field_30.field_1)) & 0x15) == 0 {
				return result
			}
			v6 = result.field_30.field_0
			if result.field_30.field_1&1 != 0 {
				if v6 >= v3 {
					if v6 == v3 {
						v7 = uint32(*(*int32)(unsafe.Pointer(&result.timerGroup_44.Timers[0].Current)))
						if v7 < v2 && v2-v7 >= 0x666 {
							v3 = result.field_30.field_0
							v4 = result
							v2 = uint32(*(*int32)(unsafe.Pointer(&result.timerGroup_44.Timers[0].Current)))
						}
					}
				} else {
					v2 = uint32(*(*int32)(unsafe.Pointer(&result.timerGroup_44.Timers[0].Current)))
					v3 = result.field_30.field_0
					v4 = result
				}
			} else if v6 < v8 {
				v8 = result.field_30.field_0
				v9 = result
			}
		}
	}
	result = v9
	if v9 == nil || v8 > v3 {
		result = v4
	}
	return result
}

func (m *AudioModule) sub_4877D0(a1p *Struct264, a2 **Struct312) *Struct312 {
	result := a1p.field_50.FirstSafe_425890()
	*a2 = result
	return result
}

func (m *AudioModule) sub_4877F0(a1 **Struct312) *Struct312 {
	if *a1 != nil {
		*a1 = (*a1).NextSafe_4258A0()
	}
	return *a1
}

func (m *AudioModule) sub_487360(a1 int32, a2 **Struct88, a3 *int32) {
	var (
		result *Struct88
		i      int32
		v5     int32
		v6     *Struct88
	)
	result = m.Sub_4870E0(&v6)
	for i = a1; result != nil; result = m.Sub_487100(&v6) {
		v5 = result.field_5
		if i < v5 {
			break
		}
		i -= v5
	}
	*a2 = result
	if result != nil {
		*a3 = i
	} else {
		*a3 = -1
	}
}

func (m *AudioModule) sub_487750(a1p *Struct264) *Struct312 {
	if a1p.field_48 >= a1p.field_49 {
		return nil
	}
	v1p := m.Sub_4BD720(a1p)
	v2p := v1p
	if v1p == nil {
		return nil
	}
	m.sub_486E30(a1p, v1p)
	return v2p
}

func (m *AudioModule) Sub_487790(a1p *Struct264, a2 int32) int32 {
	var (
		v2 int32
		v3 int32
	)
	v2 = 0
	if m.sub_487750(a1p) != nil {
		v3 = a2
		for {
			v2++
			v3--
			if v3 == 0 || m.sub_487750(a1p) == nil {
				break
			}
		}
	}
	return v2
}

func (m *AudioModule) sub_486E30(a1p *Struct264, a2p *Struct312) int32 {
	a2p.field_33 = a1p
	a1p.field_48++
	a1p.field_53++
	a1p.field_50.Append_4258E0(&a2p.ListElement)
	result := a1p.field_53 - 1
	a1p.field_53 = result
	if result < 0 {
		a1p.field_53 = 0
	}
	return result
}

func (m *AudioModule) Sub_486FE0(a1 *Struct587000_94032) *Struct88 {
	v1p, _ := alloc.Calloc(1, 0x58)
	v1pp := (*Struct88)(v1p)
	alloc.Memset(v1p, 0, 0x58)
	v1pp.field_0.Init_425770()
	v1pp.field_4 = 0
	v1pp.field_3 = a1

	if ccall.CallIntPtr(a1.field_5, unsafe.Pointer(v1pp)) == 0 {
		return v1pp
	}
	if v1pp != nil {
		m.sub_487030(v1pp)
	}
	return nil
}

func (m *AudioModule) sub_487030(a1p *Struct88) {
	// var lpMem unsafe.Pointer = unsafe.Pointer(a1p)
	ccall.CallVoidPtr(a1p.field_3.field_6, unsafe.Pointer(a1p))
	a1p.field_3.field_3 &= 0xFFFFFFFE
	alloc.Free(a1p)
}

func (m *AudioModule) Sub_487050(a1p *Struct88) {
	(*m.Externs.Dword_587000_155144).field_0.Append_4258E0(&a1p.field_0)
}

func (m *AudioModule) Sub_4870A0() {
	var (
		v1 *Struct88
		v2 *Struct88
		v3 *Struct88
	)
	v1 = m.Sub_4870E0(&v3)
	if v1 != nil {
		for {
			v2 = m.Sub_487100(&v3)
			m.sub_487070(v1)
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}

func (m *AudioModule) Sub_4870E0(a1 **Struct88) *Struct88 {
	result := (*m.Externs.Dword_587000_155144).field_0.FirstSafe_425890()
	*a1 = result
	return result
}

func (m *AudioModule) sub_487070(a1 *Struct88) {
	m.sub_487090(a1)
	m.sub_487030(a1)
	*m.Externs.Ptr_uint32_5d4594_1193332 = 0
}

func (m *AudioModule) sub_487090(a1 *Struct88) {
	a1.field_0.Remove_425920()
}

func (m *AudioModule) Sub_487100(a1 **Struct88) *Struct88 {
	if *a1 != nil {
		res := (*a1).field_0.NextSafe_4258A0()
		*a1 = res
	}
	return *a1
}

func (m *AudioModule) sub_4875B0(a1 **Struct264) *Struct264 {
	result := (*m.Externs.Dword_587000_155144).field_3.FirstSafe_425890()
	*a1 = result
	return result
}

func (m *AudioModule) sub_4875D0(a1 **Struct264) *Struct264 {
	if *a1 != nil {
		*a1 = (*a1).field_0.NextSafe_4258A0()
	}
	return *a1
}

func (m *AudioModule) Sub_4875F0() int32 {
	var (
		v0     *Struct264
		v1     *Struct264
		result int32
		v3     *Struct264
	)
	(*m.Externs.Dword_587000_155144).field_6 += 1
	v0 = m.sub_4875B0(&v3)
	if v0 != nil {
		for {
			v1 = m.sub_4875D0(&v3)
			m.Sub_487680(v0)
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
	result = int32((*m.Externs.Dword_587000_155144).field_6 - 1)
	(*m.Externs.Dword_587000_155144).field_6 = uint32(result)
	if result < 0 {
		(*m.Externs.Dword_587000_155144).field_6 = 0
	}
	return result
}

func (m *AudioModule) Sub_487D60(a1p *Struct24[[0x2000]byte]) {
	a1p.field_5 = nil
}

func (m *AudioModule) Sub_4BD680(a1p *Struct84) int32 {
	return int32(a1p.field_3)
}

func (m *AudioModule) Sub_487C50(a1p *Struct84Field6, a2 *Struct24[[0x2000]byte]) int32 {
	var result int32
	(&a1p.field_2).Append_4258E0(&a2.ListElement)
	result = int32(a2.field_4 + a1p.field_1)
	a1p.field_1 = uint32(result)
	a2.field_5 = a1p
	return result
}

func (m *AudioModule) Sub_4BD690(a1p *Struct84) {
	if *(*uint32)(unsafe.Pointer(&a1p.prev)) != uint32(int32(uintptr(unsafe.Pointer(a1p)))) {
		a1p.Remove_425920()
	}
	for it := a1p.field_6.field_2.NextSafe_425940(); it != nil; it = a1p.field_6.field_2.NextSafe_425940() {
		it.Remove_425920()
		m.Sub_487D60(it)
		a1p.field_13.field_1.Release_4BD300(it)
	}
	a1p.field_13.field_2.Release_4BD300(a1p)
}

func (m *AudioModule) Sub_4BD600(a1p *Struct28[[0x2000]byte]) int32 {
	v1 := a1p.field_3.PrevSafe_425960()
	if v1 == nil {
		return 0
	}
	for m.Sub_4BD680(v1) != 0 {
		v1 = v1.PrevSafe_425960()
		if v1 == nil {
			return 0
		}
	}
	m.Sub_4BD690(v1)
	return 1
}

func Sub_425900[T any](a1p *ListElement[T], a2p *ListElement[T]) {
	a2p.prev = a1p
	a2p.next = a1p.next
	a1p.next = a2p
	a2p.next.prev = a2p
}

func (m *AudioModule) Sub_487D30(a1p *Struct24[[0x2000]byte], a2 *[0x2000]byte, a3 int32) *uint32 {
	var result *uint32
	*(*uint32)(unsafe.Pointer(&a1p.field_4)) = uint32(a3)
	a1p.field_3 = a2
	((*UnknownListElement)(unsafe.Pointer(a1p))).Init_425770()
	*(*uint32)(unsafe.Pointer(&a1p.field_5)) = 0
	return result
}

func createFreeList_4BD280[T any](a1 int32) *FreeList[T] {
	var (
		v2     int32
		result *FreeList[T]
	)
	var v T
	var a2 = int32(unsafe.Sizeof(v))

	v2 = a2 + 4
	res, _ := alloc.Calloc(1, uintptr(a1*(a2+4)+4))
	result = (*FreeList[T])(res)
	if result != nil {
		v4 := &result.item0
		result.first = &result.item0
		for i := int32(0); i+1 < a1; i++ {
			v5 := (*FreeListItem[T])(unsafe.Add(unsafe.Pointer(v4), v2)) //nolint: unsafeadd
			v4.next = v5
			v4 = v5
		}
		v4.next = nil
	}
	return result
}

func sub_4BD2E0[T any](a1 *FreeList[T]) *T {
	result := a1.first
	if result == nil {
		return nil
	}
	v2 := result.next
	ret := &result.item
	a1.first = v2
	return ret
}

func (m *AudioModule) Sub_4BD2D0(lpMem unsafe.Pointer) {
	alloc.FreePtr(lpMem)
}

func (m *AudioModule) Sub_4866D0(a1p *AudioStructXxx, a2 int32) *AudioStructYyy {
	return (*AudioStructYyy)(unsafe.Add(unsafe.Pointer(a1p.Arr0), a2*36)) //nolint: unsafeadd
}

func (m *AudioModule) Sub_487C30(a1p *Struct84Field6) {
	a1p.field_0 = nil
	a1p.field_1 = 0
	a1p.field_5 = nil
	a1p.field_6 = 0
	a1p.field_2.Clear_425760()
}

func (m *AudioModule) Sub_487D00(a1p *[7]uint32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(a1p[1])
	result = int32(a1p[2] * a1p[3] * a1p[4])
	a1p[5] = uint32(result)
	if v1 == 1 {
		result >>= 2
		a1p[5] = uint32(result)
	}
	return result
}

func sub_4BD340[T any](m *AudioModule, a1p *AudioStructXxx, a2 int32, a3 int32) *Struct28[T] {
	var itemT T
	a4 := int32(unsafe.Sizeof(itemT))
	var v4 *Struct28[T]
	res, _ := alloc.Calloc(1, 0x1C)
	v4 = (*Struct28[T])(res)
	alloc.Memset(unsafe.Pointer(v4), 0, 0x1C)
	v4.field_0 = a1p
	v4.field_6 = uint32(a4)
	v4.field_1 = createFreeList_4BD280[Struct24[T]](a2 / (a4 + 24))
	v4.field_2 = createFreeList_4BD280[Struct84](a3)
	v4.field_3.Clear_425760()
	if v4.field_1 != nil && v4.field_2 != nil {
		return v4
	}
	Sub_4BD3C0(m, v4)
	return nil
}

func Sub_4BD3C0[T any](m *AudioModule, a1 *Struct28[T]) {
	for i := a1.field_3.NextSafe_425940(); i != nil; i = a1.field_3.NextSafe_425940() {
		m.Sub_4BD690(i)
	}
	if a1.field_1 != nil {
		m.Sub_4BD2D0(unsafe.Pointer(a1.field_1))
	}
	if a1.field_2 != nil {
		m.Sub_4BD2D0(unsafe.Pointer(a1.field_2))
	}
	alloc.Free(a1)
}

func (m *AudioModule) Sub_486AA0(a1p *AudioStructXxx, a2 int32, a3p *Struct84Field14) int32 {
	var (
		result int32
	)
	var v3p = m.Sub_4866D0(a1p, a2)

	a3p.field_0 = 4
	a3p.field_2 = *(*uint32)(unsafe.Pointer(&v3p.Field24))
	a3p.field_3 = uint32(bool2int32((*(*uint32)(unsafe.Pointer(&v3p.Field28))&1) != 0) + 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3p), 4*6)) = *(*uint32)(unsafe.Pointer(&v3p.Field32))
	if *(*uint32)(unsafe.Pointer(&v3p.Field28))&8 != 0 {
		result = 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3p), 4*1)) = 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3p), 4*4)) = 2
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3p), 4*1)) = 0
		result = bool2int32((*(*uint32)(unsafe.Pointer(&v3p.Field28))&4) != 0) + 1
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3p), 4*4)) = uint32(result)
	}
	return result
}

func (m *AudioModule) Sub_4BD420(a1p *Struct28[[0x2000]byte], a2 int32) *Struct84 {
	result := a1p.field_3.next
	if result == &a1p.field_3 {
		return nil
	}
	res := result.PromoteUnsafe()
	for res.field_4 != uint32(a2) || res.field_5 == 0 {
		result = res.next
		if result == &a1p.field_3 {
			return nil
		}
		res = result.PromoteUnsafe()
	}
	return res
}

func (m *AudioModule) Sub_4BD470(a1 *Struct28[[0x2000]byte], a2 int32) *Struct84 {
	v2 := m.Sub_4BD420(a1, a2)
	v3 := v2
	if v2 != nil {
		v2.Remove_425920()
		Sub_425900(&a1.field_3, &v3.ListElement)
		return v3
	}

	if m.Sub_486B60(a1.field_0, a2) == 0 {
		return nil
	}
	v5p := sub_4BD2E0(a1.field_2)
	if v5p == nil {
		m.Sub_4BD600(a1)
		v5p = sub_4BD2E0(a1.field_2)
		if v5p == nil {
			m.Sub_486E00(a1.field_0)
			return nil
		}
	}
	v5p.field_4 = uint32(a2)
	v5p.field_13 = a1
	v5p.Init_425770()
	v5p.field_3 = 0
	m.Sub_487C30(&v5p.field_6)
	v5p.field_6.field_5 = &v5p.field_14

	v6 := a1.field_0.Field284
	v10 := a1.field_0.Field284
	if v6 == 0 {
		m.Sub_486AA0(a1.field_0, *(*int32)(unsafe.Pointer(&v5p.field_4)), &v5p.field_14)
		Sub_425900(&a1.field_3, &v5p.ListElement)
		v5p.field_5 = 1
		m.Sub_486E00(a1.field_0)
		return v5p
	}
	for {
		v7 := *(*int32)(unsafe.Pointer(&a1.field_6))
		if v7 > v6 {
			v7 = v6
		}
		v8 := sub_4BD2E0(a1.field_1)
		if v8 == nil {
			found := false
			for m.Sub_4BD600(a1) != 0 {

				v8 = sub_4BD2E0(a1.field_1)
				if v8 != nil {
					found = true
					break
				}
			}
			if !found {
				m.Sub_4BD690(v5p)
				return nil
			}
		}
		// LABEL_17
		m.Sub_487D30(v8, &v8.body_6, v7)
		m.Sub_487C50(&v5p.field_6, v8)
		v9 := m.Sub_486DB0(a1.field_0, (*byte)(unsafe.Pointer(&v8.body_6)), v7)
		if v9 != v7 {
			m.Sub_4BD690(v5p)
			return nil
		}
		v10 = v10 - v9
		if v10 == 0 {
			m.Sub_486AA0(a1.field_0, *(*int32)(unsafe.Pointer(&v5p.field_4)), &v5p.field_14)
			Sub_425900(&a1.field_3, &v5p.ListElement)
			v5p.field_5 = 1
			m.Sub_486E00(a1.field_0)
			return v5p
		}
		v6 = v10
	}
}

func (m *AudioModule) Sub_486DB0(a1p *AudioStructXxx, a2 *byte, a3 int32) int32 {
	var (
		result int32
		v4     int32
	)
	if a1p.Field280 == 0 {
		return 0
	}
	v4 = a3
	if a3 > a1p.Field284 {
		v4 = int32(*(*uint32)(unsafe.Pointer(&a1p.Field284)))
	}
	if v4 <= 0 || (func() bool {
		result = m.nox_binfile_fread_raw_40ADD0(unsafe.Pointer(a2), 1, uint32(v4), *(*unsafe.Pointer)(unsafe.Pointer(&a1p.Field280)))
		return result < 0
	}()) {
		result = 0
	}
	*(*uint32)(unsafe.Pointer(&a1p.Field284)) -= uint32(result)
	return result
}

func (m *AudioModule) Sub_486E00(a1p *AudioStructXxx) unsafe.Pointer {
	var result unsafe.Pointer
	result = *(*unsafe.Pointer)(&a1p.Field272)
	a1p.Field280 = 0
	if result != nil {
		m.nox_fs_close(result)
		result = nil
		*(*uint32)(unsafe.Pointer(&a1p.Field272)) = 0
	}
	return result
}

func (m *AudioModule) Sub_486B60(a1p *AudioStructXxx, a2 int32) int32 {
	var (
		v3  unsafe.Pointer
		v6  unsafe.Pointer
		v7  unsafe.Pointer
		v8  int32
		v9  int32
		v10 int32
		v12 int32
		v13 [8]byte
		v14 [16]byte
		v15 [12]int32
	)
	v12 = 1
	v2p := m.Sub_4866D0(a1p, a2)

	m.Sub_486E00(a1p)
	v3 = *(*unsafe.Pointer)(&a1p.Bagfile268)
	a1p.Field280 = uint32(uintptr(v3))
	*(*uint32)(unsafe.Pointer(&a1p.Field284)) = v2p.Field20
	if m.nox_fs_fseek(v3, int32(v2p.Field16), 0 /* stdio.SEEK_SET */) != 0 {
		v12 = 0
	}
	if v2p.Field20 == 0 {
		v12 = 0
	}
	if a1p.Field276 == 0 {
		return v12
	}
	alloc.Strcpy(unsafe.Pointer(&v15[3]), unsafe.Pointer(&a1p.Path2_8))
	alloc.Strcat(unsafe.Pointer(&v15[3]), unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v2p))))))
	alloc.Strcat(unsafe.Pointer(&v15[3]), unsafe.Pointer(alloc.InternCString(".wav")))
	v6 = m.nox_fs_open(unsafe.Pointer(&v15[3]))
	v7 = v6
	v8 = 0
	*(*uint32)(unsafe.Pointer(&a1p.Field272)) = uint32(uintptr(v6))
	if v6 == nil {
		return v12
	}
	if m.nox_binfile_fread_raw_40ADD0(unsafe.Pointer(&v15[0]), 0xC, 1, v6) != 1 || uint32(v15[0]) != 1179011410 || uint32(v15[2]) != 1163280727 {
		log.Printf("error: '%s' is bad - cannot read\n", &v15[3])
		if *(*uint32)(unsafe.Pointer(&a1p.Field272)) != 0 {
			m.nox_fs_close(*(*unsafe.Pointer)(&a1p.Field272))
			*(*uint32)(unsafe.Pointer(&a1p.Field272)) = 0
		}
		return v12
	}
	if m.nox_binfile_fread_raw_40ADD0(unsafe.Pointer(&v13[0]), 8, 1, v7) != 1 {
		goto LABEL_18
	}
	for {
		if *(*uint32)(unsafe.Pointer(&v13[0])) == 544501094 {
			m.nox_binfile_fread_raw_40ADD0(unsafe.Pointer(&v14[0]), 0x10, 1, v7)
			m.nox_fs_fseek(v7, int32(*(*uint32)(unsafe.Pointer(&v13[4]))-16), 1 /* stdio.SEEK_CUR */)
			goto LABEL_15
		}
		if *(*uint32)(unsafe.Pointer(&v13[0])) == 1635017060 {
			break
		}
		m.nox_fs_fseek(v7, *(*int32)(unsafe.Pointer(&v13[4])), 1 /* stdio.SEEK_CUR */)
	LABEL_15:
		if m.nox_binfile_fread_raw_40ADD0(unsafe.Pointer(&v13[0]), 8, 1, v7) != 1 {
			goto LABEL_18
		}
	}
	v8 = int32(*(*uint32)(unsafe.Pointer(&v13[4])))
LABEL_18:
	v2p.Field28 = 2
	if int32(*(*uint16)(unsafe.Pointer(&v14[12])))/int32(*(*uint16)(unsafe.Pointer(&v14[2]))) == 2 {
		v2p.Field28 = 6
	}
	if int32(*(*uint16)(unsafe.Pointer(&v14[2]))) == 2 {
		v9 = int32(v2p.Field28)
		*((*uint8)(unsafe.Pointer(&v9))) = uint8(int8(v9 | 1))
		v2p.Field28 = uint32(v9)
	}
	v2p.Field24 = *(*uint32)(unsafe.Pointer(&v14[4]))
	v10 = int32(*(*uint32)(unsafe.Pointer(&a1p.Field272)))
	*(*uint32)(unsafe.Pointer(&a1p.Field284)) = uint32(v8)
	a1p.Field280 = uint32(v10)
	return 1
}

func (m *AudioModule) Sub_486FA0(a1p *Struct587000_94032) {
	result := m.Sub_486FE0(a1p)
	v2 := result
	if result != nil {
		a1p.field_3 |= 1
		m.Sub_487050(v2)
		if a1p.field_2&2 != 0 {
			*m.Externs.Ptr_uint32_5d4594_1193332 = 1
		}
		result = v2
	}
}

func (m *AudioModule) Sub_452EB0(a1 *int32) int32 {
	var result int32
	result = *a1
	if *a1 != 0 && (uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*2))) != *(*uint32)(unsafe.Pointer(uintptr(result + 36))) || uint32(*(*int32)(unsafe.Add(unsafe.Pointer(a1), 4*1))) != *(*uint32)(unsafe.Pointer(uintptr(result + 280)))) {
		result = 0
		*a1 = 0
	}
	return result
}

func (m *AudioModule) Sub_4522A0(a1p *Struct200) int32 {
	var result int32
	if *m.Externs.Dword_5d4594_1045432 != 0 {
		result = int32(a1p.field_16)
	} else {
		result = 0
	}
	return result
}

func (m *AudioModule) Sub_45A9B0(a1p, a2p unsafe.Pointer /* *nox_drawable */) {
	var (
		a1     = int32(uintptr(a1p))
		a2     = int32(uintptr(a2p))
		v2     int32
		v3     int32
		result *int32 = nil
		v7     int32
		v8     int32
		v9     int32
		v10    int32
		v11    int64
		v12    int32
		v13    *int32
		v14    *int32
		v15    *int32
		v16    int32
		v18    *int32
	)
	v2 = a1
	v3 = 0
	v16 = 0
	v4p := m.Nox_xxx_draw_452270(int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 492)))))
	v18 = (*int32)(m.nox_draw_getViewport_437250())
	if v4p != nil && v18 != nil {
		if *(*uint32)(unsafe.Pointer(uintptr(a1 + 120)))&0x1000000 != 0 && (int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 280))))&0xC) == 0 {
			v7 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 12))) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
			v8 = int32(*(*uint32)(unsafe.Pointer(uintptr(a2 + 16))) - *(*uint32)(unsafe.Pointer(uintptr(a1 + 16))))
			v9 = m.Sub_4522A0(v4p)
			v10 = v9
			if v7 < v9 && v8 < v9 && v9 > 0 {
				v11 = int64(math.Sqrt(float64(v8*v8 + v7*v7 + 1)))
				if int32(v11) < v10 {
					v12 = (v10 - int32(v11)) * 100 / v10
					v3 = v12
					if v12 <= 100 {
						if v12 < 0 {
							v3 = 0
						}
					} else {
						v3 = 100
					}
					v16 = (*(*int32)(unsafe.Pointer(uintptr(a1 + 12))) - *(*int32)(unsafe.Add(unsafe.Pointer(v18), 4*6)) - *v18) * 50 / (*m.Externs.Nox_win_width / 2)
				}
			}
			v2 = a1
		}
		v13 = (*int32)(unsafe.Pointer(uintptr(v2 + 496)))
		result = (*int32)(unsafe.Pointer(uintptr(m.Sub_452EB0(v13))))
		v14 = result
		if v3 != 0 {
			if result != nil {
				m.Sub_452FE0((*Struct576)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))), v16)
				m.Sub_452F50((*Struct576)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v14)))))), v3)
			} else {
				result = (*int32)(unsafe.Pointer(m.Nox_xxx_draw_452300(v4p)))
				v15 = result
				if result != nil {
					m.Sub_452EE0((*Struct576)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(result)))))), v3)
					m.Sub_452F80((*Struct576)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))), v16)
					m.Sub_452E90((*uint32)(unsafe.Pointer(v13)), (*Struct576)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(v15)))))))
				}
			}
		} else if result != nil {
			m.Sub_4523D0((*Struct576)(unsafe.Pointer(result)))
		}
	}
}
