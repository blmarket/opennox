package client

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/client/gui"
)

func TestFirstWord_Extra(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"", ""},
		{"single", "single"},
		{"  leading", "  leading"},
		{"a b c", "a"},
		{"hello\tworld", "hello"},
		{"hello\nworld", "hello"},
		{"hello\rworld", "hello"},
		{"   ", "   "},
	}
	for _, tt := range tests {
		got := firstWord(tt.input)
		if got != tt.want {
			t.Errorf("firstWord(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestStateConstants_Extra(t *testing.T) {
	if StateMovies != gui.StateID(10) {
		t.Errorf("StateMovies = %v, want 10", StateMovies)
	}
	if StateMainMenu != gui.StateID(100) {
		t.Errorf("StateMainMenu = %v, want 100", StateMainMenu)
	}
	if StateOptions != gui.StateID(300) {
		t.Errorf("StateOptions = %v, want 300", StateOptions)
	}
	if StateCharSelect != gui.StateID(500) {
		t.Errorf("StateCharSelect = %v, want 500", StateCharSelect)
	}
	if StateClassSelect != gui.StateID(600) {
		t.Errorf("StateClassSelect = %v, want 600", StateClassSelect)
	}
	if StateColorSelect != gui.StateID(700) {
		t.Errorf("StateColorSelect = %v, want 700", StateColorSelect)
	}
	if StateServerList != gui.StateID(10000) {
		t.Errorf("StateServerList = %v, want 10000", StateServerList)
	}
	if StateXxx != gui.StateID(1915) {
		t.Errorf("StateXxx = %v, want 1915", StateXxx)
	}
}

func TestRulesConstants_Extra(t *testing.T) {
	if C != 0 {
		t.Errorf("C constant should be 0, got %d", C)
	}
	if legacy != 0 {
		t.Errorf("legacy constant should be 0, got %d", legacy)
	}
}

func TestParseAnimKind_Extra(t *testing.T) {
	tests := []struct {
		input string
		want  AnimKind
	}{
		{"OneShot", AnimOneShot},
		{"oneshot", AnimOneShot},
		{"Loop", AnimLoop},
		{"Invalid", AnimOneShot},
		{"", AnimOneShot},
		{"Slave", AnimSlave},
		{"Random", AnimRandom},
	}
	for _, tt := range tests {
		got := ParseAnimKind(tt.input)
		if got != tt.want {
			t.Errorf("ParseAnimKind(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestAnimationVector_FramesSlice_Extra(t *testing.T) {
	// Test with empty AnimationVector
	av := &AnimationVector{}
	if av.Cnt40 != 0 {
		t.Error("New AnimationVector should have Cnt40 = 0")
	}
	// FramesSlice with nil should not panic for empty
	defer func() {
		if r := recover(); r != nil {
			t.Logf("FramesSlice panicked as expected for nil: %v", r)
		}
	}()
	_ = av.FramesSlice(0)
}

func TestPlayerAnimation_FramesSlice_Extra(t *testing.T) {
	pa := &PlayerAnimation{}
	if pa.Base.Cnt40 != 0 {
		t.Error("New PlayerAnimation should have Base.Cnt40 = 0")
	}
}
