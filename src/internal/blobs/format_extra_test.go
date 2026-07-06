package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFormatAccesses_Extra(t *testing.T) {
	// Test FormatAccesses - it should not panic even if no files
	err := FormatAccesses()
	// May return error if no files found, or nil if successful
	_ = err
}

func TestFormatAccesses_WithPath(t *testing.T) {
	// Test FormatAccesses with a specific path
	// Set a path that likely doesn't exist or has no blob accesses
	originalPath := Path()
	defer SetPath(originalPath)

	SetPath("/tmp/nonexistent_blobs_path")
	err := FormatAccesses()
	// Should not panic, may return error
	_ = err
	require.NotPanics(t, func() {
		_ = FormatAccesses()
	})
}
