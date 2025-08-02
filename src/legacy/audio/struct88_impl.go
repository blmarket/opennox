package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func (m *Struct88Module) Sub_486FE0(a1 unsafe.Pointer) *Struct88 {
	v1p, _ := alloc.Calloc(1, 0x58)
	v1pp := (*Struct88)(v1p)
	alloc.Memset(v1p, 0, 0x58)
	m.sub_425770(&v1pp.field_0)
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

func (m *Struct88Module) sub_487030(a1p *Struct88) {
	var lpMem unsafe.Pointer = unsafe.Pointer(a1p)
	ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), 4*3))) + 24))), lpMem)
	*(*uint32)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), 4*3))) + 12))) &= 0xFFFFFFFE
	alloc.Free(a1p)
}

func (m *Struct88Module) Sub_487050(a1_ *Struct88) {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	m.nox_common_list_append_4258E0((unsafe.Pointer(uintptr(*(*int32)(unsafe.Pointer(&*m.dword_587000_155144))))), (unsafe.Pointer(a1)))
}

func (m *Struct88Module) Sub_4870A0() {
	var (
		v1 unsafe.Pointer
		v2 unsafe.Pointer
		v3 unsafe.Pointer
	)
	v1 = m.Sub_4870E0(&v3)
	if v1 != nil {
		for {
			v2 = m.Sub_487100(&v3)
			m.sub_487070(unsafe.Pointer(v1))
			v1 = v2
			if v2 == nil {
				break
			}
		}
	}
}
func (m *Struct88Module) Sub_4870E0(a1 *unsafe.Pointer) unsafe.Pointer {
	result := m.nox_common_list_getFirstSafe_425890((unsafe.Pointer(*(**int32)(unsafe.Pointer(&*m.dword_587000_155144)))))
	*a1 = result
	return result
}
func (m *Struct88Module) sub_487070(lpMem unsafe.Pointer) {
	m.sub_487090((**uint32)(lpMem))
	m.sub_487030((*Struct88)(lpMem))
	*memmap.PtrUint32(0x5D4594, 1193332) = 0
}
func (m *Struct88Module) sub_487090(a1 **uint32) {
	m.nox_common_list_remove_425920(unsafe.Pointer(a1))
}
func (m *Struct88Module) Sub_487100(a1 *unsafe.Pointer) unsafe.Pointer {
	if *a1 != nil {
		*a1 = (unsafe.Pointer(m.nox_common_list_getNextSafe_4258A0((unsafe.Pointer(*a1)))))
	}
	return *a1
}
func (m *Struct88Module) sub_4875B0(a1 *int32) *int32 {
	var result *int32
	result = (*int32)(unsafe.Pointer(m.nox_common_list_getFirstSafe_425890((unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 12))))))
	*a1 = int32(uintptr(unsafe.Pointer(result)))
	return result
}
func (m *Struct88Module) sub_4875D0(a1 **int32) *int32 {
	if *a1 != nil {
		*a1 = (*int32)(unsafe.Pointer(m.nox_common_list_getNextSafe_4258A0((unsafe.Pointer(*a1)))))
	}
	return *a1
}
func (m *Struct88Module) Sub_4875F0() int32 {
	var (
		v0     *int32
		v1     *int32
		result int32
		v3     *int32
	)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24)))++
	v0 = m.sub_4875B0((*int32)(unsafe.Pointer(&v3)))
	if v0 != nil {
		for {
			v1 = m.sub_4875D0(&v3)
			m.sub_487680((*Struct264)(unsafe.Pointer(v0)))
			v0 = v1
			if v1 == nil {
				break
			}
		}
	}
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(uint32(uintptr(*m.dword_587000_155144)) + 24))) = 0
	}
	return result
}
