package input

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox-lib/client/seat"
)

type mockInput2 struct {
	onInput func(ev seat.InputEvent)
}

func (m *mockInput2) InputTick()                                          {}
func (m *mockInput2) ReplaceInputs(cfg seat.InputConfig) seat.InputConfig { return nil }
func (m *mockInput2) OnInput(fnc func(ev seat.InputEvent))                { m.onInput = fnc }
func (m *mockInput2) SetTextInput(enable bool)                            {}

func TestKeyToWCharExtra(t *testing.T) {
	mock := &mockInput2{}
	h := New(mock, false, 0)
	// Test all language codes
	for code := 0; code <= 9; code++ {
		h.SetLanguage(code)
		// Test basic keys
		_ = h.KeyToWChar(keybind.KeyA)
		_ = h.KeyToWChar(keybind.KeyLShift)
		_ = h.KeyToWChar(keybind.KeyRShift)
		_ = h.KeyToWChar(keybind.KeyCaps)
		// Test key > 0xFF
		_ = h.KeyToWChar(0x1FF)
	}
	// Test str2u16
	if v := str2u16(""); v != 0 {
		t.Error("expected 0 for empty")
	}
	if v := str2u16("A"); v != uint16('A') {
		t.Error("str2u16 failed")
	}
	if v := str2u16("AB"); v != uint16('A')|uint16('B')<<8 {
		t.Error("str2u16 multi failed")
	}
	// Test iswalpha
	if !iswalpha(uint16('A')) {
		t.Error("expected iswalpha true for A")
	}
	if iswalpha(200) {
		t.Error("expected iswalpha false for >=192")
	}
	if iswalpha(170) {
		t.Error("expected iswalpha false for 170")
	}
}
