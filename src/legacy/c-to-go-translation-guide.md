# C-to-Go Translation Guide for OpenNox Legacy Modules

This guide provides patterns and best practices for translating C modules to Go based on the existing `music/` and `dialog/` module implementations.

## Architecture Overview

### Module Structure Pattern
Each translated module follows this structure:
- **Core Module Package**: Contains the main Go implementation (e.g., `music/music.go`, `dialog/dialog.go`)
- **Import File**: Provides CGO interface and initialization (e.g., `music_imports.go`, `dialog_imports.go`)
- **External Dependencies**: Handled through CGO headers and pointer mapping

### Key Architectural Assumptions
- **GOARCH=386**: All pointer arithmetic assumes 4-byte pointers (32-bit architecture)
- **Memory Layout**: Direct memory mapping to preserve C struct layouts and data compatibility
- **CGO Integration**: Seamless bidirectional function calls between C and Go

## CGO Import File Patterns

### File Structure Template
```go
package legacy

/*
#include <stdint.h>

// External variable declarations (from C)
extern uint32_t dword_5d4594_816356;
extern void* dword_587000_81128;
// ... more C variables
*/
import "C"

import (
    "unsafe"
    // Module-specific imports
    "github.com/noxworld-dev/opennox/v1/legacy/modulename"
)

var (
    ModuleInstance *modulename.Module
)

func initModule() {
    // Initialize with CGO variable mappings
}

//export c_function_name
func c_function_name(params) {
    // Export Go functions to C
}
```

### CGO Header Guidelines
1. **No Additional Includes**: Only declare function signatures individually
2. **Variable Declarations**: Use `extern` for C global variables
3. **Type Consistency**: Maintain exact type signatures from C files
4. **Pointer Safety**: Use `unsafe.Pointer` for C pointer interop

### Variable Mapping Patterns
```go
// Pattern 1: Direct pointer mapping to C globals
var_name := (*uint32)(&C.c_variable_name)

// Pattern 2: Memory-mapped pointers using memmap
ptr_name := memmap.PtrT[Type](base_addr, offset)

// Pattern 3: Double pointer handling
dbl_ptr := (**Type)(unsafe.Pointer(&C.c_pointer_var))
```

## Module Implementation Patterns

### Constructor Pattern
```go
func NewModule(
    dir string,
    // External dependencies as parameters
    external_var1 *uint32,
    external_var2 *SomeType,
    // Callback functions
    callback1 func() ReturnType,
    callback2 func(params) ReturnType,
) *Module {
    return &Module{
        // Initialize all fields
    }
}
```

### State Management
- **Private State**: Internal module variables (counters, timers, etc.)
- **External References**: Pointers to C global variables
- **Callback Functions**: Interface with other modules via function pointers

### Memory Safety
```go
// Size assertions for struct compatibility
var _ = [1]struct{}{}[expectedSize-unsafe.Sizeof(StructType{})]
```

## Common Translation Patterns

### Function Translation
```c
// C function
int sub_43D9E0(void* a1) {
    // implementation
}
```

```go
// Go translation with CGO export
//export sub_43D9E0
func sub_43D9E0(a1p unsafe.Pointer) {
    ModuleInstance.HandleFunction(*(*ParamType)(a1p))
}
```

### State Machine Translation
```go
// Translate C switch statements to Go switch
switch *moduleState {
case 0:
    // State 0 logic
case 1:
    // State 1 logic
default:
    // Default case
}
```

### Timer Integration
```go
// Timer objects are consistently used across modules
timers := &timer.TimerGroup{}
timers.Init()
timers.Timers[0].SetParams(period, initial_value)
```

## Data Structure Patterns

### C Struct to Go Struct
```go
type ModuleState struct {
    Field1, Field2, Field3, Field4 uint32
}

// Always include size assertion
var _ = [1]struct{}{}[16-unsafe.Sizeof(ModuleState{})]
```

### Module Container Pattern
```go
type Module struct {
    // Configuration
    dir string
    
    // Private state
    internalCounter int32
    stateFlags      uint32
    
    // External references (pointers to C globals)
    externalVar1 *uint32
    externalVar2 *SomeType
    
    // Callback functions for cross-module communication
    callback1 func() ReturnType
    callback2 func(params) ReturnType
}
```

## Error Handling and Validation

### File Operations
```go
// Pattern: Try primary path, fallback if needed
s := driver.OpenStream(primaryPath)
if s == 0 {
    if fallbackCondition() {
        return false
    }
    fallbackPath := getFallbackPath()
    s = driver.OpenStream(fallbackPath)
}
```

### Parameter Validation
```go
// Always validate bounds for arrays/slices
if index <= 0 || index >= uint32(len(table)) {
    return false
}

// Clamp values to valid ranges
if volume > 100 {
    volume = 100
}
```

## Integration Guidelines

### Initialization Order
1. Initialize CGO variable mappings
2. Create module instance with dependencies
3. Set up callback functions
4. Register exported functions

### Cross-Module Communication
- Use callback functions for loose coupling
- Pass pointers to shared state variables
- Maintain consistent interfaces across modules

### Testing Considerations
- Include test files with module (see `dialog_test.go`)
- Use testdata directories for fixtures
- Test both normal and error conditions

## Best Practices

1. **Preserve C Semantics**: Maintain exact behavior of original C code
2. **Memory Layout**: Keep struct sizes and layouts identical
3. **Pointer Arithmetic**: Avoid when possible, use proper struct field access
4. **Error Propagation**: Handle C-style error codes appropriately
5. **Resource Management**: Ensure proper cleanup of streams and handles
6. **State Synchronization**: Maintain consistency between C and Go state

## Common Pitfalls

1. **Struct Padding**: Go and C struct padding may differ
2. **Pointer Arithmetic**: Avoid improper pointer arithmetic from C code
3. **String Handling**: Use proper Go string conversion utilities
4. **Memory Management**: Be careful with CGO memory allocation/deallocation
5. **Function Signatures**: Maintain exact parameter types for CGO exports

This guide should be used as a reference when creating new module translations, ensuring consistency with the established patterns in the codebase.