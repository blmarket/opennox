package legacy

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func TestCgoFsFprintf(t *testing.T) {
	handles.Init()

	// Create a temp file
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "test.txt")

	// Create file using C helper
	f := C_nox_fs_create_text(filePath)
	if f == nil {
		t.Fatalf("C_nox_fs_create_text failed for path %s", filePath)
	}

	// Print strings and digits
	C_nox_fs_fprintf_s(f, "Hello, %s!\n", "World")
	C_nox_fs_fprintf_d(f, "Number is %d.\n", 42)

	// Close file
	C_nox_fs_close(f)

	// Read content
	data, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	expected := "Hello, World!\nNumber is 42.\n"
	if string(data) != expected {
		t.Errorf("Expected content:\n%q\nGot:\n%q", expected, string(data))
	}
}
