package server

import (
	"testing"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/stretchr/testify/require"
)

func TestServerStartNAT(t *testing.T) {
	s := &Server{}
	// By default, UseNAT is false, so startNAT should return nil immediately
	err := s.startNAT()
	require.NoError(t, err)

	// Test with UseNAT true but no GameOnline flag
	s.UseNAT = true
	noxflags.ResetGame()
	err = s.startNAT()
	require.NoError(t, err)

	// Test stopNAT with nil stop func (should not panic)
	s.stopNAT()

	// Test stopNAT with a stop func
	called := false
	s.nat.stop = func() { called = true }
	s.stopNAT()
	require.True(t, called)
	require.Nil(t, s.nat.stop)
}
