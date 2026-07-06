package version

import (
	"context"
	"testing"
)

func TestLatestEmpty(t *testing.T) {
	// Test Latest (should not panic)
	_ = Latest()
}

func TestGetGithubToken(t *testing.T) {
	// Test getGithubToken (reads env, should not panic)
	_, _ = getGithubToken(context.Background())
}

func TestLatestGithub(t *testing.T) {
	// Test latestGithub with invalid data (should not panic)
	_, _ = latestGithub(context.Background())
}
