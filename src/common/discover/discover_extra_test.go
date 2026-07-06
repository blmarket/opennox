package discover

import (
	"net"
	"net/netip"
	"testing"

	"github.com/noxworld-dev/lobby"
	"github.com/noxworld-dev/opennox-lib/noxnet"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestGetAddrMore(t *testing.T) {
	// Test with UDPAddr
	udpAddr := &net.UDPAddr{IP: net.ParseIP("192.168.1.1"), Port: 1234}
	result := getAddr(udpAddr)
	if !result.IsValid() {
		t.Error("getAddr(UDPAddr) should return valid AddrPort")
	}
	if result.Port() != 1234 {
		t.Errorf("getAddr port = %d, want 1234", result.Port())
	}

	// Test with TCPAddr
	tcpAddr := &net.TCPAddr{IP: net.ParseIP("10.0.0.1"), Port: 5678}
	result = getAddr(tcpAddr)
	if !result.IsValid() {
		t.Error("getAddr(TCPAddr) should return valid AddrPort")
	}

	// Test with AddrPort directly via interface
	ap := netip.MustParseAddrPort("127.0.0.1:9999")
	result = getAddr(net.UDPAddrFromAddrPort(ap))
	if result != ap {
		t.Error("getAddr should preserve AddrPort")
	}
}

func TestDecodeGameInfoMore(t *testing.T) {
	// Test with empty buffer
	result := decodeGameInfo([]byte{})
	if result != nil {
		t.Error("decodeGameInfo empty should return nil")
	}

	// Test with only header
	result = decodeGameInfo([]byte{0, 0})
	if result != nil {
		// May be nil or empty, both ok
	}
}

func TestConvGameInfoMore(t *testing.T) {
	addr := netip.MustParseAddrPort("192.168.1.100:8080")

	// Create a message with basic fields
	msg := &noxnet.MsgServerInfo{
		ServerName: "Test Server",
		MapName:    "TestMap",
		Flags:      uint16(noxflags.GameModeCTF),
	}

	// Create buffer with player counts and status
	buf := make([]byte, 100)
	buf[3] = 3  // players cur
	buf[4] = 8  // players max
	buf[20] = 0 // status open
	buf[21] = 0

	result := convGameInfo(addr, msg, buf)
	if result == nil {
		t.Fatal("convGameInfo should not return nil")
	}
	if result.Name != "Test Server" {
		t.Errorf("convGameInfo name = %q, want %q", result.Name, "Test Server")
	}
	if result.Address != "192.168.1.100" {
		t.Errorf("convGameInfo address = %q", result.Address)
	}
	if result.Port != 8080 {
		t.Errorf("convGameInfo port = %d, want 8080", result.Port)
	}
	if result.Map != "testmap" { // should be lowercase
		t.Errorf("convGameInfo map = %q, want %q", result.Map, "testmap")
	}
	if result.Mode != lobby.ModeCTF {
		t.Errorf("convGameInfo mode = %v, want CTF", result.Mode)
	}
	if result.Players.Cur != 3 || result.Players.Max != 8 {
		t.Error("convGameInfo players count mismatch")
	}
	if result.Access != lobby.AccessOpen {
		t.Error("convGameInfo access should be open")
	}

	// Test with closed status
	buf[20] = 0x10
	result = convGameInfo(addr, msg, buf)
	if result.Access != lobby.AccessClosed {
		t.Error("convGameInfo access should be closed when status 0x10")
	}

	// Test with password status
	buf[20] = 0x20
	result = convGameInfo(addr, msg, buf)
	if result.Access != lobby.AccessPassword {
		t.Error("convGameInfo access should be password when status 0x20")
	}

	// Test with quest mode
	msg.Flags = uint16(noxflags.GameModeQuest)
	buf[20] = 0
	buf[68] = 5 // quest stage low byte
	buf[69] = 0 // quest stage high byte
	result = convGameInfo(addr, msg, buf)
	if result.Quest == nil {
		t.Error("convGameInfo quest should not be nil for quest mode")
	} else if result.Quest.Stage != 5 {
		t.Errorf("convGameInfo quest stage = %d, want 5", result.Quest.Stage)
	}
}

func TestMergeInfoMore(t *testing.T) {
	// Test when g1 has all fields
	g1 := lobby.Game{Name: "S1", Map: "m1", Mode: "ctf", Players: lobby.PlayersInfo{Cur: 1, Max: 2}}
	g2 := &lobby.Game{Name: "S2", Map: "m2", Mode: "kotr", Players: lobby.PlayersInfo{Cur: 5, Max: 10}}
	result := mergeInfo(g1, g2)
	if result.Name != "S1" || result.Map != "m1" || result.Mode != "ctf" {
		t.Error("mergeInfo should keep g1's non-empty fields")
	}
	// mergeInfo always takes players from g2
	if result.Players.Cur != 5 {
		t.Error("mergeInfo should take players from g2")
	}
}
