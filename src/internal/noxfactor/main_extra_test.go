package main

import (
	"go/ast"
	"go/token"
	"testing"
)

func TestStringExpr(t *testing.T) {
	changed := false
	s := stringExpr(&ast.BasicLit{Kind: token.STRING, Value: `"hello"`}, &changed)
	if s == nil {
		t.Fatal("stringExpr should not return nil")
	}
}

func TestIsFuncCall(t *testing.T) {
	call := &ast.CallExpr{Fun: ast.NewIdent("foo")}
	if _, ok := isFuncCall("foo", call); !ok {
		t.Error("should be func call")
	}
	if _, ok := isFuncCall("bar", call); ok {
		t.Error("should not match bar")
	}
}

func TestIsZeroIsOne(t *testing.T) {
	zero := &ast.BasicLit{Kind: token.INT, Value: "0"}
	if !isZeroInt(zero) || !isZero(zero) {
		t.Error("zero should be true")
	}
	one := &ast.BasicLit{Kind: token.INT, Value: "1"}
	if !isOne(one) {
		t.Error("one should be true")
	}
	two := &ast.BasicLit{Kind: token.INT, Value: "2"}
	if isZero(two) || isOne(two) {
		t.Error("two should be false")
	}
}

func TestIntLit(t *testing.T) {
	lit := intLit(42)
	if lit.Value != "42" {
		t.Errorf("intLit value = %s", lit.Value)
	}
}

func TestStrLit(t *testing.T) {
	lit := strLit("test")
	if lit.Value != `"test"` {
		t.Errorf("strLit value = %s", lit.Value)
	}
}

func TestRemoveUnusedExterns(t *testing.T) {
	// Test with unused extern
	data := []byte(`
extern uint32_t unused_var;
int main() { return 0; }
`)
	result := removeUnusedExterns(data)
	// Should remove the unused extern
	if string(result) == string(data) {
		t.Error("removeUnusedExterns should remove unused extern")
	}

	// Test with used extern (used more than once)
	data2 := []byte(`
extern uint32_t used_var;
int main() { return used_var + used_var; }
`)
	result2 := removeUnusedExterns(data2)
	// Should keep the used extern
	if string(result2) != string(data2) {
		t.Error("removeUnusedExterns should keep used extern")
	}

	// Test with no externs
	data3 := []byte(`int main() { return 0; }`)
	result3 := removeUnusedExterns(data3)
	if string(result3) != string(data3) {
		t.Error("removeUnusedExterns should not modify data with no externs")
	}
}
