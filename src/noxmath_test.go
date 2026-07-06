package opennox

import (
	"math"
	"testing"
)

func TestNoxXxxMathRoundDirI16(t *testing.T) {
	tests := []struct {
		v    int16
		want uint16
	}{
		{0, 0},
		{1, 1},
		{-1, 255},
		{127, 127},
		{-128, 128},
	}
	for _, tt := range tests {
		got := nox_xxx_math_roundDirI16(tt.v)
		if got != tt.want {
			t.Errorf("nox_xxx_math_roundDirI16(%d) = %d, want %d", tt.v, got, tt.want)
		}
	}
}

func TestFabs(t *testing.T) {
	tests := []struct {
		v    float32
		want float32
	}{
		{0, 0},
		{1.5, 1.5},
		{-1.5, 1.5},
		{-0, 0},
	}
	for _, tt := range tests {
		got := fabs(tt.v)
		if got != tt.want {
			t.Errorf("fabs(%f) = %f, want %f", tt.v, got, tt.want)
		}
		// Also verify it's non-negative
		if got < 0 {
			t.Errorf("fabs(%f) = %f, want non-negative", tt.v, got)
		}
	}
}

func TestFabsSpecial(t *testing.T) {
	// Test with special float values
	if !math.IsNaN(float64(fabs(float32(math.NaN())))) {
		t.Error("fabs(NaN) should be NaN")
	}
	if fabs(float32(math.Inf(1))) != float32(math.Inf(1)) {
		t.Error("fabs(+Inf) should be +Inf")
	}
	if fabs(float32(math.Inf(-1))) != float32(math.Inf(1)) {
		t.Error("fabs(-Inf) should be +Inf")
	}
}

func TestSincosTable16(t *testing.T) {
	// Verify the table has 256 entries (0-255)
	if len(sincosTable16) != 256 {
		t.Errorf("sincosTable16 length = %d, want 256", len(sincosTable16))
	}
	// Check a few known values
	if sincosTable16[0].X != 16 || sincosTable16[0].Y != 0 {
		t.Errorf("sincosTable16[0] = %v, want {16, 0}", sincosTable16[0])
	}
}
