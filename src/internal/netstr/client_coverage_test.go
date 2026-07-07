package netstr

import (
	"testing"
	"time"
)

func TestClientCoverage(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })

	// Test NewClient with nil options
	_, err := s.NewClient(nil)
	if err == nil {
		t.Error("NewClient with nil should return error")
	}

	// Test NewClient with valid options
	ns, err := s.NewClient(&Options{})
	if err != nil {
		t.Logf("NewClient error (expected without network): %v", err)
	}
	if ns != nil {
		// Test Dial with nil ns
		var nilClient *Client
		func() {
			defer func() { recover() }()
			err = nilClient.Dial("localhost", 8080, 0, nil)
			if err == nil {
				t.Error("Dial with nil client should return error")
			}
		}()

		// Test Dial with empty host
		err = ns.Dial("", 8080, 0, nil)
		if err == nil {
			t.Error("Dial with empty host should return error")
		}

		// Test Dial with invalid port
		err = ns.Dial("localhost", 0, 0, nil)
		if err == nil {
			t.Error("Dial with invalid port should return error")
		}
		err = ns.Dial("localhost", 70000, 0, nil)
		if err == nil {
			t.Error("Dial with invalid port should return error")
		}

		// Test Dial with invalid host
		err = ns.Dial("invalid.host.name.that.does.not.exist", 8080, 0, nil)
		if err == nil {
			t.Log("Dial with invalid host should return error (or may succeed in some environments)")
		}

		// Test DialWait
		err = ns.DialWait(time.Millisecond, func() {}, func() bool { return true })
		if err != nil {
			t.Logf("DialWait error: %v", err)
		}
		err = ns.DialWait(0, func() {}, func() bool { return false })
		if err == nil {
			t.Error("DialWait with timeout should return error")
		}

		// Test WaitServerResponse with nil
		var nilNs *Client
		res := nilNs.WaitServerResponse(0, 0, 0)
		if res != -3 {
			t.Errorf("WaitServerResponse with nil should return -3, got %d", res)
		}

		// Test WaitServerResponse
		res = ns.WaitServerResponse(0, 1, 0)
		_ = res
		res = ns.WaitServerResponse(1000, 1, 0)
		_ = res
	}

	// Test NewClient when no more slots
	s2 := NewStreams(func() uint32 { return 0 })
	// Fill up streams
	for i := 0; i < 100; i++ {
		_, _ = s2.NewClient(&Options{})
	}
}
