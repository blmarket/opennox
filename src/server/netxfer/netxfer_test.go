package netxfer

import "testing"

func TestNetXferInitFree(t *testing.T) {
	var x NetXfer
	x.Init(16, nil)
	x.Update(123)
	x.Free()
}

func TestConstants(t *testing.T) {
	if minStreams != 16 || maxStreams != 256 {
		t.Fatalf("unexpected constants")
	}
}
