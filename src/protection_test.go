package opennox

import (
	"testing"
)

func TestProtectBytes(t *testing.T) {
	tests := []struct {
		name string
		data []byte
		want uint32
	}{
		{"empty", []byte{}, 0},
		{"single byte", []byte{0x12}, 0}, // less than 4 bytes, no full uint32
		{"4 bytes", []byte{0x12, 0x34, 0x56, 0x78}, 0x78563412},
		{"8 bytes", []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE, 0xF0}, 0x78563412 ^ 0xF0DEBC9A},
		{"5 bytes", []byte{0x01, 0x02, 0x03, 0x04, 0x05}, 0x04030201}, // only first 4 bytes
		{"all zeros", []byte{0, 0, 0, 0}, 0},
		{"all ones", []byte{0xFF, 0xFF, 0xFF, 0xFF}, 0xFFFFFFFF},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := protectBytes(tt.data)
			if got != tt.want {
				t.Errorf("protectBytes(%v) = 0x%08x, want 0x%08x", tt.data, got, tt.want)
			}
		})
	}
}

func TestProtectStr(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want uint32
	}{
		{"empty", "", 0},
		{"short", "abc", 0},             // less than 4 bytes
		{"4 chars", "abcd", 0x64636261}, // 'a'=0x61, 'b'=0x62, etc, little endian
		{"8 chars", "abcdefgh", 0x64636261 ^ 0x68676665},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := protectStr(tt.str)
			if got != tt.want {
				t.Errorf("protectStr(%q) = 0x%08x, want 0x%08x", tt.str, got, tt.want)
			}
		})
	}
}

func TestProtectWStr(t *testing.T) {
	tests := []struct {
		name string
		str  string
	}{
		{"empty", ""},
		{"ascii", "abc"},
		{"unicode", "héllo"},
		{"emoji", "😀"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Just ensure it doesn't panic and returns a uint32
			got := protectWStr(tt.str)
			_ = got
		})
	}
}

func TestProtectWStrConsistency(t *testing.T) {
	// Same string should give same result
	s := "test string"
	got1 := protectWStr(s)
	got2 := protectWStr(s)
	if got1 != got2 {
		t.Errorf("protectWStr(%q) not consistent: %d vs %d", s, got1, got2)
	}
}
