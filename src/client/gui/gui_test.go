package gui

import (
	"image/color"
	"strings"
	"testing"
)

func TestParseNextField(t *testing.T) {
	v, rest := ParseNextField("  foo bar baz")
	if v != "foo" || strings.TrimSpace(rest) != "bar baz" {
		t.Fatalf("unexpected %q %q", v, rest)
	}
	v, rest = ParseNextField("single")
	if v != "single" || rest != "" {
		t.Fatalf("unexpected single")
	}
}

func TestParseNextUintField(t *testing.T) {
	v, rest := ParseNextUintField("123 456")
	if v != 123 || strings.TrimSpace(rest) != "456" {
		t.Fatalf("unexpected")
	}
}

func TestParseNextIntField(t *testing.T) {
	v, rest := ParseNextIntField("-42 foo")
	if v != -42 || strings.TrimSpace(rest) != "foo" {
		t.Fatalf("unexpected")
	}
}

func TestParseColor(t *testing.T) {
	r, g, b := ParseColor("10 20 30")
	if r != 10 || g != 20 || b != 30 {
		t.Fatalf("unexpected color")
	}
}

func TestParseColorTransp(t *testing.T) {
	c, ok := ParseColorTransp("TRANSPARENT")
	if !ok || c != color.Transparent {
		t.Fatalf("expected transparent")
	}
	c, ok = ParseColorTransp("1 2 3")
	if !ok {
		t.Fatalf("expected ok")
	}
	nrgba, ok := c.(color.NRGBA)
	if !ok || nrgba.R != 1 || nrgba.G != 2 || nrgba.B != 3 || nrgba.A != 0xff {
		t.Fatalf("unexpected color %v", c)
	}
}

func TestReadNextToken(t *testing.T) {
	r := strings.NewReader("  foo bar;baz")
	sr := struct {
		*strings.Reader
	}{r}
	// strings.Reader implements ReadByte
	tok, err := ReadNextToken(sr)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if tok != "foo bar" {
		t.Fatalf("expected 'foo bar', got %q", tok)
	}
	tok2, err := ReadNextToken(sr)
	if err != nil {
		t.Fatalf("unexpected err: %v", err)
	}
	if tok2 != "baz" {
		t.Fatalf("expected 'baz', got %q", tok2)
	}
}

func TestState(t *testing.T) {
	var s State
	if s.Current() != StateNone {
		t.Fatalf("expected none")
	}
	if !s.Push(1) {
		t.Fatalf("expected push true")
	}
	if s.Push(1) {
		t.Fatalf("expected push false for same")
	}
	if s.Current() != 1 {
		t.Fatalf("expected 1")
	}
	s.Push(2)
	s.PopUntil(1)
	if s.Current() != 1 {
		t.Fatalf("expected 1 after popuntil")
	}
	s.Pop()
	if s.Current() != StateNone {
		t.Fatalf("expected none after pop")
	}
}

func TestRegisterState(t *testing.T) {
	// reset global map to avoid pollution
	stateByID = make(map[StateID]*state)
	RegisterState(10, "test", func() bool { return true })
	if StateID(10).String() != "test" {
		t.Fatalf("expected test name")
	}
	if StateNone.String() != "<none>" {
		t.Fatalf("expected <none>")
	}
	if StateID(99).String() == "" {
		t.Fatalf("expected non-empty")
	}
	var s State
	s.Push(10)
	if !s.Switch() {
		t.Fatalf("expected switch true")
	}
	// test panic on invalid
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Fatalf("expected panic")
			}
		}()
		RegisterState(0, "bad", func() bool { return true })
	}()
}

func TestStatus(t *testing.T) {
	var st StatusFlags
	if st.Has(StatusEnabled) {
		t.Fatalf("should not have")
	}
	st.Set(StatusEnabled | StatusHidden)
	if !st.Has(StatusEnabled) {
		t.Fatalf("should have enabled")
	}
	if !st.HasNone(0) {
		t.Fatalf("should have none")
	}
	if st.IsEnabled() == false {
		t.Fatalf("expected enabled")
	}
	if !st.IsHidden() {
		t.Fatalf("expected hidden")
	}
	arr := st.Split()
	if len(arr) != 2 {
		t.Fatalf("unexpected split length %d", len(arr))
	}
	if st.String() == "" {
		t.Fatalf("expected non-empty string")
	}
}

func TestStyle(t *testing.T) {
	var s StyleFlags
	s.Set(StylePushButton | StyleRadioButton)
	if !s.Has(StylePushButton) {
		t.Fatalf("expected push button")
	}
	if !s.IsPushButton() || !s.IsRadioButton() {
		t.Fatalf("expected button types")
	}
	if s.IsCheckBox() {
		t.Fatalf("should not be checkbox")
	}
	s = StyleCheckBox
	if !s.IsCheckBox() {
		t.Fatalf("expected checkbox")
	}
	s = StyleVertSlider
	if !s.IsVertSlider() {
		t.Fatalf("expected vert slider")
	}
	s = StyleHorizSlider
	if !s.IsHorizSlider() {
		t.Fatalf("expected horiz slider")
	}
	s = StyleScrollListBox
	if !s.IsScrollListBox() {
		t.Fatalf("expected scroll list")
	}
	s = StyleEntryField
	if !s.IsEntryField() {
		t.Fatalf("expected entry")
	}
	s = StyleStaticText
	if !s.IsStaticText() {
		t.Fatalf("expected static")
	}
	s = StyleProgressBar
	if !s.IsProgressBar() {
		t.Fatalf("expected progress")
	}
	s = StyleUserWindow
	if !s.IsUserWindow() {
		t.Fatalf("expected user window")
	}
}
