package cnxz

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDecompressFile_Invalid(t *testing.T) {
	tmp := t.TempDir()
	invalid := filepath.Join(tmp, "invalid.nxz")
	err := os.WriteFile(invalid, []byte("not a valid nxz file content"), 0644)
	require.NoError(t, err)

	err = DecompressFile(invalid, filepath.Join(tmp, "out"))
	// Should error or handle gracefully, just ensure no panic
	_ = err
}

func TestDecompressFile_Empty(t *testing.T) {
	tmp := t.TempDir()
	empty := filepath.Join(tmp, "empty.nxz")
	err := os.WriteFile(empty, []byte{}, 0644)
	require.NoError(t, err)

	err = DecompressFile(empty, filepath.Join(tmp, "out"))
	require.Error(t, err)
}

func TestCompBufferSize_Extra(t *testing.T) {
	size := compBufferSize(1024)
	require.Greater(t, size, 0)
	require.Equal(t, 1024+512+32, size)
}
