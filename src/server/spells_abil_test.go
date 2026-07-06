package server

import (
	"testing"
)

func TestAbility_String(t *testing.T) {
	tests := []struct {
		a    Ability
		want string
	}{
		{AbilityInvalid, "ABILITY_INVALID"},
		{AbilityBerserk, "ABILITY_BERSERKER_CHARGE"},
		{AbilityWarcry, "ABILITY_WARCRY"},
		{AbilityHarpoon, "ABILITY_HARPOON"},
		{AbilityTreadLightly, "ABILITY_TREAD_LIGHTLY"},
		{AbilityInfravis, "ABILITY_EYE_OF_THE_WOLF"},
		{Ability(999), "Ability(999)"},
		{Ability(-1), "Ability(-1)"},
	}
	for _, tt := range tests {
		got := tt.a.String()
		if got != tt.want {
			t.Errorf("Ability(%d).String() = %q, want %q", tt.a, got, tt.want)
		}
	}
}

func TestAbility_Valid(t *testing.T) {
	tests := []struct {
		a    Ability
		want bool
	}{
		{AbilityInvalid, false},
		{AbilityBerserk, true},
		{AbilityWarcry, true},
		{AbilityHarpoon, true},
		{AbilityTreadLightly, true},
		{AbilityInfravis, true},
		{AbilityMax, false},
		{Ability(999), false},
		{Ability(-1), false},
	}
	for _, tt := range tests {
		got := tt.a.Valid()
		if got != tt.want {
			t.Errorf("Ability(%d).Valid() = %v, want %v", tt.a, got, tt.want)
		}
	}
}

func TestServerAbilities(t *testing.T) {
	var a serverAbilities
	a.Reset()

	if a.ByUnit == nil {
		t.Error("ByUnit should not be nil after Reset")
	}

	// Test GetFor with nil object
	if got := a.GetFor(nil); got != nil {
		t.Errorf("GetFor(nil) = %v, want nil", got)
	}

	// Test IsActive with non-existent unit
	obj := &Object{}
	if a.IsActive(obj, AbilityBerserk) {
		t.Error("IsActive should return false for non-existent unit")
	}

	// Test IsActiveVal with non-existent unit
	if a.IsActiveVal(obj, AbilityBerserk) {
		t.Error("IsActiveVal should return false for non-existent unit")
	}

	// Test IsAnyActive with non-existent unit
	if a.IsAnyActive(obj) {
		t.Error("IsAnyActive should return false for non-existent unit")
	}

	// Test Sub4FC070 with non-existent unit (should not panic)
	a.Sub4FC070(obj, AbilityBerserk, 10)

	// Test Sub4FC030 with non-existent unit
	if got := a.Sub4FC030(obj, AbilityBerserk); got != -1 {
		t.Errorf("Sub4FC030 for non-existent unit = %d, want -1", got)
	}

	// Test Sub4FC440 with non-existent unit (should not panic)
	a.Sub4FC440(obj, AbilityBerserk)

	// Test DisableAbilityAaa with nil object (should not panic)
	a.DisableAbilityAaa(nil, AbilityBerserk)

	// Test DisableAbilityAaa with invalid ability (should not panic)
	a.DisableAbilityAaa(obj, AbilityInvalid)
	a.DisableAbilityAaa(obj, AbilityMax)
	a.DisableAbilityAaa(obj, Ability(999))

	// Test DisableAbilityAaa with non-existent unit (should not panic)
	a.DisableAbilityAaa(obj, AbilityBerserk)
}

func TestServerAbilities_GetFor(t *testing.T) {
	var a serverAbilities
	a.Reset()

	obj := &Object{}

	// First call should create new unitAbilities
	d1 := a.GetFor(obj)
	if d1 == nil {
		t.Fatal("GetFor should not return nil for valid object")
	}

	// Second call should return same unitAbilities
	d2 := a.GetFor(obj)
	if d1 != d2 {
		t.Error("GetFor should return same unitAbilities for same object")
	}
}

func TestServerAbilities_IsActive(t *testing.T) {
	var a serverAbilities
	a.Reset()

	obj := &Object{}
	d := a.GetFor(obj)

	// Initially no abilities active
	if a.IsActive(obj, AbilityBerserk) {
		t.Error("IsActive should return false when no abilities")
	}

	// Add an ability to exec list
	exec := &ExecAbilityClass{
		Abil:   AbilityBerserk,
		Active: 1,
	}
	d.ExecList = exec

	if !a.IsActive(obj, AbilityBerserk) {
		t.Error("IsActive should return true when ability in exec list")
	}

	if a.IsActive(obj, AbilityWarcry) {
		t.Error("IsActive should return false for ability not in exec list")
	}
}

