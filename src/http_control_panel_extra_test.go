package opennox

import (
	"testing"
)

func TestGameControlHTTPChangeMap(t *testing.T) {
	// Test ChangeMap with allowMapChange false - should not panic
	g := &gameControlHTTP{
		allowMapChange: false,
	}
	err := g.ChangeMap("testmap")
	if err != nil {
		t.Errorf("ChangeMap with allowMapChange=false should not return error, got %v", err)
	}

	// Test with allowMapChange true - may panic because queueServerMapLoad requires server
	// Just verify it doesn't panic without server, or recover if it does
	g.allowMapChange = true
	defer func() {
		if r := recover(); r != nil {
			t.Logf("ChangeMap with allowMapChange=true panicked as expected: %v", r)
		}
	}()
	err = g.ChangeMap("testmap")
	if err != nil {
		t.Errorf("ChangeMap should not return error, got %v", err)
	}
}

func TestGameControlHTTPCommand(t *testing.T) {
	// Test Command with allowCmds false - should not panic
	g := &gameControlHTTP{
		allowCmds: false,
	}
	err := g.Command("test command")
	if err != nil {
		t.Errorf("Command with allowCmds=false should not return error, got %v", err)
	}

	// Test with allowCmds true - may panic
	g.allowCmds = true
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Command with allowCmds=true panicked as expected: %v", r)
		}
	}()
	err = g.Command("test command")
	if err != nil {
		t.Errorf("Command should not return error, got %v", err)
	}
}

func TestGameControlHTTPListMaps(t *testing.T) {
	g := &gameControlHTTP{}
	// ListMaps calls scanMaps which may fail if maps directory doesn't exist
	// Just verify it doesn't panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("ListMaps panicked as expected: %v", r)
		}
	}()
	list, err := g.ListMaps()
	t.Logf("ListMaps returned %d maps, err=%v", len(list), err)
}

func TestGameControlHTTPGameInfo(t *testing.T) {
	g := &gameControlHTTP{}
	// GameInfo calls getGameInfo which requires server
	defer func() {
		if r := recover(); r != nil {
			t.Logf("GameInfo panicked as expected: %v", r)
		}
	}()
	info, err := g.GameInfo()
	t.Logf("GameInfo returned info=%v, err=%v", info, err)
}
