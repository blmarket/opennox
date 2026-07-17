package legacy

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/stretchr/testify/require"
)

func drawParseMemfile(t *testing.T, data []byte) unsafe.Pointer {
	t.Helper()
	raw, _ := alloc.CloneSlice(data)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	t.Cleanup(f.Free)
	return f.C()
}

func TestClientDrawAnimationKinds(t *testing.T) {
	for name, want := range map[string]int{
		"OneShot":       0,
		"OneShotRemove": 1,
		"Loop":          2,
		"LoopAndFade":   3,
		"Random":        4,
		"Slave":         5,
		"unknown":       0,
	} {
		require.Equal(t, want, C_getAnimationKindID(name), name)
	}
}

func TestClientDrawVectorAnimationParsing(t *testing.T) {
	data := append([]byte{7, 9, byte(len("LoopAndFade"))}, []byte("LoopAndFade")...)
	ret, frames, delay, kind := C_loadVectorAnimatedHeader(drawParseMemfile(t, data))
	require.Equal(t, 1, ret)
	require.Equal(t, uint16(7), frames)
	require.Equal(t, uint16(9), delay)
	require.Equal(t, 3, kind)

	require.Equal(t, 1, C_loadVectorAnimatedEmptyDirections())
}

func TestClientDrawStaticRandomEmpty(t *testing.T) {
	ok, size, count := C_loadStaticRandomEmpty(drawParseMemfile(t, []byte{0}))
	require.True(t, ok)
	require.Equal(t, uint32(12), size)
	require.Zero(t, count)
}
