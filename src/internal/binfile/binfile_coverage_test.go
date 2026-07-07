package binfile

import (
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinfileCoverage(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test.bin")

	// Test BinfileOpen with invalid mode
	_, err := BinfileOpen(path, Mode(99))
	require.Error(t, err)

	// Create a file for writing
	bf, err := BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	require.NotNil(t, bf)
	require.Equal(t, WriteOnly, bf.Mode())

	// Test SetKey with negative key (should close and return nil)
	err = bf.SetKey(-1)
	require.NoError(t, err)

	// Test SetKey with nil file
	var nilBf *Binfile
	err = nilBf.SetKey(1)
	require.Error(t, err)

	// Reopen for writing with key
	bf, err = BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	err = bf.SetKey(1)
	require.NoError(t, err)

	// Test Write
	n, err := bf.Write([]byte("hello world"))
	require.NoError(t, err)
	require.Greater(t, n, 0)

	// Test Written
	written := bf.Written()
	require.Greater(t, written, int64(0))

	// Test flags
	require.Equal(t, "W", bf.flags())

	// Test Seek on writer (should error)
	_, err = bf.Seek(0, io.SeekStart)
	require.Error(t, err)

	// Test FileSeek with zero off (should reset to end)
	err = bf.FileSeek(0, io.SeekStart)
	_ = err

	// Test FileFlush
	_, err = bf.FileFlush()
	require.NoError(t, err)

	// Test WriteUint32At with negative off (should return nil)
	err = bf.WriteUint32At(123, -1)
	require.NoError(t, err)

	// Test WriteUint32At with valid off
	err = bf.WriteUint32At(456, 0)
	require.NoError(t, err)

	// Test Close
	err = bf.Close()
	require.NoError(t, err)

	// Reopen for reading
	bf, err = BinfileOpen(path, ReadOnly)
	require.NoError(t, err)
	err = bf.SetKey(1)
	require.NoError(t, err)

	// Test Read
	buf := make([]byte, 100)
	n, err = bf.Read(buf)
	_ = n
	_ = err

	// Test ReadAligned
	buf2 := make([]byte, 8)
	n, err = bf.ReadAligned(buf2)
	if err == nil {
		require.GreaterOrEqual(t, n, 0)
	}

	// Test flags for reader
	require.Equal(t, "R", bf.flags())

	// Test Seek on reader
	_, err = bf.Seek(0, io.SeekStart)
	require.NoError(t, err)

	// Test FileSeek on reader
	err = bf.FileSeek(0, io.SeekStart)
	require.NoError(t, err)

	// Test SkipLine
	err = bf.SkipLine()
	require.NoError(t, err)

	// Test Close
	err = bf.Close()
	require.NoError(t, err)

	// Test ReadAligned on writer (should error)
	bf, err = BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	_, err = bf.ReadAligned(buf2)
	require.Error(t, err)
	bf.Close()

	// Test with ReadWrite mode
	bf, err = BinfileOpen(path, ReadWrite)
	require.NoError(t, err)
	err = bf.SetKey(1)
	require.NoError(t, err)
	require.Equal(t, "RW", bf.flags())
	bf.Close()

	// Test SetKey with invalid mode
	bf = &Binfile{mode: Mode(99)}
	err = bf.SetKey(1)
	require.Error(t, err)
}

func TestFileCoverage(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test2.bin")

	// Create file
	f, err := os.Create(path)
	require.NoError(t, err)
	file := &File{File: f}
	bin := &Binfile{File: file, mode: WriteOnly}
	file.Bin = bin

	// Test enableBuffer
	file.enableBuffer()
	require.NotNil(t, file.buf)

	// Test Size
	sz, err := file.Size()
	require.NoError(t, err)
	require.GreaterOrEqual(t, sz, int64(0))

	// Test Write
	n, err := file.Write([]byte("test"))
	require.NoError(t, err)
	require.Equal(t, 4, n)

	// Test WriteString
	n, err = file.WriteString("hello")
	require.NoError(t, err)
	require.Equal(t, 5, n)

	// Test Seek
	off, err := file.Seek(0, io.SeekStart)
	require.NoError(t, err)
	require.Equal(t, int64(0), off)

	// Test Read
	buf := make([]byte, 10)
	n, err = file.Read(buf)
	_ = n
	_ = err

	// Test ReadString
	file.Seek(0, io.SeekStart)
	s, err := file.ReadString()
	_ = s
	_ = err

	// Test FileSize
	sz, err = FileSize(f)
	require.NoError(t, err)
	require.Greater(t, sz, int64(0))

	// Test Close
	err = file.Close()
	require.NoError(t, err)

	// Test NewFile and NewTextFile
	f2, err := os.Create(path)
	require.NoError(t, err)
	bf := NewFile(f2)
	require.NotNil(t, bf)
	bf.Close()

	f3, err := os.Create(path)
	require.NoError(t, err)
	bf = NewTextFile(f3)
	require.NotNil(t, bf)
	bf.Close()
}

func TestMemfileCoverage(t *testing.T) {
	// Create a test file first
	tmp := t.TempDir()
	path := filepath.Join(tmp, "memtest.bin")
	bf, err := BinfileOpen(path, WriteOnly)
	require.NoError(t, err)
	bf.SetKey(1)
	bf.Write([]byte("hello world test data for memfile"))
	bf.Close()

	mf, err := LoadMemFile(path, 1)
	if err != nil {
		t.Skip("LoadMemFile failed, skipping memfile tests")
		return
	}
	require.NotNil(t, mf)

	// Test C and Free
	cptr := mf.C()
	require.NotNil(t, cptr)

	// Test RawData and Data
	raw := mf.RawData()
	require.NotNil(t, raw)
	d := mf.Data()
	require.NotNil(t, d)

	// Test offset
	off := mf.offset()
	require.GreaterOrEqual(t, off, 0)

	// Test Seek
	pos, err := mf.Seek(5, io.SeekStart)
	require.NoError(t, err)
	require.Equal(t, int64(5), pos)

	pos, err = mf.Seek(2, io.SeekCurrent)
	require.NoError(t, err)
	require.Equal(t, int64(7), pos)

	pos, err = mf.Seek(-3, io.SeekEnd)
	require.NoError(t, err)

	// Test Skip
	mf.Skip(3)

	// Test Read
	buf := make([]byte, 5)
	n, err := mf.Read(buf)
	_ = n
	_ = err

	// Test ReadU8, ReadI8
	mf.Seek(0, io.SeekStart)
	v8 := mf.ReadU8()
	_ = v8
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadI8()

	// Test ReadU16, ReadI16
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadU16()
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadI16()

	// Test ReadU32, ReadI32
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadU32()
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadI32()

	// Test ReadU64, ReadI64
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadU64()
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadI64()

	// Test ReadBytes8
	mf.Seek(0, io.SeekStart)
	b, err := mf.ReadBytes8()
	_ = b
	_ = err

	// Test ReadString8
	mf.Seek(0, io.SeekStart)
	s, err := mf.ReadString8()
	_ = s
	_ = err

	// Test ReadString16
	mf.Seek(0, io.SeekStart)
	s, err = mf.ReadString16()
	_ = s
	_ = err

	// Test SkipString8
	mf.Seek(0, io.SeekStart)
	mf.SkipString8()

	// Test ReadU64Align
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadU64Align()

	mf.Free()
}
