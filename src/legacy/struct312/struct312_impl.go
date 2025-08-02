package struct312

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func (m *Struct312Module) Sub_4BD720(a1 int32) *uint32 {
	var v1 *uint32
	v1p, _ := alloc.Calloc(1, 0x138)
	v1 = (*uint32)(v1p)
	alloc.Memset(unsafe.Pointer(v1), 0, 0x138)
	m.sub_425770(unsafe.Pointer(v1))
	m.sub_4BDC00(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*30))))))
	m.sub_4864A0(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*44))))
	m.sub_4BD7C0(v1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*33)) = uint32(a1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*43)) = *(*uint32)(unsafe.Pointer(uintptr(a1 + 256)))
	if ccall.CallIntPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 256))) + 4))), unsafe.Pointer(v1)) == 0 {
		return v1
	}
	if v1 != nil {
		m.Sub_4BD7A0(unsafe.Pointer(v1))
	}
	return nil
}
func (m *Struct312Module) Sub_4BD7A0(lpMem unsafe.Pointer) {
	ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(lpMem)), 4*43))) + 8))), lpMem)
	alloc.FreePtr(lpMem)
}
func (m *Struct312Module) sub_4BD7C0(a1 *uint32) *uint32 {
	var result *uint32
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*69)) = m.sub_4BD8C0_ptr
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*70)) = m.sub_4BD940_ptr
	*(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a1), 4*71)) = m.sub_4BD9B0_ptr
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*34)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*35)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*36)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*38)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*3)) = 1
	m.sub_4BDC00(int32(uintptr(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*30))))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*30)) = 0
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*29)) = *memmap.PtrUint32(0x5D4594, 1193340)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*28)) = 0
	result = (*uint32)(m.sub_4864A0(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*4)))))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a1), 4*72)) = 0
	return result
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
func (m *Struct312Module) Sub_4BD940(a1 int32) int32 {
	if *(*uint32)(unsafe.Pointer(uintptr(a1 + 128))) != 0 {
		if *(*int32)(unsafe.Pointer(uintptr(a1 + 128))) != -1 {
			*(*uint32)(unsafe.Pointer(uintptr(a1 + 128)))--
		}
		m.Sub_4BDB90((*uint32)(unsafe.Pointer(uintptr(a1))), *(**uint32)(unsafe.Pointer(uintptr(a1 + 288))))
	} else {
		m.Sub_4BDB90((*uint32)(unsafe.Pointer(uintptr(a1))), nil)
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
func (m *Struct312Module) Sub_4BD9B0(a2 *uint32) int32 {
	var (
		v1     int32
		result int32
	)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*72)) = 0
	v1 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*31)))
	*((*uint8)(unsafe.Pointer(&v1))) = uint8(int8(v1 & 0xFA))
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*31)) = uint32(v1)
	*(*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*32)) = 0
	m.sub_4864A0(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(a2), 4*4))))
	v2 := *(*unsafe.Pointer)(unsafe.Add(unsafe.Pointer(a2), 4*36))
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
func (m *Struct312Module) Sub_4BDB90(a1 *uint32, a2 *uint32) {
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