package client

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/stretchr/testify/require"
)

func TestClient_ReadWalls_Extra(t *testing.T) {
	// Test ReadWalls with nil client should panic or return error
	// We just verify the function exists and handles nil gracefully
	require.Panics(t, func() {
		var c *Client
		var f *binfile.MemFile
		_ = c.ReadWalls(f)
	})
}

func TestClient_ReadWalls_NilFile(t *testing.T) {
	// Test ReadWalls with nil file
	c := &Client{}
	require.Panics(t, func() {
		_ = c.ReadWalls(nil)
	})
}
