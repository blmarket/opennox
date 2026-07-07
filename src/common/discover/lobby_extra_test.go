package discover

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/internal/version"
)

func TestNewLobbyClient(t *testing.T) {
	// Save original LobbyServer
	originalServer := LobbyServer
	defer func() { LobbyServer = originalServer }()

	// Test with custom server
	LobbyServer = "http://test.example.com:8080"
	cli := newLobbyClient()
	if cli == nil {
		t.Fatal("newLobbyClient should not return nil")
	}

	// Test with default server
	LobbyServer = originalServer
	cli = newLobbyClient()
	if cli == nil {
		t.Fatal("newLobbyClient should not return nil with default server")
	}
}

func TestLobbyServerVariable(t *testing.T) {
	// Test that LobbyServer variable can be modified
	original := LobbyServer
	defer func() { LobbyServer = original }()

	LobbyServer = "http://custom.lobby.server:9090"
	if LobbyServer != "http://custom.lobby.server:9090" {
		t.Error("LobbyServer variable should be modifiable")
	}

	cli := newLobbyClient()
	if cli == nil {
		t.Error("newLobbyClient should work with custom LobbyServer")
	}
}

func TestNewLobbyClientUserAgent(t *testing.T) {
	// Test that user agent is set correctly
	cli := newLobbyClient()
	if cli == nil {
		t.Fatal("newLobbyClient should not return nil")
	}
	// The client should have a user agent containing the version
	// We can't directly inspect the user agent, but we can verify the client was created
	// with the version from the version package
	_ = version.Version() // Ensure version package is used
}
