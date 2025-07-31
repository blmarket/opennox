package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func main() {
	var pointerFlag bool
	flag.BoolVar(&pointerFlag, "p", false, "handle as function pointer instead of function invocation")
	flag.Parse()

	args := flag.Args()
	if len(args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [-p] <module> <function_name>\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "  -p  handle as function pointer instead of function invocation\n")
		os.Exit(1)
	}

	moduleName := args[0]
	functionName := args[1]

	tool := &UpdateImportsTool{
		ModuleName:   moduleName,
		FunctionName: functionName,
		IsPointer:    pointerFlag,
	}

	if err := tool.Run(); err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Successfully updated imports for %s.%s\n", moduleName, functionName)
}

type UpdateImportsTool struct {
	ModuleName   string
	FunctionName string
	IsPointer    bool
}

func (t *UpdateImportsTool) Run() error {
	// 1. Find C function declaration
	decl, err := t.findCDeclaration()
	if err != nil {
		return fmt.Errorf("failed to find C declaration: %w", err)
	}

	// 2. Parse the declaration
	parsedDecl, err := t.parseCDeclaration(decl)
	if err != nil {
		return fmt.Errorf("failed to parse C declaration: %w", err)
	}

	// 3. Update imports file
	if err := t.updateImportsFile(parsedDecl); err != nil {
		return fmt.Errorf("failed to update imports file: %w", err)
	}

	// 4. Update module struct
	if err := t.updateModuleStruct(parsedDecl); err != nil {
		return fmt.Errorf("failed to update module struct: %w", err)
	}

	// 5. Update NewXxxModule constructor
	if err := t.updateModuleConstructor(parsedDecl); err != nil {
		return fmt.Errorf("failed to update module constructor: %w", err)
	}

	// 6. Update initXxx function
	if err := t.updateInitFunction(parsedDecl); err != nil {
		return fmt.Errorf("failed to update init function: %w", err)
	}

	return nil
}

// CDeclaration represents a parsed C function declaration
type CDeclaration struct {
	ReturnType string
	Name       string
	Parameters []CParameter
	RawDecl    string
}

type CParameter struct {
	Type string
	Name string
}

func (t *UpdateImportsTool) findCDeclaration() (string, error) {
	// Use the existing tools/find_decl.py tool
	cmd := exec.Command("python3", "-m", "tools.find_decl", t.FunctionName)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("find_decl.py failed: %w", err)
	}

	decl := strings.TrimSpace(string(output))
	if decl == "" {
		return "", fmt.Errorf("declaration not found for function %s", t.FunctionName)
	}

	return decl, nil
}

func (t *UpdateImportsTool) parseCDeclaration(decl string) (*CDeclaration, error) {
	// Remove semicolon and strip
	decl = strings.TrimSuffix(strings.TrimSpace(decl), ";")

	// Basic regex to match function declaration
	re := regexp.MustCompile(`^([^()]+?)\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)$`)
	matches := re.FindStringSubmatch(decl)
	if matches == nil {
		return nil, fmt.Errorf("invalid function declaration format: %s", decl)
	}

	returnType := strings.TrimSpace(matches[1])
	funcName := strings.TrimSpace(matches[2])
	paramsStr := strings.TrimSpace(matches[3])

	var parameters []CParameter
	if paramsStr != "" && paramsStr != "void" {
		paramParts := t.splitParameters(paramsStr)
		for i, param := range paramParts {
			param = strings.TrimSpace(param)
			if param != "" {
				paramMatch := regexp.MustCompile(`^(.+?)\s+([a-zA-Z_][a-zA-Z0-9_]*)$`).FindStringSubmatch(param)
				if paramMatch != nil {
					parameters = append(parameters, CParameter{
						Type: strings.TrimSpace(paramMatch[1]),
						Name: strings.TrimSpace(paramMatch[2]),
					})
				} else {
					// Parameter without name
					parameters = append(parameters, CParameter{
						Type: param,
						Name: fmt.Sprintf("arg%d", i),
					})
				}
			}
		}
	}

	return &CDeclaration{
		ReturnType: returnType,
		Name:       funcName,
		Parameters: parameters,
		RawDecl:    decl,
	}, nil
}

func (t *UpdateImportsTool) splitParameters(paramsStr string) []string {
	var params []string
	var current strings.Builder
	parenCount := 0

	for _, char := range paramsStr {
		if char == ',' && parenCount == 0 {
			params = append(params, current.String())
			current.Reset()
		} else {
			if char == '(' {
				parenCount++
			} else if char == ')' {
				parenCount--
			}
			current.WriteRune(char)
		}
	}

	if current.Len() > 0 {
		params = append(params, current.String())
	}

	return params
}

