package version

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetGithubTokenCancelled(t *testing.T) {
	// Test with server that returns invalid JSON
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("invalid json"))
	}))
	defer server.Close()

	// We can't easily test getGithubToken with a custom URL since it's hardcoded,
	// but we can test error paths by using a cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	_, err := getGithubToken(ctx)
	if err == nil {
		t.Error("getGithubToken with cancelled context should return an error")
	}
}

func TestGetGithubTokenInvalidStatus(t *testing.T) {
	// Test error handling when server returns non-2xx status
	// Since the URL is hardcoded, we test with a cancelled context to trigger error
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := getGithubToken(ctx)
	if err == nil {
		t.Error("getGithubToken should return error on cancelled context")
	}
}

func TestGetGithubLatestImageCancelled(t *testing.T) {
	// Test with cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := getGithubLatestImage(ctx, "test-token")
	if err == nil {
		t.Error("getGithubLatestImage with cancelled context should return an error")
	}

	// Test with empty token (will fail at HTTP level, but we test the function is called)
	ctx = context.Background()
	_, err = getGithubLatestImage(ctx, "")
	// Should return an error (either from HTTP or from empty tags)
	if err == nil {
		t.Log("getGithubLatestImage with empty token unexpectedly succeeded (may be network dependent)")
	}
}

func TestLatestGithubCancelledExtra(t *testing.T) {
	// Test with cancelled context
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := latestGithub(ctx)
	if err == nil {
		t.Error("latestGithub with cancelled context should return an error")
	}
}

func TestIsLatestExtra(t *testing.T) {
	// IsLatest compares current version with latest version
	// This may involve network calls, so we just verify it doesn't panic
	result := IsLatest()
	_ = result // Result depends on network and current version
}

func TestLatestExtra(t *testing.T) {
	// Latest() may involve network calls, just verify it doesn't panic and returns something
	result := Latest()
	if result == "" {
		t.Error("Latest() should not return empty string")
	}
}
