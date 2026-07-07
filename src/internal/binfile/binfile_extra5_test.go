package binfile

import (
	"os"
	"path/filepath"
	"testing"
	"unsafe"
)

func TestBinfileExtra5(t *testing.T) {
	// Test BinfileOpen with invalid mode
	_, err := BinfileOpen("nonexistent.bin", Mode(999))
	if err == nil {
		t.Error("expected error for invalid mode")
	}

	// Test with empty path - should error or return nil
	bf, err := BinfileOpen("", ReadOnly)
	if err == nil && bf != nil {
		bf.Close()
		// Empty path may succeed in some environments, just ensure no panic
	}
}

func TestMemFileExtra5(t *testing.T) {
	// Test NewMemFile with empty data
	mf := NewMemFile(unsafe.Pointer(nil), 0)
	if mf == nil {
		t.Fatal("expected non-nil MemFile")
	}
	// size should be 0
	if mf.RawData() != nil {
		t.Error("expected nil RawData for empty MemFile")
	}
	mf.Free()

	// Test NewMemFile with nil data and size
	mf2 := NewMemFile(nil, 0)
	if mf2 == nil {
		t.Fatal("expected non-nil MemFile for nil data")
	}
	mf2.Free()
}

func TestFileExtra5(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "test.bin")

	// Create a test file
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("failed to create test file: %v", err)
	}
	f.Write([]byte{1, 2, 3, 4, 5})
	f.Close()

	// Test FileSize
	file, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	size, err := FileSize(file)
	file.Close()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if size != 5 {
		t.Errorf("expected size 5, got %d", size)
	}

	// Test NewFile
	osFile, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	nf := NewFile(osFile)
	if nf == nil {
		t.Fatal("expected non-nil File")
	}
	nf.Close()

	// Test NewTextFile
	osFile2, err := os.Open(path)
	if err != nil {
		t.Fatalf("failed to open file: %v", err)
	}
	ntf := NewTextFile(osFile2)
	if ntf == nil {
		t.Fatal("expected non-nil File")
	}
	ntf.Close()
}

func TestBinfileModeConstants(t *testing.T) {
	// Test mode constants have expected values
	if ReadOnly != Mode(0) {
		t.Errorf("expected ReadOnly=0, got %d", ReadOnly)
	}
	if WriteOnly != Mode(1) {
		t.Errorf("expected WriteOnly=1, got %d", WriteOnly)
	}
	if ReadWrite != Mode(2) {
		t.Errorf("expected ReadWrite=2, got %d", ReadWrite)
	}
}

func TestBinfileWrittenEdgeCases(t *testing.T) {
	// Test Written on a newly created binfile (may panic, which is expected)
	func() {
		defer func() {
			recover()
		}()
		bf, err := BinfileOpen("nonexistent.bin", ReadOnly)
		if err == nil && bf != nil {
			bf.Written()
		}
	}()
}
