package blobdata

import (
	"testing"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestInitDataExtra(t *testing.T) {
	// InitData should not panic and should copy data to blobs
	// The blobs should already be registered by the noxmap init
	// Allocate data if not already allocated (as done in existing test)
	b1 := memmap.BlobByAddr(0x581450)
	if b1 == nil {
		t.Fatal("Blob at 0x581450 should exist")
	}
	if len(b1.Data) == 0 {
		b1.Data = make([]byte, 23472)
	}
	b2 := memmap.BlobByAddr(0x587000)
	if b2 == nil {
		t.Fatal("Blob at 0x587000 should exist")
	}
	if len(b2.Data) == 0 {
		b2.Data = make([]byte, 316820)
	}

	InitData()

	// Verify that data was copied to the blobs
	blob1 := memmap.BlobByAddr(0x581450)
	if blob1 == nil {
		t.Fatal("Blob at 0x581450 should exist")
	}
	if len(blob1.Data) == 0 {
		t.Error("Blob at 0x581450 should have data after InitData")
	}

	blob2 := memmap.BlobByAddr(0x587000)
	if blob2 == nil {
		t.Fatal("Blob at 0x587000 should exist")
	}
	if len(blob2.Data) == 0 {
		t.Error("Blob at 0x587000 should have data after InitData")
	}

	// Verify data was actually copied (not all zeros)
	hasNonZero := false
	for _, b := range blob1.Data {
		if b != 0 {
			hasNonZero = true
			break
		}
	}
	if !hasNonZero {
		t.Error("Blob at 0x581450 should have non-zero data after InitData")
	}

	hasNonZero = false
	for _, b := range blob2.Data {
		if b != 0 {
			hasNonZero = true
			break
		}
	}
	if !hasNonZero {
		t.Error("Blob at 0x587000 should have non-zero data after InitData")
	}
}

func TestInitDataIdempotentExtra(t *testing.T) {
	// Calling InitData multiple times should not panic
	InitData()
	InitData()
	InitData()
}
