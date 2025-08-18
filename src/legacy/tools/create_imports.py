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
    'int': 'int',
    'unsigned int': 'uint',
    'signed int': 'int32',
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
    'int**': '**C.int',
    'int*': '*C.int',
    'char*': '*C.char',
    'FILE*': '*C.FILE',
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


def parse_external_dependencies(external_deps: str) -> Tuple[List[Tuple[str, str]], List[Tuple[str, str, List[Tuple[str, str]], str]]]:
    """Parse external dependencies into variables and functions.

    Returns:
        (variables, functions) where:
        - variables: List of (type, name) tuples
        - functions: List of (return_type, func_name, parameters, original_decl) tuples
    """
    variables = []
    functions = []

    for line in external_deps.split('\n'):
        line = line.strip()
        if not line or line.startswith('#'):
            continue

        # Check if it's a function declaration (contains parentheses)
        if '(' in line and ')' in line:
            parsed = parse_function_declaration(line)
            if parsed:
                functions.append(parsed)
        # Check if it's an extern variable declaration
        elif line.startswith('extern'):
            # Parse variable declaration: extern type name;
            var_match = re.match(r'extern\s+(.+?)\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*;?$', line)
            if var_match:
                var_type = var_match.group(1).strip()
                var_name = var_match.group(2).strip()
                variables.append((var_type, var_name))

    return variables, functions


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


def generate_module_struct(module_name: str, variables: List[Tuple[str, str]], functions: List[Tuple[str, str, List[Tuple[str, str]], str]]) -> str:
    """Generate the module struct definition."""
    lines = []
    struct_name = f"{module_name.capitalize()}Module"

    lines.append(f"type {struct_name} struct {{")
    lines.append("\tmoduleName string")
    lines.append("")
    lines.append("\t// External variables")

    for var_type, var_name in variables:
        try:
            go_type = map_c_type_to_go(var_type)
            lines.append(f"\t{var_name} *{go_type}")
        except (ValueError, RuntimeError):
            # Fall back to unsafe.Pointer for unknown types
            lines.append(f"\t{var_name} unsafe.Pointer")

    if functions:
        lines.append("")
        lines.append("\t// External functions")
        for return_type, func_name, parameters, _ in functions:
            # Generate function field type
            try:
                param_types = []
                for param_type, _ in parameters:
                    go_type = map_c_type_to_go(param_type)
                    param_types.append(go_type)

                go_return_type = map_c_type_to_go(return_type)
                return_annotation = go_return_type if go_return_type else ""

                if return_annotation:
                    func_type = f"func({', '.join(param_types)}) {return_annotation}"
                else:
                    func_type = f"func({', '.join(param_types)})"

                lines.append(f"\t{func_name} {func_type}")
            except (ValueError, RuntimeError):
                # Skip functions with unknown types
                lines.append(f"\t// TODO: {func_name} - unknown parameter types")

    lines.append("}")

    return "\n".join(lines)


def generate_module_constructor(module_name: str, variables: List[Tuple[str, str]], functions: List[Tuple[str, str, List[Tuple[str, str]], str]]) -> str:
    """Generate the module constructor function."""
    lines = []
    struct_name = f"{module_name.capitalize()}Module"
    constructor_name = f"New{module_name.capitalize()}Module"

    # Generate constructor signature
    params = [f"moduleName string"]

    for var_type, var_name in variables:
        try:
            go_type = map_c_type_to_go(var_type)
            params.append(f"{var_name} *{go_type}")
        except (ValueError, RuntimeError):
            params.append(f"{var_name} unsafe.Pointer")

    for return_type, func_name, parameters, _ in functions:
        try:
            param_types = []
            for param_type, _ in parameters:
                go_type = map_c_type_to_go(param_type)
                param_types.append(go_type)

            go_return_type = map_c_type_to_go(return_type)
            return_annotation = go_return_type if go_return_type else ""

            if return_annotation:
                func_type = f"func({', '.join(param_types)}) {return_annotation}"
            else:
                func_type = f"func({', '.join(param_types)})"

            params.append(f"{func_name} {func_type}")
        except (ValueError, RuntimeError):
            # Skip functions with unknown types
            continue

    lines.append(f"func {constructor_name}(")
    for i, param in enumerate(params):
        if i == len(params) - 1:
            lines.append(f"\t{param},")
        else:
            lines.append(f"\t{param},")
    lines.append(f") *{struct_name} {{")

    # Generate constructor body
    lines.append(f"\treturn &{struct_name}{{")
    lines.append(f"\t\tmoduleName: moduleName,")

    for var_type, var_name in variables:
        lines.append(f"\t\t{var_name}: {var_name},")

    for return_type, func_name, parameters, _ in functions:
        try:
            # Check if we can map all types
            for param_type, _ in parameters:
                map_c_type_to_go(param_type)
            map_c_type_to_go(return_type)
            lines.append(f"\t\t{func_name}: {func_name},")
        except (ValueError, RuntimeError):
            # Skip functions with unknown types
            continue

    lines.append("\t}")
    lines.append("}")

    return "\n".join(lines)


