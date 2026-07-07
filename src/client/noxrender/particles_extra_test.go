package noxrender

import (
	"image"
	"image/color"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestParticles_Extra(t *testing.T) {
	r := NewRender(nil)
	require.NotNil(t, r)

	// initParticles is called in NewRender
	require.True(t, r.Part.RenderGlow)
	require.NotNil(t, r.Part.byOpts)

	// Test NewParticle
	pix := noximage.NewImage16(image.Rect(0, 0, 100, 100))
	r.SetPixBuffer(pix)
	d := newRenderData(100, 100)
	r.SetData(d)
	r.Data().SetField262(5)
	r.Data().SetColorInt54(RGB{R: 100, G: 150, B: 200})

	p := r.NewParticle(0, 10)
	require.NotNil(t, p)
	require.NotNil(t, p.img)

	// Test cached particle
	p2 := r.NewParticle(0, 10)
	require.Equal(t, p, p2)

	// Test DrawAt
	p.DrawAt(image.Point{X: 10, Y: 10})

	// Test DrawGlow
	r.DrawGlow(image.Point{X: 20, Y: 20}, color.White, 3, 4)

	// Test DrawGlow with RenderGlow false
	r.Part.RenderGlow = false
	r.DrawGlow(image.Point{X: 30, Y: 30}, color.White, 3, 4)
	r.Part.RenderGlow = true

	// Test DrawParticles49ED80 - just test the false path
	r.ClearPoints()
	require.False(t, r.DrawParticles49ED80(2))

	// Test genParticle with rad 1
	opt := particleOpt{rad: 1, blur: 0, intens: 10, color: RGB{R: 100, G: 100, B: 100}}
	img := genParticle(opt)
	require.NotNil(t, img)
	img.Free()

	// Test Free
	p.Free()
	require.Nil(t, p.r)

	// Test Part.Free
	r.Part.Free()
}
