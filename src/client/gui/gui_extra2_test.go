package gui

import (
	"testing"
)

func TestParseColor_Extra2(t *testing.T) {
	tests := []struct {
		input    string
		r, g, b int
	}{
		{"255 0 0", 255, 0, 0},
		{"0 255 0", 0, 255, 0},
		{"0 0 255", 0, 0, 255},
		{"", 0, 0, 0},
		{"invalid", 0, 0, 0},
	}
	for _, tt := range tests {
		r, g, b := ParseColor(tt.input)
		if r != tt.r || g != tt.g || b != tt.b {
			t.Errorf("ParseColor(%q) = %d,%d,%d, want %d,%d,%d", tt.input, r, g, b, tt.r, tt.g, tt.b)
		}
	}
}

func TestParseNextField_Extra2(t *testing.T) {
	v, rest := ParseNextField("hello world")
	if v != "hello" {
		t.Errorf("ParseNextField v = %q, want %q", v, "hello")
	}
	if rest != "world" {
		t.Errorf("ParseNextField rest = %q, want %q", rest, "world")
	}
	v, rest = ParseNextField("")
	if v != "" {
		t.Errorf("ParseNextField empty v = %q, want empty", v)
	}
}

func TestParseNextUintField_Extra2(t *testing.T) {
	v, rest := ParseNextUintField("123 abc")
	if v != 123 {
		t.Errorf("ParseNextUintField v = %d, want 123", v)
	}
	if rest != "abc" {
		t.Errorf("ParseNextUintField rest = %q, want %q", rest, "abc")
	}
}

func TestParseNextIntField_Extra2(t *testing.T) {
	v, rest := ParseNextIntField("-42 xyz")
	if v != -42 {
		t.Errorf("ParseNextIntField v = %d, want -42", v)
	}
	if rest != "xyz" {
		t.Errorf("ParseNextIntField rest = %q, want %q", rest, "xyz")
	}
}
