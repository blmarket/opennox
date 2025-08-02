package main

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"
)

// FunctionInfo holds information about a function
type FunctionInfo struct {
	Name       string
	Parameters []Parameter
	ReturnType string
	HasReturn  bool
}

// Parameter represents a function parameter
type Parameter struct {
	Name string
	Type string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run expose_function.go <function_name>")
		os.Exit(1)
	}

	functionName := os.Args[1]
	implFile := "struct312/struct312_impl.go"
	importsFile := "struct312_imports.go"

	// Find and parse the function in impl file
	funcInfo, err := findFunction(implFile, functionName)
	if err != nil {
		log.Fatalf("Error finding function: %v", err)
	}

	// Make the function name uppercase in impl file
	err = makeUppercase(implFile, functionName)
	if err != nil {
		log.Fatalf("Error making function uppercase: %v", err)
	}

	// Add export entry to imports file
	err = addExportEntry(importsFile, funcInfo)
	if err != nil {
		log.Fatalf("Error adding export entry: %v", err)
	}

	fmt.Printf("Successfully exposed function '%s' as '%s'\n", functionName, strings.Title(functionName))
}

// findFunction parses the Go file and finds the specified function
func findFunction(filename, functionName string) (*FunctionInfo, error) {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %v", err)
	}

	var funcInfo *FunctionInfo

	ast.Inspect(node, func(n ast.Node) bool {
		if fn, ok := n.(*ast.FuncDecl); ok {
			if fn.Name.Name == functionName {
				funcInfo = &FunctionInfo{
					Name: functionName,
				}

				// Extract parameters
				if fn.Type.Params != nil {
					for _, param := range fn.Type.Params.List {
						paramType := getTypeString(param.Type)
						for _, name := range param.Names {
							funcInfo.Parameters = append(funcInfo.Parameters, Parameter{
								Name: name.Name,
								Type: paramType,
							})
						}
					}
				}

				// Extract return type
				if fn.Type.Results != nil && len(fn.Type.Results.List) > 0 {
					funcInfo.HasReturn = true
					result := fn.Type.Results.List[0]
					funcInfo.ReturnType = getTypeString(result.Type)
				}
			}
		}
		return true
	})

	if funcInfo == nil {
		return nil, fmt.Errorf("function '%s' not found", functionName)
	}

	return funcInfo, nil
}

// getTypeString converts an ast.Expr to a string representation
func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + getTypeString(t.X)
	case *ast.SelectorExpr:
		return getTypeString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		return "[]" + getTypeString(t.Elt)
	default:
		return "interface{}" // fallback
	}
}

// makeUppercase changes the first letter of the function name to uppercase
func makeUppercase(filename, functionName string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	// Create regex to match the function declaration
	funcRegex := regexp.MustCompile(`^func\s+\(m\s+\*Struct312Module\)\s+` + regexp.QuoteMeta(functionName) + `\s*\(`)

	for scanner.Scan() {
		line := scanner.Text()
		if funcRegex.MatchString(line) {
			// Replace the function name with uppercase version
			newFuncName := strings.Title(functionName)
			line = strings.Replace(line, functionName+"(", newFuncName+"(", 1)
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write back to file
	output := strings.Join(lines, "\n")
	return os.WriteFile(filename, []byte(output), 0644)
}

// addExportEntry adds the export function to the imports file
func addExportEntry(filename string, funcInfo *FunctionInfo) error {
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

	// Generate the export function
	exportFunc := generateExportFunction(funcInfo)

	// Add the export function at the end of the file
	lines = append(lines, "")
	lines = append(lines, exportFunc...)

	// Write back to file
	output := strings.Join(lines, "\n")
	return os.WriteFile(filename, []byte(output), 0644)
}

// generateExportFunction creates the export function code
func generateExportFunction(funcInfo *FunctionInfo) []string {
	upperFuncName := strings.Title(funcInfo.Name)

	var lines []string

	// Add export comment
	lines = append(lines, fmt.Sprintf("//export %s", funcInfo.Name))

	// Build function signature
	var params []string
	var args []string

	for _, param := range funcInfo.Parameters {
		params = append(params, fmt.Sprintf("%s %s", param.Name, param.Type))
		args = append(args, param.Name)
	}

	paramStr := strings.Join(params, ", ")
	argStr := strings.Join(args, ", ")

	// Build function declaration
	var funcDecl string
	if funcInfo.HasReturn {
		funcDecl = fmt.Sprintf("func %s(%s) %s {", funcInfo.Name, paramStr, funcInfo.ReturnType)
		lines = append(lines, funcDecl)
		lines = append(lines, fmt.Sprintf("\treturn Struct312Module.%s(%s)", upperFuncName, argStr))
	} else {
		funcDecl = fmt.Sprintf("func %s(%s) {", funcInfo.Name, paramStr)
		lines = append(lines, funcDecl)
		lines = append(lines, fmt.Sprintf("\tStruct312Module.%s(%s)", upperFuncName, argStr))
	}

	lines = append(lines, "}")

	return lines
}
