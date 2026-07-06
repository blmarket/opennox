package noxflags

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseEngineFlag(t *testing.T) {
	// Single flag by name
	f, err := ParseEngineFlag("ShowExtents")
	require.NoError(t, err)
	require.Equal(t, EngineShowExtents, f)

	// With Engine prefix
	f, err = ParseEngineFlag("EngineShowExtents")
	require.NoError(t, err)
	require.Equal(t, EngineShowExtents, f)

	// Numeric
	f, err = ParseEngineFlag("0x2")
	require.NoError(t, err)
	require.Equal(t, EngineShowExtents, f)

	// Comma separated
	f, err = ParseEngineFlag("ShowExtents,Admin")
	require.NoError(t, err)
	require.True(t, f.Has(EngineShowExtents))
	require.True(t, f.Has(EngineAdmin))

	// Pipe separated
	f, err = ParseEngineFlag("ShowExtents|Admin")
	require.NoError(t, err)
	require.True(t, f.Has(EngineShowExtents))
	require.True(t, f.Has(EngineAdmin))

	// Invalid
	_, err = ParseEngineFlag("InvalidFlagName12345")
	require.Error(t, err)
}

func TestEngineFlagString(t *testing.T) {
	require.Equal(t, "ShowExtents", EngineShowExtents.String())
	require.Equal(t, "Admin", EngineAdmin.String())

	// Combined
	f := EngineShowExtents | EngineAdmin
	s := f.String()
	require.Contains(t, s, "ShowExtents")
	require.Contains(t, s, "Admin")
}

func TestEngineFlagGoString(t *testing.T) {
	s := EngineShowExtents.GoString()
	require.Contains(t, s, "EngineShowExtents")

	f := EngineShowExtents | EngineAdmin
	s = f.GoString()
	require.Contains(t, s, "EngineShowExtents")
	require.Contains(t, s, "EngineAdmin")
}

func TestEngineFlagSplit(t *testing.T) {
	f := EngineShowExtents | EngineAdmin | EngineGodMode
	list := f.Split()
	require.Len(t, list, 3)
	require.Contains(t, list, EngineShowExtents)
	require.Contains(t, list, EngineAdmin)
	require.Contains(t, list, EngineGodMode)
}

func TestEngineFlagHas(t *testing.T) {
	f := EngineShowExtents | EngineAdmin
	require.True(t, f.Has(EngineShowExtents))
	require.True(t, f.Has(EngineAdmin))
	require.False(t, f.Has(EngineGodMode))

	require.True(t, f.HasAll(EngineShowExtents|EngineAdmin))
	require.False(t, f.HasAll(EngineShowExtents|EngineGodMode))
}

func TestEngineFlagJSON(t *testing.T) {
	f := EngineShowExtents | EngineAdmin
	data, err := json.Marshal(f)
	require.NoError(t, err)

	var f2 EngineFlag
	err = json.Unmarshal(data, &f2)
	require.NoError(t, err)
	require.Equal(t, f, f2)

	// Numeric JSON
	err = json.Unmarshal([]byte("2"), &f2)
	require.NoError(t, err)
	require.Equal(t, EngineShowExtents, f2)
}

func TestEngineFlagGetSet(t *testing.T) {
	ResetEngine()
	require.Equal(t, EngineFlag(0), GetEngine())

	SetEngine(EngineShowExtents)
	require.True(t, HasEngine(EngineShowExtents))

	SetEngine(EngineAdmin)
	require.True(t, HasEngine(EngineAdmin))
	require.True(t, HasEngine(EngineShowExtents))

	UnsetEngine(EngineShowExtents)
	require.False(t, HasEngine(EngineShowExtents))
	require.True(t, HasEngine(EngineAdmin))

	ToggleEngine(EngineAdmin)
	require.False(t, HasEngine(EngineAdmin))

	ToggleEngine(EngineGodMode)
	require.True(t, HasEngine(EngineGodMode))

	ResetEngine()
	require.Equal(t, EngineFlag(0), GetEngine())
}

func TestParseGameFlag(t *testing.T) {
	f, err := ParseGameFlag("Coop")
	require.NoError(t, err)
	require.NotEqual(t, GameFlag(0), f)

	// Numeric
	f, err = ParseGameFlag("1")
	require.NoError(t, err)

	// Invalid
	_, err = ParseGameFlag("InvalidGameFlagXYZ")
	require.Error(t, err)
}

func TestGameFlagString(t *testing.T) {
	var f GameFlag = 1
	s := f.String()
	require.NotEmpty(t, s)

	s = f.GoString()
	require.NotEmpty(t, s)
}

func TestGameFlagGetSet(t *testing.T) {
	ResetGame()
	SetGame(1)
	require.True(t, HasGame(1))
	UnsetGame(1)
	require.False(t, HasGame(1))
}

func TestGameplayFlag(t *testing.T) {
	// Just ensure functions exist and don't panic
	_ = GetGamePlay()
	_ = HasGamePlay(1)
	SetGamePlay(1)
	UnsetGamePlay(1)
}
