package legacy

import (
	"runtime"
	"strings"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/stretchr/testify/require"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/common/memmap/nox/blobdata"
	"github.com/noxworld-dev/opennox/v1/common/ntype"
	"github.com/noxworld-dev/opennox/v1/server"
)

// gameLogicHarness builds the ABI-backed values used by isolated GAME*.c
// logic. Keep it Go-only: tests should enter C through production wrappers,
// not through test-specific C functions.
type gameLogicHarness struct {
	t      testing.TB
	pinner runtime.Pinner
}

type gameLogicServer struct {
	Server
	srv *server.Server
}

func (s *gameLogicServer) S() *server.Server { return s.srv }

func newGameLogicHarness(t testing.TB) *gameLogicHarness {
	t.Helper()
	require.Equal(t, uintptr(56), unsafe.Offsetof(server.Object{}.PosVec))
	require.Equal(t, uintptr(172), unsafe.Offsetof(server.Object{}.Shape))
	require.Equal(t, uintptr(748), unsafe.Offsetof(server.Object{}.UpdateData))
	require.Equal(t, uintptr(276), unsafe.Offsetof(server.PlayerUpdateData{}.Player))
	require.Equal(t, uintptr(2164), unsafe.Offsetof(server.Player{}.GoldVal))
	require.Equal(t, uintptr(1128), unsafe.Offsetof(server.MonsterUpdateData{}.Field282_0))
	oldGameFlags := noxflags.GetGame()
	oldEngineFlags := noxflags.GetEngine()
	oldSettingsUpdated := Nox_server_gameDoSwitchMap_40A680()
	t.Cleanup(func() {
		Nox_server_gameUnsetMapLoad_40A690()
		if oldSettingsUpdated != 0 {
			Nox_server_gameSettingsUpdated_40A670()
		}
	})
	t.Cleanup(func() {
		noxflags.ResetGame()
		noxflags.SetGame(oldGameFlags)
	})
	t.Cleanup(func() {
		noxflags.ResetEngine()
		noxflags.SetEngine(oldEngineFlags)
	})
	h := &gameLogicHarness{t: t}
	t.Cleanup(h.pinner.Unpin)
	return h
}

func (h *gameLogicHarness) pin(v any) {
	h.t.Helper()
	h.pinner.Pin(v)
}

func (h *gameLogicHarness) setGameFlags(flags noxflags.GameFlag) {
	h.t.Helper()
	noxflags.ResetGame()
	noxflags.SetGame(flags)
}

func (h *gameLogicHarness) setEngineFlags(flags noxflags.EngineFlag) {
	h.t.Helper()
	noxflags.ResetEngine()
	noxflags.SetEngine(flags)
}

func (h *gameLogicHarness) loadBlobData() {
	h.t.Helper()
	for _, addr := range []uintptr{0x581450, 0x587000} {
		data := memmap.BlobByAddr(addr).Data
		snapshot := append([]byte(nil), data...)
		h.t.Cleanup(func() { copy(data, snapshot) })
	}
	blobdata.InitData()
}

func (h *gameLogicHarness) snapshotMem(base, off uintptr, size int) {
	h.t.Helper()
	data := memmap.BlobByAddr(base).Data[off : off+uintptr(size)]
	snapshot := append([]byte(nil), data...)
	h.t.Cleanup(func() { copy(data, snapshot) })
}

func (h *gameLogicHarness) setServer(frame, tickRate uint32) *server.Server {
	h.t.Helper()
	oldGetServer := GetServer
	srv := &server.Server{}
	srv.SetFrame(frame)
	srv.SetTickRate(tickRate)
	stub := &gameLogicServer{srv: srv}
	GetServer = func() Server { return stub }
	h.t.Cleanup(func() { GetServer = oldGetServer })
	return srv
}

func (h *gameLogicHarness) object(pos types.Pointf, shape server.Shape) *server.Object {
	h.t.Helper()
	obj := &server.Object{PosVec: pos, Shape: shape}
	h.pin(obj)
	return obj
}

func (h *gameLogicHarness) point(x, y float32) *server.Object {
	h.t.Helper()
	return h.object(types.Ptf(x, y), server.Shape{Kind: server.ShapeKindCenter})
}

func (h *gameLogicHarness) circle(x, y, radius float32) *server.Object {
	h.t.Helper()
	shape := server.Shape{Kind: server.ShapeKindCircle}
	shape.Circle.R = radius
	return h.object(types.Ptf(x, y), shape)
}

func (h *gameLogicHarness) box(x, y, width, height float32) *server.Object {
	h.t.Helper()
	shape := server.Shape{Kind: server.ShapeKindBox}
	shape.Box.W = width
	shape.Box.H = height
	return h.object(types.Ptf(x, y), shape)
}

func (h *gameLogicHarness) player() (*server.Object, *server.PlayerUpdateData, *server.Player) {
	h.t.Helper()
	pl := &server.Player{}
	ud := &server.PlayerUpdateData{Player: pl}
	h.pin(pl)
	h.pin(ud)
	obj := h.point(0, 0)
	obj.ObjClass = object.ClassPlayer
	obj.UpdateData = unsafe.Pointer(ud)
	return obj, ud, pl
}

