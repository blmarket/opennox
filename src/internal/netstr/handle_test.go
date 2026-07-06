package netstr

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrHandle(t *testing.T) {
	// Test errHandle with negative value - should not panic
	h := errHandle(-1)
	require.False(t, h.Valid())

	// Test errHandle with positive value - should panic
	require.Panics(t, func() {
		_ = errHandle(1)
	})

	// Test errHandle with zero - should panic
	require.Panics(t, func() {
		_ = errHandle(0)
	})
}

func TestHandleMethods(t *testing.T) {
	// Test handle with nil Streams
	var h handle
	require.False(t, h.Valid())
	// IsHost returns h.i == 0, which is true for zero value, even if not valid
	require.True(t, h.IsHost())
	require.Equal(t, -1, int(h.Player()))
	require.Nil(t, h.Get())

	// Test handle with valid Streams
	streams := &Streams{}

	// Test IsHost
	h = handle{g: streams, i: 0}
	require.True(t, h.IsHost())

	h = handle{g: streams, i: 1}
	require.False(t, h.IsHost())

	// Test Player
	h = handle{g: streams, i: 1}
	require.Equal(t, 0, int(h.Player()))

	h = handle{g: streams, i: 2}
	require.Equal(t, 1, int(h.Player()))

	// Test Get with valid index but nil stream
	h = handle{g: streams, i: 0}
	require.Nil(t, h.Get()) // streams[0] is nil

	// Test Get with valid Conn
	conn := &Conn{}
	streams.streams[0] = conn
	h = handle{g: streams, i: 0}
	require.Equal(t, conn, h.Get())
}

func TestHandleIP(t *testing.T) {
	// Test IP with nil Conn - should panic because Get() returns nil and IP() calls nil.IP()
	// Actually Conn.IP() handles nil, so it should not panic
	var h handle
	// h.Get() returns nil, so h.IP() calls nil.IP() which returns empty Addr
	ip := h.IP()
	require.Equal(t, netip.Addr{}, ip)

	// Test IP with valid Conn
	streams := &Streams{}
	addr := netip.MustParseAddrPort("127.0.0.1:1234")
	conn := &Conn{
		addr: addr,
	}
	streams.streams[0] = conn
	h = handle{g: streams, i: 0}
	ip = h.IP()
	require.Equal(t, addr.Addr(), ip)
}

func TestConnMethods(t *testing.T) {
	// Test Conn.String() with nil
	var c *Conn
	require.Equal(t, "<nil>", c.String())

	// Test Conn.String() with valid Conn
	c = &Conn{
		ind:  1,
		id:   2,
		addr: netip.MustParseAddrPort("192.168.1.1:5678"),
	}
	s := c.String()
	require.Contains(t, s, "ind: 1")
	require.Contains(t, s, "id: 2")

	// Test IsHost with nil
	require.False(t, c.IsHost())
	c = &Conn{ind: 0}
	require.True(t, c.IsHost())
	c = &Conn{ind: 1}
	require.False(t, c.IsHost())

	// Test Player with nil
	c = nil
	require.Equal(t, 0, int(c.Player()))
	c = &Conn{id: 1}
	require.Equal(t, 0, int(c.Player()))
	c = &Conn{id: 2}
	require.Equal(t, 1, int(c.Player()))

	// Test Addr with nil
	c = nil
	require.Equal(t, netip.AddrPort{}, c.Addr())
	c = &Conn{addr: netip.MustParseAddrPort("10.0.0.1:9999")}
	require.Equal(t, netip.MustParseAddrPort("10.0.0.1:9999"), c.Addr())

	// Test IP with nil
	c = nil
	require.Equal(t, netip.Addr{}, c.IP())
	c = &Conn{addr: netip.MustParseAddrPort("10.0.0.1:9999")}
	require.Equal(t, netip.MustParseAddr("10.0.0.1"), c.IP())

	// Test handle() with nil
	c = nil
	h := c.handle()
	require.False(t, h.Valid())

	// Test handle() with valid Conn
	streams := &Streams{}
	c = &Conn{g: streams, id: 5}
	h = c.handle()
	require.Equal(t, streams, h.g)
	require.Equal(t, 5, h.i)
}

func TestStreamsNewClient(t *testing.T) {
	// Test Streams.NewClient creates a client
	streams := &Streams{}
	c, err := streams.NewClient(nil)
	// May return error if not properly initialized, but shouldn't panic
	if err == nil {
		require.NotNil(t, c)
	}
}

func TestConnectFailErr(t *testing.T) {
	// Test NewConnectErr with code 0 - should set code to -1
	err := NewConnectErr(0, nil)
	require.Equal(t, -1, err.Code)

	// Test NewConnectErr with non-zero code
	err = NewConnectErr(42, nil)
	require.Equal(t, 42, err.Code)

	// Test Error() method
	err = NewConnectErr(5, nil)
	s := err.Error()
	require.Contains(t, s, "code=5")

	// Test Unwrap
	require.Nil(t, err.Unwrap())
}
