package client

import (
	"testing"
	"unsafe"
)

func TestDrawableFX_C(t *testing.T) {
	fx := &DrawableFX{
		Field0: 1,
		Field4: 2,
	}
	ptr := fx.C()
	if ptr == nil {
		t.Fatal("expected non-nil pointer")
	}
	if ptr != unsafe.Pointer(fx) {
		t.Fatalf("expected pointer to fx, got %v", ptr)
	}
}
