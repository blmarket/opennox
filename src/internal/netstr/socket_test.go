package netstr

import (
	"net"
	"net/netip"
	"testing"
	"time"
)

func TestGetAddr(t *testing.T) {
	// nil
	ap := GetAddr(nil)
	if ap.IsValid() {
		t.Fatal("nil should return invalid AddrPort")
	}

	// TCPAddr
	tcp := &net.TCPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 8080,
	}
	ap = GetAddr(tcp)
	if !ap.IsValid() || ap.Port() != 8080 {
		t.Fatalf("expected valid TCP addr with port 8080, got %v", ap)
	}

	// UDPAddr
	udp := &net.UDPAddr{
		IP:   net.IPv4(10, 0, 0, 1),
		Port: 9090,
	}
	ap = GetAddr(udp)
	if !ap.IsValid() || ap.Port() != 9090 {
		t.Fatalf("expected valid UDP addr with port 9090, got %v", ap)
	}

	// unsupported type
	ap = GetAddr(&net.IPAddr{})
	if ap.IsValid() {
		t.Fatal("unsupported type should return invalid")
	}
}

type mockAddrPort struct {
	ap netip.AddrPort
}

func (m mockAddrPort) AddrPort() netip.AddrPort {
	return m.ap
}

func (m mockAddrPort) Network() string { return "udp" }
func (m mockAddrPort) String() string  { return m.ap.String() }

func TestGetAddr_Interface(t *testing.T) {
	ap := netip.AddrPortFrom(netip.AddrFrom4([4]byte{1, 2, 3, 4}), 1234)
	m := mockAddrPort{ap: ap}
	result := GetAddr(m)
	if result != ap {
		t.Fatalf("expected %v, got %v", ap, result)
	}
}

func TestListenOnFreePort_Invalid(t *testing.T) {
	// port too low
	pc, port, err := listenOnFreePort(Log, 0)
	if err == nil {
		t.Fatal("expected error for port 0")
	}
	if pc != nil || port != 0 {
		t.Fatal("expected nil pc and 0 port on error")
	}

	// port too high
	pc, port, err = listenOnFreePort(Log, 0x10001)
	if err == nil {
		t.Fatal("expected error for high port")
	}
}

func TestWriteTo_Nil(t *testing.T) {
	n, err := writeTo(false, Log, nil, []byte{1, 2, 3}, netip.AddrPort{})
	if err == nil {
		t.Fatal("expected error for nil conn")
	}
	if n != 0 {
		t.Fatalf("expected 0 bytes, got %d", n)
	}
}

func TestCanReadConn_Nil(t *testing.T) {
	n, err := canReadConn(false, Log, nil)
	if err == nil {
		t.Fatal("expected error for nil conn")
	}
	if n != 0 {
		t.Fatalf("expected 0, got %d", n)
	}
}

func TestCanReadConn_InvalidType(t *testing.T) {
	// mock that doesn't implement syscall.Conn should panic
	defer func() {
		if r := recover(); r == nil {
			t.Fatal("expected panic for invalid type")
		}
	}()
	var pc net.PacketConn = &mockPacketConn{}
	_, _ = canReadConn(false, Log, pc)
}

func TestWaitForLobbyResults(t *testing.T) {
	// nil conn with RecvCanRead should return error from canReadConn
	_, err := WaitForLobbyResults(nil, netip.Addr{}, RecvCanRead, LobbyWaitOptions{})
	if err == nil {
		t.Fatal("expected error for nil conn")
	}

	// No RecvCanRead flag should set argp=1 and then try readFrom with nil conn
	// This will panic or return error - we just verify it doesn't panic unexpectedly
	// Actually readFrom with nil conn will panic, so we skip this case
}

type mockPacketConn struct{}

func (m *mockPacketConn) ReadFrom(p []byte) (n int, addr net.Addr, err error) { return }
func (m *mockPacketConn) WriteTo(p []byte, addr net.Addr) (n int, err error)  { return }
func (m *mockPacketConn) Close() error                                        { return nil }
func (m *mockPacketConn) LocalAddr() net.Addr                                 { return nil }
func (m *mockPacketConn) SetDeadline(t time.Time) error                       { return nil }
func (m *mockPacketConn) SetReadDeadline(t time.Time) error                   { return nil }
func (m *mockPacketConn) SetWriteDeadline(t time.Time) error                  { return nil }
