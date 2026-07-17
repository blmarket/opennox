package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestScriptBuiltinHandleTable(t *testing.T) {
	first, last, count := C_sub_512E80Sequence()
	require.Zero(t, first)
	require.Equal(t, 1023, last)
	require.Equal(t, uint32(1024), count)
}

func TestScriptBuiltinGuardedObjectCommands(t *testing.T) {
	C_scriptBuiltinGuardedObjects()
}

func TestScriptBuiltinMissingObjectPaths(t *testing.T) {
	vm := setScriptBuiltinServer(t)
	for i := 0; i < 11; i++ {
		vm.stack = nil
		require.Zero(t, C_scriptBuiltinCase(i), "case %d", i)
		switch i {
		case 4:
			require.Equal(t, []uint32{17}, vm.stack)
		case 7:
			require.Equal(t, []uint32{23}, vm.stack)
		case 9, 10:
			require.Equal(t, []uint32{0}, vm.stack)
		}
	}
}
