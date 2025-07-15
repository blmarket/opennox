#!/usr/bin/env python3
"""
C Function Parameter Type Updater

Updates C function signatures by changing parameter types and adding compatibility
variable declarations in the function body.

Usage:
    python update.py <function_name> <param_index> <new_type>

Example:
    python update.py func1 0 uint32_t
"""

import sys
import os
import argparse
from pathlib import Path
import tree_sitter
from tree_sitter import Language, Parser

# Try to import tree-sitter-c
try:
    import tree_sitter_c
except ImportError:
    print("Error: tree-sitter-c not found. Install with: pip install tree-sitter-c")
    sys.exit(1)


class CFunctionUpdater:
    def __init__(self):
        # Initialize tree-sitter C parser
        self.c_language = Language(tree_sitter_c.language(), 'c')
        self.parser = Parser()
        self.parser.set_language(self.c_language)
    
    def find_c_files(self, directory="."):
        """Find all C and header files in the given directory."""
        path = Path(directory)
        c_files = list(path.glob("**/*.c")) + list(path.glob("**/*.h"))
        return c_files
    
    def parse_file(self, file_path):
        """Parse a C file and return the AST."""
        with open(file_path, 'rb') as f:
            content = f.read()
        
        tree = self.parser.parse(content)
        return tree, content
    
    def find_function_nodes(self, tree, function_name):
        """Find all function declaration and definition nodes with the given name."""
        matches = []
        
        def visit_node(node):
            # Check for function declarations and definitions
            if node.type in ['function_definition', 'declaration']:
                if node.type == 'function_definition':
                    # For function definitions, get the declarator
                    declarator = node.child_by_field_name('declarator')
                    if declarator and declarator.type == 'function_declarator':
                        func_name_node = declarator.child_by_field_name('declarator')
                        if func_name_node and func_name_node.text.decode() == function_name:
                            matches.append(node)
                
                elif node.type == 'declaration':
                    # For declarations, look for init_declarator or function_declarator
                    declarator = node.child_by_field_name('declarator')
                    if not declarator:
                        # Check if it's a direct function declarator
                        for child in node.children:
                            if child.type == 'function_declarator':
                                func_name_node = child.child_by_field_name('declarator')
                                if func_name_node and func_name_node.text.decode() == function_name:
                                    matches.append(node)
                    else:
                        if declarator.type == 'function_declarator':
                            func_name_node = declarator.child_by_field_name('declarator')
                            if func_name_node and func_name_node.text.decode() == function_name:
                                matches.append(node)
                        elif declarator.type == 'init_declarator':
                            func_declarator = declarator.child_by_field_name('declarator')
                            if func_declarator and func_declarator.type == 'function_declarator':
                                func_name_node = func_declarator.child_by_field_name('declarator')
                                if func_name_node and func_name_node.text.decode() == function_name:
                                    matches.append(node)
            
            # Recursively visit children
            for child in node.children:
                visit_node(child)
        
        visit_node(tree.root_node)
        return matches
    
    def get_parameter_info(self, function_node, param_index):
        """Extract parameter information from a function node."""
        # Find the function declarator
        declarator = None
        if function_node.type == 'function_definition':
            declarator = function_node.child_by_field_name('declarator')
        elif function_node.type == 'declaration':
            for child in function_node.children:
                if child.type == 'function_declarator':
                    declarator = child
                    break
            if not declarator:
                declarator = function_node.child_by_field_name('declarator')
                if declarator and declarator.type == 'init_declarator':
                    declarator = declarator.child_by_field_name('declarator')
        
        if not declarator or declarator.type != 'function_declarator':
            return None
        
        # Get parameters
        parameters_node = declarator.child_by_field_name('parameters')
        if not parameters_node:
            return None
        
        params = []
        for child in parameters_node.children:
            if child.type == 'parameter_declaration':
                params.append(child)
        
        if param_index >= len(params):
            return None
        
        target_param = params[param_index]
        
        # Extract type and name
        param_type = target_param.child_by_field_name('type')
        param_declarator = target_param.child_by_field_name('declarator')
        
        if not param_type or not param_declarator:
            return None
        
        old_type = param_type.text.decode().strip()
        param_name = param_declarator.text.decode().strip()
        
        return {
            'old_type': old_type,
            'param_name': param_name,
            'param_node': target_param,
            'type_node': param_type,
            'declarator_node': param_declarator
        }
    
    def update_function_signature(self, content, function_node, param_index, new_type):
        """Update function signature and add variable declaration if it's a definition."""
        param_info = self.get_parameter_info(function_node, param_index)
        if not param_info:
            return content
        
        old_type = param_info['old_type']
        param_name = param_info['param_name']
        new_param_name = param_name + '_'
        
        # Update the parameter type and name in the signature
        type_start = param_info['type_node'].start_byte
        type_end = param_info['type_node'].end_byte
        declarator_start = param_info['declarator_node'].start_byte
        declarator_end = param_info['declarator_node'].end_byte
        
        # Replace type
        content = content[:type_start] + new_type.encode() + content[type_end:]
        
        # Adjust positions after type replacement
        type_len_diff = len(new_type.encode()) - (type_end - type_start)
        declarator_start += type_len_diff
        declarator_end += type_len_diff
        
        # Replace parameter name
        content = content[:declarator_start] + new_param_name.encode() + content[declarator_end:]
        
        # If this is a function definition, add variable declaration in the body
        if function_node.type == 'function_definition':
            # Find the compound statement (function body)
            body = function_node.child_by_field_name('body')
            if body and body.type == 'compound_statement':
                # Find the opening brace and add the declaration after it
                body_start = body.start_byte + type_len_diff + len(new_param_name.encode()) - len(param_name.encode())
                
                # Find the position after the opening brace
                brace_pos = body_start
                while brace_pos < len(content) and content[brace_pos:brace_pos+1] != b'{':
                    brace_pos += 1
                brace_pos += 1  # Move past the opening brace
                
                # Skip any whitespace after the opening brace
                while brace_pos < len(content) and content[brace_pos:brace_pos+1] in b' \t\n\r':
                    brace_pos += 1
                
                # Add the variable declaration
                var_declaration = f"  {old_type} {param_name} = {new_param_name};\n"
                content = content[:brace_pos] + var_declaration.encode() + content[brace_pos:]
        
        return content
    
    def update_function_in_file(self, file_path, function_name, param_index, new_type):
        """Update all instances of the function in a single file."""
        tree, content = self.parse_file(file_path)
        function_nodes = self.find_function_nodes(tree, function_name)
        
        if not function_nodes:
            return False
        
        # Process nodes in reverse order to maintain byte positions
        function_nodes.reverse()
        
        for node in function_nodes:
            content = self.update_function_signature(content, node, param_index, new_type)
        
        # Write back to file
        with open(file_path, 'wb') as f:
            f.write(content)
        
        return True
    
    def update_function_in_codebase(self, function_name, param_index, new_type, directory="."):
        """Update all instances of the function in the entire codebase."""
        c_files = self.find_c_files(directory)
        updated_files = []
        
        for file_path in c_files:
            if self.update_function_in_file(file_path, function_name, param_index, new_type):
                updated_files.append(file_path)
        
        return updated_files


def main():
    parser = argparse.ArgumentParser(description='Update C function parameter types')
    parser.add_argument('function_name', help='Name of the function to update')
    parser.add_argument('param_index', type=int, help='Index of parameter to update (0-based)')
    parser.add_argument('new_type', help='New type for the parameter')
    parser.add_argument('--directory', '-d', default='.', help='Directory to search for C files')
    
    args = parser.parse_args()
    
    updater = CFunctionUpdater()
    updated_files = updater.update_function_in_codebase(
        args.function_name, 
        args.param_index, 
        args.new_type,
        args.directory
    )
    
    if updated_files:
        print(f"Updated function '{args.function_name}' in {len(updated_files)} files:")
        for file_path in updated_files:
            print(f"  - {file_path}")
    else:
        print(f"Function '{args.function_name}' not found in any C files.")


if __name__ == '__main__':
    main()
