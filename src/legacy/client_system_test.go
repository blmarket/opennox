package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClientRandomNames(t *testing.T) {
	setGame52Server(t, 123)
	names, initialized := C_clientRandomNames(40)
	require.Equal(t, uint32(33), initialized)
	require.Len(t, names, 40)
	allowed := map[string]bool{
		"Dweezle": true, "Glork": true, "Floogle": true, "Goombah": true,
		"Kraun": true, "Kloog": true, "Zurg": true, "Darg": true,
		"Arfingle": true, "Buurl": true, "Gurgin": true, "Grok": true,
		"Hurlong": true, "Luric": true, "Lupis": true, "Mallik": true,
		"Thrall": true, "Norwood": true, "Nulik": true, "Orin": true,
		"Olaf": true, "Orguk": true, "Pervis": true, "Paavik": true,
		"Qix": true, "Xevin": true, "Xurcon": true, "Markoan": true,
		"Yuric": true, "Yoovis": true, "Yalek": true, "Zug": true,
		"Zivik": true,
	}
	for _, name := range names {
		require.True(t, allowed[name], name)
	}
}
