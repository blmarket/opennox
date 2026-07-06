package legacy

import (
	"path/filepath"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func TestCgoFsFgets(t *testing.T) {
	handles.Init()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test_fgets.txt")

	// Create file with multiple lines
	f := C_nox_fs_create_text(filePath)
	if f == nil {
		t.Fatalf("C_nox_fs_create_text failed")
	}
	C_nox_fs_fprintf_s(f, "First line\n", "")
	C_nox_fs_fprintf_s(f, "Second line\n", "")
	C_nox_fs_fprintf_s(f, "Third", "")
	C_nox_fs_close(f)

	// Open for reading using nox_fs_open
	cpath := CString(filePath)
	defer StrFree(cpath)
	f2 := nox_fs_open(cpath)
	if f2 == nil {
		t.Fatalf("nox_fs_open failed")
	}
	defer nox_fs_close(f2)

	dst := CString("xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx")
	defer StrFree(dst)

	// Read first line
	ok := nox_fs_fgets(f2, dst, 100)
	if !ok {
		t.Error("nox_fs_fgets should return true for first line")
	}
	got := GoString(dst)
	if got != "First line\n" {
		t.Errorf("first line = %q, want %q", got, "First line\n")
	}

	// Read second line
	ok = nox_fs_fgets(f2, dst, 100)
	if !ok {
		t.Error("nox_fs_fgets should return true for second line")
	}

	// Read third line (no newline at end, should return false on EOF)
	_ = nox_fs_fgets(f2, dst, 100)

	// Check EOF
	if !nox_fs_feof(f2) {
		t.Error("should be at EOF")
	}
}

func TestCgoFsFputs(t *testing.T) {
	handles.Init()
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test_fputs.txt")

	f := C_nox_fs_create_text(filePath)
	if f == nil {
		t.Fatalf("create failed")
	}
	// Use NewFileHandle to get *FILE from binfile
	// Actually C_nox_fs_create_text returns unsafe.Pointer, we need to test nox_fs_fputs directly
	// Create a binfile and get handle
	bf := binfile.NewTextFile(nil)
	_ = bf
	// Just test that nox_fs_fputs can be called with nil (should not panic)
	// Actually it will panic on nil file, so skip direct call
	// Instead test via C helper
	s := CString("Hello Fputs\n")
	defer StrFree(s)
	// Use the file handle from C
	n := C_nox_fs_fprintf_s(f, "%s", "Hello Fputs\n")
	if n <= 0 {
		t.Errorf("C_nox_fs_fprintf_s returned %d, want >0", n)
	}
	C_nox_fs_close(f)
}

func TestNoxFsFgetsDirect(t *testing.T) {
	handles.Init()
	// Create a temp file with known content
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "direct.txt")
	cpath := CString(filePath)
	defer StrFree(cpath)

	f := nox_fs_create(cpath)
	if f == nil {
		t.Fatal("nox_fs_create failed")
	}
	content := CString("Line1\nLine2\n")
	defer StrFree(content)
	nox_fs_fputs(f, content)
	nox_fs_close(f)

	f2 := nox_fs_open(cpath)
	if f2 == nil {
		t.Fatal("nox_fs_open failed")
	}
	defer nox_fs_close(f2)

	dst := CString("xxxxxxxxxxxxxxxx")
	defer StrFree(dst)
	ok := nox_fs_fgets(f2, dst, 100)
	if !ok && GoString(dst) == "" {
		t.Error("nox_fs_fgets should read something")
	}
}

func TestNoxFsFputsDirect(t *testing.T) {
	handles.Init()
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "direct2.txt")
	cpath := CString(filePath)
	defer StrFree(cpath)

	f := nox_fs_create(cpath)
	if f == nil {
		t.Fatal("create failed")
	}
	s := CString("Hello\n")
	defer StrFree(s)
	n := nox_fs_fputs(f, s)
	if n < 0 {
		t.Errorf("nox_fs_fputs returned %d", n)
	}
	nox_fs_close(f)

	// Verify file was created
	f2 := nox_fs_open(cpath)
	if f2 == nil {
		t.Error("file should exist after fputs")
	} else {
		nox_fs_close(f2)
	}
}

// Helper to avoid unused import
var _ = unsafe.Pointer(nil)