func (h *gameLogicHarness) monster() (*server.Object, *server.MonsterUpdateData) {
	h.t.Helper()
	ud := &server.MonsterUpdateData{}
	h.pin(ud)
	obj := h.point(0, 0)
	obj.ObjClass = object.ClassMonster
	obj.UpdateData = unsafe.Pointer(ud)
	return obj, ud
}

func TestGameLogicObjectDistance(t *testing.T) {
	h := newGameLogicHarness(t)
	tests := []struct {
		name   string
		first  *server.Object
		second *server.Object
		want   float32
	}{
		{name: "points", first: h.point(0, 0), second: h.point(3, 4), want: 5},
		{name: "first circle", first: h.circle(0, 0, 1), second: h.point(3, 4), want: 4},
		{name: "second circle", first: h.point(0, 0), second: h.circle(3, 4, 1), want: 4},
		{name: "first tall box", first: h.box(0, 0, 2, 6), second: h.point(3, 4), want: 2},
		{name: "first wide box", first: h.box(0, 0, 6, 2), second: h.point(3, 4), want: 2},
		{name: "second tall box", first: h.point(0, 0), second: h.box(3, 4, 2, 6), want: 2},
		{name: "second wide box", first: h.point(0, 0), second: h.box(3, 4, 6, 2), want: 2},
		{name: "overlap clamps to epsilon", first: h.circle(0, 0, 3), second: h.circle(3, 4, 4), want: 0.01},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Nox_xxx_calcDistance_4E6C00(tc.first, tc.second)
			require.InDelta(t, tc.want, got, 0.000001)
		})
	}
}

func TestGameLogicWeaponStamina(t *testing.T) {
	newGameLogicHarness(t)
	tests := []struct {
		name string
		kind int
		want int
	}{
		{name: "default", kind: 0, want: 10},
		{name: "flag 0x200", kind: 0x200, want: 70},
		{name: "flag 0x4000", kind: 0x4000, want: 100},
		{name: "flag 0x800", kind: 0x800, want: 50},
		{name: "flag 0x100", kind: 0x100, want: 45},
		{name: "flag 0x1000", kind: 0x1000, want: 75},
		{name: "flag 0x2000", kind: 0x2000, want: 100},
		{name: "ranged flag range", kind: 0x8000, want: 45},
		{name: "flag 0x400", kind: 0x400, want: 75},
		{name: "precedence", kind: 0x200 | 0x4000, want: 70},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, Nox_xxx_weaponGetStaminaByType_4F7E80(tc.kind))
		})
	}
}

func TestGameLogicIsUnit(t *testing.T) {
	h := newGameLogicHarness(t)
	obj := h.point(0, 0)

	h.setGameFlags(0)
	obj.ObjClass = object.ClassMonster
	obj.ObjSubClass = object.SubClass(0x100)
	require.Equal(t, 1, Nox_xxx_isUnit_4E5B50(obj))

	obj.ObjSubClass = 0
	require.Equal(t, 0, Nox_xxx_isUnit_4E5B50(obj))

	obj.ObjClass = object.ClassPlayer
	obj.ObjSubClass = object.SubClass(0x100)
	require.Equal(t, 0, Nox_xxx_isUnit_4E5B50(obj))

	h.setGameFlags(noxflags.GameOnline)
	obj.ObjClass = object.ClassMonster
	require.Equal(t, 0, Nox_xxx_isUnit_4E5B50(obj))
}

func TestGameLogicDirectionFromVector(t *testing.T) {
	newGameLogicHarness(t)
	tests := []struct {
		name string
		vec  types.Pointf
		want int
	}{
		{name: "east wraps", vec: types.Ptf(1, 0), want: 0},
		{name: "north east", vec: types.Ptf(1, 1), want: 32},
		{name: "north", vec: types.Ptf(0, 1), want: 64},
		{name: "west", vec: types.Ptf(-1, 0), want: 128},
		{name: "south", vec: types.Ptf(0, -1), want: 192},
		{name: "south east", vec: types.Ptf(1, -1), want: 224},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.want, Nox_xxx_math_509ED0(tc.vec))
		})
	}
}

func TestGameLogicServerSubFlags(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setGameFlags(0)
	old := Nox_xxx_getServerSubFlags_409E60()
	t.Cleanup(func() {
		cur := Nox_xxx_getServerSubFlags_409E60()
		Sub_409EC0(int(cur &^ old))
		Sub_409E70(int(old &^ cur))
	})

	const testFlag = uint32(0x40000000)
	Sub_409EC0(int(testFlag))
	Nox_server_gameUnsetMapLoad_40A690()

	Sub_409EC0(int(testFlag))
	require.Zero(t, Nox_server_gameDoSwitchMap_40A680())

	Sub_409E70(int(testFlag))
	require.NotZero(t, Nox_xxx_getServerSubFlags_409E60()&testFlag)
	require.Equal(t, 1, Nox_server_gameDoSwitchMap_40A680())

	Nox_server_gameUnsetMapLoad_40A690()
	Sub_409E70(int(testFlag))
	require.Zero(t, Nox_server_gameDoSwitchMap_40A680())

	Sub_409EC0(int(testFlag))
	require.Zero(t, Nox_xxx_getServerSubFlags_409E60()&testFlag)
	require.Equal(t, 1, Nox_server_gameDoSwitchMap_40A680())
}

