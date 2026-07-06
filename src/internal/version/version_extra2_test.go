package version

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetGithubTokenError2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := getGithubToken(ctx)
	require.Error(t, err)
}

func TestGetGithubLatestImageError2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := getGithubLatestImage(ctx, "token")
	require.Error(t, err)
}

func TestLatestGithubError2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := latestGithub(ctx)
	require.Error(t, err)
}

func TestSemverFunctions2(t *testing.T) {
	require.True(t, semverIsValid("v1.2.3"))
	require.False(t, semverIsValid("invalid"))
	require.True(t, semverIsDev("v1.2.3-dev"))
	require.False(t, semverIsDev("v1.2.3"))
	require.Equal(t, "v1.2.3", semverLatestFromList([]string{"v1.0.0", "v1.2.3", "v1.1.0"}))
}
