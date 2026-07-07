package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReformatC(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.c")
	content := []byte("extern uint32_t some_var_123;\nint main() { return 0; }\n")
	require.NoError(t, os.WriteFile(path, content, 0644))

	r := &Refactorer{}
	err := r.reformatC(path)
	require.NoError(t, err)
}

func TestReformatC2Go(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.go")
	content := []byte("package main\n")
	require.NoError(t, os.WriteFile(path, content, 0644))

	r := &Refactorer{}
	err := r.reformatC2Go(path)
	require.NoError(t, err)
}

func TestPreProcessFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "test.go")
	content := []byte("package main\nfunc foo() {}\n")
	require.NoError(t, os.WriteFile(path, content, 0644))

	r := &Refactorer{defined: make(map[string]struct{})}
	err := r.preProcessFile(path)
	require.NoError(t, err)
	require.Contains(t, r.defined, "foo")
}

func TestProcessDir(t *testing.T) {
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	require.NoError(t, os.MkdirAll(srcDir, 0755))

	r := &Refactorer{}
	err := r.ProcessDir(dir)
	_ = err
}
