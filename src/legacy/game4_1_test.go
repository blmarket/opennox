package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame41Pure(t *testing.T) {
	t.Run("nox_xxx_utilNormalizeVector_509F20", func(t *testing.T) {
		x, y := C_nox_xxx_utilNormalizeVector_509F20(3, 4)
		require.InDelta(t, 0.6, x, 0.0001)
		require.InDelta(t, 0.8, y, 0.0001)

		x2, y2 := C_nox_xxx_utilNormalizeVector_509F20(0, 5)
		require.InDelta(t, 0, x2, 0.0001)
		require.InDelta(t, 1, y2, 0.0001)
	})

	t.Run("nox_xxx_mobActionGet_50A020", func(t *testing.T) {
		require.Equal(t, 0x12345678, C_nox_xxx_mobActionGet_50A020())
	})

	t.Run("nox_xxx_monsterActionIsCondition_50A010", func(t *testing.T) {
		require.Equal(t, 0, C_nox_xxx_monsterActionIsCondition_50A010(0))
		require.Equal(t, 0, C_nox_xxx_monsterActionIsCondition_50A010(38))
		require.Equal(t, 0, C_nox_xxx_monsterActionIsCondition_50A010(39))
		require.Equal(t, 1, C_nox_xxx_monsterActionIsCondition_50A010(40))
		require.Equal(t, 1, C_nox_xxx_monsterActionIsCondition_50A010(100))
	})

	t.Run("sub_509FF0 clears when flag 0x20 set", func(t *testing.T) {
		ret, out := C_sub_509FF0(0x20)
		require.NotEqual(t, 0, ret) // returns a1 pointer as int, non-zero
		require.Equal(t, uint32(0), out)

		ret2, out2 := C_sub_509FF0(0)
		require.NotEqual(t, 0, ret2)
		// when flag not set, *buf should remain pointer value (non-zero, not 0)
		require.NotEqual(t, uint32(0), out2)
	})
}
