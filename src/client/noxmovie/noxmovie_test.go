package noxmovie

import (
	"testing"

	"github.com/timshannon/go-openal/openal"
	"github.com/youpy/go-wav"
)

func TestConvertSampleToData(t *testing.T) {
	samples := []wav.Sample{
		{Values: [2]int{0x1234, 0x5678}},
		{Values: [2]int{0x9ABC, 0xDEF0}},
	}
	out := convertSampleToData(samples, openal.FormatStereo16)
	if len(out) != len(samples)*4 {
		t.Fatalf("unexpected length %d", len(out))
	}
	// first sample left 0x34 0x12, right 0x78 0x56
	if out[0] != 0x34 || out[1] != 0x12 || out[2] != 0x78 || out[3] != 0x56 {
		t.Fatalf("unexpected output %v", out[:4])
	}
	out2 := convertSampleToData(samples, openal.FormatMono16)
	if len(out2) != len(samples)*4 {
		t.Fatalf("unexpected length")
	}
	// mono should duplicate left channel
	if out2[2] != out2[0] || out2[3] != out2[1] {
		t.Fatalf("expected mono duplication")
	}
}
