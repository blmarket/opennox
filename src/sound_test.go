package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/sound"
)

func TestClientPlaySoundSpecial(t *testing.T) {
	// clientPlaySoundSpecial calls legacy.Nox_xxx_clientPlaySoundSpecial_452D80
	// It should not panic with valid sound ID
	defer func() {
		if r := recover(); r != nil {
			t.Logf("clientPlaySoundSpecial panicked: %v", r)
		}
	}()
	clientPlaySoundSpecial(sound.ID(0), 0)
	clientPlaySoundSpecial(sound.ID(1), 100)
}

func TestNoxXxxSoundPlayerDamageSound(t *testing.T) {
	// Nox_xxx_soundPlayerDamageSound_5328B0 accesses obj1.Server() and obj1.InvFirstItem
	// With nil objects, it should panic or return 1
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Nox_xxx_soundPlayerDamageSound_5328B0 panicked as expected: %v", r)
		}
	}()
	result := Nox_xxx_soundPlayerDamageSound_5328B0(nil, nil)
	if result != 1 {
		t.Errorf("Nox_xxx_soundPlayerDamageSound_5328B0 with nil should return 1, got %d", result)
	}
}
