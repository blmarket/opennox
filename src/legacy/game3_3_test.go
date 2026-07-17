package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGame33Pure covers pure helpers in GAME3_3.c.
func TestGame33Pure(t *testing.T) {
	t.Run("safe object and inventory paths", func(t *testing.T) {
		got := C_game33ObjectSafePaths()
		require.Zero(t, got.predicateNil)
		require.Equal(t, 1, got.predicateObject)
		require.Zero(t, got.crownEmpty)
		require.Zero(t, got.gameballEmpty)
		require.Zero(t, got.slavesMissingMask)
		require.Equal(t, 2, got.slavesMatching)
		require.Equal(t, 2, got.inventoryAll)
		require.Equal(t, 1, got.inventoryType)
		require.Zero(t, got.inventoryActive)
		require.Zero(t, got.itemMatchNil)
		require.Equal(t, 1, got.itemMatch)
		require.Equal(t, 1, got.inventoryMatch)
		require.Equal(t, 113, got.pixieNil)
		require.NotZero(t, got.flagSlot)
		require.Zero(t, got.maxHPNil)
		require.NotZero(t, got.maxHPResult)
		require.Equal(t, uint16(321), got.maxHP)
		require.Zero(t, got.maxManaNil)
		require.NotZero(t, got.maxManaNonPlayer)
		require.NotZero(t, got.maxManaResult)
		require.Equal(t, uint16(654), got.maxMana)
	})
	t.Run("sub_4F2570 range check", func(t *testing.T) {
		require.Equal(t, 0, C_sub_4F2570(0))
		require.Equal(t, 0, C_sub_4F2570(-1))
		require.Equal(t, 1, C_sub_4F2570(1))
		require.Equal(t, 1, C_sub_4F2570(5))
		require.Equal(t, 0, C_sub_4F2570(6))
		require.Equal(t, 0, C_sub_4F2570(100))
	})

	t.Run("nox_xxx_objectGetMass_4E4A70 reads float at +120", func(t *testing.T) {
		require.Equal(t, float64(0), C_nox_xxx_objectGetMass_4E4A70(0))
		require.Equal(t, float64(1.5), C_nox_xxx_objectGetMass_4E4A70(1.5))
		require.Equal(t, float64(-2.25), C_nox_xxx_objectGetMass_4E4A70(-2.25))
		require.Equal(t, float64(123.125), C_nox_xxx_objectGetMass_4E4A70(123.125))
	})

	t.Run("nox_xxx_inventoryGetFirst_4E7980 reads u32 at +504", func(t *testing.T) {
		require.Equal(t, 0, C_nox_xxx_inventoryGetFirst_4E7980(0))
		require.Equal(t, 1, C_nox_xxx_inventoryGetFirst_4E7980(1))
		require.Equal(t, 0x7ABCDEF0, C_nox_xxx_inventoryGetFirst_4E7980(0x7ABCDEF0))
		require.Equal(t, 0x12345678, C_nox_xxx_inventoryGetFirst_4E7980(0x12345678))
	})
}

func TestGame33CollisionAndTransferGuards(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setServer(0, 30)
	got := C_game33CollisionGuardPaths()
	require.Zero(t, got.pickupNil)
	require.NotZero(t, got.pentagram)
	require.Equal(t, uint32(1), got.pentagramState)
	require.Zero(t, got.spellPedestalNil)
	require.Zero(t, got.slave2Nil)
	require.Zero(t, got.slave2Empty)
	require.Zero(t, got.slaveNil)
	require.Zero(t, got.slaveEmpty)
	require.Zero(t, got.findServerObject)
	require.Zero(t, got.subordinateGuard)
	require.Zero(t, got.dropNil)
	require.Zero(t, got.dropAllEmpty)
	require.Zero(t, got.dropFoodInvalid)
	require.Zero(t, got.chestInventory)
	require.Zero(t, got.shapePoint)
	require.Equal(t, 7.5, got.shapeCircle)
	require.Equal(t, 4.0, got.shapeBoxWidth)
	require.Equal(t, 2.0, got.shapeBoxHeight)
	require.Equal(t, uint8(3), got.poisonRemaining)
	require.Equal(t, uint32(32), got.sparkFirst)
	require.Equal(t, uint32(32), got.sparkSecond)
}

func TestGame33ItemAndSpellClassification(t *testing.T) {
	got := C_game33ClassificationPaths()
	require.Zero(t, got.spellTableMissing)
	require.Zero(t, got.guideTableMissing)
	require.Zero(t, got.foodClass)
	require.Equal(t, 1, got.weaponClass)
	require.Equal(t, 1, got.creatureSpellClass)
	require.Zero(t, got.badCreatureClass)
	require.Equal(t, 1, got.firstModifier)
	require.Equal(t, 1, got.secondModifier)
	require.Equal(t, 1, got.coloredEquipment)
	require.Equal(t, 1, got.inventoryLimitNil)
	require.Equal(t, 1, got.specialSpell)
	require.Equal(t, 1, got.rangeSpell)
	require.Zero(t, got.ordinarySpell)
	require.Zero(t, got.guideAndListMissing)
	require.Zero(t, got.engageEmpty)
	require.Zero(t, got.disengageEmpty)
	require.Zero(t, got.crownPickupGuard)
	require.Zero(t, got.extraLifeGuard)
}
