package alloc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringU16Coverage(t *testing.T) {
	// Test CString16
	p, free := CString16("test")
	require.NotNil(t, p)
	require.NotNil(t, free)
	free()

	// Test InternCString16
	p1 := InternCString16("test1")
	require.NotNil(t, p1)
	p2 := InternCString16("test1")
	require.Equal(t, p1, p2)
	p3 := InternCString16("test2")
	require.NotNil(t, p3)

	// Test GoString16
	require.Equal(t, "", GoString16(nil))
	s := GoString16(p1)
	require.NotEmpty(t, s)

	// Test GoString16S
	require.Equal(t, "", GoString16S(nil))
	require.Equal(t, "", GoString16S([]uint16{}))
	arr := []uint16{'t', 'e', 's', 't', 0}
	require.Equal(t, "test", GoString16S(arr))

	// Test GoString16B
	require.Equal(t, "", GoString16B(nil))
	require.Equal(t, "", GoString16B([]byte{}))
	b := []byte{'t', 0, 'e', 0, 's', 0, 't', 0, 0, 0}
	require.Equal(t, "test", GoString16B(b))

	// Test StrCopy16P
	dst := make([]uint16, 10)
	n := StrCopy16P(dst, p1)
	require.Greater(t, n, 0)

	// Test StrCopyZero16P
	dst2 := make([]uint16, 10)
	n = StrCopyZero16P(dst2, p1)
	require.Greater(t, n, 0)

	// Test StrCopy16S
	src := []uint16{'a', 'b', 0}
	dst3 := make([]uint16, 10)
	n = StrCopy16S(dst3, src)
	require.Equal(t, 2, n)

	// Test StrCopyZero16S
	dst4 := make([]uint16, 10)
	n = StrCopyZero16S(dst4, src)
	require.Equal(t, 2, n)
	dst5 := make([]uint16, 2)
	n = StrCopyZero16S(dst5, []uint16{'a', 'b', 'c', 0})
	require.Equal(t, 1, n)

	// Test StrCopy16
	dst6 := make([]uint16, 10)
	n = StrCopy16(dst6, "hello")
	require.Equal(t, 5, n)

	// Test StrCopyZero16
	dst7 := make([]uint16, 10)
	n = StrCopyZero16(dst7, "world")
	require.Equal(t, 5, n)
	dst8 := make([]uint16, 3)
	n = StrCopyZero16(dst8, "hello")
	require.Equal(t, 2, n)

	// Test StrCopy16B
	dst9 := make([]byte, 20)
	n = StrCopy16B(dst9, "test")
	require.Equal(t, 4, n)
	dst10 := make([]byte, 20)
	n = StrCopy16B(dst10, "")
	require.Equal(t, 0, n)

	// Test freeIntern16
	freeIntern16()
}
