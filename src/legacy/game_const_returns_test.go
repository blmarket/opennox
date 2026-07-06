package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Characterize the tiny decompiled stubs that ignore their arguments and always
// return the same constant. Covering them is cheap and pins the return value in
// case a future re-decompile changes the body.
func TestGameConstReturns(t *testing.T) {
	require.Equal(t, 0, C_nox_xxx_wndRetNULL_46A8A0())
	require.Equal(t, 0, C_nox_xxx_wndRetNULL_0_46A8B0())

	// book_45BD30 returns 1 for any inputs.
	require.Equal(t, 1, C_nox_xxx_book_45BD30(0, 0))
	require.Equal(t, 1, C_nox_xxx_book_45BD30(42, -7))
}
