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

// titleCase converts a string to title case (first letter uppercase)
func titleCase(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
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

	updatesNeeded := false

	// 3. Update imports file
	importsUpdated, err := t.updateImportsFile(parsedDecl)
	if err != nil {
		return fmt.Errorf("failed to update imports file: %w", err)
	}
	if importsUpdated {
		updatesNeeded = true
	}

	// 4. Update module struct
	structUpdated, err := t.updateModuleStruct(parsedDecl)
	if err != nil {
		return fmt.Errorf("failed to update module struct: %w", err)
	}
	if structUpdated {
		updatesNeeded = true
	}

	// 5. Update NewXxxModule constructor
	constructorUpdated, err := t.updateModuleConstructor(parsedDecl)
	if err != nil {
		return fmt.Errorf("failed to update module constructor: %w", err)
	}
	if constructorUpdated {
		updatesNeeded = true
	}

	// 6. Update initXxx function (only if other updates were made)
	if updatesNeeded {
		if err := t.updateInitFunction(parsedDecl); err != nil {
			return fmt.Errorf("failed to update init function: %w", err)
		}
	}

	if !updatesNeeded {
		fmt.Printf("No updates needed - %s already exists in all required locations\n", t.FunctionName)
	}

	return nil
}

// CDeclaration represents a parsed C function or variable declaration
type CDeclaration struct {
	ReturnType string
	Name       string
	Parameters []CParameter
	RawDecl    string
	IsVariable bool   // true if this is a variable declaration, false if function
	VarType    string // variable type for variable declarations
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

	// Check if this is a function declaration (contains parentheses)
	if strings.Contains(decl, "(") && strings.Contains(decl, ")") {
		return t.parseFunctionDeclaration(decl)
	} else {
		return t.parseVariableDeclaration(decl)
	}
}

func (t *UpdateImportsTool) parseFunctionDeclaration(decl string) (*CDeclaration, error) {
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
		IsVariable: false,
	}, nil
}

func (t *UpdateImportsTool) parseVariableDeclaration(decl string) (*CDeclaration, error) {
	// Parse variable declaration: extern type name or type name
	// Remove extern keyword if present
	cleaned := regexp.MustCompile(`\bextern\s+`).ReplaceAllString(decl, "")
	cleaned = strings.TrimSpace(cleaned)

	// Match variable declaration: type name
	re := regexp.MustCompile(`^(.+?)\s+([a-zA-Z_][a-zA-Z0-9_]*(?:\[[^\]]*\])*)$`)
	matches := re.FindStringSubmatch(cleaned)
	if matches == nil {
		return nil, fmt.Errorf("invalid variable declaration format: %s", decl)
	}

	varType := strings.TrimSpace(matches[1])
	varName := strings.TrimSpace(matches[2])

	return &CDeclaration{
		Name:       varName,
		VarType:    varType,
		RawDecl:    decl,
		IsVariable: true,
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

func (t *UpdateImportsTool) updateImportsFile(decl *CDeclaration) (bool, error) {
	filename := fmt.Sprintf("%s_imports.go", t.ModuleName)

	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, fmt.Errorf("failed to read %s: %w", filename, err)
	}

	lines := strings.Split(string(content), "\n")
	declToAdd := decl.RawDecl + ";"

	// Check if declaration already exists
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == declToAdd || trimmed == decl.RawDecl {
			fmt.Printf("Declaration %s already exists in CGO block, skipping\n", decl.Name)
			return false, nil
		}
	}

	// Find the CGO comment block and add the declaration
	var updatedLines []string
	inCGOBlock := false
	cgoBlockEnd := -1

	for i, line := range lines {
		if strings.TrimSpace(line) == "/*" {
			inCGOBlock = true
		} else if strings.TrimSpace(line) == "*/" && inCGOBlock {
			// Insert the new declaration before the closing */
			updatedLines = append(updatedLines, declToAdd)
			cgoBlockEnd = i
			inCGOBlock = false
		}
		updatedLines = append(updatedLines, line)
	}

	if cgoBlockEnd == -1 {
		return false, fmt.Errorf("CGO comment block not found in %s", filename)
	}

	// Write back to file
	err = os.WriteFile(filename, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return false, fmt.Errorf("failed to write %s: %w", filename, err)
	}

	return true, nil
}