func (t *UpdateImportsTool) updateImportsFile(decl *CDeclaration) error {
	filename := fmt.Sprintf("%s_imports.go", t.ModuleName)

	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}

	lines := strings.Split(string(content), "\n")

	// Find the CGO comment block and add the declaration
	var updatedLines []string
	inCGOBlock := false
	cgoBlockEnd := -1

	for i, line := range lines {
		if strings.TrimSpace(line) == "/*" {
			inCGOBlock = true
		} else if strings.TrimSpace(line) == "*/" && inCGOBlock {
			// Insert the new declaration before the closing */
			updatedLines = append(updatedLines, decl.RawDecl+";")
			cgoBlockEnd = i
			inCGOBlock = false
		}
		updatedLines = append(updatedLines, line)
	}

	if cgoBlockEnd == -1 {
		return fmt.Errorf("CGO comment block not found in %s", filename)
	}

	// Write back to file
	err = os.WriteFile(filename, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return fmt.Errorf("failed to write %s: %w", filename, err)
	}

	return nil
}

func (t *UpdateImportsTool) updateModuleStruct(decl *CDeclaration) error {
	filename := fmt.Sprintf("%s/%s.go", t.ModuleName, t.ModuleName)

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", filename, err)
	}

	// Find the module struct and add the field
	structName := strings.Title(t.ModuleName) + "Module"

	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok && typeSpec.Name.Name == structName {
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				// Add the new field
				fieldName := t.FunctionName
				if t.IsPointer {
					fieldName += "_ptr"
				}

				var fieldType ast.Expr
				if t.IsPointer {
					// For function pointers, use unsafe.Pointer
					fieldType = &ast.SelectorExpr{
						X:   &ast.Ident{Name: "unsafe"},
						Sel: &ast.Ident{Name: "Pointer"},
					}
				} else {
					// For function invocations, create function type
					fieldType = t.createFunctionType(decl)
				}

				newField := &ast.Field{
					Names: []*ast.Ident{{Name: fieldName}},
					Type:  fieldType,
				}

				structType.Fields.List = append(structType.Fields.List, newField)
			}
		}
		return true
	})

	// Write the modified AST back to file
	return t.writeGoFile(filename, fset, node)
}

func (t *UpdateImportsTool) updateModuleConstructor(decl *CDeclaration) error {
	filename := fmt.Sprintf("%s/%s.go", t.ModuleName, t.ModuleName)

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", filename, err)
	}

	constructorName := "New" + strings.Title(t.ModuleName) + "Module"

	ast.Inspect(node, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Name.Name == constructorName {
			// Add parameter to function signature
			fieldName := t.FunctionName
			if t.IsPointer {
				fieldName += "_ptr"
			}

			var paramType ast.Expr
			if t.IsPointer {
				paramType = &ast.SelectorExpr{
					X:   &ast.Ident{Name: "unsafe"},
					Sel: &ast.Ident{Name: "Pointer"},
				}
			} else {
				paramType = t.createFunctionType(decl)
			}

			newParam := &ast.Field{
				Names: []*ast.Ident{{Name: fieldName}},
				Type:  paramType,
			}

			funcDecl.Type.Params.List = append(funcDecl.Type.Params.List, newParam)

			// Add field assignment in return statement
			if len(funcDecl.Body.List) > 0 {
				if retStmt, ok := funcDecl.Body.List[len(funcDecl.Body.List)-1].(*ast.ReturnStmt); ok {
					if len(retStmt.Results) > 0 {
						if unaryExpr, ok := retStmt.Results[0].(*ast.UnaryExpr); ok {
							if compLit, ok := unaryExpr.X.(*ast.CompositeLit); ok {
								newElt := &ast.KeyValueExpr{
									Key:   &ast.Ident{Name: fieldName},
									Value: &ast.Ident{Name: fieldName},
								}
								compLit.Elts = append(compLit.Elts, newElt)
							}
						}
					}
				}
			}
		}
		return true
	})

	return t.writeGoFile(filename, fset, node)
}

func (t *UpdateImportsTool) createFunctionType(decl *CDeclaration) ast.Expr {
	// Create parameters list
	var params []*ast.Field
	for _, param := range decl.Parameters {
		goType := t.mapCTypeToGo(param.Type)
		params = append(params, &ast.Field{
			Type: goType,
		})
	}

	// Create return type
	var results *ast.FieldList
	if decl.ReturnType != "void" {
		goReturnType := t.mapCTypeToGo(decl.ReturnType)
		results = &ast.FieldList{
			List: []*ast.Field{{Type: goReturnType}},
		}
	}

	return &ast.FuncType{
		Params:  &ast.FieldList{List: params},
		Results: results,
	}
}

