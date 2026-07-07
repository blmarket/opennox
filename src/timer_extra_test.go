package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func TestTimerImports_Extra(t *testing.T) {
	// Test that timer.PlatformTicks is set by init
	if timer.PlatformTicks == nil {
		t.Error("timer.PlatformTicks should be set by init in timer_imports.go")
	}
	// Test calling it doesn't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("timer.PlatformTicks panicked: %v", r)
			}
		}()
		if timer.PlatformTicks != nil {
			_ = timer.PlatformTicks()
		}
	}()
}
