#!/usr/bin/env python3
"""
C Function Parameter Update Tool

Updates C function signatures to change parameter types while maintaining
backward compatibility by creating local variables with original types.
"""

import sys
import os
import argparse
from pathlib import Path
import tree_sitter
from tree_sitter import Language, Parser


def setup_tree_sitter():
    """Set up tree-sitter C parser"""
    try:
        # Try to import tree_sitter_c
        import tree_sitter_c as tsc
        C_LANGUAGE = Language(tsc.language(), "c")
    except ImportError:
        # Fallback to manual language setup if needed
        raise ImportError("Please install tree-sitter-c: pip install tree-sitter-c")
    
    parser = Parser()
    parser.set_language(C_LANGUAGE)
    return parser


def find_c_files(directory):
    """Find all C and header files in directory"""
    c_files = []
    for ext in ['*.c', '*.h']:
        c_files.extend(Path(directory).glob(f'**/{ext}'))
    return c_files


def find_function_nodes(tree, function_name):
    """Find all function declaration and definition nodes with given name"""
    function_nodes = []
    
    def traverse(node):
        if node.type == 'function_definition':
            declarator = node.child_by_field_name('declarator')
            if declarator and extract_function_name(declarator) == function_name:
                function_nodes.append(('definition', node))
        elif node.type == 'declaration':
            # Check if this is a function declaration
            declarator = node.child_by_field_name('declarator')
            if declarator and extract_function_name(declarator) == function_name:
                function_nodes.append(('declaration', node))
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return function_nodes


def extract_function_name(declarator_node):
    """Extract function name from declarator node"""
    # Handle direct function declarator
    if declarator_node.type == 'function_declarator':
        declarator = declarator_node.child_by_field_name('declarator')
        if declarator and declarator.type == 'identifier':
            return declarator.text.decode()
    
    # Handle pointer declarator wrapping function declarator
    if declarator_node.type == 'pointer_declarator':
        child = declarator_node.children[-1]  # Last child is the actual declarator
        return extract_function_name(child)
    
    # Handle direct identifier
    if declarator_node.type == 'identifier':
        return declarator_node.text.decode()
    
    return None


def get_parameter_list(declarator_node):
    """Get parameter list node from function declarator"""
    if declarator_node.type == 'function_declarator':
        return declarator_node.child_by_field_name('parameters')
    
    # Handle pointer declarator
    if declarator_node.type == 'pointer_declarator':
        child = declarator_node.children[-1]
        return get_parameter_list(child)
    
    return None


def update_function_signature(source_code, function_node, param_index, new_type):
    """Update function signature by changing parameter type and name, or return type if param_index is -1"""
    node_type, node = function_node
    
    # Handle return type update when param_index is -1
    if param_index == -1:
        return update_return_type(source_code, function_node, new_type)
    
    # Get the declarator
    declarator = node.child_by_field_name('declarator')
    if not declarator:
        return source_code, None, None
    
    # Get parameter list
    param_list = get_parameter_list(declarator)
    if not param_list:
        return source_code, None, None
    
    # Get parameters (excluding parentheses and commas)
    parameters = [child for child in param_list.children 
                 if child.type == 'parameter_declaration']
    
    if param_index >= len(parameters):
        print(f"Warning: Parameter index {param_index} out of range for function with {len(parameters)} parameters")
        return source_code, None, None
    
    param_node = parameters[param_index]
    
    # Extract current parameter type and name
    param_text = param_node.text.decode()
    
    # Find the declarator within the parameter
    param_declarator = param_node.child_by_field_name('declarator')
    if not param_declarator:
        # Handle the case where there's no explicit declarator (just type)
        param_name = f"param{param_index}"
        original_type = param_text.strip()
    else:
        param_name = extract_parameter_name(param_declarator)
        original_type = extract_parameter_type(param_node, param_declarator)
    
    # Create new parameter name with suffix
    new_param_name = param_name + "_"
    
    # Create new parameter text
    new_param_text = f"{new_type} {new_param_name}"
    
    # Replace the parameter in source code
    start_byte = param_node.start_byte
    end_byte = param_node.end_byte
    
    new_source = (source_code[:start_byte] + 
                 new_param_text.encode() + 
                 source_code[end_byte:])
    
    return new_source, original_type, param_name


def update_return_type(source_code, function_node, new_type):
    """Update function return type"""
    node_type, node = function_node
    
    # Find the type specifier(s) before the declarator
    type_nodes = []
    declarator = node.child_by_field_name('declarator')
    
    for child in node.children:
        if child == declarator:
            break
        if child.type in ['type_qualifier', 'storage_class_specifier', 'primitive_type', 
                         'type_identifier', 'struct_specifier', 'union_specifier']:
            type_nodes.append(child)
    
    if not type_nodes:
        return source_code, None, None
    
    # Get the original return type including any pointer indicators from the declarator
    original_type = ' '.join(child.text.decode() for child in type_nodes)
    
    # Check if the declarator is a pointer_declarator and include the pointer part in original type
    if declarator and declarator.type == 'pointer_declarator':
        # Count the pointer levels
        pointer_count = 0
        node_ptr = declarator
        while node_ptr and node_ptr.type == 'pointer_declarator':
            pointer_count += 1
            # Find the next declarator
            for child in node_ptr.children:
                if child.type in ['function_declarator', 'identifier', 'pointer_declarator']:
                    node_ptr = child
                    break
            else:
                break
        original_type += '*' * pointer_count
        
        # For pointer declarators, we need to replace more than just the type specifiers
        # We need to replace up to the start of the function declarator
        end_byte = declarator.start_byte
        # Find the actual function declarator within the pointer declarator
        func_declarator = None
        node_search = declarator
        while node_search:
            for child in node_search.children:
                if child.type == 'function_declarator':
                    func_declarator = child
                    break
                elif child.type == 'pointer_declarator':
                    node_search = child
                    break
            else:
                break
            if func_declarator:
                break
        if func_declarator:
            end_byte = func_declarator.start_byte
    else:
        # For non-pointer return types, just replace the type specifiers
        end_byte = type_nodes[-1].end_byte
    
    # Calculate the range to replace
    start_byte = type_nodes[0].start_byte
    
    # Add a space after the new type if it doesn't end with one
    if not new_type.endswith(' '):
        new_type += ' '
    
    # Replace the return type
    new_source = (source_code[:start_byte] + 
                 new_type.encode() + 
                 source_code[end_byte:])
    
    return new_source, original_type, None


