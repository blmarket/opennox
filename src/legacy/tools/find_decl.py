#!/usr/bin/env python3

import sys
import os
import glob
import argparse
from pathlib import Path
import tree_sitter_c as tsc
from tree_sitter import Language, Parser

def find_c_files(directory="."):
    """Find all .c and .h files in the given directory and subdirectories."""
    c_files = []
    for pattern in ["**/*.c", "**/*.h"]:
        c_files.extend(glob.glob(os.path.join(directory, pattern), recursive=True))
    return c_files

def get_source_text(node, source_bytes):
    """Extract the source text for a given node."""
    return source_bytes[node.start_byte:node.end_byte].decode('utf-8')

def find_function_declarations(tree, source_bytes):
    """Find all function declarations in the parse tree."""
    declarations = []
    
    def find_function_declarator(node):
        """Recursively find function_declarator nodes."""
        if node.type == 'function_declarator':
            return node
        for child in node.children:
            result = find_function_declarator(child)
            if result:
                return result
        return None
    
    def traverse(node):
        if node.type == 'declaration':
            # Look for function declarator in this declaration
            func_declarator = find_function_declarator(node)
            if func_declarator:
                # This is a function declaration
                decl_text = get_source_text(node, source_bytes).strip()
                # Extract function name
                identifier_node = None
                for child in func_declarator.children:
                    if child.type == 'identifier':
                        identifier_node = child
                        break
                
                if identifier_node:
                    func_name = get_source_text(identifier_node, source_bytes)
                    declarations.append((func_name, decl_text))
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return declarations

def find_extern_variable_declarations(tree, source_bytes):
    """Find all extern variable declarations in the parse tree."""
    declarations = []
    
    def find_identifier_in_declarator(node):
        """Recursively find identifier in declarator."""
        if node.type == 'identifier':
            return node
        for child in node.children:
            result = find_identifier_in_declarator(child)
            if result:
                return result
        return None
    
    def traverse(node):
        if node.type == 'declaration':
            # Check if this is an extern declaration
            has_extern = False
            variable_name = None
            
            for child in node.children:
                if child.type == 'storage_class_specifier' and get_source_text(child, source_bytes) == 'extern':
                    has_extern = True
                elif child.type in ['pointer_declarator', 'init_declarator', 'identifier']:
                    # Find the variable name in the declarator
                    identifier_node = find_identifier_in_declarator(child)
                    if identifier_node:
                        variable_name = get_source_text(identifier_node, source_bytes)
            
            if has_extern and variable_name:
                decl_text = get_source_text(node, source_bytes).strip()
                declarations.append((variable_name, decl_text))
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return declarations

def parse_file(file_path, parser):
    """Parse a single C file and extract declarations."""
    try:
        with open(file_path, 'rb') as f:
            source_bytes = f.read()
        
        tree = parser.parse(source_bytes)
        
        # Find function declarations
        func_declarations = find_function_declarations(tree, source_bytes)
        
        # Find extern variable declarations
        var_declarations = find_extern_variable_declarations(tree, source_bytes)
        
        return func_declarations, var_declarations
    
    except Exception as e:
        print(f"Error parsing {file_path}: {e}", file=sys.stderr)
        return [], []

def find_declaration(target_identifier, search_dir="."):
    """Find the declaration for a given identifier."""
    # Set up tree-sitter
    C_LANGUAGE = Language(tsc.language(), "c")
    parser = Parser()
    parser.set_language(C_LANGUAGE)
    
    # Find all C files
    c_files = find_c_files(search_dir)
    
    all_func_declarations = []
    all_var_declarations = []
    
    # Parse each file
    for file_path in c_files:
        func_decls, var_decls = parse_file(file_path, parser)
        all_func_declarations.extend(func_decls)
        all_var_declarations.extend(var_decls)
    
    # Search for the target identifier
    # First check function declarations
    for func_name, declaration in all_func_declarations:
        if func_name == target_identifier:
            return declaration
    
    # Then check extern variable declarations
    for var_name, declaration in all_var_declarations:
        if var_name == target_identifier:
            return declaration
    
    return None

def main():
    parser = argparse.ArgumentParser(
        description="Find proper declaration of global variable or function using tree-sitter"
    )
    parser.add_argument("identifier", help="Target identifier (function or variable name)")
    parser.add_argument("-d", "--directory", default=".", help="Directory to search (default: current directory)")
    
    args = parser.parse_args()
    
    declaration = find_declaration(args.identifier, args.directory)
    
    if declaration:
        print(declaration)
    else:
        print(f"Declaration for '{args.identifier}' not found", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()