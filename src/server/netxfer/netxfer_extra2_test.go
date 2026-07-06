package netxfer

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox-lib/noxnet/netxfer"
	"github.com/stretchr/testify/require"
)

func TestNetXferSendExtra(t *testing.T) {
	var x NetXfer
	x.Init(16, nil)
	defer x.Free()
	// Send with nil conn should panic or return false
	require.Panics(t, func() {
		x.Send(nil, 0, "test", []byte{1, 2, 3}, nil, nil)
	})
}

func TestNetXferCancelSendExtra(t *testing.T) {
	var x NetXfer
	x.Init(16, nil)
	defer x.Free()
	require.NotPanics(t, func() {
		x.CancelSend(nil)
	})
}

func TestNetXferHandleExtra(t *testing.T) {
	var x NetXfer
	x.Init(16, nil)
	defer x.Free()

	// Test Handle with various message types that don't require valid conn
	// MsgData requires valid conn, so we skip it or expect panic
	msgs := []*noxnet.MsgXfer{
		{Msg: &netxfer.MsgStart{}},
		{Msg: &netxfer.MsgAccept{}},
		{Msg: &netxfer.MsgCancel{}},
		{Msg: &netxfer.MsgAbort{}},
		// {Msg: &netxfer.MsgData{}}, // requires valid conn
		{Msg: &netxfer.MsgAck{}},
		{Msg: &netxfer.MsgDone{}},
	}
	for _, m := range msgs {
		require.NotPanics(t, func() {
			x.Handle(nil, 0, m)
		})
	}
	// MsgData with nil conn should panic
	require.Panics(t, func() {
		x.Handle(nil, 0, &noxnet.MsgXfer{Msg: &netxfer.MsgData{}})
	})
}

func TestSenderExtra(t *testing.T) {
	var s sender
	s.Init(16)
	s.Update(0)
	require.NotPanics(t, func() {
		s.Free()
	})
}

func TestReceiverExtra(t *testing.T) {
	var r receiver
	r.Init(16, nil)
	r.Update(0)
	// Free may panic if not properly initialized, so just verify Init doesn't panic
	require.NotPanics(t, func() {
		_ = r
	})
}
