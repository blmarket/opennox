package music

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/shoenig/test/must"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func platformTicks() uint64 {
	return uint64(platform.Ticks() / time.Millisecond)
}

func newTestMusic(t *testing.T) *Module {
	timer.PlatformTicks = platformTicks

	handles.Init()
	t.Cleanup(func() {
		handles.Release()
	})
	// Initialize audio.
	ail.Startup()
	dr := ail.WaveOutOpen()

	sub_43F130 := func() ail.Driver { return dr }
	checkDialogs := func() bool { return false }
	sub_413890 := func() string { return "" }

	var (
		dword_5d4594_816368      uint32
		dword_5d4594_816372      uint32
		dword_5d4594_816376      ail.Driver
		dword_587000_93156       uint32
		dword_587000_93160       uint32
		dword_5d4594_816340      uint32
		counter_5d4594_816244    timer.TimerGroup
		ptr_counter_587000_81128 *timer.TimerGroup = new(timer.TimerGroup)
		dword_5d4594_816348      uint32
		musicStateArray          = new([3][6]MusicState)
		musicIndexArray          = new([3]uint32)
	)

	return NewModule(
		"testdata",
		&dword_5d4594_816368,
		&dword_5d4594_816372,
		musicStateArray,
		musicIndexArray,
		&dword_5d4594_816376,
		&dword_587000_93156,
		&dword_587000_93160,
		&dword_5d4594_816340,
		&counter_5d4594_816244,
		&ptr_counter_587000_81128,
		&dword_5d4594_816348,
		sub_43F130,
		checkDialogs,
		sub_413890,
		platformTicks,
	)
}

func TestMusicIntegration(t *testing.T) {
	m := newTestMusic(t)

	// Initial state should be uninitialized
	must.EqOp(t, uint32(0), *m.dword_5d4594_816340)
	must.EqOp(t, MusicState{}, m.GetCurrentBlock())

	// Initialize the music module
	m.Init()
	must.EqOp(t, uint32(1), *m.dword_5d4594_816340)
	must.EqOp(t, uint32(1), *m.dword_587000_93156)

	// Test setting next music
	testMusic := MusicState{MusicIdx: 1, Volume: 50, Position: 0, D: 0}
	m.SetNextMusic(testMusic)

	// Test getting current block
	currentBlock := m.GetCurrentBlock()
	must.EqOp(t, uint32(1), currentBlock.MusicIdx)
	must.EqOp(t, uint32(50), currentBlock.Volume)

	// Test pause/unpause functions exist
	m.Sub_43DBD0() // increment counter
	m.Sub_43DBE0() // decrement counter

	// Test state checking
	isPlaying := m.Sub_43DC40()
	must.EqOp(t, false, isPlaying) // Should be false initially

	// Test cleanup
	m.Sub_43D650()
	must.EqOp(t, uint32(0), m.currentPlaying.MusicIdx)
}

func TestMusicState(t *testing.T) {
	// Test MusicState struct size constraint
	state := MusicState{MusicIdx: 1, Volume: 100, Position: 500, D: 10}
	must.EqOp(t, uint32(1), state.MusicIdx)
	must.EqOp(t, uint32(100), state.Volume)
	must.EqOp(t, uint32(500), state.Position)
	must.EqOp(t, uint32(10), state.D)
}

func TestMusicBlock(t *testing.T) {
	// Test MusicBlock struct
	block := MusicBlock{
		Inner:  MusicState{MusicIdx: 2, Volume: 75, Position: 0, D: 0},
		Loaded: 1,
	}
	must.EqOp(t, uint32(2), block.Inner.MusicIdx)
	must.EqOp(t, uint32(75), block.Inner.Volume)
	must.EqOp(t, uint32(1), block.Loaded)
}

func TestVolumeClamp(t *testing.T) {
	m := newTestMusic(t)
	m.Init()

	// Test volume clamping to 100
	testMusic := MusicState{MusicIdx: 1, Volume: 150, Position: 0, D: 0}
	m.SetNextMusic(testMusic)

	currentBlock := m.GetCurrentBlock()
	must.EqOp(t, uint32(100), currentBlock.Volume)
}

func TestAudioTable(t *testing.T) {
	// Test that audio table has expected entries
	must.True(t, len(audioTable) > 0)
	must.EqOp(t, "", audioTable[0])
	must.EqOp(t, "chap1", audioTable[1])
	must.EqOp(t, "wander3", audioTable[23])
}
