package main

import (
	"go/ast"
	"go/token"
	"os"
	"path/filepath"
	"testing"
)

func TestAsExportedName(t *testing.T) {
	if got := asExportedName("fooBar"); got != "FooBar" {
		t.Fatalf("expected FooBar, got %s", got)
	}
}

func TestIsC(t *testing.T) {
	sel := &ast.SelectorExpr{
		X:   ast.NewIdent("C"),
		Sel: ast.NewIdent("int"),
	}
	if id, ok := isC(sel); !ok || id.Name != "int" {
		t.Fatalf("expected C.int")
	}
	sel2 := &ast.SelectorExpr{
		X:   ast.NewIdent("X"),
		Sel: ast.NewIdent("int"),
	}
	if _, ok := isC(sel2); ok {
		t.Fatalf("expected not C")
	}
	if _, ok := isC(ast.NewIdent("foo")); ok {
		t.Fatalf("expected not C")
	}
}

func TestWriteAST(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.go")
	fs := token.NewFileSet()
	f := &ast.File{
		Name: ast.NewIdent("main"),
		Decls: []ast.Decl{
			&ast.GenDecl{
				Tok: token.VAR,
				Specs: []ast.Spec{
					&ast.ValueSpec{
						Names: []*ast.Ident{ast.NewIdent("x")},
						Type:  ast.NewIdent("int"),
					},
				},
			},
		},
	}
	if err := writeAST(path, fs, f); err != nil {
		t.Fatalf("writeAST failed: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read failed: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("empty output")
	}
}

func TestExportProcessFile(t *testing.T) {
	dir := t.TempDir()
	src := filepath.Join(dir, "in.go")
	out := filepath.Join(dir, "out.go")
	content := `package legacy
//export foo
func foo() {}
`
	os.WriteFile(src, []byte(content), 0644)
	err := exportProcessFile(src, out)
	if err != nil {
		t.Fatalf("exportProcessFile failed: %v", err)
	}
	// second run should skip
	err = exportProcessFile(src, out)
	if err != nil {
		t.Fatalf("second run failed: %v", err)
	}
}
