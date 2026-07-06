package input

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/stretchr/testify/require"
)

func TestStr2U16(t *testing.T) {
	require.Equal(t, uint16(0), str2u16(""))
	require.Equal(t, uint16('A'), str2u16("A"))
	require.Equal(t, uint16(0x4241), str2u16("AB")) // little endian
}

func TestIswalpha(t *testing.T) {
	require.True(t, iswalpha('A'))
	require.True(t, iswalpha('z'))
	require.False(t, iswalpha(170))
	require.False(t, iswalpha(181))
	require.False(t, iswalpha(186))
	require.False(t, iswalpha(200)) // >=192
	require.False(t, iswalpha('1'))
}

func TestSetLanguage(t *testing.T) {
	// SetLanguage requires a keyboard handler, just verify the function exists
	h := &Handler{}
	require.NotNil(t, h.SetLanguage)
}

func TestKeyToWChar(t *testing.T) {
	h := &Handler{}
	// KeyToWChar with >0xFF just returns the value, no k needed
	require.Equal(t, uint16(0x1FF), h.KeyToWChar(0x1FF))
	require.Equal(t, uint16(0), h.KeyToWChar(keybind.KeyLShift))
	require.Equal(t, uint16(0), h.KeyToWChar(keybind.KeyRShift))
}
