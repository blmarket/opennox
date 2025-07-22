#!/usr/bin/env python3
"""
C Function Parameter Modifier

This script modifies C function signatures by changing a specific parameter's type
and name, then adds appropriate variable declarations in the function body.

Usage:
    python modify_c_function.py <function_name> <param_index> <new_type> <file_path>

Example:
    python modify_c_function.py myFunction 1 "int*" src/myfile.c
"""

import argparse
import sys
import os
from typing import List, Optional, Tuple
from pycparser import c_parser, c_ast, c_generator
from pycparser.plyparser import ParseError


class CFunctionModifier:
    def __init__(self):
        self.parser = c_parser.CParser()
        self.generator = c_generator.CGenerator()
    
    def parse_file(self, filepath: str) -> c_ast.FileAST:
        """Parse a C file and return the AST."""
        try:
            with open(filepath, 'r') as f:
                content = f.read()
            
            # Remove preprocessor directives and add basic type definitions
            lines = content.split('\n')
            filtered_lines = []
            for line in lines:
                # Skip preprocessor directives
                if line.strip().startswith('#'):
                    continue
                filtered_lines.append(line)
            
            filtered_content = '\n'.join(filtered_lines)
            
            # Add basic type definitions without any preprocessor directives
            preprocessed = f"""typedef unsigned int size_t;
typedef int ssize_t;
typedef unsigned char uint8_t;
typedef unsigned short uint16_t;
typedef unsigned int uint32_t;
typedef unsigned long long uint64_t;
typedef signed char int8_t;
typedef short int16_t;
typedef int int32_t;
typedef long long int64_t;

{filtered_content}
"""
            
            # Parse the processed content
            
            return self.parser.parse(preprocessed, filepath)
        except ParseError as e:
            print(f"Error parsing {filepath}: {e}")
            sys.exit(1)
        except FileNotFoundError:
            print(f"File not found: {filepath}")
            sys.exit(1)
    
    def find_function_definition(self, ast: c_ast.FileAST, function_name: str) -> Optional[c_ast.FuncDef]:
        """Find a function definition by name in the AST."""
        class FunctionFinder(c_ast.NodeVisitor):
            def __init__(self):
                self.found_function = None
            
            def visit_FuncDef(self, node):
                if node.decl.name == function_name:
                    self.found_function = node
                self.generic_visit(node)
        
        finder = FunctionFinder()
        finder.visit(ast)
        return finder.found_function
    
    def find_function_declaration(self, ast: c_ast.FileAST, function_name: str) -> Optional[c_ast.Decl]:
        """Find a function declaration by name in the AST."""
        class FunctionDeclFinder(c_ast.NodeVisitor):
            def __init__(self):
                self.found_declarations = []
            
            def visit_Decl(self, node):
                if (isinstance(node.type, c_ast.FuncDecl) and 
                    node.name == function_name):
                    self.found_declarations.append(node)
                self.generic_visit(node)
        
        finder = FunctionDeclFinder()
        finder.visit(ast)
        return finder.found_declarations[0] if finder.found_declarations else None
    
    def get_parameter_info(self, func_node: c_ast.FuncDef, param_index: int) -> Optional[Tuple[str, str]]:
        """Get the original type and name of a parameter at the given index."""
        if not func_node.decl.type.args:
            return None
        
        params = func_node.decl.type.args.params
        if param_index >= len(params):
            return None
        
        param = params[param_index]
        if not isinstance(param, c_ast.Decl):
            return None
        
        # Extract type information
        param_type = self.generator.visit(param.type)
        param_name = param.name
        
        return param_type, param_name
    
    def get_parameter_info_from_decl(self, func_decl: c_ast.Decl, param_index: int) -> Optional[Tuple[str, str]]:
        """Get the original type and name of a parameter from a function declaration."""
        if not func_decl.type.args:
            return None
        
        params = func_decl.type.args.params
        if param_index >= len(params):
            return None
        
        param = params[param_index]
        if not isinstance(param, c_ast.Decl):
            return None
        
        # Extract type information
        param_type = self.generator.visit(param.type)
        param_name = param.name if param.name else f"param{param_index}"
        
        return param_type, param_name
    
    def create_type_from_string(self, type_string: str, var_name: str = "dummy_var") -> c_ast.Node:
        """Create a type AST node from a string representation."""
        # Create a dummy declaration to parse the type
        dummy_code = f"{type_string} {var_name};"
        try:
            dummy_ast = self.parser.parse(dummy_code)
            decl = dummy_ast.ext[0]
            return decl.type
        except ParseError as e:
            print(f"Warning: Failed to parse type '{type_string}': {e}")
            print("Falling back to simple identifier type")
            # Fallback to simple identifier type
            return c_ast.IdentifierType(names=[type_string.split()[-1]])
    
    def modify_function_signature(self, func_node: c_ast.FuncDef, param_index: int, 
                                new_type: str, original_name: str) -> bool:
        """Modify the function signature by changing parameter type and name."""
        if not func_node.decl.type.args:
            return False
        
        params = func_node.decl.type.args.params
        if param_index >= len(params):
            return False
        
        param = params[param_index]
        if not isinstance(param, c_ast.Decl):
            return False
        
        # Create new type from string
        new_type_node = self.create_type_from_string(new_type, f"{original_name}_")
        
        # Update parameter type and name
        param.type = new_type_node
        param.name = f"{original_name}_"
        
        return True
    
    def modify_function_declaration(self, func_decl: c_ast.Decl, param_index: int, 
                                  new_type: str, original_name: str) -> bool:
        """Modify the function declaration by changing parameter type and name."""
        if not func_decl.type.args:
            return False
        
        params = func_decl.type.args.params
        if param_index >= len(params):
            return False
        
        param = params[param_index]
        if not isinstance(param, c_ast.Decl):
            return False
        
        # Create new type from string
        new_type_node = self.create_type_from_string(new_type, f"{original_name}_")
        
        # Update parameter type and name
        param.type = new_type_node
        param.name = f"{original_name}_"
        
        return True
    
    def modify_function_body(self, func_node: c_ast.FuncDef, original_type: str, 
                           original_name: str, modified_param_name: str) -> bool:
        """Add variable declaration and assignment to function body."""
        if not func_node.body:
            return False
        
        # Create variable declaration: original_type original_name;
        var_decl_type = self.create_type_from_string(original_type, original_name)
        var_decl = c_ast.Decl(
            name=original_name,
            quals=[],
            align=[],
            storage=[],
            funcspec=[],
            type=var_decl_type,
            init=None,
            bitsize=None
        )
        
        # Create assignment: original_name = modified_param_name;
        assignment = c_ast.Assignment(
            op='=',
            lvalue=c_ast.ID(name=original_name),
            rvalue=c_ast.ID(name=modified_param_name)
        )
        
        # Insert at the beginning of function body
        if not func_node.body.block_items:
            func_node.body.block_items = []
        
        # Insert declaration and assignment at the beginning
        func_node.body.block_items.insert(0, assignment)
        func_node.body.block_items.insert(0, var_decl)
        
        return True
    
    def process_file(self, filepath: str, function_name: str, param_index: int, 
                   new_type: str, output_path: str = None, dry_run: bool = False) -> bool:
        """Main processing function that modifies the C file."""
        # Parse the file
        ast = self.parse_file(filepath)
        
        # Look for function definition first
        func_def = self.find_function_definition(ast, function_name)
        func_decl = None
        
        if func_def:
            # Get original parameter info
            param_info = self.get_parameter_info(func_def, param_index)
            if not param_info:
                print(f"Error: Parameter at index {param_index} not found in function {function_name}")
                return False
            
            original_type, original_name = param_info
            modified_param_name = f"{original_name}_"
            
            print(f"Found function definition: {function_name}")
            print(f"Original parameter: {original_type} {original_name}")
            print(f"Modified parameter: {new_type} {modified_param_name}")
            
            # Modify function signature (do this even in dry run to show the result)
            if not self.modify_function_signature(func_def, param_index, new_type, original_name):
                print("Error: Failed to modify function signature")
                return False
            
            # Modify function body
            if not self.modify_function_body(func_def, original_type, original_name, modified_param_name):
                print("Error: Failed to modify function body")
                return False
        
        # Also look for function declarations (in headers)
        func_decl = self.find_function_declaration(ast, function_name)
        if func_decl and func_decl != (func_def.decl if func_def else None):
            param_info = self.get_parameter_info_from_decl(func_decl, param_index)
            if param_info:
                original_type, original_name = param_info
                print(f"Found function declaration: {function_name}")
                print(f"Original parameter: {original_type} {original_name}")
                
                if not self.modify_function_declaration(func_decl, param_index, new_type, original_name):
                    print("Error: Failed to modify function declaration")
                    return False
        
        if not func_def and not func_decl:
            print(f"Error: Function '{function_name}' not found in {filepath}")
            return False
        
        if dry_run:
            print("Dry run - no changes written to file")
            print("Generated code:")
            print(self.generator.visit(ast))
            return True
        
        # Write output
        output_file = output_path or filepath
        try:
            with open(output_file, 'w') as f:
                f.write(self.generator.visit(ast))
            print(f"Successfully modified {output_file}")
            return True
        except Exception as e:
            print(f"Error writing to {output_file}: {e}")
            return False


