package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame12(t *testing.T) {
	t.Run("sub_43AF40 returns global", func(t *testing.T) {
		old := C_game1_2_getCreateOrJoin()
		t.Cleanup(func() { C_game1_2_setCreateOrJoin(old) })

		C_game1_2_setCreateOrJoin(0)
		require.Equal(t, 0, C_sub_43AF40())

		C_game1_2_setCreateOrJoin(1)
		require.Equal(t, 1, C_sub_43AF40())

		C_game1_2_setCreateOrJoin(42)
		require.Equal(t, 42, C_sub_43AF40())
	})
}
