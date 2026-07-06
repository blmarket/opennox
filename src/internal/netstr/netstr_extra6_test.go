package netstr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewStreamsExtra2(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	require.NotNil(t, s)
	idx, _ := s.getFreeIndex()
	require.GreaterOrEqual(t, idx, 0)
	idx2 := s.getFreeNetStruct2Ind()
	require.GreaterOrEqual(t, idx2, 0)
}

func TestConnMethodsExtra2(t *testing.T) {
	c := &Conn{}
	require.NotEmpty(t, c.String())
	require.True(t, c.IsHost())
	_ = c.Player()
	c.reset()
	c.resetReliable()
}
