#!/usr/bin/env python3

import sys
import os
import argparse
import re
from pathlib import Path

# Import the find_decl module
sys.path.insert(0, os.path.dirname(os.path.abspath(__file__)))
from find_decl import find_declaration

def read_file(file_path):
    """Read file content."""
    with open(file_path, 'r', encoding='utf-8') as f:
        return f.read()

def write_file(file_path, content):
    """Write content to file."""
    with open(file_path, 'w', encoding='utf-8') as f:
        f.write(content)

def add_cgo_declaration(go_file_path, declaration):
    """Add a declaration to the CGO import section in the Go file."""
    content = read_file(go_file_path)
    
    # Find the CGO comment block (between /* and */)
    cgo_pattern = r'(/\*.*?\*/)'
    match = re.search(cgo_pattern, content, re.DOTALL)
    
    if not match:
        print(f"Error: Could not find CGO comment block in {go_file_path}", file=sys.stderr)
        return False
    
    cgo_block = match.group(1)
    
    # Check if declaration already exists
    if declaration.rstrip(';') in cgo_block:
        print(f"Declaration already exists in {go_file_path}")
        return True
    
    # Add semicolon if not present
    if not declaration.rstrip().endswith(';'):
        declaration = declaration.rstrip() + ';'
    
    # Find the end of the CGO block (before */)
    end_pattern = r'\*/'
    cgo_end_match = re.search(end_pattern, cgo_block)
    
    if not cgo_end_match:
        print(f"Error: Could not find end of CGO comment block in {go_file_path}", file=sys.stderr)
        return False
    
    # Insert the declaration before the closing */
    insert_pos = cgo_end_match.start()
    new_cgo_block = cgo_block[:insert_pos] + declaration + '\n' + cgo_block[insert_pos:]
    
    # Replace the old CGO block with the new one
    new_content = content.replace(cgo_block, new_cgo_block)
    
    # Write the updated content back to the file
    write_file(go_file_path, new_content)
    print(f"Added declaration to {go_file_path}: {declaration}")
    return True

def main():
    parser = argparse.ArgumentParser(
        description="Find function declaration and add it to CGO import section in audio_imports.go"
    )
    parser.add_argument("identifier", help="Function or variable name to find")
    parser.add_argument("-g", "--go-file", 
                       default="audio_imports.go", 
                       help="Go file to update (default: audio_imports.go)")
    parser.add_argument("-d", "--search-dir", 
                       default=".", 
                       help="Directory to search for C files (default: current directory)")
    
    args = parser.parse_args()
    
    # Find the declaration
    print(f"Searching for declaration of '{args.identifier}'...")
    declaration = find_declaration(args.identifier, args.search_dir)
    
    if not declaration:
        print(f"Error: Declaration for '{args.identifier}' not found", file=sys.stderr)
        sys.exit(1)
    
    print(f"Found declaration: {declaration}")
    
    # Check if Go file exists
    if not os.path.exists(args.go_file):
        print(f"Error: Go file '{args.go_file}' not found", file=sys.stderr)
        sys.exit(1)
    
    # Add declaration to CGO section
    if add_cgo_declaration(args.go_file, declaration):
        print("Successfully updated CGO declarations")
    else:
        print("Failed to update CGO declarations", file=sys.stderr)
        sys.exit(1)

if __name__ == "__main__":
    main()