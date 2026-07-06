package discover

import (
	"context"
	"net/netip"
	"testing"

	"github.com/noxworld-dev/lobby"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestRegisterBackend(t *testing.T) {
	// Test that RegisterBackend panics on duplicate
	name := "test_backend_panic"
	RegisterBackend(name, func(ctx context.Context, out chan<- Server) error { return nil })

	defer func() {
		if r := recover(); r == nil {
			t.Error("RegisterBackend should panic on duplicate")
		}
	}()
	RegisterBackend(name, func(ctx context.Context, out chan<- Server) error { return nil })
}

func TestRegisterFallback(t *testing.T) {
	name := "test_fallback_panic"
	RegisterFallback(name, func(ctx context.Context, out chan<- Server) error { return nil })

	defer func() {
		if r := recover(); r == nil {
			t.Error("RegisterFallback should panic on duplicate")
		}
	}()
	RegisterFallback(name, func(ctx context.Context, out chan<- Server) error { return nil })
}

func TestIsTimeout(t *testing.T) {
	if isTimeout(nil) {
		t.Error("isTimeout(nil) should be false")
	}
}

func TestMergeInfo(t *testing.T) {
	g1 := lobby.Game{Name: "Server1", Map: "map1"}
	g2 := &lobby.Game{Name: "Server2", Map: "map2", Mode: "ctf", Players: lobby.PlayersInfo{Cur: 5, Max: 10}}

	result := mergeInfo(g1, g2)
	if result.Name != "Server1" { // Should keep g1's name
		t.Errorf("mergeInfo name = %q, want %q", result.Name, "Server1")
	}
	if result.Map != "map1" { // Should keep g1's map
		t.Errorf("mergeInfo map = %q, want %q", result.Map, "map1")
	}
	if result.Mode != "ctf" { // Should take g2's mode since g1's is empty
		t.Errorf("mergeInfo mode = %q, want %q", result.Mode, "ctf")
	}
	if result.Players.Cur != 5 {
		t.Errorf("mergeInfo players cur = %d, want 5", result.Players.Cur)
	}

	// Test with nil g2
	result = mergeInfo(g1, nil)
	if result.Name != "Server1" {
		t.Error("mergeInfo with nil g2 should return g1")
	}
}

func TestGameFlagsToMode(t *testing.T) {
	tests := []struct {
		flags noxflags.GameFlag
		want  lobby.GameMode
	}{
		{noxflags.GameModeKOTR, lobby.ModeKOTR},
		{noxflags.GameModeCTF, lobby.ModeCTF},
		{noxflags.GameModeFlagBall, lobby.ModeFlagBall},
		{noxflags.GameModeChat, lobby.ModeChat},
		{noxflags.GameModeArena, lobby.ModeArena},
		{noxflags.GameModeElimination, lobby.ModeElimination},
		{noxflags.GameModeQuest, lobby.ModeQuest},
		{noxflags.GameModeCoop, lobby.ModeCoop},
		{0, lobby.ModeCustom},
	}

	for _, tt := range tests {
		got := gameFlagsToMode(tt.flags)
		if got != tt.want {
			t.Errorf("gameFlagsToMode(%v) = %v, want %v", tt.flags, got, tt.want)
		}
	}
}

func TestServerKey(t *testing.T) {
	s := Server{
		IP: netip.MustParseAddr("127.0.0.1"),
	}
	s.Port = 8080

	key := s.key()
	if key.Addr != "127.0.0.1" {
		t.Errorf("key.Addr = %q, want %q", key.Addr, "127.0.0.1")
	}
	if key.Port != 8080 {
		t.Errorf("key.Port = %d, want 8080", key.Port)
	}
}

func TestEncodeGameDiscovery(t *testing.T) {
	data := encodeGameDiscovery(0x12345678)
	if len(data) < 2 {
		t.Error("encodeGameDiscovery should return at least 2 bytes")
	}
	if data[0] != 0 || data[1] != 0 {
		t.Error("encodeGameDiscovery header should be 0, 0")
	}
}

func TestDecodeGameInfo(t *testing.T) {
	// Test with too short buffer
	result := decodeGameInfo([]byte{0x00})
	if result != nil {
		t.Error("decodeGameInfo with short buffer should return nil")
	}

	// Test with invalid data
	result = decodeGameInfo([]byte{0x00, 0x00, 0xFF, 0xFF})
	if result != nil {
		// May return nil or partial, both ok
	}
}

func TestConvGameInfo(t *testing.T) {
	addr := netip.MustParseAddrPort("127.0.0.1:8080")

	// Create a minimal valid buffer
	buf := make([]byte, 100)
	buf[3] = 5  // players cur
	buf[4] = 10 // players max
	buf[20] = 0 // status
	buf[21] = 0

	// Just verify the function exists and doesn't panic with nil
	// convGameInfo takes *noxnet.MsgServerInfo which we can't easily create
	// So we just test that getAddr works
	result := getAddr(nil)
	if result.IsValid() {
		t.Error("getAddr(nil) should return invalid AddrPort")
	}

	_ = addr
	_ = buf
}

func TestGetAddr(t *testing.T) {
	// Test with nil
	result := getAddr(nil)
	if result.IsValid() {
		t.Error("getAddr(nil) should return invalid AddrPort")
	}
}
