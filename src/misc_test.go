package opennox

import "testing"

func TestAbs(t *testing.T) {
	tests := []struct {
		v    int
		want int
	}{
		{0, 0},
		{5, 5},
		{-5, 5},
		{-1, 1},
		{123, 123},
		{-123, 123},
	}
	for _, tt := range tests {
		got := abs(tt.v)
		if got != tt.want {
			t.Errorf("abs(%d) = %d, want %d", tt.v, got, tt.want)
		}
	}
}

func TestClamp(t *testing.T) {
	tests := []struct {
		v, min, max, want int
	}{
		{5, 0, 10, 5},
		{-1, 0, 10, 0},
		{20, 0, 10, 10},
		{0, 0, 10, 0},
		{10, 0, 10, 10},
		{5, 5, 5, 5},
	}
	for _, tt := range tests {
		got := clamp(tt.v, tt.min, tt.max)
		if got != tt.want {
			t.Errorf("clamp(%d, %d, %d) = %d, want %d", tt.v, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestBool2int(t *testing.T) {
	if bool2int(true) != 1 {
		t.Error("bool2int(true) should be 1")
	}
	if bool2int(false) != 0 {
		t.Error("bool2int(false) should be 0")
	}
}

func TestRotl(t *testing.T) {
	tests := []struct {
		v    uint32
		n    int
		want uint32
	}{
		{0x12345678, 0, 0x12345678},
		{0x12345678, 4, 0x23456781},
		{0x12345678, 8, 0x34567812},
		{0xFFFFFFFF, 1, 0xFFFFFFFF},
		{0x00000001, 1, 0x00000002},
		{0x80000000, 1, 0x00000001},
	}
	for _, tt := range tests {
		got := rotl(tt.v, tt.n)
		if got != tt.want {
			t.Errorf("rotl(0x%08x, %d) = 0x%08x, want 0x%08x", tt.v, tt.n, got, tt.want)
		}
	}
}

func TestRotl16(t *testing.T) {
	tests := []struct {
		v    uint16
		n    int
		want uint16
	}{
		{0x1234, 0, 0x1234},
		{0x1234, 4, 0x2341},
		{0x1234, 8, 0x3412},
		{0xFFFF, 1, 0xFFFF},
		{0x0001, 1, 0x0002},
		{0x8000, 1, 0x0001},
		{0x1234, 16, 0x1234},
	}
	for _, tt := range tests {
		got := rotl16(tt.v, tt.n)
		if got != tt.want {
			t.Errorf("rotl16(0x%04x, %d) = 0x%04x, want 0x%04x", tt.v, tt.n, got, tt.want)
		}
	}
}

func TestSwap4(t *testing.T) {
	tests := []struct {
		in   []byte
		want []byte
	}{
		{[]byte{0x12, 0x34, 0x56, 0x78}, []byte{0x78, 0x56, 0x34, 0x12}},
		{[]byte{0x00, 0x00, 0x00, 0x00}, []byte{0x00, 0x00, 0x00, 0x00}},
		{[]byte{0xFF, 0xFF, 0xFF, 0xFF}, []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{[]byte{0xAA, 0xBB, 0xCC, 0xDD}, []byte{0xDD, 0xCC, 0xBB, 0xAA}},
	}
	for _, tt := range tests {
		p := make([]byte, 4)
		copy(p, tt.in)
		swap4(p)
		for i := range p {
			if p[i] != tt.want[i] {
				t.Errorf("swap4(%v) = %v, want %v", tt.in, p, tt.want)
				break
			}
		}
	}
}

func TestFind(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	if idx := find(slice, 3); idx != 2 {
		t.Errorf("find([1,2,3,4,5], 3) = %d, want 2", idx)
	}
	if idx := find(slice, 1); idx != 0 {
		t.Errorf("find([1,2,3,4,5], 1) = %d, want 0", idx)
	}
	if idx := find(slice, 5); idx != 4 {
		t.Errorf("find([1,2,3,4,5], 5) = %d, want 4", idx)
	}
	if idx := find(slice, 99); idx != -1 {
		t.Errorf("find([1,2,3,4,5], 99) = %d, want -1", idx)
	}
	empty := []int{}
	if idx := find(empty, 1); idx != -1 {
		t.Errorf("find([], 1) = %d, want -1", idx)
	}

	// Test with strings
	strs := []string{"a", "b", "c"}
	if idx := find(strs, "b"); idx != 1 {
		t.Errorf("find([a,b,c], b) = %d, want 1", idx)
	}
	if idx := find(strs, "z"); idx != -1 {
		t.Errorf("find([a,b,c], z) = %d, want -1", idx)
	}
}
