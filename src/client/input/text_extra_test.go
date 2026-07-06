package input

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
)

func TestStr2U16(t *testing.T) {
	if got := str2u16(""); got != 0 {
		t.Errorf("str2u16 empty = %d, want 0", got)
	}
	if got := str2u16("A"); got != uint16('A') {
		t.Errorf("str2u16 A = %d, want %d", got, 'A')
	}
	if got := str2u16("AB"); got != uint16('A')|uint16('B')<<8 {
		t.Errorf("str2u16 AB = %d", got)
	}
}

func TestIswalpha(t *testing.T) {
	if !iswalpha('A') {
		t.Error("A should be alpha")
	}
	if !iswalpha('z') {
		t.Error("z should be alpha")
	}
	if iswalpha('1') {
		t.Error("1 should not be alpha")
	}
	if iswalpha(200) {
		t.Error("200 should not be alpha")
	}
	if iswalpha(170) {
		t.Error("170 should not be alpha")
	}
}

func TestHandlerSetLanguage(t *testing.T) {
	h := &Handler{k: &keyboardHandler{}}
	h.SetLanguage(0)
	if h.textMap == nil {
		t.Error("textMap should not be nil")
	}
	h.SetLanguage(1)
	if h.modKey != 0xb8 {
		t.Error("modKey should be 0xb8 for lang 1")
	}
}

func TestHandlerKeyToWChar(t *testing.T) {
	h := &Handler{k: &keyboardHandler{}}
	h.SetLanguage(0)
	if got := h.KeyToWChar(0x1FF); got != 0x1FF {
		t.Errorf("KeyToWChar >0xFF = %d, want 0x1FF", got)
	}
	if got := h.KeyToWChar(keybind.KeyLShift); got != 0 {
		t.Errorf("Shift should return 0, got %d", got)
	}
	if got := h.KeyToWChar(0x10); got == 0 {
		t.Error("q should not return 0")
	}
}
