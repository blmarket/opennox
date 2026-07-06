package server

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerMapDebug(t *testing.T) {
	var d serverMapDebug
	d.Reset()
	d.Add("key1", "val1")
	d.Add("key1", "val2")
	d.Add("key2", "val3")
	require.Equal(t, []string{"val1", "val2"}, d.Get("key1"))
	cnt := 0
	d.Each(nil, func(key, val string) {
		cnt++
	})
	require.Equal(t, 3, cnt)
	cnt2 := 0
	d.Each([]string{"key1"}, func(key, val string) {
		cnt2++
	})
	require.Equal(t, 2, cnt2)
}
