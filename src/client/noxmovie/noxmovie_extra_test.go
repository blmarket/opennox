package noxmovie

import (
	"testing"
)

func TestMoviePlayer(t *testing.T) {
	p := &MoviePlayer{}
	p.SetAudioGain(1.0)
	// Don't call Close on uninitialized player to avoid panic
}
