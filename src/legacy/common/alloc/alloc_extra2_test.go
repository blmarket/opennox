package alloc

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllocExtra(t *testing.T) {
	p, _ := Malloc(10)
	require.NotNil(t, p)
	p2 := Realloc(p, 20)
	require.NotNil(t, p2)
	FreePtr(p2)

	c := NewDynamicClassT("test", int(0), 10)
	require.NotNil(t, c)
	obj := c.NewObject()
	require.NotNil(t, obj)
	c.FreeObjectFirst(obj)
	c.Free()
}
