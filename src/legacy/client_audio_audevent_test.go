package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientAudioEventQueue(t *testing.T) {
	t.Run("repeat limit stops an empty queue", func(t *testing.T) {
		ret, repeats := C_sub_451CF0Empty()
		require.Zero(t, ret)
		require.Equal(t, uint32(1), repeats)
	})

	t.Run("queued entry advances and is consumed", func(t *testing.T) {
		ret, remaining, index := C_sub_451CF0Queued()
		require.Equal(t, 1024, ret)
		require.Zero(t, remaining)
		require.Zero(t, index)
	})

	t.Run("active short sound remains active", func(t *testing.T) {
		require.Equal(t, 77, C_sub_451DC0Active())
	})

	t.Run("sequential selection drains in order", func(t *testing.T) {
		require.Equal(t, [3]int{0, 1, 2}, C_sub_451E80Sequence())
	})
}

func TestClientAudioEventEmptySource(t *testing.T) {
	require.Zero(t, C_sub_452580Empty())

	ret, attached := C_sub_452770NoBuffer()
	require.Zero(t, ret)
	require.Zero(t, attached)
}
