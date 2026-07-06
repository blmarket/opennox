package noxflags

import (
	"testing"
)

func TestGameplayFlags(t *testing.T) {
	// Reset
	gameplay = 0

	if GetGamePlay() != 0 {
		t.Error("initial gameplay should be 0")
	}

	SetGamePlay(GameplayFlag1)
	if !HasGamePlay(GameplayFlag1) {
		t.Error("should have GameplayFlag1")
	}
	if GetGamePlay() != GameplayFlag1 {
		t.Error("GetGamePlay should return GameplayFlag1")
	}

	SetGamePlay(GameplayFlag2)
	if !HasGamePlay(GameplayFlag2) {
		t.Error("should have GameplayFlag2")
	}

	UnsetGamePlay(GameplayFlag1)
	if HasGamePlay(GameplayFlag1) {
		t.Error("should not have GameplayFlag1 after unset")
	}
	if !HasGamePlay(GameplayFlag2) {
		t.Error("should still have GameplayFlag2")
	}

	UnsetGamePlay(GameplayFlag2 | GameplayFlag4)
	if HasGamePlay(GameplayFlag2) {
		t.Error("should not have GameplayFlag2 after unset")
	}
}
