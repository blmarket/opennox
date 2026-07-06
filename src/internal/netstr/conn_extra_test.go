package netstr

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/server/netlib"
	"github.com/stretchr/testify/require"
)

func TestConn_Data2Hdr(t *testing.T) {
	var ns *Conn
	hdr := ns.data2hdr()
	require.NotNil(t, hdr)

	ns = &Conn{}
	ns.sendBuf = make([]byte, 10)
	hdr = ns.data2hdr()
	require.NotNil(t, hdr)
}

func TestConn_TransferStats(t *testing.T) {
	var ns *Conn
	val := ns.TransferStats()
	require.Equal(t, uint32(0), val)

	ns = &Conn{g: &Streams{}}
	ns.ind = 1
	ns.g.transfer[1] = 100
	val = ns.TransferStats()
	require.Equal(t, uint32(100), val)
	require.Equal(t, uint32(0), ns.g.transfer[1])
}

func TestConn_AddTransferStats(t *testing.T) {
	ns := &Conn{g: &Streams{}}
	ns.ind = 2
	ns.addTransferStats(50)
	require.Equal(t, uint32(50), ns.g.transfer[2])
	ns.addTransferStats(25)
	require.Equal(t, uint32(75), ns.g.transfer[2])
}

func TestConn_Reset(t *testing.T) {
	var ns *Conn
	ns.reset() // should not panic

	ns = &Conn{g: &Streams{}, ind: 1, sendBuf: make([]byte, 10)}
	ns.reset()
	require.NotNil(t, ns.g)
	require.Equal(t, 1, ns.ind)
}

func TestConn_CallOnReceive(t *testing.T) {
	ns := &Conn{}
	val := ns.callOnReceive(nil, nil)
	require.Equal(t, 0, val)

	called := false
	ns.onReceive = func(conn netlib.StreamID, buf []byte) int {
		called = true
		return 42
	}
	val = ns.callOnReceive(nil, []byte{1, 2, 3})
	require.True(t, called)
	require.Equal(t, 42, val)
}

func TestConn_Close(t *testing.T) {
	var ns *Conn
	err := ns.Close()
	require.NoError(t, err)

	ns = &Conn{g: &Streams{}}
	err = ns.Close()
	require.NoError(t, err)
}
