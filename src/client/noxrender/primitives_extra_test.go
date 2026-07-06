package noxrender

import (
	"image"
	"testing"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestDrawLineHorizontalVertical(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	cl := noxcolor.RGB5551Color(255, 0, 0)

	// Horizontal lines
	r.DrawLineHorizontal(10, 20, 50, cl)
	r.DrawLineHorizontal(50, 20, 10, cl)   // reverse
	r.DrawLineHorizontal(-10, 30, 110, cl) // clipped
	r.DrawLineHorizontal(10, -5, 50, cl)   // out of bounds
	r.DrawLineHorizontal(10, 200, 50, cl)  // out of bounds

	// Vertical lines via DrawLine (which calls drawLineVertical)
	r.DrawLine(image.Pt(30, 10), image.Pt(30, 50), cl)
	r.DrawLine(image.Pt(40, 50), image.Pt(40, 10), cl) // reverse

	// With clipping
	d.SetClip(true)
	d.SetClipRect(image.Rect(20, 20, 80, 80))
	r.DrawLineHorizontal(10, 30, 90, cl)
	r.DrawLine(image.Pt(50, 10), image.Pt(50, 90), cl)
}

func TestDrawFade(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	// Draw fade (bool param)
	r.DrawFade(false)
	r.DrawFade(true)

	// Stop fade
	r.StopFade(FadeKey(0))
}

func TestDrawRect(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	cl := noxcolor.RGB5551Color(0, 255, 0)

	r.DrawRectFilledOpaque(10, 10, 40, 40, cl)
	r.DrawRectFilledAlpha(60, 60, 30, 30)

	d.SetAlpha(128)
	d.SetAlphaEnabled(true)
	r.DrawRectFilledAlpha(20, 60, 20, 20)
}

func TestDrawBorder(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	cl := noxcolor.RGB5551Color(255, 255, 0)
	r.DrawBorder(10, 10, 80, 80, cl)
}

func TestSetClip(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	d.SetClip(true)
	require.True(t, d.Clip())

	d.SetClip(false)
	require.False(t, d.Clip())

	d.SetAlphaEnabled(true)
	require.True(t, d.IsAlphaEnabled())
}
