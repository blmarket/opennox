// Package main provides functionality to enumerate variable declarations within Go functions.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

// VarDeclaration represents a variable declaration with its name and type.
type VarDeclaration struct {
	Name string
	Type string
}

// EnumerateVarDecls parses a Go file and extracts variable declarations from a specific function.
func EnumerateVarDecls(filename, functionName string) ([]VarDeclaration, error) {
	// Parse the Go source file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("failed to parse file: %v", err)
	}

	var varDecls []VarDeclaration

	// Walk through the AST to find the specified function
	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.FuncDecl:
			if node.Name.Name == functionName {
				// Extract function parameters
				if node.Type.Params != nil {
					for _, param := range node.Type.Params.List {
						typeStr := typeToString(param.Type)
						for _, name := range param.Names {
							varDecls = append(varDecls, VarDeclaration{
								Name: name.Name,
								Type: typeStr,
							})
						}
					}
				}

				// Extract variable declarations from function body
				if node.Body != nil {
					ast.Inspect(node.Body, func(n ast.Node) bool {
						switch stmt := n.(type) {
						case *ast.DeclStmt:
							if genDecl, ok := stmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
								for _, spec := range genDecl.Specs {
									if valueSpec, ok := spec.(*ast.ValueSpec); ok {
										typeStr := ""
										if valueSpec.Type != nil {
											typeStr = typeToString(valueSpec.Type)
										} else if len(valueSpec.Values) > 0 {
											// Try to infer type from value
											typeStr = inferTypeFromValue(valueSpec.Values[0])
										}
										for _, name := range valueSpec.Names {
											varDecls = append(varDecls, VarDeclaration{
												Name: name.Name,
												Type: typeStr,
											})
										}
									}
								}
							}
						case *ast.AssignStmt:
							// Handle short variable declarations (var := value)
							if stmt.Tok == token.DEFINE {
								for i, expr := range stmt.Lhs {
									if ident, ok := expr.(*ast.Ident); ok {
										typeStr := ""
										if i < len(stmt.Rhs) {
											typeStr = inferTypeFromValue(stmt.Rhs[i])
										}
										varDecls = append(varDecls, VarDeclaration{
											Name: ident.Name,
											Type: typeStr,
										})
									}
								}
							}
						}
						return true
					})
				}
				return false // Found the function, no need to continue
			}
		}
		return true
	})

	return varDecls, nil
}

// typeToString converts an AST type expression to a string representation.
func typeToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + typeToString(t.X)
	case *ast.ArrayType:
		if t.Len != nil {
			return "[" + exprToString(t.Len) + "]" + typeToString(t.Elt)
		}
		return "[]" + typeToString(t.Elt)
	case *ast.SelectorExpr:
		return typeToString(t.X) + "." + t.Sel.Name
	case *ast.ParenExpr:
		return "(" + typeToString(t.X) + ")"
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.MapType:
		return "map[" + typeToString(t.Key) + "]" + typeToString(t.Value)
	case *ast.ChanType:
		dir := ""
		if t.Dir == ast.SEND {
			dir = "chan<- "
		} else if t.Dir == ast.RECV {
			dir = "<-chan "
		} else {
			dir = "chan "
		}
		return dir + typeToString(t.Value)
	case *ast.FuncType:
		return "func" + funcTypeToString(t)
	default:
		return "unknown"
	}
}

// exprToString converts an expression to string (for array lengths, etc.)
func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return e.Value
	case *ast.Ident:
		return e.Name
	default:
		return "?"
	}
}

// funcTypeToString converts a function type to string representation.
func funcTypeToString(ft *ast.FuncType) string {
	var params []string
	if ft.Params != nil {
		for _, param := range ft.Params.List {
			typeStr := typeToString(param.Type)
			if len(param.Names) > 0 {
				for _, name := range param.Names {
					params = append(params, name.Name+" "+typeStr)
				}
			} else {
				params = append(params, typeStr)
			}
		}
	}

	var results []string
	if ft.Results != nil {
		for _, result := range ft.Results.List {
			typeStr := typeToString(result.Type)
			if len(result.Names) > 0 {
				for _, name := range result.Names {
					results = append(results, name.Name+" "+typeStr)
				}
			} else {
				results = append(results, typeStr)
			}
		}
	}

	sig := "(" + strings.Join(params, ", ") + ")"
	if len(results) > 0 {
		if len(results) == 1 && !strings.Contains(results[0], " ") {
			sig += " " + results[0]
		} else {
			sig += " (" + strings.Join(results, ", ") + ")"
		}
	}
	return sig
}

// inferTypeFromValue attempts to infer the type from a value expression.
func inferTypeFromValue(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		switch e.Kind {
		case token.INT:
			return "int32" // Default to int32 based on context
		case token.FLOAT:
			return "float64"
		case token.STRING:
			return "string"
		case token.CHAR:
			return "rune"
		}
	case *ast.CallExpr:
		// Type conversion or function call
		if ident, ok := e.Fun.(*ast.Ident); ok {
			// Check for type conversions
			if isBuiltinType(ident.Name) {
				return ident.Name
			}
		}
		if sel, ok := e.Fun.(*ast.SelectorExpr); ok {
			return typeToString(sel.X) + "." + sel.Sel.Name
		}
		return typeToString(e.Fun)
	case *ast.UnaryExpr:
		if e.Op == token.AND {
			return "*" + inferTypeFromValue(e.X)
		}
	case *ast.SelectorExpr:
		return typeToString(e.X) + "." + e.Sel.Name
	case *ast.StarExpr:
		return "*" + typeToString(e.X)
	case *ast.Ident:
		return e.Name
	}
	return "unknown"
}

// isBuiltinType checks if a name represents a built-in type.
func isBuiltinType(name string) bool {
	builtins := map[string]bool{
		"int": true, "int8": true, "int16": true, "int32": true, "int64": true,
		"uint": true, "uint8": true, "uint16": true, "uint32": true, "uint64": true, "uintptr": true,
		"float32": true, "float64": true,
		"complex64": true, "complex128": true,
		"byte": true, "rune": true,
		"string": true, "bool": true,
		"error": true,
	}
	return builtins[name]
}