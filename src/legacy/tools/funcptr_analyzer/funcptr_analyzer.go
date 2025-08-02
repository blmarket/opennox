package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strings"
)

type FuncPtrCast struct {
	FuncType      string
	CastExpr      string
	CallArgs      []string
	Position      token.Position
	SourceContext string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <go-file-path>\n", os.Args[0])
		os.Exit(1)
	}

	filePath := os.Args[1]
	
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	analyzer := &FuncPtrAnalyzer{
		fset:  fset,
		casts: []FuncPtrCast{},
	}

	ast.Walk(analyzer, node)

	for _, cast := range analyzer.casts {
		fmt.Printf("Function Pointer Cast Found:\n")
		fmt.Printf("  Position: %s\n", cast.Position)
		fmt.Printf("  Function Type: %s\n", cast.FuncType)
		fmt.Printf("  Cast Expression: %s\n", cast.CastExpr)
		fmt.Printf("  Call Arguments: %s\n", strings.Join(cast.CallArgs, ", "))
		fmt.Printf("  Source Context: %s\n", cast.SourceContext)
		fmt.Printf("\n")
	}

	if len(analyzer.casts) == 0 {
		fmt.Println("No function pointer casts found.")
	}
}

type FuncPtrAnalyzer struct {
	fset  *token.FileSet
	casts []FuncPtrCast
}

func (v *FuncPtrAnalyzer) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	// Look for call expressions that might be function pointer casts
	if callExpr, ok := n.(*ast.CallExpr); ok {
		v.analyzeCallExpr(callExpr)
	}

	return v
}

func (v *FuncPtrAnalyzer) analyzeCallExpr(callExpr *ast.CallExpr) {
	// Look for pattern: (*(*funcType)(pointer))(args...)
	// The Fun should be a parenthesized dereference of a function pointer cast
	
	parenExpr, ok := callExpr.Fun.(*ast.ParenExpr)
	if !ok {
		return
	}

	starExpr, ok := parenExpr.X.(*ast.StarExpr)
	if !ok {
		return
	}

	// Now we need to find the function pointer cast inside starExpr.X
	cast := v.findFuncPtrCast(starExpr.X, callExpr)
	if cast != nil {
		cast.Position = v.fset.Position(callExpr.Pos())
		cast.SourceContext = v.getSourceContext(callExpr)
		
		// Extract call arguments
		for _, arg := range callExpr.Args {
			cast.CallArgs = append(cast.CallArgs, v.formatExpr(arg))
		}
		
		v.casts = append(v.casts, *cast)
	}
}

// findFuncPtrCast recursively searches for function pointer casts in an expression
func (v *FuncPtrAnalyzer) findFuncPtrCast(expr ast.Expr, callExpr *ast.CallExpr) *FuncPtrCast {
	switch e := expr.(type) {
	case *ast.ParenExpr:
		return v.findFuncPtrCast(e.X, callExpr)
		
	case *ast.CallExpr:
		// This could be a type cast: (*funcType)(pointer)
		if starExpr, ok := e.Fun.(*ast.ParenExpr); ok {
			if starType, ok := starExpr.X.(*ast.StarExpr); ok {
				if funcType, ok := starType.X.(*ast.FuncType); ok {
					// Found function pointer cast!
					cast := &FuncPtrCast{
						FuncType: v.formatFuncType(funcType),
					}
					
					// Extract the cast expression (what's being cast)
					if len(e.Args) > 0 {
						cast.CastExpr = v.formatExpr(e.Args[0])
					}
					
					return cast
				}
			}
		}
		
		// Also check function without parentheses: funcType(pointer)
		if funcType, ok := e.Fun.(*ast.FuncType); ok {
			cast := &FuncPtrCast{
				FuncType: v.formatFuncType(funcType),
			}
			
			if len(e.Args) > 0 {
				cast.CastExpr = v.formatExpr(e.Args[0])
			}
			
			return cast
		}
		
		// Check if the function part contains a function type cast
		return v.findFuncPtrCast(e.Fun, callExpr)
		
	case *ast.StarExpr:
		return v.findFuncPtrCast(e.X, callExpr)
		
	default:
		return nil
	}
}

