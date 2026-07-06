package netlist

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/stretchr/testify/require"
)

func TestCheckSizes(t *testing.T) {
	l := New()
	l.Init()
	defer l.Free()

	// These are internal but we can trigger via public APIs
	// Add data to trigger size checks
	l.AddToMsgListCli(0, Kind1, []byte{1, 2, 3, 4, 5})
	out := l.CopyPacketsA(0, Kind1)
	require.NotNil(t, out)

	l.AddToMsgListSrv(0, []byte{1, 2}, func(ind ntype.PlayerInd) {})
}

func TestMsgListEachBreak(t *testing.T) {
	ml := newMsgList()
	ml.add([]byte{1}, true)
	ml.add([]byte{2}, true)
	ml.add([]byte{3}, true)

	cnt := 0
	ml.Each(func(b []byte) bool {
		cnt++
		return cnt >= 2 // break after 2
	})
	require.Equal(t, 2, cnt)
	ml.Free()
}

func TestMsgListCountSizeEmpty(t *testing.T) {
	ml := newMsgList()
	require.Equal(t, 0, ml.Count())
	require.Equal(t, 0, ml.Size())
	require.Nil(t, ml.Get())
	ml.Free()
}
