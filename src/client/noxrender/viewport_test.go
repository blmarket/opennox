package noxrender

import (
	"image"
	"testing"
	"unsafe"
)

func TestViewportC(t *testing.T) {
	vp := &Viewport{}
	ptr := vp.C()
	if ptr != unsafe.Pointer(vp) {
		t.Error("C() should return unsafe.Pointer to viewport")
	}
}

func TestViewportToScreenPos(t *testing.T) {
	vp := &Viewport{
		Screen: image.Rect(10, 20, 110, 120),
		World:  image.Rect(100, 200, 300, 400),
	}
	pos := image.Pt(150, 250)
	result := vp.ToScreenPos(pos)
	// pos.Sub(World.Min) = (50, 50), Add(Screen.Min) = (60, 70)
	expected := image.Pt(60, 70)
	if result != expected {
		t.Errorf("ToScreenPos(%v) = %v, want %v", pos, result, expected)
	}
}

func TestViewportToWorldPos(t *testing.T) {
	vp := &Viewport{
		Screen: image.Rect(10, 20, 110, 120),
		World:  image.Rect(100, 200, 300, 400),
	}
	pos := image.Pt(60, 70)
	result := vp.ToWorldPos(pos)
	// pos.Sub(Screen.Min) = (50, 50), Add(World.Min) = (150, 250)
	expected := image.Pt(150, 250)
	if result != expected {
		t.Errorf("ToWorldPos(%v) = %v, want %v", pos, result, expected)
	}
}

func TestViewportToScreenPosZero(t *testing.T) {
	vp := &Viewport{
		Screen: image.Rect(0, 0, 100, 100),
		World:  image.Rect(0, 0, 100, 100),
	}
	pos := image.Pt(50, 50)
	result := vp.ToScreenPos(pos)
	if result != pos {
		t.Errorf("ToScreenPos with identical rects should return same pos, got %v", result)
	}
}

func TestViewportToWorldPosZero(t *testing.T) {
	vp := &Viewport{
		Screen: image.Rect(0, 0, 100, 100),
		World:  image.Rect(0, 0, 100, 100),
	}
	pos := image.Pt(50, 50)
	result := vp.ToWorldPos(pos)
	if result != pos {
		t.Errorf("ToWorldPos with identical rects should return same pos, got %v", result)
	}
}
