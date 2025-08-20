package audio

import (
	unsafe "unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type Struct24[T any] struct {
	field_0 uint32
	field_1 uint32
	field_2 uint32
	field_3 uint32
	field_4 uint32
	field_5 uint32
	body_6  T
}

var _ = [1]struct{}{}[2024-unsafe.Sizeof(Struct24[[2000]byte]{})]

type Struct84 struct {
	field_0 [21]uint32
}

var _ = [1]struct{}{}[84-unsafe.Sizeof(Struct84{})]

type Struct28[T any] struct {
	field_0 uint32
	field_1 *FreeList[Struct24[T]]
	field_2 *FreeList[Struct84]
	field_3 UnknownListElement
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
	// There are more items, but cannot represent in Go...
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
	field_22   ListElement[Struct576Field3, *Struct576Field3]
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
	field_74      unsafe.Pointer
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
		*m.Externs.Dword_5d4594_1045424 = uint32(uintptr(unsafe.Pointer(m.Sub_4BD340(int32(uintptr(a3p)), 0x100000, 200, 0x2000))))
		*m.Externs.Dword_5d4594_1045436 = uint32(uintptr(unsafe.Pointer(m.Sub_4BD280(200, 576))))
	}
	if *m.Externs.Dword_5d4594_1045424 == 0 || *m.Externs.Dword_5d4594_1045420 == 0 || *m.Externs.Dword_5d4594_1045428 == nil || *m.Externs.Dword_5d4594_1045436 == 0 {
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
	a1p.field_10[a1p.field_42] = unsafe.Pointer(m.Sub_4BD470(
		(**uint32)(unsafe.Pointer(*(***uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045424)))),
		int32(a1p.field_9.field_32[a2]),
	))
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

func (m *AudioModule) sub_451FE0(a1 *Struct576) int32 {
	a1.Remove_425920()
	a1.field_70 = 0
	return m.Sub_4BD300(*(**uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)), int32(uintptr(unsafe.Pointer(a1))))
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
	var result *Struct312
	result = a1p.field_44
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
	var v3 unsafe.Pointer
	var v4 int32
	v1 := a1p.field_44
	if a1p != v1.field_38 {
		return 0
	}
	v3 = a1p.field_74
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
	var v1p = (*Struct576)(unsafe.Pointer(m.Sub_4BD2E0(*(***uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)))))
	if v1p == nil {
		m.sub_452230()
		v1p = (*Struct576)(unsafe.Pointer(m.Sub_4BD2E0(*(***uint32)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)))))
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
		result = m.sub_4BD710(int32(*(*uint32)(unsafe.Pointer(&a1.field_10[v9]))))
	} else if v3&1 != 0 {
		if v1.field_15 != 0 && (func() bool {
			v4 := a1.field_109 + 1
			a1.field_109 = v4
			return v4 >= v1.field_15
		}()) {
			result = 0
		} else {
			result = m.sub_451CA0(a1)
		}
	}
	return result
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
	res = (*Struct312)(m.Sub_452810(int32(*&v1p.field_12+uint32(v3)), 0))
	a1.field_44 = res
	if res != nil {
		v4 = int32(m.nox_common_randomIntMinMax_415FF0(int(v1p.field_19), int(*&v1p.field_20), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1482))
		a1.field_44.timerGroup_4.Timers[1].SetRaw(uint32(v4 + 100))
		m.sub_4BDB20(a1.field_44)
		a1.field_44.field_38 = a1
		a1.field_44.field_35 = m.Externs.Sub_452770_ptr
		a1.field_44.field_36 = m.Externs.Sub_4526F0_ptr
		a1.field_44.field_37 = m.Externs.Sub_4526D0_ptr
		a1.field_7 = 1
		a1.field_44.field_28 = &a1.timerGroup_46
		if int32(v1p.field_1)&8 != 0 {
			v5 = int32(m.nox_common_randomIntMinMax_415FF0(int(*&v1p.field_17), int(*&v1p.field_18), unsafe.Pointer(alloc.InternCString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c")), 1497))
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
		m.Sub_4BD3C0(*(*unsafe.Pointer)(unsafe.Pointer(m.Externs.Dword_5d4594_1045424)))
		*m.Externs.Dword_5d4594_1045424 = 0
	}
	if *m.Externs.Dword_5d4594_1045436 != 0 {
		m.Sub_4BD2D0(*(*unsafe.Pointer)(unsafe.Pointer(m.Externs.Dword_5d4594_1045436)))
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
			result = int32(v6.field_7)
			if result == 1 {
				m.sub_451DC0(v6)
				v8 = m.sub_451CA0(v6)
				v6.field_74 = unsafe.Pointer(uintptr(v8))
				if v8 == 0 {
					for {
						if m.sub_452120(v6) == false {
							break
						}
						v7 = v6.next
						m.sub_451DC0(v6)
						v9 = m.sub_451CA0(v6)
						v6.field_74 = unsafe.Pointer(uintptr(v9))
						if v9 != 0 {
							break
						}
					}
				}
				v10 = m.sub_451CA0(v6)
				v6.field_74 = unsafe.Pointer(uintptr(v10))
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
	var result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))++
	return result
}

func (m *AudioModule) sub_4BD660(a1p unsafe.Pointer) int32 {
	var result = int32(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(a1p)) + 12))) - 1)
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
			a1p.field_74 = 0
			return result
		}
	} else {
		if a1p.field_73 != nil {
			v3 := a1p.field_73.NextSafe_425940()
			a1p.field_73 = v3
			if v3 != nil {
				a1p.field_74 = *(*uint32)(unsafe.Pointer(unsafe.Add(unsafe.Pointer(v3), 12)))
				v4 = int32(*(*uint32)(unsafe.Pointer(unsafe.Add(unsafe.Pointer(v3), 16))))
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
	if a1p.field_32 != 0 {
		if a1p.field_32 != -1 {
			a1p.field_32--
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
	a2p.field_31 &= 0xFA
	a2p.field_32 = 0
	a2p.timerGroup_4.Init()
	v2 := a2p.field_36
	if v2 != nil {
		result = int32(ccall.CallIntPtr(v2, unsafe.Pointer(a2p)))
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
	var a2 = (*uint32)(a2p)
	var (
		v3 int32
		v4 *uint32
		v5 int32
	)
	a1p.field_72 = unsafe.Pointer(a2)
	if a2 != nil {
		v2 := m.sub_487C80(int32(uintptr(unsafe.Pointer(a2))))
		a1p.field_73 = v2
		if v2 != nil {
			a1p.field_74 = *(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(v2)) + 12))
			v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(v2)) + 16)))
			a1p.field_75 = uint32(v3)
			a1p.field_76 = uint32(v3)
			*a2 = 0
		} else {
			v4 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1p.field_72)))))
			a1p.field_74 = *v4
			v5 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1)))
			a1p.field_75 = uint32(v5)
			a1p.field_76 = uint32(v5)
		}
	}
}

