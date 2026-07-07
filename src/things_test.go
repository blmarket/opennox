package opennox

import (
	"testing"
)

func TestObjectTypeCode16(t *testing.T) {
	// Test with nil array
	objectTypeCode16ByInd = nil

	if v := sub_42C2E0(0); v != 0 {
		t.Errorf("sub_42C2E0 with nil array should return 0, got %d", v)
	}

	sub_42C310(0, 123) // Should do nothing with nil array

	if v := sub_42C300(); v != 0 {
		t.Errorf("sub_42C300 with nil array should return 0, got %d", v)
	}

	if v := nox_xxx_objectTOCgetTT(123); v != 0 {
		t.Errorf("nox_xxx_objectTOCgetTT with nil array should return 0, got %d", v)
	}

	// Initialize array
	objectTypeCode16ByInd = make([]uint16, 10)
	objectTypeCode16ByInd_len = 0

	// Test sub_42C310 and sub_42C2E0
	sub_42C310(0, 100)
	sub_42C310(1, 200)
	sub_42C310(5, 300)

	if v := sub_42C2E0(0); v != 100 {
		t.Errorf("sub_42C2E0(0) = %d, want 100", v)
	}
	if v := sub_42C2E0(1); v != 200 {
		t.Errorf("sub_42C2E0(1) = %d, want 200", v)
	}
	if v := sub_42C2E0(5); v != 300 {
		t.Errorf("sub_42C2E0(5) = %d, want 300", v)
	}
	if v := sub_42C2E0(2); v != 0 {
		t.Errorf("sub_42C2E0(2) = %d, want 0", v)
	}

	// Test nox_xxx_objectTOCgetTT
	if v := nox_xxx_objectTOCgetTT(100); v != 0 {
		t.Errorf("nox_xxx_objectTOCgetTT(100) = %d, want 0", v)
	}
	if v := nox_xxx_objectTOCgetTT(200); v != 1 {
		t.Errorf("nox_xxx_objectTOCgetTT(200) = %d, want 1", v)
	}
	if v := nox_xxx_objectTOCgetTT(300); v != 5 {
		t.Errorf("nox_xxx_objectTOCgetTT(300) = %d, want 5", v)
	}
	if v := nox_xxx_objectTOCgetTT(999); v != 0 {
		t.Errorf("nox_xxx_objectTOCgetTT(999) = %d, want 0", v)
	}

	// Test sub_42C300
	objectTypeCode16ByInd_len = 5
	if v := sub_42C300(); v != 5 {
		t.Errorf("sub_42C300() = %d, want 5", v)
	}

	// Test sub_42BFB0
	sub_42BFB0()
	if len(objectTypeCode16ByInd) != 10 {
		t.Errorf("sub_42BFB0 should keep array length 10, got %d", len(objectTypeCode16ByInd))
	}
	if objectTypeCode16ByInd_len != 0 {
		t.Errorf("sub_42BFB0 should reset len to 0, got %d", objectTypeCode16ByInd_len)
	}
	for i, v := range objectTypeCode16ByInd {
		if v != 0 {
			t.Errorf("sub_42BFB0 should zero array, index %d = %d", i, v)
		}
	}

	// Test nox_xxx_free_42BF80
	nox_xxx_free_42BF80()
	if objectTypeCode16ByInd != nil {
		t.Error("nox_xxx_free_42BF80 should set array to nil")
	}
}

func TestGetThingName(t *testing.T) {
	// getThingName requires noxClient to be initialized
	// With nil client, it should panic or return empty
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("getThingName panicked as expected with nil client: %v", r)
			}
		}()
		name := getThingName(0)
		if name != "" {
			t.Errorf("getThingName with nil client should return empty, got %q", name)
		}
	}()
}

func TestSub42BF10(t *testing.T) {
	// sub_42BF10 initializes objectTypeCode16ByInd based on game flags
	// With nil client/server, it may panic
	objectTypeCode16ByInd = nil
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("sub_42BF10 panicked as expected: %v", r)
			}
		}()
		sub_42BF10()
	}()

	// If array was already initialized, it should return early
	objectTypeCode16ByInd = make([]uint16, 5)
	sub_42BF10()
	if len(objectTypeCode16ByInd) != 5 {
		t.Error("sub_42BF10 should not reinitialize existing array")
	}

	// Cleanup
	objectTypeCode16ByInd = nil
}
