package legacy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/stretchr/testify/require"
)

func TestCgoFsFprintfAdditional3(t *testing.T) {
	handles.Init()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test3.txt")

	f := C_nox_fs_create_text(filePath)
	require.NotNil(t, f)

	// Test various format specifiers
	C_nox_fs_fprintf_s(f, "String: %s\n", "test string")
	C_nox_fs_fprintf_d(f, "Decimal: %d\n", 12345)
	C_nox_fs_fprintf_d(f, "Negative: %d\n", -6789)
	C_nox_fs_fprintf_d(f, "Zero: %d\n", 0)
	C_nox_fs_fprintf_d(f, "Hex: %x\n", 0xABCD)
	C_nox_fs_fprintf_d(f, "HEX: %X\n", 0xABCD)
	C_nox_fs_fprintf_d(f, "Octal: %o\n", 511)
	C_nox_fs_fprintf_d(f, "Char: %c\n", 'X')
	C_nox_fs_fprintf_s(f, "Multiple: %s %s %s\n", "a")

	C_nox_fs_close(f)

	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	content := string(data)

	require.Contains(t, content, "String: test string")
	require.Contains(t, content, "Decimal: 12345")
	require.Contains(t, content, "Negative: -6789")
	require.Contains(t, content, "Zero: 0")
	require.Contains(t, content, "Hex: abcd")
	require.Contains(t, content, "HEX: ABCD")
	require.Contains(t, content, "Octal: 777")
	require.Contains(t, content, "Char: X")
}

func TestCgoFsFprintfEmpty(t *testing.T) {
	handles.Init()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "empty.txt")

	f := C_nox_fs_create_text(filePath)
	require.NotNil(t, f)

	// Write empty string
	C_nox_fs_fprintf_s(f, "%s", "")
	C_nox_fs_close(f)

	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	require.Equal(t, "", string(data))
}

func TestCgoFsFprintfLongString(t *testing.T) {
	handles.Init()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "long.txt")

	f := C_nox_fs_create_text(filePath)
	require.NotNil(t, f)

	longStr := ""
	for i := 0; i < 100; i++ {
		longStr += "abcdefghijklmnopqrstuvwxyz"
	}

	C_nox_fs_fprintf_s(f, "Long: %s\n", longStr)
	C_nox_fs_close(f)

	data, err := os.ReadFile(filePath)
	require.NoError(t, err)
	require.Contains(t, string(data), "Long: ")
	require.Contains(t, string(data), "abcdefghijklmnopqrstuvwxyz")
}