func (m *AudioModule) sub_487C80(a1 int32) *UnknownListElement {
	v1 := (*ListElement[UnknownListElement, *UnknownListElement])(unsafe.Pointer(uintptr(a1 + 8)))
	return v1.NextSafe_425940()
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

func (m *AudioModule) Sub_4BDA80(a1 *Struct312) int32 {
	var (
		result int32 = 0
	)
	if int32(*(*uint8)(unsafe.Pointer(&a1.field_31)))&5 != 0 {
		ccall.CallVoidInt(a1.field_43.field_4, int(int32(uintptr(unsafe.Pointer(a1)))))
	}
	result2 := *(*uint32)(unsafe.Pointer(&a1.field_37))
	if result2 != 0 {
		result = int32(ccall.CallIntInt(unsafe.Pointer(uintptr(result2)), int(int32(uintptr(unsafe.Pointer(a1))))))
	}
	*(*uint32)(unsafe.Pointer(&a1.field_72)) = 0
	return result
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
		v15 *timer.TimerGroup
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
		v15 = &a3p.TimerGroup_22
		a3p.TimerGroup_22.Update()

		if a3p.field_46 != nil {
			a3p.field_46.Update()
			v17 = a3p.field_46.IsUpdated()
		}
		if a3p.field_46 == nil || v17 == false {
			v17 = a3p.TimerGroup_22.IsUpdated()
		}
		a3p.field_62 = v3

		v12 := a3p.field_50.next
		if v12 != &a3p.field_50 {
			for {
				v13 := v12.next
				v12p := v12.PromoteUnsafe()
				if (v12p.field_31&1) != 0 && v12p.field_72 != nil {
					v12p.timerGroup_4.Update()
					if v17 || v12p.timerGroup_4.IsUpdated() || (v12p.field_29 != nil && v12p.field_29.IsUpdated()) || (v12p.field_28 != nil && v12p.field_28.IsUpdated()) {
						m.sub_4BD840(v12p)
						ccall.CallVoidPtr(v12p.field_43.field_8, unsafe.Pointer(v12p))
					}
				}
				v12 = v13
				if v13 == &a3p.field_50 {
					break
				}
			}
			v11 = v15
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
	var v1p *Struct264 = a3p.field_33

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
	if int32(*(*uint8)(unsafe.Pointer(&a2p.field_31)))&5 != 0 {
		return -2146500608
	}
	if a2p.field_72 == nil {
		return -2147024896
	}
	a2p.timerGroup_4.Update()
	m.sub_4BD840(a2p)
	result = int32(ccall.CallIntPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a2p.field_43)) + 12))), unsafe.Pointer(a2p)))
	if result == 0 {
		a2p.field_31 |= 1
	}
	return result
}