func TestGameLogicPlayerLimit(t *testing.T) {
	newGameLogicHarness(t)
	old := Nox_xxx_servGetPlrLimit_409FA0()
	t.Cleanup(func() { Nox_xxx_servSetPlrLimit_409F80(old) })

	limit := 19
	if old == limit {
		limit++
	}
	Nox_xxx_servSetPlrLimit_409F80(limit)
	require.Equal(t, limit, Nox_xxx_servGetPlrLimit_409FA0())
	require.Equal(t, 1, Nox_server_gameDoSwitchMap_40A680())

	Nox_server_gameUnsetMapLoad_40A690()
	Nox_xxx_servSetPlrLimit_409F80(limit)
	require.Zero(t, Nox_server_gameDoSwitchMap_40A680())
}

func TestGameLogicServerName(t *testing.T) {
	newGameLogicHarness(t)
	old := Nox_xxx_serverOptionsGetServername_40A4C0()
	t.Cleanup(func() { Nox_xxx_gameSetServername_40A440(old) })

	name := "HarnessName"
	if strings.EqualFold(old, name) {
		name = "OtherHarness"
	}
	Nox_xxx_gameSetServername_40A440(name)
	require.Equal(t, name, Nox_xxx_serverOptionsGetServername_40A4C0())
	require.Equal(t, 1, Nox_server_gameDoSwitchMap_40A680())

	Nox_server_gameUnsetMapLoad_40A690()
	Nox_xxx_gameSetServername_40A440(strings.ToUpper(name))
	require.Equal(t, name, Nox_xxx_serverOptionsGetServername_40A4C0())
	require.Zero(t, Nox_server_gameDoSwitchMap_40A680())

	Nox_xxx_gameSetServername_40A440("12345678901234567890")
	require.Equal(t, "123456789012345", Nox_xxx_serverOptionsGetServername_40A4C0())
}

func TestGameLogicRespawnRule(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setGameFlags(0)
	old := Nox_server_doPlayersAutoRespawn_40A5F0()
	t.Cleanup(func() {
		h.setGameFlags(0)
		Nox_xxx_ruleSetNoRespawn_40A5E0(old)
	})

	Nox_xxx_ruleSetNoRespawn_40A5E0(1)
	require.Equal(t, 1, Nox_server_doPlayersAutoRespawn_40A5F0())

	h.setGameFlags(noxflags.GameModeQuest)
	require.Zero(t, Nox_server_doPlayersAutoRespawn_40A5F0())
}

func TestGameLogicSysopPassword(t *testing.T) {
	newGameLogicHarness(t)
	old := Nox_xxx_sysopGetPass_40A630()
	t.Cleanup(func() { Nox_xxx_sysopSetPass_40A610(old) })

	Nox_xxx_sysopSetPass_40A610("harness-password")
	require.Equal(t, "harness-password", Nox_xxx_sysopGetPass_40A630())
}

func TestGameLogicPairedSettings(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setGameFlags(0)
	oldMode := Sub_40A220()
	oldValue := Sub_40A6B0()
	t.Cleanup(func() {
		Sub_40A1F0(oldMode)
		Sub_40A6A0(oldValue)
	})

	Sub_40A1F0(37)
	require.Equal(t, 37, Sub_40A220())
	Sub_40A6A0(91)
	require.Equal(t, 91, Sub_40A6B0())
}

func TestGameLogicIndexedDirection(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	tests := []struct {
		direction int
		want      int
	}{
		{direction: 0, want: 5},
		{direction: 32, want: 8},
		{direction: 64, want: 7},
		{direction: 96, want: 6},
		{direction: 128, want: 3},
		{direction: 160, want: 0},
		{direction: 192, want: 1},
		{direction: 224, want: 2},
	}
	for _, tc := range tests {
		require.Equal(t, tc.want, Nox_xxx_math_509EA0(tc.direction), "direction %d", tc.direction)
	}
}

func TestGameLogicQuestStageStorage(t *testing.T) {
	newGameLogicHarness(t)
	oldGameStage := Nox_game_getQuestStage_4E3CC0()
	oldRewardStage := Nox_xxx_getQuestStage_51A930()
	t.Cleanup(func() {
		Nox_game_setQuestStage_4E3CD0(oldGameStage)
		Sub_51A920(oldRewardStage)
	})

	Nox_game_setQuestStage_4E3CD0(37)
	require.Equal(t, 37, Nox_game_getQuestStage_4E3CC0())
	Sub_51A920(41)
	require.Equal(t, 41, Nox_xxx_getQuestStage_51A930())
}

func TestGameLogicConfigurationStorage(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setEngineFlags(0)
	oldLoaded := Sub_4D0D70()
	oldDeathmatch := Sub_4D0DE0(0x100)
	oldQuest := Sub_4D0DE0(0x1000)
	t.Cleanup(func() {
		Sub_4D0D90(oldLoaded)
		Sub_4D0DC0(0x100, oldDeathmatch)
		Sub_4D0DC0(0x1000, oldQuest)
	})

	Sub_4D0D90(0)
	require.Zero(t, Sub_4D0D70())
	Sub_4D0D90(1)
	require.Equal(t, 1, Sub_4D0D70())
	Sub_4D0D90(0)
	h.setEngineFlags(noxflags.EngineNoRendering)
	require.Equal(t, 1, Sub_4D0D70())

	Sub_4D0DC0(0x100, 17)
	require.Equal(t, 17, Sub_4D0DE0(0x100))
	Sub_4D0DC0(0x1000, 23)
	require.Equal(t, 23, Sub_4D0DE0(0x1000))
}

