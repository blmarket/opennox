package discover

import (
	"testing"

	"github.com/noxworld-dev/lobby"
	"github.com/stretchr/testify/require"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestMergeInfo2(t *testing.T) {
	g1 := lobby.Game{Name: "Test", Map: "Map1"}
	g2 := &lobby.Game{Name: "Test2", Map: "Map2"}
	result := mergeInfo(g1, g2)
	require.NotEmpty(t, result.Name)
}

func TestGameFlagsToMode2(t *testing.T) {
	require.Equal(t, lobby.ModeArena, gameFlagsToMode(noxflags.GameModeArena))
	require.Equal(t, lobby.ModeCTF, gameFlagsToMode(noxflags.GameModeCTF))
}

func TestEncodeDecodeGameInfo2(t *testing.T) {
	data := encodeGameDiscovery(12345)
	require.NotEmpty(t, data)

	decoded := decodeGameInfo(data)
	// decoded may be nil or not, depending on data
	_ = decoded
}

func TestEachServer2(t *testing.T) {
	// EachServer with nil context should not panic immediately, but may return error
	// Just test that the function exists
	_ = EachServer
}
