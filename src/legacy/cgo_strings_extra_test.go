package legacy

import (
	"testing"
	"unsafe"
)

func TestStrLenBytes(t *testing.T) {
	if got := StrLenBytes([]byte("hello\x00world")); got != 5 {
		t.Errorf("StrLenBytes = %d, want 5", got)
	}
	if got := StrLenBytes([]byte("hello")); got != 5 {
		t.Errorf("StrLenBytes no null = %d, want 5", got)
	}
}

func TestStrCopyBytes(t *testing.T) {
	dst := make([]byte, 10)
	n := StrCopyBytes(dst, "hello")
	if n != 5 {
		t.Errorf("n = %d, want 5", n)
	}
	if dst[5] != 0 {
		t.Error("dst should be null terminated")
	}
}

func TestStrNCopyBytes(t *testing.T) {
	dst := make([]byte, 5)
	n := StrNCopyBytes(dst, "hello world")
	if n != 5 {
		t.Errorf("n = %d, want 5", n)
	}
}

func TestWStrLenBytes(t *testing.T) {
	if got := WStrLenBytes(nil); got != 0 {
		t.Errorf("nil = %d, want 0", got)
	}
	// "AB" in UTF-16LE: 0x41 0x00 0x42 0x00 0x00 0x00
	// WStrLenBytes stops when either byte is 0, so it returns 0 for ASCII
	b := []byte{0x41, 0x00, 0x42, 0x00, 0x00, 0x00}
	if got := WStrLenBytes(b); got != 0 {
		t.Errorf("WStrLenBytes ASCII = %d, want 0 (stops at 0 high byte)", got)
	}
	// Test with both bytes non-zero
	b2 := []byte{0x01, 0x01, 0x02, 0x02, 0x00, 0x00}
	if got := WStrLenBytes(b2); got != 2 {
		t.Errorf("WStrLenBytes = %d, want 2", got)
	}
}

func TestWStrCopySlice(t *testing.T) {
	dst := make([]uint16, 10)
	n := WStrCopySlice(dst, "Hi")
	if n != 2 {
		t.Errorf("n = %d, want 2", n)
	}
	if dst[2] != 0 {
		t.Error("should null terminate")
	}
}

func TestGoString(t *testing.T) {
	s := CString("hello")
	defer StrFree(s)
	if got := GoString(s); got != "hello" {
		t.Errorf("GoString = %q, want hello", got)
	}
	if got := GoStringP(unsafe.Pointer(s)); got != "hello" {
		t.Errorf("GoStringP = %q", got)
	}
}

func TestGoStringN(t *testing.T) {
	s := CString("hello world")
	defer StrFree(s)
	if got := GoStringN(s, 5); got != "hello" {
		t.Errorf("GoStringN = %q, want hello", got)
	}
	if got := GoStringNP(unsafe.Pointer(s), 5); got != "hello" {
		t.Errorf("GoStringNP = %q", got)
	}
}

func TestGoStringS(t *testing.T) {
	b := []byte("test\x00extra")
	if got := GoStringS(b); got != "test" {
		t.Errorf("GoStringS = %q, want test", got)
	}
}

func TestCStringArray(t *testing.T) {
	arr := CStringArray([]string{"a", "b"})
	if len(arr) != 2 {
		t.Errorf("len = %d, want 2", len(arr))
	}
	for _, s := range arr {
		StrFree(s)
	}
}

func TestCBytes(t *testing.T) {
	p := CBytes([]byte{1, 2, 3})
	if p == nil {
		t.Error("CBytes should not return nil")
	}
	StrFree((*byte)(p))
}

func TestGoWString(t *testing.T) {
	p, free := CWString("hello")
	defer free()
	if got := GoWString(p); got != "hello" {
		t.Errorf("GoWString = %q, want hello", got)
	}
	if got := GoWStringP(unsafe.Pointer(p)); got != "hello" {
		t.Errorf("GoWStringP = %q", got)
	}
	if got := GoWStringN(p, 2); got != "he" {
		t.Errorf("GoWStringN = %q, want he", got)
	}
}

func TestGoWStringBytes(t *testing.T) {
	_ = GoWStringBytes([]byte{0x48, 0x00, 0x69, 0x00, 0x00, 0x00})
	if got := GoWStringBytes(nil); got != "" {
		t.Error("nil should return empty")
	}
	// Test with non-zero high bytes, both bytes non-zero
	// U+0101 = 0x01 0x01 in little endian, so both non-zero
	b := []byte{0x01, 0x01, 0x01, 0x01, 0x00, 0x00}
	got := GoWStringBytes(b)
	if got == "" {
		t.Error("should not be empty for non-zero bytes")
	}
}

func TestCWString(t *testing.T) {
	p, free := CWString("test")
	defer free()
	if p == nil {
		t.Error("CWString should not return nil")
	}
	if got := CWLen("test"); got != 4 {
		t.Errorf("CWLen = %d, want 4", got)
	}
}

func TestCWStringCopyTo(t *testing.T) {
	p, free := CWString("hello")
	defer free()
	CWStringCopyTo(p, 10, "hi")
	if got := GoWString(p); got != "hi" {
		t.Errorf("CWStringCopyTo = %q, want hi", got)
	}
	CWStringCopyTo(p, 10, "")
	if got := GoWString(p); got != "" {
		t.Error("empty src should result in empty")
	}
}

func TestGoWStrSlice(t *testing.T) {
	arr, free := CWStrSlice([]string{"a", "b", "c"})
	defer free()
	out := GoWStrSlice((**wchar2_t)(unsafe.Pointer(&arr[0])))
	if len(out) != 3 {
		t.Errorf("len = %d, want 3", len(out))
	}
	out2 := GoWStrSliceN((**wchar2_t)(unsafe.Pointer(&arr[0])), 2)
	if len(out2) != 2 {
		t.Errorf("len2 = %d, want 2", len(out2))
	}
	if out := GoWStrSliceN(nil, 0); out != nil {
		t.Error("nil should return nil")
	}
}
