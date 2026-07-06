package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	require.Error(t, run([]string{}))
	require.Error(t, run([]string{"invalid"}))
	require.NoError(t, run([]string{"commit"}))
	require.NoError(t, run([]string{"version"}))
	require.NoError(t, run([]string{"full"}))
	*fOut = "-"
	require.NoError(t, run([]string{"commit"}))
	*fNoNewLine = true
	require.NoError(t, run([]string{"commit"}))
}
