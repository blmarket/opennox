package handles

import (
	"testing"
)

func TestHandles(t *testing.T) {
	Init()
	defer Release()

	h1 := New()
	if !IsValid(h1) {
		t.Fatalf("h1 not valid")
	}
	h2 := NewPtr()
	if h2 == nil {
		t.Fatalf("h2 nil")
	}
	AssertValid(h1)
	AssertValidPtr(h2)

	if IsValid(0) {
		t.Fatalf("0 should not be valid")
	}
	if IsValid(^uintptr(0)) {
		t.Fatalf("max should not be valid")
	}

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected panic for zero handle")
			}
		}()
		AssertValid(0)
	}()

	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected panic for invalid handle")
			}
		}()
		AssertValid(^uintptr(0))
	}()
}
