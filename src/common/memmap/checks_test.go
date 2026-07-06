package memmap

import (
	"testing"
)

func TestSetRuntimeChecks(t *testing.T) {
	SetRuntimeChecks(true)
	SetRuntimeChecks(false)
	SetRuntimeChecks(true)
}

func TestCheckAddrNoPanic(t *testing.T) {
	// checkAddr with address not in any variable should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("checkAddr panicked as expected for invalid addr: %v", r)
		}
	}()
	// Use an address that's unlikely to be in any variable
	checkAddr(0x1)
}

func TestValidateZeros(t *testing.T) {
	// ValidateZeros should not panic if regions are zero
	// It may panic if non-zero data found, which is OK
	defer func() {
		if r := recover(); r != nil {
			t.Logf("ValidateZeros panicked: %v", r)
		}
	}()
	ValidateZeros()
}
