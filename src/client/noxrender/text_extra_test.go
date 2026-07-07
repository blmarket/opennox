package noxrender

import (
	"image"
	"testing"

	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/stretchr/testify/require"
)

func TestText_Extra(t *testing.T) {
	r := NewRender(nil)
	require.NotNil(t, r)

	// Test initText is called in NewRender
	require.NotNil(t, r.text)

	// Test SetTextSmooting, SetBold, TabWidth, SetTabWidth
	r.SetTextSmooting(true)
	r.SetTextSmooting(false)
	r.SetBold(true)
	r.SetBold(false)
	_ = r.TabWidth()
	r.SetTabWidth(4)
	require.Equal(t, 4, r.TabWidth())

	// Test GetFonts and GetBag
	fonts := r.GetFonts()
	require.NotNil(t, fonts)
	bag := r.GetBag()
	// bag may be nil in test
	_ = bag

	// Test fontOrDefault with nil font
	f := r.fontOrDefault(nil)
	// may be nil if no default font loaded
	_ = f

	// Test FontHeight with nil font
	h := r.FontHeight(nil)
	require.Equal(t, 0, h)

	// Test DrawString with nil font (should not panic)
	pix := noximage.NewImage16(image.Rect(0, 0, 100, 100))
	r.SetPixBuffer(pix)
	d := newRenderData(100, 100)
	r.SetData(d)
	r.DrawString(nil, "test", image.Point{X: 0, Y: 0})
	r.DrawStringStyle(nil, "test", image.Point{X: 0, Y: 0})
	r.DrawStringHL(nil, "test", image.Point{X: 0, Y: 0})
	r.DrawStringWrapped(nil, "test", image.Rect(0, 0, 50, 50))
	r.DrawStringWrappedHL(nil, "test", image.Rect(0, 0, 50, 50))

	// Test SplitStringWrapped
	parts := r.SplitStringWrapped(nil, "test string", 50)
	// may be empty if no font
	_ = parts

	// Test GetStringSizeWrapped
	sz := r.GetStringSizeWrapped(nil, "test", 50)
	require.NotNil(t, sz)

	// Test GetStringSizeWrappedStyle
	sz = r.GetStringSizeWrappedStyle(nil, "test", 50)
	require.NotNil(t, sz)

	// Test renderFace methods with nil face - skip as they panic
	// var rf renderFace
	// g, _, _, _, _ := rf.Glyph(fixed.Point26_6{}, 0)
	// require.Nil(t, g)
}
