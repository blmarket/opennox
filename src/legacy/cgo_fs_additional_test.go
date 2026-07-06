package legacy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/stretchr/testify/require"
)

func TestCgoFsFprintfAdditional(t *testing.T) {
	handles.Init()

	tmpDir := t.TempDir()

	// Test with different format strings
	filePath := filepath.Join(tmpDir, "test2.txt")
	f := C_nox_fs_create_text(filePath)
	require.NotNil(t, f)

	C_nox_fs_fprintf_s(f, "String: %s\n", "test string")
	C_nox_fs_fprintf_d(f, "Positive: %d\n", 12345)
	C_nox_fs_fprintf_d(f, "Negative: %d\n", -6789)
	C_nox_fs_fprintf_d(f, "Zero: %d\n", 0)
	C_nox_fs_close(f)

	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	expected := "String: test string\nPositive: 12345\nNegative: -6789\nZero: 0\n"
	require.Equal(t, expected, string(data))

	// Test empty file
	filePath2 := filepath.Join(tmpDir, "empty.txt")
	f2 := C_nox_fs_create_text(filePath2)
	require.NotNil(t, f2)
	C_nox_fs_close(f2)

	data, err = os.ReadFile(filePath2)
	require.NoError(t, err)
	require.Equal(t, "", string(data))

	// Test with special characters
	filePath3 := filepath.Join(tmpDir, "special.txt")
	f3 := C_nox_fs_create_text(filePath3)
	require.NotNil(t, f3)
	C_nox_fs_fprintf_s(f3, "Special: %s\n", "a/b\\c:d*e?f\"g<h>i|j")
	C_nox_fs_close(f3)

	data, err = os.ReadFile(filePath3)
	require.NoError(t, err)
	require.Contains(t, string(data), "Special:")
}
