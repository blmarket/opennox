package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProtectBytes(t *testing.T) {
	require.Equal(t, uint32(0), protectBytes([]byte{}))
	require.Equal(t, uint32(0x01020304), protectBytes([]byte{0x04, 0x03, 0x02, 0x01}))
	data := []byte{1, 0, 0, 0, 2, 0, 0, 0}
	require.Equal(t, uint32(3), protectBytes(data))
	require.Equal(t, uint32(1), protectBytes([]byte{1, 0, 0, 0, 1}))
}
