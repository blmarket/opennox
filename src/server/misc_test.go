package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFirstWord(t *testing.T) {
	require.Equal(t, "hello", firstWord("hello world"))
	require.Equal(t, "test", firstWord("test\t\n"))
	require.Equal(t, "single", firstWord("single"))
}

func TestAbs(t *testing.T) {
	require.Equal(t, 5, abs(-5))
	require.Equal(t, 5, abs(5))
	require.Equal(t, 3.5, abs(-3.5))
	require.Equal(t, uint(3), abs(uint(3)))
}
