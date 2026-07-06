package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestGame1ModifNext(t *testing.T) {
	ret, exp := C_nox_xxx_modifNext_4133C0()
	require.Equal(t, exp, ret)
}

func TestGame1Sub413F60(t *testing.T) {
	require.Equal(t, 5, C_sub_413F60(10, 5))
	require.Equal(t, -3, C_sub_413F60(2, 5))
	require.Equal(t, 0, C_sub_413F60(7, 7))
}

func TestGame1ListNext(t *testing.T) {
	require.Equal(t, uintptr(0), C_sub_4168F0(false))
	require.NotEqual(t, uintptr(0), C_sub_4168F0(true))
	require.Equal(t, uintptr(0), C_sub_416910(false))
	require.NotEqual(t, uintptr(0), C_sub_416910(true))
}

func TestGame1WallSecretNext(t *testing.T) {
	require.Equal(t, uint32(0), C_nox_xxx_wallSecretNext_410790(false, 0))
	require.Equal(t, uint32(42), C_nox_xxx_wallSecretNext_410790(true, 42))
}

func TestGame1MapLoadFlag(t *testing.T) {
	old := C_nox_server_gameDoSwitchMap_40A680()
	t.Cleanup(func() {
		C_nox_server_gameUnsetMapLoad_40A690()
		if old != 0 {
			C_nox_server_gameSettingsUpdated_40A670()
		}
	})

	C_nox_server_gameUnsetMapLoad_40A690()
	require.Equal(t, 0, C_nox_server_gameDoSwitchMap_40A680())

	C_nox_server_gameSettingsUpdated_40A670()
	require.Equal(t, 1, C_nox_server_gameDoSwitchMap_40A680())
}

func TestGame1PauseFlagChecks(t *testing.T) {
	old := noxflags.GetGame()
	t.Cleanup(func() {
		noxflags.ResetGame()
		noxflags.SetGame(old)
	})
	noxflags.ResetGame()

	require.Equal(t, 0, C_nox_xxx_checkGameFlagPause_413A50())
	require.Equal(t, 0, C_sub_40A740_withSettingsByte(0x80))

	noxflags.SetGame(noxflags.GamePause)
	require.Equal(t, 1, C_nox_xxx_checkGameFlagPause_413A50())

	noxflags.ResetGame()
	noxflags.SetGame(0x80)
	require.Equal(t, 0, C_sub_40A740_withSettingsByte(0x7f))
	require.Equal(t, 1, C_sub_40A740_withSettingsByte(0x80))
}

func TestGame1ServerSettingsByteMask(t *testing.T) {
	require.Equal(t, byte(0xef), C_sub_4169F0_resultByte(0xff))
	require.Equal(t, byte(0x01), C_sub_4169F0_resultByte(0x11))
}

func TestGame1GetServerSubFlags(t *testing.T) {
	old := C_game1Globals()
	t.Cleanup(func() { C_game1SetGlobals(old) })

	C_game1SetGlobals(Game1Globals{v3484: 0x12345678, v251744: old.v251744})
	require.Equal(t, 0x12345678, C_nox_xxx_getServerSubFlags_409E60())

	C_game1SetGlobals(Game1Globals{v3484: 0, v251744: old.v251744})
	require.Equal(t, 0, C_nox_xxx_getServerSubFlags_409E60())
}

func TestGame1Sub40A6B0(t *testing.T) {
	// sub_40A6B0 returns *getMemU32Ptr(0x5D4594, 3588)
	// Just verify it doesn't crash and returns an int
	_ = C_sub_40A6B0()
}

func TestGame1RateGet(t *testing.T) {
	// nox_xxx_rateGet_40A6C0 returns *getMemU32Ptr(0x587000, 4728)
	_ = C_nox_xxx_rateGet_40A6C0()
}

func TestGame1Sub4139B0(t *testing.T) {
	old := C_game1Globals()
	t.Cleanup(func() { C_game1SetGlobals(old) })

	C_game1SetGlobals(Game1Globals{v3484: old.v3484, v251744: 0})
	require.Equal(t, 0, C_sub_4139B0())

	C_game1SetGlobals(Game1Globals{v3484: old.v3484, v251744: 1})
	require.Equal(t, 1, C_sub_4139B0())

	C_game1SetGlobals(Game1Globals{v3484: old.v3484, v251744: 0xFFFFFFFF})
	require.Equal(t, 1, C_sub_4139B0())
}

func TestGame1Sub409B50And409B80(t *testing.T) {
	// sub_409B50 copies string to memory and returns length (including null terminator)
	// sub_409B80 returns pointer to that memory
	testCases := []struct {
		input    string
		expected uint
	}{
		{"", 1},
		{"a", 2},
		{"hello", 6},
		{"test string", 12},
	}
	for _, tc := range testCases {
		result := C_sub_409B50(tc.input)
		require.Equal(t, tc.expected, result)
		// Verify sub_409B80 returns the same string
		got := C_sub_409B80()
		require.Equal(t, tc.input, got)
	}
}

func TestGame1Sub40A6A0(t *testing.T) {
	// sub_40A6A0 takes an int and returns something based on memory
	// Just verify it doesn't crash for various inputs
	_ = C_sub_40A6A0(0)
	_ = C_sub_40A6A0(1)
	_ = C_sub_40A6A0(-1)
	_ = C_sub_40A6A0(999)
}

func TestGame1Sub409E40(t *testing.T) {
	// sub_409E40 takes an int and returns a value
	_ = C_sub_409E40(0)
	_ = C_sub_409E40(1)
	_ = C_sub_409E40(5)
}

func TestGame1Sub409E70(t *testing.T) {
	// sub_409E70 takes an int and returns a value
	_ = C_sub_409E70(0)
	_ = C_sub_409E70(1)
	_ = C_sub_409E70(10)
}

func TestGame1Sub42EBA0(t *testing.T) {
	// sub_42EBA0 returns *getMemU32Ptr(0x5D4594, 754052)
	_ = C_sub_42EBA0()
}

func TestGame1NoxXxxCursor(t *testing.T) {
	// nox_xxx_cursor_430B00 returns nox_xxx_useAudio_587000_80772
	_ = C_nox_xxx_cursor_430B00()
}

func TestGame1Sub431370(t *testing.T) {
	// sub_431370 returns sub_488B60() != 0
	_ = C_sub_431370()
}
