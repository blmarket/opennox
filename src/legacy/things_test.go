package legacy

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

var thingParserSentinel = []byte{0xfe, 0xed, 0xfa, 0xce}

func newThingParserMemFile(t *testing.T, payload []byte) *binfile.MemFile {
	t.Helper()
	data := append(append([]byte(nil), payload...), thingParserSentinel...)
	raw, _ := alloc.CloneSlice(data)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	t.Cleanup(f.Free)
	return f
}

func checkThingParserCursor(t *testing.T, f *binfile.MemFile, want int) {
	t.Helper()
	got, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		t.Fatalf("current offset: %v", err)
	}
	if got != int64(want) {
		t.Fatalf("current offset = %d, want %d", got, want)
	}
	if remaining := f.Data(); !bytes.Equal(remaining, thingParserSentinel) {
		t.Fatalf("remaining data = % x, want sentinel % x", remaining, thingParserSentinel)
	}
}

func appendThingU32(dst []byte, value uint32) []byte {
	var buf [4]byte
	binary.LittleEndian.PutUint32(buf[:], value)
	return append(dst, buf[:]...)
}

func resetThingDefinitions(t *testing.T) {
	t.Helper()
	Sub_485CF0()
	Sub_485F30()
	LoadAllBinFileSectionsResetCounters()
	t.Cleanup(func() {
		Sub_485CF0()
		Sub_485F30()
		LoadAllBinFileSectionsResetCounters()
	})
}

func TestNoxThingSkipAUD(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "zero records",
			payload: []byte{0, 0, 0, 0},
		},
		{
			name:    "negative record count",
			payload: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			name: "multiple records and strings",
			payload: []byte{
				2, 0, 0, 0, // record count
				3, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1a, 0x1b,
				2, 'a', 'b', 1, 'c', 0,
				0, 0x20, 0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
				0,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newThingParserMemFile(t, tc.payload)
			Nox_thing_skip_AUD_414D40(f)
			checkThingParserCursor(t, f, len(tc.payload))
		})
	}
}

func TestNoxThingSkipAVNT(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "empty name and terminator",
			payload: []byte{0, 0},
		},
		{
			name: "all supported opcode layouts",
			payload: []byte{
				3, 'a', 'u', 'd',
				1, 0xa1,
				2, 0xa2,
				3, 0xa3,
				4, 0xa4,
				5, 0xa5,
				6, 0xb1, 0xb2,
				9, 0xb3, 0xb4,
				10, 0xb5, 0xb6,
				7, 2, 'x', 'y', 1, 'z', 0,
				8, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8,
				0,
			},
		},
		{
			name:    "unknown opcode terminates",
			payload: []byte{1, 'x', 0x7f},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newThingParserMemFile(t, tc.payload)
			Nox_thing_skip_AVNT_452B00(f)
			checkThingParserCursor(t, f, len(tc.payload))
		})
	}
}

func TestNoxThingReadImageCursor(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "zero records",
			payload: []byte{0, 0, 0, 0},
		},
		{
			name:    "negative record count",
			payload: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			name: "single direct image",
			payload: []byte{
				1, 0, 0, 0,
				2, 'a', 'b', 1,
				0x12, 0x34, 0x56, 0x78,
			},
		},
		{
			name: "type two with no frames",
			payload: []byte{
				1, 0, 0, 0,
				0, 2, 0, 0xaa, 2, 'm', 'n',
			},
		},
		{
			name: "type two direct and named frames",
			payload: []byte{
				1, 0, 0, 0,
				0, 2, 2, 0xaa, 1, 'm',
				0x12, 0x34, 0x56, 0x78,
				0xff, 0xff, 0xff, 0xff, 7, 2, 'x', 'y',
			},
		},
		{
			name: "unknown type uses one image",
			payload: []byte{
				1, 0, 0, 0,
				0, 3,
				0x12, 0x34, 0x56, 0x78,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newThingParserMemFile(t, tc.payload)
			Nox_thing_read_image_415240(f)
			checkThingParserCursor(t, f, len(tc.payload))
		})
	}
}

func TestNoxThingReadAbilityCursor(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "zero records",
			payload: []byte{0, 0, 0, 0},
		},
		{
			name:    "negative record count",
			payload: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			name: "direct and named references",
			payload: []byte{
				1, 0, 0, 0,
				2, 'a', 'b', 0xee,
				0x12, 0x34, 0x56, 0x78,
				0xff, 0xff, 0xff, 0xff, 7, 2, 'x', 'y',
				0x21, 0x43, 0x65, 0x87,
				2, 'c', 'd',
				3, 0, 'e', 'f', 'g',
				1, 'h', 0, 2, 'i', 'j',
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newThingParserMemFile(t, tc.payload)
			Nox_thing_read_ability_415320(f)
			checkThingParserCursor(t, f, len(tc.payload))
		})
	}
}