func TestGameLogicGameModeSettings(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	const mode = uint16(noxflags.GameModeArena)
	require.GreaterOrEqual(t, Nox_xxx_servGamedataGet_40A020(mode), 0)
	require.GreaterOrEqual(t, Sub_40A180(noxflags.GameModeArena), 0)
}

func TestGameLogicUnitsHaveSameTeam(t *testing.T) {
	h := newGameLogicHarness(t)
	a := h.point(0, 0)
	b := h.point(0, 0)

	require.False(t, Nox_xxx_unitsHaveSameTeam_4EC520(nil, b))
	require.True(t, Nox_xxx_unitsHaveSameTeam_4EC520(a, a))

	a.TeamVal.ID = server.TeamID(1)
	b.TeamVal.ID = server.TeamID(1)
	require.True(t, Nox_xxx_unitsHaveSameTeam_4EC520(a, b))

	b.TeamVal.ID = server.TeamID(2)
	require.False(t, Nox_xxx_unitsHaveSameTeam_4EC520(a, b))

	ownerA := h.point(0, 0)
	ownerB := h.point(0, 0)
	ownerA.TeamVal.ID = server.TeamID(3)
	ownerB.TeamVal.ID = server.TeamID(3)
	a.ObjOwner = ownerA
	b.ObjOwner = ownerB
	require.True(t, Nox_xxx_unitsHaveSameTeam_4EC520(a, b))
}

func TestGameLogicPlayerMovementRules(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setGameFlags(0)
	obj, ud, _ := h.player()

	require.Equal(t, 1, Nox_xxx_playerCanMove_4F9BC0(obj))
	obj.Buffs = 1 << 25
	require.Zero(t, Nox_xxx_playerCanMove_4F9BC0(obj))
	obj.Buffs = 1 << 5
	require.Zero(t, Nox_xxx_playerCanMove_4F9BC0(obj))
	obj.Buffs = 0

	h.setGameFlags(noxflags.GameModeQuest)
	trade := &server.TradeSession{}
	h.pin(trade)
	ud.Trade70 = trade
	require.Zero(t, Nox_xxx_playerCanMove_4F9BC0(obj))
	ud.Trade70 = nil
	h.setGameFlags(0)

	ud.State = server.PlayerState1
	require.Equal(t, 1, Nox_xxx_playerCanMove_4F9BC0(obj))
	weapon := h.point(0, 0)
	weapon.ObjClass = object.ClassWeapon
	weapon.ObjSubClass = object.SubClass(8)
	ud.Field26 = uint32(uintptr(unsafe.Pointer(weapon)))
	require.Zero(t, Nox_xxx_playerCanMove_4F9BC0(obj))
}

func TestGameLogicPlayerAttackRules(t *testing.T) {
	h := newGameLogicHarness(t)
	obj, ud, _ := h.player()

	ud.State = server.PlayerState0
	require.Equal(t, 1, Nox_xxx_playerCanAttack_4F9C40(obj))
	ud.State = server.PlayerState23
	require.Zero(t, Nox_xxx_playerCanAttack_4F9C40(obj))
	obj.Buffs = 1 << 25
	require.Zero(t, Nox_xxx_playerCanAttack_4F9C40(obj))
}

func TestGameLogicPlayerTargetSlot(t *testing.T) {
	h := newGameLogicHarness(t)
	obj, ud, _ := h.player()

	require.Zero(t, Sub_4F9A80(obj))
	ud.Field42 = 1
	require.Equal(t, 1, Sub_4F9A80(obj))
}

func TestGameLogicStaminaByUnitClass(t *testing.T) {
	h := newGameLogicHarness(t)
	other := h.point(0, 0)
	require.Equal(t, 1, Nox_xxx_playerSubStamina_4F7D30(other, 10))

	player, playerData, _ := h.player()
	playerData.Field22_3 = 9
	require.Zero(t, Nox_xxx_playerSubStamina_4F7D30(player, 10))

	monster, monsterData := h.monster()
	monsterData.Field282_0 = 15
	require.Equal(t, 1, Nox_xxx_playerSubStamina_4F7D30(monster, 10))
	require.Equal(t, uint8(5), monsterData.Field282_0)
	require.Zero(t, Nox_xxx_playerSubStamina_4F7D30(monster, 10))
}

func TestGameLogicGold(t *testing.T) {
	h := newGameLogicHarness(t)
	require.Zero(t, Nox_object_getGold_4FA6D0(nil))

	nonPlayer := h.point(0, 0)
	Nox_object_setGold_4FA620(nonPlayer, 50)
	require.Zero(t, Nox_object_getGold_4FA6D0(nonPlayer))

	player, _, pl := h.player()
	pl.GoldVal = 20
	require.Equal(t, 20, Nox_object_getGold_4FA6D0(player))
	Nox_object_setGold_4FA620(player, 15)
	require.Equal(t, 35, Nox_object_getGold_4FA6D0(player))
	Nox_object_setGold_4FA620(player, -10)
	require.Equal(t, 25, Nox_object_getGold_4FA6D0(player))
	Nox_object_setGold_4FA620(player, -100)
	require.Zero(t, Nox_object_getGold_4FA6D0(player))
}

