package gui

import (
	"bytes"
	"image/color"
	"io"
	"testing"

	"github.com/stretchr/testify/require"
)

type byteReader struct {
	*bytes.Reader
}

func (b *byteReader) ReadByte() (byte, error) {
	return b.Reader.ReadByte()
}

func TestReadNextToken_Extra(t *testing.T) {
	// Test normal token with semicolon
	r := &byteReader{bytes.NewReader([]byte("hello;world"))}
	tok, err := ReadNextToken(r)
	require.NoError(t, err)
	require.Equal(t, "hello", tok)

	// Test token with spaces
	r = &byteReader{bytes.NewReader([]byte("  foo bar  ;"))}
	tok, err = ReadNextToken(r)
	require.NoError(t, err)
	require.Equal(t, "foo bar  ", tok)

	// Test EOF with data
	r = &byteReader{bytes.NewReader([]byte("token"))}
	tok, err = ReadNextToken(r)
	require.NoError(t, err)
	require.Equal(t, "token", tok)

	// Test EOF without data
	r = &byteReader{bytes.NewReader([]byte(""))}
	tok, err = ReadNextToken(r)
	require.Equal(t, io.EOF, err)
	require.Equal(t, "", tok)

	// Test multiple tokens
	r = &byteReader{bytes.NewReader([]byte("a;b;c"))}
	tok, err = ReadNextToken(r)
	require.NoError(t, err)
	require.Equal(t, "a", tok)
	tok, err = ReadNextToken(r)
	require.NoError(t, err)
	require.Equal(t, "b", tok)
	tok, err = ReadNextToken(r)
	require.NoError(t, err)
	require.Equal(t, "c", tok)
}

func TestParseColor_Extra(t *testing.T) {
	r, g, b := ParseColor("255 128 64")
	require.Equal(t, 255, r)
	require.Equal(t, 128, g)
	require.Equal(t, 64, b)

	r, g, b = ParseColor("0 0 0")
	require.Equal(t, 0, r)
	require.Equal(t, 0, g)
	require.Equal(t, 0, b)
}

func TestParseColorTransp_Extra(t *testing.T) {
	// Test TRANSPARENT
	c, ok := ParseColorTransp("TRANSPARENT")
	require.True(t, ok)
	require.Equal(t, color.Transparent, c)

	// Test normal color
	c, ok = ParseColorTransp("100 150 200")
	require.True(t, ok)
	nrgba, ok := c.(color.NRGBA)
	require.True(t, ok)
	require.Equal(t, uint8(100), nrgba.R)
	require.Equal(t, uint8(150), nrgba.G)
	require.Equal(t, uint8(200), nrgba.B)
	require.Equal(t, uint8(0xff), nrgba.A)
}

func TestParseNextField_Extra(t *testing.T) {
	// Test with tabs and newlines
	v, rest := ParseNextField("\t\n  hello\tworld\n")
	require.Equal(t, "hello", v)
	require.Equal(t, "world\n", rest)

	// Test single field
	v, rest = ParseNextField("single")
	require.Equal(t, "single", v)
	require.Equal(t, "", rest)
}
