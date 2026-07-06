package cnxz

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestSub57DD90AndSub57DDC0(t *testing.T) {
	// sub57DD90 allocates a 0x224 byte buffer and stores the pointer in *this
	// sub57DDC0 frees the buffer pointed to by *this

	thisPtr, bufPtr := C_sub57DD90()
	require.NotEqual(t, uintptr(0), thisPtr)
	require.NotEqual(t, uintptr(0), bufPtr)
	require.Equal(t, thisPtr, bufPtr)

	// Verify the buffer is zeroed (0x224 = 548 bytes)
	buf := unsafe.Slice((*byte)(unsafe.Pointer(bufPtr)), 0x224)
	for i, b := range buf {
		require.Equal(t, byte(0), b, "byte at offset %d should be 0", i)
	}

	// Write some data to the buffer to verify it's writable
	buf[0] = 0xAB
	buf[1] = 0xCD
	buf[0x223] = 0xEF
	require.Equal(t, byte(0xAB), buf[0])
	require.Equal(t, byte(0xCD), buf[1])
	require.Equal(t, byte(0xEF), buf[0x223])

	// Free the buffer using sub57DDC0
	C_sub57DDC0(bufPtr)

	// Note: After free, the memory should not be accessed.
	// We don't verify the free succeeded as that would be undefined behavior.
}

func TestSub57DD90MultipleAllocations(t *testing.T) {
	// Test multiple allocations to ensure each is independent
	this1, buf1 := C_sub57DD90()
	this2, buf2 := C_sub57DD90()
	require.NotEqual(t, buf1, buf2, "each allocation should return a different buffer")

	// Write different data to each buffer
	buf1Slice := unsafe.Slice((*byte)(unsafe.Pointer(buf1)), 0x224)
	buf2Slice := unsafe.Slice((*byte)(unsafe.Pointer(buf2)), 0x224)
	buf1Slice[0] = 0x11
	buf2Slice[0] = 0x22
	require.Equal(t, byte(0x11), buf1Slice[0])
	require.Equal(t, byte(0x22), buf2Slice[0])

	// Free both
	C_sub57DDC0(buf1)
	C_sub57DDC0(buf2)

	_ = this1
	_ = this2
}
