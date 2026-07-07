package gui

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

func TestAnimDoOutDoIn(t *testing.T) {
	animList = nil
	defer func() {
		animList = nil
	}()

	// Create a window manually
	win, _ := alloc.New(Window{})
	if win == nil {
		t.Fatal("failed to create window")
	}

	// Test doOut
	anim := NewAnim(win, image.Pt(0, 0), image.Pt(10, 10), image.Pt(1, 1), image.Pt(2, 2))
	if anim == nil {
		t.Fatal("failed to create anim")
	}
	anim.SetState(AnimOut)
	win.SetPos(image.Pt(0, 0))
	for i := 0; i < 20; i++ {
		anim.doOut()
		if anim.State() == AnimOutDone {
			break
		}
	}
	if anim.State() != AnimOutDone {
		t.Errorf("expected AnimOutDone, got %d", anim.State())
	}
	anim.Free()

	// Test doIn
	anim2 := NewAnim(win, image.Pt(0, 0), image.Pt(10, 10), image.Pt(1, 1), image.Pt(2, 2))
	anim2.SetState(AnimIn)
	win.SetPos(image.Pt(10, 10))
	for i := 0; i < 20; i++ {
		anim2.doIn()
		if anim2.State() == AnimInDone {
			break
		}
	}
	if anim2.State() != AnimInDone {
		t.Errorf("expected AnimInDone, got %d", anim2.State())
	}
	anim2.Free()

	// Test AnimTick with valid window
	anim3 := NewAnim(win, image.Pt(0, 0), image.Pt(5, 5), image.Pt(1, 1), image.Pt(1, 1))
	anim3.SetState(AnimOut)
	win.SetPos(image.Pt(0, 0))
	AnimTick()
	anim3.SetState(AnimIn)
	AnimTick()
	anim3.Free()

	alloc.Free(win)
}
