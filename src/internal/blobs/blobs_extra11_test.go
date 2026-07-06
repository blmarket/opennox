package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBlobExtra(t *testing.T) {
	SetPath(t.TempDir())
	err := SplitBlob(0x1234, 10, 5)
	require.Error(t, err)
}

func TestReadBlobsExtra2(t *testing.T) {
	SetPath(t.TempDir())
	_, err := ReadBlobs()
	require.Error(t, err)
}

func TestWriteExtra(t *testing.T) {
	SetPath(t.TempDir())
	b := &Blobs{}
	err := b.Write()
	_ = err
}

func TestFormatAccessesExtra(t *testing.T) {
	SetPath(t.TempDir())
	_ = FormatAccesses()
}

func TestRewriteAccessExtra(t *testing.T) {
	SetPath(t.TempDir())
	err := RewriteAccess(func(a *Access) (bool, error) {
		return false, nil
	})
	_ = err
}

func TestReadDataExtra(t *testing.T) {
	SetPath(t.TempDir())
	b := &Blobs{}
	err := b.readData()
	require.Error(t, err)
}

func TestWriteDataExtra(t *testing.T) {
	SetPath(t.TempDir())
	b := &Blobs{}
	_ = b.writeData()
}
