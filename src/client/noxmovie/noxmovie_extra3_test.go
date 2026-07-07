package noxmovie

import (
	"testing"

	"github.com/timshannon/go-openal/openal"
	"github.com/youpy/go-wav"
)

func TestConvertSampleToData_Extra2(t *testing.T) {
	// Test empty samples
	out := convertSampleToData([]wav.Sample{}, openal.FormatStereo16)
	if len(out) != 0 {
		t.Errorf("Empty samples should return empty output, got len %d", len(out))
	}

	// Test single sample stereo
	samples := []wav.Sample{
		{Values: [2]int{0x0001, 0x0002}},
	}
	out = convertSampleToData(samples, openal.FormatStereo16)
	if len(out) != 4 {
		t.Fatalf("Single stereo sample should return 4 bytes, got %d", len(out))
	}
	if out[0] != 0x01 || out[1] != 0x00 {
		t.Errorf("Unexpected left channel bytes: %v", out[:2])
	}

	// Test mono format duplicates left channel
	outMono := convertSampleToData(samples, openal.FormatMono16)
	if len(outMono) != 4 {
		t.Fatalf("Single mono sample should return 4 bytes, got %d", len(outMono))
	}
	if outMono[0] != outMono[2] || outMono[1] != outMono[3] {
		t.Error("Mono format should duplicate left channel")
	}

	// Test with max values
	samples = []wav.Sample{
		{Values: [2]int{0x7FFF, 0x7FFF}},
		{Values: [2]int{-0x8000, -0x8000}},
	}
	out = convertSampleToData(samples, openal.FormatStereo16)
	if len(out) != 8 {
		t.Errorf("Two samples should return 8 bytes, got %d", len(out))
	}
}

func TestMoviePlayer_SetAudioGain_Extra2(t *testing.T) {
	p := &MoviePlayer{}
	gains := []float32{-1.0, 0.0, 0.5, 1.0, 2.0}
	for _, g := range gains {
		p.SetAudioGain(g)
	}
	var nilPlayer *MoviePlayer
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Close on nil panicked as expected: %v", r)
			}
		}()
		nilPlayer.Close()
	}()
}
