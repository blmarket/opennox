package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func (m *Struct264Module) sub_4871C0(a1 int32, a2 int32, a3 unsafe.Pointer) *Struct264 {
	var (
		v3 int32
		v4 *uint32
	)
	v3 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))))
	v4p, _ := alloc.Calloc(1, 0x108)
	v4m := (*Struct264)(v4p)
	v4 = (*uint32)(v4p)
	alloc.Memset(unsafe.Pointer(v4m), 0, 0x108)
	m.sub_425770(unsafe.Pointer(&v4m.field_0))
	v4m.field_6 = uint32(a2)
	v4m.field_5 = unsafe.Pointer(uintptr(a1))
	v4m.field_4 = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 16)))++
	*(*uint32)(unsafe.Pointer(uintptr(a1 + a2*4 + 24))) = uint32(uintptr(unsafe.Pointer(v4m)))
	v4m.field_64 = *(*uint32)(unsafe.Pointer(uintptr(v3 + 36)))
	m.nox_common_list_clear_425760(unsafe.Pointer(&v4m.field_50))
	m.sub_4864A0(unsafe.Pointer(&v4m.TimerGroup_22))
	v4m.field_53 = 0
	v4m.field_56 = 33
	v4m.field_60 = 0
	v4m.field_58 = 0
	v4m.field_62 = 0
	v4m.field_54 = m.sub_4873C0_ptr
	v4m.field_57 = 0
	v4m.field_61 = 0
	v4m.field_59 = 0
	v4m.field_63 = 0
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

func (m *Struct264Module) Sub_487150(a1 int32, a2 unsafe.Pointer) *Struct264 {
	var (
		v2 int32
		v6 int32
	)
	v2 = a1
	if a1 == -1 {
		v2 = 0
	}
	m.sub_487360(v2, (**int32)(unsafe.Pointer(&a1)), &v6)
	if a1 == 0 {
		return nil
	}
	v3m := *(**Struct264)(unsafe.Pointer(uintptr(a1 + v6*4 + 24)))
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
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24)))++
	m.nox_common_list_append_4258E0((unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 12))), (unsafe.Pointer(a1_)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) = 0
	}
	return result
}

func (m *Struct264Module) Sub_431270() {
	if *m.dword_5d4594_805984 != nil {
		m.Sub_487680((*Struct264)(*m.dword_5d4594_805984))
		*m.dword_5d4594_805984 = nil
	}
}
func (m *Struct264Module) Sub_487680(lpMem_ *Struct264) {
	var lpMem unsafe.Pointer = unsafe.Pointer(lpMem_)
	m.sub_4876A0((*Struct264)(unsafe.Pointer((**uint32)(lpMem))))
	m.sub_4872C0((*Struct264)(lpMem))
}
func (m *Struct264Module) Sub_431290() {
	if *m.dword_5d4594_805984 != nil {
		m.sub_487970((*Struct264)(*m.dword_5d4594_805984), -1)
	}
}

func (m *Struct264Module) sub_487970(a1_ *Struct264, a2 int32) *int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		result *int32
		v3     *int32
		v4     int32
		v5     *int32
	)
	result = m.sub_4877D0((*Struct264)(unsafe.Pointer(uintptr(a1))), &a1)
	v3 = result
	if result != nil {
		v4 = a2
		for {
			result = m.sub_4877F0((**int32)(unsafe.Pointer(&a1)))
			v5 = result
			if v4 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*3)) == v4 {
				result = (*int32)(unsafe.Pointer(uintptr(m.sub_4BDA80(int(uintptr(unsafe.Pointer(v3)))))))
			}
			v3 = v5
			if v5 == nil {
				break
			}
		}
	}
	return result
}

