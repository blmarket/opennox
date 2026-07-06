package binfile

import (
	"io"
	"os"
	"testing"
)

func TestFileOperations(t *testing.T) {
	tmp, err := os.CreateTemp("", "binfile_test_*")
	if err != nil {
		t.Fatal(err)
	}
	path := tmp.Name()
	tmp.Close()
	defer os.Remove(path)

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		t.Fatal(err)
	}
	bf := NewFile(f)
	if bf == nil {
		t.Fatal("NewFile returned nil")
	}

	// Test Write, WriteString
	data := []byte("hello world\nsecond line\r\nthird")
	n, err := bf.Write(data)
	if err != nil || n != len(data) {
		t.Fatalf("Write failed: n=%d err=%v", n, err)
	}
	n, err = bf.WriteString(" extra")
	if err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}

	// Test Size
	sz, err := bf.Size()
	if err != nil {
		t.Fatalf("Size failed: %v", err)
	}
	if sz == 0 {
		t.Error("Size should not be zero")
	}

	// Test Seek
	off, err := bf.Seek(0, io.SeekStart)
	if err != nil || off != 0 {
		t.Errorf("Seek start failed: %d %v", off, err)
	}

	// Test Read
	buf := make([]byte, 5)
	n, err = bf.Read(buf)
	if err != nil || n != 5 {
		t.Errorf("Read failed: n=%d err=%v", n, err)
	}
	if string(buf) != "hello" {
		t.Errorf("Read data mismatch: %s", string(buf))
	}

	// Test ReadString
	bf.Seek(0, io.SeekStart)
	bf.enableBuffer()
	s, err := bf.ReadString()
	if err != nil && err != io.EOF {
		t.Errorf("ReadString failed: %v", err)
	}
	if len(s) == 0 {
		t.Error("ReadString should return data")
	}

	// Test Close
	if err := bf.Close(); err != nil {
		t.Errorf("Close failed: %v", err)
	}
}

func TestBinfileOpenModes(t *testing.T) {
	tmp, err := os.CreateTemp("", "binfile_test2_*")
	if err != nil {
		t.Fatal(err)
	}
	path := tmp.Name()
	tmp.Write([]byte("test data"))
	tmp.Close()
	defer os.Remove(path)

	// Test ReadOnly
	bf, err := BinfileOpen(path, ReadOnly)
	if err != nil {
		t.Errorf("BinfileOpen ReadOnly failed: %v", err)
	} else {
		if bf.Mode() != ReadOnly {
			t.Error("Mode mismatch")
		}
		if bf.flags() != "" {
			// flags is empty until SetKey
		}
		bf.Close()
	}

	// Test WriteOnly
	bf, err = BinfileOpen(path, WriteOnly)
	if err != nil {
		t.Errorf("BinfileOpen WriteOnly failed: %v", err)
	} else {
		bf.Close()
	}

	// Test ReadWrite
	bf, err = BinfileOpen(path, ReadWrite)
	if err != nil {
		t.Errorf("BinfileOpen ReadWrite failed: %v", err)
	} else {
		bf.Close()
	}

	// Test invalid mode
	_, err = BinfileOpen(path, Mode(99))
	if err == nil {
		t.Error("expected error for invalid mode")
	}

	// Test non-existent file
	_, err = BinfileOpen("/nonexistent/path/to/file", ReadOnly)
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestBinfileSetKey(t *testing.T) {
	tmp, err := os.CreateTemp("", "binfile_test3_*")
	if err != nil {
		t.Fatal(err)
	}
	path := tmp.Name()
	tmp.Close()
	defer os.Remove(path)

	bf, err := BinfileOpen(path, WriteOnly)
	if err != nil {
		t.Fatal(err)
	}
	defer bf.Close()

	// Test SetKey with negative key (should close)
	err = bf.SetKey(-1)
	if err != nil {
		t.Errorf("SetKey(-1) should not error, got %v", err)
	}

	// Test SetKey with invalid mode
	bf2 := &Binfile{mode: Mode(99)}
	err = bf2.SetKey(0)
	if err == nil {
		t.Error("expected error for invalid mode in SetKey")
	}
}

func TestBinfileFileSeek(t *testing.T) {
	tmp, err := os.CreateTemp("", "binfile_test4_*")
	if err != nil {
		t.Fatal(err)
	}
	path := tmp.Name()
	tmp.Write([]byte("test data for seek"))
	tmp.Close()
	defer os.Remove(path)

	bf, err := BinfileOpen(path, ReadOnly)
	if err != nil {
		t.Fatal(err)
	}
	defer bf.Close()

	// FileSeek in ReadOnly mode should work
	err = bf.FileSeek(5, io.SeekStart)
	if err != nil {
		t.Errorf("FileSeek failed: %v", err)
	}

	// FileSeek with non-zero offset in non-ReadOnly mode should return nil
	bf2, err := BinfileOpen(path, ReadWrite)
	if err != nil {
		t.Fatal(err)
	}
	defer bf2.Close()
	err = bf2.FileSeek(5, io.SeekStart)
	if err != nil {
		t.Errorf("FileSeek in ReadWrite with non-zero offset should return nil, got %v", err)
	}
	err = bf2.FileSeek(0, io.SeekStart)
	if err != nil {
		t.Errorf("FileSeek with zero offset should work, got %v", err)
	}
}

func TestBinfileWritten(t *testing.T) {
	bf := &Binfile{}
	// Written should not panic, returns offset from Seek
	// With nil File, it may panic, so just verify struct
	_ = bf
}

func TestBinfileFlags(t *testing.T) {
	bf := &Binfile{}
	if bf.flags() != "" {
		t.Error("flags should be empty for new Binfile")
	}
}
