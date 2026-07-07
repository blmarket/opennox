package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func TestTimerImportsInitExtra(t *testing.T) {
	// The init function in timer_imports.go should set timer.PlatformTicks to platformTicks
	if timer.PlatformTicks == nil {
		t.Error("timer.PlatformTicks should be set by init function in timer_imports.go")
	}

	// Call the function to ensure it doesn't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("timer.PlatformTicks panicked: %v", r)
			}
		}()
		if timer.PlatformTicks != nil {
			_ = timer.PlatformTicks()
		}
	}()
}

func TestPlatformTicksExtra(t *testing.T) {
	// platformTicks function should return a uint32 without panicking
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("platformTicks panicked: %v", r)
			}
		}()
		result := platformTicks()
		_ = result // Result is a uint32 timestamp, just verify it doesn't panic
	}()
}
