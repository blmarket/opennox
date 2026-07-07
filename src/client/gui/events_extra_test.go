package gui

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox/v1/client/input"
)

func TestWindowEventsExtra(t *testing.T) {
	// RawEvent
	re := &RawEvent{Event: 99, Arg1: 1, Arg2: 2}
	if re.EventCode() != 99 {
		t.Error("RawEvent EventCode failed")
	}
	a1, a2 := re.EventArgsC()
	if a1 != 1 || a2 != 2 {
		t.Error("RawEvent EventArgsC failed")
	}

	// WindowInit
	wi := WindowInit{}
	if wi.EventCode() != 1 {
		t.Error("WindowInit EventCode failed")
	}
	if a1, a2 := wi.EventArgsC(); a1 != 0 || a2 != 0 {
		t.Error("WindowInit EventArgsC failed")
	}

	// WindowDestroy
	wd := WindowDestroy{}
	if wd.EventCode() != 2 {
		t.Error("WindowDestroy EventCode failed")
	}

	// WindowMouseState
	wms := &WindowMouseState{State: input.NOX_MOUSE_LEFT_DOWN, Pos: image.Pt(10, 20)}
	if wms.EventCode() != int(input.NOX_MOUSE_LEFT_DOWN) {
		t.Error("WindowMouseState EventCode failed")
	}
	if _, a2 := wms.EventArgsC(); a2 != 0 {
		t.Error("WindowMouseState EventArgsC failed")
	}

	// WindowMouseUnk
	wmu := &WindowMouseUnk{Event: 17, Pos: image.Pt(5, 15)}
	if wmu.EventCode() != 17 {
		t.Error("WindowMouseUnk EventCode failed")
	}
	wmu.EventArgsC()

	// WindowKeyPress
	wkp := WindowKeyPress{Key: keybind.KeyA, Pressed: true}
	if wkp.EventCode() != 21 {
		t.Error("WindowKeyPress EventCode failed")
	}
	if a1, a2 := wkp.EventArgsC(); a1 != uintptr(keybind.KeyA) || a2 != 2 {
		t.Error("WindowKeyPress EventArgsC pressed failed")
	}
	wkp2 := WindowKeyPress{Key: keybind.KeyB, Pressed: false}
	if _, a2 := wkp2.EventArgsC(); a2 != 1 {
		t.Error("WindowKeyPress EventArgsC not pressed failed")
	}

	// WindowNewChild
	wnc := WindowNewChild{ID: 42}
	if wnc.EventCode() != 22 {
		t.Error("WindowNewChild EventCode failed")
	}
	if a1, a2 := wnc.EventArgsC(); a1 != 42 || a2 != 0 {
		t.Error("WindowNewChild EventArgsC failed")
	}

	// WindowFocus
	wf := WindowFocus(true)
	if wf.EventCode() != 23 {
		t.Error("WindowFocus EventCode failed")
	}
	if a1, _ := wf.EventArgsC(); a1 != 1 {
		t.Error("WindowFocus EventArgsC true failed")
	}
	wf2 := WindowFocus(false)
	if a1, _ := wf2.EventArgsC(); a1 != 0 {
		t.Error("WindowFocus EventArgsC false failed")
	}

	// AsWindowEvent
	ev := AsWindowEvent(1, 0, 0)
	if _, ok := ev.(WindowInit); !ok {
		t.Error("AsWindowEvent WindowInit failed")
	}
	ev2 := AsWindowEvent(2, 0, 0)
	if _, ok := ev2.(WindowDestroy); !ok {
		t.Error("AsWindowEvent WindowDestroy failed")
	}
	ev3 := AsWindowEvent(21, uintptr(keybind.KeyA), 2)
	if kp, ok := ev3.(WindowKeyPress); !ok || kp.Key != keybind.KeyA || !kp.Pressed {
		t.Error("AsWindowEvent WindowKeyPress failed")
	}
	pos := uintptr(uint32(10) | (uint32(20) << 16))
	ev4 := AsWindowEvent(int(input.NOX_MOUSE_LEFT_DOWN), pos, 0)
	if _, ok := ev4.(*WindowMouseState); !ok {
		t.Error("AsWindowEvent MouseState failed")
	}
	ev5 := AsWindowEvent(17, pos, 0)
	if _, ok := ev5.(*WindowMouseUnk); !ok {
		t.Error("AsWindowEvent MouseUnk failed")
	}
	ev6 := AsWindowEvent(999, 1, 2)
	if re, ok := ev6.(*RawEvent); !ok || re.Event != 999 {
		t.Error("AsWindowEvent RawEvent failed")
	}

	if EventRespBool(nil) {
		t.Error("EventRespBool nil should be false")
	}
	if !EventRespBool(RawEventResp(1)) {
		t.Error("EventRespBool should be true")
	}
	if EventRespInt(nil) != 0 {
		t.Error("EventRespInt nil should be 0")
	}
	if EventRespInt(RawEventResp(5)) != 5 {
		t.Error("EventRespInt failed")
	}
	if EventRespPtr(nil) != nil {
		t.Error("EventRespPtr nil should be nil")
	}
	if EventRespPtr(RawEventResp(0x1234)) == nil {
		t.Error("EventRespPtr should not be nil")
	}
	if RawEventResp(42).EventRespC() != 42 {
		t.Error("RawEventResp EventRespC failed")
	}
}
