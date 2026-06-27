package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame43SpellTurnUndead(t *testing.T) {
	require.Equal(t, 0, C_nox_xxx_spellTurnUndeadUpdate_531410())
}

func TestGame43Sub534020(t *testing.T) {
	require.Equal(t, 0, C_sub_534020(0))
	require.Equal(t, 1, C_sub_534020(1<<10))
	require.Equal(t, 0, C_sub_534020(1<<9))
}

func TestGame43MonsterMoving(t *testing.T) {
	require.Equal(t, 0, C_nox_xxx_monsterIsMoveing_534320(0))
	require.Equal(t, 0, C_nox_xxx_monsterIsMoveing_534320(0.009))
	require.Equal(t, 1, C_nox_xxx_monsterIsMoveing_534320(0.02))
	require.Equal(t, 1, C_nox_xxx_monsterIsMoveing_534320(1))
}

func TestGame43Sub534440(t *testing.T) {
	require.Equal(t, 1, C_sub_534440(0))
	require.Equal(t, 1, C_sub_534440(0.07))
	require.Equal(t, 0, C_sub_534440(0.08))
	require.Equal(t, 0, C_sub_534440(1))
}

func TestGame43Sub534470(t *testing.T) {
	require.InDelta(t, 2.5, C_sub_534470(2.5), 0.0001)
	require.InDelta(t, -1.25, C_sub_534470(-1.25), 0.0001)
}

func TestGame43Sub5347A0(t *testing.T) {
	require.Equal(t, 0, C_sub_5347A0(0))
	require.Equal(t, 1, C_sub_5347A0(1<<9))
	require.Equal(t, 0, C_sub_5347A0(1<<8))
}

func TestGame43IsNotPoisoned(t *testing.T) {
	require.Equal(t, 0, C_nox_xxx_isNotPoisoned_5347F0(0))
	require.Equal(t, 1, C_nox_xxx_isNotPoisoned_5347F0(1))
	require.Equal(t, 1, C_nox_xxx_isNotPoisoned_5347F0(255))
}

func TestGame43Sub537580(t *testing.T) {
	require.Equal(t, 0, C_sub_537580(0))
	require.Equal(t, 1, C_sub_537580(1))
	require.Equal(t, 0, C_sub_537580(2))
	require.Equal(t, 1, C_sub_537580(3))
}

func TestGame43Sub540D20(t *testing.T) {
	require.Equal(t, 0, C_sub_540D20(74))
	require.Equal(t, 1, C_sub_540D20(75))
	require.Equal(t, 1, C_sub_540D20(114))
	require.Equal(t, 0, C_sub_540D20(115))
}

func TestGame43Sub534750(t *testing.T) {
	ret, after := C_sub_534750(0)
	require.Equal(t, 0x4000, ret)
	require.Equal(t, uint32(0x4000), after)

	ret, after = C_sub_534750(0x10000)
	require.Equal(t, 0x10000, ret)
	require.Equal(t, uint32(0x10000), after)
}

func TestGame43Sub534780(t *testing.T) {
	ret, after := C_sub_534780(0x8000)
	require.Equal(t, 0x8000, ret)
	require.Equal(t, uint32(0x8000), after)

	ret, after = C_sub_534780(0)
	require.Equal(t, 0, ret)
	require.Equal(t, uint32(0), after)
}

func TestGame43MonsterCanAttack(t *testing.T) {
	require.Equal(t, 0, C_nox_xxx_monsterCanAttackAtWill_534390(0.66))
	require.Equal(t, 1, C_nox_xxx_monsterCanAttackAtWill_534390(0.67))
	require.Equal(t, 1, C_nox_xxx_monsterCanAttackAtWill_534390(1))
}
