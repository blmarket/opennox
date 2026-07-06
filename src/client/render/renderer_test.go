package render

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/noxworld-dev/opennox-lib/noximage"
)

type mockScreen struct {
	sz   image.Point
	mode seat.ScreenMode
}

func (m *mockScreen) ScreenSize() image.Point            { return m.sz }
func (m *mockScreen) ScreenMaxSize() image.Point         { return m.sz }
func (m *mockScreen) ResizeScreen(sz image.Point)        { m.sz = sz }
func (m *mockScreen) SetScreenMode(mode seat.ScreenMode) { m.mode = mode }
func (m *mockScreen) SetGamma(v float32)                 {}
func (m *mockScreen) NewSurface(sz image.Point, filtering bool) seat.Surface {
	return &mockSurface{sz: sz}
}
func (m *mockScreen) Clear()                                             {}
func (m *mockScreen) Present()                                           {}
func (m *mockScreen) OnScreenResize(fnc func(sz image.Point))            {}
func (m *mockScreen) ReplaceInputs(in seat.InputConfig) seat.InputConfig { return nil }
func (m *mockScreen) OnInput(fnc func(ev seat.InputEvent))               {}
func (m *mockScreen) InputTick()                                         {}
func (m *mockScreen) SetTextInput(enable bool)                           {}

type mockSurface struct {
	sz image.Point
}

func (m *mockSurface) Size() image.Point            { return m.sz }
func (m *mockSurface) Update(img *noximage.Image16) {}
func (m *mockSurface) Draw(view image.Rectangle)    {}
func (m *mockSurface) Destroy()                     {}

func TestRendererWindowMode(t *testing.T) {
	ms := &mockScreen{sz: image.Pt(800, 600)}
	r, err := New(ms)
	if err != nil {
		t.Fatalf("new failed: %v", err)
	}
	if r.WindowMode() != -4 {
		t.Fatalf("expected -4")
	}
	r.SetWindowMode(1)
	if !r.IsFullScreen() || r.WindowMode() != 1 {
		t.Fatalf("expected fullscreen")
	}
	r.SetWindowMode(-2)
	if !r.IsFullScreen() {
		t.Fatalf("expected fullscreen")
	}
	r.ToggleWindowMode()
	if r.IsFullScreen() {
		t.Fatalf("expected windowed after toggle")
	}
	r.ToggleWindowMode()
	if !r.IsFullScreen() {
		t.Fatalf("expected fullscreen after toggle")
	}
	r.SetWindowMode(-3)
	r.SetStretched(true)
	if !r.GetStretched() {
		t.Fatalf("expected stretched")
	}
	r.SetFiltering(false)
	if r.GetFiltering() {
		t.Fatalf("expected no filtering")
	}
	if r.Ticks() != 0 {
		t.Fatalf("expected 0 ticks")
	}
}
