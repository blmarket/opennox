package memguard

import "testing"

func TestMemguard(t *testing.T) {
	ps := PageSize()
	if ps <= 0 {
		t.Fatalf("invalid page size %d", ps)
	}
	data, free := New(ps * 2)
	if len(data) == 0 {
		t.Fatalf("empty data")
	}
	if free == nil {
		t.Fatalf("nil free")
	}
	free()
}
