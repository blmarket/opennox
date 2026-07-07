package opennox

import (
	"testing"
)

func TestNoxVersion_Build(t *testing.T) {
	tests := []struct {
		v    NoxVersion
		want int
	}{
		{NoxVersion(0x0001039A), 0x039A},
		{NoxVersion(0x000F039A), 0x039A},
		{NoxVersion(0x01020304), 0x0304},
		{0, 0},
	}
	for _, tt := range tests {
		got := tt.v.Build()
		if got != tt.want {
			t.Errorf("NoxVersion(%#x).Build() = %d, want %d", uint32(tt.v), got, tt.want)
		}
	}
}

func TestNoxVersion_Major(t *testing.T) {
	tests := []struct {
		v    NoxVersion
		want int
	}{
		{NoxVersion(0x0001039A), 0},
		{NoxVersion(0x000F039A), 0},
		{NoxVersion(0x01020304), 1},
		{NoxVersion(0x0A0B0C0D), 0x0A},
	}
	for _, tt := range tests {
		got := tt.v.Major()
		if got != tt.want {
			t.Errorf("NoxVersion(%#x).Major() = %d, want %d", uint32(tt.v), got, tt.want)
		}
	}
}

func TestNoxVersion_Minor(t *testing.T) {
	tests := []struct {
		v    NoxVersion
		want int
	}{
		{NoxVersion(0x0001039A), 1},
		{NoxVersion(0x000F039A), 15},
		{NoxVersion(0x01020304), 2},
		{NoxVersion(0x0A0B0C0D), 0x0B},
	}
	for _, tt := range tests {
		got := tt.v.Minor()
		if got != tt.want {
			t.Errorf("NoxVersion(%#x).Minor() = %d, want %d", uint32(tt.v), got, tt.want)
		}
	}
}

func TestNoxVersion_String(t *testing.T) {
	tests := []struct {
		v    NoxVersion
		want string
	}{
		{NoxVersion(0x0001039A), "0.1.922"},
		{NoxVersion(0x000F039A), "0.15.922"},
		{NoxVersion(0x01020304), "1.2.772"},
	}
	for _, tt := range tests {
		got := tt.v.String()
		if got != tt.want {
			t.Errorf("NoxVersion(%#x).String() = %q, want %q", uint32(tt.v), got, tt.want)
		}
	}
}

func TestGetVersionCode(t *testing.T) {
	v := getVersionCode()
	if v != versionCode {
		t.Errorf("getVersionCode() = %#x, want %#x", uint32(v), uint32(versionCode))
	}
}

func TestSetVersionCode(t *testing.T) {
	orig := versionCode
	defer func() { versionCode = orig }()

	newVer := NoxVersion(0x01020304)
	setVersionCode(newVer)
	if versionCode != newVer {
		t.Errorf("setVersionCode() did not set version correctly, got %#x, want %#x", uint32(versionCode), uint32(newVer))
	}

	// Setting same version should not trigger legacy call (just ensure no panic)
	setVersionCode(newVer)
}

func TestNoxClientSetVersion(t *testing.T) {
	orig := versionCode
	defer func() { versionCode = orig }()

	nox_client_setVersion_409AE0(0x0001039A)
	if versionCode != NoxVersion(0x0001039A) {
		t.Errorf("nox_client_setVersion_409AE0() failed, got %#x", uint32(versionCode))
	}
}
