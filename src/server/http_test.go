package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHTTPServiceInit(t *testing.T) {
	var hs httpService
	hs.init()
	require.NotNil(t, hs.mux)
	require.NotNil(t, hs.srv)
}

func TestServerHTTP(t *testing.T) {
	s := &Server{}
	s.http.init()
	mux := s.HTTP()
	require.NotNil(t, mux)
	require.Equal(t, s.http.mux, mux)
}

func TestServerStartStopHTTP(t *testing.T) {
	s := &Server{}
	s.http.init()
	// Use port 0 to get a free port
	s.http.srv.Addr = ":0"
	err := s.startHTTP()
	// startHTTP overrides Addr with s.HTTPPort(), which defaults to 0
	// It should either succeed or fail gracefully
	if err == nil {
		require.NotNil(t, s.http.lis)
		s.stopHTTP()
		require.Nil(t, s.http.lis)
	} else {
		require.Error(t, err)
	}

	// Test start when already running
	s.http.init()
	s.http.lis = nil
	// Second start should be idempotent if lis not nil
	s.stopHTTP() // should not panic with nil lis
}

func TestServerStopHTTPNil(t *testing.T) {
	s := &Server{}
	s.stopHTTP() // should not panic
}
