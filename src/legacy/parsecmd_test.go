package legacy

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/console"
	"github.com/stretchr/testify/require"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestParseCmdArgumentGuards(t *testing.T) {
	require.Equal(t, make([]int, 19), C_parsecmdInvalidCounts())
	require.Equal(t, 1, C_parsecmdNameWithoutArguments())
}

func TestParseCmdNilPlayerHelpers(t *testing.T) {
	h := newGameLogicHarness(t)
	t.Cleanup(noxflags.ResetGame)
	h.setGameFlags(0)
	got := C_parsecmdNilPlayerHelpers()
	require.Equal(t, [7]int{}, got)

	h.setGameFlags(noxflags.GameClient)
	got = C_parsecmdNilPlayerHelpers()
	require.Equal(t, [7]int{}, got)
}

func TestParseCmdSafeStateCommands(t *testing.T) {
	h := newGameLogicHarness(t)
	t.Cleanup(noxflags.ResetGame)
	t.Cleanup(noxflags.ResetEngine)

	set, unset := C_parsecmdEngineDebug()
	require.Equal(t, 1, set)
	require.Equal(t, 1, unset)

	h.setGameFlags(0)
	require.Equal(t, 1, C_parsecmdShowRank())

	h.setGameFlags(noxflags.GameOnline | noxflags.GameFlag4)
	require.Equal(t, [5]int{1, 1, 1, 1, 1}, C_parsecmdSafeOnlineCommands())

	h.setGameFlags(0)
	require.Zero(t, C_parsecmdCheatLevelOffline())
}

func TestParseCmdControlledValidPaths(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	setGame52Server(t, 100)
	t.Cleanup(noxflags.ResetGame)
	oldGetConsole := GetConsole
	quietConsole := console.New(nil)
	GetConsole = func() *console.Console { return quietConsole }
	t.Cleanup(func() { GetConsole = oldGetConsole })

	for _, i := range []int{0, 1, 2, 3, 4, 5, 6, 7, 12, 13, 14, 15, 16} {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			if i == 3 {
				h.setGameFlags(noxflags.GameModeChat)
			} else if i == 5 || i == 6 {
				h.setGameFlags(noxflags.GameClient)
			} else {
				h.setGameFlags(0)
			}
			got := C_parsecmdValidCase(i)
			if i <= 2 {
				require.Zero(t, got)
			} else {
				require.Equal(t, 1, got)
			}
		})
	}
}
