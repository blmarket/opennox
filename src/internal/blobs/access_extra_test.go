package blobs

import (
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAccessStringPanics(t *testing.T) {
	// Test panic on no name
	require.Panics(t, func() {
		a := Access{}
		_ = a.String()
	})

	// Test panic on no blob
	require.Panics(t, func() {
		a := Access{Name: StringPos{Val: "foo"}}
		_ = a.String()
	})

	// Test panic on no offset
	require.Panics(t, func() {
		a := Access{
			Name: StringPos{Val: "foo"},
			Blob: AddrPos{Val: 0x1234},
			Expr: StringPos{Val: "test"},
		}
		_ = a.String()
	})
}

func TestAccessStringWithIndex(t *testing.T) {
	// Test with index that doesn't start with + or -
	a := Access{
		Name: StringPos{Val: "foo"},
		Blob: AddrPos{Val: 0x1234},
		Off:  AddrPos{Val: 10},
		Index: StringPos{
			Val: "bar",
		},
		Expr: StringPos{Val: "test"},
	}
	result := a.String()
	require.Contains(t, result, "+ bar")

	// Test with index starting with +
	a2 := Access{
		Name: StringPos{Val: "foo"},
		Blob: AddrPos{Val: 0x1234},
		Off:  AddrPos{Val: 10},
		Index: StringPos{
			Val: "+ bar",
		},
		Expr: StringPos{Val: "test"},
	}
	result2 := a2.String()
	require.Contains(t, result2, "+ bar")

	// Test with index starting with -
	a3 := Access{
		Name: StringPos{Val: "foo"},
		Blob: AddrPos{Val: 0x1234},
		Off:  AddrPos{Val: 10},
		Index: StringPos{
			Val: "- bar",
		},
		Expr: StringPos{Val: "test"},
	}
	result3 := a3.String()
	require.Contains(t, result3, "- bar")
}

func TestAccessStringWithZeroOffset(t *testing.T) {
	// Test with zero offset but with Off.Pos non-nil
	a := Access{
		Name: StringPos{Val: "foo"},
		Blob: AddrPos{Val: 0x1234},
		Off: AddrPos{
			Val: 0,
			Pos: []token.Pos{1, 2}, // non-nil to avoid panic
		},
		Expr: StringPos{Val: "test"},
	}
	result := a.String()
	require.Contains(t, result, "0")
}