def extract_parameter_name(declarator_node):
    """Extract parameter name from parameter declarator"""
    if declarator_node.type == 'identifier':
        return declarator_node.text.decode()
    elif declarator_node.type == 'pointer_declarator':
        # Handle pointer parameters
        for child in declarator_node.children:
            if child.type == 'identifier':
                return child.text.decode()
    return "param"


def extract_parameter_type(param_node, declarator_node):
    """Extract the original parameter type"""
    # Get the type specifier
    type_parts = []
    for child in param_node.children:
        if child.type in ['type_qualifier', 'storage_class_specifier', 'primitive_type', 
                         'type_identifier', 'struct_specifier', 'union_specifier']:
            type_parts.append(child.text.decode())
        elif child == declarator_node:
            break
    
    # Handle pointer declarators
    if declarator_node.type == 'pointer_declarator':
        pointer_count = 0
        node = declarator_node
        while node.type == 'pointer_declarator':
            pointer_count += 1
            node = node.children[-1] if node.children else None
            if not node:
                break
        type_parts.append('*' * pointer_count)
    
    return ' '.join(type_parts)


def add_variable_declaration(source_code, function_node, original_type, param_name, new_param_name):
    """Add variable declaration at the beginning of function body"""
    node_type, node = function_node
    
    if node_type != 'definition':
        return source_code  # Only add to function definitions, not declarations
    
    # Find the compound statement (function body)
    body = node.child_by_field_name('body')
    if not body or body.type != 'compound_statement':
        return source_code
    
    # Find the opening brace
    opening_brace = None
    for child in body.children:
        if child.type == '{':
            opening_brace = child
            break
    
    if not opening_brace:
        return source_code
    
    # Create the variable declaration
    var_declaration = f"\n\t{original_type} {param_name} = {new_param_name};"
    
    # Insert after the opening brace
    insert_pos = opening_brace.end_byte
    
    new_source = (source_code[:insert_pos] + 
                 var_declaration.encode() + 
                 source_code[insert_pos:])
    
    return new_source


def process_file(file_path, function_name, param_index, new_type):
    """Process a single C file"""
    parser = setup_tree_sitter()
    
    with open(file_path, 'rb') as f:
        source_code = f.read()
    
    tree = parser.parse(source_code)
    function_nodes = find_function_nodes(tree, function_name)
    
    if not function_nodes:
        return False  # No functions found
    
    modified = False
    current_source = source_code
    
    # Process function occurrences from last to first to preserve byte offsets
    function_nodes_reverse = list(reversed(function_nodes))
    
    for function_node in function_nodes_reverse:
        # Update function signature
        new_source, original_type, param_name = update_function_signature(
            current_source, function_node, param_index, new_type)
        
        if new_source != current_source and original_type:
            current_source = new_source
            modified = True
            
            # For function definitions and parameter changes (not return type), add variable declaration
            if function_node[0] == 'definition' and param_index != -1:
                # Re-parse the modified source
                tree = parser.parse(current_source)
                updated_function_nodes = find_function_nodes(tree, function_name)
                
                # Find the corresponding function definition
                for updated_node in updated_function_nodes:
                    if updated_node[0] == 'definition':
                        new_param_name = param_name + "_"
                        current_source = add_variable_declaration(
                            current_source, updated_node, original_type, param_name, new_param_name)
                        break
    
    if modified:
        with open(file_path, 'wb') as f:
            f.write(current_source)
        print(f"Updated {file_path}")
    
    return modified


def main():
    if len(sys.argv) != 4:
        print("Usage: python update.py <function_name> <param_index> <new_type>")
        print("Example: python update.py sub_451920 0 'struct200*'")
        print("Use -1 for param_index to update return type:")
        print("Example: python update.py sub_451920 -1 'struct264*'")
        sys.exit(1)
    
    function_name = sys.argv[1]
    param_index = int(sys.argv[2])
    new_type = sys.argv[3]
    
    # Find all C files in current directory
    c_files = find_c_files('.')
    
    updated_count = 0
    for file_path in c_files:
        try:
            if process_file(file_path, function_name, param_index, new_type):
                updated_count += 1
        except Exception as e:
            print(f"Error processing {file_path}: {e}")
    
    print(f"Updated {updated_count} files")


if __name__ == "__main__":
    main()