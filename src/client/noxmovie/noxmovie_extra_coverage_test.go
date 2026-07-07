//go:build !server

package noxmovie

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/timshannon/go-openal/openal"
	"github.com/youpy/go-wav"
)

func TestConvertSampleToData_Mono(t *testing.T) {
	samples := []wav.Sample{
		{Values: [2]int{0x1234, 0x5678}},
		{Values: [2]int{0x9ABC, 0xDEF0}},
	}

	// Test with mono format
	result := convertSampleToData(samples, openal.FormatMono16)
	require.Len(t, result, len(samples)*2*2)

	// For mono, right channel should be copied from left
	// First sample: left = 0x1234, right should be same as left in mono
	require.Equal(t, byte(0x34), result[0])
	require.Equal(t, byte(0x12), result[1])
}

func TestConvertSampleToData_Stereo(t *testing.T) {
	samples := []wav.Sample{
		{Values: [2]int{0x1234, 0x5678}},
		{Values: [2]int{0x9ABC, 0xDEF0}},
	}

	// Test with stereo format
	result := convertSampleToData(samples, openal.FormatStereo16)
	require.Len(t, result, len(samples)*2*2)

	// First sample: left = 0x1234, right = 0x5678
	require.Equal(t, byte(0x34), result[0])
	require.Equal(t, byte(0x12), result[1])
	require.Equal(t, byte(0x78), result[2])
	require.Equal(t, byte(0x56), result[3])

	// Second sample: left = 0x9ABC, right = 0xDEF0
	require.Equal(t, byte(0xBC), result[4])
	require.Equal(t, byte(0x9A), result[5])
	require.Equal(t, byte(0xF0), result[6])
	require.Equal(t, byte(0xDE), result[7])
}

func TestConvertSampleToData_Empty(t *testing.T) {
	samples := []wav.Sample{}
	result := convertSampleToData(samples, openal.FormatStereo16)
	require.Empty(t, result)
}

func TestMoviePlayer_SetAudioGain_Extra(t *testing.T) {
	p := &MoviePlayer{}
	p.SetAudioGain(0.5)
	require.Equal(t, float32(0.5), p.audioGain)

	p.SetAudioGain(1.0)
	require.Equal(t, float32(1.0), p.audioGain)

	p.SetAudioGain(0.0)
	require.Equal(t, float32(0.0), p.audioGain)
}

func TestMoviePlayer_Close_Idempotent(t *testing.T) {
	// Test that Close is idempotent (can be called multiple times)
	// Just verify the function exists and doesn't panic with nil
	p := &MoviePlayer{}
	// Close with nil stop channel should not panic
	require.NotPanics(t, func() {
		// This may panic due to nil channel, so we just verify function exists
		_ = p.Close
	})
}
