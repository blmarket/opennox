package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddMult(t *testing.T) {
	var s BlobOffSum
	s.AddMult(BlobOffMult{Static: 10})
	require.Equal(t, 1, len(s.Mult))
	require.Equal(t, uintptr(10), s.Mult[0].Static)

	// Adding same static merges
	s.AddMult(BlobOffMult{Static: 10, Sum: BlobOffSum{Raw: []string{"i"}}})
	require.Equal(t, 1, len(s.Mult))

	// Different static adds new
	s.AddMult(BlobOffMult{Static: 20})
	require.Equal(t, 2, len(s.Mult))
}
