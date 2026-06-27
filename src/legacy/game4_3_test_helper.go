package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int nox_xxx_spellTurnUndeadUpdate_531410();
int sub_534020(int a1);
int nox_xxx_monsterIsMoveing_534320(int a1);
int sub_534440(int a1);
double sub_534470(int a1);
int sub_5347A0(void* a1);
int nox_xxx_isNotPoisoned_5347F0(int a1);
int sub_537580(int a1);
int sub_540D20(int a1);
int sub_534750(int a1);
int sub_534780(int a1);
int nox_xxx_monsterCanAttackAtWill_534390(void* a1);
*/
import "C"
import "unsafe"

func C_nox_xxx_spellTurnUndeadUpdate_531410() int {
	return int(C.nox_xxx_spellTurnUndeadUpdate_531410())
}

func C_sub_534020(val uint32) int {
	buf := C.malloc(16)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 12)) = val
	return int(C.sub_534020(C.int(uintptr(buf))))
}

func C_nox_xxx_monsterIsMoveing_534320(f float32) int {
	buf := C.malloc(560)
	defer C.free(buf)
	*(*float32)(unsafe.Pointer(uintptr(buf) + 548)) = f
	return int(C.nox_xxx_monsterIsMoveing_534320(C.int(uintptr(buf))))
}

func C_sub_534440(f float32) int {
	a1 := C.malloc(800)
	defer C.free(a1)
	inner := C.malloc(1400)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 748)) = uint32(uintptr(inner))
	*(*float32)(unsafe.Pointer(uintptr(inner) + 1304)) = f
	return int(C.sub_534440(C.int(uintptr(a1))))
}

func C_sub_534470(f float32) float64 {
	a1 := C.malloc(800)
	defer C.free(a1)
	inner1 := C.malloc(512)
	inner2 := C.malloc(128)
	defer C.free(inner1)
	defer C.free(inner2)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 748)) = uint32(uintptr(inner1))
	*(*uint32)(unsafe.Pointer(uintptr(inner1) + 484)) = uint32(uintptr(inner2))
	*(*float32)(unsafe.Pointer(uintptr(inner2) + 112)) = f
	return float64(C.sub_534470(C.int(uintptr(a1))))
}

func C_sub_5347A0(val uint32) int {
	a1 := C.malloc(800)
	defer C.free(a1)
	inner := C.malloc(1500)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 748)) = uint32(uintptr(inner))
	*(*uint32)(unsafe.Pointer(uintptr(inner) + 1440)) = val
	return int(C.sub_5347A0(unsafe.Pointer(a1)))
}

func C_nox_xxx_isNotPoisoned_5347F0(b uint8) int {
	buf := C.malloc(560)
	defer C.free(buf)
	*(*uint8)(unsafe.Pointer(uintptr(buf) + 540)) = b
	return int(C.nox_xxx_isNotPoisoned_5347F0(C.int(uintptr(buf))))
}

func C_sub_537580(b uint8) int {
	buf := C.malloc(480)
	defer C.free(buf)
	*(*uint8)(unsafe.Pointer(uintptr(buf) + 464)) = b
	return int(C.sub_537580(C.int(uintptr(buf))))
}

func C_sub_540D20(a1 int) int {
	return int(C.sub_540D20(C.int(a1)))
}

func C_sub_534750(initial uint32) (ret int, after uint32) {
	a1 := C.malloc(800)
	inner := C.malloc(1500)
	defer C.free(a1)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 748)) = uint32(uintptr(inner))
	*(*uint32)(unsafe.Pointer(uintptr(inner) + 1440)) = initial
	ret = int(C.sub_534750(C.int(uintptr(a1))))
	after = *(*uint32)(unsafe.Pointer(uintptr(inner) + 1440))
	return
}

func C_sub_534780(initial uint32) (ret int, after uint32) {
	a1 := C.malloc(800)
	inner := C.malloc(1500)
	defer C.free(a1)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 748)) = uint32(uintptr(inner))
	*(*uint32)(unsafe.Pointer(uintptr(inner) + 1440)) = initial
	ret = int(C.sub_534780(C.int(uintptr(a1))))
	after = *(*uint32)(unsafe.Pointer(uintptr(inner) + 1440))
	return
}

func C_nox_xxx_monsterCanAttackAtWill_534390(f float32) int {
	a1 := C.malloc(800)
	inner := C.malloc(1400)
	defer C.free(a1)
	defer C.free(inner)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 748)) = uint32(uintptr(inner))
	*(*float32)(unsafe.Pointer(uintptr(inner) + 1304)) = f
	return int(C.nox_xxx_monsterCanAttackAtWill_534390(unsafe.Pointer(a1)))
}
