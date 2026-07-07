package netstr

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
)

func TestEncodeDecodeMessage(t *testing.T) {
	out := make([]byte, 100)
	n := encodeMessage(out, nil)
	if n != 0 {
		t.Error("expected 0 for nil message")
	}

	msg := &noxnet.MsgDiscover{Token: 12345}
	n = encodeMessage(out, msg)
	if n == 0 {
		t.Error("expected non-zero for valid message")
	}

	dst := &noxnet.MsgDiscover{}
	ok := decodeMessage(out[:n], dst)
	if !ok {
		t.Error("expected decode to succeed")
	}

	// Invalid data with enough length but bad content
	ok = decodeMessage([]byte{1, 2, 3, 4, 5}, dst)
	_ = ok
}
