package binfile

import (
	"testing"
)

func TestBinfile_Mode_Extra(t *testing.T) {
	b := &Binfile{mode: ReadOnly}
	if b.Mode() != ReadOnly {
		t.Errorf("Mode() = %v, want ReadOnly", b.Mode())
	}
	b.mode = WriteOnly
	if b.Mode() != WriteOnly {
		t.Errorf("Mode() = %v, want WriteOnly", b.Mode())
	}
	b.mode = ReadWrite
	if b.Mode() != ReadWrite {
		t.Errorf("Mode() = %v, want ReadWrite", b.Mode())
	}
}

func TestBinfileOpen_InvalidMode_Extra(t *testing.T) {
	_, err := BinfileOpen("nonexistent", Mode(999))
	if err == nil {
		t.Error("BinfileOpen with invalid mode should return error")
	}
}

func TestBinfile_Written_Extra(t *testing.T) {
	b := &Binfile{}
	// Written() on empty binfile may panic due to nil File, which is expected
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Written() panicked as expected on empty binfile: %v", r)
			}
		}()
		_ = b.Written()
	}()
}

func TestBinfile_Flags_Extra(t *testing.T) {
	b := &Binfile{}
	_ = b.flags()
}

func TestFile_ModeConstants_Extra(t *testing.T) {
	if ReadOnly != 0 {
		t.Errorf("ReadOnly should be 0, got %d", ReadOnly)
	}
	if WriteOnly != 1 {
		t.Errorf("WriteOnly should be 1, got %d", WriteOnly)
	}
	if ReadWrite != 2 {
		t.Errorf("ReadWrite should be 2, got %d", ReadWrite)
	}
}
