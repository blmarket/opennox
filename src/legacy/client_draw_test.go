package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientDrawWrappersSkipDrawing(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setServer(0, 30)

	require.Equal(t, 1, C_nox_thing_static_draw_skip())
	require.Equal(t, 1, C_nox_thing_weapon_draw_skip())
	require.Equal(t, 1, C_nox_thing_armor_draw_skip())
	require.Equal(t, 1, C_nox_thing_base_draw_skip())
}