func TestNoxThingSkipSpellsCursor(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{
			name:    "zero records",
			payload: []byte{0, 0, 0, 0},
		},
		{
			name:    "negative record count",
			payload: []byte{0xff, 0xff, 0xff, 0xff},
		},
		{
			name: "direct and named references",
			payload: []byte{
				1, 0, 0, 0,
				2, 'a', 'b', 0xe1, 0xe2, 0xe3,
				1, 'c',
				0xff, 0xff, 0xff, 0xff, 7, 2, 'x', 'y',
				0x12, 0x34, 0x56, 0x78,
				0xd1, 0xd2, 0xd3, 0xd4, 2, 'd', 'e',
				3, 0, 'f', 'g', 'h',
				1, 'i', 0, 2, 'j', 'k',
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newThingParserMemFile(t, tc.payload)
			Nox_thing_skip_spells_415100(f)
			checkThingParserCursor(t, f, len(tc.payload))
		})
	}
}

func TestNoxThingReadAudioCursor(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{name: "zero records", payload: []byte{0, 0, 0, 0}},
		{name: "negative record count", payload: []byte{0xff, 0xff, 0xff, 0xff}},
		{
			name: "unknown sound is skipped",
			payload: []byte{
				1, 0, 0, 0,
				0,                         // empty sound name
				0, 0, 0, 0, 0, 0, 0, 0, 0, // fixed fields
				0, // no sample names
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			f := newThingParserMemFile(t, tc.payload)
			buf := make([]byte, 256*1024)
			if got := Nox_thing_read_audio_415660(f, buf); got != 1 {
				t.Fatalf("result = %d, want 1", got)
			}
			checkThingParserCursor(t, f, len(tc.payload))
		})
	}
}

func TestNoxThingReadAVNTUnknownSoundCursor(t *testing.T) {
	payload := []byte{
		0, // empty sound name
		1, 0xa1,
		6, 0xb1, 0xb2,
		7, 2, 'x', 'y', 0,
		8, 0xc1, 0xc2, 0xc3, 0xc4, 0xc5, 0xc6, 0xc7, 0xc8,
		9, 0xd1, 0xd2,
		10, 0xe1, 0xe2,
		0,
	}
	f := newThingParserMemFile(t, payload)
	buf := make([]byte, 256*1024)
	if got := Nox_thing_read_AVNT_452890(f, buf); got != 1 {
		t.Fatalf("result = %d, want 1", got)
	}
	checkThingParserCursor(t, f, len(payload))
}

func TestNoxThingReadLegacyFloorCursor(t *testing.T) {
	tests := []struct {
		name       string
		images     []byte
		terminator uint32
		want       int
	}{
		{name: "empty image grid", terminator: 0x454e4420, want: 1},
		{
			name: "direct and named images",
			images: []byte{
				0x12, 0x34, 0x56, 0x78,
				0xff, 0xff, 0xff, 0xff, 0xaa, 2, 'x', 'y',
			},
			terminator: 0x454e4420,
			want:       1,
		},
		{name: "invalid terminator", terminator: 0, want: 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			payload := make([]byte, 17)
			if len(tc.images) != 0 {
				payload = append(payload, 1, 1, 2, 0)
			} else {
				payload = append(payload, 0, 0, 0, 0)
			}
			payload = append(payload, tc.images...)
			payload = appendThingU32(payload, tc.terminator)
			f := newThingParserMemFile(t, payload)
			if got := Nox_thing_read_FLOR_414DB0(f); got != tc.want {
				t.Fatalf("result = %d, want %d", got, tc.want)
			}
			checkThingParserCursor(t, f, len(payload))
		})
	}
}

