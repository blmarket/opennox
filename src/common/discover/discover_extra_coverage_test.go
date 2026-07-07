package discover

import (
	"testing"
	"time"

	"github.com/noxworld-dev/xwis"
	"github.com/stretchr/testify/require"
)

func TestInfoSetDefaults_Extra(t *testing.T) {
	// Test with empty info (all defaults should be set)
	info := &xwis.GameInfo{}
	infoSetDefaults(info)
	require.Equal(t, "OpenNox Server", info.Name)
	require.Equal(t, "estate", info.Map)
	require.Equal(t, xwis.MapTypeChat, info.MapType)
	require.Equal(t, xwis.Res1024x768, info.Resolution)
	require.Equal(t, 32, info.MaxPlayers)

	// Test with partial info (only empty fields should be set)
	info2 := &xwis.GameInfo{
		Name:       "Custom Server",
		Map:        "custom",
		MaxPlayers: 16,
	}
	infoSetDefaults(info2)
	require.Equal(t, "Custom Server", info2.Name)
	require.Equal(t, "custom", info2.Map)
	require.Equal(t, xwis.MapTypeChat, info2.MapType) // Should be set because MapType and Flags are 0
	require.Equal(t, xwis.Res1024x768, info2.Resolution)
	require.Equal(t, 16, info2.MaxPlayers)

	// Test with MapType set (should not overwrite)
	info3 := &xwis.GameInfo{
		MapType: xwis.MapTypeArena,
	}
	infoSetDefaults(info3)
	require.Equal(t, xwis.MapTypeArena, info3.MapType)

	// Test with Flags set (should not overwrite MapType)
	info4 := &xwis.GameInfo{
		Flags: 1,
	}
	infoSetDefaults(info4)
	require.Equal(t, xwis.MapType(0), info4.MapType)
}

func TestNewServerAndClose(t *testing.T) {
	info := GameInfo{}
	s := NewServer(info)
	require.NotNil(t, s)

	// Test Update
	newInfo := GameInfo{}
	newInfo.XWIS.Name = "Test Server"
	s.Update(newInfo)

	// Test Close
	err := s.Close()
	require.NoError(t, err)

	// Test Close again (should be no-op)
	err = s.Close()
	require.NoError(t, err)
}

func TestRegServer_Info(t *testing.T) {
	info := GameInfo{}
	info.XWIS.Name = "Test"
	s := NewServer(info)
	defer s.Close()

	got := s.info()
	require.Equal(t, "Test", got.XWIS.Name)
}

func TestGameRegHost_GameInfo(t *testing.T) {
	info := GameInfo{}
	info.XWIS.Name = "Host Test"
	s := NewServer(info)
	defer s.Close()

	host := gameRegHost{s: s}
	game, err := host.GameInfo(nil)
	require.NoError(t, err)
	require.NotNil(t, game)
}

func TestRegServer_UpdateNonBlocking(t *testing.T) {
	s := NewServer(GameInfo{})
	defer s.Close()

	// Fill the ticks channel to test non-blocking send
	s.ticks <- time.Now()

	// Update should not block even if ticks channel is full
	info := GameInfo{}
	info.XWIS.Name = "Non-blocking"
	s.Update(info)

	// Fill the xwis channel
	s.xwis <- struct{}{}

	// Update should not block even if xwis channel is full
	s.Update(info)
}
