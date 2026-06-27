package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame22(t *testing.T) {
	t.Run("sub_486640 field shift", func(t *testing.T) {
		require.Equal(t, 4660, C_sub_486640(0x12340000, 100))
		require.Equal(t, 2330, C_sub_486640(0x12340000, 50))
		require.Equal(t, 0, C_sub_486640(0, 100))
		require.Equal(t, 1, C_sub_486640(0x00010000, 100))
	})
	t.Run("sub_487C80 list next", func(t *testing.T) {
		require.Equal(t, 0, C_sub_487C80_empty())
		require.Equal(t, 1, C_sub_487C80_withNext())
	})
	t.Run("sub_480250 pack bytes", func(t *testing.T) {
		require.Equal(t, uint16(0), C_sub_480250(0, 0, 0))
		require.Equal(t, uint16(1), C_sub_480250(8, 0, 0))
		require.Equal(t, uint16(2016), C_sub_480250(0, 0xFC, 0))
		require.Equal(t, uint16(63488), C_sub_480250(0, 0, 0xF8))
		require.Equal(t, uint16(65535), C_sub_480250(0xFF, 0xFF, 0xFF))
	})
	t.Run("sub_487590 memcpy", func(t *testing.T) {
		ret, copied := C_sub_487590()
		require.NotEqual(t, 0, ret)
		var expected [28]byte
		for i := 0; i < 28; i++ {
			expected[i] = byte(i + 1)
		}
		require.Equal(t, expected, copied)
	})
	t.Run("sub_487090 list remove", func(t *testing.T) {
		require.Equal(t, 1, C_sub_487090())
	})
}
