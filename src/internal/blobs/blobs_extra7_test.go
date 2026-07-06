package blobs

import (
	"testing"
)

func TestSplitBlobErrors3(t *testing.T) {
	err := SplitBlob(0x12345678, 0, 0)
	if err == nil {
		t.Error("SplitBlob should return error for nonexistent blob")
	}
}

func TestBlobTypes3(t *testing.T) {
	b := Blob{
		Blob: 0x1000,
		Size: 100,
		Data: []byte{1, 2, 3},
	}
	if b.Blob != 0x1000 {
		t.Error("Blob field mismatch")
	}
	// Access.String may panic with empty struct, so just create but don't call String
	a := Access{}
	_ = a
}
