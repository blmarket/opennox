package netstr

import (
	"errors"
	"fmt"
	"syscall"
	"testing"
)

func TestErrIsInUse_Extra2(t *testing.T) {
	// Test with wrapped error
	wrapped := fmt.Errorf("wrapped: %w", syscall.EADDRINUSE)
	if !ErrIsInUse(wrapped) {
		t.Error("ErrIsInUse should return true for wrapped EADDRINUSE")
	}
	// Test with other syscall errors
	if ErrIsInUse(syscall.EAGAIN) {
		t.Error("ErrIsInUse should return false for EAGAIN")
	}
	if ErrIsInUse(syscall.ECONNREFUSED) {
		t.Error("ErrIsInUse should return false for ECONNREFUSED")
	}
}

func TestNewConnectErr_Extra2(t *testing.T) {
	// Test with nil error
	e := NewConnectErr(5, nil)
	if e.Code != 5 {
		t.Errorf("Code = %d, want 5", e.Code)
	}
	if e.Err != nil {
		t.Error("Err should be nil")
	}
	// Test Error() with nil
	got := e.Error()
	if got == "" {
		t.Error("Error() should not be empty even with nil error")
	}

	// Test with negative code
	e = NewConnectErr(-10, errors.New("test"))
	if e.Code != -10 {
		t.Errorf("Code = %d, want -10", e.Code)
	}
}

func TestConnectFailErr_Error_Extra2(t *testing.T) {
	// Test with different error messages
	tests := []struct {
		err  error
		code int
		want string
	}{
		{errors.New("timeout"), 1, "CONNECT_SERVER failed: timeout (code=1)"},
		{errors.New(""), 0, "CONNECT_SERVER failed:  (code=0)"},
	}
	for _, tt := range tests {
		e := &ConnectFailErr{Err: tt.err, Code: tt.code}
		got := e.Error()
		if got != tt.want {
			t.Errorf("Error() = %q, want %q", got, tt.want)
		}
	}
}

func TestConnectFailErr_Unwrap_Extra2(t *testing.T) {
	// Test unwrap with nil
	e := &ConnectFailErr{Err: nil, Code: 1}
	if e.Unwrap() != nil {
		t.Error("Unwrap() should return nil when Err is nil")
	}
	// Test unwrap preserves error chain
	orig := errors.New("original")
	e = &ConnectFailErr{Err: orig, Code: 2}
	if !errors.Is(e, orig) {
		t.Error("errors.Is should work with Unwrap")
	}
}
