package sound

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSoundExtra(t *testing.T) {
	s := ByName("NONEXISTENT")
	require.Equal(t, SoundNone, s)
	s2 := ByName("AnchorCast")
	require.Equal(t, SoundAnchorCast, s2)
	require.NotEmpty(t, s2.GoString())
	require.NotEmpty(t, s2.String())
}
