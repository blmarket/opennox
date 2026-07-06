package main

import (
	"go/ast"
	"testing"
)

func TestIsCgoExport(t *testing.T) {
	if isCgoExport(nil) {
		t.Error("isCgoExport(nil) should be false")
	}
	// Test with non-export decl
	decl := &ast.FuncDecl{}
	if isCgoExport(decl) {
		t.Error("isCgoExport should be false for non-export")
	}
}

func TestDeclName(t *testing.T) {
	// declName panics on nil, so test with valid decl
	decl := &ast.FuncDecl{Name: ast.NewIdent("test")}
	name := declName(decl)
	if name != "test" {
		t.Errorf("declName should return 'test', got %q", name)
	}
}

func TestWalkInDecl(t *testing.T) {
	// walkInDecl requires valid file, just ensure function exists
	// Should not panic when called with empty file
	f := &ast.File{}
	walkInDecl(f, "", nil, nil, nil)
}

func TestExecIn(t *testing.T) {
	// Test execIn with invalid command (should not panic)
	_ = execIn("", "false")
}

func TestRun_InvalidRoot(t *testing.T) {
	// Test with invalid root (no src directory)
	err := run("/nonexistent/path")
	if err == nil {
		t.Error("run with invalid root should return error")
	}

	// Test with empty root
	err = run("")
	if err == nil {
		t.Error("run with empty root should return error")
	}
}

func TestRun_MissingSrc(t *testing.T) {
	// Test with root that doesn't have src directory
	err := run("/tmp")
	if err == nil {
		t.Error("run with root without src should return error")
	}
}
