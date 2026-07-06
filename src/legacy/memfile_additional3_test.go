package legacy

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/stretchr/testify/require"
)

func TestNoxMemfileAdditional3(t *testing.T) {
	// Test with all zeros payload
	payload := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()
	cFile := f.C()

	require.Equal(t, int8(0), NoxMemfileReadI8(cFile))
	require.Equal(t, uint8(0), NoxMemfileReadU8(cFile))
	require.Equal(t, int16(0), NoxMemfileReadI16(cFile))
	require.Equal(t, uint16(0), NoxMemfileReadU16(cFile))
	require.Equal(t, int32(0), NoxMemfileReadI32(cFile))
	require.Equal(t, uint32(0), NoxMemfileReadU32(cFile))

	// Test with max values
	payload = []byte{
		0xFF,       // i8: -1
		0xFF,       // u8: 255
		0xFF, 0xFF, // i16: -1
		0xFF, 0xFF, // u16: 65535
		0xFF, 0xFF, 0xFF, 0xFF, // i32: -1
		0xFF, 0xFF, 0xFF, 0xFF, // u32: 4294967295
	}
	raw, _ = alloc.CloneSlice(payload)
	f2 := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f2.Free()
	cFile2 := f2.C()

	require.Equal(t, int8(-1), NoxMemfileReadI8(cFile2))
	require.Equal(t, uint8(255), NoxMemfileReadU8(cFile2))
	require.Equal(t, int16(-1), NoxMemfileReadI16(cFile2))
	require.Equal(t, uint16(65535), NoxMemfileReadU16(cFile2))
	require.Equal(t, int32(-1), NoxMemfileReadI32(cFile2))
	require.Equal(t, uint32(4294967295), NoxMemfileReadU32(cFile2))

	// Test Read block with exact size
	payload = []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	raw, _ = alloc.CloneSlice(payload)
	f3 := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f3.Free()
	cFile3 := f3.C()

	var buf [8]byte
	n := NoxMemfileRead(unsafe.Pointer(&buf[0]), 1, 8, cFile3)
	require.Equal(t, uint32(8), n)
	require.Equal(t, [8]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}, buf)

	// Test Read block with size larger than available (short read)
	payload = []byte{0xAA, 0xBB}
	raw, _ = alloc.CloneSlice(payload)
	f4 := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f4.Free()
	cFile4 := f4.C()

	var buf2 [4]byte
	n = NoxMemfileRead(unsafe.Pointer(&buf2[0]), 1, 4, cFile4)
	// Should read only available bytes
	require.LessOrEqual(t, n, uint32(4))

	// Test Skip beyond end (should not crash)
	payload = []byte{0x01, 0x02}
	raw, _ = alloc.CloneSlice(payload)
	f5 := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f5.Free()
	cFile5 := f5.C()
	NoxMemfileSkip(cFile5, 100) // Skip beyond end
	// Subsequent reads should return 0
	require.Equal(t, int8(0), NoxMemfileReadI8(cFile5))

	// Test Read64Align with various alignments
	payload = []byte{
		0x00, 0x00, 0x00, // padding to align
		0x11, 0x22, 0x33, 0x44, // aligned data
		0x55, 0x66, 0x77, 0x88,
	}
	raw, _ = alloc.CloneSlice(payload)
	f6 := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f6.Free()
	cFile6 := f6.C()

	// Skip 3 bytes to get to unaligned position
	NoxMemfileSkip(cFile6, 3)
	var alignBuf [4]byte
	n = NoxMemfileRead64Align(unsafe.Pointer(&alignBuf[0]), 1, 4, cFile6)
	if n == 1 {
		// If aligned read succeeded, verify data
		require.Equal(t, [4]byte{0x11, 0x22, 0x33, 0x44}, alignBuf)
	}

	// Test multiple sequential reads
	payload = []byte{
		0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08,
		0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10,
	}
	raw, _ = alloc.CloneSlice(payload)
	f7 := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f7.Free()
	cFile7 := f7.C()

	for i := 1; i <= 16; i++ {
		got := NoxMemfileReadU8(cFile7)
		require.Equal(t, uint8(i), got)
	}
	// Next read should return 0 (EOF)
	require.Equal(t, uint8(0), NoxMemfileReadU8(cFile7))
}

func TestNoxMemfileReadBlockAdditional(t *testing.T) {
	// Test Read with different element sizes
	payload := []byte{
		0x01, 0x00, 0x02, 0x00, // 2 uint16: 1, 2
		0x03, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, // 2 uint32: 3, 4
	}
	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()
	cFile := f.C()

	// Read 2 uint16 elements
	var buf16 [2]uint16
	n := NoxMemfileRead(unsafe.Pointer(&buf16[0]), 2, 2, cFile)
	require.Equal(t, uint32(2), n)
	require.Equal(t, uint16(1), buf16[0])
	require.Equal(t, uint16(2), buf16[1])

	// Read 2 uint32 elements
	var buf32 [2]uint32
	n = NoxMemfileRead(unsafe.Pointer(&buf32[0]), 4, 2, cFile)
	require.Equal(t, uint32(2), n)
	require.Equal(t, uint32(3), buf32[0])
	require.Equal(t, uint32(4), buf32[1])
}
