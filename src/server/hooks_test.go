package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTickHooks(t *testing.T) {
	var s Server
	called := 0
	s.TickHook(func() { called++ })
	s.TickCallback(func() { called++ })
	s.RunTickHooks()
	require.Equal(t, 2, called)
	s.RunTickHooks()
	require.Equal(t, 3, called) // only persistent hook runs again
}
