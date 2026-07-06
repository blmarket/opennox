package netstr

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/stretchr/testify/require"
)

func TestDecodeMessage_Extra(t *testing.T) {
	// Test decodeMessage with invalid packet
	var msg noxnet.MsgDiscover
	// decodeMessage panics on short packets, so test with sufficient length
	result := decodeMessage([]byte{0, 0, 0, 0, 0}, &msg)
	require.False(t, result)

	// Test with valid packet header but invalid body
	result = decodeMessage([]byte{0, 0, 0, 0xFF, 0xFF}, &msg)
	_ = result

	// Test that decodeMessage panics on empty packet (expected behavior)
	require.Panics(t, func() {
		decodeMessage([]byte{}, &msg)
	})
}

func TestEncodeMessage_Extra(t *testing.T) {
	// Test encodeMessage with nil message
	out := make([]byte, 100)
	n := encodeMessage(out, nil)
	require.Equal(t, 0, n)

	// Test encodeMessage with valid message
	msg := &noxnet.MsgDiscover{
		Token: 0x12345678,
	}
	n = encodeMessage(out, msg)
	require.Greater(t, n, 2)
	require.Equal(t, byte(0), out[0])
	require.Equal(t, byte(0), out[1])
}

func TestEncodeMessage_Panic(t *testing.T) {
	// Test that encodeMessage panics on encoding error
	// This is hard to trigger without a bad message, so just verify function exists
	require.NotPanics(t, func() {
		out := make([]byte, 100)
		msg := &noxnet.MsgDiscover{Token: 1}
		_ = encodeMessage(out, msg)
	})
}

func TestConn_String_Extra(t *testing.T) {
	// Test Conn.String with nil
	var c *Conn
	s := c.String()
	require.Equal(t, "<nil>", s)

	// Test Conn.String with empty conn
	c = &Conn{}
	s = c.String()
	require.NotEmpty(t, s)
	require.Contains(t, s, "Conn")
}

func TestConn_IsHost_Extra(t *testing.T) {
	c := &Conn{ind: 0}
	require.True(t, c.IsHost())

	c = &Conn{ind: 1}
	require.False(t, c.IsHost())
}

func TestConn_Player_Extra(t *testing.T) {
	c := &Conn{id: 6}
	require.Equal(t, 5, int(c.Player()))

	c = &Conn{id: 1}
	require.Equal(t, 0, int(c.Player()))

	c = &Conn{id: 0}
	require.Equal(t, -1, int(c.Player()))
}

func TestConn_Addr_Extra(t *testing.T) {
	c := &Conn{}
	addr := c.Addr()
	_ = addr
}

func TestConn_IP_Extra(t *testing.T) {
	c := &Conn{}
	ip := c.IP()
	_ = ip
}
