package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBlobOffSumAddMultMerge(t *testing.T) {
	var s BlobOffSum
	// First add
	s.AddMult(BlobOffMult{Static: 2, Sum: BlobOffSum{Raw: []string{"i"}}})
	require.Equal(t, 1, len(s.Mult))
	require.Equal(t, uintptr(2), s.Mult[0].Static)

	// Second add with same Static should merge instead of append
	s.AddMult(BlobOffMult{Static: 2, Sum: BlobOffSum{Raw: []string{"j"}}})
	require.Equal(t, 1, len(s.Mult), "AddMult with same Static should merge, not append")
	require.Equal(t, uintptr(2), s.Mult[0].Static)
}
