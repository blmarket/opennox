package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGame33Pure covers pure helpers in GAME3_3.c.
func TestGame33Pure(t *testing.T) {
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
