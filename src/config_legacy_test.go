package opennox

import (
	"testing"
)

func TestConfigParseVolume(t *testing.T) {
	tests := []struct {
		val     string
		ind     VolumeControl
		wantErr bool
	}{
		{"0", VolumeFX, false},
		{"100", VolumeFX, false},
		{"16384", VolumeFX, false}, // VolumeMax = 0x4000 = 16384
		{"20000", VolumeFX, false}, // Should clamp to max
		{"-1", VolumeFX, true},     // ParseUint fails on negative
		{"invalid", VolumeFX, true},
		{"", VolumeFX, true},
	}

	for _, tt := range tests {
		err := configParseVolume(tt.val, tt.ind)
		if (err != nil) != tt.wantErr {
			t.Errorf("configParseVolume(%q, %v) error = %v, wantErr %v", tt.val, tt.ind, err, tt.wantErr)
		}
	}
}

func TestConfigParseVolumeClamping(t *testing.T) {
	// Test that values above max are clamped
	err := configParseVolume("99999", VolumeFX)
	if err != nil {
		t.Errorf("configParseVolume with large value should not error, got %v", err)
	}
	val := configGetVolume(VolumeFX)
	if val != VolumeMax {
		t.Errorf("configParseVolume with large value should clamp to %d, got %d", VolumeMax, val)
	}
}

func TestConfigSetGetVolume(t *testing.T) {
	tests := []struct {
		val int
		ind VolumeControl
	}{
		{0, VolumeFX},
		{100, VolumeFX},
		{VolumeMax, VolumeFX},
		{0, VolumeDialog},
		{5000, VolumeDialog},
		{0, VolumeMusic},
		{8000, VolumeMusic},
	}

	for _, tt := range tests {
		configSetVolume(tt.val, tt.ind)
		got := configGetVolume(tt.ind)
		if got != tt.val {
			t.Errorf("configGetVolume after configSetVolume(%d, %v) = %d, want %d", tt.val, tt.ind, got, tt.val)
		}
	}
}

func TestConfigGetVolumeInvalid(t *testing.T) {
	// Invalid volume control should return VolumeMax
	val := configGetVolume(VolumeControl(999))
	if val != VolumeMax {
		t.Errorf("configGetVolume with invalid ind should return %d, got %d", VolumeMax, val)
	}
}

func TestConfigGetVolumeGain(t *testing.T) {
	configSetVolume(0, VolumeFX)
	gain := configGetVolumeGain(VolumeFX)
	if gain != 0.0 {
		t.Errorf("configGetVolumeGain with 0 volume should be 0.0, got %f", gain)
	}

	configSetVolume(VolumeMax, VolumeFX)
	gain = configGetVolumeGain(VolumeFX)
	if gain != 1.0 {
		t.Errorf("configGetVolumeGain with max volume should be 1.0, got %f", gain)
	}

	configSetVolume(VolumeMax/2, VolumeFX)
	gain = configGetVolumeGain(VolumeFX)
	expected := float32(0.5)
	if gain != expected {
		t.Errorf("configGetVolumeGain with half volume should be %f, got %f", expected, gain)
	}
}

func TestVolumeControlConstants(t *testing.T) {
	if VolumeMax != 0x4000 {
		t.Errorf("VolumeMax should be 0x4000, got %d", VolumeMax)
	}
	if VolumeFX != 0 {
		t.Errorf("VolumeFX should be 0, got %d", VolumeFX)
	}
	if VolumeDialog != 1 {
		t.Errorf("VolumeDialog should be 1, got %d", VolumeDialog)
	}
	if VolumeMusic != 2 {
		t.Errorf("VolumeMusic should be 2, got %d", VolumeMusic)
	}
}
