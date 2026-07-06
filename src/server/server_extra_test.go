package server

import (
	"testing"
)

func TestMapTraceFlags_Has(t *testing.T) {
	tests := []struct {
		f    MapTraceFlags
		f2   MapTraceFlags
		want bool
	}{
		{MapTraceFlag1, MapTraceFlag1, true},
		{MapTraceFlag1 | MapTraceFlag2, MapTraceFlag1, true},
		{MapTraceFlag1 | MapTraceFlag2, MapTraceFlag3, false},
		{0, MapTraceFlag1, false},
		{MapTraceFlag8, MapTraceFlag8, true},
		{0xFF, MapTraceFlag4, true},
	}
	for _, tt := range tests {
		got := tt.f.Has(tt.f2)
		if got != tt.want {
			t.Errorf("MapTraceFlags(%v).Has(%v) = %v, want %v", tt.f, tt.f2, got, tt.want)
		}
	}
}

func TestServerBalance(t *testing.T) {
	var b serverBalance
	b.Free()

	tag := b.Tag()
	_ = tag

	v := b.Float("nonexistent")
	if v != 0 {
		t.Errorf("Float(nonexistent) = %v, want 0", v)
	}

	v2 := b.FloatInd("nonexistent", 0)
	if v2 != 0 {
		t.Errorf("FloatInd(nonexistent) = %v, want 0", v2)
	}
}

func TestTickHooksExtra(t *testing.T) {
	var s Server
	called := 0
	s.TickCallback(func() {
		called++
	})
	s.TickHook(func() {
		called++
	})
	s.RunTickHooks()
	if called != 2 {
		t.Errorf("called = %v, want 2", called)
	}
}
