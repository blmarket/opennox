package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
)

func (m *Phase6Module) Sub_4BDA80(a1_ *Struct312) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		result int32 = 0
	)
	if int32(*(*uint8)(unsafe.Pointer(uintptr(a1 + 124))))&5 != 0 {
		ccall.CallVoidInt(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 172))) + 16))), int(a1))
	}
	result2 := *(*uint32)(unsafe.Pointer(uintptr(a1 + 148)))
	if result2 != 0 {
		result = int32(ccall.CallIntInt(unsafe.Pointer(uintptr(result2)), int(a1)))
	}
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 288))) = 0
	return result
}

func (m *Phase6Module) sub_486E90(a1_ *Struct312) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1_)))
		v1     int32
		result int32
	)
	v1 = int32(*(*uint32)(unsafe.Pointer(uintptr(a1 + 132))))
	m.nox_common_list_remove_425920(unsafe.Pointer(uintptr(a1)))
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 192)))--
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 212)))++
	m.nox_common_list_remove_425920(unsafe.Pointer(uintptr(a1)))
	result = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 212))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(v1 + 212))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 212))) = 0
	}
	return result
}

func (m *Phase6Module) Sub_4BDA60(lpMem_ *Struct312) {
	var lpMem unsafe.Pointer = unsafe.Pointer(lpMem_)
	m.Sub_4BDA80((*Struct312)(unsafe.Pointer(uintptr(int32(uintptr(lpMem))))))
	m.sub_486E90((*Struct312)(unsafe.Pointer(uintptr(int32(uintptr(lpMem))))))
	m.sub_4BD7A0(lpMem)
}

func (m *Phase6Module) Sub_4873C0(a3 int32) int32 {
	var (
		v1  int32
		v3  int64
		v4  uint32
		v5  int32
		v6  bool
		v7  uint32
		v8  uint32
		v9  int32
		v10 uint32
		v11 *uint32
		v12 int32
		v13 int32
		v14 *uint32
		v15 int32
		v16 int32
		v17 int32
	)
	v1 = a3
	if *(*uint32)(unsafe.Pointer(uintptr(a3 + 212))) != 0 {
		return -2146304000
	}
	v3 = int64(m.nox_platform_get_ticks())
	v4 = *(*uint32)(unsafe.Pointer(uintptr(a3 + 248)))
	v5 = int32(v3)
	v6 = uint32(int32(v3)) < v4
	v7 = uint32(int32(v3 - int64(v4)))
	v8 = *(*uint32)(unsafe.Pointer(uintptr(a3 + 224)))
	v16 = int32(*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), 4*1)))
	v9 = int32((*(*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(&v3))), 4*1))) - (uint32(bool2int32(v6)) + *(*uint32)(unsafe.Pointer(uintptr(a3 + 252)))))
	v10 = *(*uint32)(unsafe.Pointer(uintptr(a3 + 228)))
	if (((uint64(v9)) << 32) | uint64(v7)) >= (((uint64(v10)) << 32) | uint64(v8)) {
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 232))) = v7
		*(*uint32)(unsafe.Pointer(uintptr(a3 + 236))) = uint32(v9)
		if *(*uint64)(unsafe.Pointer(uintptr(a3 + 240))) > (((uint64(v10))<<32)|uint64(v8))*10 {
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 240))) = 0
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 244))) = 0
		}
		if (((uint64(v9)) << 32) | uint64(v7)) > *(*uint64)(unsafe.Pointer(uintptr(a3 + 240))) {
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 240))) = v7
			*(*uint32)(unsafe.Pointer(uintptr(a3 + 244))) = uint32(v9)
		}
		v11 = (*uint32)(unsafe.Pointer(uintptr(a3 + 88)))
		v15 = a3 + 88
		m.sub_486520(unsafe.Pointer(uintptr(a3 + 88)))
		if *(*uint32)(unsafe.Pointer(uintptr(a3 + 184))) != 0 {
			m.sub_486520(unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(a3 + 184)))))
		}
		if *(*uint32)(unsafe.Pointer(uintptr(a3 + 184))) == 0 || (func() int32 {
			v17 = int32(m.sub_486550(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(a3 + 184))))))
			return v17
		}()) == 0 {
			v17 = int32(m.sub_486550(unsafe.Pointer(uintptr(v1 + 88))))
		}
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 248))) = uint32(v5)
		*(*uint32)(unsafe.Pointer(uintptr(v1 + 252))) = uint32(v16)
		v12 = int32(*(*uint32)(unsafe.Pointer(uintptr(v1 + 200))))
		if v12 != v1+200 {
			for {
				v13 = int32(*(*uint32)(unsafe.Pointer(uintptr(v12))))
				if int32(*(*uint8)(unsafe.Pointer(uintptr(v12 + 124))))&1 != 0 && *(*uint32)(unsafe.Pointer(uintptr(v12 + 288))) != 0 {
					if (func() int32 {
						m.sub_486520(unsafe.Pointer(uintptr(v12 + 16)))
						return v17
					}()) != 0 || m.sub_486550(unsafe.Pointer(uintptr(v12+16))) != 0 || *(*uint32)(unsafe.Pointer(uintptr(v12 + 116))) != 0 && m.sub_486550(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(v12 + 116))))) != 0 || *(*uint32)(unsafe.Pointer(uintptr(v12 + 112))) != 0 && m.sub_486550(unsafe.Pointer(*(**uint8)(unsafe.Pointer(uintptr(v12 + 112))))) != 0 {
						m.sub_4BD840((*Struct312)(unsafe.Pointer(uintptr(v12))))
						ccall.CallVoidInt(*(*unsafe.Pointer)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v12 + 172))) + 32))), int(v12))
					}
				}
				v12 = v13
				if v13 == v1+200 {
					break
				}
			}
			v11 = (*uint32)(unsafe.Pointer(uintptr(v15)))
		}
		v14 = *(**uint32)(unsafe.Pointer(uintptr(v1 + 184)))
		if v14 != nil {
			m.sub_486620(unsafe.Pointer(v14))
		}
		m.sub_486620(unsafe.Pointer(v11))
	}
	return 0
}

func (m *Phase6Module) sub_4BD840(a3p *Struct312) {
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

func (m *Phase6Module) Sub_4BDB30(a1 *Struct312) {
	a1.field_31 &= 0xFFFFFFEF
}

func (m *Phase6Module) Sub_4BDB40(a2p *Struct312) int32 {
	var result int32
	if int32(*(*uint8)(unsafe.Pointer(&a2p.field_31)))&5 != 0 {
		return -2146500608
	}
	if a2p.field_72 == 0 {
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
