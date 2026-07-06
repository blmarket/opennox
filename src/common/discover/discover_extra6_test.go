package discover

import (
	"context"
	"testing"
	"time"
)

func TestEachServerEmptyAdditional(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	err := EachServer(ctx, func(srv Server) error {
		return nil
	})
	_ = err
}

func TestListServersEmptyAdditional(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	servers, err := ListServers(ctx)
	_ = servers
	_ = err
}

func TestIsTimeoutAdditional(t *testing.T) {
	if !isTimeout(context.DeadlineExceeded) {
		t.Error("isTimeout should return true for DeadlineExceeded")
	}
}
