package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame23Pure(t *testing.T) {
	t.Run("sub_4A0020 returns non-nil global address", func(t *testing.T) {
		addr := C_sub_4A0020()
		require.NotEqual(t, uintptr(0), addr)
		// calling twice should return same address
		require.Equal(t, addr, C_sub_4A0020())
	})
}

func TestSub48D4F0(t *testing.T) {
	// sub_48D4F0 checks if a2 is within a range of a1
	// v2 = 10000; if (a1 - 10000 < 0) { ... } return a2 < a1 && a2 >= a1 - v2;

	// When a1 >= 10000, v2 stays 10000, return a2 < a1 && a2 >= a1 - 10000
	require.Equal(t, 1, C_sub_48D4F0(15000, 12000)) // 12000 < 15000 && 12000 >= 5000
	require.Equal(t, 0, C_sub_48D4F0(15000, 16000)) // 16000 < 15000 false
	require.Equal(t, 0, C_sub_48D4F0(15000, 4000))  // 4000 >= 5000 false

	// When a1 < 10000, v2 = a1, return a2 < a1 && a2 >= 0 (since a1 - v2 = 0)
	// Also checks if a2 >= 0xFFFF - (10000 - a1), if so return 1
	require.Equal(t, 1, C_sub_48D4F0(5000, 3000)) // 3000 < 5000 && 3000 >= 0
	require.Equal(t, 0, C_sub_48D4F0(5000, 6000)) // 6000 < 5000 false
	require.Equal(t, 1, C_sub_48D4F0(5000, 0))    // 0 < 5000 && 0 >= 0 true

	// Edge case: a2 >= 0xFFFF - (10000 - a1) returns 1
	// For a1=5000, 10000-5000=5000, 0xFFFF-5000=60535, so if a2 >= 60535 return 1
	require.Equal(t, 1, C_sub_48D4F0(5000, 60535))
	require.Equal(t, 1, C_sub_48D4F0(5000, 65535))
}

func TestSub48C580(t *testing.T) {
	// sub_48C580 sorts pixels in ascending order (smallest first)
	// It uses a selection sort-like algorithm with atomic exchanges

	// Test with simple array
	pixels := []uint32{3, 1, 4, 1, 5, 9, 2, 6}
	C_sub_48C580(pixels)
	// After sorting, should be in ascending order: 1, 1, 2, 3, 4, 5, 6, 9
	require.Equal(t, uint32(1), pixels[0])
	require.Equal(t, uint32(1), pixels[1])
	require.Equal(t, uint32(2), pixels[2])
	require.Equal(t, uint32(3), pixels[3])
	require.Equal(t, uint32(4), pixels[4])
	require.Equal(t, uint32(5), pixels[5])
	require.Equal(t, uint32(6), pixels[6])
	require.Equal(t, uint32(9), pixels[7])

	// Test with already sorted array (ascending)
	pixels2 := []uint32{1, 2, 3, 4, 5}
	C_sub_48C580(pixels2)
	require.Equal(t, []uint32{1, 2, 3, 4, 5}, pixels2)

	// Test with single element
	pixels3 := []uint32{42}
	C_sub_48C580(pixels3)
	require.Equal(t, []uint32{42}, pixels3)

	// Test with empty array (should not crash)
	pixels4 := []uint32{}
	C_sub_48C580(pixels4)
	require.Equal(t, []uint32{}, pixels4)

	// Test with reverse sorted array
	pixels5 := []uint32{5, 4, 3, 2, 1}
	C_sub_48C580(pixels5)
	require.Equal(t, []uint32{1, 2, 3, 4, 5}, pixels5)
}
