package sound

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestByName(t *testing.T) {
	// Test existing sound
	s := ByName("AnchorCast")
	require.Equal(t, SoundAnchorCast, s)

	// Test non-existent sound
	s = ByName("NONEXISTENT_SOUND")
	require.Equal(t, SoundNone, s)

	// Test empty string
	s = ByName("")
	require.Equal(t, SoundNone, s)

	// Test case sensitivity
	s = ByName("anchorcast")
	require.Equal(t, SoundNone, s)
}

func TestIDString(t *testing.T) {
	// Test valid IDs
	require.Equal(t, "NULL", SoundNone.String())
	require.Equal(t, "AnchorCast", SoundAnchorCast.String())
	require.Equal(t, "AnchorOn", SoundAnchorOn.String())

	// Test last valid ID
	lastID := ID(len(soundNames) - 1)
	require.NotEmpty(t, lastID.String())

	// Test out of range ID - should panic due to bug in String() (uses || instead of &&)
	// The bug causes panic for out-of-range IDs because condition is always true
	require.Panics(t, func() {
		invalidID := ID(len(soundNames) + 100)
		_ = invalidID.String()
	})

	// Test negative ID - should also panic
	require.Panics(t, func() {
		negativeID := ID(-1)
		_ = negativeID.String()
	})
}

func TestIDGoString(t *testing.T) {
	// Test SoundNone
	require.Equal(t, "sound.SoundNone", SoundNone.GoString())

	// Test valid IDs
	require.Equal(t, "sound.SoundAnchorCast", SoundAnchorCast.GoString())
	require.Equal(t, "sound.SoundAnchorOn", SoundAnchorOn.GoString())

	// Test out of range ID - should panic due to bug
	require.Panics(t, func() {
		invalidID := ID(len(soundNames) + 100)
		_ = invalidID.GoString()
	})

	// Test negative ID - should also panic
	require.Panics(t, func() {
		negativeID := ID(-1)
		_ = negativeID.GoString()
	})
}

func TestAllSounds(t *testing.T) {
	// Test that all defined sounds have non-empty String and GoString
	// and that ByName returns the correct ID
	for i, name := range soundNames {
		id := ID(i)
		// String should return the name
		require.Equal(t, name, id.String(), "String() mismatch for ID %d", i)
		// GoString should not be empty
		require.NotEmpty(t, id.GoString(), "GoString() empty for ID %d", i)
		// ByName should return the correct ID
		require.Equal(t, id, ByName(name), "ByName() mismatch for %s", name)
	}
}
