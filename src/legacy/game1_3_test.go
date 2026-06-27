package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestGame13SimpleStateHelpers(t *testing.T) {
	h := newGameLogicHarness(t)
	oldGlobals := C_game13Globals()
	t.Cleanup(func() { C_game13SetGlobals(oldGlobals) })
	h.snapshotMem(0x5D4594, 815092, 4)
	h.snapshotMem(0x5D4594, 816408, 4)
	h.snapshotMem(0x587000, 91880, 8)
	oldPlatformTicks := PlatformTicks
	PlatformTicks = func() uint64 { return 100 }
	t.Cleanup(func() { PlatformTicks = oldPlatformTicks })

	C_game13SetGlobals(game13Globals{menuState: 0x12345678})
	require.Equal(t, 0x12345678, C_sub_43B6D0())

	*memmap.PtrUint32(0x5D4594, 815092) = 0x76543210
	require.Equal(t, 0x76543210, C_sub_43BDB0())

	ret, text := C_nox_sprintAddrPort_43BC80("127.0.0.1", 18590)
	require.Equal(t, len(text), ret)
	require.Equal(t, "127.0.0.1:18590", text)

	require.Equal(t, 91, C_sub_43E8C0(91))
	require.Equal(t, uint32(91), *memmap.PtrUint32(0x5D4594, 816408))

	C_game13SetGlobals(game13Globals{tickCount: 0})
	require.Equal(t, 1, C_sub_43C650())
	C_sub_43CEB0()
	require.Equal(t, uint64(33), *memmap.PtrUint64(0x587000, 91880))
	require.Equal(t, uint32(0), *memmap.PtrUint32(0x587000, 91884))
}

func TestGame13MusicStateHelpersWithoutMusicModule(t *testing.T) {
	oldGlobals := C_game13Globals()
	t.Cleanup(func() { C_game13SetGlobals(oldGlobals) })

	C_game13SetGlobals(game13Globals{musicSlot: 6})
	require.Equal(t, 0, C_sub_43DA80())
	require.Equal(t, uint32(6), C_game13Globals().musicSlot)

	C_game13SetGlobals(game13Globals{musicSlot: 0})
	C_sub_43DAD0()
	require.Equal(t, uint32(0), C_game13Globals().musicSlot)

	require.Equal(t, 4, C_sub_43DB30(4))
	require.Equal(t, 4, C_sub_43DB20())

	C_game13SetGlobals(game13Globals{musicDepth: 3})
	require.Equal(t, 3, C_sub_43DB60())
	require.Equal(t, uint32(3), C_game13Globals().musicDepth)

	C_game13SetGlobals(game13Globals{musicDepth: 0, musicSlot: 5})
	C_sub_43DBA0()
	require.Equal(t, uint32(0), C_game13Globals().musicDepth)
	require.Equal(t, uint32(5), C_game13Globals().musicSlot)

	C_game13SetGlobals(game13Globals{musicReady: 7, musicFlag: 0})
	require.Equal(t, 7, C_sub_43DC10())
	require.Equal(t, uint32(1), C_game13Globals().musicFlag)
}

func TestGame13AudioBranchClassifier(t *testing.T) {
	require.Equal(t, 0, C_sub_43F0E0(1, 0, 0))
	require.Equal(t, 7, C_sub_43F0E0(2, 2, 0))
	require.Equal(t, 5, C_sub_43F0E0(2, 3, 0))
	require.Equal(t, 3, C_sub_43F0E0(0, 2, 2))
	require.Equal(t, 1, C_sub_43F0E0(0, 0, 2))
}

func TestGame13GuiCheckWithoutQuitMenu(t *testing.T) {
	require.Equal(t, 0, C_nox_gui_xxx_check_446360())
}
