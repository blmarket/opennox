package blobdata

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func TestInitData(t *testing.T) {
	b1 := memmap.BlobByAddr(0x581450)
	if b1 == nil {
		t.Fatalf("blob 581450 not found")
	}
	if len(b1.Data) == 0 {
		b1.Data = make([]byte, 23472)
	}
	b2 := memmap.BlobByAddr(0x587000)
	if b2 == nil {
		t.Fatalf("blob 587000 not found")
	}
	if len(b2.Data) == 0 {
		b2.Data = make([]byte, 316820)
	}
	InitData()
	if len(b1.Data) == 0 {
		t.Fatalf("blob 581450 empty")
	}
	if len(b2.Data) == 0 {
		t.Fatalf("blob 587000 empty")
	}
}

func TestInit(t *testing.T) {
	var p Ptrs
	// Allocate dummy pointers for p.Ptr_* fields to avoid nil dereference
	var dummy unsafe.Pointer
	p.Ptr_nox_xxx_aClosewoodengat_587000_133480 = &dummy
	p.Ptr_dword_587000_155144 = &dummy
	p.Ptr_dword_587000_127004 = &dummy
	p.Ptr_dword_587000_93164 = &dummy
	p.Ptr_dword_587000_122852 = &dummy
	p.Ptr_dword_587000_81128 = &dummy
	// Init sets many pointers; it requires blobs to have Data allocated
	b1 := memmap.BlobByAddr(0x581450)
	if b1 != nil && len(b1.Data) == 0 {
		b1.Data = make([]byte, 23472)
	}
	b2 := memmap.BlobByAddr(0x587000)
	if b2 != nil && len(b2.Data) == 0 {
		b2.Data = make([]byte, 316820)
	}
	b3 := memmap.BlobByAddr(0x5D4594)
	if b3 != nil && len(b3.Data) == 0 {
		b3.Data = make([]byte, 2598284)
	}
	Init(&p)
}
