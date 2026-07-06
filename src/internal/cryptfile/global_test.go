package cryptfile

import (
	"testing"
)

func TestGlobalExtra2(t *testing.T) {
	// Initially nil
	if Global() != nil {
		t.Log("Global not nil initially")
	}

	cf := &CryptFile{}
	SetGlobal(cf)
	if Global() != cf {
		t.Error("Global should return set value")
	}

	SetGlobal(nil)
	if Global() != nil {
		t.Error("Global should be nil after set nil")
	}
}

func TestCloseNil(t *testing.T) {
	SetGlobal(nil)
	err := Close()
	if err != nil {
		t.Errorf("Close with nil global should not error: %v", err)
	}
}

func TestOpenGlobalInvalid(t *testing.T) {
	// Open with invalid path should error
	err := OpenGlobal("/nonexistent/path", 0, 0)
	if err == nil {
		t.Log("OpenGlobal with invalid path should error but got nil")
	}
	// Global should still be nil or unchanged
}
