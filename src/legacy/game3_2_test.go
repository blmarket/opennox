package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame32Const(t *testing.T) {
	require.Equal(t, 0, C_sub_4E14A0())
	require.Equal(t, 1, C_sub_479D00())
}

func TestGame32SafeEmptyAndGeometryPaths(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	h.setServer(0, 30)
	got := C_game32SafePaths()
	require.Zero(t, got.listContains)
	require.Zero(t, got.listRemoved)
	require.Equal(t, [8]uint32{4, 1, 1, 4, 7, 4, 4, 7}, got.sorted)
	require.Zero(t, got.transformNil)
	require.InDelta(t, -4080.7, got.transform[0], 0.1)
	require.InDelta(t, 16.26, got.transform[1], 0.1)
	require.Equal(t, []int{-1, 3, 0, 1, 6, -1, 2, 7, 8, 5, -1}, got.directions)
	require.Equal(t, uint32(10), got.nameLen)
	require.Equal(t, "Synthetic", got.name)
	require.Equal(t, 1, got.cloud)
	require.Equal(t, 1, got.cloudAlt)
	require.Zero(t, got.gripNil)
	require.Zero(t, got.specialNil)
	require.Zero(t, got.specialFalse)
	require.Equal(t, 1, got.specialTrue)
}

func TestGame32Arith4866D0(t *testing.T) {
	require.Equal(t, 10, C_sub_4866D0(10, 0))
	require.Equal(t, 46, C_sub_4866D0(10, 1))
	require.Equal(t, 82, C_sub_4866D0(10, 2))
}

func TestGame32MapFlags(t *testing.T) {
	// nox_mapToGameFlags mapping: bit1->0x200, 2->0x1000, 4->0x100, 8->0x20, 16->0x10, 32->0x400, 64->0x40, negative->0x80
	check := func(v uint32, expected int) {
		require.Equal(t, expected, C_nox_xxx_mapGetTypeMB_4CFFA0(v))
		require.Equal(t, expected, C_sub_4CFFC0(v))
	}
	check(0, 0)
	check(1, 0x200)
	check(2, 0x1000)
	check(4, 0x100)
	check(8, 0x20)
	check(16, 0x10)
	check(32, 0x400)
	check(64, 0x40)
	check(3, 0x200+0x1000)
	check(0x7f, 0x200+0x1000+0x100+0x20+0x10+0x400+0x40)
	// negative as uint32 max
	check(0xffffffff, 0x17f0)
}

func TestGame32CE340(t *testing.T) {
	ret, after := C_sub_4CE340(10, 5)
	require.Equal(t, 1, ret)
	require.Equal(t, uint16(15), after)

	ret, after = C_sub_4CE340(0, 255)
	require.Equal(t, 1, ret)
	require.Equal(t, uint16(255), after)

	ret, after = C_sub_4CE340(1000, 0)
	require.Equal(t, 1, ret)
	require.Equal(t, uint16(1000), after)
}

func TestGame32EffectDamageMultiplier(t *testing.T) {
	require.InDelta(t, 5.0, C_nox_xxx_effectDamageMultiplier_4E04C0(2.5, 2.0), 0.0001)
	require.InDelta(t, 0.0, C_nox_xxx_effectDamageMultiplier_4E04C0(10, 0), 0.0001)
	require.InDelta(t, -3.0, C_nox_xxx_effectDamageMultiplier_4E04C0(1.5, -2), 0.0001)
}

func TestGame32EffectProjectileSpeed(t *testing.T) {
	require.InDelta(t, 6.0, C_nox_xxx_effectProjectileSpeed_4E09B0(3, 2), 0.0001)
	require.InDelta(t, 0.0, C_nox_xxx_effectProjectileSpeed_4E09B0(5, 0), 0.0001)
	require.InDelta(t, 2.5, C_nox_xxx_effectProjectileSpeed_4E09B0(0.5, 5), 0.0001)
}

func TestGame32InversionGrip(t *testing.T) {
	ret, out := C_nox_xxx_inversionEffect_4E03D0(0)
	require.Equal(t, 0, ret)
	require.Equal(t, 0, out)

	ret, out = C_nox_xxx_inversionEffect_4E03D0(1)
	require.Equal(t, 1, ret)
	require.Equal(t, 1, out)

	ret, out = C_nox_xxx_inversionEffect_4E03D0(5)
	require.Equal(t, 1, ret)
	require.Equal(t, 1, out)

	ret, out = C_nox_xxx_gripEffect_4E0480(0)
	require.Equal(t, 1, ret)
	require.Equal(t, 1, out)

	ret, out = C_nox_xxx_gripEffect_4E0480(1)
	require.Equal(t, 0, ret)
	require.Equal(t, 0, out)

	ret, out = C_nox_xxx_gripEffect_4E0480(10)
	require.Equal(t, 0, ret)
	require.Equal(t, 0, out)
}
