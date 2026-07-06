package legacy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerScriptFileReadWriteInt(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test_int.bin")

	// Write integers
	f := C_nox_fs_open_write(path)
	require.NotNil(t, f)
	require.Equal(t, 1, C_nox_script_writeInt(42, f))
	require.Equal(t, 1, C_nox_script_writeInt(-123, f))
	require.Equal(t, 1, C_nox_script_writeInt(0, f))
	require.Equal(t, 1, C_nox_script_writeInt(2147483647, f))
	C_nox_fs_close(f)

	// Read integers back
	f = C_nox_fs_open_read(path)
	require.NotNil(t, f)
	require.Equal(t, 42, C_nox_script_readInt(f))
	require.Equal(t, -123, C_nox_script_readInt(f))
	require.Equal(t, 0, C_nox_script_readInt(f))
	require.Equal(t, 2147483647, C_nox_script_readInt(f))
	C_nox_fs_close(f)
}

func TestServerScriptFileReadWriteFloat(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test_float.bin")

	// Write floats
	f := C_nox_fs_open_write(path)
	require.NotNil(t, f)
	require.Equal(t, 1, C_nox_script_writeFloat(3.14, f))
	require.Equal(t, 1, C_nox_script_writeFloat(-2.5, f))
	require.Equal(t, 1, C_nox_script_writeFloat(0.0, f))
	C_nox_fs_close(f)

	// Read floats back
	f = C_nox_fs_open_read(path)
	require.NotNil(t, f)
	require.InDelta(t, 3.14, C_nox_script_readFloat(f), 0.001)
	require.InDelta(t, -2.5, C_nox_script_readFloat(f), 0.001)
	require.InDelta(t, 0.0, C_nox_script_readFloat(f), 0.001)
	C_nox_fs_close(f)
}

func TestServerScriptFileReadWriteMixed(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test_mixed.bin")

	// Write mixed data
	f := C_nox_fs_open_write(path)
	require.NotNil(t, f)
	require.Equal(t, 1, C_nox_script_writeInt(100, f))
	require.Equal(t, 1, C_nox_script_writeFloat(1.5, f))
	require.Equal(t, 1, C_nox_script_writeInt(200, f))
	require.Equal(t, 1, C_nox_script_writeFloat(2.5, f))
	C_nox_fs_close(f)

	// Read mixed data back
	f = C_nox_fs_open_read(path)
	require.NotNil(t, f)
	require.Equal(t, 100, C_nox_script_readInt(f))
	require.InDelta(t, 1.5, C_nox_script_readFloat(f), 0.001)
	require.Equal(t, 200, C_nox_script_readInt(f))
	require.InDelta(t, 2.5, C_nox_script_readFloat(f), 0.001)
	C_nox_fs_close(f)

	// Verify file contents directly
	data, err := os.ReadFile(path)
	require.NoError(t, err)
	require.Equal(t, 16, len(data)) // 4 ints/floats * 4 bytes each
}
