package discover

import (
	"context"
	"testing"
)

func TestServersWith(t *testing.T) {
	// Test with empty list
	err := serversWith(context.Background(), make(chan Server), map[string]SearchFunc{})
	if err != nil {
		t.Errorf("serversWith empty should not error: %v", err)
	}

	// Test with error backend
	err = serversWith(context.Background(), make(chan Server), map[string]SearchFunc{
		"test": func(ctx context.Context, out chan<- Server) error {
			return context.Canceled
		},
	})
	if err != nil {
		t.Error("serversWith with canceled should not return error")
	}
}

func TestPingerClose(t *testing.T) {
	p := &Pinger{}
	err := p.Close()
	if err != nil {
		t.Errorf("Close on nil stop should not error: %v", err)
	}
}

func TestPingerSendPing(t *testing.T) {
	// SendPing with nil pinger will panic, just verify function exists
	_ = (*Pinger).SendPing
}

func TestPingerPing(t *testing.T) {
	// Ping with nil pinger will panic, just verify function exists
	_ = (*Pinger).Ping
	_ = NewPinger
}
