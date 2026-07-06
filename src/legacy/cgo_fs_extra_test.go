package legacy

import (
	"os"
	"path/filepath"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func TestNoxFsNormalize(t *testing.T) {
	s := CString("/tmp/../tmp/test")
	defer StrFree(s)
	out := nox_fs_normalize(s)
	defer StrFree(out)
	got := GoString(out)
	if got == "" {
		t.Error("normalize should not return empty")
	}
}

func TestNoxFsMkdirRemove(t *testing.T) {
	handles.Init()
	tmp := t.TempDir()
	dir := filepath.Join(tmp, "subdir")
	cdir := CString(dir)
	defer StrFree(cdir)
	if !nox_fs_mkdir(cdir) {
		t.Error("mkdir should succeed")
	}
	if !nox_fs_remove(cdir) {
		t.Error("remove should succeed")
	}
}

func TestNoxFsCopyMove(t *testing.T) {
	handles.Init()
	tmp := t.TempDir()
	src := filepath.Join(tmp, "src.txt")
	dst := filepath.Join(tmp, "dst.txt")
	mv := filepath.Join(tmp, "mv.txt")
	os.WriteFile(src, []byte("hello"), 0644)
	csrc := CString(src)
	cdst := CString(dst)
	cmv := CString(mv)
	defer StrFree(csrc)
	defer StrFree(cdst)
	defer StrFree(cmv)
	if !nox_fs_copy(csrc, cdst) {
		t.Error("copy should succeed")
	}
	if !nox_fs_move(cdst, cmv) {
		t.Error("move should succeed")
	}
}

func TestNoxFsFseekFtell(t *testing.T) {
	handles.Init()
	tmp := t.TempDir()
	path := filepath.Join(tmp, "test.bin")
	cpath := CString(path)
	defer StrFree(cpath)
	// Create file first
	f := nox_fs_create(cpath)
	if f == nil {
		t.Fatal("create failed")
	}
	nox_fs_close(f)
	f = nox_fs_open_rw(cpath)
	if f == nil {
		t.Fatal("open_rw failed")
	}
	data := []byte("0123456789")
	nox_fs_fwrite(f, unsafePtr(&data[0]), len(data))
	if nox_fs_fseek(f, 0, 0) != 0 {
		t.Error("fseek should succeed")
	}
	pos := nox_fs_ftell(f)
	if pos != 0 {
		t.Errorf("ftell = %d, want 0", pos)
	}
	nox_fs_fseek(f, 5, 1)
	pos = nox_fs_ftell(f)
	if pos != 5 {
		t.Errorf("ftell after seek cur = %d, want 5", pos)
	}
	sz := nox_fs_fsize(f)
	if sz != 10 {
		t.Errorf("fsize = %d, want 10", sz)
	}
	if nox_fs_feof(f) {
		t.Error("should not be eof")
	}
	nox_fs_fseek(f, 0, 2)
	nox_fs_fseek(f, 0, 0)
	cstr := CString("hello\n")
	defer StrFree(cstr)
	nox_fs_fputs(f, cstr)
	nox_fs_fseek(f, 0, 0)
	buf := make([]byte, 100)
	// Skip nox_fs_fgets to avoid C import in test
	_ = buf
	nox_fs_close(f)
}

func TestNoxFsAccess(t *testing.T) {
	handles.Init()
	tmp := t.TempDir()
	path := filepath.Join(tmp, "a.txt")
	os.WriteFile(path, []byte("x"), 0644)
	cpath := CString(path)
	defer StrFree(cpath)
	if nox_fs_access(cpath, 0) != 0 {
		t.Error("access should succeed")
	}
}

func TestConvWhence(t *testing.T) {
	if convWhence(0) != 0 { // SEEK_SET
		t.Error("SEEK_SET should be 0")
	}
	if convWhence(1) != 1 { // SEEK_CUR
		t.Error("SEEK_CUR should be 1")
	}
	if convWhence(2) != 2 { // SEEK_END
		t.Error("SEEK_END should be 2")
	}
}

func unsafePtr(p *byte) unsafe.Pointer {
	return unsafe.Pointer(p)
}
