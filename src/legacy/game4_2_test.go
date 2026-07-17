package legacy

import (
	"math"
	"os"
	"path/filepath"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc/handles"
)

func mapGenInput(t *testing.T, name, data string) unsafe.Pointer {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	require.NoError(t, os.WriteFile(path, []byte(data), 0o600))
	f := C_mapGenOpenInput(path)
	require.NotNil(t, f)
	t.Cleanup(func() { C_mapGenCloseInput(f) })
	return f
}

func TestGame42Pure(t *testing.T) {
	t.Run("sub_521EB0 float rect overlap check", func(t *testing.T) {
		// a1 defines rect via indices 9,10,11,12 ; a2 defines point rect via 1,2,3,4
		// choose a1 = [0..12] with 9:0,10:0,11:10,12:10 => rect 0,0 to 10,10
		// a2 with 1:5,2:5,3:5,4:5 => center 5,5 => should pass
		a1 := make([]float32, 13)
		a1[9] = 0
		a1[10] = 0
		a1[11] = 10
		a1[12] = 10
		a2 := make([]float32, 5)
		a2[1] = 5
		a2[2] = 5
		a2[3] = 5
		a2[4] = 5
		require.Equal(t, 1, C_sub_521EB0(a1, a2))

		// outside to left
		a2[1] = -1
		a2[2] = 5
		a2[3] = -1
		a2[4] = 5
		require.Equal(t, 0, C_sub_521EB0(a1, a2))

		// outside to right
		a2[1] = 11
		a2[3] = 11
		require.Equal(t, 0, C_sub_521EB0(a1, a2))
	})

	t.Run("nox_xxx_mapGenCheckRoomType_5238F0", func(t *testing.T) {
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(2))
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(3))
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(4))
		require.Equal(t, 1, C_nox_xxx_mapGenCheckRoomType_5238F0(5))
		require.Equal(t, 0, C_nox_xxx_mapGenCheckRoomType_5238F0(1))
		require.Equal(t, 0, C_nox_xxx_mapGenCheckRoomType_5238F0(0))
		require.Equal(t, 0, C_nox_xxx_mapGenCheckRoomType_5238F0(6))
	})

	t.Run("mapGen rng helpers", func(t *testing.T) {
		// Seeding makes nox_platform_rand deterministic; both rand funcs must
		// stay within their documented [a1, a2] / [a1-a2, a1+a2] bounds.
		C_nox_xxx_mapGenSetRngSeed_526AB0(12345)

		// result = a1 + (a2-a1+1)*rand()/0x7FFF, clamped to a2.
		for i := 0; i < 200; i++ {
			r := C_nox_xxx_mapGenRandFunc_526AC0(10, 20)
			require.GreaterOrEqual(t, r, 10)
			require.LessOrEqual(t, r, 20)
		}
		// Degenerate range a1==a2 always returns a1.
		for i := 0; i < 20; i++ {
			require.Equal(t, 7, C_nox_xxx_mapGenRandFunc_526AC0(7, 7))
		}
	})

	t.Run("mapGenRandFunc2", func(t *testing.T) {
		C_nox_xxx_mapGenSetRngSeed_526AB0(999)
		// a2==0 short-circuits and returns a1 unchanged.
		require.Equal(t, 42, C_nox_xxx_mapGenRandFunc2_526B00(42, 0))
		// Otherwise result = v4 + a1 with v4 in [-a2, a2].
		for i := 0; i < 100; i++ {
			r := C_nox_xxx_mapGenRandFunc2_526B00(100, 5)
			require.GreaterOrEqual(t, r, 95)
			require.LessOrEqual(t, r, 105)
		}
	})

	t.Run("nox_xxx_isObjectMovable_52E020", func(t *testing.T) {
		// field16 & 0x8068 != 0  => not movable (0)
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0, 0x8000))
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0, 0x0008))
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0, 0x0060))
		// else result = (~field8 >> 22) & 1
		// field8 bit22 clear => ~ has bit22 set => 1
		require.Equal(t, 1, C_nox_xxx_isObjectMovable_52E020(0, 0))
		// field8 bit22 set => ~ clears it => 0
		require.Equal(t, 0, C_nox_xxx_isObjectMovable_52E020(0x00400000, 0))
		// field16 bit not in mask (0x10) => still falls to else branch
		require.Equal(t, 1, C_nox_xxx_isObjectMovable_52E020(0, 0x10))
	})

	t.Run("sub_526AA0 stride", func(t *testing.T) {
		// returns base + (a1 << 6); successive indices differ by 64 bytes.
		base := C_sub_526AA0(0)
		require.Equal(t, uintptr(64), C_sub_526AA0(1)-base)
		require.Equal(t, uintptr(640), C_sub_526AA0(10)-base)
	})
}

