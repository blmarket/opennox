package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseBytesExtra2(t *testing.T) {
	b := &Blobs{}
	out, err := b.parseBytes([]byte("0x1, 2, 3u, 0xff"), 4)
	require.NoError(t, err)
	require.Equal(t, []byte{1, 2, 3, 0xff}, out)
}

func TestBlobsGetVariantsExtra2(t *testing.T) {
	b := &Blobs{}
	b.mapg1.blobs1 = append(b.mapg1.blobs1, Blob{Blob: 0x1, Size: 10})
	b.mapg2.blobs = append(b.mapg2.blobs, Blob{Blob: 0x2, Size: 20})
	b.mapc.blobs1 = append(b.mapc.blobs1, Blob{Blob: 0x3, Size: 30})
	b.mapg1.blobs2 = append(b.mapg1.blobs2, Blob{Blob: 0x4})
	b.mapc.blobs2 = append(b.mapc.blobs2, Blob{Blob: 0x5})
	b.data.blobs = append(b.data.blobs, Blob{Blob: 0x6, Size: 60, Data: []byte{1}})
	require.NotNil(t, b.Get(0x1))
	require.NotNil(t, b.Get(0x2))
	require.NotNil(t, b.Get(0x3))
	require.NotNil(t, b.Get(0x4))
	require.NotNil(t, b.Get(0x5))
	require.NotNil(t, b.Get(0x6))
	require.Nil(t, b.Get(0x9999))
}
