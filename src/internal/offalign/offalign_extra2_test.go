package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseOffMultExtra(t *testing.T) {
	out, ok := parseOffMult([]byte("10"))
	require.True(t, ok)
	require.Equal(t, uintptr(10), out.Static)
	_, ok2 := parseOffMult([]byte("a+b"))
	require.False(t, ok2)
}

func TestEachFileGoExtra(t *testing.T) {
	tmp := t.TempDir()
	err := os.WriteFile(filepath.Join(tmp, "a.go"), []byte("package p"), 0644)
	require.NoError(t, err)
	var files []string
	err = eachFileGo(tmp, nil, func(path string) error {
		files = append(files, path)
		return nil
	})
	require.NoError(t, err)
	require.Len(t, files, 1)
}

func TestRewriteFileExtra2(t *testing.T) {
	tmp := t.TempDir()
	p := filepath.Join(tmp, "x.c")
	err := os.WriteFile(p, []byte(`a = getMemBytePtr(0x1a3a, 8);`), 0644)
	require.NoError(t, err)
	ok, err := rewriteFile(p, func(data []byte) []byte {
		return offsetAlign(data, 0x1A3A, 6, 2, 4)
	})
	require.NoError(t, err)
	require.True(t, ok)
}

func TestASTHelpersExtra2(t *testing.T) {
	require.NotNil(t, astUint(5))
	require.NotNil(t, astAdd(astUint(1), astUint(2)))
	require.NotNil(t, astMul(astUint(2), astUint(3)))
}
