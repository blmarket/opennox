package netstr

import (
	"net/netip"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetForPacketNil(t *testing.T) {
	var g *Streams
	result := g.getForPacket(0)
	require.Equal(t, handle{}, result)
}

func TestStruct2IndByAddrEmpty(t *testing.T) {
	var g *Streams
	// This panics on nil Streams, so we just verify the function exists
	// by checking it panics as expected
	require.Panics(t, func() {
		g.struct2IndByAddr(nilAddrPort())
	})
}

func TestHasConnWithIndAndAddrEmpty(t *testing.T) {
	var g *Streams
	var h handle
	require.Panics(t, func() {
		g.hasConnWithIndAndAddr(h, nilAddrPort())
	})
}

func nilAddrPort() netip.AddrPort {
	return netip.AddrPort{}
}
