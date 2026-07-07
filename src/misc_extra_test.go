package opennox

import (
	"testing"
)

func TestAbsExtra(t *testing.T) {
	tests := []struct {
		v    int
		want int
	}{
		{0, 0},
		{1, 1},
		{-1, 1},
		{1000000, 1000000},
		{-1000000, 1000000},
	}
	for _, tt := range tests {
		got := abs(tt.v)
		if got != tt.want {
			t.Errorf("abs(%d) = %d, want %d", tt.v, got, tt.want)
		}
		if got < 0 {
			t.Errorf("abs(%d) = %d, want non-negative", tt.v, got)
		}
	}
}

func TestClampExtra(t *testing.T) {
	tests := []struct {
		v, min, max, want int
	}{
		{-10, -5, 5, -5},
		{10, -5, 5, 5},
		{0, -5, 5, 0},
		{-5, -5, 5, -5},
		{5, -5, 5, 5},
		{100, 0, 10, 10},
		{-100, 0, 10, 0},
	}
	for _, tt := range tests {
		got := clamp(tt.v, tt.min, tt.max)
		if got != tt.want {
			t.Errorf("clamp(%d, %d, %d) = %d, want %d", tt.v, tt.min, tt.max, got, tt.want)
		}
	}
}

func TestBool2intExtra(t *testing.T) {
	if bool2int(true) != 1 {
		t.Error("bool2int(true) should be 1")
	}
	if bool2int(false) != 0 {
		t.Error("bool2int(false) should be 0")
	}
	// Test multiple times for consistency
	for i := 0; i < 10; i++ {
		if bool2int(true) != 1 || bool2int(false) != 0 {
			t.Error("bool2int not consistent")
		}
	}
}

func TestRotlExtra(t *testing.T) {
	tests := []struct {
		v    uint32
		n    int
		want uint32
	}{
		{0x00000000, 5, 0x00000000},
		{0xFFFFFFFF, 5, 0xFFFFFFFF},
		{0x12345678, 16, 0x56781234},
		{0x12345678, 32, 0x12345678}, // rotl by 32 should be same?
		{0x80000001, 1, 0x00000003},
	}
	for _, tt := range tests {
		got := rotl(tt.v, tt.n)
		if got != tt.want {
			t.Errorf("rotl(0x%08x, %d) = 0x%08x, want 0x%08x", tt.v, tt.n, got, tt.want)
		}
	}
}

func TestRotl16Extra(t *testing.T) {
	tests := []struct {
		v    uint16
		n    int
		want uint16
	}{
		{0x0000, 5, 0x0000},
		{0xFFFF, 5, 0xFFFF},
		{0x1234, 8, 0x3412},
		{0x8001, 1, 0x0003},
		{0x1234, 0, 0x1234},
	}
	for _, tt := range tests {
		got := rotl16(tt.v, tt.n)
		if got != tt.want {
			t.Errorf("rotl16(0x%04x, %d) = 0x%04x, want 0x%04x", tt.v, tt.n, got, tt.want)
		}
	}
}

func TestSwap4Extra(t *testing.T) {
	tests := []struct {
		in   []byte
		want []byte
	}{
		{[]byte{0x01, 0x02, 0x03, 0x04}, []byte{0x04, 0x03, 0x02, 0x01}},
		{[]byte{0x00, 0x01, 0x02, 0x03}, []byte{0x03, 0x02, 0x01, 0x00}},
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

func TestFindExtra(t *testing.T) {
	// Test with empty slice
	if idx := find([]int{}, 1); idx != -1 {
		t.Errorf("find([], 1) = %d, want -1", idx)
	}
	// Test with single element
	if idx := find([]int{5}, 5); idx != 0 {
		t.Errorf("find([5], 5) = %d, want 0", idx)
	}
	if idx := find([]int{5}, 3); idx != -1 {
		t.Errorf("find([5], 3) = %d, want -1", idx)
	}
	// Test with duplicates (should return first index)
	slice := []int{1, 2, 3, 2, 1}
	if idx := find(slice, 2); idx != 1 {
		t.Errorf("find([1,2,3,2,1], 2) = %d, want 1", idx)
	}
	// Test with strings
	strs := []string{"apple", "banana", "cherry"}
	if idx := find(strs, "banana"); idx != 1 {
		t.Errorf("find([apple,banana,cherry], banana) = %d, want 1", idx)
	}
}
