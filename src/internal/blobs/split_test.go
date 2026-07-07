package blobs

import (
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBlob(t *testing.T) {
	if os.Getenv("NOX_DO_SPLIT") == "" {
		t.SkipNow()
	}
	oldPath := blobPath
	SetPath("../../")
	t.Cleanup(func() { SetPath(oldPath) })
	sub := strings.Split(os.Getenv("NOX_DO_SPLIT"), ",")
	blob, err := strconv.ParseUint(sub[0], 0, 64)
	require.NoError(t, err)
	off, err := strconv.ParseUint(sub[1], 0, 64)
	require.NoError(t, err)
	var size uint64
	if len(sub) > 2 {
		size, err = strconv.ParseUint(sub[2], 0, 64)
		require.NoError(t, err)
	}
	err = SplitBlob(uintptr(blob), uintptr(off), uintptr(size))
	require.NoError(t, err)
}

func TestSplitBlob_ErrorCases(t *testing.T) {
	// Test with non-existent blob (should return error from ReadBlobs or "old blob not found")
	err := SplitBlob(0x12345678, 0, 0)
	if err == nil {
		t.Error("SplitBlob with non-existent blob should return error")
	}
}

func TestSplitBlob_InvalidArgs(t *testing.T) {
	// Test with invalid blob address
	err := SplitBlob(0, 0, 0)
	if err == nil {
		t.Error("SplitBlob with blob 0 should return error")
	}
}
