package main

import (
	"github.com/gotranspile/cxgo/runtime/libc"
	"unsafe"
)

func sub_451850(a2 int, a3p unsafe.Pointer) int {
	var (
		a3     int = int(uintptr(a3p))
		v2     int
		v3     *uint8
		result int
	)
	v2 = 0
	v3 = (*uint8)(mem_getPtr(0x5D4594, 840712))
	for {
		sub_451920((*struct200)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v3))), -int(unsafe.Sizeof(uint32(0))*21))))))
		*(*uint32)(unsafe.Pointer(v3)) = uint32(uintptr(unsafe.Pointer(nox_xxx_getSndName_40AF80(v2))))
		v3 = (*uint8)(unsafe.Add(unsafe.Pointer(v3), 200))
		v2++
		if int(uintptr(unsafe.Pointer(v3))) >= int(uintptr(mem_getPtr(0x5D4594, 1045312))) {
			break
		}
	}
	dword_5d4594_1045420 = uint32(a3)
	dword_5d4594_1045428 = uint32(a2)
	if a3 != 0 {
		dword_5d4594_1045424 = uint32(uintptr(unsafe.Pointer(sub_4BD340(a3, 0x100000, 200, 0x2000))))
		dword_5d4594_1045436 = uint32(uintptr(unsafe.Pointer(sub_4BD280(200, 576))))
	}
	if dword_5d4594_1045424 == 0 || dword_5d4594_1045420 == 0 || dword_5d4594_1045428 == 0 || dword_5d4594_1045436 == 0 {
		return 0
	}
	nox_common_list_clear_425760((*nox_list_item_t)(mem_getPtr(0x5D4594, 840612)))
	sub_4864A0(mem_getPtr(0x5D4594, 1045228))
	result = 1
	*(*uint32)(unsafe.Pointer(uintptr(dword_5d4594_1045428 + 184))) = uint32(uintptr(mem_getPtr(0x5D4594, 1045228)))
	dword_5d4594_1045432 = 1
	return result
}
func sub_451920(a2_ *struct200) int {
	var a2 *uint32 = &a2_.field_0
	_ = a2
	a2_.field_0 = 0
	a2_.field_1 = 0
	a2_.field_2 = 0
	a2_.field_14 = 0
	a2_.field_15 = 0
	a2_.field_19 = 0
	a2_.field_20 = 0
	a2_.field_12 = 1
	a2_.field_48 = 0
	a2_.field_18 = 0
	a2_.field_17 = 0
	a2_.field_25 = 0
	a2_.field_26 = 0
	a2_.field_16 = 600
	return sub_4862E0(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(&a2_.field_4))))), 0x4000)
}
func sub_452010() int {
	var (
		v0 *uint8
		v1 int
		v2 int
	)
	v0 = (*uint8)(mem_getPtr(0x5D4594, 839892))
	v1 = 6
	for {
		v2 = 10
		for {
			nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(v0)))
			v0 = (*uint8)(unsafe.Add(unsafe.Pointer(v0), 12))
			v2--
			if v2 == 0 {
				break
			}
		}
		v1--
		if v1 == 0 {
			break
		}
	}
	return int(func() uint32 {
		p_ := mem_getU32Ptr(0x5D4594, 1045444)
		*p_++
		return *p_
	}())
}
func sub_452190(a1_ *struct200) {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	nox_common_list_remove_425920(unsafe.Pointer((**uint32)(unsafe.Pointer(&a1_.field_28))))
}
func sub_4521A0(a1 int) *int {
	var (
		v1 int
		v2 *uint8
		v3 int
		v4 *int
		v5 *int
	)
	v1 = 0
	v2 = (*uint8)(mem_getPtr(0x5D4594, 839892))
	if a1 > 0 {
		for {
			v3 = 0
			v4 = (*int)(unsafe.Pointer(v2))
			for {
				v5 = (*int)(unsafe.Pointer(nox_common_list_getFirstSafe_425890((*nox_list_item_t)(unsafe.Pointer(v4)))))
				if v5 != nil {
					return (*int)(unsafe.Add(unsafe.Pointer(v5), -int(unsafe.Sizeof(int(0))*28)))
				}
				v3++
				v4 = (*int)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(int(0))*3))
				if v3 >= 10 {
					break
				}
			}
			v1++
			v2 = (*uint8)(unsafe.Add(unsafe.Pointer(v2), 120))
			if v1 < a1 {
				continue
			}
			break
		}
	}
	return nil
}
func sub_4521F0() int {
	var (
		result int
		v1     *uint8
		v2     *uint8
	)
	result = int(dword_5d4594_1045432)
	if dword_5d4594_1045432 != 0 {
		v1 = *(**uint8)(mem_getPtr(0x5D4594, 840612))
		if unsafe.Pointer(*(**uint8)(mem_getPtr(0x5D4594, 840612))) != mem_getPtr(0x5D4594, 840612) {
			for {
				v2 = *(**uint8)(unsafe.Pointer(v1))
				sub_4523D0((*struct576)(unsafe.Pointer(v1)))
				result = sub_451FE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v1)))))))
				v1 = v2
				if unsafe.Pointer(v2) == mem_getPtr(0x5D4594, 840612) {
					break
				}
			}
		}
	}
	return result
}
func sub_452230() *****int {
	var (
		result *****int
		v1     ****int
	)
	result = *(******int)(unsafe.Pointer(&dword_5d4594_1045432))
	if dword_5d4594_1045432 != 0 {
		result = *(******int)(mem_getPtr(0x5D4594, 840612))
		if unsafe.Pointer(*(**uint8)(mem_getPtr(0x5D4594, 840612))) != mem_getPtr(0x5D4594, 840612) {
			for {
				v1 = *result
				if int(uint8(uintptr(unsafe.Pointer(*(*****int)(unsafe.Add(unsafe.Pointer(result), unsafe.Sizeof((****int)(nil))*6))))))&1 != 0 {
					sub_451FE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(result)))))))
				}
				result = (*****int)(unsafe.Pointer(v1))
				if v1 == (****int)(mem_getPtr(0x5D4594, 840612)) {
					break
				}
			}
		}
	}
	return result
}
func nox_xxx_draw_452270(a1 int) *byte {
	var result *byte
	if dword_5d4594_1045432 != 0 && a1 >= 0 && a1 < 1023 {
		result = (*byte)(mem_getPtr(0x5D4594, uint32(a1*200+840628)))
	} else {
		result = nil
	}
	return result
}
func sub_4526F0(a1 int) int {
	var (
		v1 *uint32
		v2 int
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 152)))
	*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*6)) &= 0xFFFFFFFD
	v2 = 4
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) != 4 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*74)) != 0 || *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*142)) != 0 {
			v2 = 1
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) = 0
		}
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) != 0 {
			sub_452690((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v1)))))), int64(uint(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)))), v2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) = 0
			return 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*7)) = uint32(v2)
	}
	return 0
}
func sub_451CA0(a1_ *struct576) int {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	_ = a1
	var v1 int
	var v3 int
	var v4 *uint32
	v1 = int(a1_.field_42)
	a1_.field_108 = uint32(v1)
	if v1 == 0 {
		return 0
	}
	v3 = 0
	if v1 > 0 {
		v4 = &a1_.field_76[0]
		for {
			*v4 = uint32(func() int {
				p_ := &v3
				x := *p_
				*p_++
				return x
			}())
			v4 = (*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*1))
			if v3 >= int(a1_.field_108) {
				break
			}
		}
	}
	a1_.field_43 = 4294967295
	return sub_451CF0((*uint32)(unsafe.Pointer(a1_)))
}
func sub_451F30(a1_ *struct576, a2 int) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int
	var result int
	*(*uint32)(unsafe.Pointer(&a1_.field_10[a1_.field_42])) = uint32(uintptr(unsafe.Pointer(sub_4BD470(*(***uint32)(unsafe.Pointer(&dword_5d4594_1045424)), int(*(*int16)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1_.field_9)) + uint32(a2*2) + 128))))))))
	v2 = int(a1_.field_42)
	result = int(*(*uint32)(unsafe.Pointer(&a1_.field_10[v2])))
	if result != 0 {
		sub_4BD650(int(*(*uint32)(unsafe.Pointer(&a1_.field_10[v2]))))
		result = int(a1_.field_42 + 1)
		a1_.field_42 = uint32(result)
	}
	return result
}
func sub_451F90(a1_ *struct576) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v1 int
	var result int
	var v3 *int
	v1 = 0
	result = int(a1_.field_42)
	if result <= 0 {
		a1_.field_42 = 0
	} else {
		v3 = (*int)(unsafe.Pointer(&a1_.field_10[0]))
		for {
			sub_4BD660(*v3)
			*v3 = 0
			result = int(a1_.field_42)
			v1++
			v3 = (*int)(unsafe.Add(unsafe.Pointer(v3), unsafe.Sizeof(int(0))*1))
			if v1 >= result {
				break
			}
		}
		a1_.field_42 = 0
	}
	return result
}
func sub_451FE0(a1_ *struct576) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	nox_common_list_remove_425920(unsafe.Pointer(a1_))
	a1_.field_70 = 0
	return sub_4BD300(*(**uint32)(unsafe.Pointer(&dword_5d4594_1045436)), int(uintptr(unsafe.Pointer(a1_))))
}
func sub_452120(a1_ *struct576) *int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v1 int
	var result *int
	var v3 *int
	var v4 *uint8
	var v5 *uint8
	v1 = 0
	result = sub_4521A0(int(a1_.field_75 + *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1_.field_9)) + 48)))))
	v3 = result
	if result != nil {
		sub_452190((*struct200)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(result)))))))
		v4 = *(**uint8)(mem_getPtr(0x5D4594, 840612))
		if unsafe.Pointer(*(**uint8)(mem_getPtr(0x5D4594, 840612))) != mem_getPtr(0x5D4594, 840612) {
			for {
				v5 = *(**uint8)(unsafe.Pointer(v4))
				if *((**int)(unsafe.Add(unsafe.Pointer((**int)(unsafe.Pointer(v4))), unsafe.Sizeof((*int)(nil))*9))) == v3 {
					sub_4523D0((*struct576)(unsafe.Pointer(v4)))
					sub_451FE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v4)))))))
					v1 = 1
				}
				v4 = v5
				if unsafe.Pointer(v5) == mem_getPtr(0x5D4594, 840612) {
					break
				}
			}
		}
		result = (*int)(unsafe.Pointer(uintptr(v1)))
	}
	return result
}
func sub_4523D0(a1p_ *struct576) int {
	var (
		a1p    unsafe.Pointer = unsafe.Pointer(a1p_)
		a1     *uint32        = (*uint32)(a1p)
		result int            = 0
	)
	if (*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) & 1) == 0 {
		sub_452410((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(a1)))))))
		sub_451F90((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(a1)))))))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*7)) = 4
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*70)) = 0
		result = int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)))
		*((*uint8)(unsafe.Pointer(&result))) = uint8(int8(result | 1))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*6)) = uint32(result)
	}
	return result
}
func sub_452410(a1_ *struct576) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var result int
	result = int(a1_.field_44)
	if result != 0 && unsafe.Pointer(a1_) == unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(result + 152))))) {
		if int(a1_.field_6)&2 != 0 {
			sub_4BDA80(int(a1_.field_44))
		}
		sub_4BDB30(int(a1_.field_44))
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 152))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 148))) = 0
		result = int(a1_.field_44)
		*(*uint32)(unsafe.Pointer(uintptr(result + 140))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 144))) = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 112))) = 0
		a1_.field_44 = 0
	}
	return result
}
func sub_452490(a1_ *struct576) int {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	_ = a1
	var v1 int
	var v3 int
	var v4 int
	v1 = int(a1_.field_44)
	if unsafe.Pointer(a1_) != unsafe.Pointer(*(**uint32)(unsafe.Pointer(uintptr(v1 + 152)))) {
		return 0
	}
	v3 = int(a1_.field_74)
	sub_4BDB90((*uint32)(unsafe.Pointer(uintptr(v1))), (*uint32)(unsafe.Pointer(uintptr(a1_.field_74))))
	a1_.field_7 = 3
	v4 = int(a1_.field_6)
	*((*uint8)(unsafe.Pointer(&v4))) = uint8(int8(v4 | 2))
	a1_.field_6 = uint8(int8(v4))
	a1_.field_74 = 0
	if sub_4BDB40(int(a1_.field_44)) == 0 {
		return 1
	}
	a1_.field_7 = 1
	a1_.field_74 = uint32(v3)
	a1_.field_6 &= uint8(0xFFFFFFFD)
	return 0
}
func sub_452510(a3_ *struct576) {
	var a3 int = int(uintptr(unsafe.Pointer(a3_)))
	_ = a3
	var v1 int
	var v2 int
	if dword_587000_126996 == 0 {
		a3_.field_7 = 4
	}
	for {
		v1 = int(a3_.field_7)
		if v1 == 0 {
			break
		}
		v2 = v1 - 2
		if v2 != 0 {
			var v3 uint32 = uint32(v2 - 2)
			if v3 == 0 {
				sub_4523D0(a3_)
			}
			return
		}
		if uint64(nox_platform_get_ticks()) <= a3_.field_72 {
			return
		}
		a3_.field_7 = a3_.field_8
	}
	if sub_452580(a3_) == 0 {
		sub_4523D0(a3_)
	}
}
func sub_452690(a3_ *struct576, a4 int64, a5 int) int64 {
	var a3 int = int(uintptr(unsafe.Pointer(a3_)))
	_ = a3
	var result int64
	a3_.field_8 = uint32(a5)
	result = a4 + int64(nox_platform_get_ticks())
	a3_.field_72 = uint64(result)
	a3_.field_7 = 2
	return result
}
func nox_xxx_draw_452300(a1_ *struct200) *uint32 {
	if dword_5d4594_1045432 == 0 {
		return nil
	}
	if dword_587000_126996 == 0 {
		return nil
	}
	if a1_.field_0 == 0 {
		return nil
	}
	var v1p *struct576 = (*struct576)(unsafe.Pointer(sub_4BD2E0(*(***uint32)(unsafe.Pointer(&dword_5d4594_1045436)))))
	var v1 *uint32 = (*uint32)(unsafe.Pointer(v1p))
	if v1 == nil {
		sub_452230()
		v1 = sub_4BD2E0(*(***uint32)(unsafe.Pointer(&dword_5d4594_1045436)))
		if v1 == nil {
			return nil
		}
	}
	libc.MemSet(unsafe.Pointer(v1p), 0, 0x240)
	v1p.field_9 = a1_
	sub_425770(unsafe.Pointer(v1p))
	v1p.field_7 = 0
	v1p.field_75 = 0
	v1p.field_142 = 0
	v1p.field_108 = 0
	v1p.field_42 = 0
	sub_4864A0(unsafe.Pointer(&v1p.timerGroup_46))
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int(uintptr(mem_getPtr(0x5D4594, 840612)))))), (*nox_list_item_t)(unsafe.Pointer(v1p)))
	v1p.field_70 = func() uint32 {
		p_ := mem_getU32Ptr(0x587000, 127000)
		x := *p_
		*p_++
		return x
	}()
	return (*uint32)(unsafe.Pointer(v1p))
}
func sub_452E90(a1 *uint32, a2_ *struct576) int {
	var a2 int = int(uintptr(unsafe.Pointer(a2_)))
	_ = a2
	var result int
	result = int(uintptr(unsafe.Pointer(a2_)))
	*a1 = uint32(uintptr(unsafe.Pointer(a2_)))
	if a2_ != nil {
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*1)) = a2_.field_70
		result = int(*(*uint32)(unsafe.Pointer(&a2_.field_9)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*2)) = uint32(result)
	}
	return result
}
func sub_452EE0(a1_ *struct576, a2 int) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int
	v2 = int(sub_452F10(a1_, a2))
	sub_486320(unsafe.Pointer((*uint32)(unsafe.Pointer(&a1_.timerGroup_46))), v2)
	return sub_4863B0(unsafe.Pointer((*uint)(unsafe.Pointer(&a1_.timerGroup_46))))
}
func sub_452F10(a1_ *struct576, a2 int) uint {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int
	v2 = a2
	if a2 <= 100 {
		if a2 < 0 {
			v2 = 0
		}
	} else {
		v2 = 100
	}
	return uint(v2*163*int(*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(&a1_.field_9)) + 20)))>>16)) >> 14
}
func sub_452F50(a1_ *struct576, a2 int) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int
	v2 = int(sub_452F10(a1_, a2))
	return sub_486350(unsafe.Pointer(&a1_.timerGroup_46), v2)
}
func sub_452F80(a1_ *struct576, a2 int) *uint32 {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int
	v2 = sub_452FA0(a2)
	return (*uint32)(sub_486320(unsafe.Pointer(&a1_.timerGroup_46.field_16.field_0), v2))
}
func sub_451CF0(a1 *uint32) int {
	var (
		v1     int
		result int
		v3     int
		v4     int
		v5     int
		v6     int
		v7     int
		v8     *uint32
		v9     int
	)
	v1 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*9)))
	result = int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)))
	v3 = int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
	if result != 0 {
		if v3&2 != 0 {
			v5 = nox_common_randomIntMinMax_415FF0(0, result-1, libc.CString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c"), 376)
			v6 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108)) - 1)
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*43)) = *(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v5+76)))
			v7 = v5
			if v5 < v6 {
				v8 = (*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v5+76)))
				for {
					v7++
					*v8 = *(*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*1))
					v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v8), unsafe.Sizeof(uint32(0))*1))
					if v7 >= int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108))-1) {
						break
					}
				}
			}
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*43))++
		}
		v9 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*43)))
		*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*108))--
		result = sub_4BD710(int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*uintptr(v9+10)))))
	} else if v3&1 != 0 {
		if *(*uint32)(unsafe.Pointer(uintptr(v1 + 60))) != 0 && (func() bool {
			v4 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*109)) + 1)
			*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*109)) = uint32(v4)
			return v4 >= *(*int)(unsafe.Pointer(uintptr(v1 + 60)))
		}()) {
			result = 0
		} else {
			result = sub_451CA0((*struct576)(unsafe.Pointer(a1)))
		}
	}
	return result
}
func sub_451DC0(a1 int) int {
	var (
		v1     *uint32
		result int
		v3     int
		i      int
		v5     int
		v6     int
	)
	v1 = *(**uint32)(unsafe.Pointer(uintptr(a1 + 36)))
	result = int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 168))))
	v3 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*1)))
	if result != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*17)) < 0x21 {
			return result
		}
		sub_451F90((*struct576)(unsafe.Pointer(uintptr(a1))))
	}
	if v3&4 != 0 {
		if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*17)) >= 0x21 {
			v5 = sub_451E80(a1)
			result = sub_451F30((*struct576)(unsafe.Pointer(uintptr(a1))), v5)
		} else {
			result = int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*48)))
			for i = 0; i < result; i++ {
				sub_451F30((*struct576)(unsafe.Pointer(uintptr(a1))), i)
				result = int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*48)))
			}
		}
	} else if v3&2 != 0 {
		v6 = nox_common_randomIntMinMax_415FF0(0, int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*48))-1), libc.CString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c"), 536)
		result = sub_451F30((*struct576)(unsafe.Pointer(uintptr(a1))), v6)
	} else {
		result = sub_451F30((*struct576)(unsafe.Pointer(uintptr(a1))), 0)
	}
	return result
}
func sub_452580(a1_ *struct576) int {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	_ = a1
	var v1 int
	var result int
	var v3 int
	var v4 int
	var v5 int
	v1 = int(uintptr(unsafe.Pointer(a1_.field_9)))
	if *(*uint32)(unsafe.Pointer(uintptr(v1 + 192))) == 0 {
		return 0
	}
	v3 = int(a1_.field_75)
	a1_.field_109 = 0
	result = int(uintptr(unsafe.Pointer(sub_452810(int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 48)))+uint32(v3)), 0))))
	a1_.field_44 = uint32(result)
	if result != 0 {
		v4 = nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 76)))), int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 80)))), libc.CString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c"), 1482)
		sub_486320(unsafe.Pointer(uintptr(a1_.field_44+48)), v4+100)
		sub_4BDB20(int(a1_.field_44))
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 152))) = uint32(uintptr(unsafe.Pointer(a1_)))
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 140))) = uint32(libc.FuncAddr(sub_452770))
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 144))) = uint32(libc.FuncAddr(sub_4526F0))
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 148))) = uint32(libc.FuncAddr(sub_4526D0))
		a1_.field_7 = 1
		*(*uint32)(unsafe.Pointer(uintptr(a1_.field_44 + 112))) = uint32(uintptr(unsafe.Pointer(&a1_.timerGroup_46)))
		if int(*(*uint8)(unsafe.Pointer(uintptr(v1 + 4))))&8 != 0 {
			v5 = nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 68)))), int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 72)))), libc.CString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c"), 1497)
			if v5 > 33 {
				sub_452690(a1_, int64(v5), 1)
			}
		}
		result = 1
	}
	return result
}
func sub_451E80(a1 int) int {
	var (
		v1  int
		v2  int
		v3  int
		v4  int
		v5  int
		v6  int
		v7  int
		v8  int
		v9  int
		v10 int
		v11 *uint32
	)
	v1 = int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 36))))
	v2 = int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 4))))
	if *(*int)(unsafe.Pointer(uintptr(a1 + 568))) <= 0 {
		v3 = int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 192))))
		v4 = 0
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))) = uint32(v3)
		if v3 > 0 {
			v5 = a1 + 440
			for {
				v5 += 4
				v6 = v3 - func() int {
					p_ := &v4
					x := *p_
					*p_++
					return x
				}() - 1
				*(*uint32)(unsafe.Pointer(uintptr(v5 - 4))) = uint32(v6)
				v3 = int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))))
				if v4 >= v3 {
					break
				}
			}
		}
	}
	v7 = int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 568))) = uint32(v7)
	if (v2 & 2) == 0 {
		return int(*(*uint32)(unsafe.Pointer(uintptr(a1 + v7*4 + 440))))
	}
	v8 = nox_common_randomIntMinMax_415FF0(0, v7, libc.CString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c"), 431)
	v9 = int(*(*uint32)(unsafe.Pointer(uintptr(a1 + v8*4 + 440))))
	v10 = v8
	if v8 < *(*int)(unsafe.Pointer(uintptr(a1 + 568))) {
		v11 = (*uint32)(unsafe.Pointer(uintptr(a1 + v8*4 + 440)))
		for {
			v10++
			*v11 = *(*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*1))
			v11 = (*uint32)(unsafe.Add(unsafe.Pointer(v11), unsafe.Sizeof(uint32(0))*1))
			if v10 >= *(*int)(unsafe.Pointer(uintptr(a1 + 568))) {
				break
			}
		}
	}
	return v9
}
func sub_452770(a1 *uint32) int {
	var (
		v1 *uint32
		v2 *uint32
		v4 int
		v5 uint
	)
	v1 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*38)))))
	v2 = (*uint32)(unsafe.Pointer(uintptr(sub_451CF0((*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(a1), unsafe.Sizeof(uint32(0))*38)))))))))
	if *(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9)) + 72))) < 0x21 {
		sub_4BDB90(a1, v2)
		return 0
	}
	sub_4BDB90(a1, nil)
	v4 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*9)))
	if (int(*(*uint8)(unsafe.Pointer(uintptr(v4 + 4))))&8) == 0 || v2 != nil || *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*142)) != 0 {
		v5 = uint(nox_common_randomIntMinMax_415FF0(int(*(*uint32)(unsafe.Pointer(uintptr(v4 + 68)))), int(*(*uint32)(unsafe.Pointer(uintptr(v4 + 72)))), libc.CString("C:\\NoxPost\\src\\client\\Audio\\AudEvent.c"), 706))
		if v5 < 0x21 {
			sub_4BDB90(a1, v2)
			return 0
		}
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*71)) = uint32(v5)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*74)) = uint32(uintptr(unsafe.Pointer(v2)))
	}
	return 0
}
func sub_452050(a1_ *struct576) {
	var a1 *uint32 = (*uint32)(unsafe.Pointer(a1_))
	_ = a1
	var v1 *uint32
	var v2 int
	var v3 uint
	var v4 *uint8
	var result *uint32
	var v6 **uint32
	var v7 **uint32
	var v8 *uint32
	v1 = &a1_.field_9.field_0
	v2 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*12)) + a1_.field_75)
	v3 = uint((a1_.timerGroup_46.field_0.field_1 >> 16) / 0x666)
	v4 = (*uint8)(mem_getPtr(0x5D4594, uint32(v2*120+839892)))
	if *(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*26)) == *mem_getU32Ptr(0x5D4594, 1045444) {
		result = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)))))
		if v2 <= int(uintptr(unsafe.Pointer(result))) {
			if (*uint32)(unsafe.Pointer(uintptr(v2))) == result && v3 > uint(*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31))) {
				*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) = uint32(v3)
				v7 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))))
				nox_common_list_remove_425920(unsafe.Pointer(v7))
				nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), v3*12)))))))), (*nox_list_item_t)(unsafe.Pointer(v7)))
			}
		} else {
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)) = uint32(v2)
			*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) = uint32(v3)
			v6 = (**uint32)(unsafe.Pointer((*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))))
			nox_common_list_remove_425920(unsafe.Pointer(v6))
			nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), v3*12)))))))), (*nox_list_item_t)(unsafe.Pointer(v6)))
		}
	} else {
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*26)) = *mem_getU32Ptr(0x5D4594, 1045444)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*27)) = uint32(v2)
		*(*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*31)) = uint32(v3)
		v8 = (*uint32)(unsafe.Add(unsafe.Pointer(v1), unsafe.Sizeof(uint32(0))*28))
		sub_425770(unsafe.Pointer(v8))
		nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer((*uint8)(unsafe.Add(unsafe.Pointer(v4), v3*12)))))))), (*nox_list_item_t)(unsafe.Pointer(v8)))
	}
}
func sub_451BE0(a1_ *struct576) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v1 int
	var v2 int
	var v3 uint
	var v4 *uint32
	var v5 int
	var v6 int
	var v7 *uint32
	var result int
	var v9 int
	var v10 *uint32
	v1 = int(uintptr(unsafe.Pointer(a1_)))
	v2 = int(*(*uint32)(unsafe.Pointer(&a1_.field_9)))
	v3 = uint(a1_.timerGroup_46.field_0.field_1 >> 16)
	v4 = *(**uint32)(unsafe.Pointer(uintptr(v2 + 88)))
	if v4 != (*uint32)(unsafe.Pointer(uintptr(v2+88))) {
		for {
			v5 = int((*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*44)) >> 16) - uint32(v3))
			if v5 < 0 {
				v5 = int(v3 - uint(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*44))>>16))
			}
			if v5 >= int((*(*uint32)(unsafe.Pointer(uintptr(v2 + 20)))>>16)/10) {
				if *(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*44))>>16 < uint32(v3) {
					break
				}
			} else {
				v6 = int(*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*4)))
				if int(*(*uint8)(unsafe.Pointer(uintptr(v2 + 4))))&0x10 != 0 {
					if v6 != 0 {
						break
					}
				} else if v6 == 0 {
					break
				}
			}
			v4 = (*uint32)(unsafe.Pointer(uintptr(*v4)))
			if v4 == (*uint32)(unsafe.Pointer(uintptr(v2+88))) {
				break
			}
		}
		v1 = int(uintptr(unsafe.Pointer(a1_)))
	}
	v7 = (*uint32)(unsafe.Pointer(uintptr(v1 + 12)))
	sub_425770(unsafe.Pointer(uintptr(v1 + 12)))
	nox_common_list_append_4258E0((*nox_list_item_t)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v4)))))), (*nox_list_item_t)(unsafe.Pointer(v7)))
	result = int(*(*uint32)(unsafe.Pointer(uintptr(v2 + 56))))
	v9 = int(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) + 1)
	*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) = uint32(v9)
	if result != 0 {
		if v9 > result {
			v10 = (*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v2 + 92))) - 12)))
			nox_common_list_remove_425920(unsafe.Pointer(*(***uint32)(unsafe.Pointer(uintptr(v2 + 92)))))
			sub_4523D0((*struct576)(unsafe.Pointer(v10)))
			result = int(*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) - 1)
			*(*uint32)(unsafe.Pointer(uintptr(v2 + 52))) = uint32(result)
		}
	}
	return result
}
func sub_451970() {
	sub_4521F0()
	sub_452230()
	if dword_5d4594_1045424 != 0 {
		sub_4BD3C0(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1045424)))
		dword_5d4594_1045424 = 0
	}
	if dword_5d4594_1045436 != 0 {
		sub_4BD2D0(*(*unsafe.Pointer)(unsafe.Pointer(&dword_5d4594_1045436)))
		dword_5d4594_1045436 = 0
	}
	dword_5d4594_1045432 = 0
}
func sub_4519C0() {
	var (
		result int
		v1     int
		v2     int
		v3     int
		v4     int
		v5     *uint8
		v6     *uint8
		v7     *uint8
		v8     int
		v9     int
		v10    int
	)
	result = int(dword_5d4594_1045432)
	if dword_5d4594_1045432 == 0 {
		return
	}
	result = int(*mem_getU32Ptr(0x5D4594, 1045448))
	if *mem_getU32Ptr(0x5D4594, 1045448) != 0 {
		return
	}
	*mem_getU32Ptr(0x5D4594, 1045448) = 1
	sub_486520(unsafe.Pointer(*(**uint)(unsafe.Pointer(&dword_587000_127004))))
	v1 = int(*mem_getU32Ptr(0x5D4594, 840612))
	*mem_getU32Ptr(0x5D4594, 1045440)++
	if unsafe.Pointer(*(**uint8)(mem_getPtr(0x5D4594, 840612))) != mem_getPtr(0x5D4594, 840612) {
		for {
			v2 = int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))))
			if *(*uint32)(unsafe.Pointer(uintptr(v2 + 100))) != *mem_getU32Ptr(0x5D4594, 1045440) {
				nox_common_list_clear_425760((*nox_list_item_t)(unsafe.Pointer(uintptr(v2 + 88))))
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 52))) = 0
				*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 100))) = *mem_getU32Ptr(0x5D4594, 1045440)
			}
			sub_486520(unsafe.Pointer(uintptr(v1 + 184)))
			if *(*uint32)(unsafe.Pointer(uintptr(v1 + 28))) != 4 {
				sub_451BE0((*struct576)(unsafe.Pointer(uintptr(v1))))
			}
			v1 = int(*(*uint32)(unsafe.Pointer(uintptr(v1))))
			if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == mem_getPtr(0x5D4594, 840612) {
				break
			}
		}
		v1 = int(*mem_getU32Ptr(0x5D4594, 840612))
		if unsafe.Pointer(*(**uint8)(mem_getPtr(0x5D4594, 840612))) != mem_getPtr(0x5D4594, 840612) {
			for {
				sub_452510((*struct576)(unsafe.Pointer(uintptr(v1))))
				v1 = int(*(*uint32)(unsafe.Pointer(uintptr(v1))))
				if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) == mem_getPtr(0x5D4594, 840612) {
					break
				}
			}
			v1 = int(*mem_getU32Ptr(0x5D4594, 840612))
		}
	}
	v3 = 0
	sub_452010()
	if unsafe.Pointer((*uint8)(unsafe.Pointer(uintptr(v1)))) != mem_getPtr(0x5D4594, 840612) {
		for {
			v4 = int(*(*uint32)(unsafe.Pointer(uintptr(v1 + 176))))
			v5 = *(**uint8)(unsafe.Pointer(uintptr(v1)))
			if v4 == 0 || v1 != int(*(*uint32)(unsafe.Pointer(uintptr(v4 + 152)))) {
				sub_4523D0((*struct576)(unsafe.Pointer(uintptr(v1))))
			}
			if int(*(*uint8)(unsafe.Pointer(uintptr(v1 + 24))))&1 != 0 {
				sub_451FE0((*struct576)(unsafe.Pointer(uintptr(v1))))
			} else {
				v3 += int(uint((*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(v1 + 36))) + 20)))>>16)*33) >> 14)
				sub_452050((*struct576)(unsafe.Pointer(uintptr(v1))))
			}
			v1 = int(uintptr(unsafe.Pointer(v5)))
			if unsafe.Pointer(v5) == mem_getPtr(0x5D4594, 840612) {
				break
			}
		}
	}
	if v3 <= 100 {
		sub_486350(unsafe.Pointer(uintptr(int(uintptr(mem_getPtr(0x5D4594, 1045228))))), 0x4000)
	} else {
		sub_486350(unsafe.Pointer(uintptr(int(uintptr(mem_getPtr(0x5D4594, 1045228))))), 0x190000/v3)
	}
	result = sub_486520(unsafe.Pointer(mem_getU32Ptr(0x5D4594, 1045228)))
	v6 = *(**uint8)(mem_getPtr(0x5D4594, 840612))
	if unsafe.Pointer(*(**uint8)(mem_getPtr(0x5D4594, 840612))) != mem_getPtr(0x5D4594, 840612) {
		for {
			v7 = *(**uint8)(unsafe.Pointer(v6))
			result = int(*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*7))))
			if result == 1 {
				sub_451DC0(int(uintptr(unsafe.Pointer(v6))))
				v8 = sub_451CA0((*struct576)(unsafe.Pointer(v6)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*74))) = uint32(v8)
				if v8 == 0 {
					for {
						if sub_452120((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v6))))))) == nil {
							break
						}
						v7 = *(**uint8)(unsafe.Pointer(v6))
						sub_451DC0(int(uintptr(unsafe.Pointer(v6))))
						v9 = sub_451CA0((*struct576)(unsafe.Pointer(v6)))
						*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*74))) = uint32(v9)
						if v9 != 0 {
							break
						}
					}
				}
				v10 = sub_451CA0((*struct576)(unsafe.Pointer(v6)))
				*((*uint32)(unsafe.Add(unsafe.Pointer((*uint32)(unsafe.Pointer(v6))), unsafe.Sizeof(uint32(0))*74))) = uint32(v10)
				if v10 == 0 || (func() int {
					result = sub_452490((*struct576)(unsafe.Pointer(v6)))
					return result
				}()) == 0 {
					sub_4523D0((*struct576)(unsafe.Pointer(v6)))
					result = sub_451FE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v6)))))))
				}
			}
			v6 = v7
			if unsafe.Pointer(v7) == mem_getPtr(0x5D4594, 840612) {
				break
			}
		}
	}
	*mem_getU32Ptr(0x5D4594, 1045448) = 0
}
func sub_4BDB20(a1 int) int {
	var result int
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 124))) |= 0x10
	return result
}
func sub_4BD710(a1 int) int {
	return a1 + 24
}
func sub_4526D0(a1 int) int {
	*(*uint32)(unsafe.Pointer(uintptr(*(*uint32)(unsafe.Pointer(uintptr(a1 + 152))) + 28))) = 4
	return 0
}
func sub_452FE0(a1_ *struct576, a2 int) int {
	var a1 int = int(uintptr(unsafe.Pointer(a1_)))
	_ = a1
	var v2 int
	v2 = sub_452FA0(a2)
	return sub_486350(unsafe.Pointer(&a1_.timerGroup_46.field_16), v2)
}
func sub_452FA0(a1 int) int {
	var v1 int
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
func sub_4BD650(a1 int) int {
	var result int
	result = a1
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12)))++
	return result
}
func sub_4BD660(a1 int) int {
	var result int
	result = int(*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) - 1)
	*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = uint32(result)
	if result < 0 {
		*(*uint32)(unsafe.Pointer(uintptr(a1 + 12))) = 0
	}
	return result
}
func nox_xxx_clientPlaySoundSpecial_452D80(a1 int, a2 int) {
	var (
		result *uint32
		v3     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = nox_xxx_draw_452300((*struct200)(unsafe.Pointer(result)))
	v3 = result
	if result == nil {
		return
	}
	sub_452EE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(result)))))), a2)
	sub_452510((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v3)))))))
}
func sub_452DC0(a1 int, a2 int, a3 int) {
	var (
		result *uint32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = nox_xxx_draw_452300((*struct200)(unsafe.Pointer(result)))
	v4 = result
	if result == nil {
		return
	}
	sub_452EE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(result)))))), a2)
	sub_452F80((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v4)))))), a3)
	sub_452510((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v4)))))))
}
func sub_452E10(a1 int, a2 int, a3 int) {
	var (
		result *uint32
		v4     *uint32
	)
	result = (*uint32)(unsafe.Pointer(nox_xxx_draw_452270(a1)))
	if result == nil {
		return
	}
	result = nox_xxx_draw_452300((*struct200)(unsafe.Pointer(result)))
	v4 = result
	if result == nil {
		return
	}
	sub_452EE0((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(result)))))), a2)
	sub_452F80((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v4)))))), a3)
	*(*uint32)(unsafe.Add(unsafe.Pointer(v4), unsafe.Sizeof(uint32(0))*75)) = 2
	sub_452510((*struct576)(unsafe.Pointer(uintptr(int(uintptr(unsafe.Pointer(v4)))))))
}
