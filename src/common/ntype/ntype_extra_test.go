package ntype

import (
	"image"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPoint32_Extra(t *testing.T) {
	p := Point32{X: 10, Y: 20}
	pt := p.Point()
	require.Equal(t, image.Point{X: 10, Y: 20}, pt)

	p = Point32{X: -5, Y: -10}
	pt = p.Point()
	require.Equal(t, image.Point{X: -5, Y: -10}, pt)

	p = Point32{}
	pt = p.Point()
	require.Equal(t, image.Point{}, pt)
}

type testPlayer struct {
	ind PlayerInd
}

func (p *testPlayer) PlayerIndex() PlayerInd {
	return p.ind
}

func TestPlayer_Extra(t *testing.T) {
	var p Player = &testPlayer{ind: 5}
	require.Equal(t, PlayerInd(5), p.PlayerIndex())

	p = &testPlayer{ind: 0}
	require.Equal(t, PlayerInd(0), p.PlayerIndex())
}
