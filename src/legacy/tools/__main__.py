#!/usr/bin/env python3

import sys

if __name__ == '__main__':
    if len(sys.argv) < 2:
        print("Usage: python -m tools <command> [args...]")
        print("Available commands:")
        print("  update_imports <module_name> <func_name>")
        sys.exit(1)
    
    command = sys.argv[1]
    
    if command == 'update_imports':
        from .update_imports import main as update_imports_main
        # Remove the command from argv so update_imports sees the right args
        sys.argv = ['update_imports'] + sys.argv[2:]
        update_imports_main()
    else:
        print(f"Unknown command: {command}")
        sys.exit(1)