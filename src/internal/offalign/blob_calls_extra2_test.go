package main

import (
	"testing"
)

func TestBlobCallsExtra(t *testing.T) {
	data := []byte("getMemU8(123, 4) + getMem32(0x100, 20)")
	calls := BlobCalls(data)
	if len(calls) != 2 {
		t.Errorf("BlobCalls should find 2 calls, got %d", len(calls))
	}
}

func TestBlobCallParseOffsExtra(t *testing.T) {
	call := BlobCall{
		OffExpr: []byte("123+456"),
	}
	_, ok := call.ParseOffs()
	if !ok {
		t.Error("BlobCall.ParseOffs should succeed with valid expression")
	}
}

func TestBlobOffSumMethodsExtra(t *testing.T) {
	s := BlobOffSum{Static: 1, Raw: []string{"a", "b"}}
	if s.Parts() != 3 {
		t.Errorf("BlobOffSum.Parts() = %d, want 3", s.Parts())
	}

	s1 := BlobOffSum{Static: 1, Raw: []string{"a"}}
	s2 := BlobOffSum{Static: 2, Raw: []string{"b"}}
	s1.Merge(s2)
	if s1.Static != 3 {
		t.Errorf("After merge, static = %d, want 3", s1.Static)
	}
}
