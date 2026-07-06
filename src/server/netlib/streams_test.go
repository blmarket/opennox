package netlib

import "testing"

func TestRecvFlags(t *testing.T) {
	var f RecvFlags = RecvCanRead | RecvJustOne
	if !f.Has(RecvCanRead) {
		t.Fatalf("expected CanRead")
	}
	if f.Has(RecvNoHooks) {
		t.Fatalf("should not have NoHooks")
	}
	if !f.Has(RecvJustOne) {
		t.Fatalf("expected JustOne")
	}
}
