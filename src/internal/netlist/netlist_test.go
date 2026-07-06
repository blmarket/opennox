package netlist

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
)

func TestNew(t *testing.T) {
	l := New()
	require.NotNil(t, l)
}

func TestInitFree(t *testing.T) {
	l := New()
	l.Init()
	l.Free()
}

func TestByInd(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	// ByInd should return non-nil for valid indices after Init
	for i := 0; i < 3; i++ {
		ml := l.ByInd(0, Kind(i))
		// Kind0 for player 0 is nil, others should be non-nil
		if i == 0 {
			require.Nil(t, ml)
		} else {
			require.NotNil(t, ml)
		}
	}
}

func TestReset(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	l.ResetByInd(0, Kind1)
	l.ResetAllInd(Kind1)
	l.ResetAll()
	l.InitByInd(0)
}

func TestAddToMsgListCli(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	// Empty buf should return true
	ok := l.AddToMsgListCli(0, Kind1, []byte{})
	require.True(t, ok)

	// Normal buf
	ok = l.AddToMsgListCli(0, Kind1, []byte{1, 2, 3})
	require.True(t, ok)

	// Copy packets
	out := l.CopyPacketsA(0, Kind1)
	require.NotNil(t, out)
}

func TestHandlePackets(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	l.AddToMsgListCli(0, Kind1, []byte{1, 2, 3})

	called := false
	l.HandlePacketsA(0, Kind1, func(data []byte) {
		called = true
	})
	require.True(t, called)
}

func TestCopyPacketsB(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	out := l.CopyPacketsB(0)
	require.NotNil(t, out)
}

func TestClientSend0(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	ok := l.ClientSend0(0, Kind1, []byte{}, 0)
	require.True(t, ok)

	ok = l.ClientSend0(0, Kind1, []byte{1, 2}, 0)
	require.True(t, ok)
}

func TestAddToMsgListSrv(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	// Empty buf
	ok := l.AddToMsgListSrv(0, []byte{}, func(ind ntype.PlayerInd) {})
	require.True(t, ok)

	// Normal
	ok = l.AddToMsgListSrv(0, []byte{1, 2, 3}, func(ind ntype.PlayerInd) {})
	require.True(t, ok)
}

func TestMsgList(t *testing.T) {
	ml := newMsgList()
	require.NotNil(t, ml)

	require.Equal(t, 0, ml.Count())
	require.Equal(t, 0, ml.Size())

	// Add and get
	ml.add([]byte{1, 2, 3}, true)
	require.Equal(t, 1, ml.Count())
	require.Equal(t, 3, ml.Size())

	buf := ml.Get()
	require.Equal(t, []byte{1, 2, 3}, buf)
	require.Equal(t, 0, ml.Count())

	// Add multiple
	ml.add([]byte{1}, true)
	ml.add([]byte{2}, true)
	ml.add([]byte{3}, false) // prepend

	cnt := 0
	ml.Each(func(b []byte) bool {
		cnt++
		return false
	})
	require.Equal(t, 3, cnt)

	// FindAndFreeBuf
	ml2 := newMsgList()
	ml2.add([]byte{1, 2}, true)
	buf2 := ml2.Get()
	ml2.add(buf2, true)
	ml2.FindAndFreeBuf(buf2)

	ml.Free()
	ml.Reset()
}
