package main

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCcallInferType(t *testing.T) {
	// Test BasicLit INT
	lit := &ast.BasicLit{Kind: token.INT, Value: "42"}
	a, typ, fn := ccallInferType(lit)
	require.Equal(t, lit, a)
	require.NotNil(t, typ)
	require.NotNil(t, fn)

	// Test BasicLit STRING
	lit = &ast.BasicLit{Kind: token.STRING, Value: `"hello"`}
	a, typ, fn = ccallInferType(lit)
	require.Equal(t, lit, a)
	require.NotNil(t, typ)
	require.NotNil(t, fn)

	// Test Ident nil
	ident := &ast.Ident{Name: "nil"}
	a, typ, fn = ccallInferType(ident)
	require.Equal(t, ident, a)
	require.NotNil(t, typ)

	// Test Ident true/false
	ident = &ast.Ident{Name: "true"}
	a, typ, fn = ccallInferType(ident)
	require.Equal(t, ident, a)
	require.NotNil(t, typ)

	ident = &ast.Ident{Name: "false"}
	a, typ, fn = ccallInferType(ident)
	require.Equal(t, ident, a)

	// Test CallExpr with int
	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "int"},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
	}
	a, typ, fn = ccallInferType(call)
	require.NotNil(t, a)
	require.NotNil(t, typ)

	// Test CallExpr with C.int
	call = &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "C"},
			Sel: &ast.Ident{Name: "int"},
		},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.INT, Value: "1"}},
	}
	a, typ, fn = ccallInferType(call)
	require.NotNil(t, a)
	require.NotNil(t, typ)
	require.NotNil(t, fn)

	// Test unknown
	ident = &ast.Ident{Name: "unknownVar"}
	a, typ, fn = ccallInferType(ident)
	require.Equal(t, ident, a)
	require.NotNil(t, typ) // should be fixme
}

func TestRunCCall_Error(t *testing.T) {
	// runCCall expects specific directory structure, should return error in test
	err := runCCall()
	require.Error(t, err)
}
