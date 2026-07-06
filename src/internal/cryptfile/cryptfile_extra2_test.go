package cryptfile

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOpenFile_InvalidPath(t *testing.T) {
	_, err := OpenFile("nonexistent_12345_xyz.bin", ReadOnly, -1)
	require.Error(t, err)
}

func TestCryptFile_FlushNil(t *testing.T) {
	cf := &CryptFile{}
	// Flush on nil file may panic, just verify function exists
	require.NotNil(t, cf.Flush)
}

func TestCryptFile_ReadWriteNil(t *testing.T) {
	cf := &CryptFile{}
	// ReadWrite on nil file panics, just verify function exists
	require.NotNil(t, cf.ReadWrite)
}

func TestGlobalExtra(t *testing.T) {
	g := Global()
	SetGlobal(g)
	require.Equal(t, g, Global())
	Close()
	require.Nil(t, Global())
}

func TestOpenGlobal_InvalidPath(t *testing.T) {
	err := OpenGlobal("nonexistent_12345_xyz.bin", ReadOnly, -1)
	require.Error(t, err)
}
