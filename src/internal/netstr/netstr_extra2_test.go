package netstr

import (
	"errors"
	"syscall"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestErrIsInUseExtra(t *testing.T) {
	require.True(t, ErrIsInUse(syscall.EADDRINUSE))
	require.False(t, ErrIsInUse(errors.New("other")))
}

func TestNewConnectErrExtra(t *testing.T) {
	err := NewConnectErr(0, errors.New("fail"))
	require.Equal(t, -1, err.Code)
	require.Contains(t, err.Error(), "CONNECT_SERVER")
	require.Equal(t, "fail", err.Unwrap().Error())
}

func TestHandleExtra(t *testing.T) {
	h := errHandle(-5)
	require.False(t, h.Valid())
	require.False(t, h.IsHost())
}

func TestStreamsExtra(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	require.NotNil(t, s)
	s.reset()
	idx, _ := s.getFreeIndex()
	require.Equal(t, 0, idx)
}