func (t *UpdateImportsTool) mapCTypeToGo(cType string) ast.Expr {
	// Remove const, static, extern keywords
	cType = regexp.MustCompile(`\b(const|static|extern)\s+`).ReplaceAllString(cType, "")
	cType = strings.TrimSpace(cType)

	// Handle pointer types - most pointers become unsafe.Pointer
	if strings.Contains(cType, "*") {
		return &ast.SelectorExpr{
			X:   &ast.Ident{Name: "unsafe"},
			Sel: &ast.Ident{Name: "Pointer"},
		}
	}

	// Basic type mappings
	typeMap := map[string]string{
		"uint32_t": "uint32",
		"int32_t":  "int32",
		"uint16_t": "uint16",
		"int16_t":  "int16",
		"uint8_t":  "uint8",
		"int8_t":   "int8",
		"char":     "byte",
		"int":      "int",
		"short":    "int16",
		"long":     "int32",
		"float":    "float32",
		"double":   "float64",
		"void":     "",
	}

	if goType, exists := typeMap[cType]; exists {
		if goType == "" {
			return nil // void
		}
		return &ast.Ident{Name: goType}
	}

	// Default to unsafe.Pointer for unknown types
	return &ast.SelectorExpr{
		X:   &ast.Ident{Name: "unsafe"},
		Sel: &ast.Ident{Name: "Pointer"},
	}
}

func (t *UpdateImportsTool) updateInitFunction(decl *CDeclaration) error {
	filename := fmt.Sprintf("%s_imports.go", t.ModuleName)

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("failed to parse %s: %w", filename, err)
	}

	initFuncName := "init" + strings.Title(t.ModuleName)

	ast.Inspect(node, func(n ast.Node) bool {
		if funcDecl, ok := n.(*ast.FuncDecl); ok && funcDecl.Name.Name == initFuncName {
			// Find the constructor call
			ast.Inspect(funcDecl, func(callNode ast.Node) bool {
				if callExpr, ok := callNode.(*ast.CallExpr); ok {
					if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
						constructorName := "New" + strings.Title(t.ModuleName) + "Module"
						if selExpr.Sel.Name == constructorName {
							// Add the inline function as an argument
							if t.IsPointer {
								// For function pointers, pass unsafe.Pointer(C.functionName)
								arg := &ast.CallExpr{
									Fun: &ast.SelectorExpr{
										X:   &ast.Ident{Name: "unsafe"},
										Sel: &ast.Ident{Name: "Pointer"},
									},
									Args: []ast.Expr{
										&ast.SelectorExpr{
											X:   &ast.Ident{Name: "C"},
											Sel: &ast.Ident{Name: t.FunctionName},
										},
									},
								}
								callExpr.Args = append(callExpr.Args, arg)
							} else {
								// For function invocations, create inline function
								inlineFunc := t.createInlineFunction(decl)
								callExpr.Args = append(callExpr.Args, inlineFunc)
							}
						}
					}
				}
				return true
			})
		}
		return true
	})

	return t.writeGoFile(filename, fset, node)
}

func (t *UpdateImportsTool) createInlineFunction(decl *CDeclaration) ast.Expr {
	// Create parameter list for the inline function
	var params []*ast.Field
	var callArgs []ast.Expr

	for i, param := range decl.Parameters {
		paramName := param.Name
		if paramName == "" {
			paramName = fmt.Sprintf("arg%d", i)
		}

		goType := t.mapCTypeToGo(param.Type)
		params = append(params, &ast.Field{
			Names: []*ast.Ident{{Name: paramName}},
			Type:  goType,
		})

		// Create the call argument with proper type casting
		callArg := t.createCallArgument(param.Type, paramName)
		callArgs = append(callArgs, callArg)
	}

	// Create return type
	var results *ast.FieldList
	hasReturn := decl.ReturnType != "void"
	if hasReturn {
		goReturnType := t.mapCTypeToGo(decl.ReturnType)
		results = &ast.FieldList{
			List: []*ast.Field{{Type: goReturnType}},
		}
	}

	// Create the C function call
	cCall := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "C"},
			Sel: &ast.Ident{Name: t.FunctionName},
		},
		Args: callArgs,
	}

	// Create the function body
	var bodyStmts []ast.Stmt
	if hasReturn {
		// For non-void functions, return the result with type casting
		returnExpr := t.createReturnExpression(decl.ReturnType, cCall)
		bodyStmts = append(bodyStmts, &ast.ReturnStmt{
			Results: []ast.Expr{returnExpr},
		})
	} else {
		// For void functions, just call the function
		bodyStmts = append(bodyStmts, &ast.ExprStmt{X: cCall})
	}

	// Create the inline function literal
	return &ast.FuncLit{
		Type: &ast.FuncType{
			Params:  &ast.FieldList{List: params},
			Results: results,
		},
		Body: &ast.BlockStmt{List: bodyStmts},
	}
}

