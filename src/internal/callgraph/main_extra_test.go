package main

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsCgoExportWithDoc(t *testing.T) {
	decl := &ast.FuncDecl{
		Name: ast.NewIdent("foo"),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "//export foo"},
			},
		},
	}
	require.True(t, isCgoExport(decl))

	decl2 := &ast.FuncDecl{
		Name: ast.NewIdent("bar"),
		Doc: &ast.CommentGroup{
			List: []*ast.Comment{
				{Text: "// not export"},
			},
		},
	}
	require.False(t, isCgoExport(decl2))
}

func TestDeclNameWithReceiver(t *testing.T) {
	// Test with receiver
	decl := &ast.FuncDecl{
		Name: ast.NewIdent("Method"),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: &ast.StarExpr{
						X: ast.NewIdent("MyType"),
					},
				},
			},
		},
	}
	name := declName(decl)
	require.Equal(t, "MyType.Method", name)

	// Test with non-pointer receiver
	decl2 := &ast.FuncDecl{
		Name: ast.NewIdent("Method2"),
		Recv: &ast.FieldList{
			List: []*ast.Field{
				{
					Type: ast.NewIdent("MyType2"),
				},
			},
		},
	}
	name2 := declName(decl2)
	require.Equal(t, "MyType2.Method2", name2)
}

func TestWalkInDeclWithFunc(t *testing.T) {
	f := &ast.File{
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("foo"),
				Body: &ast.BlockStmt{
					List: []ast.Stmt{
						&ast.ExprStmt{
							X: ast.NewIdent("x"),
						},
					},
				},
			},
		},
	}
	called := false
	walkInDecl(f, "test.", nil, nil, func(n ast.Node) {
		called = true
	})
	require.True(t, called)
}

func TestWalkInDeclWithCur(t *testing.T) {
	f := &ast.File{
		Decls: []ast.Decl{
			&ast.FuncDecl{
				Name: ast.NewIdent("foo"),
				Body: &ast.BlockStmt{},
			},
		},
	}
	cur := ""
	walkInDecl(f, "test.", &cur, nil, func(n ast.Node) {})
	require.Equal(t, "", cur) // Should be reset after
}

func TestRunError(t *testing.T) {
	// Test run with invalid root (no src directory)
	err := run("/nonexistent/path")
	require.Error(t, err)
	require.Contains(t, err.Error(), "doesn't look like a root")

	// Test run with valid dir but no cxgo
	dir := t.TempDir()
	srcDir := filepath.Join(dir, "src")
	os.MkdirAll(srcDir, 0755)
	err = run(dir)
	require.Error(t, err)
	require.Contains(t, err.Error(), "cxgo is not installed")
}

func TestExecInSuccess(t *testing.T) {
	err := execIn("", "true")
	require.NoError(t, err)
}

func TestExecInFailure(t *testing.T) {
	err := execIn("", "false")
	require.Error(t, err)
}

func TestExecInInvalidDir(t *testing.T) {
	err := execIn("/nonexistent/dir/xyz", "true")
	require.Error(t, err)
}
