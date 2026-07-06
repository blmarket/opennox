package memmap

import (
	"testing"
	"unsafe"
)

func TestVariableContains(t *testing.T) {
	v := Variable{Addr: 0x1000, Size: 100}
	if !v.Contains(0x1000) {
		t.Error("Contains start should be true")
	}
	if !v.Contains(0x1050) {
		t.Error("Contains middle should be true")
	}
	if v.Contains(0x2000) {
		t.Error("Contains outside should be false")
	}
}

func TestBlobContains(t *testing.T) {
	b := Blob{Addr: 0x2000, Size: 200}
	if !b.Contains(0x2000) {
		t.Error("Blob Contains start should be true")
	}
	if b.Contains(0x3000) {
		t.Error("Blob Contains outside should be false")
	}
}

func TestBlobContainsPtr(t *testing.T) {
	b := Blob{Addr: 0x1000, Size: 100, Data: make([]byte, 100)}
	ptr := unsafe.Pointer(&b.Data[0])
	if _, ok := b.ContainsPtr(ptr); !ok {
		t.Error("ContainsPtr should work")
	}
	// Test with nil data
	b2 := Blob{Addr: 0x1000, Size: 100}
	if _, ok := b2.ContainsPtr(ptr); ok {
		t.Error("ContainsPtr should fail with empty data")
	}
}

func TestRegisterBlob(t *testing.T) {
	RegisterBlob(0x3000, "testBlob", 50)
	b := BlobByAddr(0x3000)
	if b == nil {
		t.Error("BlobByAddr should find registered blob")
	}
}

func TestRegisterVariable(t *testing.T) {
	RegisterVariable(0x4000, 4, "testVar", nil)
	v := VariableByAddr(0x4000)
	if v == nil {
		t.Error("VariableByAddr should find registered variable")
	}
}

func TestBlobsAndVariables(t *testing.T) {
	blobs := Blobs()
	_ = blobs
	vars := Variables()
	_ = vars
}

func TestPtrFunctions(t *testing.T) {
	// These functions require CGO and specific memory layout,
	// just verify they exist and don't panic on package init
	// Actual testing is done via memmap_nox subpackage
}
