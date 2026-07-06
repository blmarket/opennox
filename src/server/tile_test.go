package server

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/stretchr/testify/require"
)

func TestTileDefName(t *testing.T) {
	var td TileDef
	alloc.StrCopyZero(td.NameBuf[:], "test_tile")
	require.Equal(t, "test_tile", td.Name())

	var td2 TileDef
	require.Equal(t, "", td2.Name())
}
