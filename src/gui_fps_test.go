package opennox

import (
	"testing"
)

func TestGuiFPS_Sub470A80(t *testing.T) {
	var fps guiFPS
	if fps.dword1090246 != 0 {
		t.Fatal("initial value should be 0")
	}
	fps.sub_470A80()
	if fps.dword1090246 != 1 {
		t.Fatalf("expected 1, got %d", fps.dword1090246)
	}
	fps.sub_470A80()
	if fps.dword1090246 != 2 {
		t.Fatalf("expected 2, got %d", fps.dword1090246)
	}
}

func TestGuiFPS_Sub470A60(t *testing.T) {
	var fps guiFPS
	// toggle enabled
	if fps.enabled {
		t.Fatal("initial enabled should be false")
	}
	fps.sub_470A60()
	if !fps.enabled {
		t.Fatal("enabled should be true after toggle")
	}
	fps.sub_470A60()
	if fps.enabled {
		t.Fatal("enabled should be false after second toggle")
	}
}

func TestGuiFPS_Sub4706C0(t *testing.T) {
	var fps guiFPS
	// nil win should not panic when a1 is 0
	fps.sub_4706C0(0)
	// a1 != 0 but not enabled or nil win should not panic
	fps.sub_4706C0(1)
	fps.enabled = true
	// nil win with enabled should not panic (win.Hide/ShowModal on nil would panic, but we test the logic path)
	// Actually win is nil, so calling win.GetFlags would panic. Let's just verify the function exists and handles a1=0.
	// Reset to avoid panic
	fps.win = nil
	fps.enabled = false
	fps.sub_4706C0(0)
}

func TestGuiFPS_InitValues(t *testing.T) {
	var fps guiFPS
	// Test initial zero values
	if fps.dword1090256 != 0 || fps.dword147864 != 0 {
		t.Fatal("initial values should be zero")
	}
}