func TestGame42RoomConnections(t *testing.T) {
	ret, count, stored := C_sub_521900(0, 0, 0x12345678)
	require.Equal(t, 1, ret)
	require.Equal(t, uint8(1), count)
	require.Equal(t, uint32(0x12345678), stored)

	ret, count, stored = C_sub_521900(7, 3, 0x87654321)
	require.Equal(t, 1, ret)
	require.Equal(t, uint8(8), count)
	require.Equal(t, uint32(0x87654321), stored)

	ret, count, stored = C_sub_521900(8, 2, 0xffffffff)
	require.Equal(t, 0, ret)
	require.Equal(t, uint8(8), count)
	require.Zero(t, stored)
}

func TestGame42DecorConstraints(t *testing.T) {
	// A zero constraint accepts every room; otherwise at least one bit must overlap.
	require.Equal(t, 1, C_nox_xxx_mapGenDecorChkConstaint_5241C0(0, 0))
	require.Equal(t, 1, C_nox_xxx_mapGenDecorChkConstaint_5241C0(0, 0xff))
	require.Equal(t, 1, C_nox_xxx_mapGenDecorChkConstaint_5241C0(0x04, 0x0c))
	require.Equal(t, 0, C_nox_xxx_mapGenDecorChkConstaint_5241C0(0x04, 0x08))

	// The larger room dimension must fit the inclusive decor range.
	require.Equal(t, 1, C_nox_xxx_mapGenChkDecorFillsRoom_5241F0(5, 10, 5, 4))
	require.Equal(t, 1, C_nox_xxx_mapGenChkDecorFillsRoom_5241F0(5, 10, 4, 10))
	require.Equal(t, 0, C_nox_xxx_mapGenChkDecorFillsRoom_5241F0(5, 10, 4, 4))
	require.Equal(t, 0, C_nox_xxx_mapGenChkDecorFillsRoom_5241F0(5, 10, 11, 3))
}

func TestGame42RandomFloatDeterminism(t *testing.T) {
	C_nox_xxx_mapGenSetRngSeed_526AB0(0x42)
	first := C_sub_526BC0(-2.5, 7.25)
	C_nox_xxx_mapGenSetRngSeed_526AB0(0x42)
	require.Equal(t, first, C_sub_526BC0(-2.5, 7.25))
	// A degenerate interval is independent of the platform RNG width.
	require.Equal(t, 3.5, C_sub_526BC0(3.5, 3.5))
}

func TestGame42MapGeneratorGridLifecycle(t *testing.T) {
	got := C_mapGenGridLifecycle()
	require.Equal(t, 1, got.initialized)
	require.Equal(t, 1, got.added)
	require.NotZero(t, got.found)
	require.Equal(t, got.found, got.collided)
	require.Equal(t, 1, got.removed)
	require.True(t, got.topMatches)
	require.True(t, got.emptyAfter)

	require.Zero(t, C_sub_521720(123, true))
	require.Equal(t, 123, C_sub_521720(123, false))
}

