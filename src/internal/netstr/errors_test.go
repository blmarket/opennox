package netstr

import (
	"errors"
	"syscall"
	"testing"
)

func TestErrIsInUse(t *testing.T) {
	if !ErrIsInUse(syscall.EADDRINUSE) {
		t.Error("ErrIsInUse should return true for EADDRINUSE")
	}
	if ErrIsInUse(syscall.EAGAIN) {
		t.Error("ErrIsInUse should return false for EAGAIN")
	}
	if ErrIsInUse(nil) {
		t.Error("ErrIsInUse should return false for nil")
	}
}

func TestNewConnectErr(t *testing.T) {
	err := errors.New("test error")

	// Test with code 0 (should default to -1)
	e := NewConnectErr(0, err)
	if e.Code != -1 {
		t.Errorf("NewConnectErr with code 0 should default to -1, got %d", e.Code)
	}
	if e.Err != err {
		t.Error("NewConnectErr should preserve error")
	}

	// Test with non-zero code
	e = NewConnectErr(42, err)
	if e.Code != 42 {
		t.Errorf("NewConnectErr code = %d, want 42", e.Code)
	}
}

func TestConnectFailErr_Error(t *testing.T) {
	err := errors.New("connection refused")
	e := &ConnectFailErr{
		Err:  err,
		Code: 5,
	}
	got := e.Error()
	want := "CONNECT_SERVER failed: connection refused (code=5)"
	if got != want {
		t.Errorf("Error() = %q, want %q", got, want)
	}
}

func TestConnectFailErr_Unwrap(t *testing.T) {
	err := errors.New("test")
	e := &ConnectFailErr{
		Err:  err,
		Code: 1,
	}
	if e.Unwrap() != err {
		t.Error("Unwrap() should return the wrapped error")
	}
}
