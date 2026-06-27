package legacy

import (
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
}
