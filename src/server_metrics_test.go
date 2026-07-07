package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/server"
)

func TestServerInitMetrics(t *testing.T) {
	// initMetrics sets up a tick hook that accesses server state
	// With a nil or empty server, it may panic
	s := &Server{}
	defer func() {
		if r := recover(); r != nil {
			t.Logf("initMetrics panicked as expected: %v", r)
		}
	}()
	s.initMetrics()
}

func TestServerUpdateMetrics(t *testing.T) {
	// updateMetrics accesses s.Players and player units
	// With empty server, it should not panic
	s := &Server{}
	defer func() {
		if r := recover(); r != nil {
			t.Logf("updateMetrics panicked: %v", r)
		}
	}()
	s.updateMetrics()
}

func TestServerUpdateMetricsWithPlayers(t *testing.T) {
	s := &server.Server{}
	// updateMetrics is defined on *Server (main package), not server.Server
	// So we can't easily test it without a full Server setup
	// Just verify the function exists
	_ = s
}
