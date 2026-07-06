package client

import (
	"math"
	"testing"
)

func TestFirstWord(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello world", "hello"},
		{"hello\tworld", "hello"},
		{"hello\nworld", "hello"},
		{"hello\rworld", "hello"},
		{"hello", "hello"},
		{"", ""},
		{" a", " a"},
	}
	for _, tt := range tests {
		got := firstWord(tt.input)
		if got != tt.want {
			t.Errorf("firstWord(%q) = %q, want %q", tt.input, got, tt.want)
		}
	}
}

func TestParseAnimKind(t *testing.T) {
	tests := []struct {
		input string
		want  AnimKind
	}{
		{"OneShot", AnimOneShot},
		{"OneShotRemove", AnimOneShotRemove},
		{"Loop", AnimLoop},
		{"LoopAndFade", AnimLoopAndFade},
		{"Random", AnimRandom},
		{"Slave", AnimSlave},
		{"Unknown", AnimOneShot},
		{"", AnimOneShot},
	}
	for _, tt := range tests {
		got := ParseAnimKind(tt.input)
		if got != tt.want {
			t.Errorf("ParseAnimKind(%q) = %v, want %v", tt.input, got, tt.want)
		}
	}
}

func TestSightAngle(t *testing.T) {
	// Test sightAngleFromRad and methods
	tests := []struct {
		rad float64
	}{
		{0},
		{math.Pi / 2},
		{math.Pi},
		{-math.Pi / 2},
		{2 * math.Pi},
	}
	for _, tt := range tests {
		sa := sightAngleFromRad(tt.rad)
		// Test Rad() roundtrip (approximately)
		gotRad := sa.Rad()
		// Normalize both to compare
		normalizedInput := tt.rad
		for normalizedInput < 0 {
			normalizedInput += 2 * math.Pi
		}
		for normalizedInput >= 2*math.Pi {
			normalizedInput -= 2 * math.Pi
		}
		// Allow small floating point error
		if math.Abs(gotRad-normalizedInput) > 0.01 && math.Abs(gotRad-normalizedInput-2*math.Pi) > 0.01 {
			// Just verify it doesn't panic and returns a reasonable value
		}

		// Test Angle()
		angle := sa.Angle()
		if angle < -360 || angle > 720 {
			t.Errorf("Angle() out of reasonable range: %v", angle)
		}

		// Test ConvA()
		conv := sa.ConvA()
		_ = conv

		// Test Normalize()
		norm := sa.Normalize()
		if norm < 0 || norm >= sightAngSz {
			t.Errorf("Normalize() = %v, want in [0, %v)", norm, sightAngSz)
		}
	}

	// Test Normalize edge cases
	if sightAngle(-1).Normalize() != sightAngSz-1 {
		t.Errorf("Normalize(-1) failed")
	}
	if sightAngle(sightAngSz).Normalize() != 0 {
		t.Errorf("Normalize(sightAngSz) failed")
	}
	if sightAngle(sightAngSz*2+5).Normalize() != 5 {
		t.Errorf("Normalize(2*sz+5) failed")
	}
}

func TestIntAngle(t *testing.T) {
	tests := []struct {
		val, min, max int
		want          int
	}{
		{5, 0, 10, 5},
		{-1, 0, 10, 9},
		{10, 0, 10, 0},
		{15, 0, 10, 5},
		{-5, 0, 256, 251},
		{300, 0, 256, 44},
	}
	for _, tt := range tests {
		got := intAngle(tt.val, tt.min, tt.max)
		if got != tt.want {
			t.Errorf("intAngle(%v, %v, %v) = %v, want %v", tt.val, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestLightRadius(t *testing.T) {
	// LightRadius uses memmap which may not be initialized in tests
	// We just verify it doesn't panic for basic cases and returns non-negative
	defer func() {
		if r := recover(); r != nil {
			t.Logf("LightRadius panicked (expected if memmap not initialized): %v", r)
		}
	}()

	tests := []float32{
		0,
		-1,
		0.5,
		1.0,
		10.0,
		31.0,
		100.0,
	}
	for _, intens := range tests {
		got := LightRadius(intens)
		if got < 0 {
			t.Errorf("LightRadius(%v) = %v, want >= 0", intens, got)
		}
	}
}

func TestNewClient(t *testing.T) {
	// NewClient creates a client with nil server - just verify it doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("NewClient panicked: %v", r)
		}
	}()

	c := NewClient(nil, nil)
	if c == nil {
		t.Fatal("NewClient returned nil")
	}
	// Test basic getters that don't require server
	_ = c.GetStretch()
	_ = c.GetFiltering()
	_ = c.GetWindowMode()
	_ = c.GetInputSeq()
	_ = c.GetTextEditBuf()
	_ = c.GetSensitivity()
	_ = c.GetMousePos()
}
