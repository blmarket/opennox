package netstr

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEncodeDecodeMessage2(t *testing.T) {
	out := make([]byte, 100)
	n := encodeMessage(out, nil)
	require.Equal(t, 0, n)
}

func TestNetCrypt2(t *testing.T) {
	data := []byte{1, 2, 3, 4, 5, 6, 7, 8}
	netCrypt(1, data)
	require.Len(t, data, 8)
	dst := make([]byte, 8)
	netCryptDst(1, data, dst)
	require.Len(t, dst, 8)
}

func TestNewStreams2(t *testing.T) {
	s := NewStreams(func() uint32 { return 0 })
	require.NotNil(t, s)
	require.Nil(t, s.Host())
}

func TestErrIsInUse2(t *testing.T) {
	require.False(t, ErrIsInUse(nil))
}

func TestNewConnectErr2(t *testing.T) {
	err := NewConnectErr(0, nil)
	require.Equal(t, -1, err.Code)
	err2 := NewConnectErr(5, nil)
	require.Equal(t, 5, err2.Code)
	require.Contains(t, err2.Error(), "CONNECT_SERVER")
}
