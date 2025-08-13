package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func (m *Struct264Module) sub_4871C0(a1p *Struct88, a2 int32, a3 *[7]uint32) *Struct264 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
	var (
		v3 int32
		v4 *uint32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a1p.field_3)))))))
	v4p, _ := alloc.Calloc(1, 0x108)
	v4m := (*Struct264)(v4p)
	v4 = (*uint32)(v4p)
	alloc.Memset(unsafe.Pointer(v4m), 0, 0x108)
	m.sub_425770(unsafe.Pointer(&v4m.field_0))
	v4m.field_6 = uint32(a2)
	v4m.field_5 = unsafe.Pointer(uintptr(a1))
	v4m.field_4 = 0
	*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a1p.field_4))))))++
	*(*uint32)(unsafe.Pointer(uintptr(int32(uintptr(unsafe.Pointer(&a1p.field_6[a2])))))) = uint32(uintptr(unsafe.Pointer(v4m)))
	v4m.field_64 = *(*unsafe.Pointer)(unsafe.Pointer(uintptr(v3 + 36)))
	v4m.field_50.Clear_425760()
	v4m.TimerGroup_22.Init()
	v4m.field_53 = 0
	v4m.field_56 = 33
	v4m.field_60 = 0
	v4m.field_58 = 0
	v4m.field_62 = 0
	v4m.field_54 = m.externs.Sub_4873C0_ptr
	m.nullsub_10(uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*15))))))
	m.nullsub_10(uint32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v4), 4*8))))))
	if a3 != nil {
		m.sub_487590(v4m, a3)
	}
	if ccall.CallIntPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(v3 + 28))), unsafe.Pointer(v4m)) == 0 {
		return v4m
	}
	if v4 != nil {
		m.sub_4872C0(v4m)
	}
	return nil
}

func (m *Struct264Module) Sub_487150(aa1 int32, a2 *[7]uint32) *Struct264 {
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

func (m *Struct264Module) sub_487310(a1_ *Struct264) int32 {
	var (
		result int32
	)
	(*m.externs.Dword_587000_155144).field_6++
	(*m.externs.Dword_587000_155144).field_3.Append_4258E0(&a1_.field_0)
	result = int32((*m.externs.Dword_587000_155144).field_6 - 1)
	(*m.externs.Dword_587000_155144).field_6 = uint32(result)
	if result < 0 {
		(*m.externs.Dword_587000_155144).field_6 = 0
	}
	return result
}

func (m *Struct264Module) Sub_431270() {
	if *m.externs.Dword_5d4594_805984 != nil {
		m.Sub_487680((*Struct264)(*m.externs.Dword_5d4594_805984))
		*m.externs.Dword_5d4594_805984 = nil
	}
}

func (m *Struct264Module) Sub_487680(lpMem_ *Struct264) {
	var lpMem unsafe.Pointer = unsafe.Pointer(lpMem_)
	m.sub_4876A0((*Struct264)(unsafe.Pointer((**uint32)(lpMem))))
	m.sub_4872C0((*Struct264)(lpMem))
}

func (m *Struct264Module) Sub_431290() {
	if *m.externs.Dword_5d4594_805984 != nil {
		m.sub_487970((*Struct264)(*m.externs.Dword_5d4594_805984), -1)
	}
}

func (m *Struct264Module) sub_487970(a1_ *Struct264, a2 int32) *Struct312 {
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
			if v4 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*3)) == v4 {
				result = (*Struct312)(unsafe.Pointer(uintptr(m.sub_4BDA80(v3))))
			}
			v3 = v5
			if v5 == nil {
				break
			}
		}
	}
	return result
}

func (m *Struct264Module) sub_487590(a1_ *Struct264, a2 *[7]uint32) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		result int32
	)
	result = a1
	alloc.Memcpy(unsafe.Pointer(&a1_.field_15), unsafe.Pointer(a2), 0x1C)
	return result
}

func (m *Struct264Module) sub_4872C0(a1p *Struct264) {
	var (
		// lpMem unsafe.Pointer = unsafe.Pointer(a1p)
		v1 int32
		v2 int32
	)
	m.sub_487910(a1p, -1)
	ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1p.field_5) + 12)) + 32))), unsafe.Pointer(a1p))
	*(*uint32)(unsafe.Pointer(uintptr(a1p.field_5) + uintptr(a1p.field_6*4) + uintptr(24))) = 0
	v1 = int32(uintptr(a1p.field_5))
	v2 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 16))) = uint32(v2)
	if v2 < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1p.field_5) + uintptr(16))) = 0
	}
	alloc.Free(a1p)
}

func (m *Struct264Module) sub_4876A0(a1 *Struct264) {
	var (
		result int32
	)
	(*m.externs.Dword_587000_155144).field_6++
	a1.field_0.Remove_425920()
	result = int32((*m.externs.Dword_587000_155144).field_6 - 1)
	(*m.externs.Dword_587000_155144).field_6 = uint32(result)
	if result < 0 {
		(*m.externs.Dword_587000_155144).field_6 = 0
	}
}

func (m *Struct264Module) sub_487910(a1p *Struct264, a2 int32) int32 {
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
			m.sub_4BDA60(unsafe.Pointer(v2))
		}
		v2 = v4
		if v4 == nil {
			break
		}
	}
	return 0
}

