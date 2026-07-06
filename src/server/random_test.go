package server

import (
	"testing"
	"time"

	"github.com/noxworld-dev/opennox-lib/platform"
	"github.com/stretchr/testify/require"
)

type testPlatform struct {
	seed int64
}

func (p *testPlatform) TimeSeed() int64        { return p.seed }
func (p *testPlatform) Ticks() time.Duration   { return 0 }
func (p *testPlatform) Sleep(dt time.Duration) {}
func (p *testPlatform) RandInt() int           { return 0 }
func (p *testPlatform) RandSeed(seed int64)    {}
func (p *testPlatform) RandSeedTime()          {}

func TestServerRandomInit(t *testing.T) {
	var sr serverRandom
	sr.init(nil)
	require.NotNil(t, sr.Logic)
	require.NotNil(t, sr.Other)
}

func TestServerRandomReset(t *testing.T) {
	var sr serverRandom
	sr.init(&testPlatform{seed: 12345})
	sr.Reset()
	require.NotNil(t, sr.Logic)
	require.NotNil(t, sr.Other)

	// Reset with nil platform should use platform.TimeSeed()
	sr2 := serverRandom{}
	sr2.init(nil)
	sr2.Reset()
	require.NotNil(t, sr2.Logic)
}

func TestServerRandomColor3(t *testing.T) {
	var sr serverRandom
	sr.init(nil)
	sr.Reset()
	c := sr.RandomColor3()
	// Just verify it doesn't panic and returns values in range
	require.LessOrEqual(t, c.R, byte(255))
	require.LessOrEqual(t, c.G, byte(255))
	require.LessOrEqual(t, c.B, byte(255))
}

var _ platform.Platform = (*testPlatform)(nil)
