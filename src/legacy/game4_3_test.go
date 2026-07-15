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

func TestGame43MonsterStatePredicates(t *testing.T) {
	require.Equal(t, 0, C_nox_xxx_monsterCanCast_534300(0))
	require.Equal(t, 1, C_nox_xxx_monsterCanCast_534300(1<<5))
	require.Equal(t, 0, C_nox_xxx_monsterCanCast_534300(1<<4))

	for _, action := range []int8{0, 1, 4, 23, 25, 26, 27} {
		require.Equal(t, 1, C_sub_534340(action), "action %d", action)
	}
	require.Equal(t, 0, C_sub_534340(2))
	require.Equal(t, 0, C_sub_534340(24))

	require.Equal(t, 0, C_sub_5343C0(0.32))
	require.Equal(t, 1, C_sub_5343C0(0.34))
	require.Equal(t, 1, C_sub_5343C0(0.65))
	require.Equal(t, 0, C_sub_5343C0(0.67))

	require.Equal(t, 0, C_sub_534400(0.07))
	require.Equal(t, 1, C_sub_534400(0.09))
	require.Equal(t, 1, C_sub_534400(0.32))
	require.Equal(t, 0, C_sub_534400(0.33))

	for _, action := range []int8{18, 19, 20} {
		require.Equal(t, 1, C_sub_5408A0(action), "action %d", action)
	}
	require.Equal(t, 0, C_sub_5408A0(17))
	require.Equal(t, 0, C_sub_5408A0(21))
}

func TestGame43TextParsers(t *testing.T) {
	ret, floats := C_sub_536550("1.25 -2.5")
	require.Equal(t, 1, ret)
	require.Equal(t, [3]float32{1.25, 1.25, -2.5}, floats)

	ret, ints := C_sub_536580("17 -23 4096")
	require.Equal(t, 1, ret)
	require.Equal(t, [3]int32{17, -23, 4096}, ints)

	ret, value := C_sub_536600("-99")
	require.Equal(t, 1, ret)
	require.Equal(t, -99, value)

	ret, value = C_sub_536D80("12345")
	require.Equal(t, 1, ret)
	require.Equal(t, 12345, value)

	ret, byteValue := C_sub_536DE0("300")
	require.Equal(t, 1, ret)
	require.Equal(t, uint8(44), byteValue)

	ret, byteValue = C_sub_536E50("255 trailing")
	require.Equal(t, 1, ret)
	require.Equal(t, uint8(255), byteValue)
}

func TestGame43PointerAndActionAliases(t *testing.T) {
	require.Equal(t, 0, C_sub_537750(123, true))
	require.Equal(t, 123, C_sub_537750(123, false))

	set, clear := C_monsterActionFlagAliases(0)
	require.Equal(t, [2]uint32{0x4000, 0x4000}, set)
	require.Equal(t, [3]uint32{}, clear)

	set, clear = C_monsterActionFlagAliases(0x4000)
	require.Equal(t, [2]uint32{0x4000, 0x4000}, set)
	require.Equal(t, [3]uint32{}, clear)

	set, clear = C_monsterActionFlagAliases(0xC000)
	require.Equal(t, [2]uint32{0xC000, 0xC000}, set)
	require.Equal(t, [3]uint32{0xC000, 0xC000, 0xC000}, clear)
}
