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

type PointerArithmeticConverter struct {
	fset         *token.FileSet
	replacements map[ast.Expr]ast.Expr
}

func NewPointerArithmeticConverter() *PointerArithmeticConverter {
	return &PointerArithmeticConverter{
		fset:         token.NewFileSet(),
		replacements: make(map[ast.Expr]ast.Expr),
	}
}

func (pac *PointerArithmeticConverter) convertFile(filename string) error {
	src, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	file, err := parser.ParseFile(pac.fset, filename, src, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse file: %w", err)
	}

	ast.Inspect(file, pac.findReplacements)

	if len(pac.replacements) > 0 {
		file = pac.applyReplacements(file).(*ast.File)
	}

	var buf strings.Builder
	if err := format.Node(&buf, pac.fset, file); err != nil {
		return fmt.Errorf("failed to format AST: %w", err)
	}

	if err := os.WriteFile(filename, []byte(buf.String()), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("Updated file: %s\n", filename)
	return nil
}

func (pac *PointerArithmeticConverter) findReplacements(n ast.Node) bool {
	switch node := n.(type) {
	case *ast.StarExpr:
		if pac.isUnsafePointerArithmetic(node.X) {
			replacement := pac.createStarReplacement(node)
			if replacement != nil {
				pac.replacements[node] = replacement
			}
		}
	case *ast.CallExpr:
		if pac.isUnsafePointerCast(node) {
			replacement := pac.createCallReplacement(node)
			if replacement != nil {
				pac.replacements[node] = replacement
			}
		}
	}
	return true
}

func (pac *PointerArithmeticConverter) applyReplacements(node ast.Node) ast.Node {
	switch n := node.(type) {
	case *ast.File:
		for i, decl := range n.Decls {
			n.Decls[i] = pac.applyReplacements(decl).(ast.Decl)
		}
		return n
	case *ast.GenDecl:
		for i, spec := range n.Specs {
			n.Specs[i] = pac.applyReplacements(spec).(ast.Spec)
		}
		return n
	case *ast.FuncDecl:
		if n.Body != nil {
			n.Body = pac.applyReplacements(n.Body).(*ast.BlockStmt)
		}
		return n
	case *ast.BlockStmt:
		for i, stmt := range n.List {
			n.List[i] = pac.applyReplacements(stmt).(ast.Stmt)
		}
		return n
	case *ast.ExprStmt:
		n.X = pac.applyReplacements(n.X).(ast.Expr)
		return n
	case *ast.AssignStmt:
		for i, expr := range n.Lhs {
			n.Lhs[i] = pac.applyReplacements(expr).(ast.Expr)
		}
		for i, expr := range n.Rhs {
			n.Rhs[i] = pac.applyReplacements(expr).(ast.Expr)
		}
		return n
	case *ast.StarExpr:
		if replacement, exists := pac.replacements[n]; exists {
			return replacement
		}
		n.X = pac.applyReplacements(n.X).(ast.Expr)
		return n
	case *ast.CallExpr:
		if replacement, exists := pac.replacements[n]; exists {
			return replacement
		}
		n.Fun = pac.applyReplacements(n.Fun).(ast.Expr)
		for i, arg := range n.Args {
			n.Args[i] = pac.applyReplacements(arg).(ast.Expr)
		}
		return n
	case *ast.Ident:
		return n
	case *ast.SelectorExpr:
		n.X = pac.applyReplacements(n.X).(ast.Expr)
		return n
	case *ast.UnaryExpr:
		n.X = pac.applyReplacements(n.X).(ast.Expr)
		return n
	case *ast.BinaryExpr:
		n.X = pac.applyReplacements(n.X).(ast.Expr)
		n.Y = pac.applyReplacements(n.Y).(ast.Expr)
		return n
	case *ast.BasicLit:
		return n
	case *ast.ParenExpr:
		n.X = pac.applyReplacements(n.X).(ast.Expr)
		return n
	default:
		return n
	}
}

func (pac *PointerArithmeticConverter) isUnsafePointerArithmetic(expr ast.Expr) bool {
	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}

	if !pac.isPointerCast(callExpr) {
		return false
	}

	if len(callExpr.Args) != 1 {
		return false
	}

	unsafeAddCall, ok := callExpr.Args[0].(*ast.CallExpr)
	if !ok {
		return false
	}

	return pac.isUnsafeAddCall(unsafeAddCall)
}

