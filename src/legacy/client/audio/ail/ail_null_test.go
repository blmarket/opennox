//go:build server

package ail

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func TestNullImplementations(t *testing.T) {
	handles.Init()
	var dig Driver

	s := dig.AllocateSample()
	if s == 0 {
		t.Errorf("AllocateSample returned 0")
	}
	s.Release()

	st := dig.OpenStream("test")
	if st == 0 {
		t.Errorf("OpenStream returned 0")
	}
	if err := st.Close(); err != nil {
		t.Errorf("Stream Close: %v", err)
	}

	tm := RegisterTimer(func(u uint32) {})
	if tm == 0 {
		t.Errorf("RegisterTimer returned 0")
	}
	tm.Release()

	// Test Sample methods
	s = Sample(handles.New())
	if s.GetSource() != nil {
		t.Errorf("GetSource should be nil")
	}
	s.End()
	s.Init()
	if LastError() != "" {
		t.Errorf("LastError should be empty")
	}
	s.LoadBuffer(0, nil)
	st.Pause(false)
	s.RegisterEOBCallback(nil)
	s.RegisterEOSCallback(nil)
	if s.BufferReady() != -1 {
		t.Errorf("BufferReady should be -1")
	}
	if s.UserData() != nil {
		t.Errorf("UserData should be nil")
	}
	Serve()
	s.SetADPCMBlockSize(0)
	s.SetPan(0)
	s.SetPlaybackRate(0)
	s.SetType(0, 0)
	s.SetUserData(nil)
	s.SetVolume(0)
	st.SetPosition(0)
	st.SetVolume(0)
	tm.SetFrequency(0)
	Shutdown()
	st.Start()
	tm.Start()
	if Startup() != -1 {
		t.Errorf("Startup should be -1")
	}
	s.Stop()
	tm.Stop()
	if st.Position() != -1 {
		t.Errorf("Position should be -1")
	}
	if st.Status() != 2 {
		t.Errorf("Status should be 2")
	}
	if err := dig.Close(); err != nil {
		t.Errorf("Driver Close: %v", err)
	}
	dig = WaveOutOpen()
	if dig != 0 {
		t.Errorf("WaveOutOpen should return 0")
	}

	s.Release()
}
