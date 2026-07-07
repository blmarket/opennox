package c2gotok

import (
	"bytes"
	"go/token"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name string
		data []byte
	}{
		{
			name: "simple",
			data: []byte("int x = 42;"),
		},
		{
			name: "with comment",
			data: []byte("int x; // comment\n"),
		},
		{
			name: "empty",
			data: []byte(""),
		},
		{
			name: "whitespace",
			data: []byte("   \n\t  "),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			toks := Tokenize(tt.data)
			// Just verify it doesn't panic
			_ = toks
		})
	}
}

func TestTokenString(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want string
	}{
		{
			name: "illegal",
			tok:  Token{Tok: token.ILLEGAL, Lit: "illegal"},
			want: "illegal",
		},
		{
			name: "comment",
			tok:  Token{Tok: token.COMMENT, Lit: "// comment"},
			want: "// comment",
		},
		{
			name: "semicolon",
			tok:  Token{Tok: token.SEMICOLON, Lit: ";"},
			want: ";",
		},
		{
			name: "operator",
			tok:  Token{Tok: token.ADD, Lit: "+"},
			want: "+",
		},
		{
			name: "keyword",
			tok:  Token{Tok: token.INT, Lit: "int"},
			want: "int",
		},
		{
			name: "literal",
			tok:  Token{Tok: token.INT, Lit: "42"},
			want: "42",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tok.String()
			if got != tt.want {
				t.Errorf("Token.String() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestTokenIsSpace(t *testing.T) {
	tests := []struct {
		name string
		tok  Token
		want bool
	}{
		{
			name: "space",
			tok:  Token{Tok: token.ILLEGAL, Lit: "   "},
			want: true,
		},
		{
			name: "newline",
			tok:  Token{Tok: token.ILLEGAL, Lit: "\n\t"},
			want: true,
		},
		{
			name: "not space",
			tok:  Token{Tok: token.ILLEGAL, Lit: "  x  "},
			want: false,
		},
		{
			name: "not illegal",
			tok:  Token{Tok: token.INT, Lit: "   "},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.tok.IsSpace()
			if got != tt.want {
				t.Errorf("Token.IsSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrint(t *testing.T) {
	toks := []Token{
		{Tok: token.INT, Lit: "int"},
		{Tok: token.ILLEGAL, Lit: " "},
		{Tok: token.IDENT, Lit: "x"},
	}
	var buf bytes.Buffer
	Print(&buf, toks)
	got := buf.String()
	// Should contain "int", " ", and "x"
	if got == "" {
		t.Error("Print produced empty output")
	}
}

func TestTokenizeComplex(t *testing.T) {
	data := []byte(`
		// This is a comment
		int main() {
			int x = 42;
			return x;
		}
	`)
	toks := Tokenize(data)
	if len(toks) == 0 {
		t.Error("Tokenize returned no tokens for complex input")
	}

	// Verify we can print the tokens back
	var buf bytes.Buffer
	Print(&buf, toks)
	if buf.Len() == 0 {
		t.Error("Print produced no output")
	}
}