func TestGame42MapGeneratorRoomGeometry(t *testing.T) {
	gx, gy, rect := C_mapGenRoomRect(2, 3, 65.1, -65.1, true)
	require.Equal(t, int32(2), gx)
	require.Equal(t, int32(-2), gy)
	require.InDelta(t, 65.1, rect[0], 0.001)
	require.InDelta(t, -65.1, rect[1], 0.001)
	require.InDelta(t, 65.1+2*32.526913, rect[2], 0.001)
	require.InDelta(t, -65.1+3*32.526913, rect[3], 0.001)

	gx, gy, rect = C_mapGenRoomRect(1, 1, 10, 20, false)
	require.Zero(t, gx)
	require.Zero(t, gy)
	require.Equal(t, [4]float32{10, 20, 10 + 32.526913, 20 + 32.526913}, rect)

	made, prepared := C_mapGenMakeAndPrepareRooms()
	require.Equal(t, float32(1), made[0])
	require.Equal(t, float32(7), made[1])
	require.Equal(t, float32(9), made[2])
	require.InDelta(t, 7*32.526913, made[3], 0.001)
	require.InDelta(t, 9*32.526913, made[4], 0.001)
	require.Equal(t, int32(10), prepared[0])
	require.GreaterOrEqual(t, prepared[1], int32(7))
	require.LessOrEqual(t, prepared[1], int32(13))
}

func TestGame42MapGeneratorDirectionsAndEdges(t *testing.T) {
	tests := []struct {
		roomType uint32
		want     int
	}{
		{1, -1}, {2, 1}, {3, 0}, {4, 3}, {5, 2}, {6, -1},
	}
	for _, tc := range tests {
		for direction := 0; direction < 4; direction++ {
			want := 0
			if direction == tc.want {
				want = 1
			}
			require.Equal(t, want, C_sub_521AA0(tc.roomType, int32(direction)), "type=%d direction=%d", tc.roomType, direction)
		}
	}

	a, b := C_mapGenDirections()
	require.Equal(t, []int{0, 0, 0, 1, 2, 3, 0}, a)
	require.Equal(t, []int{0, 0, 3, 2, 0, 1, 0}, b)

	edges := C_mapGenRoomRandEdges()
	// Equal dimensions make each RNG range degenerate, so these values are exact.
	require.InDelta(t, 100, edges[0], 0.001)
	require.InDelta(t, 200, edges[1], 0.001)
	require.InDelta(t, 300, edges[2], 0.001)
	require.InDelta(t, 400, edges[3], 0.001)
}

func TestGame42MapGeneratorRectList(t *testing.T) {
	got := C_mapGenRectList()
	require.Equal(t, 2, got.countBefore)
	require.Equal(t, 1, got.inside)
	require.Zero(t, got.outside)
	require.NotZero(t, got.overlap)
	require.Zero(t, got.separate)
	require.Equal(t, 1, got.countAfter)
	require.Equal(t, 1, got.randomRet, "point=%v", got.random)
	// sub_526BC0 intentionally scales the platform RNG as an unsigned value;
	// characterize only that the produced coordinates are finite here.
	require.False(t, math.IsNaN(float64(got.random[0])))
	require.False(t, math.IsInf(float64(got.random[0]), 0))
	require.False(t, math.IsNaN(float64(got.random[1])))
	require.False(t, math.IsInf(float64(got.random[1]), 0))
}

func TestGame42MapGeneratorUniquePoints(t *testing.T) {
	points := make([][2]float32, 0, 19)
	points = append(points, [2]float32{1, 2}, [2]float32{1, 2}, [2]float32{1.00001, 2.00001})
	for i := 0; i < 16; i++ {
		points = append(points, [2]float32{float32(10 + i), float32(30 + i)})
	}
	counts := C_sub_522CA0(points)
	require.Equal(t, 1, counts[0])
	// The comparison epsilon is loaded from the production data blob. Without
	// provisioning that blob, it is zero and even exact duplicates are appended.
	require.Equal(t, 2, counts[1])
	// The legacy epsilon determines whether this near-equal point is deduplicated;
	// either way, the count may never exceed the fixed sixteen-point capacity.
	for _, n := range counts {
		require.LessOrEqual(t, n, 16)
		require.GreaterOrEqual(t, n, 1)
	}
	require.Equal(t, 16, counts[len(counts)-1])
}

