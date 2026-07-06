package noxrender

import (
	"image"
	"image/color"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestFadeHasAny(t *testing.T) {
	var f fadeFlags = fadeActive | fadeKeep
	require.True(t, f.Has(fadeActive))
	require.True(t, f.HasAny(fadeActive|fadeMenu))
	require.False(t, f.HasAny(fadeMenu))
	require.True(t, f.HasAny(fadeKeep))
}

func TestFadeOperations(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	// Test FadeReset
	r.FadeReset()
	for _, f := range r.fade.arr {
		require.False(t, f.flags.Has(fadeActive))
	}

	// Test FadeDisable
	r.FadeInCinema(0.5, 10, color.White)
	r.FadeDisable()
	for _, f := range r.fade.arr {
		require.False(t, f.flags.Has(fadeActive))
	}

	// Test newFade
	r.FadeReset()
	f := r.newFade(FadeInCinemaKey, 5, fadeKeep)
	require.NotNil(t, f)
	require.Equal(t, FadeInCinemaKey, f.key)
	require.Equal(t, 5, f.remaining)
	require.True(t, f.flags.Has(fadeActive))
	require.True(t, f.flags.Has(fadeKeep))

	// Fill all fade slots
	r.FadeReset()
	for i := 0; i < 4; i++ {
		f = r.newFade(FadeKey(i), 5, 0)
		require.NotNil(t, f)
	}
	// Next should fail
	f = r.newFade(FadeInCinemaKey, 5, 0)
	require.Nil(t, f)
}

func TestFadeInOutCinema(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)
	r.FadeReset()

	// Test FadeInCinema
	ok := r.FadeInCinema(0.5, 10, color.White)
	require.True(t, ok)
	require.True(t, r.CheckFade(FadeInCinemaKey))

	// Test FadeOutCinema
	r.FadeReset()
	ok = r.FadeOutCinema(0.5, 10, color.White)
	require.True(t, ok)
	require.True(t, r.CheckFade(FadeOutCinemaKey))

	// Test when fade slots are full
	r.FadeReset()
	for i := 0; i < 4; i++ {
		r.newFade(FadeKey(i), 5, 0)
	}
	ok = r.FadeInCinema(0.5, 10, color.White)
	require.False(t, ok)
}

func TestFadeClearScreen(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)
	r.FadeReset()

	// Test FadeClearScreen
	ok := r.FadeClearScreen(false, color.White)
	require.True(t, ok)
	require.True(t, r.CheckFade(FadeClearScreenKey))

	// Test with menu flag
	r.FadeReset()
	ok = r.FadeClearScreen(true, color.White)
	require.True(t, ok)

	// Test when FadeOutScreen is active
	r.FadeReset()
	r.FadeOutScreen(10, false, nil)
	ok = r.FadeClearScreen(false, color.White)
	require.True(t, ok) // Should return true immediately
}

func TestFadeInOutScreen(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)
	r.FadeReset()

	// Test FadeInScreen
	doneCalled := false
	ok := r.FadeInScreen(10, false, func() {
		doneCalled = true
	})
	require.True(t, ok)
	require.True(t, r.CheckFade(FadeInScreenKey))

	// Draw a few frames
	for i := 0; i < 5; i++ {
		r.DrawFade(false)
	}
	require.False(t, doneCalled)

	// Test FadeOutScreen
	r.FadeReset()
	doneCalled = false
	result := r.FadeOutScreen(10, false, func() {
		doneCalled = true
	})
	require.Equal(t, 1, result)
	require.True(t, r.CheckFade(FadeOutScreenKey))

	// Test with menu flag
	r.FadeReset()
	ok = r.FadeInScreen(10, true, nil)
	require.True(t, ok)
	result = r.FadeOutScreen(10, true, nil)
	require.Equal(t, 1, result)
}

func TestStopFade(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)
	r.FadeReset()

	doneCalled := false
	r.FadeInScreen(10, false, func() {
		doneCalled = true
	})

	ok := r.StopFade(FadeInScreenKey)
	require.True(t, ok)
	require.True(t, doneCalled)
	require.False(t, r.CheckFade(FadeInScreenKey))

	// Stop non-existent fade
	ok = r.StopFade(FadeOutScreenKey)
	require.False(t, ok)
}

func TestDrawFadeExtra(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)
	r.FadeReset()

	// Test DrawFade with no active fades
	result := r.DrawFade(false)
	require.Equal(t, 0, result)

	// Test with active fade
	r.FadeInCinema(0.5, 3, color.White)
	for i := 0; i < 5; i++ {
		result = r.DrawFade(false)
		require.Equal(t, 0, result)
	}

	// Test with menu flag
	r.FadeReset()
	r.FadeClearScreen(true, color.White)
	result = r.DrawFade(true)
	require.Equal(t, 0, result)
	result = r.DrawFade(false)
	require.Equal(t, 0, result)
}
