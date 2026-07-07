package opennox

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

func TestNoxXxxMapGenMakeInfoExtra(t *testing.T) {
	// Allocate a buffer large enough for the map info structure
	// The function writes to various offsets up to at least 1392+4
	buf := make([]byte, 2048)
	p := unsafe.Pointer(&buf[0])

	nox_xxx_mapGenMakeInfo_4D5DB0(p)

	// Check that "Generated Map" was written at offset 0
	name := alloc.GoString((*byte)(unsafe.Add(p, 0)))
	if name != "Generated Map" {
		t.Errorf("name at offset 0 = %q, want %q", name, "Generated Map")
	}

	// Check that URL was written at offset 656
	url := alloc.GoString((*byte)(unsafe.Add(p, 656)))
	if url != "http://www.westwood.com" {
		t.Errorf("url at offset 656 = %q, want %q", url, "http://www.westwood.com")
	}

	// Check that the version field at offset 1392 is 3
	vers := *(*uint32)(unsafe.Add(p, 1392))
	if vers != 3 {
		t.Errorf("version at offset 1392 = %d, want 3", vers)
	}
}

func TestNoxXxxMapGenMakeInfoDateExtra(t *testing.T) {
	buf := make([]byte, 2048)
	p := unsafe.Pointer(&buf[0])

	nox_xxx_mapGenMakeInfo_4D5DB0(p)

	// Check that date was written at offset 1360 (format: "Mon, Jan 2 2006")
	date := alloc.GoString((*byte)(unsafe.Add(p, 1360)))
	if len(date) == 0 {
		t.Error("date at offset 1360 is empty")
	}
	// Date should contain a comma (from format "Mon, Jan 2 2006")
	if len(date) > 0 && date[3] != ',' {
		t.Logf("date format may have changed: %q", date)
	}
}
