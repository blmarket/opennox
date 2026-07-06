package client

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
)

func TestClient_OnClientPacketOpSub(t *testing.T) {
	c := &Client{}

	// Test MSG_SERVER_CLOSE_ACK
	n, handled, err := c.OnClientPacketOpSub(ntype.PlayerInd(0), noxnet.MSG_SERVER_CLOSE_ACK, []byte{byte(noxnet.MSG_SERVER_CLOSE_ACK)}, nil, nil, CurPlayerInfo{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !handled || n != 1 {
		t.Fatalf("expected handled with n=1, got n=%d handled=%v", n, handled)
	}

	// Test unknown op
	n, handled, err = c.OnClientPacketOpSub(ntype.PlayerInd(0), noxnet.Op(255), []byte{255}, nil, nil, CurPlayerInfo{})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if handled {
		t.Fatal("expected not handled for unknown op")
	}
	if n != 0 {
		t.Fatalf("expected n=0, got %d", n)
	}

	// Test MSG_STAT_MULTIPLIERS with invalid data
	n, handled, err = c.OnClientPacketOpSub(ntype.PlayerInd(0), noxnet.MSG_STAT_MULTIPLIERS, []byte{byte(noxnet.MSG_STAT_MULTIPLIERS)}, nil, nil, CurPlayerInfo{})
	if err == nil {
		t.Fatal("expected error for invalid data")
	}

	// Test MSG_DESTROY_WALL with invalid data
	n, handled, err = c.OnClientPacketOpSub(ntype.PlayerInd(0), noxnet.MSG_DESTROY_WALL, []byte{byte(noxnet.MSG_DESTROY_WALL)}, nil, nil, CurPlayerInfo{})
	if err == nil {
		t.Fatal("expected error for invalid data")
	}

	// Test MSG_FX_JIGGLE with invalid data
	n, handled, err = c.OnClientPacketOpSub(ntype.PlayerInd(0), noxnet.MSG_FX_JIGGLE, []byte{byte(noxnet.MSG_FX_JIGGLE)}, nil, nil, CurPlayerInfo{})
	if err == nil {
		t.Fatal("expected error for invalid data")
	}

	// Test MSG_XFER_MSG with invalid data
	n, handled, err = c.OnClientPacketOpSub(ntype.PlayerInd(0), noxnet.MSG_XFER_MSG, []byte{byte(noxnet.MSG_XFER_MSG)}, nil, nil, CurPlayerInfo{})
	if err == nil {
		t.Fatal("expected error for invalid data")
	}
}
