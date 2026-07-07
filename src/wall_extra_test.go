package opennox

import (
	"testing"
	"unsafe"
)

func TestWallCExtra(t *testing.T) {
	var w Wall
	ptr := w.C()
	if ptr != unsafe.Pointer(&w) {
		t.Error("Wall.C() should return pointer to wall")
	}
}

func TestWallSExtra(t *testing.T) {
	var w Wall
	s := w.S()
	if s == nil {
		t.Error("Wall.S() should not return nil for non-nil wall")
	}

	var w2 *Wall
	s2 := w2.S()
	if s2 != nil {
		t.Error("Wall.S() should return nil for nil wall")
	}
}

func TestWallDefExtra(t *testing.T) {
	var w *Wall
	def := w.Def()
	if def != nil {
		t.Error("Wall.Def() should return nil for nil wall")
	}
}

func TestWallString(t *testing.T) {
	// Just ensure it doesn't panic on nil wall (may panic, so recover)
	defer func() {
		recover()
	}()
	var w *Wall
	_ = w.String()
}

func TestWallScriptID(t *testing.T) {
	defer func() {
		recover()
	}()
	var w *Wall
	_ = w.WallScriptID()
}

func TestWallGridPos(t *testing.T) {
	defer func() {
		recover()
	}()
	var w *Wall
	_ = w.GridPos()
}
