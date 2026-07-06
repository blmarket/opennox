package netstr

import (
	"testing"
	"time"
)

func TestNewStreamsAdditional(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	if s == nil {
		t.Fatal("NewStreams returned nil")
	}
	if s.Host() != nil {
		t.Error("Host should be nil initially")
	}
}

func TestErrIsInUseAdditional(t *testing.T) {
	if ErrIsInUse(nil) {
		t.Error("nil should not be in use")
	}
}

func TestConnectFailErrAdditional(t *testing.T) {
	e := &ConnectFailErr{Err: nil, Code: 0}
	s := e.Error()
	if s == "" {
		t.Error("Error string should not be empty")
	}
	if e.Unwrap() != nil {
		t.Error("Unwrap nil should return nil")
	}
}

func TestNewConnectErrAdditional(t *testing.T) {
	e := NewConnectErr(-5, nil)
	if e.Code != -5 {
		t.Errorf("Code = %d, want -5", e.Code)
	}
}

func TestProcessStatsAdditional(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	// Should not panic on nil or empty streams
	s.ProcessStats(time.Second, time.Second)
}

func TestStreamsUpdateAdditional(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	s.Update()
}
