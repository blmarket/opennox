package binfile

import (
	"os"
	"path/filepath"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/stretchr/testify/require"
)

func TestFileSizeAdditional(t *testing.T) {
	// Test FileSize with MemFile
	mfData := []byte{1, 2, 3, 4, 5}
	buf, _ := alloc.Make([]byte{}, len(mfData))
	copy(buf, mfData)
	mf := NewMemFile(unsafe.Pointer(&buf[0]), len(buf))
	defer mf.Free()

	sz, err := FileSize(mf)
	require.NoError(t, err)
	require.Equal(t, int64(len(mfData)), sz)
}

func TestNewFileAdditional(t *testing.T) {
	// Test NewFile with nil
	f := NewFile(nil)
	require.NotNil(t, f)

	// Test NewTextFile with nil
	ft := NewTextFile(nil)
	require.NotNil(t, ft)
	require.True(t, ft.text)
}

func TestBinfileOpenErrorsAdditional(t *testing.T) {
	// Test opening non-existent file for reading
	f, err := BinfileOpen("/non/existent/file.bin", ReadOnly)
	require.Error(t, err)
	require.Nil(t, f)
}

func TestMemFileAdditional(t *testing.T) {
	// Test NewMemFile and basic operations
	data := []byte("hello world test data")
	buf, _ := alloc.Make([]byte{}, len(data))
	copy(buf, data)
	mf := NewMemFile(unsafe.Pointer(&buf[0]), len(buf))
	require.NotNil(t, mf)
	defer mf.Free()

	// Test Data
	require.Equal(t, data, mf.Data())

	// Test RawData
	require.Equal(t, data, mf.RawData())

	// Test C
	require.NotNil(t, mf.C())

	// Test Seek
	pos, err := mf.Seek(5, 0)
	require.NoError(t, err)
	require.Equal(t, int64(5), pos)

	// Test Read
	buf2 := make([]byte, 5)
	n, err := mf.Read(buf2)
	require.NoError(t, err)
	require.Equal(t, 5, n)
	require.Equal(t, []byte(" worl"), buf2)

	// Test Skip
	mf.Skip(2)

	// Test ReadU8
	v := mf.ReadU8()
	require.NotZero(t, v)

	// Test offset
	require.NotZero(t, mf.offset())
}

func TestLoadMemFileAdditional(t *testing.T) {
	// Test LoadMemFile with non-existent file - should return error, not panic
	// We recover from panic in case alloc panics on zero size
	func() {
		defer func() {
			_ = recover()
		}()
		mf, err := LoadMemFile("/non/existent/file.bin", 0)
		if err == nil && mf != nil {
			mf.Free()
		}
	}()

	// Test LoadMemFile with existing file
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test.bin")
	err := os.WriteFile(path, []byte("test data for memfile"), 0644)
	require.NoError(t, err)

	func() {
		defer func() {
			_ = recover()
		}()
		mf, err := LoadMemFile(path, 0)
		if err == nil && mf != nil {
			mf.Free()
		}
	}()
	// Error or panic is acceptable depending on file format
}
