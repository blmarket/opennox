package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGame5Pure(t *testing.T) {
	t.Run("nox_xxx_strikeBomber_549BB0 returns 1", func(t *testing.T) {
		require.Equal(t, 1, C_nox_xxx_strikeBomber_549BB0())
	})

	t.Run("sub_5456B0 sets 0x4000 when 0x10000 clear", func(t *testing.T) {
		ret, after := C_sub_5456B0(0)
		require.Equal(t, uint32(0x4000), ret)
		require.Equal(t, uint32(0x4000), after)

		ret2, after2 := C_sub_5456B0(0x10000)
		require.Equal(t, uint32(0x10000), ret2)
		require.Equal(t, uint32(0x10000), after2)

		ret3, after3 := C_sub_5456B0(0x4000)
		require.Equal(t, uint32(0x4000), ret3)
		require.Equal(t, uint32(0x4000), after3)
	})

	t.Run("sub_5456C0 clears 0x4000 when 0x8000 clear", func(t *testing.T) {
		ret, after := C_sub_5456C0(0x4000)
		require.Equal(t, uint32(0), ret)
		require.Equal(t, uint32(0), after)

		ret2, after2 := C_sub_5456C0(0x8000 | 0x4000)
		require.Equal(t, uint32(0xC000), ret2)
		require.Equal(t, uint32(0xC000), after2)

		ret3, after3 := C_sub_5456C0(0)
		require.Equal(t, uint32(0), ret3)
		require.Equal(t, uint32(0), after3)
	})

	t.Run("sub_546410 same as 5456C0", func(t *testing.T) {
		ret, after := C_sub_546410(0x4000)
		require.Equal(t, uint32(0), after)
		require.Equal(t, uint32(0), ret)
	})

	t.Run("sub_546420 same as 5456C0", func(t *testing.T) {
		_, after := C_sub_546420(0x4000)
		require.Equal(t, uint32(0), after)
	})

	t.Run("sub_547EE0 checks flags at 480 and 477", func(t *testing.T) {
		require.Equal(t, 0, C_sub_547EE0(0, 0xFF, 1))
		require.Equal(t, 0, C_sub_547EE0(1, 0, 1))
		require.Equal(t, 1, C_sub_547EE0(1, 0x01, 0x01))
		require.Equal(t, 0, C_sub_547EE0(1, 0x02, 0x01))
		require.Equal(t, 1, C_sub_547EE0(1, 0xFF, 0x80))
	})
}
