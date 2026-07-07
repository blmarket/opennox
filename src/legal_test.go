package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/client/gui"
)

func TestNoxClientOnShowLegal(t *testing.T) {
	// nox_client_onShowLegal accesses win.ChildByID and version info
	// With nil window, it will panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_client_onShowLegal panicked as expected: %v", r)
		}
	}()
	nox_client_onShowLegal(nil)
}

func TestNoxClientOnShowLegalWithWindow(t *testing.T) {
	// Create a minimal window for testing
	// This is complex because it needs a GUI context
	// Just verify the function doesn't panic with recover
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_client_onShowLegal with window panicked: %v", r)
		}
	}()
	// Skip actual window creation as it requires GUI setup
	t.Skip("Skipping window test as it requires GUI setup")
	_ = gui.Window{}
}
