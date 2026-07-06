package client

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/client/noxrender"
)

func TestAnimationVector_FramesSlice(t *testing.T) {
	var handles [3]noxrender.ImageHandle
	for i := range handles {
		handles[i] = noxrender.ImageHandle(uintptr(i + 10))
	}
	av := &AnimationVector{}
	av.Cnt40 = 3
	av.Frames[0] = &handles[0]
	slice := av.FramesSlice(0)
	if len(slice) != 3 {
		t.Fatalf("expected len 3, got %d", len(slice))
	}
	for i, h := range slice {
		if h != handles[i] {
			t.Fatalf("mismatch at %d", i)
		}
	}
}

func TestAnimationVector_FramesSlice_Empty(t *testing.T) {
	av := &AnimationVector{}
	av.Cnt40 = 0
	var h noxrender.ImageHandle
	av.Frames[1] = &h
	slice := av.FramesSlice(1)
	if len(slice) != 0 {
		t.Fatalf("expected empty, got %d", len(slice))
	}
}
