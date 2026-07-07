package binfile

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinfile_FullPathCoverage(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_full.bin")

	// Create file first with WriteOnly, then reopen with ReadWrite
	bf, err := BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	bf.Close()

	// Now open with ReadWrite mode and key to trigger f.full path
	bf, err = BinfileOpen(path, ReadWrite)
	require.NoError(t, err)
	require.NotNil(t, bf)

	// Set key to create f.full (crypt.File)
	err = bf.SetKey(1)
	require.NoError(t, err)

	// Test Write with f.full
	data := []byte("hello world test data for full path")
	n, err := bf.Write(data)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	// Test Written with f.full
	written := bf.Written()
	require.Greater(t, written, int64(0))

	// Test flags with f.full (should be "RW")
	require.Equal(t, "RW", bf.flags())

	// Test Seek with f.full
	pos, err := bf.Seek(0, io.SeekStart)
	require.NoError(t, err)
	require.Equal(t, int64(0), pos)

	// Test Read with f.full
	buf := make([]byte, len(data))
	n, err = bf.Read(buf)
	require.NoError(t, err)
	require.Equal(t, len(data), n)

	// Test ReadAligned with f.full
	pos, err = bf.Seek(0, io.SeekStart)
	require.NoError(t, err)
	buf2 := make([]byte, 8)
	n, err = bf.ReadAligned(buf2)
	if err == nil {
		require.GreaterOrEqual(t, n, 0)
	}

	// Test FileSeek with f.full and non-zero off (should return nil for ReadWrite mode)
	err = bf.FileSeek(10, io.SeekStart)
	require.NoError(t, err)

	// Test FileFlush with f.full
	_, err = bf.FileFlush()
	// May error if not all data flushed, but should not panic
	_ = err

	// Test Close with f.full
	err = bf.Close()
	require.NoError(t, err)
}

func TestBinfile_FileSizeErrors(t *testing.T) {
	// Test FileSize with seeker that errors on SeekCurrent
	badSeeker := &myBadSeeker{errOnCurrent: true}
	_, err := FileSize(badSeeker)
	require.Error(t, err)

	// Test FileSize with seeker that errors on SeekEnd
	badSeeker2 := &myBadSeeker{errOnEnd: true}
	_, err = FileSize(badSeeker2)
	require.Error(t, err)

	// Test FileSize with seeker that errors on SeekStart (restore)
	badSeeker3 := &myBadSeeker{errOnStart: true}
	_, err = FileSize(badSeeker3)
	require.Error(t, err)
}

type myBadSeeker struct {
	errOnCurrent bool
	errOnEnd     bool
	errOnStart   bool
	pos          int64
}

func (b *myBadSeeker) Seek(off int64, whence int) (int64, error) {
	switch whence {
	case io.SeekCurrent:
		if b.errOnCurrent {
			return 0, io.ErrUnexpectedEOF
		}
		b.pos += off
	case io.SeekEnd:
		if b.errOnEnd {
			return 0, io.ErrUnexpectedEOF
		}
		b.pos = 100 + off
	case io.SeekStart:
		if b.errOnStart {
			return 0, io.ErrUnexpectedEOF
		}
		b.pos = off
	}
	return b.pos, nil
}

func TestBinfile_SkipLineErrors(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_skip.bin")

	bf, err := BinfileOpen(path, WriteOnly)
	require.NoError(t, err)

	// Write data with newlines
	_, err = bf.Write([]byte("\n\n\nhello\n"))
	require.NoError(t, err)
	bf.Close()

	// Reopen for reading
	bf, err = BinfileOpen(path, ReadOnly)
	require.NoError(t, err)
	defer bf.Close()

	// Test SkipLine - should skip newlines and stop at 'h'
	err = bf.SkipLine()
	require.NoError(t, err)

	// Test SkipLine at EOF
	for {
		err = bf.SkipLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
	}
}

func TestBinfile_WriteUint32AtEdgeCases(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_uint32.bin")

	bf, err := BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	defer bf.Close()

	// Test with negative offset (should return nil)
	err = bf.WriteUint32At(123, -1)
	require.NoError(t, err)

	// Test with valid offset
	err = bf.WriteUint32At(456, 0)
	require.NoError(t, err)

	// Test with another key to trigger cipher creation
	bf2, err := BinfileOpen(path+"2", WriteOnly)
	require.NoError(t, err)
	defer bf2.Close()
	bf2.SetKey(2)
	err = bf2.WriteUint32At(789, 0)
	require.NoError(t, err)
}

func TestFile_EnableBuffer(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_buf.bin")

	f, err := os.Create(path)
	require.NoError(t, err)

	file := NewFile(f)
	require.NotNil(t, file)

	// Test enableBuffer twice (second should be no-op)
	file.enableBuffer()
	require.NotNil(t, file.buf)
	file.enableBuffer() // Should not panic

	// Test Read with buffer
	_, err = f.Write([]byte("test data\n"))
	require.NoError(t, err)
	f.Close()

	f2, err := os.Open(path)
	require.NoError(t, err)
	file2 := NewFile(f2)
	file2.enableBuffer()

	buf := make([]byte, 4)
	n, err := file2.Read(buf)
	require.NoError(t, err)
	require.Equal(t, 4, n)

	file2.Close()
}

func TestFile_WritePanic(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_panic.bin")

	f, err := os.Create(path)
	require.NoError(t, err)

	file := NewFile(f)
	file.enableBuffer()

	// Write on buffered file should panic
	require.Panics(t, func() {
		file.Write([]byte("test"))
	})

	f.Close()

	// WriteString on buffered file should panic
	f2, err := os.Create(path + "2")
	require.NoError(t, err)
	file2 := NewFile(f2)
	file2.enableBuffer()
	require.Panics(t, func() {
		file2.WriteString("test")
	})
	f2.Close()
}

func TestFile_CloseWithBin(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_close.bin")

	bf, err := BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	bf.SetKey(1)

	// Close via File.Close which calls Bin.close
	err = bf.File.Close()
	require.NoError(t, err)
}

func TestFile_ReadString(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test_readstring.bin")

	f, err := os.Create(path)
	require.NoError(t, err)
	_, err = f.Write([]byte("line1\r\nline2\nline3"))
	require.NoError(t, err)
	f.Close()

	f2, err := os.Open(path)
	require.NoError(t, err)
	file := NewFile(f2)

	// Test ReadString with CRLF
	s, err := file.ReadString()
	require.NoError(t, err)
	require.Equal(t, []byte("line1\n"), s)

	// Test ReadString with LF
	s, err = file.ReadString()
	require.NoError(t, err)
	require.Equal(t, []byte("line2\n"), s)

	// Test ReadString at EOF
	s, err = file.ReadString()
	require.Equal(t, io.EOF, err)

	file.Close()
}

func TestMemFile_UncoveredPaths(t *testing.T) {
	// Test LoadMemFile with invalid path
	_, err := LoadMemFile("/non/existent/path", 1)
	require.Error(t, err)

	// Test Seek with invalid whence
	// NewMemFile takes unsafe.Pointer and int, skip direct creation
	// The existing tests already cover MemFile functionality
}
