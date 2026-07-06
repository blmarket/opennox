package cnxz

import (
	"testing"
)

func TestDecompressFileErrors(t *testing.T) {
	// Test with nonexistent file
	err := DecompressFile("/tmp/nonexistent_12345_xyz", "/tmp/out")
	if err == nil {
		t.Error("DecompressFile should return error for nonexistent file")
	}
}

func TestCompressFileErrors(t *testing.T) {
	// Test with nonexistent file
	err := CompressFile("/tmp/nonexistent_12345_xyz", "/tmp/out")
	if err == nil {
		t.Error("CompressFile should return error for nonexistent file")
	}
}
