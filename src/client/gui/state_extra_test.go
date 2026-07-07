package gui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestState_Extra(t *testing.T) {
	// Reset state map for clean test
	stateByID = make(map[StateID]*state)

	// Test RegisterState panics
	require.Panics(t, func() {
		RegisterState(0, "bad", func() bool { return true })
	})
	require.Panics(t, func() {
		RegisterState(-1, "bad", func() bool { return true })
	})
	require.Panics(t, func() {
		RegisterState(1, "bad", nil)
	})

	RegisterState(1, "state1", func() bool { return true })
	require.Panics(t, func() {
		RegisterState(1, "duplicate", func() bool { return true })
	})

	// Test String
	require.Equal(t, "<none>", StateNone.String())
	require.Equal(t, "state1", StateID(1).String())
	require.Equal(t, "gui.StateID(999)", StateID(999).String())

	// Test State stack
	var st State
	require.Equal(t, StateNone, st.Current())

	require.True(t, st.Push(1))
	require.Equal(t, StateID(1), st.Current())
	require.False(t, st.Push(1)) // already current

	require.True(t, st.Push(2))
	require.Equal(t, StateID(2), st.Current())

	st.Pop()
	require.Equal(t, StateID(1), st.Current())

	st.PopUntil(StateID(1))
	require.Equal(t, StateID(1), st.Current())

	st.PopUntil(StateNone)
	require.Equal(t, StateNone, st.Current())

	// Pop empty should not panic
	st.Pop()
	st.Pop()

	// Test Switch
	RegisterState(2, "state2", func() bool { return false })
	st.Push(2)
	require.False(t, st.Switch())

	st.Push(3) // not registered
	require.True(t, st.Switch())

	// Test Switch with nil func
	stateByID[4] = &state{Name: "nilfunc", Func: nil}
	st.Push(4)
	require.True(t, st.Switch())
}
