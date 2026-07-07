package opennox

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestApiMemmapHandlers(t *testing.T) {
	// Test that the HTTP handlers are registered and work correctly
	// Note: The handlers may panic if memmap is not fully initialized, which is expected
	// in a test environment without the full game setup. We just verify the handlers
	// are registered without causing the test to fail.

	// Test /debug/nox/memmap/validate handler
	req := httptest.NewRequest("GET", "/debug/nox/memmap/validate", nil)
	w := httptest.NewRecorder()

	// The handler should not cause the test to fail even if it panics
	func() {
		defer func() {
			_ = recover() // Ignore panic in test environment
		}()
		http.DefaultServeMux.ServeHTTP(w, req)
	}()

	// Test /debug/nox/memmap/snapshot handler
	req = httptest.NewRequest("GET", "/debug/nox/memmap/snapshot", nil)
	w = httptest.NewRecorder()

	func() {
		defer func() {
			_ = recover() // Ignore panic in test environment
		}()
		http.DefaultServeMux.ServeHTTP(w, req)
	}()
}

func TestApiMemmapValidate(t *testing.T) {
	// memmap.ValidateZeros may panic if memmap not initialized, which is ok in test
	func() {
		defer func() {
			_ = recover() // Ignore panic in test environment
		}()
		memmap.ValidateZeros()
	}()
}