func (t *UpdateImportsTool) createCallArgument(cType, paramName string) ast.Expr {
	// Remove const, static, extern keywords
	cType = regexp.MustCompile(`\b(const|static|extern)\s+`).ReplaceAllString(cType, "")
	cType = strings.TrimSpace(cType)

	paramIdent := &ast.Ident{Name: paramName}

	// Handle pointer types
	if strings.Contains(cType, "*") {
		// Cast unsafe.Pointer to proper C type
		return &ast.CallExpr{
			Fun: &ast.ParenExpr{
				X: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "C"},
					Sel: &ast.Ident{Name: t.mapCTypeToCGO(cType)},
				},
			},
			Args: []ast.Expr{paramIdent},
		}
	}

	// Handle basic types - cast to C type
	cTypeName := t.mapCTypeToCGO(cType)
	return &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "C"},
			Sel: &ast.Ident{Name: cTypeName},
		},
		Args: []ast.Expr{paramIdent},
	}
}

func (t *UpdateImportsTool) createReturnExpression(returnType string, cCall ast.Expr) ast.Expr {
	// Remove const, static, extern keywords
	returnType = regexp.MustCompile(`\b(const|static|extern)\s+`).ReplaceAllString(returnType, "")
	returnType = strings.TrimSpace(returnType)

	// Handle pointer return types
	if strings.Contains(returnType, "*") {
		return &ast.CallExpr{
			Fun: &ast.SelectorExpr{
				X:   &ast.Ident{Name: "unsafe"},
				Sel: &ast.Ident{Name: "Pointer"},
			},
			Args: []ast.Expr{cCall},
		}
	}

	// Handle basic types
	goType := t.mapCTypeToGoString(returnType)
	if goType != "" {
		return &ast.CallExpr{
			Fun:  &ast.Ident{Name: goType},
			Args: []ast.Expr{cCall},
		}
	}

	// Default return as-is
	return cCall
}

func (t *UpdateImportsTool) mapCTypeToCGO(cType string) string {
	// Remove const, static, extern keywords
	cType = regexp.MustCompile(`\b(const|static|extern)\s+`).ReplaceAllString(cType, "")
	cType = strings.TrimSpace(cType)

	// For pointer types, return the pointer type name
	if strings.Contains(cType, "*") {
		// Extract base type and add pointer notation
		baseType := strings.TrimSuffix(cType, "*")
		baseType = strings.TrimSpace(baseType)
		return "*" + t.mapCTypeToCGO(baseType)
	}

	// Basic type mappings to C types
	typeMap := map[string]string{
		"uint32_t": "uint32_t",
		"int32_t":  "int32_t",
		"uint16_t": "uint16_t",
		"int16_t":  "int16_t",
		"uint8_t":  "uint8_t",
		"int8_t":   "int8_t",
		"char":     "char",
		"int":      "int",
		"short":    "short",
		"long":     "long",
		"float":    "float",
		"double":   "double",
		"void":     "void",
	}

	if cgoType, exists := typeMap[cType]; exists {
		return cgoType
	}

	// Default to the type as-is
	return cType
}

func (t *UpdateImportsTool) mapCTypeToGoString(cType string) string {
	// Remove const, static, extern keywords
	cType = regexp.MustCompile(`\b(const|static|extern)\s+`).ReplaceAllString(cType, "")
	cType = strings.TrimSpace(cType)

	// Basic type mappings
	typeMap := map[string]string{
		"uint32_t": "uint32",
		"int32_t":  "int32",
		"uint16_t": "uint16",
		"int16_t":  "int16",
		"uint8_t":  "uint8",
		"int8_t":   "int8",
		"char":     "byte",
		"int":      "int",
		"short":    "int16",
		"long":     "int32",
		"float":    "float32",
		"double":   "float64",
		"void":     "",
	}

	if goType, exists := typeMap[cType]; exists {
		return goType
	}

	return ""
}

func (t *UpdateImportsTool) writeGoFile(filename string, fset *token.FileSet, node *ast.File) error {
	// Create a buffer to write the formatted Go code
	var buf strings.Builder
	err := format.Node(&buf, fset, node)
	if err != nil {
		return fmt.Errorf("failed to format Go code: %w", err)
	}

	// Write to file
	return os.WriteFile(filename, []byte(buf.String()), 0644)
}
