//go:build ccover

// Package ccover exposes the GCC profile flush used by instrumented cgo tests.
package ccover

/*
void __gcov_dump(void);
*/
import "C"

// Dump writes all GCC coverage counters registered in the current process.
// Go terminates with a direct exit system call, so GCC's normal atexit hook is
// not run for Go test binaries.
func Dump() {
	C.__gcov_dump()
}