func (m *AudioModule) sub_4871C0(a1p *Struct88, a2 int32, a3 *[7]uint32) *Struct264 {
	var (
		v3 *Struct587000_94032
	)
	v3 = a1p.field_3
	v4p, _ := alloc.Calloc(1, 0x108)
	v4m := (*Struct264)(v4p)
	alloc.Memset(unsafe.Pointer(v4m), 0, 0x108)
	v4m.field_0.Init_425770()
	v4m.field_6 = uint32(a2)
	v4m.field_5 = a1p
	v4m.field_4 = 0
	*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a1p.field_4))))))++
	*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a1p.field_6[a2])))))) = uint32(uintptr(unsafe.Pointer(v4m)))
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
		m.Sub_487680((*Struct264)(*m.Externs.Dword_5d4594_805984))
		*m.Externs.Dword_5d4594_805984 = nil
	}
}

func (m *AudioModule) Sub_487680(lpMem_ *Struct264) {
	var lpMem unsafe.Pointer = unsafe.Pointer(lpMem_)
	m.sub_4876A0((*Struct264)(unsafe.Pointer((**uint32)(lpMem))))
	m.sub_4872C0((*Struct264)(lpMem))
}

func (m *AudioModule) Sub_431290() {
	if *m.Externs.Dword_5d4594_805984 != nil {
		m.sub_487970((*Struct264)(*m.Externs.Dword_5d4594_805984), -1)
	}
}

func (m *AudioModule) sub_487970(a1_ *Struct264, a2 int32) *Struct312 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		a1x    *Struct312
		result *Struct312
		v3     *Struct312
		v4     int32
	)
	result = m.sub_4877D0((*Struct264)(unsafe.Pointer(uintptr(a1))), &a1x)
	v3 = result
	if result != nil {
		v4 = a2
		for {
			result = m.sub_4877F0(&a1x)
			v5 := result
			if v4 == -1 || *(*int32)(unsafe.Pointer(&v3.field_3)) == v4 {
				result = (*Struct312)(unsafe.Pointer(uintptr(m.Sub_4BDA80(v3))))
			}
			v3 = v5
			if v5 == nil {
				break
			}
		}
	}
	return result
}

func (m *AudioModule) sub_487590(a1_ *Struct264, a2 *[7]uint32) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		result int32
	)
	result = a1
	alloc.Memcpy(unsafe.Pointer(&a1_.field_15), unsafe.Pointer(a2), 0x1C)
	return result
}

func (m *AudioModule) sub_4872C0(a1p *Struct264) {
	var (
		// lpMem unsafe.Pointer = unsafe.Pointer(a1p)
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
			if *(*int32)(unsafe.Add(unsafe.Pointer(&v3.field_31), 0))&0x15 != 0 && *(*int32)(unsafe.Add(unsafe.Pointer(&v3.field_30), 0)) > a1 {
				return nil
			}
			m.Sub_4BDA80(v3)
			v2.field_29 = (*m.Externs.Dword_587000_127004)
			v2.field_30 = a1
			if int32(a2)&1 != 0 {
				v2.field_32 = -1
			} else {
				v2.field_32 = 0
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
		if *(*int32)(unsafe.Pointer(&result.field_3)) == a2 {
			if (*(*int32)(unsafe.Pointer(&result.field_31)) & 0x15) == 0 {
				return result
			}
			v6 = *(*int32)(unsafe.Pointer(&result.field_30))
			if *(*int32)(unsafe.Pointer(&result.field_31))&1 != 0 {
				if v6 >= v3 {
					if v6 == v3 {
						v7 = uint32(*(*int32)(unsafe.Pointer(&result.timerGroup_44.Timers[0].Current)))
						if v7 < v2 && v2-v7 >= 0x666 {
							v3 = *(*int32)(unsafe.Pointer(&result.field_30))
							v4 = result
							v2 = uint32(*(*int32)(unsafe.Pointer(&result.timerGroup_44.Timers[0].Current)))
						}
					}
				} else {
					v2 = uint32(*(*int32)(unsafe.Pointer(&result.timerGroup_44.Timers[0].Current)))
					v3 = *(*int32)(unsafe.Pointer(&result.field_30))
					v4 = result
				}
			} else if v6 < v8 {
				v8 = *(*int32)(unsafe.Pointer(&result.field_30))
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
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*5))
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
	*memmap.PtrUint32(0x5D4594, 1193332) = 0
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

func (m *AudioModule) Sub_425960(a1 int32) int32 {
	if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) + 8))) != *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) {
		return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	}
	return 0
}

