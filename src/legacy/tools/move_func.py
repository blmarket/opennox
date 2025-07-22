#!/usr/bin/env python3
import sys
import os
import glob
import tree_sitter
import tree_sitter_c as tsc

def find_function_definition(parser, source_code, function_name, file_path=None):
    try:
        tree = parser.parse(source_code.encode())
        if tree.root_node.has_error:
            if file_path:
                print(f"Warning: Parse tree has errors in {file_path}")
            else:
                print(f"Warning: Parse tree has errors")

        def traverse(node):
            if node.type == 'function_definition':
                declarator = node.child_by_field_name('declarator')
                if declarator:
                    func_name_node = get_function_name_from_declarator(declarator)
                    if func_name_node:
                        found_name = func_name_node.text.decode()
                        if found_name == function_name:
                            return node

            for child in node.children:
                result = traverse(child)
                if result:
                    return result
            return None

        return traverse(tree.root_node)
    except Exception as e:
        if file_path:
            print(f"Error parsing source code in {file_path}: {e}")
        else:
            print(f"Error parsing source code: {e}")
        return None

def get_function_name_from_declarator(declarator):
    if declarator.type == 'function_declarator':
        return declarator.child_by_field_name('declarator')
    elif declarator.type == 'identifier':
        return declarator
    elif declarator.type == 'pointer_declarator':
        return get_function_name_from_declarator(declarator.child_by_field_name('declarator'))
    return None

def extract_function_text(source_code, node):
    start_byte = node.start_byte
    end_byte = node.end_byte
    return source_code[start_byte:end_byte]

def remove_function_from_source(source_code, node):
    start_byte = node.start_byte
    end_byte = node.end_byte

    # Find the start of the line containing the function
    line_start = start_byte
    while line_start > 0 and source_code[line_start - 1] != '\n':
        line_start -= 1

    # Find the end of the line after the function
    line_end = end_byte
    while line_end < len(source_code) and source_code[line_end] != '\n':
        line_end += 1
    if line_end < len(source_code):
        line_end += 1  # Include the newline

    return source_code[:line_start] + source_code[line_end:]

def find_c_files():
    return glob.glob('**/*.c', recursive=True)

def main():
    if len(sys.argv) != 3:
        print("Usage: python move_func.py <function_name> <destination_file>")
        sys.exit(1)

    function_name = sys.argv[1]
    dest_file = sys.argv[2]

    # Initialize tree-sitter parser
    language = tree_sitter.Language(tsc.language(), 'c')
    parser = tree_sitter.Parser()
    parser.set_language(language)

    # Find all C files
    c_files = find_c_files()

    found_function = None
    source_file_path = None

    # Search for the function in all C files
    for file_path in c_files:
        if not os.path.exists(file_path):
            continue

        try:
            with open(file_path, 'r', encoding='utf-8', errors='ignore') as f:
                source_code = f.read()

            if len(source_code) == 0:
                print(f"Skipping empty file: {file_path}")
                continue

            function_node = find_function_definition(parser, source_code, function_name, file_path)
            if function_node:
                found_function = function_node
                source_file_path = file_path
                print(f"Found function '{function_name}' in {file_path}")
                break
        except Exception as e:
            print(f"Error reading {file_path}: {e}")
            continue

    if not found_function:
        print(f"Function '{function_name}' not found in any C file")
        sys.exit(1)

    # Extract function text
    with open(source_file_path, 'r', encoding='utf-8', errors='ignore') as f:
        source_code = f.read()

    function_text = extract_function_text(source_code, found_function)

    # Remove function from source file
    updated_source = remove_function_from_source(source_code, found_function)

    # Write updated source file
    with open(source_file_path, 'w', encoding='utf-8') as f:
        f.write(updated_source)

    # Append function to destination file
    with open(dest_file, 'a', encoding='utf-8') as f:
        if os.path.getsize(dest_file) > 0:
            f.write('\n')
        f.write(function_text)
        f.write('\n')

    print(f"Function '{function_name}' moved from {source_file_path} to {dest_file}")

if __name__ == "__main__":
    main()
