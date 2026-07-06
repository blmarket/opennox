package alloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestMalloc(t *testing.T) {
	ptr, free := Malloc(100)
	require.NotNil(t, ptr)
	require.NotNil(t, free)
	free()

	// Test zero size panics
	require.Panics(t, func() {
		Malloc(0)
	})
}

func TestNew(t *testing.T) {
	ptr, free := New(int(0))
	require.NotNil(t, ptr)
	require.NotNil(t, free)
	*ptr = 42
	require.Equal(t, 42, *ptr)
	free()
}

func TestMake(t *testing.T) {
	src := []int{1, 2, 3}
	out, free := Make(src, 5)
	require.NotNil(t, out)
	require.NotNil(t, free)
	require.Len(t, out, 5)
	require.Equal(t, 1, out[0])
	require.Equal(t, 2, out[1])
	require.Equal(t, 3, out[2])
	free()

	// Test negative size panics
	require.Panics(t, func() {
		Make(src, -1)
	})
}

func TestCloneSlice(t *testing.T) {
	src := []byte{1, 2, 3, 4, 5}
	out, free := CloneSlice(src)
	require.NotNil(t, out)
	require.NotNil(t, free)
	require.Equal(t, src, out)
	free()
}

func TestCalloc(t *testing.T) {
	ptr, free := Calloc(10, 4)
	require.NotNil(t, ptr)
	require.NotNil(t, free)
	free()
}

func TestFreePtr(t *testing.T) {
	ptr, _ := Malloc(50)
	require.NotNil(t, ptr)
	FreePtr(ptr)
}

func TestFree(t *testing.T) {
	ptr, _ := Malloc(50)
	require.NotNil(t, ptr)
	Free((*byte)(ptr))
}

func TestFreeSlice(t *testing.T) {
	s, free := Make([]int{1, 2, 3}, 3)
	require.NotNil(t, s)
	FreeSlice(s)
	// free() already called by FreeSlice, so don't call again
	_ = free
}

func TestMemset(t *testing.T) {
	ptr, free := Malloc(10)
	require.NotNil(t, ptr)
	defer free()
	Memset(ptr, 0xFF, 10)
}

func TestMemcpy(t *testing.T) {
	src, free1 := Malloc(10)
	defer free1()
	dst, free2 := Malloc(10)
	defer free2()
	Memcpy(dst, src, 10)
	Memcpy(nil, nil, 0)
}

func TestMemcmp(t *testing.T) {
	ptr1, free1 := Malloc(10)
	defer free1()
	ptr2, free2 := Malloc(10)
	defer free2()
	result := Memcmp(ptr1, ptr2, 10)
	_ = result
}

func TestStrlen(t *testing.T) {
	s := []byte("hello\x00")
	ptr := &s[0]
	n := strlen(ptr)
	require.Equal(t, 5, n)
}

func TestStrLen(t *testing.T) {
	s := []byte("hello\x00world")
	ptr := &s[0]
	n := StrLen(ptr)
	require.Equal(t, 5, n)
}

func TestStrcpy(t *testing.T) {
	src := []byte("hello\x00")
	dst := make([]byte, 10)
	Strcpy(unsafe.Pointer(&dst[0]), unsafe.Pointer(&src[0]))
	require.Equal(t, byte('h'), dst[0])
}

func TestStrcmp(t *testing.T) {
	s1 := []byte("hello\x00")
	s2 := []byte("hello\x00")
	result := Strcmp(unsafe.Pointer(&s1[0]), unsafe.Pointer(&s2[0]))
	require.Equal(t, 0, result)
}
