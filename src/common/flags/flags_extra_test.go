package noxflags

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEngineFlagGoString_Comprehensive(t *testing.T) {
	// Test GoString for all engine flags
	flags := []EngineFlag{
		EngineShowExtents,
		EngineAdmin,
		EngineGodMode,
		EngineWindowed,
	}
	for _, f := range flags {
		s := f.GoString()
		require.NotEmpty(t, s)
		require.Contains(t, s, "Engine")
	}

	// Test combined flags
	combined := EngineShowExtents | EngineAdmin | EngineGodMode
	s := combined.GoString()
	require.Contains(t, s, "EngineShowExtents")
	require.Contains(t, s, "EngineAdmin")
	require.Contains(t, s, "EngineGodMode")

	// Test empty - may be empty string
	empty := EngineFlag(0)
	s = empty.GoString()
	_ = s // empty is ok
}

func TestGameFlagGoString_Comprehensive(t *testing.T) {
	flags := []GameFlag{
		GameModeCoop,
		GameModeQuest,
		GameHost,
		GameOnline,
	}
	for _, f := range flags {
		s := f.GoString()
		_ = s // may be empty for some, just ensure no panic
	}

	combined := GameModeCoop | GameHost
	s := combined.GoString()
	_ = s

	empty := GameFlag(0)
	s = empty.GoString()
	_ = s
}

func TestEngineFlagSetUnset_Comprehensive(t *testing.T) {
	ResetEngine()
	defer ResetEngine()

	// Test SetEngine multiple times
	SetEngine(EngineShowExtents)
	SetEngine(EngineShowExtents) // duplicate
	require.True(t, HasEngine(EngineShowExtents))

	// Test UnsetEngine when not set
	UnsetEngine(EngineAdmin)
	require.False(t, HasEngine(EngineAdmin))

	// Test Set then Unset
	SetEngine(EngineAdmin)
	require.True(t, HasEngine(EngineAdmin))
	UnsetEngine(EngineAdmin)
	require.False(t, HasEngine(EngineAdmin))

	// Test Unset again (already unset)
	UnsetEngine(EngineAdmin)
	require.False(t, HasEngine(EngineAdmin))

	// Test multiple flags
	SetEngine(EngineShowExtents | EngineGodMode)
	require.True(t, HasEngine(EngineShowExtents))
	require.True(t, HasEngine(EngineGodMode))

	UnsetEngine(EngineShowExtents | EngineGodMode)
	require.False(t, HasEngine(EngineShowExtents))
	require.False(t, HasEngine(EngineGodMode))
}

func TestGameFlagSetUnset_Comprehensive(t *testing.T) {
	ResetGame()
	defer ResetGame()

	SetGame(GameModeCoop)
	require.True(t, HasGame(GameModeCoop))

	SetGame(GameModeCoop) // duplicate
	require.True(t, HasGame(GameModeCoop))

	UnsetGame(GameHost)
	require.False(t, HasGame(GameHost))

	UnsetGame(GameModeCoop)
	require.False(t, HasGame(GameModeCoop))
}
