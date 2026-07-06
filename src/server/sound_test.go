package server

import (
	"testing"
)

func TestSub532F70(t *testing.T) {
	tests := []struct {
		v    int32
		want int32
	}{
		{0, 1},
		{2, 1},
		{3, 1},
		{10, 1},
		{11, 1},
		{1, 0},
		{4, 0},
		{5, 0},
		{-1, 0},
	}
	for _, tt := range tests {
		got := sub_532F70(tt.v)
		if got != tt.want {
			t.Errorf("sub_532F70(%d) = %d, want %d", tt.v, got, tt.want)
		}
	}
}

func TestSub532FB0(t *testing.T) {
	tests := []struct {
		v    uint16
		want int32
	}{
		{8, 1},
		{32, 1},
		{64, 1},
		{0, 0},
		{1, 0},
		{16, 0},
		{100, 0},
	}
	for _, tt := range tests {
		got := sub_532FB0(tt.v)
		if got != tt.want {
			t.Errorf("sub_532FB0(%d) = %d, want %d", tt.v, got, tt.want)
		}
	}
}
