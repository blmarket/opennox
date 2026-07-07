package opennox

import (
	"testing"

	ns "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/wall"
)

type mockWall struct {
	pos ns.Pointf
}

func (m mockWall) Pos() ns.Pointf                           { return m.pos }
func (m mockWall) ScriptID() int                            { return 0 }
func (m mockWall) WallScriptID() int                        { return 0 }
func (m mockWall) Name() string                             { return "" }
func (m mockWall) Enable(bool)                              {}
func (m mockWall) Toggle() bool                             { return false }
func (m mockWall) Destroy()                                 {}
func (m mockWall) EachWall(bool, func(obj ns.WallObj) bool) {}
func (m mockWall) IsEnabled() bool                          { return true }
func (m mockWall) Flags() wall.Flags                        { return 0 }
func (m mockWall) GridPos() ns.Point                        { return ns.Point{} }

type mockWallCond struct {
	match bool
}

func (m mockWallCond) WallMatches(obj ns.WallObj) bool {
	return m.match
}

func TestWallAND_Matches(t *testing.T) {
	w := mockWall{pos: ns.Pointf{X: 1, Y: 2}}

	// empty should match
	var empty wallAND
	if !empty.Matches(w) {
		t.Fatal("empty wallAND should match")
	}

	// all true
	arr := wallAND{
		mockWallCond{true},
		mockWallCond{true},
	}
	if !arr.Matches(w) {
		t.Fatal("all true should match")
	}

	// one false
	arr2 := wallAND{
		mockWallCond{true},
		mockWallCond{false},
	}
	if arr2.Matches(w) {
		t.Fatal("one false should not match")
	}

	// all false
	arr3 := wallAND{
		mockWallCond{false},
		mockWallCond{false},
	}
	if arr3.Matches(w) {
		t.Fatal("all false should not match")
	}
}

func TestWallAND_Matches_Nil(t *testing.T) {
	var arr wallAND
	if !arr.Matches(nil) {
		t.Fatal("nil wall with empty conditions should match")
	}
}
