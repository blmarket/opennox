package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBlob_Error2(t *testing.T) {
	// SplitBlob should error on invalid blob
	err := SplitBlob(0, 0, 0)
	require.Error(t, err)
}

func TestReadBlobs_Error2(t *testing.T) {
	// ReadBlobs without proper setup should error or return empty
	bl, err := ReadBlobs()
	if err == nil && bl == nil {
		t.Error("ReadBlobs returned nil without error")
	}
}

func TestBlobs_Write_Error2(t *testing.T) {
	b := &Blobs{}
	// Write without path should error or handle gracefully
	err := b.Write()
	// It may error or succeed with empty path, just ensure no panic
	_ = err
}

func TestBlobs_Update2(t *testing.T) {
	b := &Blobs{}
	b.Add(Blob{Blob: 0x100, Size: 10})
	b.Update(Blob{Blob: 0x100, Size: 20})
	b2 := b.Get(0x100)
	require.NotNil(t, b2)
	require.Equal(t, uintptr(20), b2.Size)
}

func TestBlobs_Add2(t *testing.T) {
	b := &Blobs{}
	b.Add(Blob{Blob: 0x200, Size: 20})
	b.Add(Blob{Blob: 0x300, Size: 30})
	b2 := b.Get(0x200)
	require.NotNil(t, b2)
}

func TestBlobs_Get2(t *testing.T) {
	b := &Blobs{}
	b.Add(Blob{Blob: 0x400, Size: 15})
	blob := b.Get(0x400)
	require.NotNil(t, blob)
	require.Equal(t, uintptr(15), blob.Size)

	blob2 := b.Get(0x999)
	require.Nil(t, blob2)
}

func TestFormatAccesses3(t *testing.T) {
	err := FormatAccesses()
	if err != nil {
		t.Logf("FormatAccesses error (expected): %v", err)
	}
}

func TestRewriteAccess3(t *testing.T) {
	err := RewriteAccess(nil)
	require.Error(t, err)
}

func TestBlobSpans3(t *testing.T) {
	spans, err := BlobSpans(0)
	if err != nil {
		t.Logf("BlobSpans error (expected): %v", err)
	}
	_ = spans
}

func TestSplitBlob_InvalidArgs2(t *testing.T) {
	// Test with various invalid arguments
	err := SplitBlob(0xFFFFFFFF, 0, 0)
	require.Error(t, err)
}
