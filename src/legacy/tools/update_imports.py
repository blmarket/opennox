#!/usr/bin/env python3

import sys
import os
import argparse
import re
import subprocess
from typing import Dict, List, Tuple, Optional


def find_declaration(target_identifier: str) -> Optional[str]:
    """Find the declaration for a given identifier using find_decl.py tool."""
    try:
        result = subprocess.run(
            ['python', '-m', 'tools.find_decl', target_identifier],
            capture_output=True, text=True, check=True
        )
        return result.stdout.strip()
    except subprocess.CalledProcessError:
        return None
    except FileNotFoundError:
        print("Error: find_decl tool not found", file=sys.stderr)
        sys.exit(1)


def parse_function_declaration(decl: str) -> Optional[Tuple[str, str, List[Tuple[str, str]]]]:
    """Parse C function declaration and return (return_type, func_name, parameters)."""
    # Remove semicolon and strip
    decl = decl.rstrip(';').strip()
    
    # Basic regex to match function declaration
    # return_type func_name(param1, param2, ...)
    match = re.match(r'^([^()]+?)\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\(([^)]*)\)$', decl)
    if not match:
        return None

    return_type = match.group(1).strip()
    func_name = match.group(2).strip()
    params_str = match.group(3).strip()

    # Parse parameters
    parameters = []
    if params_str and params_str != 'void':
        # Split by comma, but be careful with nested parentheses and function pointers
        param_parts = []
        paren_count = 0
        current_param = ""

        for char in params_str:
            if char == ',' and paren_count == 0:
                param_parts.append(current_param.strip())
                current_param = ""
            else:
                if char == '(':
                    paren_count += 1
                elif char == ')':
                    paren_count -= 1
                current_param += char

        if current_param.strip():
            param_parts.append(current_param.strip())

        for param in param_parts:
            param = param.strip()
            if param:
                # Extract parameter type and name
                param_match = re.match(r'^(.+?)\s+([a-zA-Z_][a-zA-Z0-9_]*)$', param)
                if param_match:
                    param_type = param_match.group(1).strip()
                    param_name = param_match.group(2).strip()
                else:
                    # Parameter without name (just type)
                    param_type = param
                    param_name = f"arg{len(parameters)}"

                parameters.append((param_type, param_name))

    return return_type, func_name, parameters


