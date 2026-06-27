package music

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/client/audio/ail"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

func TestMusicModule(t *testing.T) {
	os.Setenv("NOX_E2E", "true")
	defer os.Unsetenv("NOX_E2E")

	handles.Init()
	t.Cleanup(handles.Release)

	timer.PlatformTicks = func() uint64 { return 100 }
	t.Cleanup(func() {
		timer.PlatformTicks = nil
	})

	var (
		dword_5d4594_816368      uint32
		dword_5d4594_816372      uint32
		dword_5d4594_816376      ail.Driver
		dword_587000_93156       uint32
		dword_587000_93160       uint32
		dword_5d4594_816340      uint32
		counter_5d4594_816244    timer.TimerGroup
		counter_587000_81128     timer.TimerGroup
		ptr_counter_587000_81128 = &counter_587000_81128
		dword_5d4594_816348      uint32
	)

	// Create a dummy WAV file in the temp dir
	tmpDir := t.TempDir()
	dummyWav := filepath.Join(tmpDir, "title.wav")
	if err := os.WriteFile(dummyWav, []byte("RIFFxxxxWAVEfmt xxxxdataxxxx"), 0644); err != nil {
		t.Fatalf("failed to create dummy wav: %v", err)
	}

	m := NewModule(
		tmpDir,
		&dword_5d4594_816368,
		&dword_5d4594_816372,
		&dword_5d4594_816376,
		&dword_587000_93156,
		&dword_587000_93160,
		&dword_5d4594_816340,
		&counter_5d4594_816244,
		&ptr_counter_587000_81128,
		&dword_5d4594_816348,
		func() ail.Driver {
			return ail.Driver(handles.New())
		},
		func() bool {
			return false
		},
		func() string {
			return tmpDir
		},
		func() uint64 {
			return 100
		},
	)

	// Test Init
	m.Init()
	if dword_5d4594_816340 != 1 {
		t.Errorf("expected module to be initialized, got %d", dword_5d4594_816340)
	}

	// Test SetNextMusic
	m.SetNextMusic(MusicState{MusicIdx: 13, Volume: 50}) // 13 is "title"
	block := m.GetCurrentBlock()
	if block.MusicIdx != 13 {
		t.Errorf("expected MusicIdx to be 13, got %d", block.MusicIdx)
	}

	// Test Update
	m.Update()
	m.Sub_43D2D0()
}
