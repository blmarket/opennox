package netlist

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/stretchr/testify/require"
)

func TestNetList_Extra2(t *testing.T) {
	// Test New
	nl := New()
	require.NotNil(t, nl)
	nl.Free()

	// Test Init and Free
	nl = &List{}
	nl.Init()
	require.NotNil(t, nl)
	nl.Free()

	// Test ResetAll
	nl = New()
	nl.Init()
	nl.ResetAll()
	nl.Free()

	// Test ResetAllInd
	nl = New()
	nl.Init()
	nl.ResetAllInd(0)
	// nl.ResetAllInd(10) // out of range, skip
	nl.Free()

	// Test ByInd
	nl = New()
	nl.Init()
	ml := nl.ByInd(0, Kind0)
	_ = ml
	// _ = nl.ByInd(100, Kind0) // out of range, skip
	nl.Free()

	// Test MsgList methods
	ml = newMsgList()
	require.NotNil(t, ml)
	require.Equal(t, 0, ml.Count())
	require.Equal(t, 0, ml.Size())

	ml.Reset()
	require.Equal(t, 0, ml.Count())

	ml.Free()

	// Test Each with break
	ml = newMsgList()
	count := 0
	ml.Each(func(b []byte) bool {
		count++
		return false
	})
	require.Equal(t, 0, count)
	ml.Free()

	// Test ResetByInd and InitByInd
	nl = New()
	nl.Init()
	nl.ResetByInd(ntype.PlayerInd(0), Kind0)
	nl.InitByInd(ntype.PlayerInd(0))
	nl.Free()
}
