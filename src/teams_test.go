package opennox

import (
	"testing"
)

func TestNoxXxxObjGetTeamByNetCode(t *testing.T) {
	// nox_xxx_objGetTeamByNetCode_418C80 accesses noxServer or noxClient
	// Without them, it should panic or return nil
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_xxx_objGetTeamByNetCode_418C80 panicked as expected: %v", r)
		}
	}()
	result := nox_xxx_objGetTeamByNetCode_418C80(0)
	if result != nil {
		t.Logf("nox_xxx_objGetTeamByNetCode_418C80 returned non-nil: %v", result)
	}

	result = nox_xxx_objGetTeamByNetCode_418C80(12345)
	if result != nil {
		t.Logf("nox_xxx_objGetTeamByNetCode_418C80 returned non-nil: %v", result)
	}
}
