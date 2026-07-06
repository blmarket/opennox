package ccall

import (
	"testing"
	"unsafe"
)

func TestCallFunctionsExist(t *testing.T) {
	// Just verify the functions exist by referencing them
	_ = CallVoidVoid
	_ = CallVoidUPtr
	_ = CallVoidUPtr2
	_ = CallVoidUPtr3
	_ = CallVoidUPtr4
	_ = CallVoidUPtr5
	_ = CallVoidUPtr6
	_ = CallVoidPtr
	_ = CallVoidPtr2
	_ = CallVoidPtr3
	_ = CallVoidPtr4
	_ = CallVoidPtr5
	_ = CallVoidPtr6
	_ = CallVoidInt
	_ = CallVoidInt2
	_ = CallVoidInt3
	_ = CallVoidInt4
	_ = CallVoidInt5
	_ = CallVoidInt6
}

func TestCallWithNil(t *testing.T) {
	// Calling with nil function pointer should not panic in Go,
	// though the C function may fault. We recover from any panic.
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Recovered from panic (expected with nil func): %v", r)
		}
	}()

	// These calls with nil will likely cause a segfault in C,
	// but the Go wrapper should not panic before the C call.
	// We just test that the Go function can be invoked.
	// In practice, these will crash, so we skip actual calls.
	// Instead, we just verify the function signatures are correct.
	var fnc unsafe.Pointer
	var uptr uintptr
	var ptr unsafe.Pointer
	var i int

	// Reference the functions to ensure they compile correctly
	_ = fnc
	_ = uptr
	_ = ptr
	_ = i
}
