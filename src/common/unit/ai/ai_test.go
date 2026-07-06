package ai

import (
	"testing"
)

func TestActionType_String(t *testing.T) {
	tests := []struct {
		action ActionType
		want   string
	}{
		{ACTION_IDLE, "ACTION_IDLE"},
		{ACTION_WAIT, "ACTION_WAIT"},
		{ACTION_HUNT, "ACTION_HUNT"},
		{ACTION_FIGHT, "ACTION_FIGHT"},
		{DEPENDENCY_OR, "DEPENDENCY_OR"},
		{DEPENDENCY_NOT_MOVED, "DEPENDENCY_NOT_MOVED"},
		{ActionType(999), "ai.Action(999)"},
	}
	for _, tt := range tests {
		got := tt.action.String()
		if got != tt.want {
			t.Errorf("ActionType(%d).String() = %q, want %q", tt.action, got, tt.want)
		}
	}
}

func TestActionType_IsCondition(t *testing.T) {
	tests := []struct {
		action ActionType
		want   bool
	}{
		{ACTION_IDLE, false},
		{ACTION_WAIT, false},
		{ACTION_HUNT, false},
		{DEPENDENCY_OR, true},
		{DEPENDENCY_TIME, true},
		{DEPENDENCY_NOT_MOVED, true},
		{ActionType(39), false}, // Before DEPENDENCY_OR
		{ActionType(72), false}, // After DEPENDENCY_NOT_MOVED
	}
	for _, tt := range tests {
		got := tt.action.IsCondition()
		if got != tt.want {
			t.Errorf("ActionType(%d).IsCondition() = %v, want %v", tt.action, got, tt.want)
		}
	}
}