func (t *UpdateImportsTool) updateModuleStruct(decl *CDeclaration) (bool, error) {
	filename := fmt.Sprintf("%s/%s.go", t.ModuleName, t.ModuleName)

	// Parse the Go file
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return false, fmt.Errorf("failed to parse %s: %w", filename, err)
	}

	// Find the module struct and add the field
	structName := titleCase(t.ModuleName) + "Module"
	fieldExists := false

	ast.Inspect(node, func(n ast.Node) bool {
		if typeSpec, ok := n.(*ast.TypeSpec); ok && typeSpec.Name.Name == structName {
			if structType, ok := typeSpec.Type.(*ast.StructType); ok {
				// Check if field already exists
				fieldName := t.FunctionName
				if !decl.IsVariable && t.IsPointer {
					fieldName += "_ptr"
				}

				for _, field := range structType.Fields.List {
					for _, name := range field.Names {
						if name.Name == fieldName {
							fmt.Printf("Field %s already exists in struct %s, skipping\n", fieldName, structName)
							fieldExists = true
							return false
						}
					}
				}

				// Add the new field
				var fieldType ast.Expr

				if decl.IsVariable {
					// For variables, use pointer to the Go type
					goType := t.mapCTypeToGo(decl.VarType)
					if goType != nil {
						if selExpr, ok := goType.(*ast.SelectorExpr); ok && selExpr.Sel.Name == "Pointer" {
							// Already a pointer type (unsafe.Pointer)
							fieldType = goType
						} else {
							// Create pointer to the type
							fieldType = &ast.StarExpr{X: goType}
						}
					} else {
						// Default to unsafe.Pointer for unknown types
						fieldType = &ast.SelectorExpr{
							X:   &ast.Ident{Name: "unsafe"},
							Sel: &ast.Ident{Name: "Pointer"},
						}
					}
				} else {
					// For functions
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

	if fieldExists {
		return false, nil
	}

	// Write the modified AST back to file
	return true, t.writeGoFile(filename, fset, node)
}

func (t *UpdateImportsTool) updateModuleConstructor(decl *CDeclaration) (bool, error) {
	filename := fmt.Sprintf("%s/%s.go", t.ModuleName, t.ModuleName)

	// Read the file as text
	content, err := os.ReadFile(filename)
	if err != nil {
		return false, fmt.Errorf("failed to read %s: %w", filename, err)
	}

	lines := strings.Split(string(content), "\n")
	constructorName := "New" + titleCase(t.ModuleName) + "Module"
	
	fieldName := t.FunctionName
	if !decl.IsVariable && t.IsPointer {
		fieldName += "_ptr"
	}

	// Check if parameter already exists in constructor function signature
	inConstructorParams := false
	constructorFound := false
	for _, line := range lines {
		if strings.Contains(line, "func "+constructorName+"(") {
			constructorFound = true
			inConstructorParams = true
		} else if inConstructorParams && strings.Contains(line, ") *"+titleCase(t.ModuleName)+"Module") {
			inConstructorParams = false
		} else if inConstructorParams {
			// Check if this line contains our parameter
			if strings.Contains(line, fieldName+" ") {
				fmt.Printf("Parameter %s already exists in constructor %s, skipping\n", fieldName, constructorName)
				return false, nil
			}
		}
		// Stop searching once we've passed the constructor
		if constructorFound && !inConstructorParams {
			break
		}
	}

	// Find constructor function and add parameter + field assignment
	var updatedLines []string
	inConstructor := false
	inParams := false
	inReturn := false
	
	for _, line := range lines {
		
		// Detect constructor function start
		if strings.Contains(line, "func "+constructorName+"(") {
			inConstructor = true
			inParams = true
		}
		
		// Detect end of parameters
		if inConstructor && inParams && strings.Contains(line, ") *"+titleCase(t.ModuleName)+"Module") {
			// Add new parameter before closing parenthesis
			paramType := t.getGoTypeString(decl)
			newParam := fmt.Sprintf("\t%s %s,", fieldName, paramType)
			
			// Insert parameter before the closing line
			updatedLines = append(updatedLines, newParam)
			inParams = false
		}
		
		// Detect start of return statement
		if inConstructor && strings.Contains(line, "return &"+titleCase(t.ModuleName)+"Module{") {
			inReturn = true
		}
		
		// Detect end of return statement
		if inConstructor && inReturn && strings.Contains(line, "}") && !strings.Contains(line, "{") {
			// Add field assignment before closing brace
			fieldAssignment := fmt.Sprintf("\t\t%s: %s,", fieldName, fieldName)
			updatedLines = append(updatedLines, fieldAssignment)
			inConstructor = false
			inReturn = false
		}
		
		updatedLines = append(updatedLines, line)
	}

	// Write back to file
	err = os.WriteFile(filename, []byte(strings.Join(updatedLines, "\n")), 0644)
	if err != nil {
		return false, fmt.Errorf("failed to write %s: %w", filename, err)
	}

	return true, nil
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

	// Read the file as text
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read %s: %w", filename, err)
	}

	lines := strings.Split(string(content), "\n")
	constructorName := "New" + titleCase(t.ModuleName) + "Module"
	
	// Find the last line of the constructor call (the one with closing parenthesis)
	constructorEndLine := -1
	for i, line := range lines {
		if strings.Contains(line, constructorName+"(") {
			// Found constructor start, now find the end
			for j := i; j < len(lines); j++ {
				if strings.TrimSpace(lines[j]) == ")" {
					constructorEndLine = j
					break
				}
			}
			break
		}
	}
	
	if constructorEndLine == -1 {
		return fmt.Errorf("could not find constructor call end in %s", filename)
	}

	// Check if the argument is already there
	for i := constructorEndLine; i >= 0; i-- {
		if strings.Contains(lines[i], t.FunctionName) {
			fmt.Printf("Argument %s already exists in constructor call, skipping\n", t.FunctionName)
			return nil
		}
		if strings.Contains(lines[i], constructorName+"(") {
			break
		}
	}

	// Insert the new argument before the closing parenthesis
	var newArg string
	if decl.IsVariable {
		// For variables, add variable declaration first in var block
		varFound := false
		varBlockStart := -1
		varBlockEnd := -1
		
		// Find var block inside init function
		initFuncName := "init" + titleCase(t.ModuleName)
		inInitFunc := false
		for i, line := range lines {
			if strings.Contains(line, "func "+initFuncName+"()") {
				inInitFunc = true
			} else if inInitFunc && strings.TrimSpace(line) == "var (" {
				varBlockStart = i
			} else if inInitFunc && varBlockStart != -1 && strings.TrimSpace(line) == ")" {
				varBlockEnd = i
				break
			} else if inInitFunc && strings.TrimSpace(line) == "}" && !strings.Contains(line, "{") {
				// End of init function
				break
			}
		}
		
		// Check if variable already exists in var block
		if varBlockStart != -1 && varBlockEnd != -1 {
			for i := varBlockStart; i < varBlockEnd; i++ {
				if strings.Contains(lines[i], t.FunctionName+" ") {
					varFound = true
					break
				}
			}
		}
		
		// Add variable to var block if not found
		if !varFound && varBlockEnd != -1 {
			varType := t.getGoTypeString(decl)
			var varDecl string
			if strings.Contains(varType, "unsafe.Pointer") {
				varDecl = fmt.Sprintf("\t\t%s %s = &C.%s", t.FunctionName, varType, t.FunctionName)
			} else {
				varDecl = fmt.Sprintf("\t\t%s %s = (%s)(&C.%s)", t.FunctionName, varType, varType, t.FunctionName)
			}
			
			// Insert the variable declaration
			updatedLines := make([]string, 0, len(lines)+1)
			updatedLines = append(updatedLines, lines[:varBlockEnd]...)
			updatedLines = append(updatedLines, varDecl)
			updatedLines = append(updatedLines, lines[varBlockEnd:]...)
			lines = updatedLines
			constructorEndLine++ // Adjust constructor end line due to insertion
		}
		
		newArg = fmt.Sprintf("\t\t%s,", t.FunctionName)
	} else {
		if t.IsPointer {
			// For function pointers
			newArg = fmt.Sprintf("\t\tunsafe.Pointer(C.%s),", t.FunctionName)
		} else {
			// For inline functions
			inlineFunc := t.createInlineFunctionString(decl)
			newArg = fmt.Sprintf("\t\t%s,", inlineFunc)
		}
	}

	// Insert the new argument before the closing parenthesis
	updatedLines := make([]string, 0, len(lines)+1)
	updatedLines = append(updatedLines, lines[:constructorEndLine]...)
	updatedLines = append(updatedLines, newArg)
	updatedLines = append(updatedLines, lines[constructorEndLine:]...)

	// Write back to file
	return os.WriteFile(filename, []byte(strings.Join(updatedLines, "\n")), 0644)
}

func (t *UpdateImportsTool) addVariableToInitFunction(funcDecl *ast.FuncDecl, decl *CDeclaration) bool {
	varName := t.FunctionName

	// Check if variable already exists in var block
	for _, stmt := range funcDecl.Body.List {
		if declStmt, ok := stmt.(*ast.DeclStmt); ok {
			if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
				for _, spec := range genDecl.Specs {
					if valueSpec, ok := spec.(*ast.ValueSpec); ok {
						for _, name := range valueSpec.Names {
							if name.Name == varName {
								fmt.Printf("Variable %s already exists in init function, skipping\n", varName)
								return false
							}
						}
					}
				}
			}
		}
	}

	// Add variable declaration in var block
	for _, stmt := range funcDecl.Body.List {
		if declStmt, ok := stmt.(*ast.DeclStmt); ok {
			if genDecl, ok := declStmt.Decl.(*ast.GenDecl); ok && genDecl.Tok == token.VAR {
				// Add new variable declaration
				goType := t.mapCTypeToGo(decl.VarType)
				var varType ast.Expr
				
				if goType != nil && goType.(*ast.SelectorExpr) != nil && goType.(*ast.SelectorExpr).Sel.Name == "Pointer" {
					// Already a pointer type (unsafe.Pointer)
					varType = goType
				} else {
					// Create pointer to the type
					varType = &ast.StarExpr{X: goType}
				}

				// Create the variable assignment
				var valueExpr ast.Expr
				if goType != nil && goType.(*ast.SelectorExpr) != nil && goType.(*ast.SelectorExpr).Sel.Name == "Pointer" {
					// For unsafe.Pointer types
					valueExpr = &ast.CallExpr{
						Fun: &ast.SelectorExpr{
							X:   &ast.Ident{Name: "unsafe"},
							Sel: &ast.Ident{Name: "Pointer"},
						},
						Args: []ast.Expr{
							&ast.UnaryExpr{
								Op: token.AND,
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "C"},
									Sel: &ast.Ident{Name: varName},
								},
							},
						},
					}
				} else {
					// For typed pointers
					valueExpr = &ast.CallExpr{
						Fun: &ast.ParenExpr{X: varType},
						Args: []ast.Expr{
							&ast.UnaryExpr{
								Op: token.AND,
								X: &ast.SelectorExpr{
									X:   &ast.Ident{Name: "C"},
									Sel: &ast.Ident{Name: varName},
								},
							},
						},
					}
				}

				newSpec := &ast.ValueSpec{
					Names:  []*ast.Ident{{Name: varName}},
					Type:   varType,
					Values: []ast.Expr{valueExpr},
				}

				genDecl.Specs = append(genDecl.Specs, newSpec)
				break
			}
		}
	}

	// Add variable to constructor call
	t.addArgumentToConstructorCall(funcDecl, &ast.Ident{Name: varName})
	return true
}

