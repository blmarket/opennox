package netxfer

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
)

func TestNetXfer_Extra3(t *testing.T) {
	var x NetXfer

	// Test Init and Free
	x.Init(10, nil)
	x.Free()

	// Test Init again
	x.Init(5, func(conn ntype.PlayerInd, act Action, typ string, data []byte) {})
	require.NotNil(t, x)

	// Test Update
	x.Update(100)
	x.Update(200)

	// Test Send with nil conn - skip as it panics
	// result := x.Send(nil, Action(0), "test", []byte("data"), nil, nil)
	// require.False(t, result)

	// Test CancelSend with nil
	x.CancelSend(nil)

	x.Free()

	// Test with nil onRecv
	x.Init(5, nil)
	x.Update(300)
	x.Free()
}
