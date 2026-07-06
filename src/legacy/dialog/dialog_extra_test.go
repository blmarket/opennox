package dialog

import (
	"testing"

	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/shoenig/test/must"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func TestDialog_GoString(t *testing.T) {
	d := newTestDialog(t)
	s := d.GoString()
	if s == "" {
		t.Error("GoString should not be empty")
	}
	// Should contain Dialog{dir:
	if len(s) < 10 || s[:7] != "Dialog{" {
		t.Errorf("GoString should start with 'Dialog{', got %q", s[:20])
	}
}

func TestDialog_IsFallbackMode(t *testing.T) {
	d := newTestDialog(t)
	// Initially should be 0 or some value
	mode := d.IsFallbackMode()
	_ = mode // Just verify it doesn't panic
}

func TestDialog_Sub_44D930(t *testing.T) {
	d := newTestDialog(t)

	// Initially isEnabled is 0, so should return false
	result := d.Sub_44D930()
	must.EqOp(t, false, result)
}

func TestDialog_Sub_44D8F0(t *testing.T) {
	d := newTestDialog(t)
	// Just verify it doesn't panic
	d.Sub_44D8F0()
}

func TestDialog_Sub_44D8C0(t *testing.T) {
	d := newTestDialog(t)
	// Just verify it doesn't panic
	d.Sub_44D8C0()
}

func TestDialog_Sub_44D5C0(t *testing.T) {
	d := newTestDialog(t)
	// This is already tested in existing test, but verify it doesn't panic
	// Sub_44D5C0 takes a stream and int
	d.Sub_44D5C0(0, 0)
}

func TestNewDialog(t *testing.T) {
	sm := strman.New()
	var (
		isEnabled      uint32
		state          uint32
		someFlag       uint32
		isFallbackMode uint32
		isInitialized  uint32
		audioDriver    ail.Driver
		dword830860    uint32
	)

	d := NewDialog(
		"testdir",
		&isEnabled, &state, &someFlag, &isFallbackMode, &isInitialized, &audioDriver,
		&dword830860,
		func() *strman.StringManager { return sm },
		new(timer.TimerGroup), new(timer.TimerGroup),
		new(timer.TimerGroup), new(timer.TimerGroup),
		nil,
		func() ail.Driver { return 0 },
		func() {}, func() {}, func() bool { return false },
		func(uint32) {}, func() uint32 { return 0 },
	)

	if d == nil {
		t.Fatal("NewDialog returned nil")
	}
	if d.dir != "testdir" {
		t.Errorf("dir = %q, want %q", d.dir, "testdir")
	}
}
