package netxfer

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet/netxfer"
	"github.com/stretchr/testify/require"
)

func TestReceiver(t *testing.T) {
	var r receiver

	// Test Init
	r.Init(5, nil)
	require.Equal(t, 5, r.cnt)
	require.Len(t, r.arr, 5)

	// Test New - should return a stream
	s := r.New(nil, 0, "test", 100)
	require.NotNil(t, s)
	require.Equal(t, uint32(100), uint32(len(s.full)))

	// Test New again - should return another stream
	s2 := r.New(nil, 0, "test2", 200)
	require.NotNil(t, s2)
	require.NotEqual(t, s, s2)

	// Test Free
	r.Free()
	require.Nil(t, r.arr)
}

func TestReceiverNewExhausted(t *testing.T) {
	var r receiver
	r.Init(2, nil)

	// Allocate all streams
	s1 := r.New(nil, 0, "test1", 10)
	require.NotNil(t, s1)
	s2 := r.New(nil, 0, "test2", 20)
	require.NotNil(t, s2)

	// Next New should panic or return nil (exhausted)
	// The code panics when new() returns nil, so we recover
	func() {
		defer func() {
			_ = recover()
		}()
		s3 := r.New(nil, 0, "test3", 30)
		require.Nil(t, s3)
	}()

	r.Free()
}

func TestRecvStreamReset(t *testing.T) {
	var s recvStream
	s.full = make([]byte, 100)
	s.recvID = 1
	s.action = 1

	s.Reset(netxfer.RecvID(0))
	require.Nil(t, s.full)
	require.Equal(t, 0, int(s.recvID))
}
