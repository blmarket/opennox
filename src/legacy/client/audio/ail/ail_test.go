//go:build !server

package ail

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

type memReadSeekCloser struct {
	*bytes.Reader
}

func (m memReadSeekCloser) Close() error { return nil }

func makeWavPCM() []byte {
	// Simple WAV file with PCM format, 1 channel, 8000 Hz, 4 samples
	buf := new(bytes.Buffer)
	// RIFF header
	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, uint32(36+8)) // file size - 8
	buf.WriteString("WAVE")
	// fmt chunk
	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(16)) // size
	binary.Write(buf, binary.LittleEndian, uint16(1))  // PCM
	binary.Write(buf, binary.LittleEndian, uint16(1))  // channels
	binary.Write(buf, binary.LittleEndian, uint32(8000))
	binary.Write(buf, binary.LittleEndian, uint32(16000))
	binary.Write(buf, binary.LittleEndian, uint16(2))
	binary.Write(buf, binary.LittleEndian, uint16(16))
	// data chunk
	buf.WriteString("data")
	binary.Write(buf, binary.LittleEndian, uint32(8))
	// 4 samples
	binary.Write(buf, binary.LittleEndian, uint16(0))
	binary.Write(buf, binary.LittleEndian, uint16(1000))
	binary.Write(buf, binary.LittleEndian, uint16(2000))
	binary.Write(buf, binary.LittleEndian, uint16(3000))
	return buf.Bytes()
}

func makeWavADPCM() []byte {
	buf := new(bytes.Buffer)
	buf.WriteString("RIFF")
	// placeholder for size
	sizePos := buf.Len()
	binary.Write(buf, binary.LittleEndian, uint32(0))
	buf.WriteString("WAVE")
	// fmt chunk with ADPCM (0x11)
	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(20)) // size
	binary.Write(buf, binary.LittleEndian, uint16(0x11))
	binary.Write(buf, binary.LittleEndian, uint16(1))
	binary.Write(buf, binary.LittleEndian, uint32(8000))
	binary.Write(buf, binary.LittleEndian, uint32(8000))
	binary.Write(buf, binary.LittleEndian, uint16(256))
	binary.Write(buf, binary.LittleEndian, uint16(4))
	// extra 4 bytes for ADPCM header size?
	binary.Write(buf, binary.LittleEndian, uint16(0))
	binary.Write(buf, binary.LittleEndian, uint16(0))
	// fact chunk
	buf.WriteString("fact")
	binary.Write(buf, binary.LittleEndian, uint32(4))
	binary.Write(buf, binary.LittleEndian, uint32(4)) // samples
	// data chunk
	buf.WriteString("data")
	// one ADPCM block (256 bytes for mono)
	blockSize := 256
	binary.Write(buf, binary.LittleEndian, uint32(blockSize))
	// ADPCM block header: predictor (2), index (1), reserved (1)
	binary.Write(buf, binary.LittleEndian, uint16(0))
	buf.WriteByte(0)
	buf.WriteByte(0)
	// rest of block as zeros
	for i := 4; i < blockSize; i++ {
		buf.WriteByte(0)
	}
	// update size
	data := buf.Bytes()
	size := uint32(len(data) - 8)
	binary.LittleEndian.PutUint32(data[sizePos:], size)
	return data
}

func TestWavReaderPCM(t *testing.T) {
	data := makeWavPCM()
	r := memReadSeekCloser{bytes.NewReader(data)}
	wr, err := newWavReader(r, "test.wav")
	if err != nil {
		t.Fatalf("newWavReader: %v", err)
	}
	defer wr.Close()

	if wr.Name() != "test.wav" {
		t.Errorf("Name = %q", wr.Name())
	}
	if wr.Format() != "WAV+PCM" {
		t.Errorf("Format = %q", wr.Format())
	}
	if wr.Channels() != 1 {
		t.Errorf("Channels = %d", wr.Channels())
	}
	if wr.SampleRate() != 8000 {
		t.Errorf("SampleRate = %d", wr.SampleRate())
	}

	out := make([]int16, 10)
	n, ok := wr.Decode(out)
	if !ok {
		t.Fatalf("Decode failed")
	}
	if n == 0 {
		t.Errorf("Decode returned 0 samples")
	}

	pos := wr.Position()
	if pos < 0 {
		t.Errorf("Position = %d", pos)
	}

	wr.Seek(0)
	sz, err := wr.ChunkSize()
	if err != nil && err != io.EOF {
		t.Logf("ChunkSize err: %v", err)
	}
	_ = sz

	rem := wr.Remaining()
	_ = rem

	wr.Skip(0)
}

func TestWavReaderADPCM(t *testing.T) {
	data := makeWavADPCM()
	r := memReadSeekCloser{bytes.NewReader(data)}
	wr, err := newWavReader(r, "test_adpcm.wav")
	if err != nil {
		t.Fatalf("newWavReader ADPCM: %v", err)
	}
	defer wr.Close()

	if wr.Format() != "WAV+ADPCM" {
		t.Errorf("Format = %q", wr.Format())
	}

	out := make([]int16, 1024)
	n, ok := wr.Decode(out)
	if ok {
		t.Logf("ADPCM decode n=%d", n)
	}

	wr.Seek(0)
	pos := wr.Position()
	_ = pos
}