func (m *Struct264Module) Sub_452810(a1 int32, a2 int8) *Struct312 {
	var (
		v2 *Struct312
		v3 *Struct312
	)
	v2 = nil
	if *m.externs.Dword_5d4594_1045428 != nil {
		v3 = m.sub_487810(*m.externs.Dword_5d4594_1045428, 1)
		v2 = v3
		if v3 != nil {
			if *(*int32)(unsafe.Add(unsafe.Pointer(&v3.field_31), 0))&0x15 != 0 && *(*int32)(unsafe.Add(unsafe.Pointer(&v3.field_30), 0)) > a1 {
				return nil
			}
			m.sub_4BDA80(v3)
			v2.field_29 = (*m.externs.Dword_587000_127004)
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

func (m *Struct264Module) sub_487810(a1p *Struct264, a2 int32) *Struct312 {
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
		if *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*3)) == a2 {
			if (*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*31)) & 0x15) == 0 {
				return result
			}
			v6 = *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*30))
			if *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*31))&1 != 0 {
				if v6 >= v3 {
					if v6 == v3 {
						v7 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*45)))
						if v7 < v2 && v2-v7 >= 0x666 {
							v3 = *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*30))
							v4 = result
							v2 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*45)))
						}
					}
				} else {
					v2 = uint32(*(*int32)(unsafe.Add(unsafe.Pointer(result), 4*45)))
					v3 = *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*30))
					v4 = result
				}
			} else if v6 < v8 {
				v8 = *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*30))
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

func (m *Struct264Module) sub_4877D0(a1p *Struct264, a2 **Struct312) *Struct312 {
	result := a1p.field_50.FirstSafe_425890()
	*a2 = result
	return result
}

func (m *Struct264Module) sub_4877F0(a1 **Struct312) *Struct312 {
	if *a1 != nil {
		*a1 = (*a1).NextSafe_4258A0()
	}
	return *a1
}

func (m *Struct264Module) sub_487360(a1 int32, a2 **Struct88, a3 *int32) {
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

func (m *Struct264Module) sub_487750(a1p *Struct264) *Struct312 {
	if a1p.field_48 >= a1p.field_49 {
		return nil
	}
	v1p := m.sub_4BD720(a1p)
	v2p := v1p
	if v1p == nil {
		return nil
	}
	m.sub_486E30(a1p, v1p)
	return v2p
}

func (m *Struct264Module) Sub_487790(a1p *Struct264, a2 int32) int32 {
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

func (m *Struct264Module) sub_486E30(a1p *Struct264, a2p *Struct312) int32 {
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

func (m *Struct264Module) Sub_486FE0(a1 unsafe.Pointer) *Struct88 {
	v1p, _ := alloc.Calloc(1, 0x58)
	v1pp := (*Struct88)(v1p)
	alloc.Memset(v1p, 0, 0x58)
	v1pp.field_0.Init_425770()
	v1pp.field_4 = 0
	v1pp.field_3 = a1

	if ccall.CallIntPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1) + 20)), unsafe.Pointer(v1pp)) == 0 {
		return v1pp
	}
	if v1pp != nil {
		m.sub_487030(v1pp)
	}
	return nil
}

func (m *Struct264Module) sub_487030(a1p *Struct88) {
	// var lpMem unsafe.Pointer = unsafe.Pointer(a1p)
	ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1p.field_3) + 24)), unsafe.Pointer(a1p))
	*(*uint32)(unsafe.Pointer(uintptr(a1p.field_3) + 12)) &= 0xFFFFFFFE
	alloc.Free(a1p)
}

func (m *Struct264Module) Sub_487050(a1p *Struct88) {
	(*m.externs.Dword_587000_155144).field_0.Append_4258E0(&a1p.field_0)
}

func (m *Struct264Module) Sub_4870A0() {
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

func (m *Struct264Module) Sub_4870E0(a1 **Struct88) *Struct88 {
	result := (*m.externs.Dword_587000_155144).field_0.FirstSafe_425890()
	*a1 = result
	return result
}

func (m *Struct264Module) sub_487070(a1 *Struct88) {
	m.sub_487090(a1)
	m.sub_487030(a1)
	*memmap.PtrUint32(0x5D4594, 1193332) = 0
}

func (m *Struct264Module) sub_487090(a1 *Struct88) {
	a1.field_0.Remove_425920()
}

func (m *Struct264Module) Sub_487100(a1 **Struct88) *Struct88 {
	if *a1 != nil {
		res := (*a1).field_0.NextSafe_4258A0()
		*a1 = res
	}
	return *a1
}

func (m *Struct264Module) sub_4875B0(a1 **Struct264) *Struct264 {
	result := (*m.externs.Dword_587000_155144).field_3.FirstSafe_425890()
	*a1 = result
	return result
}

func (m *Struct264Module) sub_4875D0(a1 **Struct264) *Struct264 {
	if *a1 != nil {
		*a1 = (*a1).field_0.NextSafe_4258A0()
	}
	return *a1
}

func (m *Struct264Module) Sub_4875F0() int32 {
	var (
		v0     *Struct264
		v1     *Struct264
		result int32
		v3     *Struct264
	)
	(*m.externs.Dword_587000_155144).field_6 += 1
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
	result = int32((*m.externs.Dword_587000_155144).field_6 - 1)
	(*m.externs.Dword_587000_155144).field_6 = uint32(result)
	if result < 0 {
		(*m.externs.Dword_587000_155144).field_6 = 0
	}
	return result
}
