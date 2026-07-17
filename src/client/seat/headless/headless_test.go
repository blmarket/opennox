package headless

import (
	"image"
	"reflect"
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/noxworld-dev/opennox-lib/noximage"
)

func TestSeatResizeAndInputs(t *testing.T) {
	s := New(image.Pt(640, 480))
	if got := s.ScreenSize(); got != image.Pt(640, 480) {
		t.Fatalf("screen size = %v", got)
	}
	var resized []image.Point
	s.OnScreenResize(func(sz image.Point) { resized = append(resized, sz) })
	s.ResizeScreen(image.Pt(640, 480))
	s.ResizeScreen(image.Pt(1024, 768))
	if !reflect.DeepEqual(resized, []image.Point{image.Pt(1024, 768)}) {
		t.Fatalf("resize callbacks = %v", resized)
	}

	a := func(seat.InputEvent) {}
	b := func(seat.InputEvent) {}
	if prev := s.ReplaceInputs(seat.InputConfig{a}); prev != nil {
		t.Fatalf("initial input config = %v", prev)
	}
	s.OnInput(b)
	if prev := s.ReplaceInputs(nil); len(prev) != 2 {
		t.Fatalf("input callback count = %d", len(prev))
	}
	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestSurfaceCopiesPixels(t *testing.T) {
	s := (&Seat{}).NewSurface(image.Pt(2, 2), false).(*Surface)
	img := noximage.NewImage16(image.Rect(0, 0, 2, 2))
	copy(img.Pix, []uint16{1, 2, 3, 4})
	s.Update(img)
	img.Pix[0] = 9
	if !reflect.DeepEqual(s.pix, []uint16{1, 2, 3, 4}) {
		t.Fatalf("surface pixels = %v", s.pix)
	}
	s.Destroy()
	if s.pix != nil {
		t.Fatalf("destroyed surface pixels = %v", s.pix)
	}
}

func TestSurfaceRejectsWrongSize(t *testing.T) {
	s := (&Seat{}).NewSurface(image.Pt(2, 2), false).(*Surface)
	defer func() {
		if recover() == nil {
			t.Fatal("expected size mismatch panic")
		}
	}()
	s.Update(noximage.NewImage16(image.Rect(0, 0, 1, 1)))
}
