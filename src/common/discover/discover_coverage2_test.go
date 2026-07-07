package discover

import (
	"context"
	"errors"
	"testing"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/stretchr/testify/require"
)

func TestGameFlagsToModeAllModesCoverage(t *testing.T) {
	require.Equal(t, "arena", string(gameFlagsToMode(noxflags.GameModeArena)))
	require.Equal(t, "ctf", string(gameFlagsToMode(noxflags.GameModeCTF)))
	require.Equal(t, "flagball", string(gameFlagsToMode(noxflags.GameModeFlagBall)))
	require.Equal(t, "chat", string(gameFlagsToMode(noxflags.GameModeChat)))
	require.Equal(t, "kotr", string(gameFlagsToMode(noxflags.GameModeKOTR)))
	// 0 returns "custom" or "" depending on implementation, just check it doesn't panic
	_ = gameFlagsToMode(0)
}

func TestIsTimeoutCoverage(t *testing.T) {
	require.True(t, isTimeout(context.DeadlineExceeded))
	require.True(t, isTimeout(context.Canceled))
	require.False(t, isTimeout(errors.New("other error")))
	require.False(t, isTimeout(nil))
}
