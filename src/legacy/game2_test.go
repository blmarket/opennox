package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
