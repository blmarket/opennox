package render

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/noxworld-dev/opennox-lib/noximage"
)

func TestRendererWindowMode_Extra(t *testing.T) {
	ms := &mockScreen{sz: image.Pt(800, 600)}
	r, err := New(ms)
	if err != nil {
		t.Fatalf("new failed: %v", err)
	}
	// Test initial mode
	if r.WindowMode() != -4 {
		t.Errorf("Initial WindowMode should be -4, got %d", r.WindowMode())
	}
	// Test SetWindowMode with various values
	modes := []int{-4, -3, -2, -1, 0, 1, 2}
	for _, m := range modes {
		r.SetWindowMode(m)
		if r.WindowMode() != m {
			t.Errorf("SetWindowMode(%d) failed, got %d", m, r.WindowMode())
		}
	}
}

func TestRenderer_ToggleWindowMode_Extra(t *testing.T) {
	ms := &mockScreen{sz: image.Pt(1024, 768)}
	r, err := New(ms)
	if err != nil {
		t.Fatalf("new failed: %v", err)
	}
	initial := r.IsFullScreen()
	r.ToggleWindowMode()
	if r.IsFullScreen() == initial {
		t.Error("ToggleWindowMode should change fullscreen state")
	}
	r.ToggleWindowMode()
	if r.IsFullScreen() != initial {
		t.Error("ToggleWindowMode twice should return to initial state")
	}
}

func TestRenderer_StretchFiltering_Extra(t *testing.T) {
	ms := &mockScreen{sz: image.Pt(800, 600)}
	r, err := New(ms)
	if err != nil {
		t.Fatalf("new failed: %v", err)
	}
	// Test stretch
	r.SetStretched(true)
	if !r.GetStretched() {
		t.Error("GetStretched should return true after SetStretched(true)")
	}
	r.SetStretched(false)
	if r.GetStretched() {
		t.Error("GetStretched should return false after SetStretched(false)")
	}
	// Test filtering
	r.SetFiltering(true)
	if !r.GetFiltering() {
		t.Error("GetFiltering should return true after SetFiltering(true)")
	}
	r.SetFiltering(false)
	if r.GetFiltering() {
		t.Error("GetFiltering should return false after SetFiltering(false)")
	}
}

// Mock types already defined in renderer_test.go
type mockScreen2 struct {
	sz   image.Point
	mode seat.ScreenMode
}

func (m *mockScreen2) ScreenSize() image.Point            { return m.sz }
func (m *mockScreen2) ScreenMaxSize() image.Point         { return m.sz }
func (m *mockScreen2) ResizeScreen(sz image.Point)        { m.sz = sz }
func (m *mockScreen2) SetScreenMode(mode seat.ScreenMode) { m.mode = mode }
func (m *mockScreen2) SetGamma(v float32)                 {}
func (m *mockScreen2) NewSurface(sz image.Point, filtering bool) seat.Surface {
	return &mockSurface2{sz: sz}
}
func (m *mockScreen2) Clear()                                             {}
func (m *mockScreen2) Present()                                           {}
func (m *mockScreen2) OnScreenResize(fnc func(sz image.Point))            {}
func (m *mockScreen2) ReplaceInputs(in seat.InputConfig) seat.InputConfig { return nil }
func (m *mockScreen2) OnInput(fnc func(ev seat.InputEvent))               {}
func (m *mockScreen2) InputTick()                                         {}
func (m *mockScreen2) SetTextInput(enable bool)                           {}

type mockSurface2 struct {
	sz image.Point
}

func (m *mockSurface2) Size() image.Point            { return m.sz }
func (m *mockSurface2) Update(img *noximage.Image16) {}
func (m *mockSurface2) Draw(view image.Rectangle)    {}
func (m *mockSurface2) Destroy()                     {}
