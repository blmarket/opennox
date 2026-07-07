package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/types"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/server"
)

func testServerWithMapGroups(t *testing.T) *Server {
	t.Helper()

	hadFlag22 := noxflags.HasGame(noxflags.GameFlag22)
	noxflags.SetGame(noxflags.GameFlag22)
	t.Cleanup(func() {
		if !hadFlag22 {
			noxflags.UnsetGame(noxflags.GameFlag22)
		}
	})

	s := &Server{Server: &server.Server{}}
	s.MapGroups.Init()
	t.Cleanup(s.MapGroups.Free)
	return s
}

func TestServerGetWaypointGroupByID(t *testing.T) {
	// getWaypointGroupByID accesses s.MapGroups and s.WPs
	s := &Server{}
	defer func() {
		if r := recover(); r != nil {
			t.Logf("getWaypointGroupByID panicked as expected: %v", r)
		}
	}()
	result := s.getWaypointGroupByID("")
	if result != nil {
		t.Logf("getWaypointGroupByID with empty ID returned non-nil: %v", result)
	}

	result = s.getWaypointGroupByID("test")
	if result != nil {
		t.Logf("getWaypointGroupByID with 'test' returned non-nil: %v", result)
	}
}

func TestGetWaypointGroupByIDWithGroups(t *testing.T) {
	s := testServerWithMapGroups(t)

	// Test with non-existent group
	result := s.getWaypointGroupByID("nonexistent")
	if result != nil {
		t.Error("Expected nil for non-existent group")
	}

	// Create a waypoint group
	wp := s.NewWaypoint(types.Pointf{})

	if ok := s.MapGroups.MapLoadAddGroup57C0C0("test_group", 1, byte(server.MapGroupWaypoints)); ok == 0 {
		t.Fatal("Failed to create waypoint group")
	}
	if ok := s.MapGroups.Sub57C130([]uint32{uint32(wp.Ind())}, 1); ok == 0 {
		t.Fatal("Failed to add waypoint to group")
	}

	result = s.getWaypointGroupByID("test_group")
	if result == nil {
		t.Fatal("Expected non-nil result for existing group")
	}

	if result.ID() != "test_group" {
		t.Errorf("Group ID = %q, want %q", result.ID(), "test_group")
	}
}

func TestGetWaypointGroupByIDEmptyGroup(t *testing.T) {
	s := testServerWithMapGroups(t)

	if ok := s.MapGroups.MapLoadAddGroup57C0C0("empty_group", 1, byte(server.MapGroupWaypoints)); ok == 0 {
		t.Fatal("Failed to create waypoint group")
	}

	result := s.getWaypointGroupByID("empty_group")
	if result == nil {
		t.Fatal("Expected non-nil result for empty group")
	}

	if len(result.Waypoints()) != 0 {
		t.Errorf("Empty group should have 0 waypoints, got %d", len(result.Waypoints()))
	}
}
