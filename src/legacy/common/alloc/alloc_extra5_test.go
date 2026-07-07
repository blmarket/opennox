package alloc

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/stretchr/testify/require"
)

func TestNewCloneFreeSlice(t *testing.T) {
	p, free := New(int(0))
	require.NotNil(t, p)
	*p = 42
	require.Equal(t, 42, *p)
	free()

	src := []int{1, 2, 3, 4}
	clone, free2 := CloneSlice(src)
	require.Equal(t, src, clone)
	free2()

	// Test FreeSlice separately
	clone2, _ := CloneSlice(src)
	FreeSlice(clone2)
}

func TestMemcpyMemcmp(t *testing.T) {
	src, freeSrc := Malloc(8)
	defer freeSrc()
	dst, freeDst := Malloc(8)
	defer freeDst()

	srcSlice := unsafe.Slice((*byte)(src), 8)
	for i := range srcSlice {
		srcSlice[i] = byte(i + 1)
	}
	Memcpy(dst, src, 8)
	require.Equal(t, 0, Memcmp(dst, src, 8))

	dstSlice := unsafe.Slice((*byte)(dst), 8)
	dstSlice[0] = 99
	require.NotEqual(t, 0, Memcmp(dst, src, 8))
}

func TestStrlenFunctions(t *testing.T) {
	s := "hello\x00world"
	ptr := unsafe.Pointer(unsafe.StringData(s))
	require.Equal(t, 5, Strlen(ptr))
	require.Equal(t, 5, StrLen((*byte)(ptr)))
	require.Equal(t, 3, StrLenN[byte]((*byte)(ptr), 3))
	require.Equal(t, 0, StrLenN[byte](nil, 5))
	require.Equal(t, 0, StrLen[byte](nil))

	// strlenN
	require.Equal(t, 5, strlenN((*byte)(ptr), 10))
	require.Equal(t, 0, strlenN[byte](nil, 5))

	// ZeroTermLen
	require.Equal(t, 5, ZeroTermLen((*byte)(ptr)))

	// StrLenS
	b := []byte{'a', 'b', 0, 'c'}
	require.Equal(t, 2, StrLenS(b))
	require.Equal(t, 0, StrLenS([]byte{}))
}

func TestStrcpyStrcatStrcmp(t *testing.T) {
	src, freeSrc := CString("hello")
	defer freeSrc()
	dst, freeDst := Malloc(20)
	defer freeDst()

	Strcpy(dst, unsafe.Pointer(src))
	require.Equal(t, "hello", GoString((*byte)(dst)))

	src2, freeSrc2 := CString(" world")
	defer freeSrc2()
	Strcat(dst, unsafe.Pointer(src2))
	require.Equal(t, "hello world", GoString((*byte)(dst)))

	require.Equal(t, 0, Strcmp(dst, dst))
}

func TestStringFunctions(t *testing.T) {
	// GoStringS
	b := []byte{'x', 'y', 'z', 0}
	require.Equal(t, "xyz", GoStringS(b))

	// StrCopyP
	src, freeSrc := CString("test")
	defer freeSrc()
	dst := make([]byte, 10)
	n := StrCopyP(dst, src)
	require.Equal(t, 4, n)
	require.Equal(t, byte(0), dst[4])

	// StrCopyZeroP
	dst2 := make([]byte, 4)
	n2 := StrCopyZeroP(dst2, src)
	require.Equal(t, 3, n2) // truncated to len-1
	require.Equal(t, byte(0), dst2[3])

	// StrCopyS
	srcSlice := []byte{'a', 'b', 0}
	dst3 := make([]byte, 10)
	n3 := StrCopyS(dst3, srcSlice)
	require.Equal(t, 2, n3)

	// StrCopyZeroS
	dst4 := make([]byte, 2)
	n4 := StrCopyZeroS(dst4, []byte{'1', '2', '3'})
	require.Equal(t, 1, n4)

	// StrCopy
	dst5 := make([]byte, 10)
	n5 := StrCopy(dst5, "hi")
	require.Equal(t, 2, n5)
	require.Equal(t, byte(0), dst5[2])

	// StrCopyZero
	dst6 := make([]byte, 2)
	n6 := StrCopyZero(dst6, "hello")
	require.Equal(t, 1, n6)
	require.Equal(t, byte(0), dst6[1])
}

func TestMemlogFunctions(t *testing.T) {
	EnableMemlog(true)
	EnableMemlog(false)
	// ensureLog is called internally when logging, but we don't need a file
	// Just verify functions don't panic when disabled
	logMemRead(nil, 0)
	logMemWrite(nil, 0)
	logMemReadString(nil, 0)
	logMemWriteString(nil, 0)
}

func TestClassFunctions(t *testing.T) {
	handles.Init()
	defer handles.Release()

	c := NewClass("testclass", 32, 5)
	require.NotNil(t, c)

	// UPtr
	up := c.UPtr()
	require.NotNil(t, up)
	require.Equal(t, up, c.UPtr())

	// AsClass
	require.Equal(t, c, AsClass(up))
	require.Nil(t, AsClass(nil))

	// IsDead
	require.True(t, IsDead(unsafe.Pointer(uintptr(DeadWord))))
	require.False(t, IsDead(nil))

	// NewObject and FreeObjectLast
	obj1 := c.NewObject()
	require.NotNil(t, obj1)
	c.FreeObjectLast(obj1)

	obj2 := c.NewObject()
	require.NotNil(t, obj2)
	c.FreeObjectFirst(obj2)

	// Keep
	c.Keep(0)
	c.Keep(4)

	// ResetStats
	c.ResetStats()

	// FreeAllObjects
	c.FreeAllObjects()

	// Free
	c.Free()

	// Free nil should not panic
	var nilClass *Class
	nilClass.Free()
}

func TestFreeObjectLastNil(t *testing.T) {
	var c *Class
	// Should not panic
	c = NewClass("tmp", 8, 1)
	c.FreeObjectLast(nil)
	c.FreeObjectFirst(nil)
	c.Free()
}

func TestIsDeadFalse(t *testing.T) {
	require.False(t, IsDead(unsafe.Pointer(uintptr(12345))))
}
