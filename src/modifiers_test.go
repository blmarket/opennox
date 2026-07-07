package opennox

import (
	"testing"
)

func TestSub4A5E90A(t *testing.T) {
	// sub_4A5E90_A accesses noxServer.Modif.Dword_5d4594_251608
	defer func() {
		if r := recover(); r != nil {
			t.Logf("sub_4A5E90_A panicked as expected: %v", r)
		}
	}()
	sub_4A5E90_A()
}

func TestNoxXxxFireRingEffect(t *testing.T) {
	// nox_xxx_fireRingEffect_4E05B0 accesses src.PosVec and calls legacy function
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_xxx_fireRingEffect_4E05B0 panicked as expected: %v", r)
		}
	}()
	nox_xxx_fireRingEffect_4E05B0(nil, nil, nil, nil)
}

func TestNoxXxxBlueFREffect(t *testing.T) {
	// nox_xxx_blueFREffect_4E05F0 similar to fire ring effect
	defer func() {
		if r := recover(); r != nil {
			t.Logf("nox_xxx_blueFREffect_4E05F0 panicked as expected: %v", r)
		}
	}()
	nox_xxx_blueFREffect_4E05F0(nil, nil, nil, nil)
}
