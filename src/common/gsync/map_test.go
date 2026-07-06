package gsync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	var m Map[string, int]

	// Load non-existent
	v, ok := m.Load("key1")
	require.False(t, ok)
	require.Equal(t, 0, v)

	// Store and load
	m.Store("key1", 42)
	v, ok = m.Load("key1")
	require.True(t, ok)
	require.Equal(t, 42, v)

	// Overwrite
	m.Store("key1", 100)
	v, ok = m.Load("key1")
	require.True(t, ok)
	require.Equal(t, 100, v)

	// Delete
	m.Delete("key1")
	v, ok = m.Load("key1")
	require.False(t, ok)
	require.Equal(t, 0, v)

	// Multiple keys
	m.Store("a", 1)
	m.Store("b", 2)
	m.Store("c", 3)

	v, ok = m.Load("a")
	require.True(t, ok)
	require.Equal(t, 1, v)

	v, ok = m.Load("b")
	require.True(t, ok)
	require.Equal(t, 2, v)

	m.Delete("b")
	_, ok = m.Load("b")
	require.False(t, ok)
}
