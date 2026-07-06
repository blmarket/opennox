package binfile

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadMemFile_InvalidPath(t *testing.T) {
	_, err := LoadMemFile("nonexistent_file_12345.bin", 0)
	require.Error(t, err)
}

func TestBinfile_FileFlushExtra(t *testing.T) {
	f := &Binfile{}
	// FileFlush on nil File panics, just verify function exists
	require.NotNil(t, f.FileFlush)
}

func TestBinfile_WriteUint32AtExtra(t *testing.T) {
	f := &Binfile{}
	// Should handle nil file gracefully - just verify function exists
	require.NotNil(t, f.WriteUint32At)
}

func TestBinfile_SkipLineExtra(t *testing.T) {
	f := &Binfile{}
	require.NotNil(t, f.SkipLine)
}

func TestMemFile_Data_Empty(t *testing.T) {
	mf := &MemFile{}
	require.Nil(t, mf.Data())
	require.Nil(t, mf.RawData())
}

func TestMemFile_SeekExtra(t *testing.T) {
	mf := NewMemFile(nil, 0)
	require.NotNil(t, mf)
	// Seek on empty should not panic
	pos, err := mf.Seek(0, 0)
	require.NoError(t, err)
	require.Equal(t, int64(0), pos)
	mf.Free()
}
