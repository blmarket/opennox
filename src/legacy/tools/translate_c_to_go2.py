#!/usr/bin/env python3

import sys
import argparse
import subprocess
import re
import os
import shutil
from typing import Dict, List, Set
from pathlib import Path

def run_cxgo2():
    try:
        cmd = ['go', 'run', 'github.com/gotranspile/cxgo/cmd/cxgo@latest']
        subprocess.run(cmd, cwd="../../")
    except subprocess.CalledProcessError as e:
        print(f"Error running cxgo: {e.stderr}", file=sys.stderr)
        sys.exit(1)
    except FileNotFoundError as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

def get_external_variables(c_file_path: str) -> Set[str]:
    """Get list of external variables from the external_funcs tool."""
    try:
        result = subprocess.run(
            ['python', '-m', 'tools.external_funcs', c_file_path],
            capture_output=True, text=True, check=True
        )

        # Parse the output to extract variable names
        external_vars = set()
        lines = result.stdout.strip().split('\n')
        for line in lines:
            line = line.strip()
            if line.startswith('extern ') and not '(' in line:  # Variables, not functions
                # Extract variable name from extern declaration
                # e.g., "extern uint32_t someVar;" -> "someVar"
                var_match = re.search(r'extern\s+\w+(?:\*+)?\s+(\w+);', line)
                if var_match:
                    external_vars.add(var_match.group(1))

        return external_vars
    except subprocess.CalledProcessError as e:
        print(f"Error running external_funcs tool: {e.stderr}", file=sys.stderr)
        return set()

def add_module_receivers(go_content: str, module_name: str) -> str:
    """Add module receiver to all function definitions."""
    # Pattern to match function definitions
    func_pattern = r'func\s+(\w+)\s*\('

    def replace_func(match):
        func_name = match.group(1)
        return f'func (m *{module_name}) {func_name}('

    return re.sub(func_pattern, replace_func, go_content)

def prefix_external_variables(go_content: str, external_vars: Set[str]) -> str:
    """Prefix external variable references with *m."""
    if not external_vars:
        return go_content

    # Create pattern for all external variables
    var_pattern = r'\b(' + '|'.join(re.escape(var) for var in external_vars) + r')\b'

    def replace_var(match):
        var_name = match.group(1)
        # Don't replace if it's already prefixed or in a declaration context
        before = go_content[:match.start()]

        # Skip if already has m. prefix
        if before.endswith('m.') or before.endswith('*m.'):
            return match.group(0)

        # Skip if it's in a type declaration or function parameter
        if re.search(r'\b(type|var|func)\s+[^{]*$', before.split('\n')[-1]):
            return match.group(0)

        return f'*m.{var_name}'

    return re.sub(var_pattern, replace_var, go_content)

def extract_functions_from_go(go_content: str) -> List[str]:
    """Extract individual function definitions from Go content."""
    functions = []
    lines = go_content.split('\n')
    i = 0
    
    while i < len(lines):
        line = lines[i].strip()
        # Look for function definitions with receivers
        func_match = re.match(r'func\s+\([^)]+\)\s+(\w+)\s*\(', line)
        if func_match:
            # Found function start, now collect the entire function
            func_lines = [lines[i]]
            i += 1
            brace_count = 0
            in_function = False
            
            while i < len(lines):
                current_line = lines[i]
                func_lines.append(current_line)
                
                # Count braces to find function end
                for char in current_line:
                    if char == '{':
                        brace_count += 1
                        in_function = True
                    elif char == '}':
                        brace_count -= 1
                        if in_function and brace_count == 0:
                            # Function complete
                            functions.append('\n'.join(func_lines))
                            break
                
                if in_function and brace_count == 0:
                    break
                i += 1
        else:
            i += 1
    
    return functions

def extract_function_signature(func_def: str) -> tuple:
    """Extract function name, parameters, and return type from function definition."""
    # Match function signature with receiver
    func_match = re.search(r'func\s+\([^)]+\)\s+(\w+)\s*\(([^)]*)\)(?:\s+([^{]+))?', func_def)
    if not func_match:
        return None, None, None
        
    func_name = func_match.group(1)
    params = func_match.group(2).strip() if func_match.group(2) else ""
    return_type = func_match.group(3).strip() if func_match.group(3) else ""
    
    return func_name, params, return_type

def go_type_to_c_type(go_type: str) -> str:
    """Convert Go types to C types for CGO declarations."""
    type_mapping = {
        'int32': 'int32_t',
        'uint32': 'uint32_t',
        'int64': 'int64_t',
        'uint64': 'uint64_t',
        'int16': 'int16_t',
        'uint16': 'uint16_t',
        'int8': 'int8_t',
        'uint8': 'uint8_t',
        'bool': 'bool',
        'int': 'int',
        'float32': 'float',
        'float64': 'double',
        'unsafe.Pointer': 'void*',
        '': 'void'
    }
    
    # Handle pointer types
    if go_type.startswith('*'):
        base_type = go_type[1:]
        c_base = type_mapping.get(base_type, base_type)
        return f"{c_base}*"
    
    return type_mapping.get(go_type, go_type)

