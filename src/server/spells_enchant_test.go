package server

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/spell"
	"github.com/stretchr/testify/require"
)

func TestParseEnchant(t *testing.T) {
	// Test valid enchant names
	id, ok := ParseEnchant("ENCHANT_INVISIBLE")
	require.True(t, ok)
	require.Equal(t, ENCHANT_INVISIBLE, id)

	id, ok = ParseEnchant("ENCHANT_HASTED")
	require.True(t, ok)
	require.Equal(t, ENCHANT_HASTED, id)

	// Test invalid name
	id, ok = ParseEnchant("INVALID_ENCHANT")
	require.False(t, ok)
}

func TestEnchantID_String(t *testing.T) {
	// Test valid enchant IDs
	require.Equal(t, "ENCHANT_INVISIBLE", ENCHANT_INVISIBLE.String())
	require.Equal(t, "ENCHANT_HASTED", ENCHANT_HASTED.String())
	require.Equal(t, "ENCHANT_SNEAK", ENCHANT_SNEAK.String())

	// Test invalid ID
	invalid := EnchantID(255)
	require.Equal(t, "EnchantID(255)", invalid.String())
}

func TestEnchantID_Spell(t *testing.T) {
	// Test enchant to spell mapping
	require.Equal(t, spell.SPELL_INVISIBILITY, ENCHANT_INVISIBLE.Spell())
	require.Equal(t, spell.SPELL_HASTE, ENCHANT_HASTED.Spell())
	require.Equal(t, spell.SPELL_VAMPIRISM, ENCHANT_VAMPIRISM.Spell())

	// Test enchants with no spell (0)
	require.Equal(t, spell.ID(0), ENCHANT_DETECTING.Spell())
	require.Equal(t, spell.ID(0), ENCHANT_ETHEREAL.Spell())
}
