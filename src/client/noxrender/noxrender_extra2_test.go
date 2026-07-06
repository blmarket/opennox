package noxrender

import (
	"image"
	"image/color"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestRenderDataExtra(t *testing.T) {
	d := newRenderData(100, 100)

	d.SetLightColor(RGB{10, 20, 30})
	lc := d.GetLightColor()
	require.Equal(t, RGB{10, 20, 30}, lc)

	d.SetMaterial(1, color.White)

	d.SetColorMultA(Color16{R: 128, G: 128, B: 128})
	_ = d.ColorMultA()

	d.SetField262(42)
	require.Equal(t, uint32(42), d.Field262())

	_ = d.ShouldDrawText()

	d.SetColorInt54(RGB{1, 2, 3})
	require.Equal(t, RGB{1, 2, 3}, d.ColorInt54())
}

func TestNoxRenderExtra(t *testing.T) {
	sz := image.Pt(100, 100)
	d := newRenderData(sz.X, sz.Y)
	r := NewRender(nil)
	r.SetData(d)

	img := noximage.NewImage16(image.Rectangle{Max: sz})
	r.SetPixBuffer(img)

	r.CopyPixBuffer()
	r.ClearScreen(color.White)

	ci := r.ColorIntensity(10, 20, 30)
	_ = ci

	r.Set_dword_5d4594_3799484(1)
	r.SetInterlacing(true, 0)
	r.Reset_dword_5d4594_3799476()
	_ = r.Get_dword_5d4594_3799476()
}
