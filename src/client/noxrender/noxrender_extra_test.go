package noxrender

import (
	"image"
	"image/color"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestRenderData(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r := NewRender(nil)
	r.SetPixBuffer(img)
	r.SetData(d)

	// Test Data()
	data := r.Data()
	require.NotNil(t, data)

	// Test SetData
	r.SetData(d)
	require.Equal(t, d, r.Data())

	// Test PixBuffer
	pix := r.PixBuffer()
	require.NotNil(t, pix)

	// Test PixBufferRect
	rect := r.PixBufferRect()
	require.False(t, rect.Empty())

	// Test Frame
	frame := r.Frame()
	_ = frame

	// Test RenderData methods
	d.SetClip(true)
	require.True(t, d.Clip())

	d.SetClipRect(image.Rect(0, 0, 50, 50))
	require.Equal(t, image.Rect(0, 0, 50, 50), d.ClipRect())

	d.SetClipRect2(image.Rect(10, 10, 50, 50))
	require.Equal(t, image.Rect(10, 10, 50, 50), d.ClipRect2())

	d.SetRect3(image.Rect(5, 5, 30, 30))
	require.Equal(t, image.Rect(5, 5, 30, 30), d.Rect3())

	d.SetMultiply14(1)
	require.True(t, d.Multiply14())

	require.Equal(t, 0, d.Field15())

	d.SetFlag16(true)
	require.True(t, d.Flag16())

	d.SetColorize17(1)
	require.True(t, d.Colorize17())

	d.SetColorInt44(RGB{1, 2, 3})

	bg := d.BgColor()
	_ = bg

	d.SetSelectColor(10)

	tc := d.TextColor()
	_ = tc
	d.SetTextColor(color.White)

	col := d.Color()
	_ = col
	col2 := d.Color2()
	_ = col2

	d.SetColor(color.White)
	d.SetColor2(color.Black)

	// Test C() and material()
	c := d.C()
	require.NotNil(t, c)

	m := d.material(0)
	require.NotNil(t, m)

	d.SetMaterialRGB(0, 100, 150, 200)
}

func TestNewRenderData(t *testing.T) {
	rd, free := NewRenderData()
	require.NotNil(t, rd)
	require.NotNil(t, free)
	free()
}
