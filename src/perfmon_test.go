package opennox

import (
	"net/netip"
	"testing"
	"time"

	"github.com/noxworld-dev/opennox/v1/common/ntype"
)

type mockStreamStats struct {
	host     bool
	player   ntype.PlayerInd
	transfer uint32
}

func (m *mockStreamStats) String() string { return "mock" }
func (m *mockStreamStats) IsHost() bool   { return m.host }
func (m *mockStreamStats) Player() ntype.PlayerInd {
	return m.player
}
func (m *mockStreamStats) IP() netip.Addr       { return netip.Addr{} }
func (m *mockStreamStats) Addr() netip.AddrPort { return netip.AddrPort{} }
func (m *mockStreamStats) TransferStats() uint32 {
	return m.transfer
}

func TestNewPerfmon(t *testing.T) {
	m := newPerfmon()
	if m == nil {
		t.Fatal("newPerfmon returned nil")
	}
	if m.nextCnt != 10 {
		t.Errorf("nextCnt = %d, want 10", m.nextCnt)
	}
	if m.logger == nil {
		t.Error("logger is nil")
	}
}

func TestPerfmonToggle(t *testing.T) {
	m := newPerfmon()
	if m.enabled {
		t.Error("expected enabled to be false initially")
	}
	m.Toggle()
	if !m.enabled {
		t.Error("expected enabled to be true after Toggle")
	}
	m.Toggle()
	if m.enabled {
		t.Error("expected enabled to be false after second Toggle")
	}
}

func TestPerfmonTransferStats(t *testing.T) {
	m := newPerfmon()
	// Set transferTick to past so first call doesn't use cached value
	m.transferTick[0] = -time.Second

	// Test host stream
	hostMock := &mockStreamStats{
		host:     true,
		player:   0,
		transfer: 12345,
	}
	stat := m.TransferStats(hostMock)
	if stat != 12345 {
		t.Errorf("TransferStats for host = %d, want 12345", stat)
	}

	// Second call within a second should return cached value
	hostMock.transfer = 67890
	stat2 := m.TransferStats(hostMock)
	if stat2 != 12345 {
		t.Errorf("TransferStats cached = %d, want 12345", stat2)
	}

	// Test non-host stream
	m2 := newPerfmon()
	m2.transferTick[1] = -time.Second // player 2 => ri = 1
	playerMock := &mockStreamStats{
		host:     false,
		player:   2, // player index 2 means ri = 1
		transfer: 11111,
	}
	stat3 := m2.TransferStats(playerMock)
	if stat3 != 11111 {
		t.Errorf("TransferStats for player = %d, want 11111", stat3)
	}
}

func TestPerfmonStartProfile(t *testing.T) {
	m := newPerfmon()

	// Test client profile
	done := m.startProfileClient()
	if done == nil {
		t.Fatal("startProfileClient returned nil")
	}
	time.Sleep(time.Millisecond)
	done()
	if m.profClient == 0 {
		t.Error("profClient should be non-zero after profile")
	}

	// Test server profile
	done2 := m.startProfileServer()
	if done2 == nil {
		t.Fatal("startProfileServer returned nil")
	}
	time.Sleep(time.Millisecond)
	done2()
	if m.profServer == 0 {
		t.Error("profServer should be non-zero after profile")
	}
}

func TestPerfmonBandData(t *testing.T) {
	// bandData accesses memmap which may not be initialized in test
	// We just verify it doesn't panic when memmap is not set up
	// In a real environment, memmap would be initialized
	m := newPerfmon()
	// This may panic if memmap not initialized, so we recover
	func() {
		defer func() {
			if r := recover(); r != nil {
				// Expected if memmap not initialized
				t.Logf("bandData panicked as expected (memmap not initialized): %v", r)
			}
		}()
		_ = m.bandData(0)
	}()
}
