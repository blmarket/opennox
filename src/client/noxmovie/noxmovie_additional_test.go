package noxmovie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPlayerError(t *testing.T) {
	// Test NewPlayer with non-existent file - should return error
	// NewPlayer opens the file first, so it returns error before using seat/audio
	p, err := NewPlayer("/non/existent/file.nox", nil, 0)
	require.Error(t, err)
	require.Nil(t, p)
}
