package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestGameAdditionalGAME1(t *testing.T) {
	// Test sub_409A70 with various inputs - simple lookup function
	_ = C_sub_409A70(0)
	_ = C_sub_409A70(1)
	_ = C_sub_409A70(-1)
	_ = C_sub_409A70(100)

	// Test map filename and name functions (may return empty or default values)
	_ = C_nox_server_currentMapGetFilename_409B30()
	_ = C_nox_xxx_mapGetMapName_409B40()

	// Test sub_409EF0, sub_409F40 with various inputs - simple functions
	_ = C_sub_409EF0(0)
	_ = C_sub_409EF0(1)

	_ = C_sub_409F40(0)
	_ = C_sub_409F40(1)

	// Test player limit - simple getter
	_ = C_nox_xxx_servGetPlrLimit_409FA0()

	// Test sub_40A220 - simple getter
	_ = C_sub_40A220()

	// Test sub_40AA40 - simple getter
	_ = C_sub_40AA40()

	// Test sub_40E090 (void function) - simple setter
	C_sub_40E090()

	// Test sub_413920 - simple getter
	_ = C_sub_413920()
}

func TestGameAdditionalGAME2(t *testing.T) {
	// Test sub_452010, sub_4521F0 - simple functions
	_ = C_sub_452010()
	_ = C_sub_4521F0()
}

func TestGameAdditionalGAME3(t *testing.T) {
	// GAME3 tests - functions are tested via other test files
}

func TestGameAdditionalGAME4(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()

	got := C_game4SafePaths()
	require.Zero(t, got.initResult)
	require.Equal(t, -1, got.missingName)
	require.Zero(t, got.negativeIndex)
	require.NotZero(t, got.zeroIndex)
	require.Zero(t, got.count)
	require.Equal(t, 1, got.setPrimary)
	require.Zero(t, got.clearPrimary)
	require.Equal(t, 1, got.setSecondary)
	require.Zero(t, got.clearSecondary)
	require.Zero(t, got.invalidSave)
	require.Zero(t, got.closeEmpty)
	require.Zero(t, got.seekEmpty)
	require.Equal(t, -1.0, got.invalidFirstCoord)
	require.Equal(t, -1.0, got.invalidLastCoord)
	require.Equal(t, 1, got.moveEmpty)
	require.Equal(t, 1, got.placeEmpty)
	require.Equal(t, 77, got.objectField)
	require.Equal(t, 88, got.nodeField)
	require.Zero(t, got.removeNil)
	require.Zero(t, got.removeMissing)
	require.Zero(t, got.voteGuard)
	require.Equal(t, 1, got.voteThreshold)
	require.False(t, got.votesActive)
}

func TestGame4SpellPhonemeForNonPlayerUnit(t *testing.T) {
	h := newGameLogicHarness(t)
	srv := h.setServer(0, 30)
	h.setGameFlags(noxflags.GameHost)
	obj := h.point(0, 0)
	obj.NetCode = 0x1234
	srv.Objs.SetObjects(obj)

	tests := []struct {
		phoneme int8
		want    int
	}{
		{phoneme: 0, want: 193},
		{phoneme: 1, want: 186},
		{phoneme: 2, want: 187},
		{phoneme: 3, want: 192},
		{phoneme: 4, want: 0},
		{phoneme: 5, want: 188},
		{phoneme: 6, want: 191},
		{phoneme: 7, want: 190},
		{phoneme: 8, want: 189},
		{phoneme: -1, want: 0},
	}
	for _, tc := range tests {
		require.Equal(t, tc.want, C_game4SpellPhoneme(int(obj.NetCode), tc.phoneme))
	}
}

func TestGameAdditionalGAME5(t *testing.T) {
	// GAME5 tests - functions are tested via other test files
}

func TestGameAdditionalWrappers(t *testing.T) {
	// Test various simple wrapper functions that just return memory values
	// These should not crash

	// GAME1 wrappers
	_ = C_sub_40A220()
	_ = C_sub_413920()
}
