package noxscript

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNsAbort(t *testing.T) {
	// nsAbort is a simple function that returns 1
	result := nsAbort(nil)
	require.Equal(t, 1, result)
}

func TestGridUnpack(t *testing.T) {
	// Test gridUnpack function which unpacks a packed int32 into image.Point
	// The function does: image.Pt(int(packed>>16), int(uint16(packed)))

	// Test zero
	p := gridUnpack(0)
	require.Equal(t, image.Pt(0, 0), p)

	// Test simple values
	// Packed format: high 16 bits = x, low 16 bits = y
	packed := int32((10 << 16) | 20)
	p = gridUnpack(packed)
	require.Equal(t, image.Pt(10, 20), p)

	// Test max values
	packed = int32((0x7FFF << 16) | 0xFFFF)
	p = gridUnpack(packed)
	require.Equal(t, image.Pt(0x7FFF, 0xFFFF), p)

	// Test only x
	packed = int32(5 << 16)
	p = gridUnpack(packed)
	require.Equal(t, image.Pt(5, 0), p)

	// Test only y
	packed = int32(7)
	p = gridUnpack(packed)
	require.Equal(t, image.Pt(0, 7), p)

	// Test negative values (as signed int16)
	packed = int32((-1 << 16) | 0xFFFF)
	p = gridUnpack(packed)
	// -1 in int16 is 0xFFFF, so x = -1, y = 65535
	require.Equal(t, image.Pt(-1, 65535), p)
}
