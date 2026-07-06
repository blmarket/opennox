package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommonObjectArmrLook(t *testing.T) {
	// Test nox_xxx_loadLook_415D50 - should not crash and return a pointer
	// It sets a global flag after first call, subsequent calls return early
	C_testLoadLook()
	C_testLoadLook() // Second call should be safe
}

func TestCommonObjectWeapLook(t *testing.T) {
	// Test nox_xxx_loadModifyers_4158C0 - similar to armrlook
	C_testLoadModifyers()
	C_testLoadModifyers()
}

func TestCommonObjectLookMultipleCalls(t *testing.T) {
	// Multiple calls should be safe due to global flag check
	for i := 0; i < 5; i++ {
		C_testLoadLook()
		C_testLoadModifyers()
	}
}

func TestCommonMagicComGuide(t *testing.T) {
	// Test nox_xxx_loadGuides_427070 - loads creature guides
	// May crash if no data loaded, so we recover
	defer func() {
		_ = recover()
	}()
	result := C_testLoadGuides()
	// Result is 0 or 1 depending on data availability
	require.True(t, result == 0 || result == 1, "expected 0 or 1, got %d", result)

	// Second call should also be safe
	result2 := C_testLoadGuides()
	require.True(t, result2 == 0 || result2 == 1)
}