def generate_module_file(module_name: str, variables: List[Tuple[str, str]], functions: List[Tuple[str, str, List[Tuple[str, str]], str]]) -> str:
    """Generate the module Go file (xxx/xxx.go)."""
    lines = []

    # Package declaration
    lines.append(f"package audio")
    lines.append("")

    # Imports
    lines.append("import (")
    lines.append('\t"unsafe"')
    lines.append(")")
    lines.append("")

    # Generate struct
    struct_def = generate_module_struct(module_name, variables, functions)
    lines.append(struct_def)
    lines.append("")

    # Generate constructor
    constructor_def = generate_module_constructor(module_name, variables, functions)
    lines.append(constructor_def)

    return "\n".join(lines)


def generate_imports_file(c_file_path: str, external_deps: str, module_name: str) -> str:
    """Generate the complete Go imports file."""
    lines = []
    variables, functions = parse_external_dependencies(external_deps)

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
    lines.append("")
    lines.append(f'\t"github.com/noxworld-dev/opennox/v1/legacy/{module_name}"')
    lines.append(")")
    lines.append("")

    # Package variable
    struct_name = f"{module_name.capitalize()}Module"
    var_name = f"{module_name.capitalize()}Module"
    lines.append("var (")
    lines.append(f"\t{var_name} *{module_name}.{struct_name}")
    lines.append(")")
    lines.append("")

    # Generate init function
    init_func = generate_init_function(module_name, variables, functions)
    lines.append(init_func)
    lines.append("")

    # Generate wrapper functions
    for return_type, func_name, parameters, original_decl in functions:
        try:
            wrapper = generate_go_wrapper(return_type, func_name, parameters, original_decl)
            lines.append(wrapper)
            lines.append("")
        except ValueError as e:
            print(f"Error generating wrapper for '{original_decl}': {e}", file=sys.stderr)
            print(f"Skipping function {func_name}", file=sys.stderr)
            continue

    return "\n".join(lines)


def generate_init_function(module_name: str, variables: List[Tuple[str, str]], functions: List[Tuple[str, str, List[Tuple[str, str]], str]]) -> str:
    """Generate the initXxx() function."""
    lines = []
    init_func_name = f"init{module_name.capitalize()}"
    var_name = f"{module_name.capitalize()}Module"
    constructor_name = f"New{module_name.capitalize()}Module"

    lines.append(f"func {init_func_name}() {{")

    # Generate variable declarations for CGO access
    if variables:
        lines.append("\tvar (")
        for var_type, var_name_local in variables:
            try:
                go_type = map_c_type_to_go(var_type)
                if go_type == 'unsafe.Pointer':
                    cgo_type = map_c_type_to_cgo(var_type)
                    lines.append(f"\t\t{var_name_local} {cgo_type} = ({cgo_type})(unsafe.Pointer(&C.{var_name_local}))")
                else:
                    lines.append(f"\t\t{var_name_local} *{go_type} = (*{go_type})(&C.{var_name_local})")
            except (ValueError, RuntimeError):
                lines.append(f"\t\t{var_name_local} unsafe.Pointer = unsafe.Pointer(&C.{var_name_local})")
        lines.append("\t)")

    # Generate constructor call
    lines.append(f"\t{var_name} = {module_name}.{constructor_name}(")
    lines.append(f'\t\t"{module_name}",')

    # Add variable parameters
    for var_type, var_name_local in variables:
        lines.append(f"\t\t{var_name_local},")

    # Add function parameters - use the wrapper functions as parameters
    for return_type, func_name, parameters, _ in functions:
        try:
            # Check if we can map all types
            for param_type, _ in parameters:
                map_c_type_to_go(param_type)
            map_c_type_to_go(return_type)
            lines.append(f"\t\t{func_name},")
        except (ValueError, RuntimeError):
            # Skip functions with unknown types
            continue

    lines.append("\t)")
    lines.append("}")

    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser(
        description='Generate Go imports file and module package from C source file'
    )
    parser.add_argument('c_file', help='C source file to analyze')
    parser.add_argument('-o', '--output', help='Output file path (default: <module>_imports.go)')

    args = parser.parse_args()

    # Get module name from C file
    module_name = os.path.splitext(os.path.basename(args.c_file))[0]
    output_file = args.output or f"{module_name}_imports.go"

    # Get external dependencies
    external_deps = get_external_dependencies(args.c_file)
    variables, functions = parse_external_dependencies(external_deps)

    try:
        # Create module directory
        module_dir = "audio"
        os.makedirs(module_dir, exist_ok=True)

        # Generate module file (xxx/xxx.go)
        module_content = generate_module_file(module_name, variables, functions)
        module_file_path = os.path.join(module_dir, f"{module_name}.go")

        with open(module_file_path, 'w', encoding='utf-8') as f:
            f.write(module_content)

        print(f"Generated {module_file_path}")

        # Generate imports file (xxx_imports.go)
        imports_content = generate_imports_file(args.c_file, external_deps, module_name)

        with open(output_file, 'w', encoding='utf-8') as f:
            f.write(imports_content)

        print(f"Generated {output_file}")

    except Exception as e:
        print(f"Error generating files: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == '__main__':
    main()
