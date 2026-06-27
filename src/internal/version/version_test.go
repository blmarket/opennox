package version

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValid(t *testing.T) {
	if !semverIsValid(DefVersion) {
		t.Fatal("invalid version")
	}
}

func TestSemverIsValid(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"v1.0.0", true},
		{"v1.9.0-dev", true},
		{"1.0.0", false},
		{"invalid", false},
		{"", false},
		{"v1.2.3-pre.1", true},
	}
	for _, c := range cases {
		require.Equal(t, c.want, semverIsValid(c.in), c.in)
	}
}

func TestSemverIsDev(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"v1.0.0", false},
		{"v1.9.0-dev", true},
		{"v1.2.3-pre", true},
		{"invalid", true},
		{"", true},
		{"v2.0.0+meta", false},
	}
	for _, c := range cases {
		require.Equal(t, c.want, semverIsDev(c.in), c.in)
	}
}

func TestSemverLatestFromList(t *testing.T) {
	require.Equal(t, "", semverLatestFromList(nil))
	require.Equal(t, "", semverLatestFromList([]string{"invalid", "foo"}))
	require.Equal(t, "v1.2.3", semverLatestFromList([]string{"v1.0.0", "v1.2.3", "v1.1.0"}))
	require.Equal(t, "v2.0.0", semverLatestFromList([]string{"v1.9.0", "v1.9.0-dev", "v2.0.0"}))
	require.Equal(t, "v1.10.0", semverLatestFromList([]string{"v1.9.0", "v1.10.0", "v1.2.0"}))
}

func TestVersionAndCommit(t *testing.T) {
	v := Version()
	c := Commit()
	require.NotEmpty(t, v)
	require.NotEmpty(t, c)
	require.Equal(t, DefVersion, v)
	require.Equal(t, devCommit, c)
}

func TestIsDev(t *testing.T) {
	// default is dev
	require.True(t, IsDev())
	// test non-dev via temporary globals
	oldV, oldC := version, commit
	defer func() { version, commit = oldV, oldC }()
	version = "v1.0.0"
	commit = "abc123"
	require.False(t, IsDev())
}

func TestClientVersion(t *testing.T) {
	cv := ClientVersion()
	require.Contains(t, cv, Version())
	// default dev should include commit
	require.Contains(t, cv, Commit())

	// test non-dev path
	oldV, oldC := version, commit
	defer func() { version, commit = oldV, oldC }()
	version = "v1.0.0"
	commit = "abc123"
	cv2 := ClientVersion()
	require.Equal(t, "v1.0.0", cv2)
}

func TestLogVersion(t *testing.T) {
	// should not panic
	LogVersion()
}

func TestIsLatest(t *testing.T) {
	// Without network, Latest() falls back to Version(), so IsLatest should be true
	require.True(t, IsLatest())
}
