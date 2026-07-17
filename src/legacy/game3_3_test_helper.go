package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include "GAME3_3.h"

int sub_4F2570(int a1);
double nox_xxx_objectGetMass_4E4A70(int a1);
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

type game33ObjectPathsResult struct {
	predicateNil      int
	predicateObject   int
	crownEmpty        int
	gameballEmpty     int
	slavesMissingMask int
	slavesMatching    int
	inventoryAll      int
	inventoryType     int
	inventoryActive   int
	itemMatchNil      int
	itemMatch         int
	inventoryMatch    int
	pixieNil          int
	flagSlot          uintptr
	maxHPNil          int
	maxHPResult       uintptr
	maxHP             uint16
	maxManaNil        int
	maxManaNonPlayer  uintptr
	maxManaResult     uintptr
	maxMana           uint16
}

type game33CollisionGuardResult struct {
	pickupNil        int
	pentagram        int
	pentagramState   uint32
	spellPedestalNil int
	slave2Nil        int
	slave2Empty      int
	slaveNil         int
	slaveEmpty       int
	findServerObject int
	subordinateGuard int
	dropNil          int
	dropAllEmpty     uintptr
	dropFoodInvalid  int
	chestInventory   uintptr
	shapePoint       float64
	shapeCircle      float64
	shapeBoxWidth    float64
	shapeBoxHeight   float64
	poisonRemaining  uint8
	sparkFirst       uint32
	sparkSecond      uint32
}

type game33ClassificationResult struct {
	spellTableMissing   int
	guideTableMissing   int
	foodClass           int
	weaponClass         int
	creatureSpellClass  int
	badCreatureClass    int
	firstModifier       int
	secondModifier      int
	coloredEquipment    int
	inventoryLimitNil   int
	specialSpell        int
	rangeSpell          int
	ordinarySpell       int
	guideAndListMissing int
	engageEmpty         int
	disengageEmpty      int
	crownPickupGuard    int
	extraLifeGuard      int
}

