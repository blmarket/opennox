package opennox

import (
	"testing"
)

func TestNoxVersionExtra(t *testing.T) {
	tests := []struct {
		v     NoxVersion
		major int
		minor int
		build int
		str   string
	}{
		{NoxVersion(0x00000000), 0, 0, 0, "0.0.0"},
		{NoxVersion(0xFFFFFFFF), 255, 255, 65535, "255.255.65535"},
		{NoxVersion(0x01000000), 1, 0, 0, "1.0.0"},
		{NoxVersion(0x00010000), 0, 1, 0, "0.1.0"},
		{NoxVersion(0x00000001), 0, 0, 1, "0.0.1"},
	}
	for _, tt := range tests {
		if got := tt.v.Major(); got != tt.major {
			t.Errorf("NoxVersion(%#x).Major() = %d, want %d", uint32(tt.v), got, tt.major)
		}
		if got := tt.v.Minor(); got != tt.minor {
			t.Errorf("NoxVersion(%#x).Minor() = %d, want %d", uint32(tt.v), got, tt.minor)
		}
		if got := tt.v.Build(); got != tt.build {
			t.Errorf("NoxVersion(%#x).Build() = %d, want %d", uint32(tt.v), got, tt.build)
		}
		if got := tt.v.String(); got != tt.str {
			t.Errorf("NoxVersion(%#x).String() = %q, want %q", uint32(tt.v), got, tt.str)
		}
	}
}

func TestNoxVersionConstants(t *testing.T) {
	if noxProtoVersionLegacy.String() != "0.1.922" {
		t.Errorf("noxProtoVersionLegacy.String() = %q, want %q", noxProtoVersionLegacy.String(), "0.1.922")
	}
	if noxProtoVersionHighRes.String() != "0.15.922" {
		t.Errorf("noxProtoVersionHighRes.String() = %q, want %q", noxProtoVersionHighRes.String(), "0.15.922")
	}
}
