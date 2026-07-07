package blobs

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "opennox-blobs-test-*")
	if err != nil {
		panic(err)
	}
	SetPath(dir)
	code := m.Run()
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
