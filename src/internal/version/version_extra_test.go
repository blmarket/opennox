package version

import (
	"context"
	"testing"
)

func TestVersionExtra(t *testing.T) {
	v := Version()
	if v == "" {
		t.Error("Version should not be empty")
	}
}

func TestCommitExtra(t *testing.T) {
	c := Commit()
	if c == "" {
		t.Error("Commit should not be empty")
	}
}

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
