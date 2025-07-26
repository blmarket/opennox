#!/usr/bin/env python3

import sys
import argparse
import tree_sitter_c as tsc
from tree_sitter import Language, Parser
from typing import Set, List


def get_source_text(node, source_bytes):
    """Extract the source text for a given node."""
    return source_bytes[node.start_byte:node.end_byte].decode('utf-8')


def get_function_name(node, source_bytes):
    """Extract function name from a function declarator node."""
    def find_identifier(node):
        if node.type == 'identifier':
            return get_source_text(node, source_bytes)
        for child in node.children:
            result = find_identifier(child)
            if result:
                return result
        return None
    
    return find_identifier(node)


def find_function_declarator(node):
    """Find function_declarator node within a declaration or definition."""
    if node.type == 'function_declarator':
        return node
    for child in node.children:
        result = find_function_declarator(child)
        if result:
            return result
    return None


def extract_extern_variables(tree, source_bytes):
    """Extract all extern variable declarations."""
    extern_vars = []
    
    def traverse(node):
        if node.type == 'declaration':
            # Check if it starts with extern and is not a function
            decl_text = get_source_text(node, source_bytes).strip()
            if decl_text.startswith('extern') and '(' not in decl_text:
                # This is an extern variable declaration
                extern_vars.append(decl_text)
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return extern_vars


def extract_function_declarations(tree, source_bytes):
    """Extract all function declarations (not definitions)."""
    declarations = set()
    
    def traverse(node):
        if node.type == 'declaration':
            # Look for function declarator in this declaration
            func_declarator = find_function_declarator(node)
            if func_declarator:
                # This is a function declaration, get its text
                decl_text = get_source_text(node, source_bytes).strip()
                # Ensure it ends with semicolon (declaration, not definition)
                if decl_text.endswith(';'):
                    # Extract function name for deduplication
                    func_name = get_function_name(func_declarator, source_bytes)
                    if func_name:
                        declarations.add((func_name, decl_text))
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return declarations


def extract_function_definitions(tree, source_bytes):
    """Extract all function definitions."""
    definitions = set()
    
    def traverse(node):
        if node.type == 'function_definition':
            # Find the function declarator to get the name
            func_declarator = find_function_declarator(node)
            if func_declarator:
                func_name = get_function_name(func_declarator, source_bytes)
                if func_name:
                    definitions.add(func_name)
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return definitions


def find_external_dependencies(c_file_path):
    """Find all external dependencies in a C file."""
    try:
        with open(c_file_path, 'r', encoding='utf-8') as f:
            source_code = f.read()
    except FileNotFoundError:
        print(f"Error: File '{c_file_path}' not found", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error reading file: {e}", file=sys.stderr)
        sys.exit(1)
    
    # Initialize treesitter
    C_LANGUAGE = Language(tsc.language(), "c")
    parser = Parser()
    parser.set_language(C_LANGUAGE)
    
    source_bytes = source_code.encode('utf-8')
    tree = parser.parse(source_bytes)
    
    # Extract extern variables
    extern_vars = extract_extern_variables(tree, source_bytes)
    
    # Extract function declarations and definitions
    function_declarations = extract_function_declarations(tree, source_bytes)
    function_definitions = extract_function_definitions(tree, source_bytes)
    
    # Find external function declarations (declarations - definitions)
    external_functions = []
    for func_name, decl_text in function_declarations:
        if func_name not in function_definitions:
            external_functions.append(decl_text)
    
    return extern_vars, external_functions


def format_for_cgo(extern_vars, external_functions):
    """Format the output for CGO import section."""
    output_lines = []
    
    # Add standard includes that are commonly needed  
    output_lines.append('#include "defs.h"')
    output_lines.append("")
    
    # Add extern variables
    for var in extern_vars:
        output_lines.append(var)
    
    # Add external function declarations
    for func in external_functions:
        # Ensure function declarations end with semicolon
        if not func.endswith(';'):
            func += ';'
        output_lines.append(func)
    
    return '\n'.join(output_lines)


def main():
    parser = argparse.ArgumentParser(
        description='Extract external dependencies from C code for CGO imports'
    )
    parser.add_argument('filename', help='C source file to analyze')
    
    args = parser.parse_args()
    
    extern_vars, external_functions = find_external_dependencies(args.filename)
    
    output = format_for_cgo(extern_vars, external_functions)
    print(output)


if __name__ == '__main__':
    main()