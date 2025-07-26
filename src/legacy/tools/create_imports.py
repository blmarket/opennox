#!/usr/bin/env python3

import sys
import argparse
import subprocess
import re
import os
from typing import Dict, List, Tuple, Optional


# Type mappings from C to Go
TYPE_MAPPINGS = {
    'uint32_t': 'uint32',
    'int32_t': 'int32',
    'uint16_t': 'uint16',
    'int16_t': 'int16',
    'uint8_t': 'uint8',
    'int8_t': 'int8',
    'char': 'byte',
    'int': 'int32',
    'unsigned int': 'uint32',
    'short': 'int16',
    'unsigned short': 'uint16',
    'long': 'int32',
    'unsigned long': 'uint32',
    'float': 'float32',
    'double': 'float64',
    'void': '',
    'void*': 'unsafe.Pointer',
    'nox_list_item_t*': '*C.nox_list_item_t',
    'uint32_t**': '**C.uint32_t',
    'uint32_t*': '*C.uint32_t',
    'char*': '*C.char',
}


def get_external_dependencies(c_file_path: str) -> str:
    """Get external dependencies using the external_funcs tool."""
    try:
        result = subprocess.run(
            ['python', '-m', 'tools.external_funcs', c_file_path],
            capture_output=True, text=True, check=True
        )
        return result.stdout.strip()
    except subprocess.CalledProcessError as e:
        print(f"Error running external_funcs tool: {e.stderr}", file=sys.stderr)
        sys.exit(1)
    except FileNotFoundError:
        print("Error: external_funcs tool not found", file=sys.stderr)
        sys.exit(1)


def parse_function_declaration(decl: str) -> Optional[Tuple[str, str, List[Tuple[str, str]], str]]:
    """Parse C function declaration and return (return_type, func_name, parameters, original_decl)."""
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
    
    return return_type, func_name, parameters, decl


def map_c_type_to_go(c_type: str) -> str:
    """Map C type to Go type."""
    # Remove const, static, extern keywords
    c_type = re.sub(r'\b(const|static|extern)\s+', '', c_type).strip()
    
    # Handle pointer types
    if '*' in c_type:
        # For any pointer type, use unsafe.Pointer
        return 'unsafe.Pointer'
    
    # Handle basic types
    for c_typ, go_typ in TYPE_MAPPINGS.items():
        if c_type == c_typ:
            return go_typ
    
    # If type not found, raise error
    raise ValueError(f"Unknown C type: {c_type}")


def map_c_type_to_cgo(c_type: str) -> str:
    """Map C type to CGO type (with C. prefix)."""
    # Remove const, static, extern keywords
    c_type = re.sub(r'\b(const|static|extern)\s+', '', c_type).strip()

    if c_type in TYPE_MAPPINGS:
        return TYPE_MAPPINGS[c_type]

    raise RuntimeError(f"Cannot find mapping for {c_type}")


def generate_go_wrapper(return_type: str, func_name: str, parameters: List[Tuple[str, str]], original_decl: str) -> str:
    """Generate Go wrapper function for a C function."""
    lines = []
    
    # Generate function signature
    go_params = []
    cgo_params = []
    
    for param_type, param_name in parameters:
        try:
            go_type = map_c_type_to_go(param_type)
            go_params.append(f"{param_name} {go_type}")
            
            # For CGO call, cast unsafe.Pointer parameters to proper C types
            if go_type == 'unsafe.Pointer':
                cgo_type = map_c_type_to_cgo(param_type)
                cgo_params.append(f"({cgo_type})({param_name})")
            else:
                cgo_params.append(f"C.{param_type}({param_name})")
        except ValueError as e:
            raise ValueError(f"Error processing parameter {param_name} in function {func_name}: {e}")
    
    # Generate return type
    try:
        go_return_type = map_c_type_to_go(return_type)
        go_return_annotation = f" {go_return_type}" if go_return_type else ""
    except ValueError as e:
        raise ValueError(f"Error processing return type for function {func_name}: {e}")
    
    # Generate function declaration
    params_str = ", ".join(go_params)
    lines.append(f"func {func_name}({params_str}){go_return_annotation} {{")
    
    # Generate function body
    cgo_call = f"C.{func_name}({', '.join(cgo_params)})"
    
    if go_return_type:
        if go_return_type == 'unsafe.Pointer':
            lines.append(f"\treturn unsafe.Pointer({cgo_call})")
        else:
            lines.append(f"\treturn {go_return_type}({cgo_call})")
    else:
        lines.append(f"\t{cgo_call}")
    
    lines.append("}")
    
    return "\n".join(lines)


def generate_imports_file(c_file_path: str, external_deps: str) -> str:
    """Generate the complete Go imports file."""
    lines = []
    
    # Package declaration
    lines.append("package legacy")
    lines.append("")
    
    # CGO comment block
    lines.append("/*")
    lines.extend(external_deps.split('\n'))
    lines.append("*/")
    lines.append('import "C"')
    lines.append("")
    
    # Go imports
    lines.append("import (")
    lines.append('\t"unsafe"')
    lines.append(")")
    lines.append("")
    
    # Parse external dependencies to extract functions
    func_declarations = []
    for line in external_deps.split('\n'):
        line = line.strip()
        if line and not line.startswith('#') and not line.startswith('extern') and '(' in line:
            func_declarations.append(line)
    
    # Generate wrapper functions
    for decl in func_declarations:
        try:
            parsed = parse_function_declaration(decl)
            if parsed:
                return_type, func_name, parameters, original_decl = parsed
                wrapper = generate_go_wrapper(return_type, func_name, parameters, original_decl)
                lines.append(wrapper)
                lines.append("")
        except ValueError as e:
            print(f"Error generating wrapper for '{decl}': {e}", file=sys.stderr)
            print(f"Skipping function {decl}", file=sys.stderr)
            continue
    
    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser(
        description='Generate Go imports file from C source file'
    )
    parser.add_argument('c_file', help='C source file to analyze')
    parser.add_argument('-o', '--output', help='Output file path (default: <module>_imports.go)')
    
    args = parser.parse_args()
    
    # Get module name from C file
    module_name = os.path.splitext(os.path.basename(args.c_file))[0]
    output_file = args.output or f"{module_name}_imports.go"
    
    # Get external dependencies
    external_deps = get_external_dependencies(args.c_file)
    
    # Generate imports file
    try:
        imports_content = generate_imports_file(args.c_file, external_deps)
        
        # Write output file
        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(imports_content)
        
        print(f"Generated {output_file}")
        
    except Exception as e:
        print(f"Error generating imports file: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == '__main__':
    main()