func TestGameLogicKillableHealth(t *testing.T) {
	h := newGameLogicHarness(t)
	obj := h.point(0, 0)
	require.Zero(t, Nox_xxx_checkIsKillable_528190(obj))

	tests := []struct {
		current uint16
		maximum uint16
		want    int
	}{
		{current: 10, maximum: 10, want: 1},
		{current: 10, maximum: 0, want: 1},
		{current: 0, maximum: 0, want: 1},
		{current: 0, maximum: 10, want: 0},
	}
	for _, tc := range tests {
		health := &server.HealthData{Cur: tc.current, Max: tc.maximum}
		h.pin(health)
		obj.HealthData = health
		require.Equal(t, tc.want, Nox_xxx_checkIsKillable_528190(obj), "health %d/%d", tc.current, tc.maximum)
	}
}

func TestGameLogicSetUnitHealth(t *testing.T) {
	h := newGameLogicHarness(t)
	obj := h.point(0, 0)
	Nox_xxx_unitSetHP_4E4560(obj, 12)

	health := &server.HealthData{Cur: 3, Max: 20}
	h.pin(health)
	obj.HealthData = health
	Nox_xxx_unitSetHP_4E4560(obj, 12)
	require.Equal(t, uint16(12), health.Cur)
	require.Equal(t, ^uint32(0), obj.Field38)
}

func TestGameLogicMonsterCanCast(t *testing.T) {
	h := newGameLogicHarness(t)
	obj, ud := h.monster()
	require.False(t, Nox_xxx_monsterCanCast_534300(obj))
	ud.StatusFlags = 1 << 5
	require.True(t, Nox_xxx_monsterCanCast_534300(obj))
}

func TestGameLogicPlayerActionState(t *testing.T) {
	h := newGameLogicHarness(t)
	obj, ud, pl := h.player()
	tests := []struct {
		state server.PlayerState
		want  int
	}{
		{state: server.PlayerState0, want: 4},
		{state: server.PlayerState2, want: 21},
		{state: server.PlayerState3, want: 1},
		{state: server.PlayerState4, want: 2},
		{state: server.PlayerState5, want: 6},
		{state: server.PlayerState10, want: 21},
		{state: server.PlayerState12, want: 3},
		{state: server.PlayerState13, want: 0},
		{state: server.PlayerState15, want: 40},
		{state: server.PlayerState16, want: 40},
		{state: server.PlayerState17, want: 40},
		{state: server.PlayerState18, want: 48},
		{state: server.PlayerState19, want: 49},
		{state: server.PlayerState20, want: 47},
		{state: server.PlayerState21, want: 30},
		{state: server.PlayerState23, want: 50},
		{state: server.PlayerState24, want: 19},
		{state: server.PlayerStateShakeFist, want: 20},
		{state: server.PlayerStateLaugh, want: 15},
		{state: server.PlayerState27, want: 16},
		{state: server.PlayerStatePoint, want: 16},
		{state: server.PlayerState29, want: 16},
		{state: server.PlayerState30, want: 52},
		{state: server.PlayerState31, want: 0},
		{state: server.PlayerState32, want: 54},
		{state: server.PlayerState33, want: 0},
	}
	for _, tc := range tests {
		ud.State = tc.state
		require.Equal(t, tc.want, Nox_common_mapPlrActionToStateId_4FA2B0(obj), "state %d", tc.state)
	}

	pl.WeaponEquip = 0x400
	ud.State = server.PlayerState13
	require.Equal(t, 0x26, Nox_common_mapPlrActionToStateId_4FA2B0(obj))
}

func TestGameLogicRemoveChildren(t *testing.T) {
	h := newGameLogicHarness(t)
	Nox_xxx_unitRemoveChild_4EC470(nil)

	parent := h.point(0, 0)
	first := h.point(0, 0)
	second := h.point(0, 0)
	parent.Field129 = first
	first.ObjOwner = parent
	first.Field128 = second
	second.ObjOwner = parent

	Nox_xxx_unitRemoveChild_4EC470(parent)
	require.Nil(t, parent.Field129)
	require.Nil(t, first.ObjOwner)
	require.Nil(t, first.Field128)
	require.Nil(t, second.ObjOwner)
	require.Nil(t, second.Field128)
}

func TestGameLogicShadowList(t *testing.T) {
	h := newGameLogicHarness(t)
	C_resetShadowList()
	t.Cleanup(C_resetShadowList)
	obj := h.point(0, 0)
	t.Cleanup(func() { Nox_xxx_action_4DA9F0(obj) })

	Nox_xxx_unitNewAddShadow_4DA9A0(obj)
	require.NotZero(t, obj.ObjFlags&0x10000)
	Nox_xxx_unitNewAddShadow_4DA9A0(obj)
	Nox_xxx_action_4DA9F0(obj)
	require.Zero(t, obj.ObjFlags&0x10000)

	obj.ObjFlags = 0x400000
	Nox_xxx_unitNewAddShadow_4DA9A0(obj)
	require.Zero(t, obj.ObjFlags&0x10000)
}

func TestGameLogicMonsterSoundSet(t *testing.T) {
	h := newGameLogicHarness(t)
	require.Nil(t, Nox_xxx_monsterGetSoundSet_424300(nil))
	require.Nil(t, Nox_xxx_monsterGetSoundSet_424300(h.point(0, 0)))

	obj, ud := h.monster()
	soundSet := new(byte)
	h.pin(soundSet)
	ud.SoundSet122 = unsafe.Pointer(soundSet)
	require.Equal(t, unsafe.Pointer(soundSet), Nox_xxx_monsterGetSoundSet_424300(obj))
}