func (t *UpdateImportsTool) addFunctionToConstructorCall(funcDecl *ast.FuncDecl, decl *CDeclaration) bool {
	found := false
	ast.Inspect(funcDecl, func(callNode ast.Node) bool {
		if callExpr, ok := callNode.(*ast.CallExpr); ok {
			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				constructorName := "New" + titleCase(t.ModuleName) + "Module"
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
					found = true
				}
			}
		}
		return true
	})
	return found
}

func (t *UpdateImportsTool) addArgumentToConstructorCall(funcDecl *ast.FuncDecl, arg ast.Expr) {
	ast.Inspect(funcDecl, func(callNode ast.Node) bool {
		if callExpr, ok := callNode.(*ast.CallExpr); ok {
			if selExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				constructorName := "New" + titleCase(t.ModuleName) + "Module"
				if selExpr.Sel.Name == constructorName {
					callExpr.Args = append(callExpr.Args, arg)
				}
			}
		}
		return true
	})
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

func (t *UpdateImportsTool) getGoTypeString(decl *CDeclaration) string {
	if decl.IsVariable {
		// For variables, use pointer to the Go type
		goType := t.mapCTypeToGoString(decl.VarType)
		if strings.Contains(decl.VarType, "*") {
			return "unsafe.Pointer"
		}
		if goType != "" {
			return "*" + goType
		}
		return "*unsafe.Pointer"
	} else {
		// For functions
		if t.IsPointer {
			return "unsafe.Pointer"
		} else {
			// Create function type string
			var params []string
			for _, param := range decl.Parameters {
				goType := t.mapCTypeToGoString(param.Type)
				if strings.Contains(param.Type, "*") || goType == "" {
					params = append(params, "unsafe.Pointer")
				} else {
					params = append(params, goType)
				}
			}
			
			returnType := ""
			if decl.ReturnType != "void" {
				goReturnType := t.mapCTypeToGoString(decl.ReturnType)
				if strings.Contains(decl.ReturnType, "*") || goReturnType == "" {
					returnType = " unsafe.Pointer"
				} else {
					returnType = " " + goReturnType
				}
			}
			
			return fmt.Sprintf("func(%s)%s", strings.Join(params, ", "), returnType)
		}
	}
}

