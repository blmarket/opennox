package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerBalanceFreeAndTag(t *testing.T) {
	sb := &serverBalance{}
	sb.Free()
	require.Nil(t, sb.file)
	// Tag should default to Arena when no coop flag
	tag := sb.Tag()
	require.NotEqual(t, "", string(tag))
}

func TestServerBalanceFloatInd(t *testing.T) {
	sb := &serverBalance{}
	require.Equal(t, 0.0, sb.FloatInd("key", -1))
	require.Equal(t, 0.0, sb.FloatInd("key", 100))
}
