package noxrender

import (
	"image"
	"image/color"
	"image/draw"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/noximage"
)

type RenderData interface {
	Frame() uint32
	ClipRect() image.Rectangle
	ClipRect2() image.Rectangle

	Clip() bool
	Multiply14() bool
	Flag16() bool
	Colorize17() bool

	ColorMultA() Color16
	ColorMultOp(op int) Color16

	IsAlphaEnabled() bool
	Alpha() byte

	Color() color.Color

	RenderDataText
}

func NewRender() *NoxRender {
	r := &NoxRender{}
	r.initText()
	r.initColorTablesRev()
	return r
}

type NoxRender struct {
	p   RenderData
	pix *noximage.Image16

	colors struct {
		revTable []byte // map[Color16]byte
	}
	points []image.Point
	text   noxRenderText
	fade   noxRenderFade

	dword_5d4594_3799476 int
	dword_5d4594_3799484 uint32
	interlacingY         int
	interlacing          bool

	HookImageDrawXxx func(pos image.Point, sz image.Point)
}

func (r *NoxRender) ColorIntensity(cr, cg, cb byte) byte {
	v := ((cb & 0xF8) >> 3) | ((cg & 0xF8) << 2) | ((cr & 0xF8) << 7)
	return r.colors.revTable[v]
}

func (r *NoxRender) PixBufferRect() image.Rectangle {
	return r.pix.Rect
}

func (r *NoxRender) PixBuffer() *noximage.Image16 {
	return r.pix
}

func (r *NoxRender) SetPixBuffer(pix *noximage.Image16) {
	r.pix = pix
}

func (r *NoxRender) Data() RenderData {
	return r.p
}

func (r *NoxRender) SetData(p RenderData) {
	r.p = p
}

func (r *NoxRender) Frame() uint32 {
	return r.p.Frame()
}

func (r *NoxRender) initColorTablesRev() {
	const max = 0x7FFF
	r.colors.revTable = make([]byte, max+3)
	for i := 0; i <= max; i++ {
		c := SplitColor16(uint16(i))
		r.colors.revTable[i] = byte((28*(c.B|7) + 150*(c.G|7) + 76*(c.R|7)) >> 8)
	}
}

func (r *NoxRender) CopyPixBuffer() *image.NRGBA {
	img := image.NewNRGBA(r.pix.Rect)
	draw.Draw(img, img.Rect, r.pix, r.pix.Rect.Min, draw.Src)
	return img
}

func (r *NoxRender) ClearScreen(cl color.Color) {
	u16 := noxcolor.ModelRGBA5551.Convert16(cl).Color16()
	for i := range r.pix.Pix {
		r.pix.Pix[i] = u16
	}
}

func (r *NoxRender) Set_dword_5d4594_3799484(v int) { // sub_47D370
	if v < 0 {
		v = 0
	}
	r.dword_5d4594_3799484 = uint32(v)
}

func (r *NoxRender) SetInterlacing(enable bool, y int) {
	r.interlacing = enable
	r.interlacingY = y & 0x1
}

func (r *NoxRender) Reset_dword_5d4594_3799476() {
	r.dword_5d4594_3799476 = 0
}

func (r *NoxRender) Get_dword_5d4594_3799476() int {
	return r.dword_5d4594_3799476
}
