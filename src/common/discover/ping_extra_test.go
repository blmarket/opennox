package discover

import (
	"net/netip"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"

	"github.com/stretchr/testify/require"
)

func TestConvGameInfo_Extra(t *testing.T) {
	addr := netip.MustParseAddrPort("127.0.0.1:1234")
	m := &noxnet.MsgServerInfo{
		ServerName: "TestServer",
		MapName:    "TestMap",
		Flags:      uint16(noxflags.GameModeCoop),
	}
	buf := make([]byte, 100)
	buf[3] = 2 // cur players
	buf[4] = 8 // max players
	buf[20] = 0
	buf[21] = 0

	game := convGameInfo(addr, m, buf)
	require.NotNil(t, game)
	require.Equal(t, "TestServer", game.Name)
	require.Equal(t, "127.0.0.1", game.Address)
	require.Equal(t, 1234, game.Port)
	require.Equal(t, "testmap", game.Map)
	require.Equal(t, 2, game.Players.Cur)
	require.Equal(t, 8, game.Players.Max)

	// Test closed access
	buf[20] = 0x10
	game = convGameInfo(addr, m, buf)
	require.NotNil(t, game)

	// Test password access
	buf[20] = 0x20
	game = convGameInfo(addr, m, buf)
	require.NotNil(t, game)

	// Test quest mode
	m.Flags = uint16(noxflags.GameModeQuest)
	buf[68] = 3
	buf[69] = 0
	game = convGameInfo(addr, m, buf)
	require.NotNil(t, game)
	require.NotNil(t, game.Quest)
}
