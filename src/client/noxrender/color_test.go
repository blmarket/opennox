package noxrender

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColor16(t *testing.T) {
	c := Color16{R: 100, G: 150, B: 200}
	m := c.Make()
	require.NotZero(t, m)
	require.Equal(t, uint16(m), c.Make16())

	s := SplitColor(m)
	require.InDelta(t, 100, s.R, 10)
	require.InDelta(t, 150, s.G, 10)
	require.InDelta(t, 200, s.B, 10)

	s2 := SplitColor16(uint16(m))
	require.InDelta(t, 100, s2.R, 10)

	c4444, a := SplitColor4444(0xF123)
	require.NotZero(t, c4444.R+c4444.G+c4444.B)
	require.NotZero(t, a)

	// Saturate
	cs := Color16{R: 300, G: 500, B: 1000}.Saturate()
	require.Equal(t, uint16(0xff), cs.R)
	require.Equal(t, uint16(0xff), cs.G)
	require.Equal(t, uint16(0xff), cs.B)

	// Mult
	cm := Color16{R: 128, G: 128, B: 128}.Mult(Color16{R: 128, G: 128, B: 128})
	require.Equal(t, uint16(64), cm.R)

	// MultI
	cmi := Color16{R: 200, G: 100, B: 50}.MultI(128)
	require.Equal(t, uint16(100), cmi.R)

	// Over
	co := Color16{R: 0, G: 0, B: 0}.Over(Color16{R: 100, G: 100, B: 100})
	require.Equal(t, uint16(50), co.R)

	// Over2
	co2 := Color16{R: 0, G: 0, B: 0}.Over2(Color16{R: 100, G: 0, B: 0})
	require.Equal(t, uint16(25), co2.R)

	// OverAlpha
	coa := Color16{R: 0, G: 0, B: 0}.OverAlpha(128, Color16{R: 200, G: 0, B: 0})
	require.Equal(t, uint16(100), coa.R)
}

func TestFadeFunctions(t *testing.T) {
	// basic coverage for fade methods that don't require full render context
	r := NewRender(nil)
	r.FadeReset()
	r.FadeDisable()
	r.DrawFade(false)
	r.StopFade(FadeInCinemaKey)
	r.CheckFade(FadeInCinemaKey)
	// FadeIn etc require PixBuffer; skip to avoid panic, we already covered simple paths
}
