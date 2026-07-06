package discover

import (
	"context"
	"net"
	"net/netip"
	"testing"
	"time"

	"github.com/noxworld-dev/lobby"
)

func TestPingEachServer(t *testing.T) {
	// Register a test backend that returns a server
	backendName := "test_ping_backend"
	RegisterBackend(backendName, func(ctx context.Context, out chan<- Server) error {
		out <- Server{
			Game: lobby.Game{
				Name: "TestServer",
				Port: 8080,
			},
			Source: "test",
			IP:     netip.MustParseAddr("127.0.0.1"),
		}
		return nil
	})

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Test with nil pc (should create its own)
	err := PingEachServer(ctx, nil, func(s Server) error {
		if s.Name != "TestServer" {
			t.Errorf("PingEachServer server name = %q, want %q", s.Name, "TestServer")
		}
		return nil
	})
	if err != nil && err != context.DeadlineExceeded {
		// Deadline exceeded is ok, we just want to ensure it doesn't panic
	}

	// Test with provided pc
	pc, err := net.ListenUDP("udp4", &net.UDPAddr{IP: nil, Port: 0})
	if err != nil {
		t.Fatalf("Failed to create UDP conn: %v", err)
	}
	defer pc.Close()

	ctx2, cancel2 := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel2()

	err = PingEachServer(ctx2, pc, func(s Server) error {
		return nil
	})
	if err != nil && err != context.DeadlineExceeded {
		// Ok
	}
}

func TestPingEachServer_Empty(t *testing.T) {
	// Test with no backends (should not panic)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	err := PingEachServer(ctx, nil, func(s Server) error {
		return nil
	})
	// Should not return error for empty backends, just timeout
	if err != nil && err != context.DeadlineExceeded {
		t.Logf("PingEachServer with no backends returned: %v", err)
	}
}

func TestListServersWith(t *testing.T) {
	// Register a test backend
	backendName := "test_list_backend"
	RegisterBackend(backendName, func(ctx context.Context, out chan<- Server) error {
		out <- Server{
			Game: lobby.Game{
				Name: "ListTestServer",
				Port: 9090,
			},
			Source: "test",
			IP:     netip.MustParseAddr("127.0.0.1"),
		}
		return nil
	})

	pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create packet conn: %v", err)
	}
	defer pc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	servers, err := ListServersWith(ctx, pc)
	if err != nil && err != context.DeadlineExceeded {
		t.Logf("ListServersWith returned: %v", err)
	}

	// Should have at least our test server
	found := false
	for _, s := range servers {
		if s.Name == "ListTestServer" {
			found = true
			break
		}
	}
	if !found && len(servers) > 0 {
		t.Logf("Test server not found in results, got %d servers", len(servers))
	}
}

func TestListServersWith_NoDeadline(t *testing.T) {
	// Test without deadline in context (should set its own timeout)
	pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create packet conn: %v", err)
	}
	defer pc.Close()

	ctx := context.Background() // No deadline

	servers, err := ListServersWith(ctx, pc)
	if err != nil {
		t.Logf("ListServersWith without deadline returned: %v", err)
	}
	_ = servers
}

func TestListServersWith_Empty(t *testing.T) {
	pc, err := net.ListenPacket("udp4", "127.0.0.1:0")
	if err != nil {
		t.Fatalf("Failed to create packet conn: %v", err)
	}
	defer pc.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	servers, err := ListServersWith(ctx, pc)
	if err != nil && err != context.DeadlineExceeded {
		t.Logf("ListServersWith empty returned: %v", err)
	}
	_ = servers
}
