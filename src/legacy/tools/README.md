# C-to-Go Translation Tools

This directory contains a comprehensive set of Python tools designed to assist in translating C code to Go, specifically focused on the OpenNox codebase. All tools are designed to work with the assumption of GOARCH=386 (32-bit architecture) and use tree-sitter for robust C code parsing.

## Prerequisites

- Python 3.x
- Tree-sitter and tree-sitter-c libraries (see `requirements.txt`)
- All dependencies are pre-installed as mentioned by the user

## Tool Categories

### 🔄 Core Translation Tools

#### `alias.py`
**Usage**: `python -m tools.alias [function_name]`
- Interactive C-to-Go variable alias replacement tool
- Analyzes C functions to identify pointer arithmetic patterns
- Replaces C pointer arithmetic with proper Go struct field access
- Supports struct200 and struct576 field mappings
- Main workflow tool for translating individual functions

#### `translator.py`
**Usage**: Used as a library module
- Core translation context and logic
- Provides `TranslationContext` class for managing translation state
- Handles struct field mapping and size calculations
- Extract index from multiplication patterns (e.g., "4 * index" → "index")

### 🔍 Analysis Tools

#### `extract_functions.py`
**Usage**: `python -m tools.extract_functions <file.c>`
- Extracts all function declarations AND definitions from C source files
- Returns clean function signatures suitable for Go headers
- Handles multi-line function declarations
- Removes comments and string literals to avoid false matches
- **⚠️ Limitation**: Does NOT separate declarations from definitions - extracts both together
- **Note**: Cannot be used directly to identify external dependencies, as external dependencies are declarations without corresponding definitions in the same file

#### `external_funcs.py`
**Usage**: `python -m tools.external_funcs <file.c> [--verbose]`
- **NEW**: Identifies all external function dependencies in a C file
- Finds functions that are called/declared but NOT defined in the same file
- Uses tree-sitter when available, falls back to regex-based parsing
- Returns sorted list of external dependencies
- **Perfect for creating import files** like `audio_imports.go`
- **Advantage over extract_functions.py**: Only shows external dependencies, not internal functions

#### `create_imports.py`
**Usage**: `python -m tools.create_imports <file.c> [-o output.go]`
- **NEW**: Generates complete Go import files for C modules
- Uses `external_funcs.py` to identify external dependencies
- Creates CGO header with `#include "defs.h"` and external function declarations
- Generates Go wrapper functions with proper C-to-Go type mapping
- Maps C types to Go types (e.g., `uint32_t` → `uint32`, pointers → `unsafe.Pointer`)
- **Perfect for automated creation** of files like `audio_imports.go`
- Follows the same pattern as existing `music_imports.go` and `dialog_imports.go`

#### `find_decl.py`
**Usage**: `python -m tools.find_decl <identifier> [-d directory]`
- Finds proper declarations for global variables and functions
- Searches through all .c and .h files in specified directory
- Handles both function declarations and extern variable declarations
- Essential for creating import files like `audio_imports.go`

#### `type_analysis.py`
**Usage**: Used as a library module
- Analyzes C variable types and declarations
- Extracts type information from parse trees
- Handles pointer types and complex declarations
- Provides `find_variable_type()` and `infer_from_func()` functions

### 🏗️ Structural Analysis Tools

#### `struct_utils.py`
**Usage**: Used as a library module
- Calculates struct field offsets for 32-bit architecture
- Provides field-by-offset mapping functionality
- Essential for converting C pointer arithmetic to Go field access
- Supports complex nested struct analysis

#### `ast_utils.py`
**Usage**: Used as a library module
- Tree-sitter AST navigation utilities
- Find function definitions, struct definitions, variable references
- Core parsing primitives used by other tools

### 🔨 Compilation Assistance

#### `compile_audio.py`
**Usage**: `python -m tools.compile_audio`
- Iteratively compiles audio.c and fixes missing declarations
- Automatically discovers missing symbols from compiler errors
- Uses `find_decl.py` to locate proper declarations
- Adds declarations to source files automatically
- Handles up to 10 compilation iterations

#### `check_all_audio_functions.py`
**Usage**: `python -m tools.check_all_audio_functions`
- Automated analysis of all functions in audio.c
- Runs alias.py on every function definition
- Categorizes results: successful, failed, or no replacements
- Provides comprehensive analysis reports
- Useful for bulk translation assessment

### 🐛 Debugging and Inspection Tools

#### `inspect.py`
**Usage**: Used as a library module
- Debugging utilities for tree-sitter nodes
- `dbg()` function extracts source text from nodes
- `print_tree()` visualizes parse tree structure
- Essential for understanding AST structures

#### `debug_array.py`, `debug_array2.py`, `debug_array3.py`
**Usage**: Various debugging scripts
- Specialized debugging tools for array access patterns
- Help understand specific translation scenarios

