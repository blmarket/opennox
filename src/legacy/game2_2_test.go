package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame22(t *testing.T) {
	t.Run("safe empty state paths", func(t *testing.T) {
		got := C_game22SafeStatePaths()
		require.Equal(t, ^uint32(0), got.emptySearch)
		require.NotZero(t, got.recordStride)
		require.Equal(t, 2, got.normalFormat)
		require.Equal(t, uint32(2), got.normalChannels)
		require.Equal(t, uint32(2), got.normalMode)
		require.Equal(t, 2, got.specialFormat)
		require.Equal(t, uint32(2), got.specialInput)
		require.Equal(t, uint32(2), got.specialMode)
		require.Zero(t, got.emptyRead)
		require.Zero(t, got.closeEmpty)
		require.Zero(t, got.nextEmpty)
		require.Zero(t, got.firstEmpty)
		require.Zero(t, got.selectEmpty)
		require.Zero(t, got.removeEmpty)
		require.Zero(t, got.releaseEmpty)
		require.True(t, got.listInitialized)
		require.Equal(t, 24, got.bufferSize)
		require.Equal(t, 6, got.quarterSize)
		require.Equal(t, uint32(7), got.configuredStride)
		require.Equal(t, uint32(9), got.configuredWidth)
		require.Zero(t, got.clearedSize)
	})
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
	t.Run("sub_481410 sets waypoint counter", func(t *testing.T) {
		old := C_game2_2_getWaypointCounter()
		t.Cleanup(func() { C_game2_2_setWaypointCounter(old) })

		C_game2_2_setWaypointCounter(0)
		C_sub_481410()
		require.Equal(t, uint32(0xFFFFFFFF), C_game2_2_getWaypointCounter())

		C_game2_2_setWaypointCounter(123)
		C_sub_481410()
		require.Equal(t, uint32(0xFFFFFFFF), C_game2_2_getWaypointCounter())
	})
	t.Run("sub_484450 initializes the full light buffer", func(t *testing.T) {
		got := C_game22InitializeLightBuffer(0x1234)
		require.Equal(t, 0x1234, got.ret)
		require.Equal(t, uint16(0x1234), got.first)
		require.Equal(t, uint32(0x1234), got.second)
		require.Equal(t, uint32(0x1234), got.middle)
		require.Equal(t, uint16(0x1234), got.penultimate)
		require.Equal(t, uint16(0x1234), got.last)
		require.Zero(t, got.untouchedGap)
	})
}
