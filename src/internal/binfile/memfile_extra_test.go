package binfile

import (
	"encoding/binary"
	"io"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

func TestMemFile_NewAndFree(t *testing.T) {
	data, _ := alloc.Make([]byte{}, 64)
	mf := NewMemFile(unsafe.Pointer(&data[0]), len(data))
	if mf == nil {
		t.Fatal("NewMemFile returned nil")
	}
	if mf.C() == nil {
		t.Error("C() should not be nil")
	}
	raw := mf.RawData()
	if len(raw) != 64 {
		t.Errorf("RawData len = %d, want 64", len(raw))
	}
	mf.Free()
	// After Free, the struct is freed, so we can't safely access it
}

func TestMemFile_SeekReadWrite(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	buf, _ := alloc.Make([]byte{}, len(data))
	copy(buf, data)
	mf := NewMemFile(unsafe.Pointer(&buf[0]), len(buf))
	defer mf.Free()

	// Test Data()
	d := mf.Data()
	if len(d) != len(data) {
		t.Fatalf("Data len = %d, want %d", len(d), len(data))
	}

	// Test ReadU8, ReadI8
	if v := mf.ReadU8(); v != 0 {
		t.Errorf("ReadU8 = %d, want 0", v)
	}
	if v := mf.ReadI8(); v != 1 {
		t.Errorf("ReadI8 = %d, want 1", v)
	}

	// Test ReadU16, ReadI16
	mf.Seek(2, io.SeekStart)
	if v := mf.ReadU16(); v != binary.LittleEndian.Uint16([]byte{2, 3}) {
		t.Errorf("ReadU16 mismatch")
	}
	mf.Seek(4, io.SeekStart)
	_ = mf.ReadI16()

	// Test ReadU32, ReadI32
	mf.Seek(0, io.SeekStart)
	v32 := mf.ReadU32()
	expected := binary.LittleEndian.Uint32([]byte{0, 1, 2, 3})
	if v32 != expected {
		t.Errorf("ReadU32 = %d, want %d", v32, expected)
	}
	mf.Seek(4, io.SeekStart)
	_ = mf.ReadI32()

	// Test ReadU64, ReadI64
	mf.Seek(0, io.SeekStart)
	v64 := mf.ReadU64()
	if v64 == 0 {
		t.Error("ReadU64 should not be zero")
	}
	mf.Seek(0, io.SeekStart)
	_ = mf.ReadI64()

	// Test Seek
	off, err := mf.Seek(5, io.SeekStart)
	if err != nil || off != 5 {
		t.Errorf("Seek start failed: off=%d err=%v", off, err)
	}
	off, _ = mf.Seek(2, io.SeekCurrent)
	if off != 7 {
		t.Errorf("Seek current failed: off=%d", off)
	}
	off, _ = mf.Seek(-2, io.SeekEnd)
	if off != int64(len(data)-2) {
		t.Errorf("Seek end failed: off=%d", off)
	}
	off, _ = mf.Seek(-100, io.SeekStart)
	if off != 0 {
		t.Errorf("Seek negative should clamp to 0, got %d", off)
	}
	off, _ = mf.Seek(1000, io.SeekStart)
	if off != int64(len(data)) {
		t.Errorf("Seek beyond end should clamp to size, got %d", off)
	}

	// Test Skip
	mf.Seek(0, io.SeekStart)
	mf.Skip(3)
	if mf.offset() != 3 {
		t.Errorf("Skip failed, offset=%d", mf.offset())
	}
	mf.Skip(1000)
	if mf.offset() != len(data) {
		t.Error("Skip beyond end should clamp")
	}

	// Test Read
	mf.Seek(0, io.SeekStart)
	buf2 := make([]byte, 4)
	n, err := mf.Read(buf2)
	if err != nil || n != 4 {
		t.Errorf("Read failed: n=%d err=%v", n, err)
	}
	mf.Seek(0, io.SeekStart)
	mf.Skip(len(data))
	n, err = mf.Read(buf2)
	if err != io.EOF {
		t.Errorf("Read at EOF should return io.EOF, got %v", err)
	}

	// Test SkipString8, ReadBytes8, ReadString8, ReadString16
	mf2Data := []byte{
		3, 'f', 'o', 'o',
		5, 'h', 'e', 'l', 'l', 'o',
		2, 'h', 'i',
	}
	buf3, _ := alloc.Make([]byte{}, len(mf2Data))
	copy(buf3, mf2Data)
	mf2 := NewMemFile(unsafe.Pointer(&buf3[0]), len(buf3))
	defer mf2.Free()

	mf2.SkipString8()
	if mf2.offset() != 4 {
		t.Errorf("SkipString8 offset = %d, want 4", mf2.offset())
	}
	b, err := mf2.ReadBytes8()
	if err != nil || string(b) != "hello" {
		t.Errorf("ReadBytes8 failed: %v %s", err, string(b))
	}
	s, err := mf2.ReadString8()
	if err != nil || s != "hi" {
		t.Errorf("ReadString8 failed: %v %s", err, s)
	}

	// Test ReadString16
	mf3Data := []byte{3, 0, 'b', 'a', 'r'}
	buf4, _ := alloc.Make([]byte{}, len(mf3Data))
	copy(buf4, mf3Data)
	mf3 := NewMemFile(unsafe.Pointer(&buf4[0]), len(buf4))
	defer mf3.Free()
	s, err = mf3.ReadString16()
	if err != nil || s != "bar" {
		t.Errorf("ReadString16 failed: %v %s", err, s)
	}

	// Test ReadU64Align
	alignData := make([]byte, 24)
	buf5, _ := alloc.Make([]byte{}, len(alignData))
	copy(buf5, alignData)
	binary.LittleEndian.PutUint64(buf5[8:], 0x12345678)
	mf4 := NewMemFile(unsafe.Pointer(&buf5[0]), len(buf5))
	defer mf4.Free()
	mf4.Seek(1, io.SeekStart)
	v := mf4.ReadU64Align()
	if v != 0x12345678 {
		t.Errorf("ReadU64Align = 0x%x, want 0x12345678", v)
	}
}

func TestMemFile_DataNil(t *testing.T) {
	mf := &MemFile{}
	if mf.RawData() != nil {
		t.Error("RawData should be nil for empty MemFile")
	}
	if mf.Data() != nil {
		t.Error("Data should be nil for empty MemFile")
	}
	mf.Skip(10)
}

func TestFileSize(t *testing.T) {
	mfData := []byte{1, 2, 3, 4, 5}
	buf, _ := alloc.Make([]byte{}, len(mfData))
	copy(buf, mfData)
	mf := NewMemFile(unsafe.Pointer(&buf[0]), len(buf))
	defer mf.Free()

	sz, err := FileSize(mf)
	if err != nil || sz != int64(len(mfData)) {
		t.Errorf("FileSize = %d, err %v", sz, err)
	}
}

func TestNewFile(t *testing.T) {
	f := NewFile(nil)
	if f == nil {
		t.Error("NewFile should not return nil")
	}
	ft := NewTextFile(nil)
	if ft == nil || !ft.text {
		t.Error("NewTextFile should set text=true")
	}
}
