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

func (m *Struct88Module) sub_487030(a1p *Struct88) {
	// var lpMem unsafe.Pointer = unsafe.Pointer(a1p)
	ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1p.field_3) + 24)), unsafe.Pointer(a1p))
	*(*uint32)(unsafe.Pointer(uintptr(a1p.field_3) + 12)) &= 0xFFFFFFFE
	alloc.Free(a1p)
}

func (m *Struct88Module) Sub_487050(a1p *Struct88) {
	(*m.externs.Dword_587000_155144).field_0.Append_4258E0(&a1p.field_0)
}

func (m *Struct88Module) Sub_4870A0() {
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

func (m *Struct88Module) Sub_4870E0(a1 **Struct88) *Struct88 {
	result := (*m.externs.Dword_587000_155144).field_0.FirstSafe_425890()
	*a1 = result
	return result
}

func (m *Struct88Module) sub_487070(a1 *Struct88) {
	m.sub_487090(a1)
	m.sub_487030(a1)
	*memmap.PtrUint32(0x5D4594, 1193332) = 0
}

func (m *Struct88Module) sub_487090(a1 *Struct88) {
	a1.field_0.Remove_425920()
}

func (m *Struct88Module) Sub_487100(a1 **Struct88) *Struct88 {
	if *a1 != nil {
		res := (*a1).field_0.NextSafe_4258A0()
		*a1 = res
	}
	return *a1
}

func (m *Struct88Module) sub_4875B0(a1 **Struct264) *Struct264 {
	result := (*m.externs.Dword_587000_155144).field_3.FirstSafe_425890()
	*a1 = result
	return result
}

func (m *Struct88Module) sub_4875D0(a1 **Struct264) *Struct264 {
	if *a1 != nil {
		*a1 = (*a1).field_0.NextSafe_4258A0()
	}
	return *a1
}

func (m *Struct88Module) Sub_4875F0() int32 {
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
			m.sub_487680(v0)
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
