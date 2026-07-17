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
char sub_532110(int a1);
char nox_xxx_mobActionGetUp_534A90(int a1);
void nox_xxx_lightningSpellTrapEffect_530020(int a1, int a2);
int nox_xxx_castTele_530820(int a1);
int sub_530880(int* a1);
int sub_53A680(int a1);
char sub_53B300(int a1);
void nox_xxx_updateLifetime_53B8F0(int unit);
void nox_xxx_updateBlow_53C160(int a3);
int nox_xxx_rechargeItem_53C520(int a1, int a2);
int nox_xxx_getRechargeRate_53C940(uint32_t* a1);
void nox_xxx_updateBlackPowderBarrel_53C9A0(float* a1);
void nox_xxx_updateWaterBarrel_53CB90(int a1);
void nox_xxx_waterBarrel_53CC30(float* a1, int a2);
void nox_xxx_updateBlackPowderBurn_53CCB0(int a1);
void nox_xxx_updateDeathBallFragment_53D220(int a1);
void nox_xxx_updateFist_53D400(int a1);
void nox_xxx_updateMeteorShower_53D5A0(float* a2);
void nox_xxx_updateToxicCloud_53D850(int a1);
void sub_53D8C0(int a1, int a2);
void nox_xxx_updateArachnaphobia_53DA60(int* a1);
void nox_xxx_updateBreakAndRemove_53DC30(uint32_t* a1);
char* sub_5435C0(int a1, int a2, int a3, int a4);
char* sub_543620(int a1, int a2);
int sub_543680(float* a1);
void sub_543BC0(int a1, int a2, int a3, int a4, int a5, int a6);
int nox_xxx_tile_543C50(uint32_t* a1, int a2, int a3, int a4, int a5, int a6);
int sub_543FB0(const char* a1);
int sub_544020(char* a1);
int nox_xxx_tileCheckByte3_544070(int a1);
int nox_xxx_tileCheckByte4_5440A0(int a1);
int sub_544AE0(int a1, float a2);
void sub_544B20(int a1, int a2);
int nox_xxx_useLesserFireballStaff_53F290(int a1, uint32_t* a2);
int sub_53F830(int a1, int a2);

extern uint32_t dword_5d4594_3835356;
extern uint32_t dword_5d4594_2489436;
*/
import "C"
import "unsafe"

type game43SafeUpdateResult struct {
	emptyWeapon                         int
	toggleOn, toggleOff                 uint8
	deadline                            uint32
	teleportEmpty                       int
	rechargeMissing, rechargeChanged    int
	rechargeUnchanged                   int
	charge, percent                     uint8
	rechargeRateNil                     int
	barrelNext, burnNext, toxicRemain   uint32
	longName, shortName                 string
	tileDisabled, tileDelete            int
	tileMissing, tileNone               int
	tileInvalid, tileSubtileWithoutType int
	nearbyItem                          int
	lesserStaff, readGuard              int
}