func (m *Struct264Module) sub_487590(a1_ *Struct264, a2 unsafe.Pointer) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		result int32
	)
	result = a1
	alloc.Memcpy(unsafe.Pointer(uintptr(a1+60)), a2, 0x1C)
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
func (m *Struct264Module) sub_4876A0(a1_ *Struct264) unsafe.Pointer {
	var (
		a1     **uint32 = (**uint32)(unsafe.Pointer(a1_))
		result unsafe.Pointer
	)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24)))++
	m.nox_common_list_remove_425920(unsafe.Pointer(a1))
	result = unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) - 1))
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) = uint32(uintptr(result))
	if int32(uintptr(result)) < 0 {
		result = *m.dword_587000_155144
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) = 0
	}
	return result
}
func (m *Struct264Module) sub_487910(a1_ *Struct264, a2 int32) int32 {
	var (
		a1 int32 = int32(uintptr(unsafe.Pointer(a1_)))
		v2 *int32
		v3 int32
		v4 *int32
	)
	v2 = m.sub_4877D0((*Struct264)(unsafe.Pointer(uintptr(a1))), &a1)
	if v2 == nil {
		return 0
	}
	v3 = a2
	for {
		v4 = m.sub_4877F0((**int32)(unsafe.Pointer(&a1)))
		if v3 == -1 || *(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*3)) == v3 {
			m.sub_4BDA60(unsafe.Pointer(v2))
		}
		v2 = v4
		if v4 == nil {
			break
		}
	}
	return 0
}
func (m *Struct264Module) Sub_452810(a1 int32, a2 int8) *int32 {
	var (
		v2 *int32
		v3 *int32
	)
	v2 = nil
	if *m.dword_5d4594_1045428 != 0 {
		v3 = m.sub_487810((*Struct264)(unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&*m.dword_5d4594_1045428))))), 1)
		v2 = v3
		if v3 != nil {
			if *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*31))&0x15 != 0 && *(*int32)(unsafe.Add(unsafe.Pointer(v3), 4*30)) > a1 {
				return nil
			}
			m.sub_4BDA80(int(uintptr(unsafe.Pointer(v3))))
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*29)) = int32(uintptr(*m.dword_587000_127004))
			*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*30)) = a1
			if int32(a2)&1 != 0 {
				*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*32)) = -1
			} else {
				*(*int32)(unsafe.Add(unsafe.Pointer(v2), 4*32)) = 0
			}
			m.sub_486320(unsafe.Pointer((*int32)(unsafe.Add(unsafe.Pointer(v2), 4*4))), 0x4000)
		}
	}
	return v2
}
func (m *Struct264Module) sub_487810(a1p *Struct264, a2 int32) *int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
		v2     uint32
		v3     int32
		v4     *int32
		result *int32
		v6     int32
		v7     uint32
		v8     int32
		v9     *int32
	)
	v2 = 4294967295
	if a2 == -1 {
		a2 = 1
	}
	v3 = 127
	v4 = nil
	v8 = 127
	v9 = nil
	for result = m.sub_4877D0((*Struct264)(unsafe.Pointer(uintptr(a1))), &a1); result != nil; result = m.sub_4877F0((**int32)(unsafe.Pointer(&a1))) {
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

func (m *Struct264Module) sub_4877D0(a1p *Struct264, a2 *int32) *int32 {
	var (
		result *int32
	)
	result = (*int32)(unsafe.Pointer(m.nox_common_list_getFirstSafe_425890(unsafe.Pointer(&a1p.field_50))))
	*a2 = int32(uintptr(unsafe.Pointer(result)))
	return result
}

func (m *Struct264Module) sub_4877F0(a1 **int32) *int32 {
	if *a1 != nil {
		*a1 = (*int32)(unsafe.Pointer(m.nox_common_list_getNextSafe_4258A0((unsafe.Pointer(*a1)))))
	}
	return *a1
}

func (m *Struct264Module) sub_487360(a1 int32, a2 **int32, a3 *int32) *int32 {
	var (
		result *int32
		i      int32
		v5     int32
		v6     *int32
	)
	result = (*int32)(m.sub_4870E0((unsafe.Pointer(&v6))))
	for i = a1; result != nil; result = (*int32)(m.sub_487100(unsafe.Pointer(&v6))) {
		v5 = *(*int32)(unsafe.Add(unsafe.Pointer(result), 4*5))
		if i < v5 {
			break
		}
		i -= v5
	}
	*a2 = result
	if result != nil {
		result = a3
		*a3 = i
	} else {
		*a3 = -1
	}
	return result
}
