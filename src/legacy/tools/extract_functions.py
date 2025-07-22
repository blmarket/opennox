#!/usr/bin/env python3
import re
import sys
import argparse

def extract_function_definitions(c_code):
    """
    Extract function definitions from C code.
    Returns a list of function signatures.
    """
    function_signatures = []
    
    # Remove comments and strings to avoid false matches
    code = remove_comments_and_strings(c_code)
    
    # Pattern to match function definitions
    # Matches: return_type function_name(parameters) {
    pattern = r'^\s*([a-zA-Z_][a-zA-Z0-9_*\s]*?)\s+([a-zA-Z_][a-zA-Z0-9_]*)\s*\([^)]*\)\s*\{'
    
    lines = code.split('\n')
    i = 0
    
    while i < len(lines):
        line = lines[i].strip()
        
        # Skip empty lines and preprocessor directives
        if not line or line.startswith('#'):
            i += 1
            continue
            
        # Look for potential function definition start
        match = re.match(pattern, line, re.MULTILINE)
        if match:
            # Extract the function signature
            return_type = match.group(1).strip()
            function_name = match.group(2).strip()
            
            # Get the complete function signature including parameters
            signature_lines = [line]
            brace_count = line.count('{') - line.count('}')
            
            # If the opening brace is not on the same line, look for it
            j = i + 1
            while j < len(lines) and brace_count == 0:
                next_line = lines[j].strip()
                signature_lines.append(next_line)
                if '{' in next_line:
                    brace_count = 1
                    break
                j += 1
            
            if brace_count > 0:  # Found opening brace, this is a definition
                # Extract parameter list
                full_signature = ' '.join(signature_lines)
                param_match = re.search(r'\(([^)]*)\)', full_signature)
                if param_match:
                    params = param_match.group(1).strip()
                    signature = f"{return_type} {function_name}({params});"
                    function_signatures.append(signature)
            
            i = j + 1 if j < len(lines) else i + 1
        else:
            i += 1
    
    return function_signatures

def remove_comments_and_strings(code):
    """
    Remove C-style comments and string literals to avoid false matches.
    """
    # Remove single-line comments
    code = re.sub(r'//.*$', '', code, flags=re.MULTILINE)
    
    # Remove multi-line comments
    code = re.sub(r'/\*.*?\*/', '', code, flags=re.DOTALL)
    
    # Remove string literals
    code = re.sub(r'"[^"]*"', '""', code)
    code = re.sub(r"'[^']*'", "''", code)
    
    return code

def main():
    parser = argparse.ArgumentParser(description='Extract function definitions from C code')
    parser.add_argument('filename', help='C source file to parse')
    
    args = parser.parse_args()
    
    try:
        with open(args.filename, 'r', encoding='utf-8') as f:
            c_code = f.read()
        
        signatures = extract_function_definitions(c_code)
        
        for signature in signatures:
            print(signature)
            
    except FileNotFoundError:
        print(f"Error: File '{args.filename}' not found", file=sys.stderr)
        sys.exit(1)
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

if __name__ == '__main__':
    main()