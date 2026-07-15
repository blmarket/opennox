package legacy

/*
#include <stdint.h>
#include <stdlib.h>

typedef struct nox_object_t nox_object_t;

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
int nox_xxx_monsterCanCast_534300(nox_object_t* a1);
int sub_534340(int a1);
int sub_5343C0(int a1);
int sub_534400(int a1);
int sub_536550(char* a1, uint32_t* a2);
int sub_536580(char* a1, int a2);
int sub_536600(char* a1, int a2);
int sub_536D80(char* a1, int a2);
int sub_536DE0(char* a1, uint8_t* a2);
int sub_536E50(char* a1, uint8_t* a2);
int sub_537750(int a1);
int sub_5408A0(int a1);
int sub_544740(int a1);
int sub_544750(int a1);
int nox_xxx_mobActionReturnToHome_544920(int a1);
int sub_544930(int a1);
int sub_544940(int a1);
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

func allocMonsterState() (outer, inner unsafe.Pointer) {
	outer = C.calloc(1, 800)
	inner = C.calloc(1, 1600)
	*(*uint32)(unsafe.Pointer(uintptr(outer) + 748)) = uint32(uintptr(inner))
	return outer, inner
}

func freeMonsterState(outer, inner unsafe.Pointer) {
	C.free(outer)
	C.free(inner)
}

func C_nox_xxx_monsterCanCast_534300(flags uint32) int {
	outer, inner := allocMonsterState()
	defer freeMonsterState(outer, inner)
	*(*uint32)(unsafe.Pointer(uintptr(inner) + 1440)) = flags
	return int(C.nox_xxx_monsterCanCast_534300((*C.nox_object_t)(outer)))
}

func C_sub_534340(action int8) int {
	outer, inner := allocMonsterState()
	defer freeMonsterState(outer, inner)
	*(*int8)(unsafe.Pointer(uintptr(inner) + 544)) = action
	*(*uint32)(unsafe.Pointer(uintptr(inner) + uintptr(24*int(action+23)))) = uint32(action)
	return int(C.sub_534340(C.int(uintptr(outer))))
}

func C_sub_5343C0(value float32) int {
	outer, inner := allocMonsterState()
	defer freeMonsterState(outer, inner)
	*(*float32)(unsafe.Pointer(uintptr(inner) + 1304)) = value
	return int(C.sub_5343C0(C.int(uintptr(outer))))
}

func C_sub_534400(value float32) int {
	outer, inner := allocMonsterState()
	defer freeMonsterState(outer, inner)
	*(*float32)(unsafe.Pointer(uintptr(inner) + 1304)) = value
	return int(C.sub_534400(C.int(uintptr(outer))))
}

func C_sub_536550(s string) (ret int, values [3]float32) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	buf := C.calloc(3, C.size_t(unsafe.Sizeof(uint32(0))))
	defer C.free(buf)
	ret = int(C.sub_536550(cs, (*C.uint32_t)(buf)))
	values[0] = *(*float32)(buf)
	values[1] = *(*float32)(unsafe.Pointer(uintptr(buf) + 4))
	values[2] = *(*float32)(unsafe.Pointer(uintptr(buf) + 8))
	return ret, values
}

func C_sub_536580(s string) (ret int, values [3]int32) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	buf := C.calloc(3, C.size_t(unsafe.Sizeof(int32(0))))
	defer C.free(buf)
	ret = int(C.sub_536580(cs, C.int(uintptr(buf))))
	values = *(*[3]int32)(buf)
	return ret, values
}

func cParseOneInt(s string, fn func(*C.char, C.int) C.int) (ret, value int) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	buf := C.calloc(1, C.size_t(unsafe.Sizeof(int32(0))))
	defer C.free(buf)
	ret = int(fn(cs, C.int(uintptr(buf))))
	value = int(*(*int32)(buf))
	return ret, value
}

func C_sub_536600(s string) (ret, value int) {
	return cParseOneInt(s, func(cs *C.char, p C.int) C.int { return C.sub_536600(cs, p) })
}

func C_sub_536D80(s string) (ret, value int) {
	return cParseOneInt(s, func(cs *C.char, p C.int) C.int { return C.sub_536D80(cs, p) })
}

func C_sub_536DE0(s string) (ret int, value uint8) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	ret = int(C.sub_536DE0(cs, (*C.uint8_t)(unsafe.Pointer(&value))))
	return ret, value
}

func C_sub_536E50(s string) (ret int, value uint8) {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	ret = int(C.sub_536E50(cs, (*C.uint8_t)(unsafe.Pointer(&value))))
	return ret, value
}

func C_sub_537750(value uint32, nilObject bool) int {
	if nilObject {
		return int(C.sub_537750(0))
	}
	buf := C.calloc(1, 464)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 460)) = value
	return int(C.sub_537750(C.int(uintptr(buf))))
}

func C_sub_5408A0(action int8) int {
	outer, inner := allocMonsterState()
	defer freeMonsterState(outer, inner)
	*(*int8)(unsafe.Pointer(uintptr(inner) + 544)) = action
	*(*uint32)(unsafe.Pointer(uintptr(inner) + uintptr(24*int(action+23)))) = uint32(action)
	return int(C.sub_5408A0(C.int(uintptr(outer))))
}

func C_monsterActionFlagAliases(initial uint32) (setResults [2]uint32, clearResults [3]uint32) {
	outer, inner := allocMonsterState()
	defer freeMonsterState(outer, inner)
	flags := (*uint32)(unsafe.Pointer(uintptr(inner) + 1440))
	*flags = initial
	setResults[0] = uint32(C.sub_544740(C.int(uintptr(outer))))
	*flags = initial
	setResults[1] = uint32(C.nox_xxx_mobActionReturnToHome_544920(C.int(uintptr(outer))))
	*flags = initial
	clearResults[0] = uint32(C.sub_544750(C.int(uintptr(outer))))
	*flags = initial
	clearResults[1] = uint32(C.sub_544930(C.int(uintptr(outer))))
	*flags = initial
	clearResults[2] = uint32(C.sub_544940(C.int(uintptr(outer))))
	return setResults, clearResults
}
