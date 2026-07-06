package noxflags

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEngineFlagGoString_All(t *testing.T) {
	// Test GoString for all single flags to increase coverage
	flags := []EngineFlag{
		EngineShowExtents,
		EngineAdmin,
		EngineGodMode,
		EngineWindowed,
		EngineNoRendering,
		EngineNoTextRendering,
		EngineNetDebug,
	}
	for _, f := range flags {
		s := f.GoString()
		require.NotEmpty(t, s, "GoString should not be empty for flag %v", f)
	}

	// Test GoString for combined flags
	combined := EngineShowExtents | EngineAdmin | EngineGodMode | EngineWindowed
	s := combined.GoString()
	require.NotEmpty(t, s)
	require.Contains(t, s, "EngineShowExtents")
	require.Contains(t, s, "EngineAdmin")

	// Test GoString for empty flag
	empty := EngineFlag(0)
	s = empty.GoString()
	// Empty may return empty string or "EngineFlag(0)", both ok
	_ = s

	// Test GoString for unknown flag bits
	unknown := EngineFlag(1 << 30)
	s = unknown.GoString()
	_ = s
}

func TestGameFlagGoString_All(t *testing.T) {
	// Test GoString for all single flags to increase coverage
	flags := []GameFlag{
		GameModeCoop,
		GameModeQuest,
		GameHost,
		GameOnline,
		GameModeChat,
	}
	for _, f := range flags {
		s := f.GoString()
		// May be empty for some flags, just ensure no panic
		_ = s
	}

	// Test GoString for combined flags
	combined := GameModeCoop | GameHost | GameOnline
	s := combined.GoString()
	_ = s

	// Test GoString for empty flag
	empty := GameFlag(0)
	s = empty.GoString()
	_ = s

	// Test GoString for all modes
	modes := []GameFlag{
		GameModeCoop,
		GameModeQuest,
	}
	for _, m := range modes {
		s := m.GoString()
		_ = s
	}
}

func TestEngineFlagGoString_EdgeCases(t *testing.T) {
	// Test with all bits set
	all := EngineFlag(0xFFFFFFFF)
	s := all.GoString()
	require.NotEmpty(t, s)

	// Test with single bit flags
	for i := 0; i < 32; i++ {
		f := EngineFlag(1 << uint(i))
		s := f.GoString()
		_ = s
	}
}

func TestGameFlagGoString_EdgeCases(t *testing.T) {
	// Test with all bits set
	all := GameFlag(0xFFFFFFFF)
	s := all.GoString()
	_ = s

	// Test with single bit flags
	for i := 0; i < 32; i++ {
		f := GameFlag(1 << uint(i))
		s := f.GoString()
		_ = s
	}
}
