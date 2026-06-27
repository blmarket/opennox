package blobs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlobsGetAddUpdate(t *testing.T) {
	b := &Blobs{}
	require.Nil(t, b.Get(0x1234))
	v := Blob{Blob: 0x1234, Size: 100}
	b.Add(v)
	got := b.Get(0x1234)
	require.NotNil(t, got)
	require.Equal(t, uintptr(0x1234), got.Blob)
	require.Equal(t, uintptr(100), got.Size)
	// Update
	v2 := Blob{Blob: 0x1234, Size: 200, Data: []byte{1, 2, 3}}
	b.Update(v2)
	got2 := b.Get(0x1234)
	require.Equal(t, uintptr(200), got2.Size)
}

func TestAccessString(t *testing.T) {
	a := Access{
		Name: StringPos{Val: "getMem"},
		Blob: AddrPos{Val: 0x5D4594},
		Off:  AddrPos{Val: 10},
	}
	require.Equal(t, "getMem(0x5D4594, 10)", a.String())

	a2 := Access{
		Name:  StringPos{Val: "getMem"},
		Blob:  AddrPos{Val: 0x1000},
		Off:   AddrPos{Val: 0},
		Index: StringPos{Val: "i*4"},
	}
	require.Equal(t, "getMem(0x1000, + i*4)", a2.String())

	a3 := Access{
		Name:  StringPos{Val: "getMem"},
		Blob:  AddrPos{Val: 0x2000},
		Off:   AddrPos{Val: 5},
		Index: StringPos{Val: "+ 2"},
	}
	require.Equal(t, "getMem(0x2000, 5 + 2)", a3.String())
}

func TestAccessZero(t *testing.T) {
	var ap *AddrPos
	require.True(t, ap.Zero())
	ap2 := &AddrPos{}
	require.True(t, ap2.Zero())
	ap3 := &AddrPos{Val: 1}
	require.False(t, ap3.Zero())

	var sp *StringPos
	require.True(t, sp.Zero())
	sp2 := &StringPos{}
	require.True(t, sp2.Zero())
	sp3 := &StringPos{Val: "x"}
	require.False(t, sp3.Zero())
}

func TestFormatAccesses(t *testing.T) {
	// just ensure function exists and returns error or nil without panic when path not set
	// It will try to rewrite files; we set path to a temp empty dir to avoid modifying repo
	SetPath(t.TempDir())
	_ = FormatAccesses() // ignore error, just coverage
}
