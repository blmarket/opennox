package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTileDef_Name(t *testing.T) {
	var td TileDef
	name := td.Name()
	require.Equal(t, "", name)

	copy(td.NameBuf[:], "TestTile\x00")
	name = td.Name()
	require.Equal(t, "TestTile", name)
}

func TestInterfaceAddr(t *testing.T) {
	var v any = 42
	addr := interfaceAddr(v)
	require.NotEqual(t, uintptr(0), addr)

	v = "hello"
	addr2 := interfaceAddr(v)
	require.NotEqual(t, uintptr(0), addr2)

	v = nil
	addr3 := interfaceAddr(v)
	require.Equal(t, uintptr(0), addr3)
}

func TestScriptEventType_String(t *testing.T) {
	tests := []struct {
		ev   ScriptEventType
		want string
	}{
		{0, "UnknownEvent"},
		{NoxEventCollide, "Collide"},
		{NoxEventGeneratorSpawn, "GeneratorSpawn"},
		{NoxEventMonsterIdle, "MonsterIdle"},
		{ScriptEventType(999), "server.ScriptEventType(999)"},
	}
	for _, tt := range tests {
		got := tt.ev.String()
		require.Equal(t, tt.want, got)
	}
	// Test PlayerJoin and PlayerLeave - just ensure non-empty
	require.NotEmpty(t, NoxEventPlayerJoin.String())
	require.NotEmpty(t, NoxEventPlayerLeave.String())
}

func TestServer_StartNAT(t *testing.T) {
	s := &Server{}
	s.UseNAT = false
	err := s.startNAT()
	require.NoError(t, err)

	s.UseNAT = true
	// Without GameOnline flag, should return nil
	err = s.startNAT()
	require.NoError(t, err)

	s.stopNAT() // should not panic
	s.stopNAT() // idempotent
}