func TestGame42MapGeneratorHallGeometry(t *testing.T) {
	tests := []struct {
		typ          int32
		wantDims     [2]int32
		wantFirst    [2]float32
		wantOpposite [2]float32
		wantCenter   [2]float32
	}{
		{2, [2]int32{3, 7}, [2]float32{30, 20}, [2]float32{30, 80}, [2]float32{30, 20 + 3*32.526913/2}},
		{3, [2]int32{3, 7}, [2]float32{30, 80}, [2]float32{30, 20}, [2]float32{30, 80 - 3*32.526913/2}},
		{4, [2]int32{7, 3}, [2]float32{50, 50}, [2]float32{10, 50}, [2]float32{50 - 3*32.526913/2, 50}},
		{5, [2]int32{7, 3}, [2]float32{10, 50}, [2]float32{50, 50}, [2]float32{10 + 3*32.526913/2, 50}},
	}
	for _, tc := range tests {
		dims, spans, first, opposite, center := C_mapGenHallGeometry(tc.typ, 3, 7)
		require.Equal(t, tc.wantDims, dims, "type=%d", tc.typ)
		require.InDelta(t, float64(dims[0])*32.526913, spans[0], 0.001)
		require.InDelta(t, float64(dims[1])*32.526913, spans[1], 0.001)
		for i := 0; i < 2; i++ {
			require.True(t, math.Abs(float64(first[i]-tc.wantFirst[i])) < 0.001, "type=%d first=%v", tc.typ, first)
			require.True(t, math.Abs(float64(opposite[i]-tc.wantOpposite[i])) < 0.001, "type=%d opposite=%v", tc.typ, opposite)
			require.True(t, math.Abs(float64(center[i]-tc.wantCenter[i])) < 0.001, "type=%d center=%v", tc.typ, center)
		}
	}

	// Invalid types retain zero dimensions and exercise each switch default.
	dims, spans, _, _, _ := C_mapGenHallGeometry(99, 3, 7)
	require.Zero(t, dims)
	require.Zero(t, spans)
}

func TestGame42MapGeneratorAdditionalHeadlessPrimitives(t *testing.T) {
	got := C_mapGenMoreBasics()
	require.True(t, got.bufferAllocated)
	require.True(t, got.bufferReset)
	require.True(t, got.freeEmptyTop)
	require.True(t, got.freeNonEmptyTop)
	require.Equal(t, 1, got.fitSmall)
	require.Zero(t, got.fitTooLarge)
	require.Equal(t, 1, got.fitEmpty)
	require.Zero(t, got.neighborEmpty)
	require.Equal(t, 1, got.neighborRoom)
	require.Zero(t, got.other)
	require.True(t, got.linkedObject)
	require.Contains(t, []int{1, 2, 3}, got.selectedListNode)
	require.Zero(t, got.horizontalZero)
	require.Equal(t, 1, got.verticalZero)
	require.True(t, got.fillZero)
	require.True(t, got.columnZero)
	require.True(t, got.clearHorizontal)
	require.True(t, got.clearVertical)
	require.Equal(t, -1, got.nameMissing)
	require.Equal(t, []int{0, 1, 1, 0}, got.toggleA)
	require.Equal(t, []int{0, 1, 1, 0}, got.toggleB)
	require.Equal(t, []int{0, 1, 1, 1, 0}, got.direction)
	require.Zero(t, got.coordNoOutput)
	require.Equal(t, 1, got.borderWalk)
	require.Equal(t, float32(2), got.recursionParent)
	require.Equal(t, float32(7), got.recursionChild)
}

func TestGame42MapGeneratorHeadlessAlgorithmBatch(t *testing.T) {
	got := C_mapGenHeadlessAlgorithmBatch()
	require.Equal(t, []int{1, 1, 1, 1}, got.connectionReturns)
	require.Equal(t, [][2]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}, got.connectionCounts)
	for _, direction := range got.oppositeDirections {
		require.GreaterOrEqual(t, direction, 0)
		require.LessOrEqual(t, direction, 3)
	}
	require.Equal(t, [][2]int32{{3, 7}, {3, 7}, {7, 3}, {7, 3}, {0, 0}}, got.hallDims)
	require.Equal(t, []bool{true, true, true, true, true, true}, got.selectionReturns)
	require.Equal(t, []uint8{0, 0, 0, 0, 0, 0}, got.selectionLimits)
	require.Equal(t, []uint8{1, 1, 1, 1, 1, 1}, got.selectionDisabled)
	require.Equal(t, []int32{10, 0, 0, 0, 0, 10}, got.selectionWeights)
	require.Zero(t, got.stackEmptyRet)
	require.Equal(t, 1, got.stackRet)
	require.Equal(t, [3]uint32{11, 22, 33}, got.stackValues)
	require.Equal(t, 34, got.wallCountRet)
	require.Equal(t, [2]uint32{12, 34}, got.wallMins)
	require.Equal(t, [2]int{-5, 5}, got.assignReturns)
	require.Equal(t, [2]int32{1, 5}, got.assignValues)
}

