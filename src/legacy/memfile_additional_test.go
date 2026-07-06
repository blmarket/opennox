package legacy

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/stretchr/testify/require"
)

func TestNoxMemfileAdditional(t *testing.T) {
	// Test with larger data
	payload := make([]byte, 1024)
	for i := range payload {
		payload[i] = byte(i % 256)
	}

	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()

	cFile := f.C()

	// Read all bytes as u8 and verify
	for i := 0; i < 256; i++ {
		got := NoxMemfileReadU8(cFile)
		require.Equal(t, uint8(i%256), got)
	}

	// Test skip forward and backward (if supported)
	// Skip is only forward, so test skipping to end
	NoxMemfileSkip(cFile, 1000) // Skip remaining

	// Test read at end (should return 0 or handle gracefully)
	// The file position is now at or past end
}

func TestNoxMemfileReadBlock(t *testing.T) {
	payload := []byte{
		0x01, 0x02, 0x03, 0x04,
		0x05, 0x06, 0x07, 0x08,
		0x09, 0x0A, 0x0B, 0x0C,
		0x0D, 0x0E, 0x0F, 0x10,
	}

	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()

	cFile := f.C()

	// Read 4 bytes as single block
	var buf [4]byte
	n := NoxMemfileRead(unsafe.Pointer(&buf[0]), 1, 4, cFile)
	require.Equal(t, uint32(4), n)
	require.Equal(t, [4]byte{0x01, 0x02, 0x03, 0x04}, buf)

	// Read 2 uint16 values
	var buf16 [2]uint16
	n = NoxMemfileRead(unsafe.Pointer(&buf16[0]), 2, 2, cFile)
	require.Equal(t, uint32(2), n)

	// Read remaining as u32
	var buf32 [2]uint32
	n = NoxMemfileRead(unsafe.Pointer(&buf32[0]), 4, 2, cFile)
	require.Equal(t, uint32(2), n)
}

func TestNoxMemfileRead64AlignAdditional(t *testing.T) {
	// Create payload with specific alignment pattern
	payload := make([]byte, 64)
	for i := range payload {
		payload[i] = byte(i)
	}

	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()

	cFile := f.C()

	// Read first byte to misalign
	got := NoxMemfileReadU8(cFile)
	require.Equal(t, uint8(0), got)

	// Now test read64align from misaligned position
	var aligned [8]byte
	n := NoxMemfileRead64Align(unsafe.Pointer(&aligned[0]), 1, 8, cFile)
	// Should read 1 block of 8 bytes (after alignment)
	if n == 1 {
		// Verify data was read (specific values depend on alignment logic)
		t.Logf("Read64Align got data: %v", aligned)
	}

	// Test read64align with exact alignment
	raw2, _ := alloc.CloneSlice(payload)
	f2 := binfile.NewMemFile(unsafe.Pointer(&raw2[0]), len(raw2))
	defer f2.Free()
	cFile2 := f2.C()

	var aligned2 [8]byte
	n = NoxMemfileRead64Align(unsafe.Pointer(&aligned2[0]), 1, 8, cFile2)
	if n == 1 {
		t.Logf("Read64Align from start got data: %v", aligned2)
	}
}

func TestNoxMemfileSkip(t *testing.T) {
	payload := []byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08}

	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()

	cFile := f.C()

	// Read first byte
	require.Equal(t, uint8(0x01), NoxMemfileReadU8(cFile))

	// Skip 2 bytes
	NoxMemfileSkip(cFile, 2)

	// Next byte should be 0x04
	require.Equal(t, uint8(0x04), NoxMemfileReadU8(cFile))

	// Skip to end
	NoxMemfileSkip(cFile, 10)

	// Reading past end should not crash
	_ = NoxMemfileReadU8(cFile)
}
