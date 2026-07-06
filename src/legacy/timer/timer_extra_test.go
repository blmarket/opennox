package timer

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/noxtest"
	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/stretchr/testify/require"
)

func TestTimerGroup_Update(t *testing.T) {
	PlatformTicks = func() uint64 { return uint64(platform.Ticks() / time.Millisecond) }
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	tg := &TimerGroup{}
	tg.Init()

	// Test Update
	tg.Update()

	// Test IsUpdated
	updated := tg.IsUpdated()
	_ = updated

	// Test ClearUpdated
	tg.ClearUpdated()
	require.False(t, tg.Timers[0].Flags&2 != 0)
	require.False(t, tg.Timers[1].Flags&2 != 0)
	require.False(t, tg.Timers[2].Flags&2 != 0)

	// Test TimerGroupClearUpdated
	TimerGroupClearUpdated(tg)
}

func TestTimerGroup_Mix(t *testing.T) {
	PlatformTicks = func() uint64 { return uint64(platform.Ticks() / time.Millisecond) }
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	tg1 := &TimerGroup{}
	tg1.Init()

	tg2 := &TimerGroup{}
	tg2.Init()

	// Test Mix
	tg1.Mix(tg2)

	// Test Mix with Timers[2] at 0x2000
	tg2.Timers[2].Current = 0x2000 << 16
	tg1.Mix(tg2)
}

func TestTimer_Update_Empty(t *testing.T) {
	PlatformTicks = func() uint64 { return uint64(platform.Ticks() / time.Millisecond) }
	p := noxtest.MockPlatform{}
	platform.Set(&p)

	timer := &Timer{}
	timer.Init(0)

	// Update when diff == 0 should return false
	result := timer.Update()
	require.False(t, result)
}

func TestTimer_C(t *testing.T) {
	timer := &Timer{}
	ptr := timer.C()
	require.NotNil(t, ptr)
}

func TestTimerGroup_C(t *testing.T) {
	tg := &TimerGroup{}
	ptr := tg.C()
	require.NotNil(t, ptr)
}