func (m *AudioModule) Sub_487D60(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 20))) = 0
	return result
}

func (m *AudioModule) Sub_4BD300(a1 *uint32, a2 int32) int32 {
	var result int32
	result = a2 - 4
	*(*uint32)(unsafe.Pointer(uintptr(a2 - 4))) = *a1
	*a1 = uint32(a2 - 4)
	return result
}

func (m *AudioModule) Sub_4BD680(a1 int32) int32 {
	return int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
}

func (m *AudioModule) Sub_487C50(a1 int32, a2 *uint32) int32 {
	var result int32
	m.nox_common_list_append_4258E0((unsafe.Pointer(uintptr(a1 + 8))), (unsafe.Pointer(a2)))
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4)) + *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))))
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = uint32(result)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*5)) = uint32(a1)
	return result
}

func (m *AudioModule) Sub_4BD690(a1 int32) int32 {
	var i **uint32
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) != uint32(a1) {
		m.nox_common_list_remove_425920(unsafe.Pointer(uintptr(a1)))
	}
	for i = (**uint32)(m.nox_common_list_getNext_425940((unsafe.Pointer(uintptr(a1 + 32))))); i != nil; i = (**uint32)(m.nox_common_list_getNext_425940((unsafe.Pointer(uintptr(a1 + 32))))) {
		m.nox_common_list_remove_425920(unsafe.Pointer(i))
		m.Sub_487D60(int32(uintptr(unsafe.Pointer(i))))
		m.Sub_4BD300(*(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) + 4))), int32(uintptr(unsafe.Pointer(i))))
	}
	return m.Sub_4BD300(*(**uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 52))) + 8))), a1)
}

func (m *AudioModule) Sub_4BD600(a1 int32) int32 {
	var v1 int32
	v1 = m.Sub_425960(a1 + 12)
	if v1 == 0 {
		return 0
	}
	for m.Sub_4BD680(v1) != 0 {
		v1 = m.Sub_425960(v1)
		if v1 == 0 {
			return 0
		}
	}
	m.Sub_4BD690(v1)
	return 1
}

func (m *AudioModule) Sub_425900(a1 *uint32, a2 *uint32) *uint32 {
	var result *uint32
	result = a2
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*1)) = uint32(uintptr(unsafe.Pointer(a1)))
	*a2 = *a1
	*a1 = uint32(uintptr(unsafe.Pointer(a2)))
	*(*uint32)(unsafe.Pointer(uintptr(*a2 + 4))) = uint32(uintptr(unsafe.Pointer(a2)))
	return result
}

func (m *AudioModule) Sub_487D30(a1 *uint32, a2 int32, a3 int32) *uint32 {
	var result *uint32
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)) = uint32(a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3)) = uint32(a2)
	((*UnknownListElement)(unsafe.Pointer(a1))).Init_425770()
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) = 0
	return result
}

func (m *AudioModule) Sub_4BD280(a1 int32, a2 int32) *uint32 {
	var (
		v2     int32
		result *uint32
		v4     *uint32
		v5     int32
	)
	v2 = a2 + 4
	res, _ := alloc.Calloc(1, uintptr(a1*(a2+4)+4))
	result = (*uint32)(res)
	if result != nil {
		v4 = (*uint32)(unsafe.Add(unsafe.Pointer(result), 4*1))
		*result = uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(result), 4*1)))))
		if a1 != 1 {
			v5 = a1 - 1
			for {
				v5--
				*v4 = uint32(uintptr(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v4))), v2)))))
				v4 = (*uint32)(unsafe.Pointer((*byte)(unsafe.Add(unsafe.Pointer((*byte)(unsafe.Pointer(v4))), v2))))
				if v5 == 0 {
					break
				}
			}
		}
		*v4 = 0
	}
	return result
}

