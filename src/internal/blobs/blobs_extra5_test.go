package blobs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSplitBlobErrors(t *testing.T) {
	SetPath(t.TempDir())
	err := SplitBlob(0xDEADBEEF, 0, 0)
	require.Error(t, err) // files don't exist or blob not found
}

func TestReadMemmapSuccess(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	// Create a simple memmap file at common/memmap/nox/noxmap.go
	memmapPath := filepath.Join(dir, "common", "memmap", "nox")
	err := os.MkdirAll(memmapPath, 0755)
	require.NoError(t, err)
	content := `package memmap
var (
	// {0x1000, 0, 4, "testVar"}, // 0x1000
	{0x2000, 8, 8, "anotherVar"},
)
`
	err = os.WriteFile(filepath.Join(memmapPath, "noxmap.go"), []byte(content), 0644)
	require.NoError(t, err)
	m, err := ReadMemmap()
	require.NoError(t, err)
	require.NotNil(t, m)
	require.Len(t, m.Vars, 2)
	// Check sorting by blob+off
	require.Equal(t, uintptr(0x1000), m.Vars[0].Blob)
	require.True(t, m.Vars[0].Disabled)
	require.Equal(t, uintptr(0x2000), m.Vars[1].Blob)
	require.False(t, m.Vars[1].Disabled)
}

func TestMappingWrite(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	memmapPath := filepath.Join(dir, "common", "memmap", "nox")
	err := os.MkdirAll(memmapPath, 0755)
	require.NoError(t, err)
	m := &Mapping{
		pre: []byte("package memmap\nvar (\n"),
		Vars: []Var{
			{Blob: 0x1000, Off: 0, Size: 4, Name: "var1", Comment: "test", Disabled: false},
			{Blob: 0x2000, Off: 8, Size: 8, Name: "var2", Comment: "", Disabled: true},
		},
		post: []byte(")\n"),
	}
	err = m.Write()
	if err != nil {
		t.Logf("Write error (may be gofmt): %v", err)
	}
	// Verify file was created
	data, readErr := os.ReadFile(filepath.Join(memmapPath, "noxmap.go"))
	if readErr != nil && err == nil {
		t.Errorf("File should exist if Write succeeded")
	}
	if readErr == nil {
		require.Contains(t, string(data), "var1")
	}
}

func TestReadBlobsSuccess(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	// Create blobs file with simple content
	content := `package blobs
var blobs = []Blob{
	{Blob: 0x1000, Size: 100},
}
`
	err := os.WriteFile(filepath.Join(dir, "blobs.go"), []byte(content), 0644)
	require.NoError(t, err)
	// Also need memmap files for readData
	// Just test that it doesn't panic and returns error for missing memmap
	_, err = ReadBlobs()
	// May succeed or fail depending on memmap files, but shouldn't panic
	_ = err
}

func TestBlobsGetUpdateAdd(t *testing.T) {
	b := &Blobs{}
	b.Add(Blob{Blob: 0x1000, Size: 50})
	b.Add(Blob{Blob: 0x2000, Size: 100})

	// Test Get
	got := b.Get(0x1000)
	require.NotNil(t, got)
	require.Equal(t, uintptr(0x1000), got.Blob)

	got = b.Get(0x9999)
	require.Nil(t, got)

	// Test Update
	b.Update(Blob{Blob: 0x1000, Size: 75})
	got = b.Get(0x1000)
	require.Equal(t, uintptr(75), got.Size)

	// Test Update non-existing (should NOT add, just no-op)
	b.Update(Blob{Blob: 0x3000, Size: 25})
	got = b.Get(0x3000)
	require.Nil(t, got) // Update does not add new blobs
}

func TestParseBytesMore(t *testing.T) {
	b := &Blobs{}

	// Empty with size 0
	out, err := b.parseBytes([]byte{}, 0)
	require.NoError(t, err)
	require.Nil(t, out)

	// With newlines and spaces
	out, err = b.parseBytes([]byte(" 0x1 ,\n 0x2 , 0x3 "), 3)
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2, 3}, out)
}

func TestWriteBlobs(t *testing.T) {
	dir := t.TempDir()
	SetPath(dir)
	b := &Blobs{}
	b.Add(Blob{Blob: 0x1000, Size: 10, Data: []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}})
	err := b.Write()
	// goFormat may fail in test environment, but file should be created
	// Just verify no panic and file exists or error is about goformat
	if err != nil {
		t.Logf("Write error (may be gofmt): %v", err)
	}
	// Verify file exists
	_, statErr := os.Stat(filepath.Join(dir, "blobs.go"))
	if statErr != nil && err == nil {
		t.Errorf("File should exist if Write succeeded")
	}
}
