package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file.go>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	
	// Parse the Go file
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	// Find and replace unsafe pointer arithmetic patterns
	replacements := findUnsafePatterns(fset, file)
	
	if len(replacements) == 0 {
		fmt.Println("No unsafe pointer arithmetic patterns found")
		return
	}

	// Print all found patterns
	fmt.Printf("Found %d unsafe pointer arithmetic patterns:\n", len(replacements))
	for _, repl := range replacements {
		fmt.Printf("  %s: %s -> %s\n", fset.Position(repl.pos), repl.original, repl.replacement)
	}

	// Apply replacements
	applyReplacements(file, replacements)

	// Write back to file
	f, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer f.Close()

	if err := format.Node(f, fset, file); err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated %s with %d replacements\n", filename, len(replacements))
}

type replacementInfo struct {
	pos         token.Pos
	original    string
	replacement string
	node        ast.Node
	newExpr     ast.Expr
}

func findUnsafePatterns(fset *token.FileSet, file *ast.File) []replacementInfo {
	var replacements []replacementInfo

	ast.Inspect(file, func(n ast.Node) bool {
		switch node := n.(type) {
		case *ast.StarExpr:
			// Look for *(*type)(unsafe.Add(...))
			if callExpr, ok := node.X.(*ast.CallExpr); ok {
				if repl := checkTypeCastCall(fset, node, callExpr, true); repl != nil {
					replacements = append(replacements, *repl)
				}
			}
		case *ast.CallExpr:
			// Look for (*type)(unsafe.Add(...))
			if repl := checkTypeCastCall(fset, node, node, false); repl != nil {
				replacements = append(replacements, *repl)
			}
		}
		return true
	})

	return replacements
}

func checkTypeCastCall(fset *token.FileSet, parentNode ast.Node, callExpr *ast.CallExpr, isDereference bool) *replacementInfo {
	// Check if the callExpr.Fun is a ParenExpr (like (*uint32) or (*int32))  
	parenExpr, ok := callExpr.Fun.(*ast.ParenExpr)
	if !ok {
		return nil
	}

	// Check if it's a pointer type cast (*uint32) or (*int32)
	starExpr, ok := parenExpr.X.(*ast.StarExpr)
	if !ok {
		return nil
	}

	ident, ok := starExpr.X.(*ast.Ident)
	if !ok || (ident.Name != "uint32" && ident.Name != "int32") {
		return nil
	}

	// Must have exactly 1 argument: unsafe.Add(...)
	if len(callExpr.Args) != 1 {
		return nil
	}

	unsafeAddCall, ok := callExpr.Args[0].(*ast.CallExpr)
	if !ok {
		return nil
	}

	return checkUnsafeAddPattern(fset, parentNode, unsafeAddCall, isDereference)
}

func checkUnsafeAddPattern(fset *token.FileSet, parentNode ast.Node, callExpr *ast.CallExpr, isDereference bool) *replacementInfo {
	// Check if it's unsafe.Add call
	if !isUnsafeAddCall(callExpr) {
		return nil
	}

	// Must have exactly 2 arguments
	if len(callExpr.Args) != 2 {
		return nil
	}

	// First argument should be unsafe.Pointer(...)
	unsafePtrCall, ok := callExpr.Args[0].(*ast.CallExpr)
	if !ok || !isUnsafePointerCall(unsafePtrCall) {
		return nil
	}

	// Get the variable from unsafe.Pointer(variable)
	if len(unsafePtrCall.Args) != 1 {
		return nil
	}
	
	varName := getVariableName(unsafePtrCall.Args[0])
	if varName == "" {
		return nil
	}

	// Second argument should resolve to a constant integer
	offset, ok := evaluateConstantOffset(callExpr.Args[1])
	if !ok {
		return nil
	}

	// Calculate field index (assuming 4 bytes per field for GOARCH=386)
	if offset%4 != 0 {
		return nil // Not aligned to 4-byte boundary
	}
	fieldIndex := offset / 4

	// Create replacement
	var newExpr ast.Expr
	fieldName := fmt.Sprintf("field_%d", fieldIndex)
	
	if isDereference {
		// *(*uint32)(unsafe.Add(unsafe.Pointer(v1), offset)) -> *&v1.field_N
		newExpr = &ast.StarExpr{
			X: &ast.UnaryExpr{
				Op: token.AND,
				X: &ast.SelectorExpr{
					X:   &ast.Ident{Name: varName},
					Sel: &ast.Ident{Name: fieldName},
				},
			},
		}
	} else {
		// (*uint32)(unsafe.Add(unsafe.Pointer(v1), offset)) -> &v1.field_N
		newExpr = &ast.UnaryExpr{
			Op: token.AND,
			X: &ast.SelectorExpr{
				X:   &ast.Ident{Name: varName},
				Sel: &ast.Ident{Name: fieldName},
			},
		}
	}

	original := nodeToString(fset, parentNode)
	replacement := nodeToString(fset, newExpr)

	return &replacementInfo{
		pos:         parentNode.Pos(),
		original:    original,
		replacement: replacement,
		node:        parentNode,
		newExpr:     newExpr,
	}
}

