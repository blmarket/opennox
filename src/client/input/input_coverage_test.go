package input

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/stretchr/testify/require"
)

func TestHandlerMouseKeyboard(t *testing.T) {
	mock := &mockInput{}
	h := New(mock, false, 0)
	require.NotNil(t, h)

	// Test mouse position functions
	pos := h.GetMousePos()
	_ = pos
	rel := h.GetMouseRel()
	_ = rel

	h.ChangeMousePos(image.Pt(10, 20), true)
	h.ChangeMousePos(image.Pt(5, 5), false)

	h.SetMouseBounds(image.Rect(0, 0, 100, 100))

	// Test mouse button functions
	pressed := h.IsMousePressed(seat.MouseButtonLeft)
	_ = pressed

	st := h.GetMouseState(seat.MouseButtonLeft)
	_ = st

	h.SetMouseState(seat.MouseButtonLeft, st)

	// Test mouse action
	action := h.MouseAction(0, 0)
	_ = action

	// Test dist slow
	ds := h.GetDistSlow()
	_ = ds

	// Test mouse wheel
	wheel := h.GetMouseWheel()
	_ = wheel
	h.SetMouseWheel(5)
	require.Equal(t, 5, h.GetMouseWheel())

	// Test keyboard functions
	released := h.IsReleased(0)
	_ = released

	ctrl := h.KeyModCtrl()
	_ = ctrl

	// KeyModAlt and KeyModShift already tested in other tests
	_ = h.KeyModAlt()
	_ = h.KeyModShift()

	flag := h.GetKeyFlag(0)
	_ = flag
	h.SetKeyFlag(0, true)

	keys := h.KeyboardKeys()
	_ = keys

	h.SetTextInput(true)
	h.SetTextInput(false)

	buf := h.GetTextEditBuf()
	_ = buf

	// Test callbacks
	h.OnToggleFullScreen(func() {})
	h.OnMouseWheel(func(delta int) {})
	h.OnKeyPress(func(k keybind.Key) {})
	h.OnInputString(func(s string) {})

	// Test window size
	h.SetWinSize(image.Rect(0, 0, 800, 600))
	h.SetDrawWinSize(image.Pt(800, 600))

	// Test reset and tick
	h.Reset()
	h.Tick()
}

func TestKeyboardHandler(t *testing.T) {
	mock := &mockInput{}
	h := New(mock, false, 0)
	require.NotNil(t, h)

	// Test keyboard handler methods via Handler
	h.OnToggleFullScreen(func() {})
	h.OnInputString(func(s string) {})
	h.OnKeyPress(func(k keybind.Key) {})

	buf := h.GetTextEditBuf()
	_ = buf

	flag := h.GetKeyFlag(keybind.KeyA)
	_ = flag

	released := h.IsReleased(keybind.KeyA)
	_ = released

	keys := h.KeyboardKeys()
	_ = keys

	ctrl := h.KeyModCtrl()
	_ = ctrl
}

func TestMouseHandler(t *testing.T) {
	mock := &mockInput{}
	h := New(mock, false, 0)
	require.NotNil(t, h)

	// Test mouse handler via Handler
	pos := h.GetMousePos()
	_ = pos

	rel := h.GetMouseRel()
	_ = rel

	h.ChangeMousePos(image.Pt(1, 1), true)
	h.SetMouseBounds(image.Rect(0, 0, 10, 10))

	pressed := h.IsMousePressed(seat.MouseButtonRight)
	_ = pressed

	st := h.GetMouseState(seat.MouseButtonRight)
	h.SetMouseState(seat.MouseButtonRight, st)

	action := h.MouseAction(1, 1)
	_ = action

	ds := h.GetDistSlow()
	_ = ds

	wheel := h.GetMouseWheel()
	h.SetMouseWheel(wheel + 1)
}
