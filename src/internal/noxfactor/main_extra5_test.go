package main

import (
	"testing"
)

func TestUnwrapFunc2(t *testing.T) {
	changed := false
	result := unwrapFunc(intLit(0), &changed)
	if !changed {
		t.Error("unwrapFunc should change zero to nil")
	}
	_ = result
}

func TestHelperConstructors2(t *testing.T) {
	if not(ident("x")) == nil {
		t.Error("not should not return nil")
	}
	if star(ident("x")) == nil {
		t.Error("star should not return nil")
	}
	if paren(ident("x")) == nil {
		t.Error("paren should not return nil")
	}
	if selExpr("a", "b") == nil {
		t.Error("selExpr should not return nil")
	}
	if call("f", ident("x")) == nil {
		t.Error("call should not return nil")
	}
	if callExpr(ident("f"), ident("x")) == nil {
		t.Error("callExpr should not return nil")
	}
}