func TestGameLogicPointDirectionLookup(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	directions := []int16{0, 32, 64, 96, 128, 160, 192, 224}
	points := []types.Pointf{
		types.Ptf(1, 0),
		types.Ptf(1, 1),
		types.Ptf(0, 1),
		types.Ptf(-1, 1),
		types.Ptf(-1, 0),
		types.Ptf(-1, -1),
		types.Ptf(0, -1),
		types.Ptf(1, -1),
	}
	want := [8][8]int{
		{1, 9, 8, 10, 2, 6, 4, 5},
		{5, 1, 9, 8, 10, 2, 6, 4},
		{4, 5, 1, 9, 8, 10, 2, 6},
		{6, 4, 5, 1, 9, 8, 10, 2},
		{2, 6, 4, 5, 1, 9, 8, 10},
		{10, 2, 6, 4, 5, 1, 9, 8},
		{8, 10, 2, 6, 4, 5, 1, 9},
		{9, 8, 10, 2, 6, 4, 5, 1},
	}
	for i, direction := range directions {
		for j, point := range points {
			require.Equal(t, want[i][j],
				Nox_server_testTwoPointsAndDirection_4E6E50(types.Ptf(0, 0), direction, point),
				"direction %d point %v", direction, point)
		}
	}
}

func TestGameLogicGoldAdjustments(t *testing.T) {
	h := newGameLogicHarness(t)
	obj, _, pl := h.player()
	pl.GoldVal = 20

	Nox_xxx_playerAddGold_4FA590(obj, 15)
	require.Equal(t, uint32(35), pl.GoldVal)
	Nox_xxx_playerSubGold_4FA5D0(obj, 10)
	require.Equal(t, uint32(25), pl.GoldVal)
	Nox_xxx_playerSubGold_4FA5D0(obj, 100)
	require.Zero(t, pl.GoldVal)
}

func TestGameLogicConfusedDirection(t *testing.T) {
	h := newGameLogicHarness(t)
	srv := h.setServer(0, 30)
	obj := h.point(0, 0)
	obj.Direction2 = 5

	require.Equal(t, server.Dir16(231), Nox_xxx_playerConfusedGetDirection_4F7A40(obj))
	srv.SetFrame(30)
	require.Equal(t, server.Dir16(5), Nox_xxx_playerConfusedGetDirection_4F7A40(obj))
	srv.SetFrame(20)
	obj.Direction2 = 250
	obj.BuffsPower[3] = 10
	require.Equal(t, server.Dir16(124), Nox_xxx_playerConfusedGetDirection_4F7A40(obj))
}

func TestGameLogicMoveAttemptTime(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setServer(100, 30)
	obj, ud := h.monster()

	ud.Field127 = 50
	require.Equal(t, 1, Nox_xxx_mobGetMoveAttemptTime_534810(obj))
	ud.Field127 = 0
	require.Zero(t, Nox_xxx_mobGetMoveAttemptTime_534810(obj))
}

func TestGameLogicPlayerInteractionStatus(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setServer(77, 30)
	source, _, sourcePlayer := h.player()
	target, _, targetPlayer := h.player()
	sourcePlayer.PlayerInd = 7

	Sub_4E7540(nil, target)
	Sub_4E7540(source, nil)
	Sub_4E7540(source, source)
	require.Zero(t, targetPlayer.Field3600)

	Sub_4E7540(source, target)
	require.Equal(t, uint32(1), targetPlayer.Field3600)
	require.Equal(t, uint32(7), targetPlayer.Field3604)
}

func TestGameLogicPlayerCounterReset(t *testing.T) {
	h := newGameLogicHarness(t)
	h.loadBlobData()
	h.setServer(0, 30)
	const index = 3
	const off = uintptr(1565124 + 12*index)
	h.snapshotMem(0x5D4594, off, 12)

	Nox_xxx_playerResetImportantCtr_4E4F40(ntype.PlayerInd(index))
	require.Equal(t, uint8(1), *memmap.PtrUint8(0x5D4594, off))
	require.Equal(t, uint8(2), *memmap.PtrUint8(0x5D4594, off+1))
	require.Equal(t, uint32(60), *memmap.PtrUint32(0x5D4594, off+4))
	require.Zero(t, *memmap.PtrUint32(0x5D4594, off+8))
}

func TestGameLogicMouseConfiguration(t *testing.T) {
	newGameLogicHarness(t)
	Sub_430AA0(0)
	t.Cleanup(func() {
		Sub_430AA0(0)
		require.Equal(t, 5, Nox_xxx_cursor_430B00())
	})

	tests := []struct {
		setting int
		primary int
		cursor  int
	}{
		{setting: 1, primary: 1, cursor: 9},
		{setting: 2, primary: 2, cursor: 13},
		{setting: 0, primary: 0, cursor: 5},
		{setting: 99, primary: 0, cursor: 5},
	}
	for _, tc := range tests {
		Sub_430AA0(tc.setting)
		require.Equal(t, tc.primary, Nox_client_mousePriKey_430AF0())
		require.Equal(t, tc.cursor, Nox_xxx_cursor_430B00())
	}
}

