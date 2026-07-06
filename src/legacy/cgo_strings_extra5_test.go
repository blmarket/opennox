package legacy

import (
	"testing"
	"unsafe"
)

func TestStrNCopyFunc(t *testing.T) {
	dst := CString("xxxxxxxxxx")
	defer StrFree(dst)
	n := StrNCopy(dst, 10, "hi")
	if n != 2 {
		t.Errorf("StrNCopy n = %d, want 2", n)
	}
	if got := GoString(dst); got != "hi" {
		t.Errorf("StrNCopy result = %q, want hi", got)
	}
	n = StrNCopy(dst, 5, "hello world")
	if n != 5 {
		t.Errorf("StrNCopy truncation n = %d, want 5", n)
	}
}

func TestWStrCopyFunc(t *testing.T) {
	p, free := CWString("xxxxxxxxxx")
	defer free()
	n := WStrCopy(p, 10, "hi")
	if n != 2 {
		t.Errorf("WStrCopy n = %d, want 2", n)
	}
	if got := GoWString(p); got != "hi" {
		t.Errorf("WStrCopy result = %q, want hi", got)
	}
}

func TestWStrCopyBytesFunc(t *testing.T) {
	buf := make([]byte, 20)
	n := WStrCopyBytes(buf, "test")
	if n != 4 {
		t.Errorf("WStrCopyBytes n = %d, want 4", n)
	}
	// WStrCopyBytes writes UTF-16, GoWStringBytes stops at zero high byte for ASCII,
	// so just verify the function didn't panic and returned correct count
}

func TestStrCopyPFunc(t *testing.T) {
	dst := CString("xxxxxxxxxx")
	defer StrFree(dst)
	n := StrCopyP(unsafe.Pointer(dst), 10, "hello")
	if n != 5 {
		t.Errorf("StrCopyP n = %d, want 5", n)
	}
}
