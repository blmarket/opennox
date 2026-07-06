package noxrender

import (
	"image"
	"testing"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestDrawPrimitivesExtra2(t *testing.T) {
	sz := image.Pt(64, 64)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	cl := noxcolor.RGB5551Color(255, 0, 0)
	r.DrawLineHorizontal(5, 10, 20, cl)
	r.DrawLineHorizontal(10, 5, 5, cl)
	r.DrawLine(image.Pt(10, 5), image.Pt(10, 20), cl)

	r.DrawBorder(5, 5, 25, 25, cl)
	r.DrawRectFilledOpaque(0, 0, 10, 10, cl)
	r.DrawRectFilledAlpha(10, 10, 20, 20)

	d.SetAlpha(0x80)
	r.DrawRectFilledAlpha(20, 20, 30, 30)

	r.DrawPoint(image.Pt(5, 5), 1, cl)
	r.DrawPoint(image.Pt(100, 100), 1, cl)

	r.ClearPoints()
	r.AddPoint(image.Pt(1, 1))
	r.AddPointRel(image.Pt(2, 2))
	lp, ok := r.LastPoint(true)
	require.True(t, ok)
	require.Equal(t, image.Pt(3, 3), lp)
	r.DrawLineFromPoints(cl)
	r.DrawVector(image.Pt(10, 10), image.Pt(15, 15), cl)

	r.DrawPointRad(image.Pt(32, 32), 10, cl)

	d.SetClip(true)
	d.SetClipRect(image.Rect(0, 0, 10, 10))
	r.DrawLine(image.Pt(-5, -5), image.Pt(20, 20), cl)
	d.SetClip(false)

	require.NotNil(t, r.Frame())
	r.Set_dword_5d4594_3799484(123)
}
