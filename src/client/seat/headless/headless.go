// Package headless provides an in-memory seat for running the client without
// creating a native operating-system window.
package headless

import (
	"image"
	"sync"

	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/noxworld-dev/opennox-lib/noximage"
)

// Seat implements seat.Seat entirely in memory.
type Seat struct {
	mu       sync.Mutex
	size     image.Point
	mode     seat.ScreenMode
	inputs   seat.InputConfig
	onResize []func(image.Point)
	closed   bool
}

// New creates an in-memory seat with the requested screen size.
func New(sz image.Point) *Seat {
	return &Seat{size: sz, mode: seat.Windowed}
}

func (s *Seat) ScreenSize() image.Point {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size
}

func (s *Seat) ScreenMaxSize() image.Point { return s.ScreenSize() }

func (s *Seat) ResizeScreen(sz image.Point) {
	s.mu.Lock()
	if s.size == sz {
		s.mu.Unlock()
		return
	}
	s.size = sz
	hooks := append([]func(image.Point){}, s.onResize...)
	s.mu.Unlock()
	for _, fnc := range hooks {
		fnc(sz)
	}
}

func (s *Seat) SetScreenMode(mode seat.ScreenMode) {
	s.mu.Lock()
	s.mode = mode
	s.mu.Unlock()
}

func (s *Seat) SetGamma(float32) {}

func (s *Seat) OnScreenResize(fnc func(image.Point)) {
	s.mu.Lock()
	s.onResize = append(s.onResize, fnc)
	s.mu.Unlock()
}

func (s *Seat) NewSurface(sz image.Point, _ bool) seat.Surface {
	return &Surface{size: sz}
}

func (s *Seat) Clear()     {}
func (s *Seat) Present()   {}
func (s *Seat) InputTick() {}

func (s *Seat) ReplaceInputs(cfg seat.InputConfig) seat.InputConfig {
	s.mu.Lock()
	defer s.mu.Unlock()
	prev := s.inputs
	s.inputs = cfg
	return prev
}

func (s *Seat) OnInput(fnc func(seat.InputEvent)) {
	s.mu.Lock()
	s.inputs = append(s.inputs, fnc)
	s.mu.Unlock()
}

func (s *Seat) SetTextInput(bool) {}

func (s *Seat) Close() error {
	s.mu.Lock()
	s.closed = true
	s.inputs = nil
	s.onResize = nil
	s.mu.Unlock()
	return nil
}

// Surface is the in-memory framebuffer owned by a headless seat.
type Surface struct {
	mu   sync.Mutex
	size image.Point
	pix  []uint16
}

func (s *Surface) Size() image.Point {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.size
}

func (s *Surface) Update(img *noximage.Image16) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if img.Size() != s.size {
		panic("headless seat: invalid image size")
	}
	s.pix = append(s.pix[:0], img.Pix...)
}

func (s *Surface) Draw(image.Rectangle) {}

func (s *Surface) Destroy() {
	s.mu.Lock()
	s.pix = nil
	s.mu.Unlock()
}

var _ seat.Seat = (*Seat)(nil)
var _ seat.Surface = (*Surface)(nil)
