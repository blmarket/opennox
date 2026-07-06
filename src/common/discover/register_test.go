package discover

import (
	"context"
	"testing"
	"time"

	"github.com/noxworld-dev/lobby"
	"github.com/noxworld-dev/xwis"
)

func TestNewServer(t *testing.T) {
	info := GameInfo{}
	s := NewServer(info)
	if s == nil {
		t.Fatal("NewServer should not return nil")
	}
	if s.stop == nil {
		t.Error("RegServer stop should not be nil")
	}
	if s.ticks == nil {
		t.Error("RegServer ticks should not be nil")
	}
	if s.xwis == nil {
		t.Error("RegServer xwis should not be nil")
	}
	// Close should work
	if err := s.Close(); err != nil {
		t.Errorf("Close should not error: %v", err)
	}
	// Second close should be safe
	if err := s.Close(); err != nil {
		t.Errorf("Second Close should not error: %v", err)
	}
}

func TestRegServerUpdate(t *testing.T) {
	s := NewServer(GameInfo{})
	defer s.Close()

	info := GameInfo{
		Game: lobby.Game{
			Name: "Test Server",
		},
		XWIS: xwis.GameInfo{
			Name: "Test",
		},
	}
	s.Update(info)

	// Verify info was updated
	got := s.info()
	if got.Game.Name != "Test Server" {
		t.Errorf("Update should set game name, got %q", got.Game.Name)
	}
}

func TestRegServerInfo(t *testing.T) {
	info := GameInfo{
		Game: lobby.Game{
			Name: "Info Test",
		},
	}
	s := NewServer(info)
	defer s.Close()

	got := s.info()
	if got.Game.Name != "Info Test" {
		t.Errorf("info() should return current info, got %q", got.Game.Name)
	}
}

func TestGameRegHostGameInfo(t *testing.T) {
	info := GameInfo{
		Game: lobby.Game{
			Name: "Host Test",
		},
	}
	s := NewServer(info)
	defer s.Close()

	h := gameRegHost{s: s}
	game, err := h.GameInfo(context.Background())
	if err != nil {
		t.Fatalf("GameInfo should not error: %v", err)
	}
	if game.Name != "Host Test" {
		t.Errorf("GameInfo should return game, got name %q", game.Name)
	}
}

func TestInfoSetDefaults(t *testing.T) {
	info := &xwis.GameInfo{}
	infoSetDefaults(info)

	if info.Name != "OpenNox Server" {
		t.Errorf("Name default should be 'OpenNox Server', got %q", info.Name)
	}
	if info.Map != "estate" {
		t.Errorf("Map default should be 'estate', got %q", info.Map)
	}
	if info.MapType != xwis.MapTypeChat {
		t.Errorf("MapType default should be MapTypeChat")
	}
	if info.Resolution != xwis.Res1024x768 {
		t.Errorf("Resolution default should be Res1024x768")
	}
	if info.MaxPlayers != 32 {
		t.Errorf("MaxPlayers default should be 32, got %d", info.MaxPlayers)
	}
}

func TestInfoSetDefaultsNoOverride(t *testing.T) {
	info := &xwis.GameInfo{
		Name:       "Custom",
		Map:        "custommap",
		MapType:    xwis.MapTypeArena,
		Resolution: xwis.Res800x600,
		MaxPlayers: 16,
		Flags:      1,
	}
	infoSetDefaults(info)

	if info.Name != "Custom" {
		t.Error("Name should not be overridden")
	}
	if info.Map != "custommap" {
		t.Error("Map should not be overridden")
	}
	if info.MapType != xwis.MapTypeArena {
		t.Error("MapType should not be overridden when Flags set")
	}
	if info.Resolution != xwis.Res800x600 {
		t.Error("Resolution should not be overridden")
	}
	if info.MaxPlayers != 16 {
		t.Error("MaxPlayers should not be overridden")
	}
}

func TestRegServerCloseNilStop(t *testing.T) {
	s := &RegServer{}
	if err := s.Close(); err != nil {
		t.Errorf("Close with nil stop should not error: %v", err)
	}
}

func TestRegServerUpdateNonBlocking(t *testing.T) {
	s := NewServer(GameInfo{})
	defer s.Close()

	// Fill the channels to test non-blocking behavior
	for i := 0; i < 10; i++ {
		select {
		case s.ticks <- time.Now():
		default:
		}
		select {
		case s.xwis <- struct{}{}:
		default:
		}
	}

	// Update should not block even if channels are full
	done := make(chan struct{})
	go func() {
		s.Update(GameInfo{})
		close(done)
	}()

	select {
	case <-done:
		// OK
	case <-time.After(time.Second):
		t.Error("Update should not block")
	}
}
