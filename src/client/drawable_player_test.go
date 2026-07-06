package client

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/client/noxrender"
)

func TestPlayerAnimation_FramesSlice(t *testing.T) {
	var handles [5]noxrender.ImageHandle
	for i := range handles {
		handles[i] = noxrender.ImageHandle(uintptr(i + 1))
	}
	pa := &PlayerAnimation{}
	pa.Base.Cnt40 = 5
	slice := pa.FramesSlice(&handles[0])
	if len(slice) != 5 {
		t.Fatalf("expected len 5, got %d", len(slice))
	}
	for i, h := range slice {
		if h != handles[i] {
			t.Fatalf("mismatch at %d", i)
		}
	}
}

func TestPlayerAnimation_FramesSlice_Empty(t *testing.T) {
	pa := &PlayerAnimation{}
	pa.Base.Cnt40 = 0
	var h noxrender.ImageHandle
	slice := pa.FramesSlice(&h)
	if len(slice) != 0 {
		t.Fatalf("expected empty slice, got len %d", len(slice))
	}
}
