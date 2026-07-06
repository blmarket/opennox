package music

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
	"github.com/stretchr/testify/require"
)

func TestMusicModuleExtra2(t *testing.T) {
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
	require.NotNil(t, m)

	m.Init()
	m.SetNextMusic(MusicState{MusicIdx: 1, Volume: 50})
	m.SetNextMusic(MusicState{MusicIdx: 2, Volume: 200})
	m.Sub_43DBD0()
	m.Sub_43DBE0()
	_ = m.Sub_43DC40()
	m.Sub_43DD70(3, 75)
	m.Sub_43DDA0()
	_ = m.GetCurrentBlock()
	m.Sub_43D2D0()
	m.Sub_43D650()
	_ = m.startPlay(&MusicState{MusicIdx: 0})
	_ = m.startPlay(&MusicState{MusicIdx: 999})
	v5 = 1
	v6 = 0
	m.Update()
}
