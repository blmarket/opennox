package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestGame31FieldAccessors covers small pure pointer/arithmetic helpers in
// GAME3_1.c.
func TestGame31FieldAccessors(t *testing.T) {
	t.Run("sub_4BD680 reads field +12", func(t *testing.T) {
		require.Equal(t, int(0x12345678), C_sub_4BD680(0x12345678))
	})

	t.Run("sub_4BD710 returns ptr+24", func(t *testing.T) {
		base, ret := C_sub_4BD710()
		require.Equal(t, base+24, ret)
	})

	t.Run("sub_4BD300 links node onto free list", func(t *testing.T) {
		ret, headAfter, nodeNext, oldHead := C_sub_4BD300()
		// result == new head == node (a2-4); node's next-field == old head
		require.Equal(t, uint32(ret), headAfter)
		require.Equal(t, oldHead, nodeNext)
	})

	t.Run("sub_4BDB20 sets bit 0x10", func(t *testing.T) {
		_, flags := C_sub_4BDB20(0x01)
		require.Equal(t, uint32(0x11), flags)
		_, flags = C_sub_4BDB20(0x10) // idempotent when already set
		require.Equal(t, uint32(0x10), flags)
	})

	t.Run("sub_4BDB30 clears bit 0x10", func(t *testing.T) {
		_, flags := C_sub_4BDB30(0x11)
		require.Equal(t, uint32(0x01), flags)
		_, flags = C_sub_4BDB30(0x01) // no-op when already clear
		require.Equal(t, uint32(0x01), flags)
	})
}

func TestGame31ConstReturns(t *testing.T) {
	require.Equal(t, 1, C_nox_xxx_updDrawMonsterGen_4BC920())
	require.Equal(t, 1, C_sub_4BDDA0())
	require.Equal(t, 1, C_sub_4BE320())
	require.Equal(t, 1, C_sub_4C24A0())
	require.Equal(t, 0, C_sub_4C2BD0())
	require.Equal(t, 1, C_sub_4C2BE0())
	require.Equal(t, 1, C_nox_xxx_updDrawUndeadKiller_4CCCF0())
}
