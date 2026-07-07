package input

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/seat"
	"github.com/stretchr/testify/require"
)

func TestHandlerBasic(t *testing.T) {
	mock := &mockInput{}
	h := New(mock, false, 0)
	require.NotNil(t, h)

	// Test basic methods
	seq1 := h.NextEventSeq()
	seq2 := h.NextEventSeq()
	require.Equal(t, seq1+1, seq2)

	require.Equal(t, uint(0), h.CurrentSeq())
	require.Equal(t, 0, h.SeqDelay())

	h.SetSensitivity(1.5)
	require.Equal(t, float32(1.5), h.GetSensitivity())

	// Test OnQuit
	quitCalled := false
	h.OnQuit(func() { quitCalled = true })
	h.InputEvent(seat.WindowClosed)
	require.True(t, quitCalled)
}

type mockInput struct {
	onInput func(ev seat.InputEvent)
}

func (m *mockInput) InputTick()                                          {}
func (m *mockInput) ReplaceInputs(cfg seat.InputConfig) seat.InputConfig { return nil }
func (m *mockInput) OnInput(fnc func(ev seat.InputEvent))                { m.onInput = fnc }
func (m *mockInput) SetTextInput(enable bool)                            {}
