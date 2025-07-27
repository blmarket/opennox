#!/usr/bin/env python3
"""
Tool to iteratively compile C files and fix unknown type errors using find_decl.
Usage: python -m tools.fix_compile <filename.c>
"""

import subprocess
import re
import sys
import os

def run_command(cmd):
    """Run a command and return the result."""
    try:
        result = subprocess.run(cmd, shell=True, capture_output=True, text=True)
        return result.returncode, result.stdout, result.stderr
    except Exception as e:
        return -1, "", str(e)

def find_declaration(symbol):
    """Use find_decl tool to get declaration for a symbol."""
    cmd = f"python -m tools.find_decl {symbol}"
    returncode, stdout, stderr = run_command(cmd)
    if returncode == 0 and stdout.strip():
        return stdout.strip()
    return None

def extract_unknown_symbols(error_output):
    """Extract unknown symbols from GCC error output."""
    symbols = set()
    lines = error_output.split('\n')
    
    for line in lines:
        # Skip warnings
        if 'warning:' in line:
            continue
        
        # Skip lines that don't contain errors
        if 'error:' not in line:
            continue
        
        # Pattern for "error: 'symbol' undeclared" (handles Unicode quotes)
        match = re.search(r"[''\u2018\u2019]([^'\u2018\u2019]+)[''\u2018\u2019] undeclared", line)
        if match:
            symbols.add(match.group(1))
            continue
        
        # Pattern for "error: unknown type name 'symbol'" (handles Unicode quotes)
        match = re.search(r"unknown type name [''\u2018\u2019]([^'\u2018\u2019]+)[''\u2018\u2019]", line)
        if match:
            symbols.add(match.group(1))
            continue
        
        # Pattern for "error: implicit declaration of function 'symbol'" (handles Unicode quotes)
        match = re.search(r"implicit declaration of function [''\u2018\u2019]([^'\u2018\u2019]+)[''\u2018\u2019]", line)
        if match:
            symbols.add(match.group(1))
            continue
    
    return list(symbols)

def add_declarations_to_file(filename, declarations):
    """Add declarations to the top of the C file."""
    if not declarations:
        return
    
    with open(filename, 'r') as f:
        content = f.read()
    
    # Find the position after includes
    lines = content.split('\n')
    insert_pos = 0
    
    # Find last #include line
    for i, line in enumerate(lines):
        if line.strip().startswith('#include'):
            insert_pos = i + 1
    
    # Insert declarations after includes
    decl_lines = []
    for decl in declarations:
        if decl and not decl in content:
            decl_lines.append(decl)
    
    if decl_lines:
        lines.insert(insert_pos, "")
        lines.insert(insert_pos + 1, "// Declarations added by compile_audio.py")
        for i, decl in enumerate(decl_lines):
            lines.insert(insert_pos + 2 + i, decl)
        lines.insert(insert_pos + 2 + len(decl_lines), "")
        
        with open(filename, 'w') as f:
            f.write('\n'.join(lines))
        
        print(f"Added {len(decl_lines)} declarations to {filename}")

def compile_c_file(filename):
    """Main function to compile a C file iteratively."""
    compile_cmd = f"gcc -m32 -Wno-int-conversion -Wno-incompatible-pointer-types -c {filename}"
    max_iterations = 10
    iteration = 0
    
    print(f"Starting iterative compilation of {filename}...")
    
    while iteration < max_iterations:
        iteration += 1
        print(f"\n--- Iteration {iteration} ---")
        
        # Try to compile
        returncode, stdout, stderr = run_command(compile_cmd)
        
        if returncode == 0:
            print("✓ Compilation successful!")
            return True
        
        print("Compilation failed. Analyzing errors...")
        
        # Extract unknown symbols
        unknown_symbols = extract_unknown_symbols(stderr)
        
        if not unknown_symbols:
            print("No unknown symbols found in error output.")
            print("Error output:")
            print(stderr)
            return False
        
        print(f"Found unknown symbols: {unknown_symbols}")
        
        # Find declarations for each symbol
        declarations = []
        for symbol in unknown_symbols:
            print(f"Looking up declaration for: {symbol}")
            decl = find_declaration(symbol)
            if decl:
                print(f"  Found: {decl}")
                declarations.append(decl)
            else:
                print(f"  No declaration found for: {symbol}")
        
        if not declarations:
            print("No declarations found for any unknown symbols.")
            return False

        # Sort declarations: variables (extern) first, then functions
        declarations.sort(key=lambda x: (not x.strip().startswith('extern'), x))
        
        # Add declarations to the file
        add_declarations_to_file(filename, declarations)
    
    print(f"Failed to compile after {max_iterations} iterations.")
    return False

def compile_audio():
    """Wrapper function for backward compatibility."""
    return compile_c_file("audio.c")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Usage: python -m tools.fix_compile <filename.c>")
        sys.exit(1)
    
    filename = sys.argv[1]
    
    if not os.path.exists(filename):
        print(f"Error: {filename} not found in current directory")
        sys.exit(1)
    
    success = compile_c_file(filename)
    sys.exit(0 if success else 1)