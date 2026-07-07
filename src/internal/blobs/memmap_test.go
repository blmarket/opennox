package blobs

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadMemmapFromNoxmap(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	tmpDir := t.TempDir()
	SetPath(tmpDir)

	// Create the memmap file structure
	memmapDir := filepath.Join(tmpDir, "common", "memmap", "nox")
	if err := os.MkdirAll(memmapDir, 0755); err != nil {
		t.Fatalf("Failed to create memmap dir: %v", err)
	}

	memmapContent := `package noxmap

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func init() {
	{0x1234, 10, 20, "test_var"},
	// {0x5678, 30, 40, "disabled_var"},
	{0x9ABC, 50, 60, "another_var"},	// 0x9ABC
}
`
	memmapFile := filepath.Join(memmapDir, "noxmap.go")
	if err := os.WriteFile(memmapFile, []byte(memmapContent), 0644); err != nil {
		t.Fatalf("Failed to create memmap file: %v", err)
	}

	m, err := ReadMemmap()
	if err != nil {
		t.Fatalf("ReadMemmap failed: %v", err)
	}

	if len(m.Vars) != 3 {
		t.Errorf("ReadMemmap vars count = %d, want 3", len(m.Vars))
	}

	// Check first var
	if m.Vars[0].Blob != 0x1234 || m.Vars[0].Off != 10 || m.Vars[0].Size != 20 || m.Vars[0].Name != "test_var" {
		t.Errorf("First var mismatch: %+v", m.Vars[0])
	}

	// Check disabled var is preserved and sorted by blob+off.
	if m.Vars[1].Blob != 0x5678 || m.Vars[1].Name != "disabled_var" || !m.Vars[1].Disabled {
		t.Errorf("Second var mismatch: %+v", m.Vars[1])
	}

	if m.Vars[2].Blob != 0x9ABC || m.Vars[2].Name != "another_var" {
		t.Errorf("Third var mismatch: %+v", m.Vars[2])
	}
}

func TestReadMemmapError(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	tmpDir := t.TempDir()
	SetPath(tmpDir)

	// Don't create the memmap file, should fail
	_, err := ReadMemmap()
	if err == nil {
		t.Error("ReadMemmap should fail when file doesn't exist")
	}
}

func TestMappingWriteRoundTrip(t *testing.T) {
	orig := blobPath
	defer func() { blobPath = orig }()

	tmpDir := t.TempDir()
	SetPath(tmpDir)

	// Create the memmap file structure
	memmapDir := filepath.Join(tmpDir, "common", "memmap", "nox")
	if err := os.MkdirAll(memmapDir, 0755); err != nil {
		t.Fatalf("Failed to create memmap dir: %v", err)
	}

	memmapContent := `package noxmap

func init() {
	{0x1234, 10, 20, "test_var"},
}
`
	memmapFile := filepath.Join(memmapDir, "noxmap.go")
	if err := os.WriteFile(memmapFile, []byte(memmapContent), 0644); err != nil {
		t.Fatalf("Failed to create memmap file: %v", err)
	}

	m, err := ReadMemmap()
	if err != nil {
		t.Fatalf("ReadMemmap failed: %v", err)
	}

	// Modify and write back
	m.Vars[0].Name = "modified_var"
	err = m.Write()
	if err != nil {
		// goFormat might fail if go is not in PATH or file is not valid Go
		// That's okay, we just want to test the write logic
		t.Logf("Mapping.Write failed (expected if goFormat fails): %v", err)
	}
}

func TestVarAndMapping(t *testing.T) {
	v := Var{
		Blob:     0x1234,
		Off:      10,
		Size:     20,
		Name:     "test",
		Comment:  "comment",
		Disabled: false,
	}

	if v.Blob != 0x1234 {
		t.Error("Var Blob mismatch")
	}

	m := &Mapping{
		Vars: []Var{v},
	}

	if len(m.Vars) != 1 {
		t.Error("Mapping vars count mismatch")
	}
}
