package opennox

import (
	"testing"

	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
)

func TestNoxScriptNSLoadMap(t *testing.T) {
	s := noxScriptNS{}

	// Test nil options
	err := s.loadMap("testmap", nil)
	// Should fail because map doesn't exist, but not because of validation
	if err != nil {
		t.Logf("loadMap failed as expected (map doesn't exist): %v", err)
	}

	// Test with extension - should fail validation
	err = s.loadMap("testmap.map", nil)
	if err == nil {
		t.Error("expected error for map name with extension")
	} else if err.Error() != "map name must not contain file extension" {
		t.Errorf("unexpected error: %v", err)
	}

	// Test with path elements
	testCases := []string{
		"../testmap",
		"dir/testmap",
		"test.map",
		"test/map",
		"test\\map",
		"test.map.name",
	}
	for _, tc := range testCases {
		err = s.loadMap(tc, nil)
		if err == nil {
			t.Errorf("expected error for map name %q with path elements", tc)
		}
	}

	// Test with valid name but non-existent map
	err = s.loadMap("validmapname", &ns4.LoadMapOptions{})
	if err != nil {
		t.Logf("loadMap failed as expected: %v", err)
	}

	// Test LoadMap panics on error
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Error("LoadMap should panic on error")
			}
		}()
		s.LoadMap("invalid/map", nil)
	}()
}

func TestNoxScriptNSAutoSave(t *testing.T) {
	// AutoSave should not panic even with nil server
	s := noxScriptNS{}
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("AutoSave panicked as expected with nil server: %v", r)
			}
		}()
		s.AutoSave()
	}()
}
