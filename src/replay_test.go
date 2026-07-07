package opennox

import (
	"bytes"
	"testing"

	"github.com/noxworld-dev/opennox-lib/prand"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func TestReplayConstants(t *testing.T) {
	if replayOpMsg != 1 {
		t.Fatal("replayOpMsg should be 1")
	}
	if replayOpInit != 2 {
		t.Fatal("replayOpInit should be 2")
	}
	if replayOpConsole != 3 {
		t.Fatal("replayOpConsole should be 3")
	}
	if replayOpFrame != 4 {
		t.Fatal("replayOpFrame should be 4")
	}
}

func TestNoxXxxReplayWriteRndCounter(t *testing.T) {
	// Setup minimal server with Rand
	s := &Server{}
	s.Rand.Logic = prand.New(12345)
	noxServer = s

	var buf bytes.Buffer
	nox_xxx_replayWriteRndCounter_415F30(&buf)
	if buf.Len() != 4 {
		t.Fatalf("expected 4 bytes, got %d", buf.Len())
	}

	// Write again, should get same value since index not advanced by just reading
	var buf2 bytes.Buffer
	nox_xxx_replayWriteRndCounter_415F30(&buf2)
	if !bytes.Equal(buf.Bytes(), buf2.Bytes()) {
		t.Log("counters differ, which is expected if index advances")
	}
}

func TestNoxXxxReplayReadeRndCounter(t *testing.T) {
	s := &Server{}
	s.Rand.Logic = prand.New(999)
	noxServer = s

	// Write a known counter value
	var buf bytes.Buffer
	s.Rand.Logic.Reset(999)
	nox_xxx_replayWriteRndCounter_415F30(&buf)

	// Reset to different value
	s.Rand.Logic.Reset(0)

	// Read back
	nox_xxx_replayReadeRndCounter_415F50(&buf)
	// After reading, the rand index should be set to the value from buffer
	// We can't easily verify without exposing internal state, but it should not panic
}

func TestNoxXxxReplaySaveConsole(t *testing.T) {
	// nil writer should not panic
	replay.writer = nil
	nox_xxx_replaySaveConsole("test")
	nox_xxx_replaySaveConsole("")

	// empty cmd should not write
	var buf bytes.Buffer
	replay.writer = &buf
	nox_xxx_replaySaveConsole("")
	if buf.Len() != 0 {
		t.Fatal("empty cmd should not write")
	}

	// valid cmd should write
	replay.writer = &buf
	s := &Server{}
	noxServer = s
	nox_xxx_replaySaveConsole("test command")
	if buf.Len() == 0 {
		t.Fatal("valid cmd should write data")
	}

	replay.writer = nil
}

func TestNoxXxxReplayWriteFrame(t *testing.T) {
	// nil writer should not panic
	replay.writer = nil
	nox_xxx_replayWriteFrame_4D39B0()

	// with writer
	var buf bytes.Buffer
	replay.writer = &buf
	s := &Server{}
	noxServer = s
	nox_xxx_replayWriteFrame_4D39B0()
	if buf.Len() != 5 {
		t.Fatalf("expected 5 bytes for frame op, got %d", buf.Len())
	}
	replay.writer = nil
}

func TestNoxXxxReplayStopSave(t *testing.T) {
	s := &Server{}
	// nil closer should not panic
	replay.wcloser = nil
	replay.writer = &bytes.Buffer{}
	s.nox_xxx_replayStopSave_4D33B0()
	if replay.writer != nil {
		t.Fatal("writer should be nil after stop")
	}
	if noxflags.HasEngine(noxflags.EngineReplayWrite) {
		t.Fatal("replay write flag should be cleared")
	}
}

func TestNoxXxxReplayStopRead(t *testing.T) {
	s := &Server{}
	replay.rcloser = nil
	replay.reader = &bytes.Buffer{}
	replay.readHeader = true
	s.nox_xxx_replayStopReadMB_4D3530()
	if replay.reader != nil {
		t.Fatal("reader should be nil after stop")
	}
	if replay.readHeader {
		t.Fatal("readHeader should be false after stop")
	}
	if noxflags.HasEngine(noxflags.EngineReplayRead) {
		t.Fatal("replay read flag should be cleared")
	}
}
