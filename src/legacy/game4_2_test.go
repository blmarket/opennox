package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame42Pure(t *testing.T) {
	t.Run("sub_521EB0 float rect overlap check", func(t *testing.T) {
		// a1 defines rect via indices 9,10,11,12 ; a2 defines point rect via 1,2,3,4
		// choose a1 = [0..12] with 9:0,10:0,11:10,12:10 => rect 0,0 to 10,10
		// a2 with 1:5,2:5,3:5,4:5 => center 5,5 => should pass
		a1 := make([]float32, 13)
		a1[9] = 0
		a1[10] = 0
		a1[11] = 10
		a1[12] = 10
		a2 := make([]float32, 5)
		a2[1] = 5
		a2[2] = 5
		a2[3] = 5
		a2[4] = 5
		require.Equal(t, 1, C_sub_521EB0(a1, a2))

		// outside to left
		a2[1] = -1
		a2[2] = 5
		a2[3] = -1
		a2[4] = 5
		require.Equal(t, 0, C_sub_521EB0(a1, a2))

		// outside to right
		a2[1] = 11
		a2[3] = 11
		require.Equal(t, 0, C_sub_521EB0(a1, a2))
	})

	t.Run("nox_xxx_mapGenCheckRoomType_5238F0", func(t *testing.T) {
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(2))
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(3))
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(4))
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(5))
		require.Equal(t, 0, C_nox_xxx_mapGenCheckRoomType_5238F0(1))
		require.Equal(t, 0, C_nox_xxx_mapGenCheckRoomType_5238F0(0))
		require.Equal(t, 0, C_nox_xxx_mapGenCheckRoomType_5238F0(6))
	})

	t.Run("mapGen rng helpers", func(t *testing.T) {
		// Seeding makes nox_platform_rand deterministic; both rand funcs must
		// stay within their documented [a1, a2] / [a1-a2, a1+a2] bounds.
		C_nox_xxx_mapGenSetRngSeed_526AB0(12345)

		// result = a1 + (a2-a1+1)*rand()/0x7FFF, clamped to a2.
		for i := 0; i < 200; i++ {
			r := C_nox_xxx_mapGenRandFunc_526AC0(10, 20)
			require.GreaterOrEqual(t, r, 10)
			require.LessOrEqual(t, r, 20)
		}
		// Degenerate range a1==a2 always returns a1.
		for i := 0; i < 20; i++ {
			require.Equal(t, 7, C_nox_xxx_mapGenRandFunc_526AC0(7, 7))
		}
	})

	t.Run("mapGenRandFunc2", func(t *testing.T) {
		C_nox_xxx_mapGenSetRngSeed_526AB0(999)
		// a2==0 short-circuits and returns a1 unchanged.
		require.Equal(t, 42, C_nox_xxx_mapGenRandFunc2_526B00(42, 0))
		// Otherwise result = v4 + a1 with v4 in [-a2, a2].
		for i := 0; i < 100; i++ {
			r := C_nox_xxx_mapGenRandFunc2_526B00(100, 5)
			require.GreaterOrEqual(t, r, 95)
			require.LessOrEqual(t, r, 105)
		}
	})

	t.Run("nox_xxx_isObjectMovable_52E020", func(t *testing.T) {
		// field16 & 0x8068 != 0  => not movable (0)
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0, 0x8000))
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0, 0x0008))
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0, 0x0060))
		// else result = (~field8 >> 22) & 1
		// field8 bit22 clear => ~ has bit22 set => 1
		require.Equal(t, 1, C_nox_xxx_isObjectMovable_52E020(0, 0))
		// field8 bit22 set => ~ clears it => 0
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0x00400000, 0))
		// field16 bit not in mask (0x10) => still falls to else branch
		require.Equal(t, 1, C_nox_xxx_isObjectMovable_52E020(0, 0x10))
	})

	t.Run("sub_526AA0 stride", func(t *testing.T) {
		// returns base + (a1 << 6); successive indices differ by 64 bytes.
		base := C_sub_526AA0(0)
		require.Equal(t, uintptr(64), C_sub_526AA0(1)-base)
		require.Equal(t, uintptr(640), C_sub_526AA0(10)-base)
	})
}
