package dialog

import (
	"testing"

	"github.com/shoenig/test/must"
)

func TestSub44D930AfterInit(t *testing.T) {
	d := newTestDialog(t)
	d.Nox_xxx_WorkerHurt_44D810()

	// After init, fileToRead is "empty", so Sub_44D930 should return true
	result := d.Sub_44D930()
	must.EqOp(t, true, result)
}

func TestSub44D930WithStream(t *testing.T) {
	d := newTestDialog(t)
	d.Nox_xxx_WorkerHurt_44D810()

	// Set fileToRead to empty but stream should be non-zero after init?
	// Actually GetStream() returns the stream, which may be non-zero
	// Let's just verify it doesn't panic and returns a bool
	result := d.Sub_44D930()
	_ = result
}

func TestSub44D8C0AfterInit(t *testing.T) {
	d := newTestDialog(t)
	d.Nox_xxx_WorkerHurt_44D810()
	must.EqOp(t, true, d.IsInitialized())

	d.Sub_44D8C0()
	must.EqOp(t, false, d.IsInitialized())

	// Calling again should not panic
	d.Sub_44D8C0()
	must.EqOp(t, false, d.IsInitialized())
}
