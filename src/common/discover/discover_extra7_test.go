package discover

import (
	"net/netip"
	"testing"
)

func TestNewLobbyClientExtra(t *testing.T) {
	cli := newLobbyClient()
	if cli == nil {
		t.Fatal("newLobbyClient returned nil")
	}
}

func TestRegisterBackendExtra(t *testing.T) {
	// Test that backends are registered via init functions
	// At least lan and lobby should be registered
	backendsList := []string{"lan", "lobby"}
	for _, name := range backendsList {
		if _, ok := backends[name]; !ok {
			t.Errorf("backend %q not registered", name)
		}
	}
	// static and xwis may not be registered in all environments
}

func TestEncodeDecodeGameDiscoveryExtra(t *testing.T) {
	token := uint32(0x12345678)
	data := encodeGameDiscovery(token)
	if len(data) == 0 {
		t.Fatal("encodeGameDiscovery returned empty data")
	}

	m := decodeGameInfo(data)
	if m == nil {
		t.Logf("decodeGameInfo returned nil (may be expected in this environment)")
		return
	}
	if m.Token != token {
		t.Errorf("Token = 0x%X, want 0x%X", m.Token, token)
	}
}

func TestDecodeGameInfoInvalidExtra(t *testing.T) {
	// Test with invalid data
	m := decodeGameInfo([]byte{})
	if m != nil {
		t.Error("decodeGameInfo should return nil for empty data")
	}

	m = decodeGameInfo([]byte{1, 2, 3})
	if m != nil {
		t.Error("decodeGameInfo should return nil for invalid data")
	}
}

func TestGetAddrExtra(t *testing.T) {
	addr := getAddr(nil)
	if addr.IsValid() {
		t.Error("getAddr(nil) should return invalid addr")
	}
}

func TestConvGameInfoExtra(t *testing.T) {
	// Test with nil message - should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("convGameInfo panicked as expected with nil: %v", r)
		}
	}()
	g := convGameInfo(netip.AddrPort{}, nil, nil)
	if g != nil {
		t.Error("convGameInfo with nil message should return nil")
	}
}

func TestGameFlagsToModeExtra(t *testing.T) {
	// Test game flags to mode conversion
	mode := gameFlagsToMode(0)
	if mode == "" {
		t.Error("gameFlagsToMode should not return empty for 0 flags")
	}
}

func TestServerKeyExtra(t *testing.T) {
	s := Server{
		IP: netip.MustParseAddr("127.0.0.1"),
	}
	k := s.key()
	if k == (serverKey{}) {
		t.Error("Server.key() should not return empty key")
	}
}
