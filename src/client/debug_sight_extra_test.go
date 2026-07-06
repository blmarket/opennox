package client

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClient_DebugSightAdd_Extra(t *testing.T) {
	// Test DebugSightAdd with nil client should panic
	require.Panics(t, func() {
		var c *Client
		c.DebugSightAdd()
	})
}

func TestClient_DrawDebugSight_Extra(t *testing.T) {
	// Test DrawDebugSight with nil client should panic
	require.Panics(t, func() {
		var c *Client
		c.DrawDebugSight(nil)
	})
}
