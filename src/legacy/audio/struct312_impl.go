package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func (m *Struct312Module) Sub_4BD720(a1p *Struct264) *Struct312 {
	var v1 *uint32
	v1pp, _ := alloc.Calloc(1, 0x138)
	v1p := (*Struct312)(v1pp)
	alloc.Memset(unsafe.Pointer(v1p), 0, 0x138)
	m.sub_425770(&v1p.ListItem)
	m.sub_4BDC00(int32(uintptr(unsafe.Pointer(&v1p.field_30))))
	v1p.timerGroup_44.Init()
	m.sub_4BD7C0(v1p)
	v1p.field_33 = a1p
	v1p.field_43 = a1p.field_64
	if ccall.CallIntPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1p.field_64)) + 4))), unsafe.Pointer(v1p)) == 0 {
		return v1p
	}
	if v1 != nil {
		m.Sub_4BD7A0(v1p)
	}
	return nil
}

func (m *Struct312Module) Sub_4BD7A0(lpMem *Struct312) {
	ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((lpMem)), 4*43))) + 8))), unsafe.Pointer(lpMem))
	alloc.FreePtr(unsafe.Pointer(lpMem))
}

func (m *Struct312Module) sub_4BD7C0(a1p *Struct312) {
	a1p.field_69 = m.externs.Sub_4BD8C0_ptr
	a1p.field_70 = m.externs.Sub_4BD940_ptr
	a1p.field_71 = m.externs.Sub_4BD9B0_ptr
	a1p.field_34 = 0
	a1p.field_35 = nil
	a1p.field_36 = nil
	a1p.field_38 = nil
	a1p.field_3 = 1
	m.sub_4BDC00(int32(uintptr(unsafe.Pointer(&a1p.field_30))))
	a1p.field_30 = 0
	a1p.field_29 = *m.externs.Ptr_uint32_5d4594_1193340
	a1p.field_28 = nil
	a1p.timerGroup_4.Init()
	a1p.field_72 = 0
}

func (m *Struct312Module) Sub_4BD8C0(a1 int32) int32 {
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

func (m *Struct312Module) Sub_4BD940(a1p *Struct312) int32 {
	var a1 int32 = int32(uintptr(unsafe.Pointer(a1p)))
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) != 0 {
		if *(*int32)(unsafe.Pointer(uintptr(a1 + 128))) != -1 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))--
		}
		m.Sub_4BDB90(a1p, *(**uint32)(unsafe.Pointer(uintptr(a1 + 288))))
	} else {
		m.Sub_4BDB90(a1p, nil)
	}
	v1 := *(*unsafe.Pointer)(unsafe.Pointer(uintptr(a1 + 140)))
	if v1 != nil {
		ccall.CallVoidPtr(v1, unsafe.Pointer(uintptr(a1)))
	}
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 288))) != 0 {
		ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 172))) + 36))), unsafe.Pointer(uintptr(a1)))
	}
	return 0
}

func (m *Struct312Module) Sub_4BD9B0(a2p *Struct312) int32 {
	var a2 *uint32 = (*uint32)(unsafe.Pointer(a2p))
	var (
		result int32
	)
	a2p.field_72 = 0
	a2p.field_31 &= 0xFA
	a2p.field_32 = 0
	a2p.timerGroup_4.Init()
	v2 := a2p.field_36
	if v2 != nil {
		result = int32(ccall.CallIntPtr(v2, unsafe.Pointer(a2)))
	} else {
		result = 0
	}
	return result
}

func (m *Struct312Module) sub_4BDC00(a1 int32) int32 {
	var result int32
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 8))) = 0
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 4))) = 0
	return result
}

func (m *Struct312Module) Sub_4BDB90(a1p *Struct312, a2 *uint32) {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1p))
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

func (m *Struct312Module) sub_487C80(a1 int32) int32 {
	return int32(uintptr(unsafe.Pointer(m.nox_common_list_getNext_425940((unsafe.Pointer(uintptr(a1 + 8)))))))
}
