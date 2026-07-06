package gui

import (
	"image"
	"testing"
)

func TestEventResp(t *testing.T) {
	// Test nil responses
	if EventRespBool(nil) != false {
		t.Error("EventRespBool(nil) should be false")
	}
	if EventRespInt(nil) != 0 {
		t.Error("EventRespInt(nil) should be 0")
	}
	if EventRespPtr(nil) != nil {
		t.Error("EventRespPtr(nil) should be nil")
	}
}

func TestAsWindowEvent(t *testing.T) {
	// Test with nil
	if AsWindowEvent(999, 0, 0) == nil {
		t.Error("AsWindowEvent should not be nil for unregistered event")
	}
}

func TestWindowEvents(t *testing.T) {
	// Test WindowInit
	wi := WindowInit{}
	if wi.EventCode() != 1 {
		t.Error("WindowInit.EventCode should be 1")
	}
	a1, a2 := wi.EventArgsC()
	if a1 != 0 || a2 != 0 {
		t.Error("WindowInit.EventArgsC should be 0, 0")
	}

	// Test WindowDestroy
	wd := WindowDestroy{}
	if wd.EventCode() != 2 {
		t.Error("WindowDestroy.EventCode should be 2")
	}

	// Test WindowKeyPress
	wkp := WindowKeyPress{}
	if wkp.EventCode() != 21 {
		t.Error("WindowKeyPress.EventCode should be 21")
	}

	// Test WindowNewChild
	wnc := WindowNewChild{ID: 123}
	if wnc.EventCode() != 22 {
		t.Error("WindowNewChild.EventCode should be 22")
	}

	// Test WindowFocus
	wf := WindowFocus(true)
	if wf.EventCode() != 23 {
		t.Error("WindowFocus.EventCode should be 23")
	}

	// Test WindowMouseState
	wms := WindowMouseState{}
	if wms.EventCode() != 0 {
		t.Error("WindowMouseState.EventCode should be 0")
	}

	// Test WindowMouseUnk
	wmu := WindowMouseUnk{}
	if wmu.EventCode() != 0 {
		t.Error("WindowMouseUnk.EventCode should be 0")
	}
}

func TestAnim(t *testing.T) {
	// Test SetAnimGlobalState and AnimGlobalState
	SetAnimGlobalState(5)
	if AnimGlobalState() != 5 {
		t.Error("AnimGlobalState should be 5")
	}

	// Test NewAnim
	a := NewAnim(nil, image.Point{}, image.Point{}, image.Point{}, image.Point{})
	if a == nil {
		t.Fatal("NewAnim returned nil")
	}
	if a.State() != AnimIn {
		t.Error("new anim state should be AnimIn")
	}

	// Test SetState
	a.SetState(2)
	if a.State() != 2 {
		t.Error("anim state should be 2")
	}

	// Test Window (should be nil initially)
	if a.Window() != nil {
		t.Error("anim window should be nil")
	}

	// Test Free
	a.Free()

	// Test FindAnimForStateID (should return nil for nonexistent)
	if FindAnimForStateID(999) != nil {
		t.Error("FindAnimForStateID should return nil for nonexistent")
	}

	// Test AnimTick (should not panic)
	AnimTick()
}

func TestGUI(t *testing.T) {
	// Test FocusMainBg (should not panic)
	FocusMainBg()
}