func TestServerAbilities_IsActiveVal(t *testing.T) {
	var a serverAbilities
	a.Reset()

	obj := &Object{}
	d := a.GetFor(obj)

	exec := &ExecAbilityClass{
		Abil:   AbilityBerserk,
		Active: 1,
	}
	d.ExecList = exec

	if !a.IsActiveVal(obj, AbilityBerserk) {
		t.Error("IsActiveVal should return true when active != 0")
	}

	exec.Active = 0
	if a.IsActiveVal(obj, AbilityBerserk) {
		t.Error("IsActiveVal should return false when active == 0")
	}
}

func TestServerAbilities_IsAnyActive(t *testing.T) {
	var a serverAbilities
	a.Reset()

	obj := &Object{}

	if a.IsAnyActive(obj) {
		t.Error("IsAnyActive should return false for non-existent unit")
	}

	d := a.GetFor(obj)
	if a.IsAnyActive(obj) {
		t.Error("IsAnyActive should return false when exec list is nil")
	}

	d.ExecList = &ExecAbilityClass{Abil: AbilityBerserk}
	if !a.IsAnyActive(obj) {
		t.Error("IsAnyActive should return true when exec list not nil")
	}
}

func TestServerAbilities_Sub4FC070(t *testing.T) {
	var a serverAbilities
	s := &Server{}
	a.init(s)
	a.Reset()

	obj := &Object{}
	d := a.GetFor(obj)

	exec := &ExecAbilityClass{
		Abil:  AbilityBerserk,
		Frame: 100,
	}
	d.ExecList = exec

	// This sets Frame to s.Frame() + dt
	// s.Frame() returns 0 for empty server, so Frame should be 10
	a.Sub4FC070(obj, AbilityBerserk, 10)
	if exec.Frame != 10 {
		t.Errorf("Sub4FC070 should set Frame to 10, got %d", exec.Frame)
	}
}

func TestServerAbilities_Sub4FC030(t *testing.T) {
	var a serverAbilities
	s := &Server{}
	a.init(s)
	a.Reset()

	obj := &Object{}
	d := a.GetFor(obj)

	exec := &ExecAbilityClass{
		Abil:  AbilityBerserk,
		Frame: 100,
	}
	d.ExecList = exec

	// s.Frame() is 0, so result should be 100 - 0 = 100
	got := a.Sub4FC030(obj, AbilityBerserk)
	if got != 100 {
		t.Errorf("Sub4FC030 = %d, want 100", got)
	}

	// Non-existent ability should return -1
	if got := a.Sub4FC030(obj, AbilityWarcry); got != -1 {
		t.Errorf("Sub4FC030 for non-existent ability = %d, want -1", got)
	}
}

func TestServerAbilities_Sub4FC440(t *testing.T) {
	var a serverAbilities
	a.Reset()

	obj := &Object{}
	d := a.GetFor(obj)

	exec := &ExecAbilityClass{
		Abil:   AbilityBerserk,
		Active: 1,
	}
	d.ExecList = exec

	a.Sub4FC440(obj, AbilityBerserk)
	if exec.Active != 0 {
		t.Error("Sub4FC440 should set Active to 0")
	}
}

func TestServerAbilities_DisableAbilityAaa(t *testing.T) {
	var a serverAbilities
	a.Reset()

	obj := &Object{}
	d := a.GetFor(obj)

	// Create a linked list of abilities
	exec1 := &ExecAbilityClass{Abil: AbilityBerserk}
	exec2 := &ExecAbilityClass{Abil: AbilityWarcry}
	exec3 := &ExecAbilityClass{Abil: AbilityHarpoon}

	exec1.Next = exec2
	exec2.Prev = exec1
	exec2.Next = exec3
	exec3.Prev = exec2

	d.ExecList = exec1

	// Disable middle ability
	a.DisableAbilityAaa(obj, AbilityWarcry)

	// exec2 should be removed from list and zeroed
	if exec2.Abil != 0 || exec2.Next != nil || exec2.Prev != nil {
		t.Error("Disabled ability should be zeroed")
	}

	// exec1 should now point to exec3
	if exec1.Next != exec3 {
		t.Error("exec1.Next should point to exec3 after removal")
	}
	if exec3.Prev != exec1 {
		t.Error("exec3.Prev should point to exec1 after removal")
	}

	// Disable first ability
	a.DisableAbilityAaa(obj, AbilityBerserk)
	if d.ExecList != exec3 {
		t.Error("d.ExecList should point to exec3 after removing first")
	}

	// Disable last ability
	a.DisableAbilityAaa(obj, AbilityHarpoon)
	if d.ExecList != nil {
		t.Error("d.ExecList should be nil after removing all")
	}
}
