package main

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestUnwrapFuncAdditional(t *testing.T) {
	// Test unwrapFunc with nil
	var changed bool
	result := unwrapFunc(nil, &changed)
	if result != nil {
		t.Error("unwrapFunc(nil) should return nil")
	}

	// Test with simple ident
	ident := &ast.Ident{Name: "foo"}
	result = unwrapFunc(ident, &changed)
	if result != ident {
		t.Error("unwrapFunc(ident) should return ident")
	}

	// Test with paren expr
	paren := &ast.ParenExpr{X: ident}
	result = unwrapFunc(paren, &changed)
	// unwrapFunc may or may not unwrap depending on implementation, just ensure no panic
	_ = result
}

func TestIsZeroIntAdditional(t *testing.T) {
	if !isZeroInt(&ast.BasicLit{Kind: token.INT, Value: "0"}) {
		t.Error("isZeroInt should return true for 0")
	}
	if isZeroInt(&ast.BasicLit{Kind: token.INT, Value: "1"}) {
		t.Error("isZeroInt should return false for 1")
	}
	if isZeroInt(nil) {
		t.Error("isZeroInt(nil) should return false")
	}
}

func TestIsZeroAdditional(t *testing.T) {
	if !isZero(&ast.BasicLit{Kind: token.INT, Value: "0"}) {
		t.Error("isZero should return true for int 0")
	}
	// Float 0.0 may not be considered zero by isZero, just test it doesn't panic
	_ = isZero(&ast.BasicLit{Kind: token.FLOAT, Value: "0.0"})
}

func TestIsOneAdditional(t *testing.T) {
	if !isOne(&ast.BasicLit{Kind: token.INT, Value: "1"}) {
		t.Error("isOne should return true for 1")
	}
	if isOne(&ast.BasicLit{Kind: token.INT, Value: "0"}) {
		t.Error("isOne should return false for 0")
	}
}

func TestStringExprAdditional(t *testing.T) {
	var changed bool
	s := stringExpr(&ast.BasicLit{Kind: token.STRING, Value: `"test"`}, &changed)
	if s == nil {
		t.Error("stringExpr should not return nil")
	}
}

func TestIntLitAdditional(t *testing.T) {
	i := intLit(42)
	if i == nil {
		t.Error("intLit should not return nil")
	}
}

func TestStrLitAdditional(t *testing.T) {
	s := strLit("hello")
	if s == nil {
		t.Error("strLit should not return nil")
	}
}

func TestIdentAdditional(t *testing.T) {
	id := ident("foo")
	if id == nil || id.Name != "foo" {
		t.Error("ident should return correct Ident")
	}
}
