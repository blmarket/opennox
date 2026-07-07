package discover

import (
	"net"
	"net/netip"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
)

func TestPingEncodeDecode(t *testing.T) {
	data := encodeGameDiscovery(12345)
	if len(data) == 0 {
		t.Error("expected non-empty data")
	}

	addr := &net.UDPAddr{IP: net.ParseIP("127.0.0.1"), Port: 18590}
	s := getAddr(addr)
	if !s.IsValid() {
		t.Error("expected valid addr")
	}

	if s := getAddr(nil); s.IsValid() {
		t.Error("expected invalid for nil addr")
	}

	if m := decodeGameInfo([]byte{1, 2, 3}); m != nil {
		t.Error("expected nil for short data")
	}

	if m := decodeGameInfo([]byte("NOX!")); m != nil {
		t.Error("expected nil for short NOX data")
	}

	m := &noxnet.MsgServerInfo{
		ServerName: "Test",
		MapName:    "testmap",
	}
	buf := make([]byte, 100)
	g := convGameInfo(netip.AddrPortFrom(netip.MustParseAddr("127.0.0.1"), 18590), m, buf)
	if g == nil {
		t.Error("expected non-nil game")
	}

	g2 := convGameInfo(netip.AddrPort{}, m, buf)
	_ = g2
}
