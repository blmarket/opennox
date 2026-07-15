package legacy

import (
	"image"
	"math"
	"testing"

	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/stretchr/testify/require"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/server"
)

func TestGame52NetworkCodeHelpers(t *testing.T) {
	t.Run("net get unit code handles nil and invalid code", func(t *testing.T) {
		require.Equal(t, uint32(0), C_nox_xxx_netGetUnitCodeCli_578B00_nil())
		require.Equal(t, uint32(0), C_nox_xxx_netGetUnitCodeCli_578B00(0x8000, 0))
	})

	t.Run("net get unit code marks dynamic extent units", func(t *testing.T) {
		require.Equal(t, uint32(0x1234), C_nox_xxx_netGetUnitCodeCli_578B00(0x1234, 0))
		require.Equal(t, uint32(0x9234), C_nox_xxx_netGetUnitCodeCli_578B00(0x1234, 0x20400000))
		require.Equal(t, uint32(0x8012), C_nox_xxx_netGetUnitCodeCli_578B00(0x12, 0x400000))
	})

	t.Run("clear high bit masks bit 15", func(t *testing.T) {
		require.Equal(t, 0, C_nox_xxx_netClearHighBit_578B30(0))
		require.Equal(t, 0x1234, C_nox_xxx_netClearHighBit_578B30(0x1234))
		require.Equal(t, 0, C_nox_xxx_netClearHighBit_578B30(-0x8000))
		require.Equal(t, 0x7FFF, C_nox_xxx_netClearHighBit_578B30(-1))
	})

	t.Run("test high bit reads bit 15 only", func(t *testing.T) {
		require.Equal(t, uint32(0), C_nox_xxx_netTestHighBit_578B70(0))
		require.Equal(t, uint32(0), C_nox_xxx_netTestHighBit_578B70(0x7FFF))
		require.Equal(t, uint32(1), C_nox_xxx_netTestHighBit_578B70(0x8000))
		require.Equal(t, uint32(0), C_nox_xxx_netTestHighBit_578B70(0x10000))
		require.Equal(t, uint32(1), C_nox_xxx_netTestHighBit_578B70(0x18000))
	})
}

func TestGame52OffsetAccessors(t *testing.T) {
	t.Run("waypoint next reads offset 484", func(t *testing.T) {
		require.Equal(t, 0, C_nox_xxx_waypointNext_579870_nil())
		require.Equal(t, 0x11223344, C_nox_xxx_waypointNext_579870(0x11223344))
	})

	t.Run("sub 5798A0 reads offset 484", func(t *testing.T) {
		require.Equal(t, 0, C_sub_5798A0_nil())
		require.Equal(t, 0x55667788, C_sub_5798A0(0x55667788))
	})

	t.Run("sub 57BA10 writes packed fields", func(t *testing.T) {
		ret, out0, out2, out4 := C_sub_57BA10(-2, 0x1234, 0x89ABCDEF)
		require.NotZero(t, ret)
		require.Equal(t, uint16(0xFFFE), out0)
		require.Equal(t, uint16(0x1234), out2)
		require.Equal(t, uint32(0x89ABCDEF), out4)
	})

	t.Run("next map group reads offset 88", func(t *testing.T) {
		require.Equal(t, 0, C_nox_server_getNextMapGroup_57C090_nil())
		require.Equal(t, 0x01020304, C_nox_server_getNextMapGroup_57C090(0x01020304))
	})
}

func TestGame52CollisionDirectionStateMachine(t *testing.T) {
	h := newGameLogicHarness(t)
	srv := h.setServer(0, 30)
	require.Equal(t, 1, srv.Walls.Init())
	t.Cleanup(srv.Walls.Free)

	const y = 20
	for direction := 0; direction <= 10; direction++ {
		wall := srv.Walls.CreateAtGrid(image.Pt(20+direction, y))
		require.NotNil(t, wall)
		wall.Dir0 = byte(direction)
	}

	// First-point offsets exercise every result of sub_57F2A0. The second
	// point covers each quadrant mask produced by sub_57F1D0.
	offsets := []float32{0, 5, 11, 12, 18, 22}
	quadrants := [][2]float32{{0, 0}, {22, 0}, {22, 22}, {0, 22}}
	seenOutput := make(map[[2]float32]bool)
	for direction := 0; direction <= 10; direction++ {
		x := 20 + direction
		for _, dx := range offsets {
			for _, dy := range offsets {
				for _, quadrant := range quadrants {
					ret, output := C_sub_57CDB0(
						[2]int32{int32(x), y},
						[4]float32{23*float32(x) + dx, 23*float32(y) + dy, quadrant[0], quadrant[1]},
					)
					require.Equal(t, 1, ret, "direction=%d dx=%g dy=%g quadrant=%v", direction, dx, dy, quadrant)
					require.InDelta(t, 0.70709997, abs32(output[0]), 0.000001)
					require.InDelta(t, 0.70709997, abs32(output[1]), 0.000001)
					seenOutput[output] = true
				}
			}
		}
	}
	require.Equal(t, 4, len(seenOutput), "all four diagonal outputs should be reachable")

	ret, output := C_sub_57CDB0([2]int32{-1, -1}, [4]float32{})
	require.Zero(t, ret)
	require.Equal(t, [2]float32{}, output)
}

