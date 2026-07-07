package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
)

func TestCtrlEventBinding_DefKey_Extra(t *testing.T) {
	// Test defKey with empty keys
	b := &CtrlEventBinding{}
	if b.defKey() != 0 {
		t.Error("defKey on empty binding should return 0")
	}
	// Test defKey with keys
	b.keys = []keybind.Key{keybind.KeyA, keybind.KeyB}
	if b.defKey() != keybind.KeyA {
		t.Errorf("defKey should return first key, got %v", b.defKey())
	}
}

func TestCtrlEventBinding_DefEvent_Extra(t *testing.T) {
	// Test defEvent with empty events
	b := &CtrlEventBinding{}
	if b.defEvent() != 0 {
		t.Error("defEvent on empty binding should return 0")
	}
	// Test defEvent with events
	b.events = []keybind.Event{keybind.EventAction, keybind.EventJump}
	if b.defEvent() != keybind.EventAction {
		t.Errorf("defEvent should return first event, got %v", b.defEvent())
	}
}

func TestCtrlEventHandler_Reset_Extra(t *testing.T) {
	h := &CtrlEventHandler{}
	h.bindings = &CtrlEventBinding{}
	h.flags750956 = true
	h.indA = 5
	h.indB = 3
	h.Reset()
	if h.bindings != nil {
		t.Error("Reset should set bindings to nil")
	}
	if h.flags750956 {
		t.Error("Reset should set flags750956 to false")
	}
	if h.indA != 0 || h.indB != 0 || h.indC != 0 || h.indD != 0 {
		t.Error("Reset should set all indices to 0")
	}
}

func TestCtrlEventHandler_AddBinding_Extra(t *testing.T) {
	h := &CtrlEventHandler{}
	b1 := &CtrlEventBinding{}
	h.addBinding(b1)
	if h.bindings != b1 {
		t.Error("addBinding should set bindings to new binding")
	}
	if b1.prev != nil {
		t.Error("New binding prev should be nil")
	}
	b2 := &CtrlEventBinding{}
	h.addBinding(b2)
	if h.bindings != b2 {
		t.Error("addBinding should update bindings to new binding")
	}
	if b2.next != b1 {
		t.Error("New binding next should point to previous binding")
	}
	if b1.prev != b2 {
		t.Error("Previous binding prev should point to new binding")
	}
}

func TestCtrlEventHandler_ListBindings_Extra(t *testing.T) {
	h := &CtrlEventHandler{}
	// Test empty list
	list := h.listBindings()
	if list != nil {
		t.Error("listBindings on empty handler should return nil")
	}
	// Test with bindings
	b1 := &CtrlEventBinding{}
	b2 := &CtrlEventBinding{}
	h.addBinding(b1)
	h.addBinding(b2)
	list = h.listBindings()
	if len(list) != 2 {
		t.Errorf("listBindings should return 2 items, got %d", len(list))
	}
}

func TestCtrlEventHandler_HasDefBinding_Extra(t *testing.T) {
	h := &CtrlEventHandler{}
	b := &CtrlEventBinding{
		keys:   []keybind.Key{keybind.KeyA},
		events: []keybind.Event{keybind.EventAction},
	}
	h.addBinding(b)
	if !h.hasDefBinding(keybind.EventAction, keybind.KeyA) {
		t.Error("hasDefBinding should return true for existing binding")
	}
	if h.hasDefBinding(keybind.EventJump, keybind.KeyA) {
		t.Error("hasDefBinding should return false for non-existing event")
	}
	if h.hasDefBinding(keybind.EventAction, keybind.KeyB) {
		t.Error("hasDefBinding should return false for non-existing key")
	}
}