def generate_cgo_declarations(functions: List[str]) -> List[str]:
    """Generate CGO function declarations and //export comments for the functions."""
    cgo_decls = []
    
    for func in functions:
        func_name, params, return_type = extract_function_signature(func)
        if not func_name:
            continue
            
        # Skip if function name starts with lowercase (unexported)
        if func_name[0].islower():
            continue
            
        # Convert Go types to C types for CGO
        c_params = []
        if params:
            param_list = [p.strip() for p in params.split(',')]
            for param in param_list:
                if ' ' in param:
                    param_parts = param.split()
                    param_name = param_parts[0]
                    param_type = ' '.join(param_parts[1:])
                    c_type = go_type_to_c_type(param_type)
                    c_params.append(f"{c_type} {param_name}")
        
        c_return_type = go_type_to_c_type(return_type) if return_type else "void"
        
        # Generate CGO declaration
        param_str = ", ".join(c_params) if c_params else "void"
        cgo_decl = f"//export {func_name}\n{c_return_type} {func_name}({param_str});"
        cgo_decls.append(cgo_decl)
    
    return cgo_decls

def add_export_comments(functions: List[str]) -> List[str]:
    """Add //export comments to function definitions."""
    exported_functions = []
    
    for func in functions:
        func_name, _, _ = extract_function_signature(func)
        if func_name and func_name[0].isupper():  # Only export public functions
            # Add //export comment before function
            export_comment = f"//export {func_name}\n"
            exported_func = export_comment + func
            exported_functions.append(exported_func)
        else:
            exported_functions.append(func)
    
    return exported_functions

def append_to_audio_impl(c_file_path: str, module_name: str = "AudioModule") -> tuple:
    """Extract functions from translated C file and prepare them for appending to audio_impl.go."""
    print(f"Translating {c_file_path} to Go functions...")

    # Step 1: Generate Go code from C file
    print("Running cxgo")
    run_cxgo2()

    # Determine the corresponding Go file based on C filename
    c_base_name = os.path.splitext(os.path.basename(c_file_path))[0]
    go_file_path = f"../../gonox/{c_base_name}.go"

    with open(go_file_path, 'r') as f:
        go_content = f.read()

    # Step 2: Get external variables
    print("Getting external variables...")
    external_vars = get_external_variables(c_file_path)

    # Step 3: Add module receivers to functions
    print("Adding module receivers...")
    go_content = add_module_receivers(go_content, module_name)

    # Step 4: Prefix external variable references
    print("Prefixing external variables...")
    go_content = prefix_external_variables(go_content, external_vars)

    # Step 5: Extract individual function definitions
    print("Extracting function definitions...")
    functions = extract_functions_from_go(go_content)

    # Step 6: Add //export comments to functions
    print("Adding //export comments...")
    exported_functions = add_export_comments(functions)

    # Step 7: Generate CGO declarations
    print("Generating CGO declarations...")
    cgo_declarations = generate_cgo_declarations(functions)

    return exported_functions, cgo_declarations

def append_functions_to_file(functions: List[str], cgo_declarations: List[str], target_file: str):
    """Append the new functions and CGO declarations to the existing audio_impl.go file."""
    
    # Read existing content
    with open(target_file, 'r') as f:
        existing_content = f.read()
    
    # Add CGO declarations at the top after imports if any functions need exporting
    if cgo_declarations:
        # Find the end of imports
        import_end = existing_content.find('\n)', existing_content.find('import ('))
        if import_end == -1:
            # Fallback: add after package declaration
            package_end = existing_content.find('\n', existing_content.find('package '))
            insert_pos = package_end + 1
        else:
            insert_pos = import_end + 2
        
        # Check if CGO is already imported
        if 'import "C"' not in existing_content:
            # Insert CGO declarations and import
            cgo_section = "\n/*\n" + "\n".join(cgo_declarations) + "\n*/\nimport \"C\"\n\n"
            existing_content = existing_content[:insert_pos] + cgo_section + existing_content[insert_pos:]
        else:
            # Find existing CGO section and append to it
            cgo_start = existing_content.find('/*')
            if cgo_start != -1:
                cgo_end = existing_content.find('*/', cgo_start) 
                if cgo_end != -1:
                    # Insert new declarations before the closing */
                    existing_content = (existing_content[:cgo_end] + 
                                      "\n" + "\n".join(cgo_declarations) + "\n" +
                                      existing_content[cgo_end:])
    
    # Add a comment separator and the new functions
    functions_content = "\n\n".join(functions)
    separator = "\n\n// === Functions translated from C file ===\n\n"
    updated_content = existing_content + separator + functions_content
    
    # Write back to file
    with open(target_file, 'w') as f:
        f.write(updated_content)

def main():
    parser = argparse.ArgumentParser(description='Translate C file functions and append to existing AudioModule')
    parser.add_argument('c_file', help='Path to the C file to translate')
    parser.add_argument('--module-name', default='AudioModule', help='Name of the module struct (default: AudioModule)')

    args = parser.parse_args()

    if not os.path.exists(args.c_file):
        print(f"Error: C file {args.c_file} not found", file=sys.stderr)
        sys.exit(1)

    # Target file is always audio/audio_impl.go
    target_file = "audio/audio_impl.go"
    
    if not os.path.exists(target_file):
        print(f"Error: Target file {target_file} not found", file=sys.stderr)
        sys.exit(1)

    # Generate the functions to append
    try:
        functions, cgo_declarations = append_to_audio_impl(args.c_file, args.module_name)
        
        if not functions:
            print("Warning: No functions found to append")
            return
            
        # Append to existing file
        append_functions_to_file(functions, cgo_declarations, target_file)
        
        print(f"Successfully appended {len(functions)} functions to {target_file}")
        if cgo_declarations:
            print(f"Added {len(cgo_declarations)} CGO declarations for C interop")

    except Exception as e:
        print(f"Error during translation: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == '__main__':
    main()