func C_game33ObjectSafePaths() (res game33ObjectPathsResult) {
	C.nox_xxx_unitBecomePet_4E7B00(0, 0)
	player := C.calloc(1, 752)
	defer C.free(player)
	C.nox_xxx_monsterRemoveMonitors_4E7B60((*C.nox_object_t)(player), nil)

	object := C.calloc(1, 520)
	defer C.free(object)
	res.predicateNil = int(C.sub_4E7BC0(0))
	*(*uint32)(unsafe.Pointer(uintptr(object) + 8)) = 4
	res.predicateObject = int(C.sub_4E7BC0(C.int(uintptr(object))))

	crownID := memmap.PtrUint32(0x5D4594, 1567716)
	ballID := memmap.PtrUint32(0x5D4594, 1567720)
	pixieID := memmap.PtrUint32(0x5D4594, 1567728)
	savedCrown, savedBall, savedPixie := *crownID, *ballID, *pixieID
	*crownID, *ballID, *pixieID = 111, 112, 113
	defer func() {
		*crownID, *ballID, *pixieID = savedCrown, savedBall, savedPixie
	}()
	res.crownEmpty = int(C.nox_xxx_unitIsCrown_4E7BE0(C.int(uintptr(object))))
	res.gameballEmpty = int(C.nox_xxx_unitIsGameball_4E7C30(C.int(uintptr(object))))

	slave1 := C.calloc(1, 516)
	slave2 := C.calloc(1, 516)
	defer C.free(slave1)
	defer C.free(slave2)
	*(*uint32)(unsafe.Pointer(uintptr(object) + 516)) = uint32(uintptr(slave1))
	*(*uint32)(unsafe.Pointer(uintptr(slave1) + 8)) = 0x2
	*(*uint32)(unsafe.Pointer(uintptr(slave1) + 12)) = 0x4
	*(*uint32)(unsafe.Pointer(uintptr(slave1) + 512)) = uint32(uintptr(slave2))
	*(*uint32)(unsafe.Pointer(uintptr(slave2) + 8)) = 0x2
	*(*uint32)(unsafe.Pointer(uintptr(slave2) + 12)) = 0x8
	res.slavesMissingMask = int(C.nox_xxx_unitCountSlaves_4E7CF0(C.int(uintptr(object)), 0, 0xC))
	res.slavesMatching = int(C.nox_xxx_unitCountSlaves_4E7CF0(C.int(uintptr(object)), 0x2, 0xC))

	inv1 := C.calloc(1, 500)
	inv2 := C.calloc(1, 500)
	defer C.free(inv1)
	defer C.free(inv2)
	*(*uint32)(unsafe.Pointer(uintptr(object) + 504)) = uint32(uintptr(inv1))
	*(*uint16)(unsafe.Pointer(uintptr(inv1) + 4)) = 21
	*(*uint32)(unsafe.Pointer(uintptr(inv1) + 496)) = uint32(uintptr(inv2))
	*(*uint16)(unsafe.Pointer(uintptr(inv2) + 4)) = 22
	*(*uint8)(unsafe.Pointer(uintptr(inv2) + 16)) = 0x20
	res.inventoryAll = int(C.nox_xxx_inventoryCountObjects_4E7D30(C.int(uintptr(object)), 0))
	res.inventoryType = int(C.nox_xxx_inventoryCountObjects_4E7D30(C.int(uintptr(object)), 21))
	res.inventoryActive = int(C.nox_xxx_inventoryCountObjects_4E7D30(C.int(uintptr(object)), 22))

	item := C.calloc(1, 752)
	defer C.free(item)
	*(*uint16)(unsafe.Pointer(uintptr(item) + 4)) = 21
	res.itemMatchNil = int(C.sub_4E7DE0(0, (*C.nox_object_t)(item)))
	res.itemMatch = int(C.sub_4E7DE0(C.int(uintptr(inv1)), (*C.nox_object_t)(item)))
	res.inventoryMatch = int(C.sub_4E7EC0(C.int(uintptr(object)), (*C.nox_object_t)(item)))

	res.pixieNil = int(C.sub_4E81D0(nil))
	res.flagSlot = uintptr(unsafe.Pointer(C.sub_4E8320(3)))
	door := C.calloc(1, 12)
	defer C.free(door)
	C.nox_xxx_fnFindCloseDoors_4E8340((*C.float)(door), 0)

	health := C.calloc(1, 8)
	hpUnit := C.calloc(1, 560)
	defer C.free(health)
	defer C.free(hpUnit)
	*(*uint32)(unsafe.Pointer(uintptr(hpUnit) + 556)) = uint32(uintptr(health))
	res.maxHPNil = int(C.nox_xxx_unitSetMaxHP_4EE7C0(0, 99))
	res.maxHPResult = uintptr(uint32(C.nox_xxx_unitSetMaxHP_4EE7C0(C.int(uintptr(hpUnit)), 321)))
	res.maxHP = *(*uint16)(unsafe.Pointer(uintptr(health) + 4))

	playerData := C.calloc(1, 12)
	defer C.free(playerData)
	res.maxManaNil = int(C.nox_xxx_playerSetMaxMana_4EECD0(0, 55))
	res.maxManaNonPlayer = uintptr(uint32(C.nox_xxx_playerSetMaxMana_4EECD0(C.int(uintptr(player)), 55)))
	*(*uint8)(unsafe.Pointer(uintptr(player) + 8)) = 4
	*(*uint32)(unsafe.Pointer(uintptr(player) + 748)) = uint32(uintptr(playerData))
	res.maxManaResult = uintptr(uint32(C.nox_xxx_playerSetMaxMana_4EECD0(C.int(uintptr(player)), 654)))
	res.maxMana = *(*uint16)(unsafe.Pointer(uintptr(playerData) + 8))
	return res
}

