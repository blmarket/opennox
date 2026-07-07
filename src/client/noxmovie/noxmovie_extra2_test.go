//go:build !server

package noxmovie

import (
	"testing"

	"github.com/youpy/go-wav"
)

func TestNewPlayerError2(t *testing.T) {
	// Test NewPlayer with non-existent file
	_, err := NewPlayer("/nonexistent/path/movie.vqa", nil, 0)
	if err == nil {
		t.Error("expected error for non-existent file")
	}
}

func TestMoviePlayerCloseNil(t *testing.T) {
	// Test Close on a player with stop channel already closed
	player := &MoviePlayer{
		stop: make(chan struct{}),
	}
	close(player.stop)
	// Should not panic
	player.Close()
}

func TestMoviePlayerSetAudioGain(t *testing.T) {
	player := &MoviePlayer{}
	player.SetAudioGain(0.5)
	if player.audioGain != 0.5 {
		t.Errorf("expected audioGain 0.5, got %f", player.audioGain)
	}
	player.SetAudioGain(1.0)
	if player.audioGain != 1.0 {
		t.Errorf("expected audioGain 1.0, got %f", player.audioGain)
	}
}

func TestConvertSampleToDataEmpty(t *testing.T) {
	// Test with empty samples
	samples := []wav.Sample{}
	out := convertSampleToData(samples, 0)
	if len(out) != 0 {
		t.Errorf("expected empty output, got length %d", len(out))
	}
}

func TestConvertSampleToDataMono(t *testing.T) {
	samples := []wav.Sample{
		{Values: [2]int{0x1234, 0x5678}},
	}
	out := convertSampleToData(samples, 1) // FormatMono16 = 1?
	// Just verify it doesn't panic and produces output
	if len(out) == 0 {
		t.Error("expected non-empty output")
	}
}
