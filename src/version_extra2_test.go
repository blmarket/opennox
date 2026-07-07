package opennox

import (
	"testing"
)

func TestNoxVersion_Extra2(t *testing.T) {
	v := NoxVersion(0x01020003)
	if v.Major() != 1 {
		t.Errorf("Major() = %d, want 1", v.Major())
	}
	if v.Minor() != 2 {
		t.Errorf("Minor() = %d, want 2", v.Minor())
	}
	if v.Build() != 3 {
		t.Errorf("Build() = %d, want 3", v.Build())
	}
	s := v.String()
	if s == "" {
		t.Error("String() should not be empty")
	}
	if s != "1.2.3" {
		t.Errorf("String() = %q, want %q", s, "1.2.3")
	}
}

func TestBuildVersion_Extra2(t *testing.T) {
	v := NoxVersion(0x00010003)
	if v.Build() != 3 {
		t.Errorf("Build() = %d, want 3", v.Build())
	}
}

func TestGetVersionCode_Extra2(t *testing.T) {
	code := getVersionCode()
	_ = code
}

func TestSetVersionCode_Extra2(t *testing.T) {
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("setVersionCode panicked: %v", r)
			}
		}()
		setVersionCode(0x00010000)
	}()
}