func C_game33CollisionGuardPaths() (out game33CollisionGuardResult) {
	obj := C.calloc(1, 752)
	update := C.calloc(1, 32)
	collideData := C.calloc(1, 16)
	xfer := C.calloc(1, 16)
	defer C.free(obj)
	defer C.free(update)
	defer C.free(collideData)
	defer C.free(xfer)
	*(*uint32)(unsafe.Pointer(uintptr(obj) + 748)) = uint32(uintptr(update))
	*(*uint32)(unsafe.Pointer(uintptr(obj) + 700)) = uint32(uintptr(collideData))
	*(*uint32)(unsafe.Pointer(uintptr(obj) + 692)) = uint32(uintptr(xfer))

	C.nox_xxx_collideDoor_4E8AC0(C.int(uintptr(obj)), 0)
	out.pickupNil = int(C.nox_xxx_collidePickup_4E8DF0(C.int(uintptr(obj)), 0))
	C.nox_xxx_collideManadrain_4E9490(C.int(uintptr(obj)), 0)
	C.nox_xxx_collideDie_4E99B0(C.int(uintptr(obj)), 0)
	C.sub_4EA2C0(C.int(uintptr(obj)), 0)
	C.nox_xxx_collideWebbing_4EA380(C.int(uintptr(obj)), 0)
	C.sub_4EAAA0(C.int(uintptr(obj)))
	out.pentagram = int(C.nox_xxx_collidePentagram_4EAB20(C.int(uintptr(obj))))
	out.pentagramState = *(*uint32)(unsafe.Pointer(uintptr(update) + 4))
	out.spellPedestalNil = int(C.nox_xxx_collideSpellPedestal_4EAD20(C.int(uintptr(obj)), 0))
	*(*float32)(unsafe.Pointer(uintptr(obj) + 104)) = 1
	C.nox_xxx_collideFist_4EADF0(C.int(uintptr(obj)), 0)
	C.nox_xxx_collideTeleportWake_4EAE30(C.int(uintptr(obj)), 0)
	C.nox_xxx_collideBearTrap_4EB890((*C.int)(obj), 0)
	C.nox_xxx_collidePoisonGasTrap_4EB910((*C.int)(obj), 0)
	out.subordinateGuard = int(C.sub_4EBB50(C.int(uintptr(obj)), 0))
	C.nox_xxx_collideUndeadKiller_4EBD40(C.int(uintptr(obj)), 0, 1)
	C.nox_xxx_collideMonsterGen_4EBE10(C.int(uintptr(obj)), 0)
	C.sub_4EBE40(C.int(uintptr(obj)), 0)
	C.nox_xxx_collideAnkhQuest_4EBF40(C.int(uintptr(obj)), 0)

	out.slave2Nil = int(C.nox_xxx_playerObserverFindGoodSlave2_4EC3E0(0))
	out.slave2Empty = int(C.nox_xxx_playerObserverFindGoodSlave2_4EC3E0(C.int(uintptr(obj))))
	out.slaveNil = int(C.nox_xxx_playerObserverFindGoodSlave_4EC420(0))
	out.slaveEmpty = int(C.nox_xxx_playerObserverFindGoodSlave_4EC420(C.int(uintptr(obj))))
	out.findServerObject = int(C.sub_4ECF10(123))

	out.dropNil = int(C.nox_xxx_drop_4ED790((*C.nox_object_t)(obj), nil, nil))
	out.dropAllEmpty = uintptr(unsafe.Pointer(C.nox_xxx_dropAllItems_4EDA40((*C.uint32_t)(obj))))
	out.dropFoodInvalid = int(C.nox_xxx_dropFood_4EDE50(0, 0, nil))
	C.nox_xxx_chest_4EDF00(C.int(uintptr(obj)), C.int(uintptr(obj)))
	out.chestInventory = uintptr(*(*uint32)(unsafe.Pointer(uintptr(obj) + 504)))

	*(*uint32)(unsafe.Pointer(uintptr(obj) + 172)) = 0
	out.shapePoint = float64(C.sub_4EE2A0(C.int(uintptr(obj))))
	*(*uint32)(unsafe.Pointer(uintptr(obj) + 172)) = 2
	*(*float32)(unsafe.Pointer(uintptr(obj) + 176)) = 7.5
	out.shapeCircle = float64(C.sub_4EE2A0(C.int(uintptr(obj))))
	*(*uint32)(unsafe.Pointer(uintptr(obj) + 172)) = 3
	*(*float32)(unsafe.Pointer(uintptr(obj) + 184)) = 8
	*(*float32)(unsafe.Pointer(uintptr(obj) + 188)) = 4
	out.shapeBoxWidth = float64(C.sub_4EE2A0(C.int(uintptr(obj))))
	*(*float32)(unsafe.Pointer(uintptr(obj) + 184)) = 2
	out.shapeBoxHeight = float64(C.sub_4EE2A0(C.int(uintptr(obj))))

	C.nox_xxx_unitAdjustHP_4EE460((*C.nox_object_t)(obj), 10)
	C.nox_xxx_mobInformOwnerHP_4EE4C0(nil)
	C.nox_xxx_mobInformOwnerHP_4EE4C0((*C.nox_object_t)(obj))
	*(*uint8)(unsafe.Pointer(uintptr(obj) + 540)) = 5
	C.nox_xxx_updatePoison_4EE8F0((*C.nox_object_t)(obj), 2)
	out.poisonRemaining = *(*uint8)(unsafe.Pointer(uintptr(obj) + 540))
	C.nox_xxx_unitSparkInit_4F0390(C.int(uintptr(obj)))
	out.sparkFirst = *(*uint32)(update)
	out.sparkSecond = *(*uint32)(unsafe.Pointer(uintptr(update) + 4))
	return out
}

