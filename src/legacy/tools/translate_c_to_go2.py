#!/usr/bin/env python3

import sys
import argparse
import subprocess
import re
import os
from typing import Set

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

def make_function_public(go_content: str) -> str:
    """Make all function names public by capitalizing the first letter."""
    def capitalize_func_name(match):
        receiver = match.group(1)  # The receiver part: "(m *AudioModule) "
        func_name = match.group(2)
        params = match.group(3)
        
        # Capitalize first letter of function name
        if func_name and func_name[0].islower():
            func_name = func_name[0].upper() + func_name[1:]
        
        return f'func {receiver}{func_name}({params})'
    
    # Pattern to match function definitions with receivers
    func_pattern = r'func\s+(\(m \*\w+\)\s+)(\w+)\s*\(([^)]*)\)'
    
    return re.sub(func_pattern, capitalize_func_name, go_content)

def extract_functions_only(go_content: str) -> str:
    """Extract only function definitions from Go content, excluding package, imports, and struct definitions."""
    lines = go_content.split('\n')
    function_lines = []
    in_function = False
    brace_count = 0
    
    for line in lines:
        stripped = line.strip()
        
        # Skip package and import statements
        if stripped.startswith('package ') or stripped.startswith('import ') or stripped == 'import (':
            continue
        
        # Skip single-line imports and closing import parenthesis
        if re.match(r'^\s*"[^"]*"\s*$', stripped) or stripped == ')':
            continue
        
        # Skip struct definitions and type declarations
        if stripped.startswith('type ') and ('struct' in stripped or 'interface' in stripped):
            # Skip the entire struct/interface definition
            if '{' in stripped:
                struct_brace_count = stripped.count('{') - stripped.count('}')
                if struct_brace_count > 0:
                    # Multi-line struct, skip until closing brace
                    while struct_brace_count > 0:
                        try:
                            next_line = next(iter(lines[lines.index(line)+1:]))
                            struct_brace_count += next_line.count('{') - next_line.count('}')
                        except (StopIteration, ValueError):
                            break
            continue
        
        # Check if this line starts a function
        if stripped.startswith('func '):
            in_function = True
            brace_count = stripped.count('{') - stripped.count('}')
            function_lines.append(line)
            continue
        
        # If we're in a function, add the line and track braces
        if in_function:
            function_lines.append(line)
            brace_count += stripped.count('{') - stripped.count('}')
            
            # If braces are balanced, we've finished this function
            if brace_count == 0:
                in_function = False
                function_lines.append('')  # Add empty line between functions
    
    return '\n'.join(function_lines)

def append_to_audio_impl(new_functions: str, module_name: str = "AudioModule") -> str:
    """Append new functions to the existing audio/audio_impl.go file."""
    audio_impl_path = "audio/audio_impl.go"
    
    if not os.path.exists(audio_impl_path):
        raise FileNotFoundError(f"Target file {audio_impl_path} not found")
    
    # Read existing content
    with open(audio_impl_path, 'r') as f:
        existing_content = f.read()
    
    # Append new functions
    updated_content = existing_content.rstrip() + '\n\n' + new_functions.strip() + '\n'
    
    return updated_content

def main():
    parser = argparse.ArgumentParser(description='Translate C file and append to existing AudioModule')
    parser.add_argument('c_file', help='Path to the C file to translate')
    parser.add_argument('--module-name', default='AudioModule', help='Name of the module struct (default: AudioModule)')

    args = parser.parse_args()

    if not os.path.exists(args.c_file):
        print(f"Error: C file {args.c_file} not found", file=sys.stderr)
        sys.exit(1)

    module_name = args.module_name
    output_file = "audio/audio_impl.go"

    # Check if target file exists
    if not os.path.exists(output_file):
        print(f"Error: Target file {output_file} not found", file=sys.stderr)
        sys.exit(1)

    try:
        print(f"Translating {args.c_file} to Go...")

        # Step 1: Generate Go code from C file
        print("Running cxgo")
        run_cxgo2()

        # Determine the corresponding Go file based on C filename
        c_base_name = os.path.splitext(os.path.basename(args.c_file))[0]
        go_file_path = f"../../gonox/{c_base_name}.go"

        with open(go_file_path, 'r') as f:
            go_content = f.read()

        # Step 2: Get external variables
        print("Getting external variables...")
        external_vars = get_external_variables(args.c_file)

        # Step 3: Add module receivers to functions
        print("Adding module receivers...")
        go_content = add_module_receivers(go_content, module_name)

        # Step 4: Prefix external variable references
        print("Prefixing external variables...")
        go_content = prefix_external_variables(go_content, external_vars)

        print(go_content)

        # Step 5: Make functions public
        print("Making functions public...")
        go_content = make_function_public(go_content)

        print(go_content)

        # Step 6: Extract only functions
        print("Extracting functions...")
        functions_only = extract_functions_only(go_content)

        # Step 7: Append to existing audio_impl.go
        print("Appending to audio_impl.go...")
        updated_content = append_to_audio_impl(functions_only, module_name)

        # Write updated content back to file
        with open(output_file, 'w') as f:
            f.write(updated_content)

        print(f"Successfully appended new functions to {output_file}")

    except Exception as e:
        print(f"Error during translation: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == '__main__':
    main()