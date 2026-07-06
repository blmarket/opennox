package gui

import (
	"image/color"
	"testing"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	"github.com/noxworld-dev/opennox/v1/client/noxrender"
	"github.com/stretchr/testify/require"
)

func TestStyleDefaultsSetColors(t *testing.T) {
	var s StyleDefaults
	cl := color.RGBA{R: 1, G: 2, B: 3, A: 4}
	s.SetColors(cl)
	require.Equal(t, cl, s.EnabledColor)
	require.Equal(t, cl, s.DisabledColor)
	require.Equal(t, cl, s.BackgroundColor)
	require.Equal(t, cl, s.HighlightColor)
	require.Equal(t, cl, s.SelectedColor)
	require.Equal(t, cl, s.TextColor)
}

func TestDialogFlagsHas(t *testing.T) {
	var f DialogFlags = DialogOKButton | DialogCancelButton
	require.True(t, f.Has(DialogOKButton))
	require.True(t, f.Has(DialogCancelButton))
	require.False(t, f.Has(DialogYesButton))
}

func TestNewButtonOrCheckboxRaw(t *testing.T) {
	// Nil style should return nil without panicking
	var draw WindowData
	result := NewButtonOrCheckboxRaw(nil, nil, 0, 0, 0, 0, 0, &draw)
	require.Nil(t, result)
}

func TestCWidgetData(t *testing.T) {
	// Test CWidgetData functions from various widget files
	// These should not panic and return nil for nil receiver
	require.NotPanics(t, func() {
		var d1 *EntryFieldData
		_ = d1.CWidgetData()
		var d2 *ScrollListBoxData
		_ = d2.CWidgetData()
		var d3 *RadioButtonData
		_ = d3.CWidgetData()
		var d4 *SliderData
		_ = d4.CWidgetData()
		var d5 *StaticTextData
		_ = d5.CWidgetData()
	})
}

func TestButtonSetImageNil(t *testing.T) {
	// ButtonSetImage with nil win should not panic
	require.NotPanics(t, func() {
		var h noxrender.ImageHandle
		ButtonSetImage(nil, h, h, h, h, h)
	})
}

func TestButtonProc1(t *testing.T) {
	// Test buttonProc1 with simple events
	// Create a minimal window and draw data
	d, free := NewWindowData()
	defer free()
	// buttonProc1 expects win.DrawData().Window to be set for Func94 call
	// We'll test with a simple event that doesn't require Func94
	// Actually WindowFocus event calls Func94, so we need to avoid that or expect panic
	// Let's just verify the function exists and test StaticTextSetText which doesn't call Func94
	win := &Window{}
	win.drawData = *d
	d.Window = win
	e := &StaticTextSetText{Str: "test"}
	resp := buttonProc1(win, e)
	require.Equal(t, RawEventResp(0), resp)
	require.Equal(t, "test", win.DrawData().Text())
}

func TestButtonProc2(t *testing.T) {
	win := &Window{}
	d, free := NewWindowData()
	defer free()
	win.drawData = *d
	d.Window = win

	// Test with Tab key
	e := WindowKeyPress{Key: keybind.KeyTab} // Tab keybind
	resp := ButtonProc2(win, e)
	// Should return 1 for Tab
	require.Equal(t, RawEventResp(1), resp)
}