def main():
    parser = argparse.ArgumentParser(
        description="Modify C function parameter types and names",
        formatter_class=argparse.RawDescriptionHelpFormatter,
        epilog="""
Examples:
  %(prog)s myFunction 1 "int*" src/myfile.c
    - Modify the 2nd parameter (index 1) of myFunction to type int*
  
  %(prog)s processData 0 "const char*" include/header.h
    - Modify the 1st parameter (index 0) of processData to type const char*
"""
    )
    
    parser.add_argument("function_name", help="Name of the C function to modify")
    parser.add_argument("param_index", type=int, help="0-based index of the parameter to modify")
    parser.add_argument("new_type", help="New type for the parameter (e.g., 'int*', 'const char*')")
    parser.add_argument("file_path", help="Path to the C/header file")
    parser.add_argument("--output", "-o", help="Output file path (default: overwrite input file)")
    parser.add_argument("--dry-run", action="store_true", help="Show changes without writing to file")
    
    args = parser.parse_args()
    
    if not os.path.exists(args.file_path):
        print(f"Error: File '{args.file_path}' does not exist")
        sys.exit(1)
    
    if args.param_index < 0:
        print("Error: Parameter index must be non-negative")
        sys.exit(1)
    
    if not args.function_name.strip():
        print("Error: Function name cannot be empty")
        sys.exit(1)
    
    if not args.new_type.strip():
        print("Error: New type cannot be empty")
        sys.exit(1)
    
    # Validate that the new type string looks reasonable
    invalid_chars = ['@', '#', '$', '%', '^', '&', '!', '~', '`']
    if any(char in args.new_type for char in invalid_chars):
        print(f"Warning: New type '{args.new_type}' contains potentially invalid characters")
    
    # Check file extension
    if not args.file_path.lower().endswith(('.c', '.h', '.cpp', '.hpp', '.cc', '.cxx')):
        print(f"Warning: File '{args.file_path}' doesn't appear to be a C/C++ source or header file")
    
    modifier = CFunctionModifier()
    
    print(f"Processing file: {args.file_path}")
    print(f"Looking for function: {args.function_name}")
    print(f"Parameter index: {args.param_index}")
    print(f"New type: {args.new_type}")
    
    # Process the file
    success = modifier.process_file(
        args.file_path, 
        args.function_name, 
        args.param_index, 
        args.new_type,
        args.output,
        args.dry_run
    )
    
    if success:
        print("Operation completed successfully")
        sys.exit(0)
    else:
        print("Operation failed")
        sys.exit(1)


if __name__ == "__main__":
    main()