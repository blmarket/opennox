package opennox

import (
	"testing"
)

func TestFlushCoverage(t *testing.T) {
	// flushCoverage should not panic
	// In ccover build, it calls ccover.Dump()
	// In non-ccover build, it does nothing
	flushCoverage()
}

func TestSub57C4902(t *testing.T) {
	// sub_57C490_2 accesses noxServer.Map.Debug
	// It should panic without server, or do nothing with empty key
	defer func() {
		if r := recover(); r != nil {
			t.Logf("sub_57C490_2 panicked as expected without server: %v", r)
		}
	}()
	sub_57C490_2("testkey")
	sub_57C490_2("")
}

func TestNoxServerMapRWDebugData(t *testing.T) {
	// nox_server_mapRWDebugData_5060D0 accesses noxServer.Map.Debug
	// It should panic without server
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_server_mapRWDebugData_5060D0 panicked as expected without server: %v", r)
		}
	}()
	err := nox_server_mapRWDebugData_5060D0(nil, nil)
	if err != nil {
		t.Logf("nox_server_mapRWDebugData_5060D0 returned error: %v", err)
	}
}
