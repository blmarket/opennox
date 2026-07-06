package noxflags

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGameFlagJSON(t *testing.T) {
	f := GameModeCoop | GameModeQuest
	data, err := json.Marshal(f)
	require.NoError(t, err)

	var f2 GameFlag
	err = json.Unmarshal(data, &f2)
	require.NoError(t, err)
	require.Equal(t, f, f2)

	// Numeric JSON
	err = json.Unmarshal([]byte("2048"), &f2)
	require.NoError(t, err)
	require.Equal(t, GameModeCoop, f2)

	// Invalid JSON
	err = json.Unmarshal([]byte(`["InvalidFlag"]`), &f2)
	require.Error(t, err)
}

func TestGameFlagMode(t *testing.T) {
	f := GameModeQuest | GameHost | GameOnline
	mode := f.Mode()
	require.Equal(t, GameModeQuest, mode)

	f = GameModeCTF | GameModeKOTR
	mode = f.Mode()
	// Should return both modes
	require.True(t, mode.Has(GameModeCTF) || mode.Has(GameModeKOTR))
}

func TestGameFlagModeString(t *testing.T) {
	require.Equal(t, "kotr", GameModeKOTR.ModeString())
	require.Equal(t, "ctf", GameModeCTF.ModeString())
	require.Equal(t, "flagball", GameModeFlagBall.ModeString())
	require.Equal(t, "chat", GameModeChat.ModeString())
	require.Equal(t, "arena", GameModeArena.ModeString())
	require.Equal(t, "elimination", GameModeElimination.ModeString())
	require.Equal(t, "quest", GameModeQuest.ModeString())
	require.Equal(t, "custom", GameFlag(0).ModeString())
	require.Equal(t, "custom", (GameHost | GameOnline).ModeString())
}

func TestGameFlagSplit(t *testing.T) {
	f := GameHost | GameClient | GameModeCoop
	list := f.Split()
	require.Len(t, list, 3)
	require.Contains(t, list, GameHost)
	require.Contains(t, list, GameClient)
	require.Contains(t, list, GameModeCoop)

	// Empty
	f = GameFlag(0)
	list = f.Split()
	require.Empty(t, list)
}

func TestGameFlagSplitString(t *testing.T) {
	f := GameHost | GameClient
	arr := f.SplitString()
	require.Len(t, arr, 2)
	require.Contains(t, arr, "Host")
	require.Contains(t, arr, "Client")

	arr2 := f.SplitGoString()
	require.Len(t, arr2, 2)
}

func TestGameFlagHasAll(t *testing.T) {
	f := GameHost | GameClient | GameModeCoop
	require.True(t, f.HasAll(GameHost|GameClient))
	require.False(t, f.HasAll(GameHost|GameModeQuest))
	require.True(t, f.HasAll(GameFlag(0)))
}

func TestGameFlagGetGame(t *testing.T) {
	ResetGame()
	require.Equal(t, GameFlag(0), GetGame())

	SetGame(GameHost | GameModeCoop)
	require.Equal(t, GameHost|GameModeCoop, GetGame())

	ResetGame()
}

func TestGameFlagHooks(t *testing.T) {
	ResetGame()

	var changed []GameFlag
	var set []GameFlag
	var unset []GameFlag

	OnGameChange(func(f GameFlag) {
		changed = append(changed, f)
	})
	OnGameSet(func(f GameFlag) {
		set = append(set, f)
	})
	OnGameUnset(func(f GameFlag) {
		unset = append(unset, f)
	})

	SetGame(GameHost)
	require.Len(t, changed, 1)
	require.Len(t, set, 1)
	require.Equal(t, GameHost, changed[0])

	// Setting already set flag should not trigger
	changed = nil
	set = nil
	SetGame(GameHost)
	require.Empty(t, changed)
	require.Empty(t, set)

	UnsetGame(GameHost)
	require.Len(t, changed, 1)
	require.Len(t, unset, 1)

	// Unsetting already unset should not trigger
	changed = nil
	unset = nil
	UnsetGame(GameHost)
	require.Empty(t, changed)
	require.Empty(t, unset)

	ResetGame()
}

func TestParseGameFlagExtra(t *testing.T) {
	// Single flag by name
	f, err := ParseGameFlag("Coop")
	require.NoError(t, err)
	require.Equal(t, GameModeCoop, f)

	// With Game prefix
	f, err = ParseGameFlag("GameCoop")
	require.NoError(t, err)
	require.Equal(t, GameModeCoop, f)

	// Comma separated
	f, err = ParseGameFlag("Coop,Quest")
	require.NoError(t, err)
	require.True(t, f.Has(GameModeCoop))
	require.True(t, f.Has(GameModeQuest))

	// Pipe separated
	f, err = ParseGameFlag("Coop|Quest")
	require.NoError(t, err)
	require.True(t, f.Has(GameModeCoop))
	require.True(t, f.Has(GameModeQuest))
}