func TestGameLogicClientStateAccessors(t *testing.T) {
	newGameLogicHarness(t)

	oldCollision := Sub_42EBA0()
	t.Cleanup(func() { Sub_42EB90(oldCollision) })
	Sub_42EB90(37)
	require.Equal(t, 37, Sub_42EBA0())

	oldNoxWorld := Sub_43AF80()
	t.Cleanup(func() { Sub_43AF90(oldNoxWorld) })
	Sub_43AF90(41)
	require.Equal(t, 41, Sub_43AF80())

	oldServerOptions := Sub_459D60()
	t.Cleanup(func() { Sub_459D50(oldServerOptions) })
	Sub_459D50(43)
	require.Equal(t, 43, Sub_459D60())

	oldOverlay := Sub_47A260()
	t.Cleanup(func() { Set_dword_5d4594_1123520(oldOverlay) })
	Set_dword_5d4594_1123520(47)
	require.Equal(t, 47, Sub_47A260())

	require.Zero(t, Sub_459DA0())
}

func TestGameLogicNormalizeMapRectangle(t *testing.T) {
	h := newGameLogicHarness(t)
	tests := []struct {
		name string
		in   [8]int32
		want [4]int32
	}{
		{
			name: "orders mixed coordinates",
			in:   [8]int32{1: -10, 2: 7000, 4: -20, 7: 6000},
			want: [4]int32{7000, 6000, -20, -10},
		},
		{
			name: "clamps lower maximums",
			in:   [8]int32{1: -5, 2: -10, 4: -20, 7: -15},
			want: [4]int32{0, 0, -10, -5},
		},
		{
			name: "clamps upper minimums",
			in:   [8]int32{1: 8000, 2: 7000, 4: 6000, 7: 6500},
			want: [4]int32{6000, 6500, 5887, 5887},
		},
		{
			name: "reverse ordering",
			in:   [8]int32{1: 100, 2: 10, 4: 50, 7: 20},
			want: [4]int32{10, 20, 50, 100},
		},
	}
	for _, tc := range tests {
		in := tc.in
		var got [4]int32
		h.pin(&in)
		h.pin(&got)
		Sub_428170(unsafe.Pointer(&in), unsafe.Pointer(&got))
		require.Equal(t, tc.want, got, tc.name)
	}
}

func TestGameLogicRaiseUnit(t *testing.T) {
	h := newGameLogicHarness(t)
	obj := h.point(0, 0)
	obj.ObjClass = object.ClassPlayer
	for i := range obj.Field140 {
		obj.Field140[i] = 0x123
	}

	Nox_xxx_unitRaise_4E46F0(obj, 0)
	require.Zero(t, obj.ZVal)
	Nox_xxx_unitRaise_4E46F0(obj, 7.5)
	require.Equal(t, float32(7.5), obj.ZVal)
	require.Equal(t, ^uint32(0), obj.Field38)
	for _, v := range obj.Field140 {
		require.Equal(t, uint32(0x400000), v)
	}
}

func TestGameLogicObjectOnOff(t *testing.T) {
	h := newGameLogicHarness(t)
	obj := h.point(0, 0)
	obj.ObjClass = object.Class(0x20400005)

	Nox_xxx_objectSetOn_4E75B0(obj)
	require.NotZero(t, obj.ObjFlags&0x1000000)
	require.Equal(t, ^uint32(0), obj.Field38)
	Nox_xxx_objectSetOn_4E75B0(obj)

	Nox_xxx_objectSetOff_4E7600(obj)
	require.Zero(t, obj.ObjFlags&0x1000000)
}

func TestGameLogicAdjustUnitHealth(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setGameFlags(0)
	obj := h.point(0, 0)
	health := &server.HealthData{Cur: 5, Max: 10}
	h.pin(health)
	obj.HealthData = health

	Nox_xxx_unitAdjustHP_4EE460(obj, 3)
	require.Equal(t, uint16(8), health.Cur)
	Nox_xxx_unitAdjustHP_4EE460(obj, 20)
	require.Equal(t, uint16(10), health.Cur)
	Nox_xxx_unitAdjustHP_4EE460(obj, 1)
	require.Equal(t, uint16(10), health.Cur)

	health.Cur = 1
	h.setGameFlags(noxflags.GameFlag(0x4000000))
	Nox_xxx_unitAdjustHP_4EE460(obj, 3)
	require.Equal(t, uint16(1), health.Cur)
}

func TestGameLogicNonlethalDamage(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setEngineFlags(0)
	Nox_xxx_unitDamageClear_4EE5E0(nil, 3)

	obj := h.point(0, 0)
	Nox_xxx_unitDamageClear_4EE5E0(obj, 3)
	health := &server.HealthData{Cur: 10, Max: 20}
	h.pin(health)
	obj.HealthData = health
	Nox_xxx_unitDamageClear_4EE5E0(obj, 3)
	require.Equal(t, uint16(7), health.Cur)

	player, _, _ := h.player()
	playerHealth := &server.HealthData{Cur: 10, Max: 20}
	h.pin(playerHealth)
	player.HealthData = playerHealth
	h.setEngineFlags(noxflags.EngineGodMode)
	Nox_xxx_unitDamageClear_4EE5E0(player, 3)
	require.Equal(t, uint16(10), playerHealth.Cur)
}

