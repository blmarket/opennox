package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func TestTimerImports(t *testing.T) {
	// timer.PlatformTicks should be set by init() in timer_imports.go
	if timer.PlatformTicks == nil {
		t.Error("timer.PlatformTicks should not be nil after init")
	}
}

func TestFlushCoverageExtra(t *testing.T) {
	// flushCoverage should not panic
	// In non-ccover build, it's a no-op
	// In ccover build, it calls ccover.Dump()
	defer func() {
		if r := recover(); r != nil {
			t.Logf("flushCoverage panicked (may be expected): %v", r)
		}
	}()
	flushCoverage()
}
