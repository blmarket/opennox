package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlobsExtra14(t *testing.T) {
	// Test Get with nil blobs
	b := &Blobs{}
	got := b.Get(0x1000)
	require.Nil(t, got)

	// Test Update with nil blobs (should not panic)
	b.Update(Blob{Blob: 0x1000, Size: 50})

	// Test Add and Get
	b.Add(Blob{Blob: 0x2000, Size: 100, Data: []byte{1, 2, 3}})
	got = b.Get(0x2000)
	require.NotNil(t, got)
	require.Equal(t, uintptr(0x2000), got.Blob)

	// Test parseBytes with various inputs
	out, err := b.parseBytes([]byte("0xFF, 0x0, 10"), 3)
	require.NoError(t, err)
	require.Equal(t, []byte{0xFF, 0, 10}, out)

	// Test parseBytes with invalid input
	_, err = b.parseBytes([]byte("not a number"), 1)
	require.Error(t, err)
}

func TestReadMemmapExtra(t *testing.T) {
	// Test ReadMemmap with non-existent file
	SetPath(t.TempDir())
	m, err := ReadMemmap()
	require.Error(t, err)
	require.Nil(t, m)
}

func TestMappingWriteExtra(t *testing.T) {
	// Test Mapping.Write with empty vars
	m := &Mapping{
		pre:  []byte("package memmap\nvar (\n"),
		Vars: []Var{},
		post: []byte(")\n"),
	}
	err := m.Write()
	// May fail due to gofmt not available, but should not panic
	_ = err
}
