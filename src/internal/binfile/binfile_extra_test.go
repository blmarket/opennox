package binfile

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileFlush(t *testing.T) {
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test.bin")
	f, err := BinfileOpen(path, WriteOnly)
	if err != nil {
		t.Fatalf("BinfileOpen: %v", err)
	}
	if err := f.SetKey(5); err != nil {
		t.Fatalf("SetKey: %v", err)
	}
	_, err = f.Write([]byte("hello world"))
	if err != nil {
		t.Fatalf("Write: %v", err)
	}
	n, err := f.FileFlush()
	if err != nil {
		t.Fatalf("FileFlush: %v", err)
	}
	if n == 0 {
		t.Error("FileFlush should return non-zero")
	}
	f.Close()
	st, err := os.Stat(path)
	if err != nil {
		t.Fatalf("Stat: %v", err)
	}
	if st.Size() == 0 {
		t.Error("file should not be empty")
	}
}
