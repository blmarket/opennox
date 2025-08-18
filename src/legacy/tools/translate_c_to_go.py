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


def extract_struct_definitions(defs_content: str, target_structs: List[str]) -> str:
    """Extract specific struct definitions from defs.go content."""
    extracted_structs = []

    for struct_name in target_structs:
        # Look for struct definition pattern
        pattern = rf'type\s+{re.escape(struct_name)}\s+struct\s*\{{[^}}]*\}}'
        match = re.search(pattern, defs_content, re.DOTALL)
        if match:
            extracted_structs.append(match.group(0))
        else:
            print(f"Warning: struct {struct_name} not found in defs.go", file=sys.stderr)

    return '\n\n'.join(extracted_structs)

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

def infer_module_name_from_filename(c_file_path: str) -> str:
    """Infer module name from C filename: 'audio.c' => 'AudioModule'."""
    base_name = os.path.splitext(os.path.basename(c_file_path))[0]
    # Capitalize first letter and add 'Module' suffix
    return base_name.capitalize() + 'Module'

def create_audio_impl_go(c_file_path: str, module_name: str = "AudioModule") -> str:
    """Create the complete audio_impl.go content."""
    print(f"Translating {c_file_path} to Go...")

    # Step 1: Generate Go code from C file
    print("Running cxgo")
    run_cxgo2()

    # Determine the corresponding Go file based on C filename
    c_base_name = os.path.splitext(os.path.basename(c_file_path))[0]
    go_file_path = f"../../gonox/{c_base_name}.go"

    with open(go_file_path, 'r') as f:
        go_content = f.read()
    with open("../../gonox/defs.go", "r") as f:
        defs_go_content = f.read()
        # target_structs = ['struct200', 'struct576', 'nox_list_item_t']
        target_structs = [] # No hardcoded struct definitions for now.
        struct_definitions = extract_struct_definitions(defs_go_content, target_structs)

    # Step 4: Get external variables
    print("Getting external variables...")
    external_vars = get_external_variables(c_file_path)

    # Step 5: Add module receivers to functions
    print("Adding module receivers...")
    go_content = add_module_receivers(go_content, module_name)

    # Step 6: Prefix external variable references
    print("Prefixing external variables...")
    go_content = prefix_external_variables(go_content, external_vars)

    # Step 7: Combine everything
    final_content = []

    # Add package declaration and imports (extract from original audio.go)
    package_section = []
    lines = go_content.split('\n')
    in_imports = False

    for line in lines:
        if line.startswith('package '):
            # package_section.append(f'package {c_file_path.replace(".c", "")}')
            package_section.append(f'package audio')
        elif line.startswith('import '):
            in_imports = True
            package_section.append(line)
        elif in_imports and (line.startswith('\t') or line.strip() in ['(', ')']):
            package_section.append(line)
        elif in_imports and line.strip() == '':
            package_section.append(line)
            break
        elif not in_imports and line.strip() == '':
            package_section.append(line)

    final_content.extend(package_section)

    # Add struct definitions if available
    if struct_definitions:
        final_content.append('\n// Extracted struct definitions from defs.h')
        final_content.append(struct_definitions)
        final_content.append('')

    # Add the main audio implementation
    # Remove package and import sections from audio_go_content
    impl_lines = []
    skip_until_empty = False
    can_skip = True
    for line in lines:
        if line.startswith('import '):
            can_skip = False
            skip_until_empty = True
        if not can_skip and line.strip() == '':
            skip_until_empty = False
            continue
        if can_skip or skip_until_empty:
            continue
        impl_lines.append(line)

    final_content.extend(impl_lines)

    return '\n'.join(final_content)

def main():
    parser = argparse.ArgumentParser(description='Translate C file to Go implementation')
    parser.add_argument('c_file', help='Path to the C file to translate')
    parser.add_argument('--module-name', help='Name of the module struct (default: inferred from filename)')

    args = parser.parse_args()

    if not os.path.exists(args.c_file):
        print(f"Error: C file {args.c_file} not found", file=sys.stderr)
        sys.exit(1)

    # Infer module name from filename if not provided
    module_name = args.module_name or infer_module_name_from_filename(args.c_file)

    # Get base name for output directory
    base_name = os.path.splitext(os.path.basename(args.c_file))[0]
    output_dir = "audio"
    output_file = os.path.join(output_dir, f'{base_name}_impl.go')

    # Create output directory
    os.makedirs(output_dir, exist_ok=True)

    # Generate the Go implementation
    try:
        go_content = create_audio_impl_go(args.c_file, module_name)

        # Write to output file
        with open(output_file, 'w') as f:
            f.write(go_content)

        print(f"Successfully created {output_file}")

    except Exception as e:
        print(f"Error during translation: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == '__main__':
    main()
