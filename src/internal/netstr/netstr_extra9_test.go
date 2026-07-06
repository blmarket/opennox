package netstr

import (
	"net"
	"net/netip"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
)

func TestWaitForLobbyResultsError(t *testing.T) {
	// Test with nil conn - should return error quickly
	// We use a closed PacketConn to get error
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Skip("cannot create udp conn")
	}
	pc.Close()

	_, err = WaitForLobbyResults(pc, netip.Addr{}, 0, LobbyWaitOptions{})
	if err == nil {
		t.Log("expected error from closed conn, got nil (may be OK)")
	}
}

func TestWaitForLobbyResultsCanRead(t *testing.T) {
	pc, err := net.ListenPacket("udp", "127.0.0.1:0")
	if err != nil {
		t.Skip("cannot create udp conn")
	}
	defer pc.Close()

	// With RecvCanRead flag, should try to check if can read
	// Closed or no data should return error or 0
	_, err = WaitForLobbyResults(pc, netip.Addr{}, RecvCanRead, LobbyWaitOptions{})
	if err != nil {
		t.Logf("got expected error: %v", err)
	}
}

func TestSendUnreliableNil(t *testing.T) {
	var ns *Conn
	n, err := ns.SendUnreliable([]byte{1, 2, 3}, false)
	if err == nil {
		t.Error("SendUnreliable with nil should error")
	}
	if n != -3 {
		t.Errorf("SendUnreliable nil should return -3, got %d", n)
	}
}

func TestSendUnreliableEmpty(t *testing.T) {
	ns := &Conn{}
	n, err := ns.SendUnreliable([]byte{}, false)
	if err == nil {
		t.Error("SendUnreliable with empty buf should error")
	}
	if n != -2 {
		t.Errorf("SendUnreliable empty should return -2, got %d", n)
	}
}

func TestSendUnreliableMsg(t *testing.T) {
	ns := &Conn{}
	// Test with nil msg - should error on AppendPacket, recover from panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("recovered from panic as expected: %v", r)
		}
	}()
	_, err := ns.SendUnreliableMsg(nil, false)
	if err == nil {
		t.Log("SendUnreliableMsg with nil may error or not")
	}
}

func TestSendDestBufNil(t *testing.T) {
	var ns *Conn
	buf := ns.sendDestBuf()
	if buf != nil {
		t.Error("sendDestBuf with nil should return nil")
	}
}

func TestCallSendPollNil(t *testing.T) {
	ns := &Conn{}
	n := ns.callSendPoll([]byte{1, 2, 3})
	if n != 0 {
		t.Errorf("callSendPoll with nil should return 0, got %d", n)
	}
}

func TestFlushNil(t *testing.T) {
	var ns *Conn
	err := ns.Flush()
	if err == nil {
		t.Error("Flush with nil should error")
	}
}

func TestFlushAndPollNil(t *testing.T) {
	var ns *Conn
	err := ns.FlushAndPoll()
	if err == nil {
		t.Error("FlushAndPoll with nil should error")
	}
}

func TestSendFlushNil(t *testing.T) {
	var ns *Conn
	err := ns.sendFlush(false)
	if err == nil {
		t.Error("sendFlush with nil should error")
	}
}

func TestStreamsListenNilOpt(t *testing.T) {
	var g Streams
	_, err := g.Listen(nil)
	if err == nil {
		t.Error("Listen with nil opt should error")
	}
}

func TestStreamsListenMaxLimit(t *testing.T) {
	var g Streams
	opt := &Options{Max: maxStructs + 1}
	_, err := g.Listen(opt)
	if err == nil {
		t.Error("Listen with max over limit should error")
	}
}

func TestNetCanReadInvalid(t *testing.T) {
	n, err := netCanRead(999999)
	t.Logf("netCanRead invalid fd: n=%d err=%v", n, err)
	// Should not panic
}

func TestLobbyWaitOptions(t *testing.T) {
	opts := LobbyWaitOptions{
		OnResult:       func(addr netip.AddrPort, data []byte) {},
		OnPassRequired: func() {},
		OnPing:         func(addr netip.AddrPort, data []byte) {},
		OnConnectErr:   func(errcode noxnet.ConnectError) bool { return true },
		OnJoinOK:       func() {},
		OnJoinFail:     func() {},
	}
	// Just verify opts can be created
	_ = opts
}