func isUnsafeAddCall(callExpr *ast.CallExpr) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	
	return ident.Name == "unsafe" && sel.Sel.Name == "Add"
}

func isUnsafePointerCall(callExpr *ast.CallExpr) bool {
	sel, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	
	ident, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	
	return ident.Name == "unsafe" && sel.Sel.Name == "Pointer"
}

func getVariableName(expr ast.Expr) string {
	if ident, ok := expr.(*ast.Ident); ok {
		return ident.Name
	}
	return ""
}

func evaluateConstantOffset(expr ast.Expr) (int, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.INT {
			val, err := strconv.Atoi(e.Value)
			if err != nil {
				return 0, false
			}
			return val, true
		}
	case *ast.BinaryExpr:
		if e.Op == token.MUL {
			left, leftOk := evaluateConstantOffset(e.X)
			right, rightOk := evaluateConstantOffset(e.Y)
			if leftOk && rightOk {
				return left * right, true
			}
		}
	case *ast.ParenExpr:
		return evaluateConstantOffset(e.X)
	}
	return 0, false
}

func nodeToString(fset *token.FileSet, node ast.Node) string {
	var buf strings.Builder
	format.Node(&buf, fset, node)
	return buf.String()
}

func applyReplacements(file *ast.File, replacements []replacementInfo) {
	for _, repl := range replacements {
		replaceNodeInAST(file, repl.node, repl.newExpr)
	}
}

func replaceNodeInAST(file *ast.File, oldNode ast.Node, newExpr ast.Expr) {
	ast.Inspect(file, func(n ast.Node) bool {
		switch parent := n.(type) {
		case *ast.AssignStmt:
			for i, expr := range parent.Lhs {
				if expr == oldNode {
					parent.Lhs[i] = newExpr
					return false
				}
			}
			for i, expr := range parent.Rhs {
				if expr == oldNode {
					parent.Rhs[i] = newExpr
					return false
				}
			}
		case *ast.CallExpr:
			for i, arg := range parent.Args {
				if arg == oldNode {
					parent.Args[i] = newExpr
					return false
				}
			}
		case *ast.ReturnStmt:
			for i, result := range parent.Results {
				if result == oldNode {
					parent.Results[i] = newExpr
					return false
				}
			}
		case *ast.ExprStmt:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
		case *ast.BinaryExpr:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
			if parent.Y == oldNode {
				parent.Y = newExpr
				return false
			}
		case *ast.UnaryExpr:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
		case *ast.ParenExpr:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
		case *ast.StarExpr:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
		case *ast.IndexExpr:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
			if parent.Index == oldNode {
				parent.Index = newExpr
				return false
			}
		case *ast.SelectorExpr:
			if parent.X == oldNode {
				parent.X = newExpr
				return false
			}
		}
		return true
	})
}