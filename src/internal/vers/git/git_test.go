package git

import (
	"os"
	"testing"
)

func TestVersionWithEnv(t *testing.T) {
	os.Setenv("GIT_SHA", "abc123")
	os.Setenv("GIT_TAG", "v1.2.3")
	defer os.Unsetenv("GIT_SHA")
	defer os.Unsetenv("GIT_TAG")

	vers, sha := Version()
	if vers != "v1.2.3" {
		t.Fatalf("expected vers v1.2.3, got %s", vers)
	}
	if sha != "" {
		t.Fatalf("expected empty sha when tag set, got %s", sha)
	}
	fv := FullVersion()
	if fv != "v1.2.3" {
		t.Fatalf("expected full version v1.2.3, got %s", fv)
	}
}

func TestSHAWithEnv(t *testing.T) {
	os.Setenv("GIT_SHA", "def456")
	defer os.Unsetenv("GIT_SHA")
	if got := SHA(); got != "def456" {
		t.Fatalf("expected def456, got %s", got)
	}
	if got := ShortSHA(); got != "def456" {
		t.Fatalf("expected def456, got %s", got)
	}
}

func TestVersionNoEnv(t *testing.T) {
	os.Unsetenv("GIT_SHA")
	os.Unsetenv("GIT_TAG")
	vers, sha := Version()
	if vers == "" {
		t.Fatalf("expected default version")
	}
	_ = sha
	_ = FullVersion()
	_ = SHA()
	_ = ShortSHA()
}
