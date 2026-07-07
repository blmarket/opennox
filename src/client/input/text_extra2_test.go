package input

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/stretchr/testify/require"
)

type mockInput3 struct {
	onInput func(ev seat.InputEvent)
}

func (m *mockInput3) InputTick()                                          {}
func (m *mockInput3) ReplaceInputs(cfg seat.InputConfig) seat.InputConfig { return nil }
func (m *mockInput3) OnInput(fnc func(ev seat.InputEvent))                { m.onInput = fnc }
func (m *mockInput3) SetTextInput(enable bool)                            {}

func TestText_Extra2(t *testing.T) {
	mock := &mockInput3{}
	h := New(mock, false, 0)

	// Test SetLanguage
	for code := 0; code <= 5; code++ {
		h.SetLanguage(code)
	}

	// Test KeyToWChar
	c := h.KeyToWChar(keybind.KeyA)
	_ = c

	c = h.KeyToWChar(keybind.KeyLShift)
	_ = c

	// Test str2u16
	s := str2u16("")
	require.Equal(t, uint16(0), s)

	s = str2u16("A")
	require.Equal(t, uint16('A'), s)

	s = str2u16("AB")
	require.Equal(t, uint16('A')|uint16('B')<<8, s)

	// Test iswalpha
	require.True(t, iswalpha(uint16('a')))
	require.True(t, iswalpha(uint16('Z')))
	require.False(t, iswalpha(uint16('1')))
	require.False(t, iswalpha(200))
}
