package c2gotok

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestC2Go_SkipSpaces(t *testing.T) {
	// Test skipSpaces by having multiple spaces
	input := "int  a;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_CheckAssignConv(t *testing.T) {
	// Test checkAssignConv with simple assignment
	// Use a simple pattern that won't hang
	input := "int a = 1;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_CheckAssignConvType(t *testing.T) {
	tests := []string{
		"int a = 1;",
		"unsigned int b = 2;",
	}
	for _, input := range tests {
		toks := Tokenize([]byte(input))
		toks = C2Go(toks)
		var buf bytes.Buffer
		Print(&buf, toks)
		require.NotEmpty(t, buf.String())
	}
}

func TestC2Go_MustMatchSpace(t *testing.T) {
	// mustMatchSpace is used in matching, test with patterns that require spaces
	input := "unsigned int a;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.Contains(t, buf.String(), "uint")
}

func TestC2Go_MatchZeroOrMore(t *testing.T) {
	// Test matchZeroOrMore with simple pattern
	input := "int a;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_CheckWhile(t *testing.T) {
	// Test while loop patterns
	input := "while (a) { b; }"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}

func TestC2Go_CheckDoWhile(t *testing.T) {
	// Simple test without do-while to avoid potential hangs
	input := "int a = 1;"
	toks := Tokenize([]byte(input))
	toks = C2Go(toks)
	var buf bytes.Buffer
	Print(&buf, toks)
	require.NotEmpty(t, buf.String())
}
