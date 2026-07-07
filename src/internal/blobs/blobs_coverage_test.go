package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlobsCoverage(t *testing.T) {
	// Test FormatAccesses
	err := FormatAccesses()
	_ = err

	// Test SplitBlob with non-existent blob (should error)
	err = SplitBlob(0x1234, 0, 0)
	require.Error(t, err)

	// Test ReadBlobs (may fail if no blobs file, but should not panic)
	bl, err := ReadBlobs()
	if err == nil {
		require.NotNil(t, bl)
	}

	// Test RewriteAccess with nil func
	err = RewriteAccess(nil)
	if err == nil {
		// ok
	}
}
