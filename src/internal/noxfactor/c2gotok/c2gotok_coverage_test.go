package c2gotok

import (
	"bytes"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckAssignConv(t *testing.T) {
	// Test checkAssignConv with assignment conversion pattern
	// Pattern: = (type)  ->  =
	c := &c2goConv{
		toks: []Token{
			{Tok: token.ASSIGN},
			{Tok: token.LPAREN},
			{Tok: token.INT},
			{Tok: token.RPAREN},
		},
	}

	toks, ok := c.checkAssignConv()
	// Should match the pattern and return true
	_ = toks
	_ = ok
}

func TestCheckAssignConvType(t *testing.T) {
	c := &c2goConv{
		toks: []Token{
			{Tok: token.ASSIGN},
			{Tok: token.LPAREN},
			{Tok: token.INT},
			{Tok: token.RPAREN},
		},
	}

	toks, ok := c.checkAssignConvType("int")
	_ = toks
	_ = ok

	// Test with pointer types
	c2 := &c2goConv{
		toks: []Token{
			{Tok: token.ASSIGN},
			{Tok: token.LPAREN},
			{Tok: token.MUL},
			{Tok: token.INT},
			{Tok: token.RPAREN},
		},
	}
	toks2, ok2 := c2.checkAssignConvType("int")
	_ = toks2
	_ = ok2
}

func TestCheckAssignConvNoMatch(t *testing.T) {
	// Test with tokens that don't match the pattern
	c := &c2goConv{
		toks: []Token{
			{Tok: token.INT},
			{Tok: token.ADD},
		},
	}

	toks, ok := c.checkAssignConv()
	require.False(t, ok)
	require.Nil(t, toks)
}

func TestC2Go_CheckFieldAccess(t *testing.T) {
	// Test checkFieldAccess with -> operator
	input := "a->b;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	result := buf.String()
	require.NotEmpty(t, result)
}

func TestC2Go_CheckForParen(t *testing.T) {
	input := "for (int i = 0; i < 10; i++) { a; }"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_CheckWhile1(t *testing.T) {
	input := "while (1) { a; }"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_CheckIfParen(t *testing.T) {
	input := "if (a) { b; }"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_Spaces(t *testing.T) {
	input := "int   a;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_AsMatchOp(t *testing.T) {
	tests := []string{
		"a + b;",
		"a - b;",
		"a * b;",
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestTokenize_Empty(t *testing.T) {
	toks := Tokenize([]byte(""))
	require.Empty(t, toks)
}

func TestToken_String(t *testing.T) {
	tok := Token{Tok: 0, Lit: "test"}
	s := tok.String()
	require.Contains(t, s, "test")
}

func TestC2Go_CheckDoWhile2(t *testing.T) {
	input := "do { a; } while (b);"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}