func TestGame42MapGeneratorTokenReader(t *testing.T) {
	handles.Init()
	f := mapGenInput(t, "tokens.txt", "  alpha\tbeta // ignored words\n gamma\n")
	require.Equal(t, []string{"alpha", "beta", "gamma"}, C_mapGenReadTokens(f, true))

	f = mapGenInput(t, "conditional.txt", "IF 0% hidden ELSE shown ENDIF tail\n")
	require.Equal(t, []string{"shown", "tail"}, C_mapGenReadTokens(f, true))

	// A nested false branch forces sub_51E780 to delegate to sub_51E720.
	f = mapGenInput(t, "nested.txt", "IF 0% IF 0% hidden ENDIF ELSE visible ENDIF\n")
	require.Equal(t, []string{"visible"}, C_mapGenReadTokens(f, true))
}

func TestGame42MapGeneratorAlgorithmParser(t *testing.T) {
	handles.Init()
	// The decompiled success path stores a token through a one-byte local and
	// trips modern stack protection. Characterize the safe empty section here;
	// richer algorithm parsing remains an end-to-end-only path.
	f := mapGenInput(t, "algorithm.txt", "END\n")
	ret, got := C_mapGenReadAlgorithmData(f)
	require.Equal(t, 1, ret)
	require.Zero(t, got.ints)
	require.Zero(t, got.mapSz)
}

func TestGame42MapGeneratorConditionOperators(t *testing.T) {
	handles.Init()
	f := mapGenInput(t, "operators.txt", "< > == != invalid\n")
	rets, values := C_mapGenReadOperators(f, 5)
	require.Equal(t, []int{1, 1, 1, 1, 0}, rets)
	require.Equal(t, []int{0, 1, 2, 3, 99}, values)

	f = mapGenInput(t, "conditions.txt", "EXPERIENCE_LEVEL < 4 EXPERIENCE_LEVEL > 4 EXPERIENCE_LEVEL == 3 EXPERIENCE_LEVEL != 3 100% 0% invalid\n")
	rets, values = C_mapGenReadConditions(f, 7)
	require.Equal(t, []int{1, 1, 1, 1, 1, 1, 0}, rets)
	require.Equal(t, []int{1, 0, 1, 0, 1, 0, 99}, values)
}

func TestGame42MapGeneratorThemeSectionParsers(t *testing.T) {
	handles.Init()
	spell := mapGenInput(t, "spells.txt", "END\n")
	// Nonempty weapon/armor sections call sub_51F230, whose decompiled 75 KiB
	// stack frame overflows under GCC coverage instrumentation. The empty
	// section paths remain deterministic and characterize their lifecycle.
	weapon := mapGenInput(t, "weapons.txt", "END\n")
	armor := mapGenInput(t, "armor.txt", "END\n")
	exit := mapGenInput(t, "exit.txt", "OBJECT NorthExit NORTH OBJECT SouthExit SOUTH OBJECT EastExit EAST OBJECT WestExit WEST LINKDATA PortalLink END\n")
	got := C_mapGenReadThemeSections(spell, weapon, armor, exit)
	require.Equal(t, 1, got.spellRet)
	require.Equal(t, 1, got.weaponRet)
	require.Equal(t, 1, got.armorRet)
	require.Equal(t, 1, got.exitRet)
	require.Zero(t, got.spellCount)
	require.Zero(t, got.weaponCount)
	require.Empty(t, got.weaponName)
	require.Empty(t, got.weaponBase)
	require.Zero(t, got.armorCount)
	require.Empty(t, got.armorName)
	require.Empty(t, got.armorBase)
	require.Equal(t, uint32(4), got.exitCount)
	require.Equal(t, [4]string{"NorthExit", "SouthExit", "EastExit", "WestExit"}, got.exitNames)
	require.Equal(t, [4]uint32{0, 1, 2, 3}, got.exitDirections)
	require.Equal(t, "PortalLink", got.linkData)
}