func (m *AudioModule) Sub_4BD2E0(a1 **uint32) *uint32 {
	var (
		result *uint32
		v2     *uint32
	)
	result = *a1
	if *a1 != nil {
		v2 = (*uint32)(unsafe.Pointer(uintptr(*result)))
		result = (*uint32)(unsafe.Add(unsafe.Pointer(result), 4*1))
		*a1 = v2
	}
	return result
}

func (m *AudioModule) Sub_4BD2D0(lpMem unsafe.Pointer) {
	alloc.FreePtr(lpMem)
}

func (m *AudioModule) Sub_4866D0(a1 *uint32, a2 int32) int32 {
	return int32(*a1 + uint32(a2*36))
}

func (m *AudioModule) Sub_487C30(a1 *uint32) {
	*a1 = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*6)) = 0
	((*UnknownListElement)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2))))).Clear_425760()
}

func (m *AudioModule) Sub_487D00(a1 *uint32) int32 {
	var (
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*1)))
	result = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*2)) * *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3)) * *(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) = uint32(result)
	if v1 == 1 {
		result >>= 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*5)) = uint32(result)
	}
	return result
}

func (m *AudioModule) Sub_4BD340(a1 int32, a2 int32, a3 int32, a4 int32) *uint32 {
	var v4 *uint32
	res, _ := alloc.Calloc(1, 0x1C)
	v4 = (*uint32)(res)
	alloc.Memset(unsafe.Pointer(v4), 0, 0x1C)
	*v4 = uint32(a1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*6)) = uint32(a4)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1)) = uint32(uintptr(unsafe.Pointer(m.Sub_4BD280(a2/(a4+24), a4+24))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*2)) = uint32(uintptr(unsafe.Pointer(m.Sub_4BD280(a3, 84))))
	((*UnknownListElement)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*3))))).Clear_425760()
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*1)) != 0 && *(*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*2)) != 0 {
		return v4
	}
	m.Sub_4BD3C0(unsafe.Pointer(v4))
	return nil
}

func (m *AudioModule) Sub_4BD3C0(lpMem unsafe.Pointer) {
	var i int32
	for i = int32(uintptr(unsafe.Pointer(m.nox_common_list_getNext_425940((unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer((*int32)(lpMem)), 4*3)))))))); i != 0; i = int32(uintptr(unsafe.Pointer(m.nox_common_list_getNext_425940((unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer((*int32)(lpMem)), 4*3)))))))) {
		m.Sub_4BD690(i)
	}
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), 4*1))) != 0 {
		m.Sub_4BD2D0(*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(lpMem)), unsafe.Sizeof(unsafe.Pointer(nil))*1))))
	}
	if *((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), 4*2))) != 0 {
		m.Sub_4BD2D0(*((*unsafe.Pointer)(unsafe.Add(unsafe.Pointer((*unsafe.Pointer)(lpMem)), unsafe.Sizeof(unsafe.Pointer(nil))*2))))
	}
	alloc.FreePtr(lpMem)
}

func (m *AudioModule) Sub_486AA0(a1 *uint32, a2 int32, a3 *uint32) int32 {
	var (
		v3     *uint32
		result int32
	)
	v3 = (*uint32)(unsafe.Pointer(uintptr(m.Sub_4866D0(a1, a2))))
	*a3 = 4
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*2)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*6))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*3)) = uint32(bool2int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*7))&1) != 0) + 1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*6)) = *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*8))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*7))&8 != 0 {
		result = 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*1)) = 2
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*4)) = 2
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*1)) = 0
		result = bool2int32((*(*uint32)(unsafe.Add(unsafe.Pointer(v3), 4*7))&4) != 0) + 1
		*(*uint32)(unsafe.Add(unsafe.Pointer(a3), 4*4)) = uint32(result)
	}
	return result
}

