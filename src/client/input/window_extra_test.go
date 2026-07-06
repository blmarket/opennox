package input

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/stretchr/testify/require"
)

func TestClampf(t *testing.T) {
	r := image.Rect(0, 0, 100, 100)

	// Test within bounds
	p := types.Pointf{X: 50, Y: 50}
	result := clampf(r, p)
	require.Equal(t, float32(50), result.X)
	require.Equal(t, float32(50), result.Y)

	// Test below min
	p = types.Pointf{X: -10, Y: -20}
	result = clampf(r, p)
	require.Equal(t, float32(0), result.X)
	require.Equal(t, float32(0), result.Y)

	// Test above max
	p = types.Pointf{X: 150, Y: 200}
	result = clampf(r, p)
	require.Equal(t, float32(100), result.X)
	require.Equal(t, float32(100), result.Y)

	// Test at max boundary
	p = types.Pointf{X: 100, Y: 100}
	result = clampf(r, p)
	require.Equal(t, float32(100), result.X)
	require.Equal(t, float32(100), result.Y)
}

func TestClamp(t *testing.T) {
	r := image.Rect(10, 10, 100, 100)

	// Test within bounds
	p := image.Pt(50, 50)
	result := clamp(r, p)
	require.Equal(t, 50, result.X)
	require.Equal(t, 50, result.Y)

	// Test below min
	p = image.Pt(5, 5)
	result = clamp(r, p)
	require.Equal(t, 10, result.X)
	require.Equal(t, 10, result.Y)

	// Test above max
	p = image.Pt(150, 200)
	result = clamp(r, p)
	require.Equal(t, 100, result.X)
	require.Equal(t, 100, result.Y)
}
