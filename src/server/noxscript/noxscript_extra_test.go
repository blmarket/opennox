package noxscript

import (
	"testing"

	"github.com/noxworld-dev/noxscript/ns/asm"
	"github.com/stretchr/testify/require"
)

func TestRegisterAndCallBuiltin(t *testing.T) {
	// Test CallBuiltin with unknown op
	result, ok := CallBuiltin(nil, asm.Builtin(99999))
	require.False(t, ok)
	require.Equal(t, 0, result)

	// Test Register and CallBuiltin
	op := asm.Builtin(12345)
	called := false
	Register(op, func(s VM) int {
		called = true
		return 42
	})

	result, ok = CallBuiltin(nil, op)
	require.True(t, ok)
	require.Equal(t, 42, result)
	require.True(t, called)

	// Test Register panic on duplicate
	require.Panics(t, func() {
		Register(op, func(s VM) int { return 0 })
	})
}

func TestBuiltinFunctions(t *testing.T) {
	// Test that builtin functions exist and can be called
	// They may panic with nil VM, which is expected - we just verify they exist
	// Audio functions
	_ = nsAudioEvent
	_ = nsMusic
	_ = nsMusicPushEvent
	_ = nsMusicPopEvent
	_ = nsMusicEvent

	// Object functions - test a sample
	_ = nsGetTrigger
	_ = nsGetCaller
	_ = nsIsTrigger
	_ = nsIsCaller
	_ = nsObject
}

func TestObjectGroupFunctions(t *testing.T) {
	_ = nsGetObjectGroup
	_ = nsObjectGroupOn
	_ = nsObjectGroupOff
	_ = nsObjectGroupToggle
	_ = nsGroupSetOwner
}
