package sound

import (
	"testing"
)

func TestByName_Extra2(t *testing.T) {
	id := ByName("AnchorCast")
	if id != SoundAnchorCast {
		t.Errorf("ByName(AnchorCast) = %v, want %v", id, SoundAnchorCast)
	}
	id = ByName("NonExistentSoundXYZ")
	if id != 0 {
		t.Errorf("ByName non-existent should return 0, got %v", id)
	}
	id = ByName("")
	if id != 0 {
		t.Errorf("ByName empty should return 0, got %v", id)
	}
}

func TestID_String_Extra2(t *testing.T) {
	s := SoundNone.String()
	if s == "" {
		t.Error("SoundNone.String() should not be empty")
	}
	// Test out of range ID - should panic or return formatted string
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Out of range ID.String() panicked as expected: %v", r)
			}
		}()
		id := ID(99999)
		s := id.String()
		if s == "" {
			t.Error("Out of range ID.String() should not be empty")
		}
	}()
	// Test negative ID
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Negative ID.String() panicked as expected: %v", r)
			}
		}()
		id := ID(-1)
		s := id.String()
		if s == "" {
			t.Error("Negative ID.String() should not be empty")
		}
	}()
}

func TestID_GoString_Extra2(t *testing.T) {
	if SoundNone.GoString() != "sound.SoundNone" {
		t.Errorf("SoundNone.GoString() = %q, want %q", SoundNone.GoString(), "sound.SoundNone")
	}
	s := SoundAnchorCast.GoString()
	if s != "sound.SoundAnchorCast" {
		t.Errorf("SoundAnchorCast.GoString() = %q, want %q", s, "sound.SoundAnchorCast")
	}
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("Out of range ID.GoString() panicked as expected: %v", r)
			}
		}()
		id := ID(99999)
		s = id.GoString()
		if s == "" {
			t.Error("Out of range ID.GoString() should not be empty")
		}
	}()
}

func TestAllSounds_Extra2(t *testing.T) {
	for id, name := range soundNames {
		if name == "" {
			continue
		}
		gotID := ByName(name)
		if gotID != ID(id) {
			t.Errorf("ByName(%q) = %v, want %v", name, gotID, id)
		}
		sid := ID(id)
		if sid.String() != name {
			t.Errorf("ID(%d).String() = %q, want %q", id, sid.String(), name)
		}
	}
}
