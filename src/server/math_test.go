package server

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/stretchr/testify/require"
)

func TestRoundDir(t *testing.T) {
	require.Equal(t, Dir16(0), RoundDir(0))
	require.Equal(t, Dir16(255), RoundDir(255))
	require.Equal(t, Dir16(1), RoundDir(257))
}

func TestRoundCoord(t *testing.T) {
	require.Equal(t, -1, RoundCoord(-1.5))
	require.Equal(t, 0, RoundCoord(0))
	v := RoundCoord(85)
	require.True(t, v >= 0)
}

func TestFabs(t *testing.T) {
	require.Equal(t, float32(1.5), fabs(-1.5))
	require.Equal(t, float32(2.0), fabs(2.0))
}

func TestRoundPos(t *testing.T) {
	p := RoundPos(types.Pointf{X: 0, Y: 0})
	require.Equal(t, image.Pt(0, 0), p)
}

func TestDirFromVec(t *testing.T) {
	d := DirFromVec(types.Pointf{X: 1, Y: 0})
	require.True(t, d <= 255)
	d2 := DirFromVec(types.Pointf{X: 0, Y: 1})
	require.True(t, d2 <= 255)
}

func TestSinCosDir(t *testing.T) {
	c, s := SinCosDir(0)
	require.InDelta(t, 1.0, c, 0.001)
	require.InDelta(t, 0.0, s, 0.001)
	c2, s2 := SinCosDir(64)
	require.True(t, c2 < 1.0)
	_ = s2
}

func TestDoorSize(t *testing.T) {
	p := DoorSize(0)
	require.Equal(t, image.Pt(-23, -23), p)
	p2 := DoorSize(31)
	require.NotEqual(t, image.Pt(0, 0), p2)
}

func TestPointOnTheLine(t *testing.T) {
	p1 := types.Pointf{X: 0, Y: 0}
	p2 := types.Pointf{X: 10, Y: 0}
	a := types.Pointf{X: 5, Y: 0}
	out, ok := PointOnTheLine(p1, p2, a)
	require.True(t, ok)
	require.InDelta(t, 5.0, out.X, 0.001)
	// outside segment
	a2 := types.Pointf{X: 20, Y: 0}
	_, ok2 := PointOnTheLine(p1, p2, a2)
	require.False(t, ok2)
}
