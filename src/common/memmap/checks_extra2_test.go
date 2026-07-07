package memmap

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCheckAddrPanic(t *testing.T) {
	// Register a blob and variable, then checkAddr should panic on intersection
	RegisterBlob(0x5001, "testBlobCheck", 100)
	RegisterVariable(0x5011, 4, "testVarCheck", nil)

	require.Panics(t, func() {
		checkAddr(0x5011)
	})

	// Cleanup
	for i, v := range variables {
		if v.Addr == 0x5011 && v.Name == "testVarCheck" {
			variables = append(variables[:i], variables[i+1:]...)
			varsSorted = false
			break
		}
	}
}

func TestValidateZerosPanic(t *testing.T) {
	// Register a blob with non-zero data using RegisterBlobData
	data := make([]byte, 10)
	RegisterBlobData(0x6001, "testBlobZero", data)
	b := BlobByAddr(0x6001)
	require.NotNil(t, b)
	// Make data non-zero but initial zero
	b.Data[0] = 1
	RegisterVariable(0x6001, 1, "testVarZero", nil)
	require.Panics(t, func() {
		ValidateZeros()
	})
	// Restore and cleanup: set data to zero and remove variable
	b.Data[0] = 0
	for i, v := range variables {
		if v.Addr == 0x6001 && v.Name == "testVarZero" {
			variables = append(variables[:i], variables[i+1:]...)
			varsSorted = false
			break
		}
	}
}

func TestValidateZerosNoPanic(t *testing.T) {
	// With no variables or zero data, should not panic
	require.NotPanics(t, func() {
		ValidateZeros()
	})
}

func TestSetRuntimeChecksCoverage(t *testing.T) {
	SetRuntimeChecks(true)
	SetRuntimeChecks(false)
	SetRuntimeChecks(true)
	SetRuntimeChecks(false)
}
