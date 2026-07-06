package version

import (
	"context"
	"testing"
)

func TestLatestExtra5(t *testing.T) {
	_ = Latest()
	_, _ = getGithubToken(context.Background())
	_, _ = latestGithub(context.Background())
	_, _ = getGithubLatestImage(context.Background(), "token")
}
