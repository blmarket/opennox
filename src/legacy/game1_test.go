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
