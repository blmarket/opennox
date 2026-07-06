package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/opennox/v1/client"
)

// thingCheckClient is a minimal Client stub used by TestGame2ThingChecks. Only
// Cli() is called (via nox_get_thing); the embedded nil Client satisfies the
// rest of the interface without them ever being invoked.
type thingCheckClient struct {
	Client
	cli *client.Client
}

func (c *thingCheckClient) Cli() *client.Client { return c.cli }

// TestGame2ConstReturns covers pure constant-return helpers in GAME2.c.
func TestGame2ConstReturns(t *testing.T) {
	require.Equal(t, 1, C_sub_44E8D0())
	require.Equal(t, 1, C_nox_xxx_quickbarButtonBookDraw_45EF30())
	require.Equal(t, 0, C_sub_45EF40())
}

// TestGame2Quickbar covers nox_xxx_quickbar_45F8D0 pure logic on a2.
func TestGame2Quickbar(t *testing.T) {
	cases := []struct {
		a2   int
		want int
	}{
		{8, 0},
		{12, 0},
		{16, 0},
		{0, 1},
		{7, 1},
		{9, 1},
		{15, 1},
		{17, 1},
	}
	for _, c := range cases {
		require.Equal(t, c.want, C_nox_xxx_quickbar_45F8D0(0, c.a2, 0, 0))
	}
}

// TestGame2Bitmask454000 covers sub_454000 bit test on caller buffer.
func TestGame2Bitmask454000(t *testing.T) {
	require.Equal(t, 1, C_sub_454000(5, true))
	require.Equal(t, 0, C_sub_454000(5, false))
	require.Equal(t, 1, C_sub_454000(33, true))
	require.Equal(t, 0, C_sub_454000(100, false))
}

// TestGame2DrawableField104 covers sub_45A010 returning dr->field_104.
func TestGame2DrawableField104(t *testing.T) {
	ret, exp := C_sub_45A010()
	require.Equal(t, exp, ret)
}

// TestGame2Bit45F500 covers sub_45F500 bit extraction.
func TestGame2Bit45F500(t *testing.T) {
	require.Equal(t, 1, C_sub_45F500(0, true))
	require.Equal(t, 0, C_sub_45F500(0, false))
	require.Equal(t, 1, C_sub_45F500(3, true))
}

// TestGame2Write4526D0 covers sub_4526D0 writing 4 to nested offset.
func TestGame2Write4526D0(t *testing.T) {
	ret, written := C_sub_4526D0()
	require.Equal(t, 0, ret)
	require.Equal(t, uint32(4), written)
}

// TestGame2Flags459DB0 covers sub_459DB0 flag checks at +112 and +116.
func TestGame2Flags459DB0(t *testing.T) {
	require.Equal(t, 1, C_sub_459DB0(0x400000, 8))
	require.Equal(t, 0, C_sub_459DB0(0, 8))
	require.Equal(t, 0, C_sub_459DB0(0x400000, 0))
	require.Equal(t, 0, C_sub_459DB0(0, 0))
}

// TestGame2SpriteSetActive covers nox_xxx_spriteSetActiveMB_45A990_drawable setting bit.
func TestGame2SpriteSetActive(t *testing.T) {
	ret, flags := C_nox_xxx_spriteSetActiveMB_45A990_drawable()
	require.NotEqual(t, 0, ret)
	require.Equal(t, uint32(4), flags&4)
}

// TestGame2ThingChecks covers sub_44D040, sub_44D060, sub_44D090 with invalid thing IDs.
// These functions call nox_get_thing which returns 0 for invalid IDs, so they should return 0.
func TestGame2ThingChecks(t *testing.T) {
	// nox_get_thing (called by these C funcs) dereferences GetClient().Cli();
	// without a client installed it segfaults. Install a stub whose empty
	// client.Client makes Things.TypeByInd return nil for any invalid index.
	old := GetClient
	t.Cleanup(func() { GetClient = old })
	GetClient = func() Client { return &thingCheckClient{cli: &client.Client{}} }

	t.Run("sub_44D040 returns 0 for invalid thing", func(t *testing.T) {
		require.Equal(t, 0, C_sub_44D040(0))
		require.Equal(t, 0, C_sub_44D040(-1))
		require.Equal(t, 0, C_sub_44D040(999999))
	})
	t.Run("sub_44D060 returns 0 for invalid thing", func(t *testing.T) {
		require.Equal(t, 0, C_sub_44D060(0))
		require.Equal(t, 0, C_sub_44D060(-1))
		require.Equal(t, 0, C_sub_44D060(999999))
	})
	t.Run("sub_44D090 returns 0 for invalid thing", func(t *testing.T) {
		require.Equal(t, 0, C_sub_44D090(0))
		require.Equal(t, 0, C_sub_44D090(-1))
		require.Equal(t, 0, C_sub_44D090(999999))
	})
}

// TestGame2GlobalVars covers sub_44D960, sub_44D970, sub_44D990 which manipulate global variables.
func TestGame2GlobalVars(t *testing.T) {
	old := C_game2Globals()
	t.Cleanup(func() { C_game2SetGlobals(old) })

	t.Run("sub_44D960 sets v122848 to 0", func(t *testing.T) {
		C_game2SetGlobals(Game2Globals{v831092: 0, v122848: 123, v831220: 0})
		C_sub_44D960()
		require.Equal(t, 0, C_sub_44D990())
	})

	t.Run("sub_44D970 returns v831092 and sets v122848", func(t *testing.T) {
		C_game2SetGlobals(Game2Globals{v831092: 0, v122848: 0, v831220: 0})
		require.Equal(t, 0, C_sub_44D970())
		require.Equal(t, 0, C_sub_44D990())

		C_game2SetGlobals(Game2Globals{v831092: 42, v122848: 0, v831220: 0})
		require.Equal(t, 42, C_sub_44D970())
		require.Equal(t, 1, C_sub_44D990())
	})

	t.Run("sub_44D990 returns v122848", func(t *testing.T) {
		C_game2SetGlobals(Game2Globals{v831092: 0, v122848: 77, v831220: 0})
		require.Equal(t, 77, C_sub_44D990())
	})
}

// TestGame2Sub44E8B0 covers sub_44E8B0 which returns 1.0 if v831220 == 255 else 0.0.
func TestGame2Sub44E8B0(t *testing.T) {
	old := C_game2Globals()
	t.Cleanup(func() { C_game2SetGlobals(old) })

	C_game2SetGlobals(Game2Globals{v831092: 0, v122848: 0, v831220: 255})
	require.Equal(t, 1.0, C_sub_44E8B0())

	C_game2SetGlobals(Game2Globals{v831092: 0, v122848: 0, v831220: 0})
	require.Equal(t, 0.0, C_sub_44E8B0())

	C_game2SetGlobals(Game2Globals{v831092: 0, v122848: 0, v831220: 100})
	require.Equal(t, 0.0, C_sub_44E8B0())
}

// TestGame2WndProc covers nox_xxx_wndProc_44E6E0 which returns a2 == 23.
func TestGame2WndProc(t *testing.T) {
	require.Equal(t, 1, C_nox_xxx_wndProc_44E6E0(0, 23, 0, 0))
	require.Equal(t, 0, C_nox_xxx_wndProc_44E6E0(0, 0, 0, 0))
	require.Equal(t, 0, C_nox_xxx_wndProc_44E6E0(0, 22, 0, 0))
	require.Equal(t, 0, C_nox_xxx_wndProc_44E6E0(0, 24, 0, 0))
	require.Equal(t, 1, C_nox_xxx_wndProc_44E6E0(123, 23, 456, 789))
}
