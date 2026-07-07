package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNoxScriptNSUnknown(t *testing.T) {
	var ns NoxScriptNS

	// Test all unused functions - they should not panic
	require.NotPanics(t, func() {
		ns.Unused1f(1)
		ns.Unused20(2)
		ns.Unused50()
		ns.Unused58(1, 2)
		ns.Unused59(3, 4)
		ns.Unused5a(5, 6)
		ns.Unused5b(7, 8)
		ns.Unused5c(9, 10)
		ns.Unused5d(11, 12)
	})
}
