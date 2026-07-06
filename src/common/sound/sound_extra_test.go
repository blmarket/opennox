package sound

import (
	"testing"
)

func TestSoundString(t *testing.T) {
	// Test String with valid ID
	s := ID(0).String()
	if s == "" {
		t.Error("String should not be empty")
	}

	// Test GoString with valid ID
	gs := ID(0).GoString()
	if gs == "" {
		t.Error("GoString should not be empty")
	}

	// Test ByName with invalid name
	id := ByName("nonexistent_sound_xyz")
	if id != 0 {
		t.Error("ByName with invalid name should return 0")
	}

	// Test ByName with valid name
	id = ByName("Silent")
	_ = id
}
