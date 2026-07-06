package blobs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadMemmapCError(t *testing.T) {
	SetPath(t.TempDir())
	b := &Blobs{}
	err := b.readMemmapC()
	require.Error(t, err)
}

func TestWriteMemmapCError(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	// Make path a file to cause error on os.Create
	p := Path(memmapC)
	require.NoError(t, os.MkdirAll(filepath.Dir(p), 0755))
	require.NoError(t, os.WriteFile(p, []byte{}, 0644))
	// Now make it a directory to cause error
	os.Remove(p)
	require.NoError(t, os.Mkdir(p, 0755))
	b := &Blobs{}
	err := b.writeMemmapC()
	require.Error(t, err)
}

func TestReadMemmapGo2Error(t *testing.T) {
	SetPath(t.TempDir())
	b := &Blobs{}
	err := b.readMemmapGo2()
	require.Error(t, err)
}

func TestWriteMemmapGo2(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	p := Path(memmapGo2)
	require.NoError(t, os.MkdirAll(filepath.Dir(p), 0755))
	require.NoError(t, os.WriteFile(p, []byte("package test\n"), 0644))
	b := &Blobs{}
	b.mapg2.pre = []byte("pre")
	b.mapg2.post = []byte("post")
	b.mapg2.blobs = []Blob{{Blob: 0x123, Size: 4}}
	// goFormat will fail because file is not valid go, but we test error path
	err := b.writeMemmapGo2()
	require.Error(t, err)
}

func TestReadDataError(t *testing.T) {
	SetPath(t.TempDir())
	b := &Blobs{}
	err := b.readData()
	require.Error(t, err)
}

func TestWriteDataError(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	p := Path(dataC)
	require.NoError(t, os.MkdirAll(filepath.Dir(p), 0755))
	require.NoError(t, os.Mkdir(p, 0755)) // directory instead of file
	b := &Blobs{}
	err := b.writeData()
	require.Error(t, err)
}

func TestSplitBlobMore(t *testing.T) {
	SetPath(t.TempDir())
	// Test not found
	err := SplitBlob(0xDEAD, 0, 0)
	require.Error(t, err)
}

func TestFormatAccessesMore(t *testing.T) {
	// Test with empty
	err := FormatAccesses()
	require.Error(t, err) // no path set, should error
}

func TestRewriteAccessError(t *testing.T) {
	SetPath("/nonexistent/path/12345")
	err := RewriteAccess(func(a *Access) (bool, error) { return false, nil })
	require.Error(t, err)
}
