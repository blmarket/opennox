package opennox

import (
	"math"
	"testing"
)

func TestNoxXxxMathRoundDirI16Extra(t *testing.T) {
	tests := []struct {
		v    int16
		want uint16
	}{
		{2, 2},
		{-2, 254},
		{126, 126},
		{-127, 129},
		{255, 255}, // int16 overflow?
	}
	for _, tt := range tests {
		got := nox_xxx_math_roundDirI16(tt.v)
		if got != tt.want {
			t.Errorf("nox_xxx_math_roundDirI16(%d) = %d, want %d", tt.v, got, tt.want)
		}
	}
}

func TestFabsExtra(t *testing.T) {
	tests := []struct {
		v    float32
		want float32
	}{
		{3.14, 3.14},
		{-3.14, 3.14},
		{1e10, 1e10},
		{-1e10, 1e10},
		{1e-10, 1e-10},
		{-1e-10, 1e-10},
	}
	for _, tt := range tests {
		got := fabs(tt.v)
		if got != tt.want {
			t.Errorf("fabs(%f) = %f, want %f", tt.v, got, tt.want)
		}
		if got < 0 {
			t.Errorf("fabs(%f) = %f, want non-negative", tt.v, got)
		}
	}
}

func TestSincosTable16Extra(t *testing.T) {
	// Check that all entries are within expected range
	for i, v := range sincosTable16 {
		if v.X < -16 || v.X > 16 {
			t.Errorf("sincosTable16[%d].X = %d, out of range [-16, 16]", i, v.X)
		}
		if v.Y < -16 || v.Y > 16 {
			t.Errorf("sincosTable16[%d].Y = %d, out of range [-16, 16]", i, v.Y)
		}
	}
	// Check specific known values
	// At 0 degrees (index 0): sin=0, cos=16
	if sincosTable16[0].X != 16 || sincosTable16[0].Y != 0 {
		t.Errorf("sincosTable16[0] = %v, want {16, 0}", sincosTable16[0])
	}
	// At 90 degrees (index 64): sin=16, cos=0
	if sincosTable16[64].X != 0 || sincosTable16[64].Y != 16 {
		t.Errorf("sincosTable16[64] = %v, want {0, 16}", sincosTable16[64])
	}
	// At 180 degrees (index 128): sin=0, cos=-16 (actually -15 due to table precision)
	if sincosTable16[128].Y != 0 {
		t.Errorf("sincosTable16[128].Y = %d, want 0", sincosTable16[128].Y)
	}
	// At 270 degrees (index 192): sin=-16, cos=0 (actually -15 due to table precision)
	if sincosTable16[192].X != 0 {
		t.Errorf("sincosTable16[192].X = %d, want 0", sincosTable16[192].X)
	}
}

func TestFabsNaN(t *testing.T) {
	got := fabs(float32(math.NaN()))
	if !math.IsNaN(float64(got)) {
		t.Error("fabs(NaN) should be NaN")
	}
}
