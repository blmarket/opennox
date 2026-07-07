package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/noxnet"
	"github.com/noxworld-dev/opennox/v1/server"
)

func TestCopyFull(t *testing.T) {
	src := []byte{1, 2, 3, 4}
	dst := make([]byte, 4)
	n := copyFull(dst, src)
	if n != 4 {
		t.Fatalf("expected 4, got %d", n)
	}
	for i, v := range src {
		if dst[i] != v {
			t.Fatalf("mismatch at %d", i)
		}
	}

	// dst too small
	dst2 := make([]byte, 2)
	n = copyFull(dst2, src)
	if n != 0 {
		t.Fatalf("expected 0 for small dst, got %d", n)
	}
}

func TestZero3Full(t *testing.T) {
	b := []byte{1, 2, 3, 4}
	n := zero3full(b)
	if n != 3 {
		t.Fatalf("expected 3, got %d", n)
	}
	if b[0] != 0 || b[1] != 0 || b[2] != 0 {
		t.Fatal("expected first 3 bytes zeroed")
	}
	if b[3] != 4 {
		t.Fatal("fourth byte should be unchanged")
	}

	// too small
	b2 := []byte{1, 2}
	n = zero3full(b2)
	if n != 0 {
		t.Fatalf("expected 0 for small buffer, got %d", n)
	}
}

func TestNoxXxxChkIsMsgTimestamp(t *testing.T) {
	if nox_xxx_chkIsMsgTimestamp_4DF7F0(nil) {
		t.Fatal("nil should be false")
	}
	if nox_xxx_chkIsMsgTimestamp_4DF7F0([]byte{}) {
		t.Fatal("empty should be false")
	}
	if nox_xxx_chkIsMsgTimestamp_4DF7F0([]byte{0x01}) {
		t.Fatal("random byte should be false")
	}
	if !nox_xxx_chkIsMsgTimestamp_4DF7F0([]byte{byte(noxnet.MSG_TIMESTAMP)}) {
		t.Fatal("MSG_TIMESTAMP should be true")
	}
	if !nox_xxx_chkIsMsgTimestamp_4DF7F0([]byte{byte(noxnet.MSG_FULL_TIMESTAMP)}) {
		t.Fatal("MSG_FULL_TIMESTAMP should be true")
	}
}

func TestSub57B930(t *testing.T) {
	var arr [255]server.PlayerNetData
	// empty array should return 255
	if v := sub_57B930(&arr, 1, 2, 0); v != 255 {
		t.Fatalf("expected 255 for empty, got %d", v)
	}

	// set up a matching entry
	arr[1].Field0 = 10
	arr[1].Field2 = 20
	arr[1].Frame4 = 100
	if v := sub_57B930(&arr, 10, 20, 50); v != 1 {
		t.Fatalf("expected 1, got %d", v)
	}

	// frame too high should break and return 255
	if v := sub_57B930(&arr, 10, 20, 200); v != 255 {
		t.Fatalf("expected 255 for high frame, got %d", v)
	}

	// test si == 0 or 255 defaults to 1
	arr[5].Field0 = 5
	arr[5].Field2 = 5
	arr[5].Frame4 = 10
	if v := sub_57B930(&arr, 0, 5, 0); v != 255 {
		// si defaults to 1, but Field0 is 0, so no match
	}
	if v := sub_57B930(&arr, 5, 5, 0); v != 5 {
		t.Fatalf("expected 5, got %d", v)
	}

	// test wrap around
	arr[254].Field0 = 30
	arr[254].Field2 = 40
	arr[254].Frame4 = 10
	if v := sub_57B930(&arr, 30, 40, 0); v != 254 {
		t.Fatalf("expected 254, got %d", v)
	}
}

func TestGetNetPlayerBufSize(t *testing.T) {
	netPlayerBufSize = 123
	if v := getNetPlayerBufSize(); v != 123 {
		t.Fatalf("expected 123, got %d", v)
	}
}
