package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBytes(t *testing.T) {
	b := &Blobs{}

	// Empty
	out, err := b.parseBytes([]byte{}, 0)
	require.NoError(t, err)
	require.Nil(t, out)

	// Single zero
	out, err = b.parseBytes([]byte("0"), 1)
	require.NoError(t, err)
	require.Nil(t, out) // {0} becomes nil

	// Simple bytes
	out, err = b.parseBytes([]byte("0x1, 0x2, 0x3"), 3)
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2, 3}, out)

	// With 'u' suffix
	out, err = b.parseBytes([]byte("1u, 2u, 3u"), 3)
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2, 3}, out)

	// Decimal and hex mixed
	out, err = b.parseBytes([]byte("10, 0xA, 255"), 3)
	require.NoError(t, err)
	require.Equal(t, []byte{10, 10, 255}, out)

	// Invalid
	_, err = b.parseBytes([]byte("xyz"), 1)
	require.Error(t, err)
}

func TestReadBlobs(t *testing.T) {
	// Set path to temp dir to avoid reading real files
	SetPath(t.TempDir())
	_, err := ReadBlobs()
	// Should fail because files don't exist, but shouldn't panic
	require.Error(t, err)
}

func TestSplitBlobError(t *testing.T) {
	SetPath(t.TempDir())
	err := SplitBlob(0x1234, 10, 5)
	require.Error(t, err) // files don't exist
}

func TestReadMemmap(t *testing.T) {
	SetPath(t.TempDir())
	_, err := ReadMemmap()
	require.Error(t, err)
}
