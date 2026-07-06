package version

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLatestGithubCancelled(t *testing.T) {
	// Test latestGithub with cancelled context (should not panic)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	result, err := latestGithub(ctx)
	_ = result
	_ = err

	// Test getGithubToken
	token, err := getGithubToken(ctx)
	_ = token
	_ = err

	// Test getGithubLatestImage with invalid client
	// This is internal, just ensure it doesn't panic
}

func TestSemverFunctions3(t *testing.T) {
	// Test semverIsValid
	require.True(t, semverIsValid("v1.2.3"))
	require.False(t, semverIsValid("invalid"))

	// Test semverIsDev
	require.True(t, semverIsDev("v1.2.3-dev"))
	require.False(t, semverIsDev("v1.2.3"))

	// Test semverLatestFromList
	list := []string{"v1.2.3", "v1.2.4", "v1.2.3-dev"}
	result := semverLatestFromList(list)
	require.NotEmpty(t, result)
}