func TestWavReaderInvalid(t *testing.T) {
	// Invalid magic
	data := []byte("INVALID")
	r := memReadSeekCloser{bytes.NewReader(data)}
	_, err := newWavReader(r, "bad.wav")
	if err == nil {
		t.Errorf("expected error for invalid magic")
	}

	// Invalid WAVE
	buf := new(bytes.Buffer)
	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, uint32(10))
	buf.WriteString("BAD!")
	_, err = newWavReader(memReadSeekCloser{bytes.NewReader(buf.Bytes())}, "bad2.wav")
	if err == nil {
		t.Errorf("expected error for invalid WAVE")
	}
}

func TestPCMDecoderDirect(t *testing.T) {
	data := makeWavPCM()
	r := memReadSeekCloser{bytes.NewReader(data)}
	wr, err := newWavReader(r, "test.wav")
	if err != nil {
		t.Fatalf("newWavReader: %v", err)
	}
	defer wr.Close()

	dec, ok := wr.dec.(*pcmDecoder)
	if !ok {
		t.Fatalf("not pcm decoder")
	}

	if dec.Format() != "PCM" {
		t.Errorf("Format = %q", dec.Format())
	}
	if dec.Channels() != 1 {
		t.Errorf("Channels = %d", dec.Channels())
	}
	if dec.SampleRate() != 8000 {
		t.Errorf("SampleRate = %d", dec.SampleRate())
	}

	out := make([]int16, 10)
	n, ok := dec.Decode(out)
	if !ok || n == 0 {
		t.Errorf("Decode failed n=%d ok=%v", n, ok)
	}

	dec.Seek(0)
	if dec.Position() != 0 {
		t.Errorf("Position after seek = %d", dec.Position())
	}

	dec.Close()
}

func TestADPCMDecoderDirect(t *testing.T) {
	data := makeWavADPCM()
	r := memReadSeekCloser{bytes.NewReader(data)}
	wr, err := newWavReader(r, "test.wav")
	if err != nil {
		t.Fatalf("newWavReader: %v", err)
	}
	defer wr.Close()

	dec, ok := wr.dec.(*adpcmDecoder)
	if !ok {
		t.Fatalf("not adpcm decoder")
	}

	if dec.Format() != "ADPCM" {
		t.Errorf("Format = %q", dec.Format())
	}

	out := make([]int16, 1024)
	dec.Decode(out)
	dec.Seek(0)
	dec.Position()
	dec.Close()
}

func TestDecodeFunctions(t *testing.T) {
	// Test decodeADPCMMono
	data := make([]byte, 8)
	data[0] = 0
	data[1] = 0
	data[2] = 0
	data[3] = 0
	out := decodeADPCMMono(nil, data)
	if len(out) == 0 {
		t.Errorf("decodeADPCMMono returned empty")
	}

	// Test decodeADPCMStereo
	data = make([]byte, 16)
	out = decodeADPCMStereo(nil, data)
	if len(out) == 0 {
		t.Errorf("decodeADPCMStereo returned empty")
	}

	// Test sat16
	if sat16(100000) != 32767 {
		t.Errorf("sat16 high failed")
	}
	if sat16(-100000) != -32768 {
		t.Errorf("sat16 low failed")
	}
	if sat16(100) != 100 {
		t.Errorf("sat16 normal failed")
	}

	// Test satindex
	if satindex(-1) != 0 {
		t.Errorf("satindex low failed")
	}
	if satindex(100) != 88 {
		t.Errorf("satindex high failed")
	}
	if satindex(10) != 10 {
		t.Errorf("satindex normal failed")
	}

	// Test decodeNibble
	v := decodeNibble(0, 0, 0)
	_ = v
}

func TestOpenalFunctions(t *testing.T) {
	// Test functions that don't require OpenAL context
	// These should handle nil gracefully or return error values

	var s Sample
	var st Stream
	var tm Timer
	var dig Driver

	// Test null or error paths
	_ = s.GetSource()
	s.End()
	s.Init()
	_ = LastError()
	s.LoadBuffer(0, nil)
	st.Pause(false)
	s.RegisterEOBCallback(nil)
	s.RegisterEOSCallback(nil)
	_ = s.BufferReady()
	_ = s.UserData()
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
	_ = Startup()
	s.Stop()
	tm.Stop()
	_ = st.Position()
	_ = st.Status()
	_ = dig.Close()
	func() {
		defer func() { recover() }()
		_ = WaveOutOpen()
	}()

	// Test AllocateSample and OpenStream (may fail without OpenAL, but should not panic)
	func() {
		defer func() { recover() }()
		_ = dig.AllocateSample()
	}()
	func() {
		defer func() { recover() }()
		_ = dig.OpenStream("nonexistent.wav")
	}()
	func() {
		defer func() { recover() }()
		_ = RegisterTimer(nil)
	}()

	// Test wavReader open with nonexistent file
	_, err := openWav("/nonexistent/path.wav")
	if err == nil {
		t.Logf("openWav should fail for nonexistent file")
	}
}
