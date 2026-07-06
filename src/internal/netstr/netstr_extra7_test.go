package netstr

import (
	"net/netip"
	"testing"
)

func TestNewStreamsExtra(t *testing.T) {
	s := NewStreams(func() uint32 { return 12345 })
	if s == nil {
		t.Fatal("NewStreams returned nil")
	}
	s.reset()
}

func TestConnMethodsExtra(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	c := &Conn{g: s, ind: 1, id: 1}

	// These should not panic with nil checks
	var c2 *Conn
	c2.SendServerClose()
	c2.SendClientClose()
	c2.SendCode6()
	c2.SendSelfRaw([]byte{1, 2, 3})

	c.RecvLoop(false)
}

func TestGetFreeMethodsExtra(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })

	idx, _ := s.getFreeIndex()
	if idx < 0 {
		t.Error("getFreeIndex should return non-negative")
	}

	idx2 := s.getFreeNetStruct2Ind()
	if idx2 < 0 {
		t.Error("getFreeNetStruct2Ind should return non-negative")
	}

	addr := netip.MustParseAddrPort("127.0.0.1:1234")
	idx3, _ := s.getFreeIndexWithAddr(addr)
	if idx3 < 0 {
		t.Error("getFreeIndexWithAddr should return non-negative")
	}
}

func TestMaybeSendCode11Extra(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	s.maybeSendCode11()
}
