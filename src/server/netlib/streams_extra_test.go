package netlib

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRecvFlags_Extra(t *testing.T) {
	var f RecvFlags = RecvCanRead | RecvNoHooks
	require.True(t, f.Has(RecvCanRead))
	require.True(t, f.Has(RecvNoHooks))
	require.False(t, f.Has(RecvJustOne))

	f = RecvJustOne
	require.True(t, f.Has(RecvJustOne))
	require.False(t, f.Has(RecvCanRead))

	f = 0
	require.False(t, f.Has(RecvCanRead))
	require.False(t, f.Has(RecvNoHooks))
	require.False(t, f.Has(RecvJustOne))

	f = RecvCanRead | RecvNoHooks | RecvJustOne
	require.True(t, f.Has(RecvCanRead))
	require.True(t, f.Has(RecvNoHooks))
	require.True(t, f.Has(RecvJustOne))
}
