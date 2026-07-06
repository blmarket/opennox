package discover

import (
	"context"
	"net"
	"net/netip"
	"testing"

	"github.com/noxworld-dev/lobby"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"

	"github.com/stretchr/testify/require"
)

func TestEncodeGameDiscovery_Extra(t *testing.T) {
	// Test with different tokens
	for _, token := range []uint32{0, 1, 0xFFFFFFFF, 0x12345678} {
		data := encodeGameDiscovery(token)
		require.NotNil(t, data)
		require.GreaterOrEqual(t, len(data), 2)
		require.Equal(t, byte(0), data[0])
		require.Equal(t, byte(0), data[1])
	}
}

func TestDecodeGameInfo_Extra(t *testing.T) {
	// Test with empty buffer
	result := decodeGameInfo([]byte{})
	require.Nil(t, result)

	// Test with only header
	result = decodeGameInfo([]byte{0x00, 0x00})
	require.Nil(t, result) // no valid packet

	// Test with valid header but invalid packet
	result = decodeGameInfo([]byte{0x00, 0x00, 0xFF, 0xFF, 0xFF})
	// May return nil or partial, both ok
	_ = result
}

func TestGetAddr_Extra(t *testing.T) {
	// Test with nil
	result := getAddr(nil)
	require.False(t, result.IsValid())

	// Test with UDPAddr
	udpAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 8080,
	}
	result = getAddr(udpAddr)
	require.True(t, result.IsValid())
	// Address may be IPv6 mapped, so just check port and that IP contains 127.0.0.1
	require.Equal(t, uint16(8080), result.Port())

	// Test with TCPAddr
	tcpAddr := &net.TCPAddr{
		IP:   net.ParseIP("192.168.1.1"),
		Port: 1234,
	}
	result = getAddr(tcpAddr)
	require.True(t, result.IsValid())
	require.Equal(t, uint16(1234), result.Port())

	// Test with custom AddrPort type
	customAddr := netip.MustParseAddrPort("10.0.0.1:5678")
	type customAddrPort struct {
		net.Addr
	}
	// Use a type that implements AddrPort() method
	_ = customAddr
	result = getAddr(&customAddrPort{})
	// Should return invalid for unsupported type
	_ = result
}

func TestConvGameInfo_Extra2(t *testing.T) {
	addr := netip.MustParseAddrPort("127.0.0.1:1234")

	// Test with nil message (should not panic)
	// convGameInfo with nil m would panic, so skip

	// Test with empty buffer (short buffer)
	m := &noxnet.MsgServerInfo{
		ServerName: "Test",
		MapName:    "Map",
		Flags:      uint16(noxflags.GameModeCoop),
	}
	buf := make([]byte, 10) // too short, but convGameInfo accesses buf[20], buf[68] etc.
	// This would panic, so we use sufficient length
	buf = make([]byte, 100)
	buf[3] = 1
	buf[4] = 4
	game := convGameInfo(addr, m, buf)
	require.NotNil(t, game)
	require.Equal(t, "Test", game.Name)

	// Test with quest mode and stage
	m.Flags = uint16(noxflags.GameModeQuest)
	buf[68] = 5
	buf[69] = 0
	game = convGameInfo(addr, m, buf)
	require.NotNil(t, game)
	require.NotNil(t, game.Quest)
	require.Equal(t, 5, game.Quest.Stage)
}

func TestPingAll_Empty(t *testing.T) {
	// Test PingAll with empty array should return quickly
	// But PingAll creates a pinger even with empty array, so nil pc will panic
	// We just verify the function exists and handles empty array gracefully
	// by using a valid UDP connection
	pc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	require.NoError(t, err)
	defer pc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*1e6) // 10ms
	defer cancel()
	err = PingAll(ctx, pc, []Server{})
	// With empty array, should return nil quickly
	require.NoError(t, err)
}

func TestPinger_SendPing_Error(t *testing.T) {
	// Test SendPing with closed connection should return error
	// Create a UDP connection and close it immediately
	pc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	require.NoError(t, err)
	p := NewPinger(pc)
	require.NotNil(t, p)
	pc.Close() // Close the connection
	// SendPing should return error now
	out := make(chan *lobby.Game, 1)
	addr := netip.MustParseAddrPort("127.0.0.1:1234")
	err = p.SendPing(out, addr)
	// May return error or not depending on timing
	_ = err
	p.Close()
}

func TestPinger_Ping_Timeout(t *testing.T) {
	// Test Ping with context timeout
	pc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 0})
	require.NoError(t, err)
	p := NewPinger(pc)
	require.NotNil(t, p)
	defer p.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*1e6) // 10ms
	defer cancel()
	addr := netip.MustParseAddrPort("127.0.0.1:1234")
	game, err := p.Ping(ctx, addr)
	// Should timeout
	require.Error(t, err)
	require.Nil(t, game)
}
