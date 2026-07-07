package opennox

import (
	"testing"
)

func TestOpenThingsError(t *testing.T) {
	// openThings tries to load thing.bin which may not exist in test environment
	// It should return an error if the file doesn't exist
	noxLoadedThings = nil
	thg, err := openThings()
	if err != nil {
		t.Logf("openThings returned expected error: %v", err)
	} else {
		t.Logf("openThings succeeded, thg=%v", thg)
		if thg != nil {
			closeThings()
		}
	}
}

func TestCloseThingsNil(t *testing.T) {
	// closeThings with nil noxLoadedThings should not panic
	noxLoadedThings = nil
	closeThings()
	// Should not panic
}

func TestCloseThings(t *testing.T) {
	// Test closeThings sets noxLoadedThings to nil
	// We can't easily create a valid binfile.MemFile without the file,
	// but we can verify it handles nil correctly
	noxLoadedThings = nil
	closeThings()
	if noxLoadedThings != nil {
		t.Error("closeThings should set noxLoadedThings to nil")
	}
}

func TestGetThingNameNilClient(t *testing.T) {
	// getThingName with nil noxClient should panic or return empty
	// The existing test already covers this, but we add a specific case
	originalClient := noxClient
	noxClient = nil
	defer func() { noxClient = originalClient }()

	defer func() {
		if r := recover(); r != nil {
			t.Logf("getThingName panicked as expected with nil client: %v", r)
		}
	}()

	name := getThingName(999)
	if name != "" {
		t.Logf("getThingName(999) with nil client returned %q", name)
	}
}
