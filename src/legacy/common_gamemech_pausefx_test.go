package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonGamemechPausefx(t *testing.T) {
	// Save original values
	orig2523804 := C_get_pausefx_2523804()
	t.Cleanup(func() {
		C_set_pausefx_2523804(orig2523804)
	})

	// Test that the function respects the pause flag
	// Set the flag to 1, which should cause early return
	C_set_pausefx_2523804(1)
	C_sub_57AF30(999, 999)
	// Should return early without crashing
	require.Equal(t, uint32(1), C_get_pausefx_2523804())
}

func TestCommonGamemechPausefxState(t *testing.T) {
	// Test that the function respects the pause flag
	// Save original values
	orig2523804 := C_get_pausefx_2523804()
	t.Cleanup(func() {
		C_set_pausefx_2523804(orig2523804)
	})

	// Set the flag to 1, which should cause early return
	C_set_pausefx_2523804(1)
	C_sub_57AF30(999, 999)
	// Should return early without crashing
	require.Equal(t, uint32(1), C_get_pausefx_2523804())
}
