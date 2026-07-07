package gui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStatusFlags_Extra(t *testing.T) {
	var s StatusFlags
	require.False(t, s.Has(StatusActive))
	require.True(t, s.HasNone(StatusActive))

	s.Set(StatusActive | StatusEnabled)
	require.True(t, s.Has(StatusActive))
	require.True(t, s.Has(StatusEnabled))
	require.False(t, s.HasNone(StatusEnabled))
	require.True(t, s.IsEnabled())
	require.False(t, s.IsHidden())

	s.Set(StatusHidden)
	require.True(t, s.IsHidden())

	// Test Split
	s = StatusActive | StatusToggle | StatusEnabled
	parts := s.Split()
	require.Len(t, parts, 3)
	require.Contains(t, parts, StatusActive)
	require.Contains(t, parts, StatusToggle)
	require.Contains(t, parts, StatusEnabled)

	// Test String
	s = StatusActive | StatusEnabled | StatusHidden
	str := s.String()
	require.Contains(t, str, "Active")
	require.Contains(t, str, "Enabled")
	require.Contains(t, str, "Hidden")

	// Test String with high bits
	s = StatusFlags(1 << 20)
	str = s.String()
	require.Contains(t, str, "0x")

	// Test empty
	s = 0
	require.Equal(t, "", s.String())
	require.Empty(t, s.Split())
}
