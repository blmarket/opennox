package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerMapGroups(t *testing.T) {
	var s ServerMapGroups
	s.Init()
	require.NotNil(t, s.AllocGroup)
	require.NotNil(t, s.AllocItem)
	require.Nil(t, s.GetFirstMapGroup())
	require.Equal(t, uint32(0), s.NextMapGroupIndex())
	s.Reset()
	s.Free()
}

func TestMapGroupMethods(t *testing.T) {
	var g MapGroup
	require.NotNil(t, g.C())
	require.Equal(t, MapGroupKind(0), g.GroupType())
	require.Equal(t, uint32(0), g.Index())
	require.Equal(t, "", g.ID())
	g.SetID("test")
	require.Equal(t, "test", g.ID())
	require.Nil(t, g.Next())
}
