package discover

import (
	"context"
	"testing"

	"github.com/noxworld-dev/lobby"
	"github.com/stretchr/testify/require"
)

func TestIsTimeoutExtra(t *testing.T) {
	require.False(t, isTimeout(nil))
	require.True(t, isTimeout(context.Canceled))
	require.True(t, isTimeout(context.DeadlineExceeded))
	require.False(t, isTimeout(context.Background().Err()))
}

func TestMergeInfoExtra(t *testing.T) {
	// Test empty g1 fields filled from g2
	g1 := lobby.Game{}
	g2 := &lobby.Game{Name: "Test", Map: "map1", Mode: lobby.ModeArena, Players: lobby.PlayersInfo{Cur: 3, Max: 8}}
	result := mergeInfo(g1, g2)
	require.Equal(t, "Test", result.Name)
	require.Equal(t, "map1", result.Map)
	require.Equal(t, lobby.ModeArena, result.Mode)
	require.Equal(t, 3, result.Players.Cur)

	// Test g1 fields preserved
	g1 = lobby.Game{Name: "Keep", Map: "keepmap", Mode: lobby.ModeKOTR}
	g2 = &lobby.Game{Name: "Ignore", Map: "ignoremap", Mode: lobby.ModeCTF, Players: lobby.PlayersInfo{Cur: 1, Max: 2}}
	result = mergeInfo(g1, g2)
	require.Equal(t, "Keep", result.Name)
	require.Equal(t, "keepmap", result.Map)
	require.Equal(t, lobby.ModeKOTR, result.Mode)
	require.Equal(t, 1, result.Players.Cur)
}

func TestServerKeyExtra(t *testing.T) {
	s := Server{}
	s.Port = 0
	key := s.key()
	require.Equal(t, "invalid IP", key.Addr)
	require.Equal(t, 0, key.Port)
}
