package gsync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapExtra(t *testing.T) {
	var m Map[int, string]

	// Test Load on empty map
	v, ok := m.Load(1)
	require.False(t, ok)
	require.Equal(t, "", v)

	// Test Store and Load
	m.Store(1, "one")
	v, ok = m.Load(1)
	require.True(t, ok)
	require.Equal(t, "one", v)

	// Test Delete
	m.Delete(1)
	v, ok = m.Load(1)
	require.False(t, ok)
	require.Equal(t, "", v)
}

func TestPoolExtra(t *testing.T) {
	var p Pool[[]byte]

	// Test Get on empty pool
	v := p.Get()
	require.NotNil(t, v)

	// Test Put and Get
	data := []byte{1, 2, 3}
	p.Put(&data)
	v = p.Get()
	require.NotNil(t, v)

	// Test Get again - should be nil or reused
	v2 := p.Get()
	_ = v2
}
