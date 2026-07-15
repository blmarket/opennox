package legacy

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame3Pure(t *testing.T) {
	t.Run("sub_4A3090 shifts array left", func(t *testing.T) {
		// count=4, arr=[10,20,30,40], idx=1 => expect [10,30,40,0xFFFFFFFF]
		ret, out := C_sub_4A3090(4, []uint32{10, 20, 30, 40}, 1)
		require.NotEqual(t, uintptr(0), ret)
		require.Equal(t, []uint32{10, 30, 40, 0xFFFFFFFF}, out)

		ret2, out2 := C_sub_4A3090(3, []uint32{1, 2, 3}, 0)
		require.Equal(t, []uint32{2, 3, 0xFFFFFFFF}, out2)
		_ = ret2
	})

	t.Run("distance and popup position helpers", func(t *testing.T) {
		require.Equal(t, 1, C_sub_4A2560(10, 20, 10, 20, 5))
		require.Zero(t, C_sub_4A2560(10, 20, 16, 20, 5))
		require.Equal(t, [2]int32{216, 27}, C_sub_4A2830(100, 20))
		require.Equal(t, [2]int32{400, 251}, C_sub_4A2830(700, 500))
		require.Equal(t, [2]int32{300, 80}, C_sub_4A2830(400, 100))
	})

	t.Run("selection and state globals", func(t *testing.T) {
		empty, indexed, missing := C_game3SelectionGlobals()
		require.True(t, empty)
		require.Equal(t, uint32(0x12345678), indexed)
		require.Zero(t, missing)
		mapState, mode := C_game3SetState(41)
		require.Equal(t, 41, mapState)
		require.Equal(t, 42, mode)
	})

	t.Run("configuration string copies", func(t *testing.T) {
		for _, tc := range []struct {
			fn  int
			off uintptr
		}{
			{0, 1308644}, {1, 1308172}, {2, 1308352}, {4, 1308324}, {5, 1308364},
		} {
			require.Zero(t, C_game3StringFunction(tc.fn, nil))
			value := "fixture-value"
			require.Equal(t, 1, C_game3StringFunction(tc.fn, &value))
			require.Equal(t, value, C_game3BlobString(tc.off))
		}
	})

	t.Run("configuration integers and pair", func(t *testing.T) {
		for _, tc := range []struct {
			fn  int
			off uintptr
		}{
			{3, 1308740}, {6, 1308188}, {8, 1308728}, {10, 1308348},
		} {
			require.Zero(t, C_game3StringFunction(tc.fn, nil))
			value := "-123"
			require.Equal(t, 1, C_game3StringFunction(tc.fn, &value))
			require.Equal(t, int32(-123), C_game3BlobInt(tc.off))
		}
		require.Zero(t, C_game3StringFunction(7, nil))
		pair := "12,34"
		require.Equal(t, 1, C_game3StringFunction(7, &pair))
		require.Equal(t, int32(12), C_game3BlobInt(1308732))
		require.Equal(t, int32(34), C_game3BlobInt(1308736))
		incomplete := "12"
		require.Zero(t, C_game3StringFunction(7, &incomplete))
		zero := "0,34"
		require.Zero(t, C_game3StringFunction(7, &zero))
	})

	t.Run("bounded string and name lookup", func(t *testing.T) {
		require.Zero(t, C_game3StringFunction(9, nil))
		short := "short value"
		require.Equal(t, 1, C_game3StringFunction(9, &short))
		require.Equal(t, short, C_game3BlobString(1308192))
		long := strings.Repeat("x", 129)
		require.Zero(t, C_game3StringFunction(9, &long))
		known, unknown, stored := C_game3NameLookup("fixture-name")
		require.Equal(t, 1, known)
		require.Equal(t, 1, unknown)
		require.Equal(t, uint32(0xaabbccdd), stored)
	})

	t.Run("small event handlers", func(t *testing.T) {
		require.Equal(t, [7]int{0, 1, 0, 1, 0, 1, 0}, C_game3SimpleEventHandlers())
	})

	t.Run("sine table and sprite bounce", func(t *testing.T) {
		h := newGameLogicHarness(t)
		h.loadBlobData()
		C_sub_4AEE30()
		setGame52Server(t, 2)
		result, position, velocity := C_sub_4B69F0(10, 3)
		require.Equal(t, int16(13), result)
		require.Equal(t, uint16(13), position)
		require.Equal(t, uint8(3), velocity)
		result, position, velocity = C_sub_4B69F0(0x7fff, 3)
		require.Equal(t, int16(13091), result)
		require.Zero(t, position)
		require.Zero(t, velocity)
	})
}
