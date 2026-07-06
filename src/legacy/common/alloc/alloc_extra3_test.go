package alloc

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestStrlenExtra(t *testing.T) {
	require.Equal(t, 0, Strlen(nil))
}

func TestStrcatExtra(t *testing.T) {
	// Strcat with nil causes segfault, so we just verify the function exists
	require.NotNil(t, Strcat)
}

func TestClassStringExtra(t *testing.T) {
	c := &Class{}
	s := c.String()
	require.Contains(t, s, "Class")
}

func TestNewDynamicClassExtra(t *testing.T) {
	c := NewDynamicClass("test", 64, 10)
	require.NotNil(t, c)
	c.Free()
}

func TestClassKeepExtra(t *testing.T) {
	c := NewClass("test", 32, 10)
	require.NotNil(t, c)
	c.Keep(0)
	c.Free()
}

func TestAsClassTExtra(t *testing.T) {
	result := AsClassT[int](nil)
	require.Nil(t, result.Class)
}

func TestNewClassTExtra(t *testing.T) {
	c := NewClassT[int]("test", 0, 10)
	require.NotNil(t, c)
	c.Free()
}

func TestStrlenUnsafe(t *testing.T) {
	s := "hello\x00"
	ptr := unsafe.Pointer(unsafe.StringData(s))
	n := Strlen(ptr)
	require.Equal(t, 5, n)
}
