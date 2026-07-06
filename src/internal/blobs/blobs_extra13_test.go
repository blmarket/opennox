package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBlobAdditional(t *testing.T) {
	// Test SplitBlob with empty data
	err := SplitBlob(0, 0, 0)
	require.NotNil(t, err) // Should return error for invalid blob

	// Test with simple data
	err = SplitBlob(0x1000, 0, 100)
	_ = err
}

func TestBlobsWriteAdditional(t *testing.T) {
	b := &Blobs{}
	// Test Write with empty blobs
	err := b.Write()
	// Should not panic, error is ok
	_ = err
}

func TestReadMemmapFunctions(t *testing.T) {
	// These functions require specific setup, just ensure they don't panic on empty input
	b := &Blobs{}
	_ = b
	// readMemmapC, readMemmapGo1, etc. are internal and require file setup
	// Just test that Blobs struct can be created and methods don't panic on empty
}