def update_cgo_comment_block(imports_file: str, func_declaration: str, func_name: str) -> bool:
    """Add function declaration to CGO comment block in imports file if not already present."""
    try:
        with open(imports_file, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Check if function already exists in CGO block
        if func_name in content:
            print(f"Function {func_name} already exists in CGO block, skipping...")
            return True
        
        # Find the CGO comment block
        cgo_start = content.find('/*')
        cgo_end = content.find('*/', cgo_start) + 2
        
        if cgo_start == -1 or cgo_end == -1:
            print(f"Error: Could not find CGO comment block in {imports_file}", file=sys.stderr)
            return False
        
        # Extract the existing CGO block content
        before_cgo = content[:cgo_start]
        cgo_block = content[cgo_start:cgo_end]
        after_cgo = content[cgo_end:]
        
        # Add the new function declaration before the closing */
        new_func_line = func_declaration + (';\n' if not func_declaration.endswith(';') else '\n')
        updated_cgo_block = cgo_block[:-2] + new_func_line + '*/'
        
        # Write the updated content back
        updated_content = before_cgo + updated_cgo_block + after_cgo
        
        with open(imports_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        
        return True
        
    except Exception as e:
        print(f"Error updating CGO comment block: {e}", file=sys.stderr)
        return False


def add_field_to_module_struct(module_file: str, func_name: str, func_signature: str) -> bool:
    """Add field to the module struct in module/module.go file if not already present."""
    try:
        with open(module_file, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Check if field already exists
        if func_name in content:
            print(f"Field {func_name} already exists in module struct, skipping...")
            return True
        
        # Find the struct definition
        struct_pattern = r'(type\s+\w+Module\s+struct\s*\{[^}]+)(})'
        match = re.search(struct_pattern, content, re.DOTALL)
        
        if not match:
            print(f"Error: Could not find module struct in {module_file}", file=sys.stderr)
            return False
        
        struct_content = match.group(1)
        struct_end = match.group(2)
        
        # Add new field before the closing brace
        new_field = f"\t{func_name} {func_signature}\n"
        updated_struct = struct_content + new_field + struct_end
        
        # Replace in content
        updated_content = content.replace(match.group(0), updated_struct)
        
        with open(module_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        
        return True
        
    except Exception as e:
        print(f"Error updating module struct: {e}", file=sys.stderr)
        return False


def update_module_constructor(module_file: str, func_name: str, func_signature: str) -> bool:
    """Update NewXxxModule function to add the new field if not already present."""
    try:
        with open(module_file, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Check if constructor already has this parameter/field
        constructor_has_param = f"{func_name} {func_signature}" in content
        constructor_has_assignment = f"{func_name}: {func_name}," in content
        
        if constructor_has_param and constructor_has_assignment:
            print(f"Constructor already has {func_name} parameter and assignment, skipping...")
            return True
        
        # Find the constructor function
        constructor_pattern = r'(func\s+New\w+Module\s*\([^}]+\)\s+\*\w+Module\s*\{[^}]+return\s+&\w+Module\s*\{[^}]+)(})'
        match = re.search(constructor_pattern, content, re.DOTALL)
        
        if not match:
            print(f"Error: Could not find constructor function in {module_file}", file=sys.stderr)
            return False
        
        # Add parameter to function signature
        func_content = match.group(1)
        func_end = match.group(2)
        
        # Add new field assignment before the closing brace if not already present
        if not constructor_has_assignment:
            new_assignment = f"\t\t{func_name}: {func_name},\n"
            func_content = func_content + new_assignment
        
        updated_func = func_content + "\t" + func_end
        
        # Also need to add parameter to function signature if not already present
        if not constructor_has_param:
            # Find the parameter list
            param_pattern = r'(func\s+New\w+Module\s*\([^)]+)(,?\s*\)\s+\*\w+Module)'
            param_match = re.search(param_pattern, updated_func)
            
            if param_match:
                param_list = param_match.group(1)
                param_end = param_match.group(2)
                
                # Add new parameter
                new_param = f",\n\t{func_name} {func_signature}"
                updated_param = param_list + new_param + param_end
                updated_func = updated_func.replace(param_match.group(0), updated_param)
        
        # Replace in content
        updated_content = content.replace(match.group(0), updated_func)
        
        with open(module_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        
        return True
        
    except Exception as e:
        print(f"Error updating module constructor: {e}", file=sys.stderr)
        return False


def map_c_type_to_go(c_type: str) -> str:
    """Map C type to Go type."""
    # Remove const, static, extern keywords
    c_type = re.sub(r'\b(const|static|extern)\s+', '', c_type).strip()
    
    # Type mappings from C to Go
    type_mappings = {
        'uint32_t': 'uint32',
        'int32_t': 'int32',
        'uint16_t': 'uint16',
        'int16_t': 'int16',
        'uint8_t': 'uint8',
        'int8_t': 'int8',
        'char': 'byte',
        'int': 'int',
        'unsigned int': 'uint',
        'short': 'int16',
        'unsigned short': 'uint16',
        'long': 'int32',
        'unsigned long': 'uint32',
        'float': 'float32',
        'double': 'float64',
        'void': '',
    }
    
    # Handle pointer types
    if '*' in c_type:
        return 'unsafe.Pointer'
    
    # Handle basic types
    for c_typ, go_typ in type_mappings.items():
        if c_type == c_typ:
            return go_typ
    
    # If type not found, use unsafe.Pointer as fallback
    return 'unsafe.Pointer'


def generate_go_function_signature(return_type: str, parameters: List[Tuple[str, str]]) -> str:
    """Generate Go function signature from C function info."""
    go_params = []
    
    for param_type, param_name in parameters:
        go_type = map_c_type_to_go(param_type)
        go_params.append(go_type)
    
    go_return_type = map_c_type_to_go(return_type)
    
    params_str = ", ".join(go_params)
    
    if go_return_type:
        return f"func({params_str}) {go_return_type}"
    else:
        return f"func({params_str})"


def generate_inline_function_wrapper(func_name: str, return_type: str, parameters: List[Tuple[str, str]]) -> str:
    """Generate inline function wrapper for initXxx function."""
    go_params = []
    cgo_params = []
    
    for i, (param_type, param_name) in enumerate(parameters):
        go_type = map_c_type_to_go(param_type)
        param_var = f"a{i+1}"
        go_params.append(f"{param_var} {go_type}")
        
        # For CGO call, cast parameters appropriately
        if go_type == 'unsafe.Pointer':
            cgo_params.append(f"({param_type})({param_var})")
        else:
            cgo_params.append(f"C.{param_type}({param_var})")
    
    go_return_type = map_c_type_to_go(return_type)
    params_str = ", ".join(go_params)
    cgo_call = f"C.{func_name}({', '.join(cgo_params)})"
    
    lines = []
    if go_return_type:
        return_annotation = f" {go_return_type}"
        if go_return_type == 'unsafe.Pointer':
            body = f"return unsafe.Pointer({cgo_call})"
        else:
            body = f"return {go_return_type}({cgo_call})"
    else:
        return_annotation = ""
        body = cgo_call
    
    wrapper = f"func({params_str}){return_annotation} {{\n\t\t\t{body}\n\t\t}}"
    return wrapper


def update_init_function(imports_file: str, module_name: str, func_name: str, inline_wrapper: str) -> bool:
    """Update initXxx function to include the new function if not already present."""
    try:
        with open(imports_file, 'r', encoding='utf-8') as f:
            content = f.read()
        
        # Check if the inline wrapper already exists
        if inline_wrapper.replace('\n\t\t\t', ' ').replace('\n\t\t', ' ') in content.replace('\n', ' '):
            print(f"Inline wrapper for {func_name} already exists in init function, skipping...")
            return True
        
        # Find the init function - improved pattern
        init_func_name = f"init{module_name.capitalize()}"
        # More flexible pattern that handles the actual structure
        init_pattern = rf'(func\s+{init_func_name}\s*\(\)\s*\{{.*?{module_name}\.New{module_name.capitalize()}Module\s*\([^)]*?)(\s*\)\s*\}})'
        match = re.search(init_pattern, content, re.DOTALL)
        
        if not match:
            print(f"Error: Could not find init function {init_func_name} in {imports_file}", file=sys.stderr)
            return False
        
        func_params = match.group(1)
        func_end = match.group(2)
        
        # Add new function parameter
        new_param = f",\n\t\t{inline_wrapper}"
        updated_func = func_params + new_param + func_end
        
        # Replace in content
        updated_content = content.replace(match.group(0), updated_func)
        
        with open(imports_file, 'w', encoding='utf-8') as f:
            f.write(updated_content)
        
        return True
        
    except Exception as e:
        print(f"Error updating init function: {e}", file=sys.stderr)
        return False


def main():
    parser = argparse.ArgumentParser(
        description='Update Go imports file to add a new dependency'
    )
    parser.add_argument('module_name', help='Module name (e.g., audio)')
    parser.add_argument('function_name', help='Function name to add (e.g., sub_4862E0)')
    
    args = parser.parse_args()
    
    module_name = args.module_name
    func_name = args.function_name
    
    # File paths
    imports_file = f"{module_name}_imports.go"
    module_file = f"{module_name}/{module_name}.go"
    
    # Check if files exist
    if not os.path.exists(imports_file):
        print(f"Error: {imports_file} not found", file=sys.stderr)
        sys.exit(1)
    
    if not os.path.exists(module_file):
        print(f"Error: {module_file} not found", file=sys.stderr)
        sys.exit(1)
    
    # Find the function declaration
    print(f"Finding declaration for {func_name}...")
    func_declaration = find_declaration(func_name)
    
    if not func_declaration:
        print(f"Error: Declaration for '{func_name}' not found", file=sys.stderr)
        sys.exit(1)
    
    print(f"Found declaration: {func_declaration}")
    
    # Parse the function declaration
    parsed = parse_function_declaration(func_declaration)
    if not parsed:
        print(f"Error: Could not parse function declaration: {func_declaration}", file=sys.stderr)
        sys.exit(1)
    
    return_type, parsed_func_name, parameters = parsed
    
    # Generate Go function signature
    go_func_signature = generate_go_function_signature(return_type, parameters)
    print(f"Generated Go signature: {go_func_signature}")
    
    # Generate inline function wrapper
    inline_wrapper = generate_inline_function_wrapper(func_name, return_type, parameters)
    print(f"Generated inline wrapper: {inline_wrapper}")
    
    # Update files
    print("Updating CGO comment block...")
    if not update_cgo_comment_block(imports_file, func_declaration, func_name):
        sys.exit(1)
    
    print("Adding field to module struct...")
    if not add_field_to_module_struct(module_file, func_name, go_func_signature):
        sys.exit(1)
    
    print("Updating module constructor...")
    if not update_module_constructor(module_file, func_name, go_func_signature):
        sys.exit(1)
    
    print("Updating init function...")
    if not update_init_function(imports_file, module_name, func_name, inline_wrapper):
        sys.exit(1)
    
    print(f"Successfully updated {module_name} module to include {func_name}")


if __name__ == '__main__':
    main()