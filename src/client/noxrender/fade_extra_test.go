package noxrender

import (
	"image"
	"image/color"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestFadeFlags_Extra(t *testing.T) {
	var f fadeFlags = fadeActive | fadeKeep
	require.True(t, f.Has(fadeActive))
	require.True(t, f.Has(fadeKeep))
	require.False(t, f.Has(fadeMenu))
	require.True(t, f.HasAny(fadeMenu|fadeKeep))
	require.False(t, f.HasAny(fadeMenu))
}

func TestFade_Extra(t *testing.T) {
	r := NewRender(nil)
	require.NotNil(t, r)

	// Test HasAny
	var f fadeFlags = fadeActive
	require.True(t, f.HasAny(fadeActive|fadeMenu))
	require.False(t, f.HasAny(fadeMenu))

	// Test FadeReset and FadeDisable
	r.FadeReset()
	r.FadeDisable()

	// Test CheckFade and StopFade on empty
	require.False(t, r.CheckFade(FadeInCinemaKey))
	require.False(t, r.StopFade(FadeInCinemaKey))

	// Test newFade
	fade := r.newFade(FadeInCinemaKey, 10, fadeKeep)
	require.NotNil(t, fade)
	require.True(t, r.CheckFade(FadeInCinemaKey))

	// Test StopFade
	require.True(t, r.StopFade(FadeInCinemaKey))
	require.False(t, r.CheckFade(FadeInCinemaKey))

	// Test DrawFade with no active fades
	result := r.DrawFade(false)
	require.Equal(t, 0, result)

	// Test FadeInCinema with PixBuffer
	pix := noximage.NewImage16(image.Rect(0, 0, 100, 100))
	r.SetPixBuffer(pix)
	d := newRenderData(100, 100)
	r.SetData(d)
	ok := r.FadeInCinema(0.5, 5, color.Black)
	require.True(t, ok)
	require.True(t, r.CheckFade(FadeInCinemaKey))

	// Draw a few frames
	r.DrawFade(false)
	r.DrawFade(false)

	// Stop it
	r.StopFade(FadeInCinemaKey)

	// Test FadeOutCinema
	ok = r.FadeOutCinema(0.5, 5, color.Black)
	require.True(t, ok)
	r.DrawFade(false)
	r.StopFade(FadeOutCinemaKey)

	// Test FadeClearScreen
	ok = r.FadeClearScreen(false, color.Black)
	require.True(t, ok)
	require.True(t, r.CheckFade(FadeClearScreenKey))
	r.DrawFade(false)
	r.StopFade(FadeClearScreenKey)

	// Test FadeInScreen
	doneCalled := false
	ok = r.FadeInScreen(5, false, func() { doneCalled = true })
	require.True(t, ok)
	for i := 0; i < 6; i++ {
		r.DrawFade(false)
	}
	require.True(t, doneCalled)

	// Test FadeOutScreen
	doneCalled = false
	result = r.FadeOutScreen(5, false, func() { doneCalled = true })
	require.Equal(t, 1, result)
	for i := 0; i < 6; i++ {
		r.DrawFade(false)
	}
	require.True(t, doneCalled)

	// Test newFade when full
	r.FadeReset()
	for i := 0; i < 4; i++ {
		r.newFade(FadeKey(i), 1, 0)
	}
	fade = r.newFade(FadeInCinemaKey, 1, 0)
	require.Nil(t, fade)
}
