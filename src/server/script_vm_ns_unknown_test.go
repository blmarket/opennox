package server

import "testing"

func TestNoxScriptNS_Unused(t *testing.T) {
	var s NoxScriptNS
	s.Unused1f(1)
	s.Unused20(2)
	s.Unused50()
	s.Unused58(1, 2)
	s.Unused59(3, 4)
	s.Unused5a(5, 6)
	s.Unused5b(7, 8)
	s.Unused5c(9, 10)
	s.Unused5d(11, 12)
}
