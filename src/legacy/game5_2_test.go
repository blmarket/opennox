package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame52NetworkCodeHelpers(t *testing.T) {
	t.Run("net get unit code handles nil and invalid code", func(t *testing.T) {
		require.Equal(t, uint32(0), C_nox_xxx_netGetUnitCodeCli_578B00_nil())
		require.Equal(t, uint32(0), C_nox_xxx_netGetUnitCodeCli_578B00(0x8000, 0))
	})

	t.Run("net get unit code marks dynamic extent units", func(t *testing.T) {
		require.Equal(t, uint32(0x1234), C_nox_xxx_netGetUnitCodeCli_578B00(0x1234, 0))
		require.Equal(t, uint32(0x9234), C_nox_xxx_netGetUnitCodeCli_578B00(0x1234, 0x20400000))
		require.Equal(t, uint32(0x8012), C_nox_xxx_netGetUnitCodeCli_578B00(0x12, 0x400000))
	})

	t.Run("clear high bit masks bit 15", func(t *testing.T) {
		require.Equal(t, 0, C_nox_xxx_netClearHighBit_578B30(0))
		require.Equal(t, 0x1234, C_nox_xxx_netClearHighBit_578B30(0x1234))
		require.Equal(t, 0, C_nox_xxx_netClearHighBit_578B30(-0x8000))
		require.Equal(t, 0x7FFF, C_nox_xxx_netClearHighBit_578B30(-1))
	})

	t.Run("test high bit reads bit 15 only", func(t *testing.T) {
		require.Equal(t, uint32(0), C_nox_xxx_netTestHighBit_578B70(0))
		require.Equal(t, uint32(0), C_nox_xxx_netTestHighBit_578B70(0x7FFF))
		require.Equal(t, uint32(1), C_nox_xxx_netTestHighBit_578B70(0x8000))
		require.Equal(t, uint32(0), C_nox_xxx_netTestHighBit_578B70(0x10000))
		require.Equal(t, uint32(1), C_nox_xxx_netTestHighBit_578B70(0x18000))
	})
}

func TestGame52OffsetAccessors(t *testing.T) {
	t.Run("waypoint next reads offset 484", func(t *testing.T) {
		require.Equal(t, 0, C_nox_xxx_waypointNext_579870_nil())
		require.Equal(t, 0x11223344, C_nox_xxx_waypointNext_579870(0x11223344))
	})

	t.Run("sub 5798A0 reads offset 484", func(t *testing.T) {
		require.Equal(t, 0, C_sub_5798A0_nil())
		require.Equal(t, 0x55667788, C_sub_5798A0(0x55667788))
	})

	t.Run("sub 57BA10 writes packed fields", func(t *testing.T) {
		ret, out0, out2, out4 := C_sub_57BA10(-2, 0x1234, 0x89ABCDEF)
		require.NotZero(t, ret)
		require.Equal(t, uint16(0xFFFE), out0)
		require.Equal(t, uint16(0x1234), out2)
		require.Equal(t, uint32(0x89ABCDEF), out4)
	})

	t.Run("next map group reads offset 88", func(t *testing.T) {
		require.Equal(t, 0, C_nox_server_getNextMapGroup_57C090_nil())
		require.Equal(t, 0x01020304, C_nox_server_getNextMapGroup_57C090(0x01020304))
	})
}
