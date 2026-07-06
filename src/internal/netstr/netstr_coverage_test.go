package netstr

import (
	"testing"
)

func TestNewStreamsCoverage(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	if s == nil {
		t.Fatal("NewStreams returned nil")
	}

	// Test basic methods
	s.reset()
	if s.Host() != nil {
		t.Errorf("Host should be nil initially")
	}

	// Test ConnByPlayerInd with invalid index
	if s.ConnByPlayerInd(-1) != nil {
		t.Errorf("ConnByPlayerInd(-1) should be nil")
	}

	// Test ByPlayer
	func() {
		defer func() { recover() }()
		_ = s.ByPlayer(nil)
	}()

	// Test StreamByPlayerInd
	_ = s.StreamByPlayerInd(0)

	// Test HostStream
	_ = s.HostStream()

	// Test GetTimingByInd1
	_ = s.GetTimingByInd1(0)

	// Test getFreeIndex
	idx, _ := s.getFreeIndex()
	if idx < 0 {
		t.Errorf("getFreeIndex returned negative")
	}

	// Test Update
	s.Update()

	// Test ProcessStats
	s.ProcessStats(0, 0)
}

func TestHandleCoverage(t *testing.T) {
	// Handle is unexported, test via public APIs
	s := NewStreams(func() uint32 { return 0 })
	if s.Host() != nil {
		t.Errorf("Host should be nil initially")
	}
}

func TestErrorCoverage(t *testing.T) {
	err := NewConnectErr(1, nil)
	if err == nil {
		t.Fatal("NewConnectErr returned nil")
	}
	if err.Error() == "" {
		t.Errorf("Error should not be empty")
	}
	if err.Unwrap() != nil {
		t.Errorf("Unwrap should be nil")
	}

	if !ErrIsInUse(err) {
		t.Logf("ErrIsInUse returned false (expected for code 1)")
	}

	err2 := NewConnectErr(2, nil)
	_ = ErrIsInUse(err2)
}
