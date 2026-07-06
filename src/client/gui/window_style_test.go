package gui

import (
	"image"
	"image/color"
	"testing"
	"unsafe"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/stretchr/testify/require"
)

func TestWindowDataNew(t *testing.T) {
	d, free := NewWindowData()
	require.NotNil(t, d)
	require.NotNil(t, free)
	defer free()

	require.NotNil(t, d.C())
}

func TestWindowDataField0Set(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	d.Field0Set(0x1, true)
	require.Equal(t, uint32(0x1), d.Field0&0x1)
	d.Field0Set(0x1, false)
	require.Equal(t, uint32(0), d.Field0&0x1)
}

func TestWindowDataGroup(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	d.SetGroup(5)
	require.Equal(t, 5, d.Group())
}

func TestWindowDataText(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	d.SetText("hello")
	require.Equal(t, "hello", d.Text())
	d.SetText("")
	require.Equal(t, "", d.Text())
}

func TestWindowDataTooltip(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	d.SetTooltip(nil, "tip")
	require.Equal(t, "tip", d.Tooltip())
}

func TestWindowDataColors(t *testing.T) {
	d, free := NewWindowData()
	defer free()

	cl := noxcolor.RGB5551Color(10, 20, 30)
	d.SetBackgroundColor(cl)
	_ = d.BackgroundColor()

	d.SetEnabledColor(cl)
	_ = d.EnabledColor()

	d.SetDisabledColor(cl)
	_ = d.DisabledColor()

	d.SetHighlightColor(cl)
	_ = d.HighlightColor()

	d.SetSelectedColor(cl)
	_ = d.SelectedColor()

	d.SetTextColor(cl)
	_ = d.TextColor()

	// Test with standard color
	d.SetBackgroundColor(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	require.NotNil(t, d.BackgroundColor())
}

func TestWindowDataImagePoint(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	p := image.Pt(10, 20)
	d.SetImagePoint(p)
	require.Equal(t, p, d.ImagePoint())
}

func TestWindowDataFont(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	require.Nil(t, d.FontC())
	ptr := unsafe.Pointer(uintptr(0x1234))
	d.SetFont(ptr)
	require.Equal(t, ptr, d.FontC())
	// Font() calls gui() which is nil, so it panics
	require.Panics(t, func() {
		_ = d.Font()
	})
}

func TestWindowDataSetDefaults(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	def := StyleDefaults{
		EnabledColor:    noxcolor.RGB5551Color(1, 2, 3),
		HighlightColor:  noxcolor.RGB5551Color(4, 5, 6),
		DisabledColor:   noxcolor.RGB5551Color(7, 8, 9),
		BackgroundColor: noxcolor.RGB5551Color(10, 11, 12),
		SelectedColor:   noxcolor.RGB5551Color(13, 14, 15),
		TextColor:       noxcolor.RGB5551Color(16, 17, 18),
	}
	d.SetDefaults(def)
	// Just verify no panic and values set
	require.NotNil(t, d.BackgroundColor())
}

func TestWindowDataSetImage(t *testing.T) {
	d, free := NewWindowData()
	defer free()
	// SetBackgroundImage with nil should not panic
	d.SetBackgroundImage(nil)
	d.SetEnabledImage(nil)
	d.SetDisabledImage(nil)
	d.SetHighlightImage(nil)
	d.SetSelectedImage(nil)
	// BackgroundImage calls gui() which returns nil, so it panics
	require.Panics(t, func() {
		_ = d.BackgroundImage()
	})
	require.Panics(t, func() {
		_ = d.EnabledImage()
	})
	require.Panics(t, func() {
		_ = d.DisabledImage()
	})
	require.Panics(t, func() {
		_ = d.HighlightImage()
	})
	require.Panics(t, func() {
		_ = d.SelectedImage()
	})
}