func C_game33ClassificationPaths() (out game33ClassificationResult) {
	tableOffsets := []uintptr{207032, 207108, 207796}
	idOffsets := []uintptr{
		1568308, 1568312, 1568316, 1568320, 1568324,
		1568328, 1568332, 1568336, 1568340, 1568344, 1568348, 1568352,
		1568356, 1568360, 1568364, 1568368, 1568372, 1568376, 1568380,
		1568384, 1568388, 1568392, 1568396, 1568400, 1568404,
	}
	savedTables := make([]uint32, len(tableOffsets))
	savedIDs := make([]uint32, len(idOffsets))
	for i, off := range tableOffsets {
		p := memmap.PtrUint32(0x587000, off)
		savedTables[i], *p = *p, 0
	}
	for i, off := range idOffsets {
		p := memmap.PtrUint32(0x5D4594, off)
		savedIDs[i], *p = *p, uint32(1000+i)
	}
	defer func() {
		for i, off := range tableOffsets {
			*memmap.PtrUint32(0x587000, off) = savedTables[i]
		}
		for i, off := range idOffsets {
			*memmap.PtrUint32(0x5D4594, off) = savedIDs[i]
		}
	}()

	item := C.calloc(1, 752)
	initData := C.calloc(1, 16)
	useData := C.calloc(1, 8)
	holder := C.calloc(1, 752)
	holderUpdate := C.calloc(1, 16)
	defer C.free(item)
	defer C.free(initData)
	defer C.free(useData)
	defer C.free(holder)
	defer C.free(holderUpdate)
	*(*uint32)(unsafe.Pointer(uintptr(item) + 692)) = uint32(uintptr(initData))
	*(*uint32)(unsafe.Pointer(uintptr(item) + 736)) = uint32(uintptr(useData))
	*(*uint32)(unsafe.Pointer(uintptr(holder) + 748)) = uint32(uintptr(holderUpdate))

	out.spellTableMissing = int(C.sub_4F24E0(17))
	out.guideTableMissing = int(C.sub_4F2530(17))

	*(*uint32)(unsafe.Pointer(uintptr(item) + 8)) = 0x40
	out.foodClass = int(C.sub_4F2590(C.int(uintptr(item))))
	*(*uint32)(unsafe.Pointer(uintptr(item) + 8)) = 0x10
	*(*uint32)(unsafe.Pointer(uintptr(item) + 12)) = 8
	out.weaponClass = int(C.sub_4F2590(C.int(uintptr(item))))
	*(*uint32)(unsafe.Pointer(uintptr(item) + 8)) = 0x100
	*(*uint32)(unsafe.Pointer(uintptr(item) + 12)) = 4
	*(*uint8)(useData) = 1
	out.creatureSpellClass = int(C.sub_4F2700(C.int(uintptr(item))))
	*(*uint8)(useData) = 6
	out.badCreatureClass = int(C.sub_4F2700(C.int(uintptr(item))))

	*(*uint32)(unsafe.Pointer(uintptr(item) + 8)) = 0
	out.firstModifier = int(C.sub_4F27E0(C.int(uintptr(item))))
	out.secondModifier = int(C.sub_4F28C0(C.int(uintptr(item))))
	out.coloredEquipment = int(C.sub_4F2B60(C.int(uintptr(item))))
	out.inventoryLimitNil = int(C.sub_4F2C30(0))
	out.specialSpell = int(C.nox_xxx_spell_4F2E70(46))
	out.rangeSpell = int(C.nox_xxx_spell_4F2E70(100))
	out.ordinarySpell = int(C.nox_xxx_spell_4F2E70(17))
	out.guideAndListMissing = int(C.sub_4F2EF0(17))

	out.engageEmpty = int(C.nox_xxx_itemApplyEngageEffect_4F2FF0((*C.nox_object_t)(item), 0))
	out.disengageEmpty = int(C.nox_xxx_itemApplyDisengageEffect_4F3030((*C.nox_object_t)(item), 0))
	C.nox_xxx_inventoryPutImpl_4F3070(nil, nil, 0)
	out.crownPickupGuard = int(C.sub_4F3400(C.int(uintptr(item)), C.int(uintptr(holder)), 0))
	out.extraLifeGuard = int(C.sub_4F3DD0(C.int(uintptr(item)), C.int(uintptr(holder))))
	return out
}

// C_sub_4F2570 is pure arithmetic: returns 1 iff 0 < a1 < 6.
func C_sub_4F2570(a1 int) int {
	return int(C.sub_4F2570(C.int(a1)))
}

// C_nox_xxx_objectGetMass_4E4A70 reads float at offset +120 and returns as double.
func C_nox_xxx_objectGetMass_4E4A70(val float32) float64 {
	buf := C.malloc(256)
	defer C.free(buf)
	*(*float32)(unsafe.Pointer(uintptr(buf) + 120)) = val
	return float64(C.nox_xxx_objectGetMass_4E4A70(C.int(uintptr(buf))))
}

// C_nox_xxx_inventoryGetFirst_4E7980 reads uint32 at offset +504.
func C_nox_xxx_inventoryGetFirst_4E7980(val uint32) int {
	buf := C.malloc(1024)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 504)) = val
	return int(C.nox_xxx_inventoryGetFirst_4E7980(C.int(uintptr(buf))))
}
