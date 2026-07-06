package main

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestGetMultExprExtra(t *testing.T) {
	v := &offAligner{elem: 4}

	// Test nil cases
	if y, r := v.getMultExpr(nil); y != nil || r != nil {
		t.Error("getMultExpr(nil) should return nil, nil")
	}
}

func TestAstCommutExtra(t *testing.T) {
	// Test panic on nil
	func() {
		defer func() {
			if recover() == nil {
				t.Error("astCommut with nil should panic")
			}
		}()
		astCommut(token.ADD, nil)
	}()

	// Test single part
	lit := &ast.BasicLit{Kind: token.INT, Value: "1"}
	result := astCommut(token.ADD, lit)
	if result != lit {
		t.Error("astCommut with single part should return that part")
	}
}
