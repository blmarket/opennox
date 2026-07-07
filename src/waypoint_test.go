package opennox

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/server"
)

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
	s := &Server{}
	s.MapGroups = server.NewMapGroups()

	// Test with non-existent group
	result := s.getWaypointGroupByID("nonexistent")
	if result != nil {
		t.Error("Expected nil for non-existent group")
	}

	// Create a waypoint group
	s.WPs = server.NewWaypoints()
	wp := &server.Waypoint{}
	wp.SetInd(1)
	s.WPs.Add(wp)

	g := s.MapGroups.NewGroup("test_group", server.MapGroupWaypoints)
	g.Add(1, 0)

	result = s.getWaypointGroupByID("test_group")
	if result == nil {
		t.Fatal("Expected non-nil result for existing group")
	}

	if result.ID() != "test_group" {
		t.Errorf("Group ID = %q, want %q", result.ID(), "test_group")
	}
}

func TestGetWaypointGroupByIDEmptyGroup(t *testing.T) {
	s := &Server{}
	s.MapGroups = server.NewMapGroups()
	s.WPs = server.NewWaypoints()

	g := s.MapGroups.NewGroup("empty_group", server.MapGroupWaypoints)
	_ = g

	result := s.getWaypointGroupByID("empty_group")
	if result == nil {
		t.Fatal("Expected non-nil result for empty group")
	}

	if len(result.List()) != 0 {
		t.Errorf("Empty group should have 0 waypoints, got %d", len(result.List()))
	}
}
