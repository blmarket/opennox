package opennox

import (
	"testing"
)

func TestFlushCoverage_Extra(t *testing.T) {
	// Test flushCoverage doesn't panic
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("flushCoverage panicked: %v", r)
			}
		}()
		flushCoverage()
	}()
}

func TestFlushCoverage_MultipleCalls_Extra(t *testing.T) {
	// Multiple calls should not panic
	for i := 0; i < 3; i++ {
		flushCoverage()
	}
}
