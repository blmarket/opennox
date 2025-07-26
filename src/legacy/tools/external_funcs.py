#!/usr/bin/env python3

import sys
import argparse
from pathlib import Path
import tree_sitter_c as tsc
from tree_sitter import Language, Parser


def get_source_text(node, source_bytes):
    """Extract the source text for a given node."""
    return source_bytes[node.start_byte:node.end_byte].decode('utf-8')


def find_function_name_from_declarator(func_declarator, source_bytes):
    """Extract function name from a function_declarator node."""
    for child in func_declarator.children:
        if child.type == 'identifier':
            return get_source_text(child, source_bytes)
    return None


def find_function_declarations(tree, source_bytes):
    """Find all function declarations in the parse tree."""
    declarations = {}
    
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
                func_name = find_function_name_from_declarator(func_declarator, source_bytes)
                if func_name:
                    decl_text = get_source_text(node, source_bytes).strip()
                    declarations[func_name] = decl_text
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return declarations


def find_function_definitions(tree, source_bytes):
    """Find all function definitions in the parse tree."""
    definitions = {}
    
    def find_function_declarator(node):
        """Recursively find function_declarator nodes."""
        if node.type == 'function_declarator':
            return node
        # Handle pointer return types
        if node.type == 'pointer_declarator':
            for child in node.children:
                if child.type == 'function_declarator':
                    return child
        for child in node.children:
            result = find_function_declarator(child)
            if result:
                return result
        return None
    
    def traverse(node):
        if node.type == 'function_definition':
            # Look for function declarator in this definition
            func_declarator = find_function_declarator(node)
            if func_declarator:
                func_name = find_function_name_from_declarator(func_declarator, source_bytes)
                if func_name:
                    # Extract just the signature part (without the body)
                    for child in node.children:
                        if child.type == 'compound_statement':
                            # Everything before the compound statement is the signature
                            signature_end = child.start_byte
                            signature_text = source_bytes[:signature_end].decode('utf-8')
                            # Find the start of this function definition
                            signature_start = node.start_byte
                            func_signature = source_bytes[signature_start:signature_end].decode('utf-8').strip()
                            definitions[func_name] = func_signature + ";"
                            break
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return definitions


def find_external_functions(file_path):
    """Find external function dependencies in a C file."""
    # Set up tree-sitter
    C_LANGUAGE = Language(tsc.language(), "c")
    parser = Parser()
    parser.set_language(C_LANGUAGE)
    
    try:
        with open(file_path, 'rb') as f:
            source_bytes = f.read()
        
        tree = parser.parse(source_bytes)
        
        # Find all function declarations
        declarations = find_function_declarations(tree, source_bytes)
        
        # Find all function definitions
        definitions = find_function_definitions(tree, source_bytes)
        
        # External functions are declared but not defined
        external_functions = {}
        for func_name, decl_text in declarations.items():
            if func_name not in definitions:
                external_functions[func_name] = decl_text
        
        return external_functions, declarations, definitions
        
    except Exception as e:
        print(f"Error parsing {file_path}: {e}", file=sys.stderr)
        return {}, {}, {}


def main():
    parser = argparse.ArgumentParser(
        description="Find external function dependencies in a C file using tree-sitter"
    )
    parser.add_argument("file_path", help="C source file to analyze")
    parser.add_argument("-v", "--verbose", action="store_true", 
                       help="Show all declarations and definitions")
    parser.add_argument("--declarations-only", action="store_true",
                       help="Show only declarations")
    parser.add_argument("--definitions-only", action="store_true", 
                       help="Show only definitions")
    
    args = parser.parse_args()
    
    if not Path(args.file_path).exists():
        print(f"Error: File '{args.file_path}' not found", file=sys.stderr)
        sys.exit(1)
    
    external_functions, declarations, definitions = find_external_functions(args.file_path)
    
    if args.declarations_only:
        for func_name, decl_text in sorted(declarations.items()):
            print(f"{func_name}: {decl_text}")
    elif args.definitions_only:
        for func_name, def_text in sorted(definitions.items()):
            print(f"{func_name}: {def_text}")
    else:
        # Default: show external functions
        if external_functions:
            for func_name, decl_text in sorted(external_functions.items()):
                print(f"{func_name}: {decl_text}")
        
        if args.verbose:
            print(f"\n=== Summary ===")
            print(f"Total declarations: {len(declarations)}")
            print(f"Total definitions: {len(definitions)}")
            print(f"External dependencies: {len(external_functions)}")
            
            if declarations:
                print(f"\n=== All Declarations ===")
                for func_name, decl_text in sorted(declarations.items()):
                    print(f"{func_name}: {decl_text}")
            
            if definitions:
                print(f"\n=== All Definitions ===")
                for func_name, def_text in sorted(definitions.items()):
                    print(f"{func_name}: {def_text}")


if __name__ == "__main__":
    main()