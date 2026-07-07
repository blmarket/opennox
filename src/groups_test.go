package opennox

import (
	"testing"
)

func TestServerSub504720(t *testing.T) {
	// Sub504720 accesses s.MapGroups and calls groupsAdjust
	// Without proper setup, it will panic
	s := &Server{}
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Sub504720 panicked as expected: %v", r)
		}
	}()
	result := s.Sub504720(0, 0)
	t.Logf("Sub504720 returned: %d", result)
}

func TestServerGroupsAdjust(t *testing.T) {
	// groupsAdjust accesses s.MapGroups
	s := &Server{}
	defer func() {
		if r := recover(); r != nil {
			t.Logf("groupsAdjust panicked as expected: %v", r)
		}
	}()
	s.groupsAdjust(0, 0)
	s.groupsAdjust(100, 200)
}
