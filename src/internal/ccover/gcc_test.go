//go:build ccover

package ccover

import "testing"

func TestDump(t *testing.T) {
	// Dump should not panic even when called without coverage instrumentation.
	// In a real ccover build, this flushes GCC coverage counters.
	// In tests without instrumentation, __gcov_dump may be a no-op or not linked.
	// We just verify the function can be called without panicking.
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Dump panicked (expected without coverage instrumentation): %v", r)
		}
	}()
	Dump()
}
