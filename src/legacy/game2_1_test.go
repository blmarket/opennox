package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame21Pure(t *testing.T) {
	require.Equal(t, 1, C_sub_464B40(0, 0))
	require.Equal(t, 1, C_sub_464B40(3, 20))
	require.Equal(t, 0, C_sub_464B40(-1, 0))
	require.Equal(t, 0, C_sub_464B40(4, 0))
	require.Equal(t, 0, C_sub_464B40(0, 21))

	require.Equal(t, 0, C_nox_xxx_XorEaxEaxSub_464BA0())
	require.Equal(t, 1, C_nox_xxx_movEax1Sub_4661C0())
	require.Equal(t, 0, C_sub_46F060())
	require.Equal(t, 0, C_nox_xxx_Proc_46F070())

	require.Equal(t, 0, C_nox_xxx_inventoryWndProc_464BB0(0, 8))
	require.Equal(t, 0, C_nox_xxx_inventoryWndProc_464BB0(0, 12))
	require.Equal(t, 0, C_nox_xxx_inventoryWndProc_464BB0(0, 16))
	require.Equal(t, 1, C_nox_xxx_inventoryWndProc_464BB0(0, 0))
	require.Equal(t, 1, C_nox_xxx_inventoryWndProc_464BB0(0, 7))
}

func TestGame21Write46AEC0(t *testing.T) {
	ret, written := C_sub_46AEC0(true, 123)
	require.Equal(t, 0, ret)
	require.Equal(t, uint32(123), written)

	ret, _ = C_sub_46AEC0(false, 5)
	require.Equal(t, -2, ret)
}

func TestGame21Arith4739E0(t *testing.T) {
	// a3 = a2 + *a1 - a1[4] ; a3+4 = a2_4 + a1[1] - a1[5]
	ret, out0, out4 := C_sub_4739E0(10, 20, 3, 5, 100, 200)
	require.Equal(t, uint32(200), ret)
	require.Equal(t, uint32(107), out0) // 100+10-3
	require.Equal(t, uint32(215), out4) // 200+20-5
}

func TestGame21Arith473A10(t *testing.T) {
	ret, out0, out1 := C_sub_473A10(10, 20, 3, 5, 100, 200)
	require.Equal(t, uint32(200), ret)
	require.Equal(t, uint32(93), out0)  // 100+3-10
	require.Equal(t, uint32(185), out1) // 200+5-20
}

func TestGame21Read46AF40(t *testing.T) {
	require.Equal(t, uintptr(0), C_sub_46AF40(false, 0))
	v := C_sub_46AF40(true, 0xAABBCCDD)
	require.Equal(t, uintptr(0xAABBCCDD), v)
}
