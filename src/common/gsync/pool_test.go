package gsync

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPool(t *testing.T) {
	var p Pool[int]

	// Get from empty pool should create new
	v := p.Get()
	require.NotNil(t, v)
	require.Equal(t, 0, *v)

	// Put and get
	*v = 42
	p.Put(v)

	v2 := p.Get()
	require.NotNil(t, v2)
	// May reuse the same object or create new, both are valid
	_ = v2

	// Put nil should not panic
	p.Put(nil)

	// Test with custom New function
	var p2 Pool[string]
	called := false
	p2.New = func() *string {
		called = true
		s := "custom"
		return &s
	}

	v3 := p2.Get()
	require.True(t, called)
	require.NotNil(t, v3)
	require.Equal(t, "custom", *v3)
}

func TestPoolPutNil(t *testing.T) {
	var p Pool[int]
	// Should not panic
	p.Put(nil)
}
