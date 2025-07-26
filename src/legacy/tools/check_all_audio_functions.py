#!/usr/bin/env python3
"""
Automated script to check all function declarations in audio.c 
against tools/alias.py for potential alias variable replacements.
"""

import os
import sys
import subprocess
import time
from pathlib import Path
import tree_sitter_c as tsc
from tree_sitter import Language, Parser

# Change to the parent directory so we can access audio.c
os.chdir(os.path.join(os.path.dirname(__file__), ".."))

def get_source_text(node, source_bytes):
    """Extract the source text for a given node."""
    return source_bytes[node.start_byte:node.end_byte].decode('utf-8')

def find_function_definitions(tree, source_bytes):
    """Find all function definitions (not just declarations) in the parse tree."""
    function_names = []
    
    def traverse(node):
        if node.type == 'function_definition':
            # Find the function declarator to get the function name
            for child in node.children:
                if child.type == 'function_declarator':
                    # Find the identifier (function name)
                    for subchild in child.children:
                        if subchild.type == 'identifier':
                            func_name = get_source_text(subchild, source_bytes)
                            # Filter to keep local functions (sub_*, nox_*, and other local functions)
                            # Skip standard library or system functions
                            if not func_name.startswith('_') and func_name not in ['main', 'printf', 'malloc', 'free']:
                                function_names.append(func_name)
                            break
                    break
        
        for child in node.children:
            traverse(child)
    
    traverse(tree.root_node)
    return function_names

def extract_function_names_from_audio():
    """Extract all function definition names from audio.c using tree-sitter."""
    
    # Set up tree-sitter
    C_LANGUAGE = Language(tsc.language(), "c")
    parser = Parser()
    parser.set_language(C_LANGUAGE)
    
    # Parse audio.c
    try:
        with open("audio.c", 'rb') as f:
            source_bytes = f.read()
        
        tree = parser.parse(source_bytes)
        function_names = find_function_definitions(tree, source_bytes)
        
        return sorted(function_names)
    
    except Exception as e:
        print(f"Error parsing audio.c: {e}", file=sys.stderr)
        return []

def run_alias_check(func_name, timeout=30):
    """Run tools/alias.py on a specific function name."""
    try:
        # Change to tools directory and run alias.py with proper Python path
        env = os.environ.copy()
        env['PYTHONPATH'] = str(Path('.').resolve())
        
        result = subprocess.run([
            sys.executable, 'alias.py', func_name
        ], capture_output=True, text=True, timeout=timeout, cwd='tools', env=env)
        
        return {
            'returncode': result.returncode,
            'stdout': result.stdout.strip(),
            'stderr': result.stderr.strip()
        }
    
    except subprocess.TimeoutExpired:
        return {
            'returncode': -1,
            'stdout': '',
            'stderr': f'Timeout after {timeout} seconds'
        }
    except Exception as e:
        return {
            'returncode': -2,
            'stdout': '',
            'stderr': str(e)
        }

def categorize_result(func_name, result):
    """Categorize the result of alias checking."""
    if result['returncode'] == 0:
        output = result['stdout'].lower()
        if 'replaced' in output or 'replacement' in output or 'alias' in output or 'field' in output:
            return 'successful', result['stdout']
        else:
            return 'no_replacements', 'No alias replacements found'
    else:
        # Check if there's useful output even with error (e.g., showed references but failed on replacement)
        error_output = result['stderr']
        if 'found' in result['stdout'].lower() and ('references' in result['stdout'].lower() or 'variable' in result['stdout'].lower()):
            # This means it found the function and variables, but maybe failed during replacement
            if 'list index out of range' in error_output or 'IndexError' in error_output:
                return 'no_replacements', 'Function analyzed but no replacements possible (insufficient references)'
            else:
                return 'failed', error_output or 'Unknown error'
        else:
            return 'failed', error_output or 'Unknown error'

def main():
    print("=== Audio.c Function Alias Analysis ===\n")
    
    # Extract function names
    print("📁 Parsing audio.c to extract function names...")
    function_names = extract_function_names_from_audio()
    
    if not function_names:
        print("❌ No function names found in audio.c")
        sys.exit(1)
    
    print(f"📋 Found {len(function_names)} function definitions")
    print(f"Functions to analyze: {', '.join(function_names[:5])}{'...' if len(function_names) > 5 else ''}\n")
    
    # Process each function
    successful_functions = []
    failed_functions = []
    no_replacement_functions = []
    
    for i, func_name in enumerate(function_names, 1):
        print(f"🔍 [{i:2d}/{len(function_names)}] Checking {func_name}...", end=' ')
        
        result = run_alias_check(func_name)
        category, message = categorize_result(func_name, result)
        
        if category == 'successful':
            successful_functions.append((func_name, message))
            print("✅ Found replacements")
        elif category == 'failed':
            failed_functions.append((func_name, message))
            print("❌ Error")
        else:
            no_replacement_functions.append((func_name, message))
            print("⚪ No replacements")
        
        # Small delay to avoid overwhelming the system
        time.sleep(0.1)
    
    # Generate summary report
    print(f"\n{'='*60}")
    print("📊 ANALYSIS RESULTS")
    print(f"{'='*60}")
    
    if successful_functions:
        print(f"\n✅ SUCCESSFUL REPLACEMENTS ({len(successful_functions)} functions):")
        for func_name, message in successful_functions:
            # Show first line of output to keep it concise
            first_line = message.split('\n')[0] if message else 'Found alias replacements'
            print(f"   • {func_name}: {first_line}")
    
    if failed_functions:
        print(f"\n❌ FAILED ANALYSIS ({len(failed_functions)} functions):")
        for func_name, error in failed_functions:
            # Show first line of error to keep it concise
            first_line = error.split('\n')[0] if error else 'Unknown error'
            print(f"   • {func_name}: {first_line}")
    
    if no_replacement_functions:
        print(f"\n⚪ NO REPLACEMENTS FOUND ({len(no_replacement_functions)} functions):")
        # Show first 10 to avoid cluttering output
        for func_name, _ in no_replacement_functions[:10]:
            print(f"   • {func_name}")
        if len(no_replacement_functions) > 10:
            print(f"   ... and {len(no_replacement_functions) - 10} more")
    
    # Summary statistics
    print(f"\n📈 SUMMARY:")
    print(f"   Total functions analyzed: {len(function_names)}")
    print(f"   Successful alias replacements: {len(successful_functions)} functions")
    print(f"   Functions with errors: {len(failed_functions)} functions")
    print(f"   Functions with no aliases: {len(no_replacement_functions)} functions")
    
    if successful_functions:
        print(f"\n💡 TIP: Run 'python tools/alias.py <function_name>' on successful functions for detailed output")

if __name__ == '__main__':
    main()