func TestGame42MapGeneratorFullDecorParser(t *testing.T) {
	handles.Init()
	for kind, section := range []string{"ROOM", "HALL", "TEMPLATE", "BACKDROP"} {
		f := mapGenInput(t, section+".txt", section+" Decor"+section+" OCCUR_CONSTRAINT NONE OCCUR_LIMIT 9 MUST_OCCUR FREQUENCY RARE DOOR DoorName DOUBLE_DOOR DoubleDoorName ROOM_SIZE_CONSTRAINT 3 8 END\n")
		got := C_mapGenReadFullDecor(f)
		require.Equal(t, 1, got.ret, "section=%s", section)
		require.Equal(t, kind, got.kind, "section=%s", section)
		require.Equal(t, 1, got.count, "section=%s", section)
		require.Equal(t, "Decor"+section, got.name)
		require.Zero(t, got.constraint)
		require.Equal(t, uint8(9), got.limit)
		require.Equal(t, uint8(1), got.must)
		require.Equal(t, uint32(100), got.frequency)
		require.Equal(t, int32(3), got.min)
		require.Equal(t, int32(8), got.max)
		require.Equal(t, "DoorName", got.door)
		require.Equal(t, "DoubleDoorName", got.doubleDoor)
	}
}

func TestGame42MapGeneratorDecorParsers(t *testing.T) {
	handles.Init()
	f := mapGenInput(t, "decor.txt", "NONE 7 VERY_RARE * 12 OakDoor DoubleDoor\n")
	rets, got := C_mapGenReadDecorPrimitives(f)
	require.Equal(t, [6]int{1, 1, 1, 1, 1, 1}, rets)
	require.Zero(t, got.constraint)
	require.Equal(t, uint8(7), got.limit)
	require.Equal(t, uint32(10), got.frequency)
	require.Zero(t, got.min)
	require.Equal(t, int32(12), got.max)
	require.Equal(t, "OakDoor", got.door)
	require.Equal(t, "DoubleDoor", got.doubleDoor)

	f = mapGenInput(t, "frequencies.txt", "COMMON UNCOMMON RARE VERY_RARE HARDLY_EVER UNKNOWN\n")
	rets2, values := C_mapGenReadFrequencies(f, 6)
	require.Equal(t, []int{1, 1, 1, 1, 1, 1}, rets2)
	require.Equal(t, []int{1000, 500, 100, 10, 1, 1}, values)

	success, sums, failures := C_mapGenCheckSyntheticSettings()
	require.Equal(t, 1, success)
	require.Equal(t, [6]int32{3, 4, 5, 6, 7, 8}, sums)
	require.Equal(t, []int{0, 0, 0, 0, 0, 0}, failures)
}

func TestGame42MapGeneratorCompositeParsers(t *testing.T) {
	handles.Init()
	wall := mapGenInput(t, "wall.txt", "Stone Marble\n")
	appendFile := mapGenInput(t, "append.txt", "first second\n")
	setFile := mapGenInput(t, "set.txt", "2 5 END\n")
	containsFile := mapGenInput(t, "contains.txt", "* END\n")
	prefabFile := mapGenInput(t, "prefab.txt", "MUST_OCCUR AREAMAP AreaOne END\n")
	got := C_mapGenCompositeParsers(wall, appendFile, setFile, containsFile, prefabFile)
	require.Equal(t, 1, got.wallFloorRet)
	require.Equal(t, "Stone", got.wallName)
	require.Equal(t, "Marble", got.floorName)
	require.Equal(t, 1, got.appendRet)
	require.Equal(t, uint32(4), got.appendType)
	require.Equal(t, "first", got.appendFirst)
	require.Equal(t, "second", got.appendSecond)
	require.Equal(t, 1, got.setRet)
	require.Equal(t, int32(2), got.setMin)
	require.Equal(t, int32(5), got.setMax)
	require.Zero(t, got.setCount)
	require.Equal(t, uint32(101), got.containsWeight)
	require.Equal(t, 1, got.prefabRet)
	require.Equal(t, 1, got.areaRet)
	require.Equal(t, "AreaOne", got.prefabName)
}
