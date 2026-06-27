package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGame11NumericConversions covers the pure float/double conversion helpers
// in GAME1_1.c (truncation toward zero, abs, narrowing).
func TestGame11NumericConversions(t *testing.T) {
	t.Run("nox_float2int16 truncates toward zero", func(t *testing.T) {
		require.Equal(t, int16(3), C_nox_float2int16(3.9))
		require.Equal(t, int16(-3), C_nox_float2int16(-3.9))
		require.Equal(t, int16(0), C_nox_float2int16(0))
	})
	t.Run("nox_float2int16_abs takes magnitude", func(t *testing.T) {
		require.Equal(t, int16(3), C_nox_float2int16_abs(-3.9))
		require.Equal(t, int16(7), C_nox_float2int16_abs(7.2))
	})
	t.Run("nox_double2float narrows", func(t *testing.T) {
		require.Equal(t, float32(1.5), C_nox_double2float(1.5))
		require.Equal(t, float32(-2.25), C_nox_double2float(-2.25))
	})
	t.Run("nox_double2int truncates toward zero", func(t *testing.T) {
		require.Equal(t, 2, C_nox_double2int(2.9))
		require.Equal(t, -2, C_nox_double2int(-2.9))
	})
}

// TestGame11ListHelpers covers pure list manipulation helpers in GAME1_1.c
// that operate solely on caller-supplied buffers via raw offsets.
func TestGame11ListHelpers(t *testing.T) {
	t.Run("nox_common_list_clear_425760 sets self pointers", func(t *testing.T) {
		base, f0, f1, f2 := C_nox_common_list_clear_425760()
		require.Equal(t, base, f0)
		require.Equal(t, base, f1)
		require.Equal(t, base, f2)
	})
	t.Run("sub_425770 initializes node", func(t *testing.T) {
		base, f0, f1, f2, ret := C_sub_425770()
		require.Equal(t, base, ret)
		require.Equal(t, base, f0)
		require.Equal(t, base, f1)
		require.Equal(t, uint32(0), f2)
	})
	t.Run("nox_common_list_getNextSafe_4258A0 nil returns 0", func(t *testing.T) {
		require.Equal(t, 0, C_nox_common_list_getNextSafe_4258A0_nil())
	})
	t.Run("nox_common_list_getNextSafe_4258A0 empty returns 0", func(t *testing.T) {
		_, ret := C_nox_common_list_getNextSafe_4258A0_empty()
		require.Equal(t, 0, ret)
	})
	t.Run("nox_common_list_getNextSafe_4258A0 returns next", func(t *testing.T) {
		_, node, ret := C_nox_common_list_getNextSafe_4258A0_next(false)
		require.Equal(t, node, ret)
	})
	t.Run("nox_common_list_getNextSafe_4258A0 sentinel returns 0", func(t *testing.T) {
		_, _, ret := C_nox_common_list_getNextSafe_4258A0_next(true)
		require.Equal(t, 0, ret)
	})
	t.Run("nox_common_list_getFirstSafe_425890 nil", func(t *testing.T) {
		require.Equal(t, 0, C_nox_common_list_getFirstSafe_425890_nil())
	})
	t.Run("nox_common_list_getFirstSafe_425890 empty", func(t *testing.T) {
		_, ret := C_nox_common_list_getFirstSafe_425890_empty()
		require.Equal(t, 0, ret)
	})
	t.Run("nox_common_list_getFirstSafe_425890 next", func(t *testing.T) {
		_, node, ret := C_nox_common_list_getFirstSafe_425890_next(false)
		require.Equal(t, node, ret)
	})
	t.Run("sub_425A60 nil and next", func(t *testing.T) {
		require.Equal(t, 0, C_sub_425A60_nil())
		_, node, ret := C_sub_425A60_next(false)
		require.Equal(t, node, ret)
		_, _, ret2 := C_sub_425A60_next(true)
		require.Equal(t, 0, ret2)
	})
	t.Run("sub_425BE0 nil and next", func(t *testing.T) {
		require.Equal(t, 0, C_sub_425BE0_nil())
		_, node, ret := C_sub_425BE0_next(false)
		require.Equal(t, node, ret)
	})
	t.Run("sub_425BC0 empty and next", func(t *testing.T) {
		_, ret := C_sub_425BC0_empty()
		require.Equal(t, 0, ret)
		_, node, ret2 := C_sub_425BC0_next(false)
		require.Equal(t, node, ret2)
		_, _, ret3 := C_sub_425BC0_next(true)
		require.Equal(t, 0, ret3)
	})
	t.Run("sub_425960 compares pointer field", func(t *testing.T) {
		retNe := C_sub_425960(false)
		require.NotEqual(t, 0, retNe)
		retEq := C_sub_425960(true)
		require.Equal(t, 0, retEq)
	})
	t.Run("nox_common_list_remove_425920 rewires neighbors", func(t *testing.T) {
		node, prev, next, nodeF0, nodeF1, prevNext, nextPrev := C_nox_common_list_remove_425920()
		require.Equal(t, node, nodeF0)
		require.Equal(t, node, nodeF1)
		require.Equal(t, next, prevNext)
		require.Equal(t, prev, nextPrev)
	})
}
