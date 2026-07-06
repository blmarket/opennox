package main

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringExpr3(t *testing.T) {
	changed := false
	s := stringExpr(&ast.BasicLit{Kind: token.STRING, Value: `"hello"`}, &changed)
	require.NotNil(t, s)
	require.False(t, changed)

	call := &ast.CallExpr{
		Fun:  &ast.Ident{Name: "CString"},
		Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"test"`}},
	}
	s2 := stringExpr(call, &changed)
	require.True(t, changed)
	require.NotNil(t, s2)
}

func TestUnwrapFunc3(t *testing.T) {
	changed := false
	result := unwrapFunc(nil, &changed)
	require.Nil(t, result)

	zero := &ast.BasicLit{Kind: token.INT, Value: "0"}
	result = unwrapFunc(zero, &changed)
	require.True(t, changed)
	require.NotNil(t, result)
}

func TestIsFuncCall3(t *testing.T) {
	call := &ast.CallExpr{Fun: ast.NewIdent("foo")}
	_, ok := isFuncCall("foo", call)
	require.True(t, ok)
	_, ok = isFuncCall("bar", call)
	require.False(t, ok)

	call2 := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "pkg"},
			Sel: &ast.Ident{Name: "bar"},
		},
	}
	_, ok = isFuncCall("pkg.bar", call2)
	require.True(t, ok)
}

func TestHelperLits3(t *testing.T) {
	require.NotNil(t, intLit(1))
	require.NotNil(t, strLit("x"))
	require.NotNil(t, ident("y"))
	require.NotNil(t, not(ident("z")))
	require.NotNil(t, star(ident("a")))
	require.NotNil(t, paren(ident("b")))
	require.NotNil(t, selExpr("a", "b"))
	require.NotNil(t, recvCall(ident("r"), "m"))
	require.NotNil(t, call("f", ident("x")))
	require.NotNil(t, callExpr(ident("f"), ident("x")))
	require.NotNil(t, sliceExprLeft(ident("s"), ident("e")))
}
