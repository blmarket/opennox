package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func (m *Phase6Module) Sub_4BDA80(a1p *Struct312) int32 {
	var (
		a1     int32 = int32(uintptr(unsafe.Pointer(a1p)))
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

func (m *Phase6Module) Sub_4873C0(a3p *Struct264) int32 {
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
						ccall.CallVoidPtr(*(*unsafe.Pointer)(unsafe.Add(v12p.field_43, 32)), unsafe.Pointer(v12p))
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

func (m *Phase6Module) Sub_4BDB40(a2p *Struct312) int32 {
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
