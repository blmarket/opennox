package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <file.go>\n", os.Args[0])
		os.Exit(1)
	}

	filename := os.Args[1]

	// Parse the Go file to extract module methods
	moduleMethods, err := extractModuleMethods(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d module methods\n", len(moduleMethods))
	for method := range moduleMethods {
		fmt.Printf("  - %s\n", method)
	}

	// Fix function calls in the file
	err = fixFunctionCalls(filename, moduleMethods)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fixing function calls: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully fixed function calls in %s\n", filename)
}

// extractModuleMethods parses both files to extract all method names and external function names
func extractModuleMethods(filename string) (map[string]bool, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	methods := make(map[string]bool)

	// Extract module methods from the _impl.go file
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// Check if this is a method on *Struct264Module
			if x.Recv != nil && len(x.Recv.List) > 0 {
				if starExpr, ok := x.Recv.List[0].Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == "Struct264Module" {
						methods[x.Name.Name] = true
					}
				}
			}
		}
		return true
	})

	// Also parse the struct264.go file to find external functions
	structFile := strings.Replace(filename, "_impl.go", ".go", 1)
	if structFile != filename {
		structNode, err := parser.ParseFile(fset, structFile, nil, parser.ParseComments)
		if err == nil {
			ast.Inspect(structNode, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.StructType:
					// Look for external function fields in Struct264Module
					for _, field := range x.Fields.List {
						if field.Type != nil {
							if funcType, ok := field.Type.(*ast.FuncType); ok && funcType != nil {
								// This is a function field
								for _, name := range field.Names {
									methods[name.Name] = true
								}
							}
						}
					}
				}
				return true
			})
		}
	}

	return methods, nil
}

// fixFunctionCalls reads the file, identifies function calls that should be module calls, and fixes them
func fixFunctionCalls(filename string, moduleMethods map[string]bool) error {
	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Regular expression to match function calls that are not already prefixed with m.
	// This matches: function_name( but not m.function_name( or other.function_name(
	// Also excludes function declarations
	callRegex := regexp.MustCompile(`(?:^|[^a-zA-Z_.])([a-zA-Z_][a-zA-Z0-9_]*)\s*\(`)

	modified := false
	for i, line := range lines {
		originalLine := line

		// Skip function declarations - look for "func " pattern
		if strings.Contains(line, "func (") || strings.Contains(line, "func ") {
			continue
		}

		// Find all function calls in the line
		matches := callRegex.FindAllStringSubmatch(line, -1)

		// Process matches in reverse order to avoid offset issues when replacing
		for j := len(matches) - 1; j >= 0; j-- {
			match := matches[j]
			funcName := match[1]

			// Check if this function is a module method
			if moduleMethods[funcName] {
				// Find the position and replace with m.funcName
				fullMatch := match[0]
				pos := strings.LastIndex(line, fullMatch)
				if pos >= 0 {
					// Check if it's already prefixed with m. or another identifier
					prefix := ""
					if pos > 0 {
						prefix = line[:pos+len(fullMatch)-len(funcName)-1] // Everything before funcName(
					} else {
						prefix = line[:len(fullMatch)-len(funcName)-1]
					}

					// Only replace if it's not already prefixed with an identifier followed by dot
					// and it's not already m.funcName
					if !strings.HasSuffix(strings.TrimSpace(prefix), ".") && !strings.Contains(fullMatch, "m."+funcName) {
						// Handle case where match starts with non-identifier character
						beforeFunc := strings.TrimSuffix(fullMatch, funcName+"(")
						newCall := beforeFunc + "m." + funcName + "("
						line = line[:pos] + newCall + line[pos+len(fullMatch):]
					}
				}
			}
		}

		if line != originalLine {
			lines[i] = line
			modified = true
			fmt.Printf("Line %d: %s\n", i+1, strings.TrimSpace(line))
		}
	}

	if !modified {
		fmt.Println("No changes needed")
		return nil
	}

	// Write the modified content back to the file
	outFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer outFile.Close()

	writer := bufio.NewWriter(outFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}
