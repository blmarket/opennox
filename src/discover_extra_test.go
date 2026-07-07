package opennox

import (
	"context"
	"testing"
	"time"

	"github.com/noxworld-dev/lobby"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestGameModeToFlagsExtra(t *testing.T) {
	tests := []struct {
		mode lobby.GameMode
		want uint16
	}{
		{lobby.GameMode(""), 0},
		{lobby.GameMode("unknown"), 0},
	}
	for _, tt := range tests {
		got := gameModeToFlags(tt.mode)
		if got != tt.want {
			t.Errorf("gameModeToFlags(%q) = %v, want %v", tt.mode, got, tt.want)
		}
	}
}

func TestGameModeToFlagsAll(t *testing.T) {
	modes := []lobby.GameMode{
		lobby.ModeKOTR,
		lobby.ModeCTF,
		lobby.ModeFlagBall,
		lobby.ModeChat,
		lobby.ModeArena,
		lobby.ModeElimination,
		lobby.ModeQuest,
		lobby.ModeCoop,
	}
	for _, m := range modes {
		got := gameModeToFlags(m)
		if got == 0 {
			t.Errorf("gameModeToFlags(%v) = 0, want non-zero", m)
		}
	}
}

func TestIsCtxTimeoutExtra(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if !isCtxTimeout(ctx.Err()) {
		t.Error("isCtxTimeout should return true for cancelled context")
	}

	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Nanosecond)
	defer cancel()
	time.Sleep(5 * time.Millisecond)
	if !isCtxTimeout(ctx.Err()) {
		t.Error("isCtxTimeout should return true for timed out context")
	}

	ctx = context.Background()
	if isCtxTimeout(ctx.Err()) {
		t.Error("isCtxTimeout should return false for normal context")
	}
}

func TestIsCtxTimeoutWithValue(t *testing.T) {
	ctx := context.WithValue(context.Background(), "key", "value")
	if isCtxTimeout(ctx.Err()) {
		t.Error("isCtxTimeout should return false for context with value but not cancelled")
	}
}

func TestNoxFlagsGameMode(t *testing.T) {
	if noxflags.GameModeKOTR == 0 {
		t.Error("GameModeKOTR should not be 0")
	}
}
