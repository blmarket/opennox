package opennox

import (
	"testing"

	"github.com/noxworld-dev/noxscript/ns/asm"
)

func TestNoxScriptShouldReadMoreXxx(t *testing.T) {
	// Test true cases
	trueCases := []asm.Builtin{9, 10, 46, 47, 190, 126}
	for _, fi := range trueCases {
		if !nox_script_shouldReadMoreXxx(fi) {
			t.Fatalf("expected true for builtin %d", fi)
		}
	}

	// Test false cases
	falseCases := []asm.Builtin{0, 1, 8, 11, 45, 48, 125, 127, 189, 191}
	for _, fi := range falseCases {
		if nox_script_shouldReadMoreXxx(fi) {
			t.Fatalf("expected false for builtin %d", fi)
		}
	}
}

func TestNoxScriptShouldReadEvenMoreXxx(t *testing.T) {
	if !nox_script_shouldReadEvenMoreXxx(asm.BuiltinSetDialog) {
		t.Fatal("expected true for BuiltinSetDialog")
	}

	if nox_script_shouldReadEvenMoreXxx(0) {
		t.Fatal("expected false for builtin 0")
	}
	if nox_script_shouldReadEvenMoreXxx(asm.Builtin(999)) {
		t.Fatal("expected false for unknown builtin")
	}
}

func TestErrStopScript(t *testing.T) {
	if errStopScript == nil {
		t.Fatal("errStopScript should not be nil")
	}
	if errStopScript.Error() != "noxscript: exit" {
		t.Fatalf("unexpected error message: %s", errStopScript.Error())
	}
}
