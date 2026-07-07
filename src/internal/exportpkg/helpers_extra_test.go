package main

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func TestAsExportedNameExtra(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"a", "A"},
		{"testName", "TestName"},
		{"alreadyExported", "AlreadyExported"},
	}

	for _, tt := range tests {
		got := asExportedName(tt.input)
		if got != tt.want {
			t.Errorf("asExportedName(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestIsCExtra(t *testing.T) {
	// Test with C.SelectorExpr
	src := `package test
import "C"
var _ = C.foo
var _ = C.bar
var _ = notC.baz
`
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "test.go", src, 0)
	if err != nil {
		t.Fatalf("Failed to parse test source: %v", err)
	}

	var cIdents []string
	var nonCIdents []string

	ast.Inspect(file, func(n ast.Node) bool {
		if sel, ok := n.(*ast.SelectorExpr); ok {
			if ident, ok := isC(sel); ok {
				cIdents = append(cIdents, ident.Name)
			} else {
				if id, ok := sel.X.(*ast.Ident); ok {
					nonCIdents = append(nonCIdents, id.Name+"."+sel.Sel.Name)
				}
			}
		}
		return true
	})

	if len(cIdents) != 2 {
		t.Errorf("Expected 2 C identifiers, got %d: %v", len(cIdents), cIdents)
	}

	foundFoo := false
	foundBar := false
	for _, name := range cIdents {
		if name == "foo" {
			foundFoo = true
		}
		if name == "bar" {
			foundBar = true
		}
	}
	if !foundFoo || !foundBar {
		t.Errorf("Expected to find 'foo' and 'bar' in C identifiers, got %v", cIdents)
	}
}

func TestWriteASTExtra(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "output.go")

	src := `package test
func Hello() {}
`
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, "test.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("Failed to parse source: %v", err)
	}

	err = writeAST(path, fs, file)
	if err != nil {
		t.Fatalf("writeAST failed: %v", err)
	}

	// Verify file was created
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read output file: %v", err)
	}
	if len(data) == 0 {
		t.Error("Output file should not be empty")
	}
}

func TestWriteASTExtraInvalid(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "output.go")

	// Create an invalid AST that will fail format.Node
	fs := token.NewFileSet()
	file := &ast.File{
		Name: ast.NewIdent("test"),
		// Missing required fields to cause format error
	}

	err := writeAST(path, fs, file)
	// Should either succeed with printer fallback or return an error
	// Either outcome is acceptable as long as it doesn't panic
	_ = err
}

func TestRunInvalidModeExtra(t *testing.T) {
	// Save original flag values
	origMode := *fMode
	defer func() { *fMode = origMode }()

	*fMode = "invalid_mode"
	err := run()
	if err == nil {
		t.Error("run() with invalid mode should return an error")
	}
}

func TestRunNoModeExtra(t *testing.T) {
	origMode := *fMode
	defer func() { *fMode = origMode }()

	*fMode = ""
	err := run()
	if err == nil {
		t.Error("run() with empty mode should return an error")
	}
}

func TestEachSrcFileExtra(t *testing.T) {
	tmpDir := t.TempDir()

	// Create some test .go files
	for _, name := range []string{"a.go", "b.go", "c.txt"} {
		path := filepath.Join(tmpDir, name)
		if err := os.WriteFile(path, []byte("package test"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	// Create legacy directory
	legacyDir := filepath.Join(tmpDir, "legacy")
	if err := os.Mkdir(legacyDir, 0755); err != nil {
		t.Fatalf("Failed to create legacy dir: %v", err)
	}

	origDir := *fDir
	defer func() { *fDir = origDir }()
	*fDir = tmpDir

	var processed []string
	err := eachSrcFile(func(path, epath string) error {
		processed = append(processed, filepath.Base(path))
		return nil
	})
	if err != nil {
		t.Fatalf("eachSrcFile failed: %v", err)
	}

	// Should process only .go files (a.go and b.go, not c.txt)
	if len(processed) != 2 {
		t.Errorf("eachSrcFile should process 2 .go files, got %d: %v", len(processed), processed)
	}
}

func TestEachSrcFileExtraNonExistentDir(t *testing.T) {
	origDir := *fDir
	defer func() { *fDir = origDir }()
	*fDir = "/non/existent/directory"

	err := eachSrcFile(func(path, epath string) error {
		return nil
	})
	if err == nil {
		t.Error("eachSrcFile with non-existent dir should return an error")
	}
}

func TestEachDstFileExtra(t *testing.T) {
	tmpDir := t.TempDir()
	legacyDir := filepath.Join(tmpDir, "legacy")
	if err := os.MkdirAll(legacyDir, 0755); err != nil {
		t.Fatalf("Failed to create legacy dir: %v", err)
	}

	// Create some test files in legacy
	for _, name := range []string{"x.go", "y.go", "z.txt"} {
		path := filepath.Join(legacyDir, name)
		if err := os.WriteFile(path, []byte("package legacy"), 0644); err != nil {
			t.Fatalf("Failed to create test file: %v", err)
		}
	}

	origDir := *fDir
	defer func() { *fDir = origDir }()
	*fDir = tmpDir

	var processed []string
	err := eachDstFile(func(epath string) error {
		processed = append(processed, filepath.Base(epath))
		return nil
	})
	if err != nil {
		t.Fatalf("eachDstFile failed: %v", err)
	}

	if len(processed) != 2 {
		t.Errorf("eachDstFile should process 2 .go files, got %d: %v", len(processed), processed)
	}
}

func TestEachDstFileExtraNonExistent(t *testing.T) {
	origDir := *fDir
	defer func() { *fDir = origDir }()
	*fDir = "/non/existent"

	err := eachDstFile(func(epath string) error {
		return nil
	})
	if err == nil {
		t.Error("eachDstFile with non-existent dir should return an error")
	}
}
