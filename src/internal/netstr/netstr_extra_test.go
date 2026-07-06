package netstr

import (
	"fmt"
	"net/netip"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
)

func TestConnString(t *testing.T) {
	var c *Conn
	if got := c.String(); got != "<nil>" {
		t.Errorf("nil Conn String = %q, want <nil>", got)
	}
	c = &Conn{ind: 1, id: 42, addr: netip.MustParseAddrPort("127.0.0.1:1234")}
	s := c.String()
	if s == "" || s == "<nil>" {
		t.Error("Conn String should not be empty for non-nil")
	}
}

func TestConnIsHost(t *testing.T) {
	var c *Conn
	if c.IsHost() {
		t.Error("nil Conn should not be host")
	}
	c = &Conn{ind: 0}
	if !c.IsHost() {
		t.Error("ind 0 should be host")
	}
	c = &Conn{ind: 1}
	if c.IsHost() {
		t.Error("ind 1 should not be host")
	}
}

func TestConnPlayer(t *testing.T) {
	var c *Conn
	if c.Player() != 0 {
		t.Error("nil Conn Player should be 0")
	}
	c = &Conn{id: 1}
	if c.Player() != ntype.PlayerInd(0) {
		t.Errorf("Player for id 1 = %d, want 0", c.Player())
	}
	c = &Conn{id: 2}
	if c.Player() != ntype.PlayerInd(1) {
		t.Errorf("Player for id 2 = %d, want 1", c.Player())
	}
}

func TestConnAddrIP(t *testing.T) {
	addr := netip.MustParseAddrPort("192.168.1.1:5678")
	c := &Conn{addr: addr}
	if c.Addr() != addr {
		t.Errorf("Addr = %v, want %v", c.Addr(), addr)
	}
	if c.IP() != addr.Addr() {
		t.Errorf("IP = %v, want %v", c.IP(), addr.Addr())
	}
	c.setAddr(netip.MustParseAddrPort("10.0.0.1:9999"))
	if c.Addr().String() != "10.0.0.1:9999" {
		t.Error("setAddr failed")
	}
}

func TestConnHandle(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	c := &Conn{g: s, id: 5}
	h := c.handle()
	if !h.Valid() {
		t.Error("handle should be valid")
	}
	if h.IsHost() {
		t.Error("id 5 should not be host")
	}
	if h.Player() != ntype.PlayerInd(4) {
		t.Errorf("Player = %d, want 4", h.Player())
	}
}

func TestHandle(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	h := handle{g: s, i: 0}
	if !h.Valid() {
		t.Error("handle with i=0 should be valid")
	}
	if !h.IsHost() {
		t.Error("i=0 should be host")
	}
	if h.Player() != ntype.PlayerInd(-1) {
		t.Errorf("Player for host = %d", h.Player())
	}
	h2 := handle{g: s, i: 1}
	if h2.IsHost() {
		t.Error("i=1 should not be host")
	}
	if h2.Player() != ntype.PlayerInd(0) {
		t.Errorf("Player for i=1 = %d, want 0", h2.Player())
	}
	h3 := handle{g: nil, i: 0}
	if h3.Valid() {
		t.Error("nil g should be invalid")
	}
	if h3.Get() != nil {
		t.Error("Get on invalid should return nil")
	}
	eh := errHandle(-1)
	if eh.Valid() {
		t.Error("errHandle should be invalid")
	}
}

func TestDecodeEncodeMessage(t *testing.T) {
	out := make([]byte, 10)
	n := encodeMessage(out, nil)
	if n != 0 {
		t.Errorf("encode nil = %d, want 0", n)
	}
	packet := []byte{0, 0, 0, 0}
	dst := &testMsgErr{}
	ok := decodeMessage(packet, dst)
	if ok {
		t.Error("decode should fail on invalid packet")
	}
}

type testMsgErr struct{}

func (m *testMsgErr) NetOp() noxnet.Op                { return 0 }
func (m *testMsgErr) EncodeSize() int                 { return 0 }
func (m *testMsgErr) Encode(data []byte) (int, error) { return 0, nil }
func (m *testMsgErr) Decode(p []byte) (int, error) {
	return 0, fmt.Errorf("error")
}

func TestStreamsMethods(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	if s == nil {
		t.Fatal("NewStreams returned nil")
	}
	s.reset()
	if s.Host() != nil {
		t.Error("Host should be nil after reset")
	}
	if s.ConnByPlayerInd(ntype.PlayerInd(100)) != nil {
		t.Error("ConnByPlayerInd out of range should return nil")
	}
	if s.ByPlayer(&testPlayer{ind: 0}) != nil {
		t.Error("ByPlayer should be nil when no conn")
	}
	if s.GetTimingByInd1(0) != 0 {
		t.Error("GetTimingByInd1 should return 0")
	}
}

type testPlayer struct {
	ind ntype.PlayerInd
}

func (p *testPlayer) PlayerIndex() ntype.PlayerInd { return p.ind }
