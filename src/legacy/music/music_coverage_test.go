package music

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
	"github.com/shoenig/test/must"
)

func newTestModuleCoverage(t *testing.T) *Module {
	timer.PlatformTicks = func() uint64 { return 100 }
	t.Cleanup(func() { timer.PlatformTicks = nil })

	var v1, v2, v3, v4, v5, v6 uint32
	var drv ail.Driver
	tg := &timer.TimerGroup{}
	var ptg *timer.TimerGroup = tg

	m := NewModule(
		"testdir",
		&v1, &v2, &drv,
		&v3, &v4, &v5,
		tg, &ptg, &v6,
		func() ail.Driver { return 0 },
		func() bool { return false },
		func() string { return "" },
		func() uint64 { return 100 },
	)
	return m
}

func TestModulePauseUnpause(t *testing.T) {
	m := newTestModuleCoverage(t)

	// pause and unpause should not panic with nil stream
	m.pause()
	m.unpause()

	// GetCurrentBlock should not panic
	block := m.GetCurrentBlock()
	_ = block
}

func TestModuleSub43D650(t *testing.T) {
	m := newTestModuleCoverage(t)

	// Sub_43D650 should not panic
	m.Sub_43D650()
	must.EqOp(t, uint32(0), m.currentPlaying.MusicIdx)
}

func TestModuleSub43DC40(t *testing.T) {
	m := newTestModuleCoverage(t)

	// Sub_43DC40 should not panic
	result := m.Sub_43DC40()
	_ = result
}
