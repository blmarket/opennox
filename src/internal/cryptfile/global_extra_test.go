package cryptfile

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGlobal_Extra(t *testing.T) {
	// Test Global and SetGlobal
	require.Nil(t, Global())

	cf := &CryptFile{}
	SetGlobal(cf)
	require.Equal(t, cf, Global())

	SetGlobal(nil)
	require.Nil(t, Global())

	// Test Close with nil global
	require.NoError(t, Close())

	// Test OpenGlobal with invalid path
	err := OpenGlobal("/nonexistent/path/file.bin", 0, 0)
	require.Error(t, err)
	require.Nil(t, Global())

	// Test OpenGlobal with valid file
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.bin")

	// Create a simple crypt file
	cf, err = OpenFile(path, WriteOnly, -1)
	require.NoError(t, err)
	require.NotNil(t, cf)
	cf.WriteString8("test")
	cf.Close()

	// Test OpenGlobal
	err = OpenGlobal(path, ReadOnly, -1)
	require.NoError(t, err)
	require.NotNil(t, Global())

	// Test Close
	require.NoError(t, Close())
	require.Nil(t, Global())

	// Test Close again with nil
	require.NoError(t, Close())

	// Cleanup
	os.Remove(path)
}
