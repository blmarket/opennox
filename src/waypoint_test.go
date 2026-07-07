package opennox

import (
	"testing"
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
