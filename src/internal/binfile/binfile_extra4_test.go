package binfile

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBinfileExtra2(t *testing.T) {
	b := &Binfile{mode: ReadOnly}
	require.Equal(t, ReadOnly, b.Mode())
	require.Equal(t, "", b.flags())
	require.Error(t, (*Binfile)(nil).SetKey(1))
	b2 := &Binfile{mode: ReadOnly}
	require.NoError(t, b2.SetKey(-1))
	f, err := os.CreateTemp("", "binfile_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	defer f.Close()
	sz, err := FileSize(f)
	require.NoError(t, err)
	require.True(t, sz >= 0)
	bf := &Binfile{File: &File{File: f}, mode: ReadWrite}
	bf.File.Bin = bf
	bf.File.enableBuffer()
	require.NotNil(t, bf.File.buf)
	sz2, _ := bf.File.Size()
	require.True(t, sz2 >= 0)
	n, err := bf.File.WriteString("hello")
	require.NoError(t, err)
	require.Equal(t, 5, n)
	_, _ = bf.File.ReadString()
	_, _ = bf.Seek(0, 0)
	require.Equal(t, int64(0), bf.Written())
	_ = bf.SkipLine()
}
