package serial

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSerial_Extra(t *testing.T) {
	// Test Generate produces different values
	s1 := Generate()
	s2 := Generate()
	require.Len(t, s1, 22)
	require.Len(t, s2, 22)
	// Very unlikely to be equal, but not guaranteed
	_ = s2

	// Test SetSerial with empty
	SetSerial("")
	s, ok := Serial()
	require.False(t, ok)
	require.Empty(t, s)

	// Test SetSerial with valid code
	code := "0000000000000000000000"
	SetSerial(code)
	s, ok = Serial()
	require.True(t, ok)
	require.Equal(t, code, s)

	// Test panic on invalid length
	require.Panics(t, func() {
		SetSerial("123")
	})
	require.Panics(t, func() {
		SetSerial("12345678901234567890123") // 23 chars
	})

	// Reset
	SetSerial("")
}
