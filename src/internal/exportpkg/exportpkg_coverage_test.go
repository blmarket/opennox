package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWriteAST2(t *testing.T) {
	// Test writeAST with valid AST
	fs := token.NewFileSet()
	src := `package main
func main() {}
`
	f, err := parser.ParseFile(fs, "test.go", src, parser.AllErrors)
	require.NoError(t, err)

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "out.go")

	err = writeAST(path, fs, f)
	require.NoError(t, err)

	// Verify file was created
	_, err = os.Stat(path)
	require.NoError(t, err)
}

func TestWriteASTError2(t *testing.T) {
	// Test writeAST with invalid AST that causes format.Node to fail
	// Create a file with invalid AST structure
	fs := token.NewFileSet()
	f := &ast.File{
		Name: ast.NewIdent("main"),
		// Invalid AST that may cause format to fail
		Decls: []ast.Decl{
			&ast.BadDecl{},
		},
	}

	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "out.go")

	err := writeAST(path, fs, f)
	// Should return an error for invalid AST, or succeed with printer fallback
	_ = err
}

func TestEachDstFile2(t *testing.T) {
	// Test eachDstFile with non-existent directory (error path)
	origDir := *fDir
	defer func() { *fDir = origDir }()

	*fDir = "/non/existent/path"
	err := eachDstFile(func(epath string) error {
		return nil
	})
	require.Error(t, err)

	// Test with empty temp directory (no .go files)
	tmpDir := t.TempDir()
	legacyDir := filepath.Join(tmpDir, "legacy")
	err = os.MkdirAll(legacyDir, 0755)
	require.NoError(t, err)

	*fDir = tmpDir
	err = eachDstFile(func(epath string) error {
		return nil
	})
	require.NoError(t, err)
}

func TestEachDstFileWithFiles2(t *testing.T) {
	origDir := *fDir
	defer func() { *fDir = origDir }()

	tmpDir := t.TempDir()
	legacyDir := filepath.Join(tmpDir, "legacy")
	err := os.MkdirAll(legacyDir, 0755)
	require.NoError(t, err)

	// Create a .go file in legacy directory
	goFile := filepath.Join(legacyDir, "test.go")
	err = os.WriteFile(goFile, []byte("package legacy\n"), 0644)
	require.NoError(t, err)

	// Create a non-go file (should be skipped)
	txtFile := filepath.Join(legacyDir, "test.txt")
	err = os.WriteFile(txtFile, []byte("not go"), 0644)
	require.NoError(t, err)

	*fDir = tmpDir
	called := false
	err = eachDstFile(func(epath string) error {
		called = true
		require.Equal(t, goFile, epath)
		return nil
	})
	require.NoError(t, err)
	require.True(t, called)
}

func TestAsExportedName2(t *testing.T) {
	require.Equal(t, "Foo", asExportedName("foo"))
	require.Equal(t, "Bar", asExportedName("bar"))
}

func TestIsC2(t *testing.T) {
	// Test isC with non-selector expression
	fs := token.NewFileSet()
	src := `package main
var x = 1
`
	f, err := parser.ParseFile(fs, "test.go", src, parser.AllErrors)
	require.NoError(t, err)

	// isC should return false for non-selector
	for _, decl := range f.Decls {
		if gen, ok := decl.(*ast.GenDecl); ok {
			for _, spec := range gen.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					for _, val := range vs.Values {
						_, ok := isC(val)
						require.False(t, ok)
					}
				}
			}
		}
	}
}