func (pac *PointerArithmeticConverter) isUnsafePointerCast(expr *ast.CallExpr) bool {
	if !pac.isPointerCast(expr) {
		return false
	}

	if len(expr.Args) != 1 {
		return false
	}

	unsafeAddCall, ok := expr.Args[0].(*ast.CallExpr)
	if !ok {
		return false
	}

	return pac.isUnsafeAddCall(unsafeAddCall)
}

func (pac *PointerArithmeticConverter) isPointerCast(expr *ast.CallExpr) bool {
	parenExpr, ok := expr.Fun.(*ast.ParenExpr)
	if !ok {
		return false
	}

	starExpr, ok := parenExpr.X.(*ast.StarExpr)
	if !ok {
		return false
	}

	ident, ok := starExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	return ident.Name == "uint32" || ident.Name == "int32"
}

func (pac *PointerArithmeticConverter) isUnsafeAddCall(expr *ast.CallExpr) bool {
	selExpr, ok := expr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	packageIdent, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	if packageIdent.Name != "unsafe" || selExpr.Sel.Name != "Add" {
		return false
	}

	if len(expr.Args) != 2 {
		return false
	}

	firstArg := expr.Args[0]
	secondArg := expr.Args[1]

	return pac.isUnsafePointerCall(firstArg) && pac.isConstantOffset(secondArg)
}

func (pac *PointerArithmeticConverter) isUnsafePointerCall(expr ast.Expr) bool {
	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return false
	}

	selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	packageIdent, ok := selExpr.X.(*ast.Ident)
	if !ok {
		return false
	}

	return packageIdent.Name == "unsafe" && selExpr.Sel.Name == "Pointer"
}

func (pac *PointerArithmeticConverter) isConstantOffset(expr ast.Expr) bool {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.INT {
			return true
		}
	case *ast.BinaryExpr:
		if e.Op == token.MUL {
			return pac.isConstantOffset(e.X) && pac.isConstantOffset(e.Y)
		}
	}
	return false
}

func (pac *PointerArithmeticConverter) calculateOffset(expr ast.Expr) (int, bool) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.INT {
			val, err := strconv.Atoi(e.Value)
			if err == nil {
				return val, true
			}
		}
	case *ast.BinaryExpr:
		if e.Op == token.MUL {
			x, xOk := pac.calculateOffset(e.X)
			y, yOk := pac.calculateOffset(e.Y)
			if xOk && yOk {
				return x * y, true
			}
		}
	}
	return 0, false
}

func (pac *PointerArithmeticConverter) createStarReplacement(node *ast.StarExpr) ast.Expr {
	callExpr := node.X.(*ast.CallExpr)
	unsafeAddCall := callExpr.Args[0].(*ast.CallExpr)

	varName := pac.extractVarName(unsafeAddCall.Args[0])
	offset, ok := pac.calculateOffset(unsafeAddCall.Args[1])
	if !ok {
		return nil
	}

	fieldIndex := offset / 4

	return &ast.StarExpr{
		X: &ast.UnaryExpr{
			Op: token.AND,
			X: &ast.SelectorExpr{
				X:   &ast.Ident{Name: varName},
				Sel: &ast.Ident{Name: fmt.Sprintf("field_%d", fieldIndex)},
			},
		},
	}
}

func (pac *PointerArithmeticConverter) createCallReplacement(node *ast.CallExpr) ast.Expr {
	unsafeAddCall := node.Args[0].(*ast.CallExpr)

	varName := pac.extractVarName(unsafeAddCall.Args[0])
	offset, ok := pac.calculateOffset(unsafeAddCall.Args[1])
	if !ok {
		return nil
	}

	fieldIndex := offset / 4

	return &ast.UnaryExpr{
		Op: token.AND,
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: varName},
			Sel: &ast.Ident{Name: fmt.Sprintf("field_%d", fieldIndex)},
		},
	}
}

func (pac *PointerArithmeticConverter) extractVarName(expr ast.Expr) string {
	callExpr, ok := expr.(*ast.CallExpr)
	if !ok {
		return "unknown"
	}

	if len(callExpr.Args) == 0 {
		return "unknown"
	}

	if ident, ok := callExpr.Args[0].(*ast.Ident); ok {
		return ident.Name
	}

	return "unknown"
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file.go>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]
	converter := NewPointerArithmeticConverter()

	if err := converter.convertFile(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
