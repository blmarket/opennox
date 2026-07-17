package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerTradeEmptyAndCapacityGuards(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	h.setGameFlags(0)

	empty, missing := C_sub_50E7A0Missing()
	require.Zero(t, empty)
	require.Zero(t, missing)
	require.True(t, C_nox_xxx_servSendShopItemsEmpty())

	same, state := C_nox_xxx_tradeSetPlayerNonPlayer()
	require.True(t, same)
	require.Equal(t, uint32(1), state)

	emptyCapacity, sameType, overflow := C_sub_50FD60Capacity()
	require.Equal(t, 1, emptyCapacity)
	require.Equal(t, 1, sameType)
	require.Zero(t, overflow)
}

func TestServerTradeUtilityGuards(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	h.setGameFlags(0)

	C_sub_510320Guards()
	require.Equal(t, 1, C_sub_510540NonQuest())
	require.Equal(t, [4]int{1, 1, 1, 0}, C_sub_5105D0Cached())
}