func C_game43SafeUpdatePaths() (out game43SafeUpdateResult) {
	unit := C.calloc(1, 752)
	data := C.calloc(1, 512)
	defer C.free(unit)
	defer C.free(data)
	*(*uint32)(unsafe.Pointer(uintptr(unit) + 748)) = uint32(uintptr(data))
	_ = C.sub_532110(C.int(uintptr(unit)))
	_ = C.nox_xxx_mobActionGetUp_534A90(C.int(uintptr(unit)))
	C.nox_xxx_lightningSpellTrapEffect_530020(C.int(uintptr(unit)), C.int(uintptr(unit)))

	teleport := C.calloc(1, 72)
	defer C.free(teleport)
	out.teleportEmpty = int(C.nox_xxx_castTele_530820(C.int(uintptr(teleport))))
	out.deadline = *(*uint32)(unsafe.Pointer(uintptr(teleport) + 68))
	teleData := C.calloc(5, 4)
	defer C.free(teleData)
	_ = C.sub_530880((*C.int)(teleData))

	out.emptyWeapon = int(C.sub_53A680(C.int(uintptr(unit))))
	out.toggleOn = uint8(C.sub_53B300(C.int(uintptr(unit))))
	*(*uint32)(unsafe.Pointer(uintptr(unit) + 16)) |= 0x1000000
	out.toggleOff = uint8(C.sub_53B300(C.int(uintptr(unit))))

	lifetime := C.calloc(1, 4)
	defer C.free(lifetime)
	*(*uint32)(lifetime) = 100
	*(*uint32)(unsafe.Pointer(uintptr(unit) + 748)) = uint32(uintptr(lifetime))
	C.nox_xxx_updateLifetime_53B8F0(C.int(uintptr(unit)))
	C.nox_xxx_updateBlow_53C160(C.int(uintptr(unit)))

	item := C.calloc(1, 740)
	useData := C.calloc(1, 116)
	defer C.free(item)
	defer C.free(useData)
	out.rechargeMissing = int(C.nox_xxx_rechargeItem_53C520(C.int(uintptr(item)), 50))
	*(*uint32)(unsafe.Pointer(uintptr(item) + 736)) = uint32(uintptr(useData))
	*(*uint8)(unsafe.Pointer(uintptr(useData) + 109)) = 10
	out.rechargeChanged = int(C.nox_xxx_rechargeItem_53C520(C.int(uintptr(item)), 50))
	out.rechargeUnchanged = int(C.nox_xxx_rechargeItem_53C520(C.int(uintptr(item)), 0))
	out.charge = *(*uint8)(unsafe.Pointer(uintptr(useData) + 108))
	out.percent = *(*uint8)(unsafe.Pointer(uintptr(useData) + 112))
	out.rechargeRateNil = int(C.nox_xxx_getRechargeRate_53C940(nil))

	update := C.calloc(1, 752)
	updateData := C.calloc(1, 8)
	defer C.free(update)
	defer C.free(updateData)
	*(*uint32)(unsafe.Pointer(uintptr(update) + 136)) = 2
	C.nox_xxx_updateBlackPowderBarrel_53C9A0((*C.float)(update))
	out.barrelNext = *(*uint32)(unsafe.Pointer(uintptr(update) + 136))
	C.nox_xxx_updateWaterBarrel_53CB90(C.int(uintptr(update)))
	C.nox_xxx_waterBarrel_53CC30((*C.float)(update), 0)
	*(*uint32)(unsafe.Pointer(uintptr(update) + 136)) = 0
	C.nox_xxx_updateBlackPowderBurn_53CCB0(C.int(uintptr(update)))
	out.burnNext = *(*uint32)(unsafe.Pointer(uintptr(update) + 136))
	C.nox_xxx_updateDeathBallFragment_53D220(C.int(uintptr(update)))
	*(*float32)(unsafe.Pointer(uintptr(update) + 104)) = 1
	C.nox_xxx_updateFist_53D400(C.int(uintptr(update)))
	*(*uint32)(unsafe.Pointer(uintptr(update) + 136)) = 1
	C.nox_xxx_updateMeteorShower_53D5A0((*C.float)(update))
	*(*uint32)(unsafe.Pointer(uintptr(update) + 748)) = uint32(uintptr(updateData))
	*(*uint32)(updateData) = 2
	*(*uint32)(unsafe.Pointer(uintptr(update) + 136)) = 0
	C.nox_xxx_updateToxicCloud_53D850(C.int(uintptr(update)))
	out.toxicRemain = *(*uint32)(updateData)
	C.sub_53D8C0(C.int(uintptr(update)), C.int(uintptr(unit)))
	*(*uint32)(unsafe.Pointer(uintptr(update) + 136)) = 1
	C.nox_xxx_updateArachnaphobia_53DA60((*C.int)(update))
	C.nox_xxx_updateBreakAndRemove_53DC30((*C.uint32_t)(update))

	longInput := C.CString("name")
	none := C.CString("NONE")
	defer C.free(unsafe.Pointer(longInput))
	defer C.free(unsafe.Pointer(none))
	out.longName = C.GoString(C.sub_5435C0(C.int(uintptr(unsafe.Pointer(longInput))), 1, 2, 3))
	out.shortName = C.GoString(C.sub_543620(C.int(uintptr(unsafe.Pointer(longInput))), 4))
	savedType := C.dword_5d4594_3835356
	savedSet := C.dword_5d4594_2489436
	C.dword_5d4594_3835356 = 255
	C.dword_5d4594_2489436 = 0
	defer func() {
		C.dword_5d4594_3835356 = savedType
		C.dword_5d4594_2489436 = savedSet
	}()
	point := [2]C.float{0, 0}
	out.tileDisabled = int(C.sub_543680(&point[0]))
	C.sub_543BC0(0, 1, 0, 0, 0, 0)
	tile := C.calloc(5, 4)
	defer C.free(tile)
	out.tileDelete = int(C.nox_xxx_tile_543C50((*C.uint32_t)(tile), 255, 0, 0, 0, 0))
	out.tileMissing = int(C.sub_543FB0(nil))
	out.tileNone = int(C.sub_544020(none))
	out.tileInvalid = int(C.nox_xxx_tileCheckByte3_544070(-1))
	C.dword_5d4594_2489436 = 0
	out.tileSubtileWithoutType = int(C.nox_xxx_tileCheckByte4_5440A0(99))

	out.nearbyItem = int(C.sub_544AE0(C.int(uintptr(unit)), 1))
	C.sub_544B20(C.int(uintptr(item)), C.int(uintptr(unit)))
	*(*uint32)(unsafe.Pointer(uintptr(item) + 736)) = 0
	out.lesserStaff = int(C.nox_xxx_useLesserFireballStaff_53F290(C.int(uintptr(unit)), (*C.uint32_t)(item)))
	out.readGuard = int(C.sub_53F830(C.int(uintptr(unit)), C.int(uintptr(item))))
	return out
}

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
