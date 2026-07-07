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

func TestMoviePlayer_SetAudioGain(t *testing.T) {
	p := &MoviePlayer{}
	p.SetAudioGain(0.5)
	require.Equal(t, float32(0.5), p.audioGain)

	p.SetAudioGain(1.0)
	require.Equal(t, float32(1.0), p.audioGain)

	p.SetAudioGain(0.0)
	require.Equal(t, float32(0.0), p.audioGain)
}
