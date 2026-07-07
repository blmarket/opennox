package opennox

import (
	"context"
	"errors"
	"testing"

	"github.com/noxworld-dev/lobby"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestIsCtxTimeout(t *testing.T) {
	if isCtxTimeout(nil) {
		t.Error("nil should not be timeout")
	}
	if !isCtxTimeout(context.Canceled) {
		t.Error("context.Canceled should be timeout")
	}
	if !isCtxTimeout(context.DeadlineExceeded) {
		t.Error("context.DeadlineExceeded should be timeout")
	}
	if isCtxTimeout(errors.New("other")) {
		t.Error("other error should not be timeout")
	}
}

func TestGameModeToFlags(t *testing.T) {
	tests := []struct {
		mode lobby.GameMode
		want uint16
	}{
		{lobby.ModeKOTR, uint16(noxflags.GameModeKOTR)},
		{lobby.ModeCTF, uint16(noxflags.GameModeCTF)},
		{lobby.ModeFlagBall, uint16(noxflags.GameModeFlagBall)},
		{lobby.ModeChat, uint16(noxflags.GameModeChat)},
		{lobby.ModeArena, uint16(noxflags.GameModeArena)},
		{lobby.ModeElimination, uint16(noxflags.GameModeElimination)},
		{lobby.ModeQuest, uint16(noxflags.GameModeQuest)},
		{lobby.ModeCoop, uint16(noxflags.GameModeCoop)},
		{lobby.GameMode("invalid"), 0},
	}
	for _, tt := range tests {
		got := gameModeToFlags(tt.mode)
		if got != tt.want {
			t.Errorf("gameModeToFlags(%v) = %d, want %d", tt.mode, got, tt.want)
		}
	}
}
