package serial

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	s := Generate()
	require.Len(t, s, 22)
	for _, c := range s {
		require.True(t, c >= '0' && c <= '9', "serial should contain only digits")
	}
}

func TestSerial(t *testing.T) {
	// Initially empty
	s, ok := Serial()
	require.False(t, ok)
	require.Empty(t, s)

	// Set valid serial
	code := "1234567890123456789012"
	SetSerial(code)
	s, ok = Serial()
	require.True(t, ok)
	require.Equal(t, code, s)

	// Set empty (allowed)
	SetSerial("")
	s, ok = Serial()
	require.False(t, ok)
	require.Empty(t, s)

	// Invalid length should panic
	require.Panics(t, func() {
		SetSerial("short")
	})
}