func abs32(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func TestGame52CollisionMath(t *testing.T) {
	require.Equal(t, [2]float32{4, 3}, C_nox_xxx_collideReflect_57B810([2]float32{1, -1}, [2]float32{3, 4}))
	require.Equal(t, [2]float32{-4, -3}, C_nox_xxx_collideReflect_57B810([2]float32{1, 1}, [2]float32{3, 4}))

	inside := [11]float32{}
	inside[5], inside[6] = -1, -2
	inside[7], inside[8] = 0, 1
	inside[9], inside[10] = 1, 1
	require.Equal(t, 1, C_nox_xxx_map_57B850([2]float32{}, inside, [2]float32{}))

	for _, mutate := range []func(*[11]float32){
		func(v *[11]float32) { v[5], v[6] = -2, -1 },
		func(v *[11]float32) { v[7], v[8] = 1, 0 },
		func(v *[11]float32) { v[9], v[10] = -1, -1 },
		func(v *[11]float32) { v[5], v[6] = 1, -0.5 },
	} {
		data := inside
		mutate(&data)
		require.Equal(t, 0, C_nox_xxx_map_57B850([2]float32{}, data, [2]float32{}))
	}

	ret, allZero := C_sub_57B920()
	require.Zero(t, ret)
	require.True(t, allZero)
}

func TestGame52AliasSelection(t *testing.T) {
	require.Equal(t, int8(5), C_nox_xxx_cliGenerateAlias_57B9A0(5, -1, false))
	require.Equal(t, int8(8), C_nox_xxx_cliGenerateAlias_57B9A0(5, 8, false))
	require.Equal(t, int8(-1), C_nox_xxx_cliGenerateAlias_57B9A0(5, -1, true))
	// Reserved aliases zero and 255 both start their search at slot one.
	require.Equal(t, int8(1), C_nox_xxx_cliGenerateAlias_57B9A0(0, -1, false))
	require.Equal(t, int8(1), C_nox_xxx_cliGenerateAlias_57B9A0(255, -1, false))
}

func TestGame52LineProjection(t *testing.T) {
	for _, tc := range []struct {
		line  [4]float32
		point [2]float32
		scale float32
		want  [2]float32
	}{
		{line: [4]float32{0, 0, 10, 0}, point: [2]float32{5, 5}, scale: 10, want: [2]float32{5, 0}},
		{line: [4]float32{10, 0, 0, 0}, point: [2]float32{15, 2}, scale: 10, want: [2]float32{10, 0}},
		{line: [4]float32{0, 10, 0, 0}, point: [2]float32{2, -5}, scale: 10, want: [2]float32{0, 0}},
		{line: [4]float32{0, 0, 0, 10}, point: [2]float32{2, 15}, scale: 10, want: [2]float32{0, 10}},
	} {
		require.Equal(t, tc.want, C_sub_57C790(tc.line, tc.point, tc.scale))
	}

	ret, output := C_nox_xxx_mathPointOnTheLine_57C8A0([4]float32{0, 0, 10, 0}, [2]float32{5, 4})
	require.Equal(t, 1, ret)
	require.Equal(t, [2]float32{5, 0}, output)

	ret, output = C_nox_xxx_mathPointOnTheLine_57C8A0([4]float32{10, 10, 0, 0}, [2]float32{12, 12})
	require.Zero(t, ret)
	require.Equal(t, [2]float32{12, 12}, output)

	ret, output = C_nox_xxx_mathPointOnTheLine_57C8A0([4]float32{0, 10, 10, 0}, [2]float32{5, 5})
	require.Equal(t, 1, ret)
	require.Equal(t, [2]float32{5, 5}, output)
}

func TestGame52ProtectionPrimitives(t *testing.T) {
	require.Zero(t, C_nox_xxx_protectionStringCRC_56FAC0(nil, 0))
	require.Equal(t, int(0x11111111^0x22222222^0x44444444), C_nox_xxx_protectionStringCRC_56FAC0(
		[]uint32{0x11111111, 0x22222222, 0x44444444, 0x88888888}, 12,
	))
	// Partial trailing words are intentionally ignored.
	require.Equal(t, 0x11111111, C_nox_xxx_protectionStringCRC_56FAC0([]uint32{0x11111111, 0x22222222}, 7))

	require.Zero(t, C_sub_56FCB0(7, false))
	require.Equal(t, 1, C_sub_56FCB0(0, true))
	require.Equal(t, 1<<7, C_sub_56FCB0(7, true))
	require.Equal(t, uint32(1<<31), uint32(C_sub_56FCB0(31, true)))
	require.Equal(t, 1, C_sub_56FCB0(32, true))

	left, right := C_sub_56F720([2]uint32{1, 2}, [2]uint32{3, 4})
	require.Equal(t, [2]uint32{3, 4}, left)
	require.Equal(t, [2]uint32{1, 2}, right)
	C_sub_56F720_nil()
}

func TestGame52ProtectionRandomState(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()

	for _, seed := range []int{0, 1, 0x12345678, -1} {
		C_sub_56FF00(seed)
		for _, bounds := range [][2]int{{0, 0}, {1, 10}, {-10, 10}, {100, 99}} {
			got := C_sub_56FF80(bounds[0], bounds[1])
			low, high := bounds[0], bounds[1]
			if high < low {
				low, high = high, low
			}
			require.GreaterOrEqual(t, got, low)
			require.LessOrEqual(t, got, high+1)
		}
	}
}

func TestGame52ProtectionListLifecycle(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	setGame52Server(t, 0)

	got := C_protectionListLifecycle()
	require.True(t, got.EmptyFind)
	require.True(t, got.FirstCreated)
	require.True(t, got.SecondCreated)
	require.Equal(t, uint16(2), got.CountAfterAdd)
	require.Equal(t, uint32(42), got.DecodedFirst)
	require.Equal(t, uint32(math.Float32bits(3.5)), got.DecodedSecond)
	require.True(t, got.IndexedBoth)
	require.True(t, got.MissingIndexNil)
	require.True(t, got.RemovedFirst)
	require.False(t, got.RemovedAgain)
	require.True(t, got.ClearedHandle)
	require.Zero(t, got.CountAfterClear)
}

func TestGame52ProtectionTypedUpdates(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	setGame52Server(t, 1234)
	C_sub_56FF00(0x1234)

	got := C_protectionTypedUpdates()
	require.Equal(t, 11, got.UpdatesFound)
	require.Equal(t, 1, got.ValidCRC)
	require.Zero(t, got.InvalidCRC)
	require.Equal(t, 1, got.ValidFlags)
	require.Zero(t, got.InvalidFlags)
	require.Zero(t, got.NilItemCRC)
	require.Equal(t, 8, got.MissingCalls)
	require.Equal(t, uint16(1), got.FinalCount)
}

func TestGame52ProtectionInitialization(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	setGame52Server(t, 77)

	handle, count := C_protectionInitialize()
	require.GreaterOrEqual(t, handle, 657757279)
	require.Equal(t, uint16(9), count)
}

func TestGame52RemainingUtilities(t *testing.T) {
	allocated, flagSet := C_sub_579E70()
	require.True(t, allocated)
	require.True(t, flagSet)
	require.Equal(t, 0x1234, C_nox_xxx_packetDynamicUnitCode_578B40(0x1234))

	h := newGameLogicHarness(t)
	h.loadBlobData()
	h.setGameFlags(0)
	require.Zero(t, C_nox_xxx_playerCanTalkMB_57A160(1<<3, true))
	require.Zero(t, C_nox_xxx_playerCanTalkMB_57A160(1<<3, false))
	h.setGameFlags(noxflags.GameClient)
	require.Equal(t, 1, C_nox_xxx_playerCanTalkMB_57A160(1<<3, false))
	require.Zero(t, C_nox_xxx_playerCanTalkMB_57A160(0, false))

	lookup := C_game52RuleLookups()
	require.NotNil(t, lookup.KnownRule)
	require.Equal(t, 0x80, lookup.KnownMask)
	require.Nil(t, lookup.UnknownRule)
	require.Zero(t, lookup.UnknownMask)

	get, afterReset, _ := C_game52SimpleGlobals(0)
	require.Zero(t, get)
	require.Zero(t, afterReset)
	require.Equal(t, 4, C_sub_57B190(1, 0))
	require.Equal(t, 0, C_sub_57B190(100, 100))
	require.Contains(t, []int{1, 2, 3}, C_sub_57B190(1, 100))

	ret, fields := C_game52RuleDefaults(0x1234)
	require.Equal(t, int8(0x34), ret)
	require.Equal(t, [7]int32{-1, -1, -1, -1, -1, -1, -1}, fields)
}

func setGame52Server(t *testing.T, frame uint32) *server.Server {
	t.Helper()
	oldGetServer := GetServer
	srv := server.New(nil, strman.New())
	srv.SetFrame(frame)
	GetServer = func() Server { return &gameLogicServer{srv: srv} }
	t.Cleanup(func() {
		GetServer = oldGetServer
		srv.Close()
	})
	return srv
}