func (v *FuncPtrAnalyzer) formatFuncType(funcType *ast.FuncType) string {
	var parts []string
	parts = append(parts, "func")

	// Parameters
	if funcType.Params != nil {
		var params []string
		for _, field := range funcType.Params.List {
			typeStr := v.formatExpr(field.Type)
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					params = append(params, name.Name+" "+typeStr)
				}
			} else {
				params = append(params, typeStr)
			}
		}
		parts = append(parts, "("+strings.Join(params, ", ")+")")
	} else {
		parts = append(parts, "()")
	}

	// Results
	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		var results []string
		for _, field := range funcType.Results.List {
			typeStr := v.formatExpr(field.Type)
			if len(field.Names) > 0 {
				for _, name := range field.Names {
					results = append(results, name.Name+" "+typeStr)
				}
			} else {
				results = append(results, typeStr)
			}
		}
		if len(results) == 1 && !strings.Contains(results[0], " ") {
			parts = append(parts, " "+results[0])
		} else {
			parts = append(parts, " ("+strings.Join(results, ", ")+")")
		}
	}

	return strings.Join(parts, "")
}

func (v *FuncPtrAnalyzer) formatExpr(expr ast.Expr) string {
	if expr == nil {
		return ""
	}

	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.SelectorExpr:
		return v.formatExpr(e.X) + "." + e.Sel.Name
	case *ast.CallExpr:
		fun := v.formatExpr(e.Fun)
		var args []string
		for _, arg := range e.Args {
			args = append(args, v.formatExpr(arg))
		}
		return fun + "(" + strings.Join(args, ", ") + ")"
	case *ast.ParenExpr:
		return "(" + v.formatExpr(e.X) + ")"
	case *ast.StarExpr:
		return "*" + v.formatExpr(e.X)
	case *ast.UnaryExpr:
		return e.Op.String() + v.formatExpr(e.X)
	case *ast.BinaryExpr:
		return v.formatExpr(e.X) + " " + e.Op.String() + " " + v.formatExpr(e.Y)
	case *ast.TypeAssertExpr:
		return v.formatExpr(e.X) + ".(" + v.formatExpr(e.Type) + ")"
	case *ast.IndexExpr:
		return v.formatExpr(e.X) + "[" + v.formatExpr(e.Index) + "]"
	case *ast.BasicLit:
		return e.Value
	case *ast.FuncType:
		return v.formatFuncType(e)
	case *ast.ArrayType:
		if e.Len != nil {
			return "[" + v.formatExpr(e.Len) + "]" + v.formatExpr(e.Elt)
		}
		return "[]" + v.formatExpr(e.Elt)
	case *ast.MapType:
		return "map[" + v.formatExpr(e.Key) + "]" + v.formatExpr(e.Value)
	case *ast.ChanType:
		switch e.Dir {
		case ast.SEND:
			return "chan<- " + v.formatExpr(e.Value)
		case ast.RECV:
			return "<-chan " + v.formatExpr(e.Value)
		default:
			return "chan " + v.formatExpr(e.Value)
		}
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{...}"
	case *ast.SliceExpr:
		result := v.formatExpr(e.X) + "["
		if e.Low != nil {
			result += v.formatExpr(e.Low)
		}
		result += ":"
		if e.High != nil {
			result += v.formatExpr(e.High)
		}
		if e.Slice3 && e.Max != nil {
			result += ":" + v.formatExpr(e.Max)
		}
		result += "]"
		return result
	default:
		// Try to create a reasonable representation for unknown types
		return fmt.Sprintf("<%T>", expr)
	}
}

func (v *FuncPtrAnalyzer) getSourceContext(expr ast.Expr) string {
	start := v.fset.Position(expr.Pos())
	end := v.fset.Position(expr.End())
	return fmt.Sprintf("Line %d:%d-%d:%d", start.Line, start.Column, end.Line, end.Column)
}