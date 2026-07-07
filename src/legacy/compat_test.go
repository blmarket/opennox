//go:build !windows

package legacy

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCompatFindFileA(t *testing.T) {
	dir := t.TempDir()
	mustWriteFile(t, filepath.Join(dir, "alpha.txt"), []byte("alpha"))
	mustWriteFile(t, filepath.Join(dir, "beta.bin"), []byte("bb"))
	if err := os.Mkdir(filepath.Join(dir, "subdir"), 0o700); err != nil {
		t.Fatal(err)
	}
	t.Chdir(dir)

	h, first, ok := compatFindFirst("*.*")
	if !ok {
		t.Fatal("FindFirstFileA did not match temp directory contents")
	}
	seen := map[string]compatFindData{first.FileName: first}
	for {
		next, ok := compatFindNext(h)
		if !ok {
			break
		}
		seen[next.FileName] = next
	}
	if !closeCompatFind(h) {
		t.Fatal("FindClose returned false")
	}

	alpha, ok := seen["alpha.txt"]
	if !ok {
		t.Fatalf("alpha.txt not found, saw %#v", seen)
	}
	if alpha.Attributes&compatFileAttributeDirectory != 0 {
		t.Fatalf("alpha.txt reported as directory: %#x", alpha.Attributes)
	}
	if uint32(alpha.Size) != 5 {
		t.Fatalf("alpha.txt low size word = %d, want 5", uint32(alpha.Size))
	}

	subdir, ok := seen["subdir"]
	if !ok {
		t.Fatalf("subdir not found, saw %#v", seen)
	}
	if subdir.Attributes&compatFileAttributeDirectory == 0 {
		t.Fatalf("subdir was not reported as directory: %#x", subdir.Attributes)
	}

	if _, _, ok := compatFindFirst("missing*.dat"); ok {
		t.Fatal("FindFirstFileA unexpectedly matched missing*.dat")
	}
}

func mustWriteFile(t testing.TB, path string, data []byte) {
	t.Helper()
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}
}
