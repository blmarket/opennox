package netstr

import (
	"net/netip"
	"testing"
	"time"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
)

func TestNetCrypt(t *testing.T) {
	// Test empty slice
	netCrypt(0xAB, []byte{})
	// Test single byte
	p := []byte{0x12}
	netCrypt(0xFF, p)
	if p[0] != 0xED {
		t.Errorf("netCrypt single byte = 0x%X, want 0xED", p[0])
	}
	// Test roundtrip
	orig := []byte{1, 2, 3, 4, 5}
	p2 := append([]byte(nil), orig...)
	netCrypt(0x5A, p2)
	netCrypt(0x5A, p2)
	for i, v := range p2 {
		if v != orig[i] {
			t.Errorf("netCrypt roundtrip failed at %d: got %d, want %d", i, v, orig[i])
		}
	}
	// Test key 0 (no change)
	p3 := []byte{10, 20, 30}
	netCrypt(0, p3)
	if p3[0] != 10 || p3[1] != 20 || p3[2] != 30 {
		t.Error("netCrypt with key 0 should not change data")
	}
}

func TestStreamsGetFreeIndex(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	ind, ok := s.getFreeIndex()
	if !ok || ind != 0 {
		t.Errorf("getFreeIndex empty = %d, %v, want 0, true", ind, ok)
	}
	// Fill first slot
	s.streams[0] = &Conn{}
	ind, ok = s.getFreeIndex()
	if !ok || ind != 1 {
		t.Errorf("getFreeIndex after fill = %d, %v, want 1, true", ind, ok)
	}
}

func TestStreamsGetFreeIndexWithAddr(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	addr1 := netip.MustParseAddrPort("127.0.0.1:1234")
	addr2 := netip.MustParseAddrPort("127.0.0.1:5678")
	ind, ok := s.getFreeIndexWithAddr(addr1)
	if !ok || ind != 0 {
		t.Errorf("getFreeIndexWithAddr empty = %d, %v", ind, ok)
	}
	s.streams[0] = &Conn{addr: addr1}
	ind, ok = s.getFreeIndexWithAddr(addr1)
	if ok {
		t.Error("getFreeIndexWithAddr existing addr should return false")
	}
	if ind != 0 {
		t.Errorf("getFreeIndexWithAddr existing = %d, want 0", ind)
	}
	ind, ok = s.getFreeIndexWithAddr(addr2)
	if !ok || ind != 1 {
		t.Errorf("getFreeIndexWithAddr new addr = %d, %v, want 1, true", ind, ok)
	}
}

func TestStreamsGetFreeNetStruct2Ind(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	ind := s.getFreeNetStruct2Ind()
	if ind != 0 {
		t.Errorf("getFreeNetStruct2Ind empty = %d, want 0", ind)
	}
	s.streams2[0].active = true
	ind = s.getFreeNetStruct2Ind()
	if ind != 1 {
		t.Errorf("getFreeNetStruct2Ind after fill = %d, want 1", ind)
	}
}

func TestStreamsConnByAddr(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	addr := netip.MustParseAddrPort("10.0.0.1:9999")
	if s.connByAddr(addr) != nil {
		t.Error("connByAddr empty should return nil")
	}
	c := &Conn{addr: addr}
	s.streams[5] = c
	if s.connByAddr(addr) != c {
		t.Error("connByAddr should find connection")
	}
}

func TestStreamsStruct2IndByAddr(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	addr := netip.MustParseAddrPort("10.0.0.1:9999")
	if s.struct2IndByAddr(addr) != -1 {
		t.Error("struct2IndByAddr empty should return -1")
	}
	s.streams2[3].addr = addr
	if s.struct2IndByAddr(addr) != 3 {
		t.Error("struct2IndByAddr should find index 3")
	}
}

func TestStreamsHasConnWithIndAndAddr(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	addr := netip.MustParseAddrPort("1.2.3.4:1111")
	s.streams[2] = &Conn{addr: addr}
	h := handle{g: s, i: 2}
	if !s.hasConnWithIndAndAddr(h, addr) {
		t.Error("hasConnWithIndAndAddr should return true")
	}
	h2 := handle{g: s, i: 3}
	if s.hasConnWithIndAndAddr(h2, addr) {
		t.Error("hasConnWithIndAndAddr wrong ind should return false")
	}
}

func TestStreamsProcessStreamOp9(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	h := handle{g: s, i: 0}
	// Too short packet
	if s.processStreamOp9(h, []byte{1, 2, 3}) != 0 {
		t.Error("processStreamOp9 short packet should return 0")
	}
	// Valid packet but dv out of range
	packet := []byte{0, 0, 0, 0}
	s.timing[0].field4 = 0
	if s.processStreamOp9(h, packet) != 0 {
		t.Error("processStreamOp9 dv=0 should return 0")
	}
}

func TestStream2NextPingPacket(t *testing.T) {
	var nx stream2
	nx.cur = 5
	tm := time.Duration(12345) * time.Millisecond
	buf := nx.nextPingPacket(tm)
	if len(buf) != 8 {
		t.Errorf("nextPingPacket len = %d, want 8", len(buf))
	}
	if buf[3] != 5 {
		t.Errorf("nextPingPacket cur = %d, want 5", buf[3])
	}
	if nx.ticks != tm {
		t.Error("nextPingPacket should set ticks")
	}
}

func TestStreamsMaybeSendCode11(t *testing.T) {
	s := NewStreams(func() uint32 { return 1000 })
	s.Now = func() time.Duration { return 0 }
	// No connections - should not panic
	s.maybeSendCode11()
	// Connection with field38 != 1 - should not send
	s.streams[0] = &Conn{field38: 0, frame40: 0}
	s.maybeSendCode11()
}

func TestStreamsGetTimingByInd1(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	s.timing[1].field28 = 42
	if s.GetTimingByInd1(ntype.PlayerInd(0)) != 42 {
		t.Error("GetTimingByInd1 should return field28")
	}
}

func TestStreamsByPlayer(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	p := &testPlayer{ind: 0}
	if s.ByPlayer(p) != nil {
		t.Error("ByPlayer empty should return nil")
	}
	c := &Conn{}
	s.streams[1] = c
	if s.ByPlayer(p) != c {
		t.Error("ByPlayer should return conn at ind+1")
	}
}

func TestStreamsStreamByPlayerInd(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	// StreamByPlayerInd returns netlib.Stream, nil *Conn becomes non-nil interface
	// Just verify it doesn't panic
	_ = s.StreamByPlayerInd(0)
}

func TestStreamsHostStream(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	// HostStream returns netlib.Stream, nil *Conn becomes non-nil interface
	_ = s.HostStream()
	c := &Conn{}
	s.streams[0] = c
	if s.HostStream() == nil {
		t.Error("HostStream should not return nil when streams[0] is set")
	}
}