func (t *UpdateImportsTool) createInlineFunctionString(decl *CDeclaration) string {
	// Create parameter list for the inline function
	var params []string
	var callArgs []string

	for i, param := range decl.Parameters {
		paramName := param.Name
		if paramName == "" {
			paramName = fmt.Sprintf("arg%d", i)
		}

		goType := t.mapCTypeToGoString(param.Type)
		if strings.Contains(param.Type, "*") || goType == "" {
			params = append(params, paramName+" unsafe.Pointer")
			// For pointer parameters, cast the unsafe.Pointer to the C type
			cgoType := t.mapCTypeToCGO(param.Type)
			callArgs = append(callArgs, fmt.Sprintf("(*C.%s)(%s)", strings.TrimPrefix(cgoType, "*"), paramName))
		} else {
			params = append(params, paramName+" "+goType)
			callArgs = append(callArgs, fmt.Sprintf("C.%s(%s)", t.mapCTypeToCGO(param.Type), paramName))
		}
	}

	// Create return type
	hasReturn := decl.ReturnType != "void"
	returnType := ""
	if hasReturn {
		goReturnType := t.mapCTypeToGoString(decl.ReturnType)
		if strings.Contains(decl.ReturnType, "*") || goReturnType == "" {
			returnType = " unsafe.Pointer"
		} else {
			returnType = " " + goReturnType
		}
	}

	// Create the C function call
	cCall := fmt.Sprintf("C.%s(%s)", t.FunctionName, strings.Join(callArgs, ", "))

	// Create the function body
	var body string
	if hasReturn {
		// For non-void functions, return the result with type casting
		if strings.Contains(decl.ReturnType, "*") {
			body = fmt.Sprintf("return unsafe.Pointer(%s)", cCall)
		} else {
			goReturnType := t.mapCTypeToGoString(decl.ReturnType)
			if goReturnType != "" {
				body = fmt.Sprintf("return %s(%s)", goReturnType, cCall)
			} else {
				body = fmt.Sprintf("return %s", cCall)
			}
		}
	} else {
		// For void functions, just call the function
		body = cCall
	}

	return fmt.Sprintf("func(%s)%s { %s }", strings.Join(params, ", "), returnType, body)
}

func (t *UpdateImportsTool) writeGoFile(filename string, fset *token.FileSet, node *ast.File) error {
	// Create a buffer to write the formatted Go code
	var buf strings.Builder
	err := format.Node(&buf, fset, node)
	if err != nil {
		// If formatting fails, it's often due to AST formatting issues
		// In this case, fall back to reading the original file and doing string-based edits
		return fmt.Errorf("failed to format Go code: %w", err)
	}

	// Write to file
	return os.WriteFile(filename, []byte(buf.String()), 0644)
}
