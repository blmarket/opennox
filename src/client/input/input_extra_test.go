package input

import (
	"testing"
)

func TestMouseButton(t *testing.T) {
	b := NOX_MOUSE_LEFT
	s := b.String()
	if s != "left" && s != "Left" {
		t.Errorf("MouseLeft.String() = %q", s)
	}
}

func TestHandler(t *testing.T) {
	_ = New
}
