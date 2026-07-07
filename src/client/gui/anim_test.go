package gui

import (
	"image"
	"testing"
)

func TestFindAnimForStateID(t *testing.T) {
	// Ensure clean state
	animList = nil

	// Not found should return nil
	if got := FindAnimForStateID(999); got != nil {
		t.Errorf("expected nil for non-existent state, got %v", got)
	}

	// Create an anim with nil window (like existing test)
	anim := NewAnim(nil, image.Pt(0, 0), image.Pt(10, 10), image.Pt(1, 1), image.Pt(1, 1))
	if anim == nil {
		t.Fatal("failed to create anim")
	}
	defer anim.Free()

	// Set a specific StateID
	anim.StateID = 42

	// Should find it
	found := FindAnimForStateID(42)
	if found == nil {
		t.Error("expected to find anim with StateID 42, got nil")
	} else if found != anim {
		t.Error("found anim does not match created anim")
	}

	// Not found for other ID
	if got := FindAnimForStateID(12345); got != nil {
		t.Errorf("expected nil for non-existent ID, got %v", got)
	}
}

func TestAnimState(t *testing.T) {
	// Nil anim should return -1
	var nilAnim *Anim
	if s := nilAnim.State(); s != -1 {
		t.Errorf("expected -1 for nil anim state, got %d", s)
	}
	if w := nilAnim.Window(); w != nil {
		t.Errorf("expected nil window for nil anim, got %v", w)
	}

	anim := NewAnim(nil, image.Pt(0, 0), image.Pt(5, 5), image.Pt(1, 1), image.Pt(1, 1))
	if anim == nil {
		t.Fatal("failed to create anim")
	}
	defer anim.Free()

	// Test SetState and State
	anim.SetState(AnimOut)
	if s := anim.State(); s != AnimOut {
		t.Errorf("expected AnimOut, got %d", s)
	}
	anim.SetState(AnimIn)
	if s := anim.State(); s != AnimIn {
		t.Errorf("expected AnimIn, got %d", s)
	}

	// Test Window() (should be nil as we passed nil)
	if w := anim.Window(); w != nil {
		t.Error("anim.Window() should be nil")
	}

	// Test Free nil
	var nilAnim2 *Anim
	nilAnim2.Free() // should not panic
}

func TestAnimGlobalState(t *testing.T) {
	SetAnimGlobalState(AnimInDone)
	if s := AnimGlobalState(); s != AnimInDone {
		t.Errorf("expected AnimInDone, got %d", s)
	}
	SetAnimGlobalState(AnimOutDone)
	if s := AnimGlobalState(); s != AnimOutDone {
		t.Errorf("expected AnimOutDone, got %d", s)
	}
}

func TestAnimTick(t *testing.T) {
	// Ensure clean state
	animList = nil

	// Create anim in AnimOut state with nil window
	// Note: doOut/doIn will panic with nil window, so we test with empty list
	AnimTick() // should not panic with empty list

	anim := NewAnim(nil, image.Pt(0, 0), image.Pt(10, 10), image.Pt(1, 1), image.Pt(1, 1))
	if anim == nil {
		t.Fatal("failed to create anim")
	}
	defer anim.Free()

	// Set state to something that won't call doOut/doIn with nil window
	// (doOut/doIn require valid window, so we set to done states)
	anim.SetState(AnimOutDone)
	AnimTick() // should not call doOut

	anim.SetState(AnimInDone)
	AnimTick() // should not call doIn

	// Clean up
	animList = nil
}