func TestGameLogicNPCAnimationLocalAction(t *testing.T) {
	h := newGameLogicHarness(t)
	obj, ud := h.monster()
	obj.ObjSubClass = object.SubClass(0x10)
	ud.Field120_3 = 1
	ud.AIStackInd = 0
	ud.AIStack[0].Action = 16

	Nox_xxx_updateNPCAnimData_50A850(obj)
	require.Zero(t, ud.Field120_3)
}

func TestGameLogicMapPointBounds(t *testing.T) {
	newGameLogicHarness(t)
	require.Equal(t, -1, Nox_xxx_tileNFromPoint_411160(types.Ptf(0, 0)))
	require.Equal(t, -1, Nox_xxx_tileNFromPoint_411160(types.Ptf(6000, 6000)))
}

func TestGameLogicFrameThreshold(t *testing.T) {
	h := newGameLogicHarness(t)
	h.setServer(1000, 30)
	h.snapshotMem(0x5D4594, 3520, 4)
	*memmap.PtrUint32(0x5D4594, 3520) = 500
	require.Zero(t, Sub_40AA00())
	*memmap.PtrUint32(0x5D4594, 3520) = 300
	require.Equal(t, 1, Sub_40AA00())

	h.snapshotMem(0x5D4594, 3476, 4)
	*memmap.PtrUint32(0x5D4594, 3476) = 53
	require.Equal(t, 53, Sub_40AA40())
}

func TestGameLogicPaletteConversion(t *testing.T) {
	h := newGameLogicHarness(t)
	source := make([]byte, 256*3)
	converted := make([]byte, 256*4)
	for i := 0; i < 256; i++ {
		source[3*i+0] = byte(i)
		source[3*i+1] = byte(255 - i)
		source[3*i+2] = byte(i ^ 0x55)
	}
	h.pin(&source[0])
	h.pin(&converted[0])
	Sub_435120(unsafe.Pointer(&converted[0]), unsafe.Pointer(&source[0]))
	for i := 0; i < 256; i++ {
		require.Equal(t, source[3*i:3*i+3], converted[4*i:4*i+3])
		require.Equal(t, byte(4), converted[4*i+3])
	}

	roundTrip := make([]byte, 256*3)
	h.pin(&roundTrip[0])
	Sub_435150(unsafe.Pointer(&roundTrip[0]), unsafe.Pointer(&converted[0]))
	require.Equal(t, source, roundTrip)
}

func TestGameLogicTileRowInterpolation(t *testing.T) {
	h := newGameLogicHarness(t)
	h.snapshotMem(0x852978, 8, 4)
	input := [7]byte{6: 2}
	h.pin(&input)
	*memmap.PtrPtr(0x852978, 8) = nil
	require.Equal(t, 57, Sub_476080(unsafe.Pointer(&input[0])))

	target := [5]int32{}
	h.pin(&target)
	*memmap.PtrPtr(0x852978, 8) = unsafe.Pointer(&target[0])

	input = [7]byte{0, 0, 0, 0, 0, 0, 0}
	require.Equal(t, 22, Sub_476080(unsafe.Pointer(&input[0])))
	input[0] = 3
	require.Equal(t, 22, Sub_476080(unsafe.Pointer(&input[0])))
	input[0] = 11
	require.Equal(t, 22, Sub_476080(unsafe.Pointer(&input[0])))

	input[0] = 1
	require.Zero(t, Sub_476080(unsafe.Pointer(&input[0])))
	input[0] = 4
	require.Zero(t, Sub_476080(unsafe.Pointer(&input[0])))
	input[0] = 12
	require.Zero(t, Sub_476080(unsafe.Pointer(&input[0])))

	input[0] = 2
	require.Equal(t, 11, Sub_476080(unsafe.Pointer(&input[0])))
	target[3] = 100
	input[0] = 1
	require.Equal(t, 22, Sub_476080(unsafe.Pointer(&input[0])))
}

func TestGameLogicFixedBufferReset(t *testing.T) {
	h := newGameLogicHarness(t)
	buf := make([]byte, 0x7f8)
	for i := range buf {
		buf[i] = 0xff
	}
	h.pin(&buf[0])
	Sub_57B920(unsafe.Pointer(&buf[0]))
	require.Equal(t, make([]byte, len(buf)), buf)
}

func TestGameLogicGauntletStateAccessors(t *testing.T) {
	newGameLogicHarness(t)
	oldGate := Sub_4D7430()
	t.Cleanup(func() { Sub_4D7440(oldGate) })
	Sub_4D7440(59)
	require.Equal(t, 59, Sub_4D7430())

	oldWarp := Sub_4D76F0()
	t.Cleanup(func() { Sub_4D76E0(oldWarp) })
	Sub_4D76E0(61)
	require.Equal(t, 61, Sub_4D76F0())
}

func TestGameLogicMemmapResets(t *testing.T) {
	h := newGameLogicHarness(t)
	h.snapshotMem(0x5D4594, 1548428, 48)
	config := memmap.BlobByAddr(0x5D4594).Data[1548428 : 1548428+48]
	for i := range config {
		config[i] = 0xff
	}
	Sub_4D0DA0()
	require.Equal(t, make([]byte, len(config)), config)

	h.snapshotMem(0x5D4594, 1565524, 64)
	important := memmap.BlobByAddr(0x5D4594).Data[1565524 : 1565524+64]
	for i := range important {
		important[i] = 0xff
	}
	Sub_4E4ED0()
	require.Equal(t, make([]byte, len(important)), important)
}