func (m *AudioModule) Sub_4BD420(a1 int32, a2 int32) *uint32 {
	var result *uint32
	result = *(**uint32)(unsafe.Pointer(uintptr(a1 + 12)))
	if result == (*uint32)(unsafe.Pointer(uintptr(a1+12))) {
		return nil
	}
	for *(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*4)) != uint32(a2) || *(*uint32)(unsafe.Add(unsafe.Pointer(result), 4*5)) == 0 {
		result = (*uint32)(unsafe.Pointer(uintptr(*result)))
		if result == (*uint32)(unsafe.Pointer(uintptr(a1+12))) {
			return nil
		}
	}
	return result
}

func (m *AudioModule) Sub_4BD470(a1 **uint32, a2 int32) *uint32 {
	v2 := m.Sub_4BD420(int32(uintptr(unsafe.Pointer(a1))), a2)
	v3 := v2
	if v2 != nil {
		m.nox_common_list_remove_425920(unsafe.Pointer(v2))
		m.Sub_425900((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1))+4*3)), v3)
		return v3
	}

	if m.sub_486B60(int(uintptr(unsafe.Pointer(*a1))), int(a2)) == 0 {
		return nil
	}
	v5 := unsafe.Pointer(m.Sub_4BD2E0(*(***uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1)) + 4*2))))
	if v5 == nil {
		m.Sub_4BD600(int32(uintptr(unsafe.Pointer(a1))))
		v5 = unsafe.Pointer(m.Sub_4BD2E0(*(***uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1)) + 4*2))))
		if v5 == nil {
			m.sub_486E00(*(*int)(unsafe.Pointer(a1)))
			return nil
		}
	}
	*(*uint32)(unsafe.Add(v5, 4*4)) = uint32(a2)
	*(*uint32)(unsafe.Add(v5, 4*13)) = uint32(uintptr(unsafe.Pointer(a1)))
	((*UnknownListElement)(v5)).Init_425770()
	*(*uint32)(unsafe.Add(v5, 4*3)) = 0
	m.Sub_487C30((*uint32)(unsafe.Add(v5, 4*6)))
	*(*uint32)(unsafe.Add(v5, 4*11)) = uint32(uintptr(unsafe.Add(v5, 4*14)))

	v6 := *(*int32)(unsafe.Add(unsafe.Pointer(*a1), 4*71))
	v10 := *(*int32)(unsafe.Add(unsafe.Pointer(*a1), 4*71))
	if v6 == 0 {
		m.Sub_486AA0(*a1, *(*int32)(unsafe.Add(v5, 4*4)), (*uint32)(unsafe.Add(v5, 4*14)))
		m.Sub_425900((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1))+4*3)), (*uint32)(v5))
		*(*uint32)(unsafe.Add(v5, 4*5)) = 1
		m.sub_486E00(int(uintptr(unsafe.Pointer(*a1))))
		return (*uint32)(v5)
	}
	for {
		v7 := (*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1)) + 4*6)))
		if v7 > v6 {
			v7 = v6
		}
		v8 := unsafe.Pointer(m.Sub_4BD2E0(*(***uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1)) + 4*1))))
		if v8 == nil {
			found := false
			for {
				if m.Sub_4BD600(int32(uintptr(unsafe.Pointer(a1)))) == 0 {
					break
				}
				v8 = unsafe.Pointer(m.Sub_4BD2E0(*(***uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1)) + 4*1))))
				if v8 != nil {
					found = true
					break
				}
			}
			if !found {
				m.Sub_4BD690(int32(uintptr(v5)))
				return nil
			}
		}
		// LABEL_17
		m.Sub_487D30((*uint32)(v8), int32(uintptr(unsafe.Add(v8, 24))), v7)
		m.Sub_487C50(int32(uintptr(unsafe.Add(v5, 4*6))), (*uint32)(v8))
		v9 := int32(m.sub_486DB0(int(uintptr(unsafe.Pointer(*a1))), unsafe.Add(v8, 24), int(v7)))
		if v9 != v7 {
			m.Sub_4BD690(int32(uintptr(v5)))
			return nil
		}
		v10 = v10 - v9
		if v10 == 0 {
			m.Sub_486AA0(*a1, *(*int32)(unsafe.Add(v5, 4*4)), (*uint32)(unsafe.Add(v5, 4*14)))
			m.Sub_425900((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(a1))+4*3)), (*uint32)(v5))
			*(*uint32)(unsafe.Add(v5, 4*5)) = 1
			m.sub_486E00(int(uintptr(unsafe.Pointer(*a1))))
			return (*uint32)(v5)
		}
		v6 = v10
	}
}
