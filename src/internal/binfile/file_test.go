package binfile

import (
	"io"
	"os"
	"path/filepath"
	"testing"
)

func TestFileSizeFunc(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.bin")

	// Create a test file with known size
	content := []byte("1234567890")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()

	size, err := FileSize(f)
	if err != nil {
		t.Fatalf("FileSize failed: %v", err)
	}
	if size != int64(len(content)) {
		t.Errorf("FileSize = %d, want %d", size, len(content))
	}
}

func TestFileSizeError(t *testing.T) {
	// Test with a closed file
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.bin")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	f.Close()

	_, err = FileSize(f)
	if err == nil {
		t.Error("FileSize should fail on closed file")
	}
}

func TestNewFileAndNewTextFile(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.bin")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()

	bf := NewFile(f)
	if bf == nil {
		t.Fatal("NewFile returned nil")
	}
	if bf.text {
		t.Error("NewFile should not set text mode")
	}

	bf2 := NewTextFile(f)
	if bf2 == nil {
		t.Fatal("NewTextFile returned nil")
	}
	if !bf2.text {
		t.Error("NewTextFile should set text mode")
	}
}

func TestFileMethods(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.bin")
	content := []byte("Hello, World!\nSecond line\n")
	if err := os.WriteFile(testFile, content, 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.OpenFile(testFile, os.O_RDWR, 0644)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()
	defer os.Remove(testFile)

	bf := NewFile(f)

	// Test Size
	size, err := bf.Size()
	if err != nil {
		t.Fatalf("Size failed: %v", err)
	}
	if size != int64(len(content)) {
		t.Errorf("Size = %d, want %d", size, len(content))
	}

	// Test Seek
	pos, err := bf.Seek(0, io.SeekStart)
	if err != nil {
		t.Fatalf("Seek failed: %v", err)
	}
	if pos != 0 {
		t.Errorf("Seek position = %d, want 0", pos)
	}

	// Test Read
	buf := make([]byte, 5)
	n, err := bf.Read(buf)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if n != 5 || string(buf) != "Hello" {
		t.Errorf("Read = %q, want %q", string(buf), "Hello")
	}

	// Test Write
	_, err = bf.Seek(0, io.SeekEnd)
	if err != nil {
		t.Fatalf("Seek to end failed: %v", err)
	}
	n, err = bf.Write([]byte("APPENDED"))
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}
	if n != 8 {
		t.Errorf("Write n = %d, want 8", n)
	}

	// Test WriteString
	n, err = bf.WriteString(" MORE")
	if err != nil {
		t.Fatalf("WriteString failed: %v", err)
	}
	if n != 5 {
		t.Errorf("WriteString n = %d, want 5", n)
	}
}

func TestFileReadString(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.txt")
	content := "Line 1\nLine 2\r\nLine 3\n"
	if err := os.WriteFile(testFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()

	bf := NewFile(f)

	// Read first line
	line, err := bf.ReadString()
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	if string(line) != "Line 1\n" {
		t.Errorf("ReadString line 1 = %q, want %q", string(line), "Line 1\n")
	}

	// Read second line (with \r\n, should be converted to \n)
	line, err = bf.ReadString()
	if err != nil {
		t.Fatalf("ReadString failed: %v", err)
	}
	// \r\n should be converted to \n
	if string(line) != "Line 2\n" {
		t.Errorf("ReadString line 2 = %q, want %q", string(line), "Line 2\n")
	}
}

func TestFileClose(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.bin")
	if err := os.WriteFile(testFile, []byte("test"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}

	bf := NewFile(f)
	err = bf.Close()
	if err != nil {
		t.Fatalf("Close failed: %v", err)
	}
}

func TestFileEnableBuffer(t *testing.T) {
	tmpDir := t.TempDir()
	testFile := filepath.Join(tmpDir, "test.bin")
	if err := os.WriteFile(testFile, []byte("test content"), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	f, err := os.Open(testFile)
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer f.Close()

	bf := NewFile(f)
	if bf.buf != nil {
		t.Error("Buffer should be nil initially")
	}

	bf.enableBuffer()
	if bf.buf == nil {
		t.Error("Buffer should not be nil after enableBuffer")
	}

	// Calling again should not create a new buffer
	buf1 := bf.buf
	bf.enableBuffer()
	if bf.buf != buf1 {
		t.Error("enableBuffer should not recreate buffer if already exists")
	}
}