#### `debug_fields.py`
**Usage**: Debug script for struct field analysis
- Analyzes struct field access patterns
- Helps validate field offset calculations

### ⚙️ Code Manipulation Tools

#### `move_func.py`
**Usage**: `python -m tools.move_func`
- Moves function definitions between files
- Maintains proper dependencies and declarations
- Handles tree-sitter parsing for accurate extraction

#### `modify_c_function.py`
**Usage**: Tool for modifying C function implementations
- Programmatic function modification capabilities
- Part of the automated translation pipeline

### 🧪 Testing and Validation

#### `test_struct_utils.py`
**Usage**: `pytest test_struct_utils.py`
- Unit tests for struct utility functions
- Validates field offset calculations
- Ensures struct200 and struct576 mappings are correct

#### `test_infer_from_func.py`
**Usage**: `pytest test_infer_from_func.py`
- Tests for function analysis and inference
- Validates type analysis functionality

### 📁 Utility Files

#### `conftest.py`
**Usage**: Pytest configuration
- Test configuration and fixtures

#### `pytest.ini`
**Usage**: Pytest settings
- Test runner configuration

#### `requirements.txt`
**Usage**: `pip install -r requirements.txt`
- Python dependencies (tree-sitter==0.21.3, tree-sitter-c==0.21.4)

## Common Workflows

### 1. Creating External Dependencies (like audio_imports.go)

```bash
# Method 1: FULLY AUTOMATED - Use create_imports.py (RECOMMENDED)
python -m tools.create_imports audio.c

# This single command:
# - Identifies external dependencies using external_funcs.py
# - Creates CGO header with #include "defs.h" 
# - Generates Go wrapper functions with proper type mapping
# - Creates complete audio_imports.go file following existing patterns

# Method 2: Manual analysis using external_funcs.py
python -m tools.external_funcs audio.c --verbose

# Method 3: Use compilation tool to automatically discover missing dependencies  
python -m tools.compile_audio

# Method 4: Manual approach using older tools (when other methods aren't sufficient)
# Step 1: Extract all function signatures from C file (includes both declarations and definitions)
python -m tools.extract_functions audio.c > audio_functions.txt

# Step 2: Manually identify which are external dependencies by checking if definitions exist
# (External dependencies = declarations without definitions in the same file)
# This requires manual analysis

# Step 3: Find proper declarations for identified external dependencies
python -m tools.find_decl some_external_function
```

**✅ Recommended**: Use `create_imports.py` for fully automated Go import file generation.

### 2. Translating Individual Functions

```bash
# Analyze and translate a specific function
python -m tools.alias sub_451920

# Check all functions in bulk
python -m tools.check_all_audio_functions
```

### 3. Debugging Translation Issues

```python
# Use in Python REPL or scripts
from tools.inspect import dbg, print_tree
from tools.ast_utils import find_function_definition

# Debug specific parse tree structures
print_tree(node, source_code, max_depth=4)
```

## Architecture Notes

- All tools assume **GOARCH=386** (32-bit, 4-byte pointers)
- Designed for OpenNox codebase structure with struct200 and struct576
- Handle improper C pointer arithmetic → proper Go struct field access
- Use tree-sitter for robust, syntax-aware parsing
- Maintain original code structure and conventions

## Integration with Claude Agents

These tools are designed to be used by Claude agents for automated C-to-Go translation:

- Tools can be run via `python -m tools.{tool_name}` format
- Most tools provide both CLI and library interfaces
- Error handling and progress reporting suitable for automation
- Designed to work together as a translation pipeline

## Quick Start for audio_imports.go Creation

**✅ READY TO USE (AUTOMATED):**
```bash
python -m tools.create_imports audio.c
```

This single command creates a complete `audio_imports.go` file with:
- CGO header with `#include "defs.h"`
- All external function declarations  
- Go wrapper functions with proper type mapping
- Following the same pattern as `music_imports.go` and `dialog_imports.go`

**Alternative Approaches (if needed):**

**Manual Analysis Approach:**
1. Use `external_funcs.py audio.c` to directly identify all external dependencies
2. Use `find_decl.py` to locate proper declarations for each external dependency
3. Create audio_imports.go following the pattern of music_imports.go and dialog_imports.go

**Compilation-Based Approach:**
1. Use `compile_audio.py` to iteratively discover missing declarations through compilation
2. Use `find_decl.py` to locate proper declarations for discovered missing symbols
3. Create audio_imports.go following existing patterns

**Legacy Manual Approach (if needed):**
1. Use `extract_functions.py` to get all function signatures from audio.c
2. Manually analyze output to identify declarations without definitions (external dependencies)
3. Use `find_decl.py` to find external dependencies
4. Create audio_imports.go following existing patterns

**Note**: The `create_imports.py` tool is now available for fully automated import file generation.