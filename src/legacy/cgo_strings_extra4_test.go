package legacy

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestStrNCopyExtra(t *testing.T) {
	dst := make([]byte, 10)
	n := StrNCopyBytes(dst, "hello")
	require.Equal(t, 5, n)
}

func TestWStrCopyExtra(t *testing.T) {
	dst := make([]uint16, 10)
	n := WStrCopySlice(dst, "hello")
	require.Equal(t, 5, n)
}

func TestGoStringNPExtra(t *testing.T) {
	s := "hello\x00world"
	ptr := unsafe.Pointer(unsafe.StringData(s))
	result := GoStringNP(ptr, 11)
	require.Equal(t, "hello", result)
}

func TestGoWStringNExtra(t *testing.T) {
	ws, free := CWString("hello")
	defer free()
	result := GoWStringN(ws, 10)
	require.Equal(t, "hello", result)
}

func TestNoxFsCloseExtra(t *testing.T) {
	Nox_fs_close(nil)
	nox_fs_close(nil)
}

func TestCWStringCopyToExtra(t *testing.T) {
	dst := make([]uint16, 10)
	CWStringCopyTo((*wchar2_t)(unsafe.Pointer(&dst[0])), 10, "")
	require.Equal(t, uint16(0), dst[0])
}
