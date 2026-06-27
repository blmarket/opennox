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
