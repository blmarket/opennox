package noxrender

import (
	"image"
	"testing"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/noximage"
)

func TestDrawPoint(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	cl := noxcolor.RGB5551Color(255, 0, 0)

	// Test DrawPoint at various positions
	r.DrawPoint(image.Pt(10, 10), 1, cl)
	r.DrawPoint(image.Pt(0, 0), 2, cl)
	r.DrawPoint(image.Pt(99, 99), 3, cl)
	r.DrawPoint(image.Pt(-1, -1), 1, cl)   // out of bounds
	r.DrawPoint(image.Pt(200, 200), 1, cl) // out of bounds
	r.DrawPoint(image.Pt(50, 50), 4, cl)
	r.DrawPoint(image.Pt(50, 50), 5, cl)
	r.DrawPoint(image.Pt(50, 50), 6, cl)
}

func TestDrawLineHorizontalAdditional(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	cl := noxcolor.RGB5551Color(0, 255, 0)
	r.DrawLineHorizontal(0, 0, 99, cl)
	r.DrawLineHorizontal(0, 99, 99, cl)
}

func TestSetAlphaEnabledAdditional(t *testing.T) {
	sz := image.Pt(10, 10)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	d.SetAlphaEnabled(true)
	d.SetAlphaEnabled(false)
}

func TestSetFlag16Additional(t *testing.T) {
	sz := image.Pt(10, 10)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	d.SetFlag16(true)
	d.SetFlag16(false)
}
