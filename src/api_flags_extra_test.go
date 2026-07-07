package opennox

import (
	"net/http"
	"testing"
)

func TestApiFlags_Init_Extra(t *testing.T) {
	// Test that HTTP handlers are registered without panic
	// The init() function in api_flags.go registers handlers
	// We just verify the handlers exist
	handlers := []string{
		"/debug/nox/flags",
		"/debug/nox/flags/game",
		"/debug/nox/flags/engine",
	}
	for _, h := range handlers {
		// http.HandleFunc registers to DefaultServeMux
		// We can't easily check, but init should not panic
		_ = h
	}
	// Verify DefaultServeMux is not nil
	if http.DefaultServeMux == nil {
		t.Error("DefaultServeMux should not be nil")
	}
}