func TestNoxThingReadLegacyEdgeCursor(t *testing.T) {
	tests := []struct {
		name       string
		images     []byte
		terminator uint32
		want       int
	}{
		{name: "empty image grid", terminator: 0x454e4420, want: 1},
		{
			name: "direct and named images",
			images: []byte{
				0x12, 0x34, 0x56, 0x78,
				0xff, 0xff, 0xff, 0xff, 0xaa, 2, 'x', 'y',
			},
			terminator: 0x454e4420,
			want:       1,
		},
		{name: "unsupported layout", terminator: 0x454e4420, want: 0},
		{name: "invalid terminator", terminator: 0, want: 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			payload := make([]byte, 14)
			if tc.name == "unsupported layout" {
				payload = append(payload, 0, 0, 1)
			} else if len(tc.images) != 0 {
				payload = append(payload, 1, 0, 0, 1, 0)
			} else {
				payload = append(payload, 0, 0, 0, 0, 0)
			}
			if tc.name != "unsupported layout" {
				payload = append(payload, tc.images...)
				payload = appendThingU32(payload, tc.terminator)
			}
			f := newThingParserMemFile(t, payload)
			buf := make([]byte, 256*1024)
			if got := Nox_thing_read_EDGE_414E70(f, buf); got != tc.want {
				t.Fatalf("result = %d, want %d", got, tc.want)
			}
			checkThingParserCursor(t, f, len(payload))
		})
	}
}

func TestNoxThingReadWallDefinition(t *testing.T) {
	// The wall format contains sixteen 8-byte-aligned records. Zero counts
	// exercise the complete outer layout without requiring image lookups.
	payload := make([]byte, 160)
	payload = appendThingU32(payload, 0x454e4420)
	f := newThingParserMemFile(t, payload)
	buf := make([]byte, 256*1024)
	if got := Nox_thing_read_WALL_414F60(f, buf); got != 1 {
		t.Fatalf("result = %d, want 1", got)
	}
	checkThingParserCursor(t, f, len(payload))
}

func TestNoxThingReadFloorDefinition(t *testing.T) {
	resetThingDefinitions(t)
	payload := make([]byte, 4)
	payload = append(payload, 4, 't', 'e', 's', 't')
	payload = append(payload, 0xff, 0, 0xff)
	payload = appendThingU32(payload, 0x11223344)
	payload = appendThingU32(payload, 0x55667788)
	payload = append(payload, 1, 0, 0, 0, 7)
	f := newThingParserMemFile(t, payload)
	buf := make([]byte, 256*1024)
	if got := Nox_thing_read_FLOR_411540(f, buf); got != 1 {
		t.Fatalf("result = %d, want 1", got)
	}
	checkThingParserCursor(t, f, len(payload))
}

func TestNoxThingReadEdgeDefinition(t *testing.T) {
	resetThingDefinitions(t)
	payload := make([]byte, 4)
	payload = append(payload, 4, 'e', 'd', 'g', 'e')
	payload = appendThingU32(payload, 0x11223344)
	payload = appendThingU32(payload, 0x55667788)
	payload = append(payload, 1, 0, 0, 0, 0, 0)
	payload = appendThingU32(payload, 0x454e4420)
	f := newThingParserMemFile(t, payload)
	buf := make([]byte, 256*1024)
	if got := Nox_thing_read_EDGE_411850(f, buf); got != 1 {
		t.Fatalf("result = %d, want 1", got)
	}
	checkThingParserCursor(t, f, len(payload))
}

func TestLegacySettingsDisconnected(t *testing.T) {
	oldConnected := Nox_client_isConnected
	Nox_client_isConnected = func() bool { return false }
	t.Cleanup(func() { Nox_client_isConnected = oldConnected })

	index := Sub_409A70(0)
	value := memmap.PtrUint16(0x5D4594, uintptr(3488+2*index))
	oldValue := *value
	oldUpdated := Nox_server_gameDoSwitchMap_40A680()
	t.Cleanup(func() {
		*value = oldValue
		Nox_server_gameUnsetMapLoad_40A690()
		if oldUpdated != 0 {
			Nox_server_gameSettingsUpdated_40A670()
		}
	})

	*value = 17
	Nox_server_gameUnsetMapLoad_40A690()
	Sub_409FB0_settings(0, 17)
	if Nox_server_gameDoSwitchMap_40A680() != 0 {
		t.Fatal("unchanged setting marked the settings as updated")
	}

	Sub_409FB0_settings(0, 23)
	if *value != 23 || Nox_server_gameDoSwitchMap_40A680() != 1 {
		t.Fatalf("changed setting = %d, updated = %d", *value, Nox_server_gameDoSwitchMap_40A680())
	}

	Nox_server_gameUnsetMapLoad_40A690()
	Sub_409FB0_settings(0, 1000)
	if *value != 999 || Nox_server_gameDoSwitchMap_40A680() != 1 {
		t.Fatalf("clamped setting = %d, updated = %d", *value, Nox_server_gameDoSwitchMap_40A680())
	}
}
