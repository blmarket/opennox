package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlobAccesses(t *testing.T) {
	// Test FormatAccesses
	err := FormatAccesses()
	// Should fail because files don't exist, but shouldn't panic
	_ = err

	// Test BlobSpans
	spans, err := BlobSpans(0x1000)
	// Should fail because files don't exist, but shouldn't panic
	_ = spans
	_ = err
}

func TestSplitBlobValidation(t *testing.T) {
	SetPath(t.TempDir())

	// Test with non-existent blob
	err := SplitBlob(0xFFFFFFFF, 0, 0)
	require.Error(t, err)
}

func TestReadWriteMemmap(t *testing.T) {
	SetPath(t.TempDir())

	// Test ReadMemmap with non-existent file
	_, err := ReadMemmap()
	require.Error(t, err)

	// Test Write with empty memmap
	m := &Mapping{}
	err = m.Write()
	// May fail because path not set properly, but shouldn't panic
	_ = err
}

func TestBlobsWrite(t *testing.T) {
	SetPath(t.TempDir())

	b := &Blobs{}
	err := b.Write()
	// Should fail because no path, but shouldn't panic
	_ = err
}

func TestBlobsGet(t *testing.T) {
	b := &Blobs{}

	// Get non-existent
	bl := b.Get(0x9999)
	require.Nil(t, bl)
}

func TestBlobsUpdate(t *testing.T) {
	b := &Blobs{}

	// Update non-existent should not panic
	b.Update(Blob{Blob: 0x9999, Size: 50})
}

func TestBlobsAdd(t *testing.T) {
	b := &Blobs{}
	b.Add(Blob{Blob: 0x1000, Size: 100})
	b.Add(Blob{Blob: 0x2000, Size: 200})
	// Should not panic
}

func TestRewriteAccess(t *testing.T) {
	SetPath(t.TempDir())

	// Test with no files
	err := RewriteAccess(func(a *Access) (bool, error) {
		return false, nil
	})
	// Should fail because files don't exist, but shouldn't panic
	_ = err
}
