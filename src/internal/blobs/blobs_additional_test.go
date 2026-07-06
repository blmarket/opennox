package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBlobErrorsAdditional(t *testing.T) {
	// Test SplitBlob with non-existent blob - should return error
	// This covers the error path when old blob not found
	// Note: It may fail earlier if blob files don't exist, which is also an error
	err := SplitBlob(0xDEADBEEF, 0, 0)
	require.Error(t, err)
}

func TestFormatAccessesAdditional(t *testing.T) {
	// Test FormatAccesses - it calls RewriteAccess which may fail if no files
	// We just test that it doesn't panic
	_ = FormatAccesses()
	// Error is expected if no source files found, but shouldn't panic
}
