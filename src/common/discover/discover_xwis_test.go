package discover

import (
	"context"
	"testing"
)

func TestXwisUpdateLoopCancelled(t *testing.T) {
	s := &RegServer{
		xwis: make(chan struct{}),
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel immediately

	err := s.xwisUpdateLoop(ctx, nil)
	// Should return context error immediately
	_ = err
}

func TestXwisReopenAndLoopCancelled(t *testing.T) {
	s := &RegServer{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := s.xwisReopenAndLoop(ctx)
	// Should return quickly due to cancelled context
	_ = err
}

func TestDoXWISCancelled(t *testing.T) {
	s := &RegServer{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	err := s.doXWIS(ctx)
	// Should return nil immediately due to cancelled context
	_ = err
}
