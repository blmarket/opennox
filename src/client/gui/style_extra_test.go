package gui

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStyleFlags_Extra(t *testing.T) {
	var s StyleFlags
	require.False(t, s.Has(StylePushButton))

	s.Set(StylePushButton | StyleCheckBox)
	require.True(t, s.Has(StylePushButton))
	require.True(t, s.Has(StyleCheckBox))
	require.True(t, s.IsPushButton())
	require.True(t, s.IsCheckBox())
	require.False(t, s.IsRadioButton())

	s = StyleRadioButton
	require.True(t, s.IsRadioButton())

	s = StyleVertSlider
	require.True(t, s.IsVertSlider())

	s = StyleHorizSlider
	require.True(t, s.IsHorizSlider())

	s = StyleScrollListBox
	require.True(t, s.IsScrollListBox())

	s = StyleEntryField
	require.True(t, s.IsEntryField())

	s = StyleStaticText
	require.True(t, s.IsStaticText())

	s = StyleProgressBar
	require.True(t, s.IsProgressBar())

	s = StyleUserWindow
	require.True(t, s.IsUserWindow())

	// Test multiple
	s = StylePushButton | StyleRadioButton | StyleCheckBox | StyleVertSlider | StyleHorizSlider
	require.True(t, s.IsPushButton())
	require.True(t, s.IsRadioButton())
	require.True(t, s.IsCheckBox())
	require.True(t, s.IsVertSlider())
	require.True(t, s.IsHorizSlider())
}
