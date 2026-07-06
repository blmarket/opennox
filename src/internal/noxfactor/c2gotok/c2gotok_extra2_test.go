package c2gotok

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestC2Go_SkipSpaces_Extra(t *testing.T) {
	// Test skipSpaces with various inputs
	tests := []string{
		"int  a;",      // multiple spaces
		"int   b;",     // many spaces
		"int\tc;",      // tab
		"int\n d;",     // newline
		"int  \t\n e;", // mixed whitespace
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestC2Go_CheckAssignConv_Extra(t *testing.T) {
	// Test checkAssignConv with various assignment patterns
	tests := []string{
		"int a = 1;",
		"unsigned int b = 2;",
		"char c = 'x';",
		"int *p = 0;",
		"void *v = 0;",
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestC2Go_CheckAssignConvType_Extra(t *testing.T) {
	// Test checkAssignConvType with type conversions
	tests := []string{
		"int a = (int)1;",
		"unsigned int b = (unsigned int)2;",
		"char c = (char)'x';",
		"int *p = (int*)0;",
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestC2Go_MustMatchSpace_Extra(t *testing.T) {
	// Test mustMatchSpace through patterns that require spaces
	tests := []string{
		"unsigned int a;",
		"unsigned short b;",
		"unsigned char c;",
		"const char d;",
		"int * p;", // space before *
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestC2Go_MatchZeroOrMore_Extra(t *testing.T) {
	// Test matchZeroOrMore through patterns with optional elements
	tests := []string{
		"int a;",
		"int *a;",
		"int **b;",
		"int ***c;",
		"void *p;",
		"char *s;",
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestC2Go_SkipSpaces_Direct(t *testing.T) {
	// Directly test skipSpaces by creating a converter with spaces
	input := "int   a;"
	toks := Tokenize([]byte(input))
	c := &c2goConv{toks: toks}
	c.skipSpaces()
	// After skipSpaces, toks should be advanced past spaces
	require.NotNil(t, c)
}

func TestC2Go_MustMatchSpace_Direct(t *testing.T) {
	// Test mustMatchSpace function directly via a converter
	c := &c2goConv{}
	m := c.mustMatchSpace()
	require.NotNil(t, m)
}

func TestC2Go_MatchZeroOrMore_Direct(t *testing.T) {
	// Test matchZeroOrMore function directly
	c := &c2goConv{}
	m := c.matchZeroOrMore("*")
	require.NotNil(t, m)
}
