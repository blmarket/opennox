package noxflags

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameplay_Extra(t *testing.T) {
	// Reset
	gameplay = 0

	require.Equal(t, GameplayFlag(0), GetGamePlay())
	require.False(t, HasGamePlay(GameplayFlag1))

	SetGamePlay(GameplayFlag1)
	require.True(t, HasGamePlay(GameplayFlag1))
	require.Equal(t, GameplayFlag1, GetGamePlay())

	SetGamePlay(GameplayFlag2)
	require.True(t, HasGamePlay(GameplayFlag1))
	require.True(t, HasGamePlay(GameplayFlag2))
	require.Equal(t, GameplayFlag1|GameplayFlag2, GetGamePlay())

	SetGamePlay(GameplayFlag4)
	require.True(t, HasGamePlay(GameplayFlag4))

	UnsetGamePlay(GameplayFlag2)
	require.False(t, HasGamePlay(GameplayFlag2))
	require.True(t, HasGamePlay(GameplayFlag1))
	require.True(t, HasGamePlay(GameplayFlag4))

	UnsetGamePlay(GameplayFlag1 | GameplayFlag4)
	require.False(t, HasGamePlay(GameplayFlag1))
	require.False(t, HasGamePlay(GameplayFlag4))
	require.Equal(t, GameplayFlag(0), GetGamePlay())

	// Test Has with multiple flags
	SetGamePlay(GameplayFlag1 | GameplayFlag2)
	require.True(t, HasGamePlay(GameplayFlag1|GameplayFlag2))
	require.True(t, HasGamePlay(GameplayFlag1))
	require.False(t, HasGamePlay(GameplayFlag4))

	// Reset
	gameplay = 0
}
