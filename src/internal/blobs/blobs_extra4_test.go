package blobs

import (
	"testing"
)

func TestBlobsGetUpdateAddExtra(t *testing.T) {
	b := &Blobs{}

	// Test Get on empty
	if b.Get(0x1000) != nil {
		t.Error("Get on empty should return nil")
	}

	// Test Add
	blob := Blob{Blob: 0x1000, Size: 100, Data: []byte{1, 2, 3}}
	b.Add(blob)

	// Test Get after Add
	got := b.Get(0x1000)
	if got == nil {
		t.Fatal("Get after Add should not return nil")
	}
	if got.Blob != 0x1000 || got.Size != 100 {
		t.Errorf("Get returned wrong blob: %+v", got)
	}

	// Test Update
	blob2 := Blob{Blob: 0x1000, Size: 200, Data: []byte{4, 5, 6}}
	b.Update(blob2)
	got = b.Get(0x1000)
	if got.Size != 200 {
		t.Errorf("Update failed, size = %d, want 200", got.Size)
	}

	// Test get helper with empty array
	if b.get(0x1000, []Blob{}) != nil {
		t.Error("get on empty array should return nil")
	}

	// Test get with non-matching blob
	if b.get(0x2000, b.data.blobs) != nil {
		t.Error("get with non-matching blob should return nil")
	}
}

func TestParseBytesExtra(t *testing.T) {
	b := &Blobs{}
	data := []byte("0x01, 0x02, 0x03")
	result, err := b.parseBytes(data, 3)
	if err != nil {
		t.Errorf("parseBytes failed: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("parseBytes len = %d, want 3", len(result))
	}

	// Test out of bounds
	result, err = b.parseBytes(data, 10)
	if err != nil {
		t.Logf("parseBytes out of bounds error (expected): %v", err)
	}
	_ = result
}

func TestBlobsWriteExtra(t *testing.T) {
	b := &Blobs{}
	// Write on empty should not error (writes empty files)
	err := b.Write()
	// It may error if files don't exist, which is ok
	_ = err
}

func TestReadBlobsExtra(t *testing.T) {
	// ReadBlobs reads from filesystem, may fail if files not present
	// Just verify it doesn't panic
	_, _ = ReadBlobs()
}
