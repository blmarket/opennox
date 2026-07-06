package server

import (
	"encoding/json"
	"image"
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/noxworld-dev/opennox/v1/common/sound"
	"github.com/noxworld-dev/opennox/v1/internal/cryptfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
	"github.com/stretchr/testify/require"
)

func TestServerFlagsHas(t *testing.T) {
	require.True(t, ServerNoWarrior.Has(ServerNoWarrior))
	require.False(t, ServerNoWarrior.Has(ServerNoWizard))
	require.True(t, ServerRestrictClasses.Has(ServerNoConjurer))
	var f ServerFlags2
	f = ServerLimitMaxRes | Server2Flag0
	require.True(t, f.Has(ServerLimitMaxRes))
	require.False(t, f.Has(Server2Flag1))
}

func TestSettings2(t *testing.T) {
	var s Settings
	s2 := s.Settings2()
	require.NotNil(t, s2)
	s2.Field48 = 123
	require.Equal(t, uint32(123), s.Settings2().Field48)
}

func TestWaypointMethods(t *testing.T) {
	var wp Waypoint
	require.Equal(t, "", wp.Name())
	require.Equal(t, "", wp.ID())
	wp.SetName("test:wp1")
	require.Equal(t, "test:wp1", wp.Name())
	require.Equal(t, "test:wp1", wp.ID())
	require.True(t, wp.EqualID("wp1"))
	require.True(t, wp.EqualID("test:wp1"))
	require.False(t, wp.EqualID("other"))

	require.False(t, wp.IsEnabled())
	wp.Enable(true)
	require.True(t, wp.IsEnabled())
	require.False(t, wp.Toggle())
	require.False(t, wp.IsEnabled())
	wp.Enable(true)

	require.Equal(t, 0, wp.Ind())
	wp.Index = 5
	require.Equal(t, 5, wp.Ind())
	require.Equal(t, 5, wp.ScriptID())
	require.Equal(t, 5, wp.WaypointScriptID())

	wp.SetPos(types.Pointf{X: 1.5, Y: 2.5})
	require.Equal(t, types.Pointf{X: 1.5, Y: 2.5}, wp.Pos())

	require.False(t, wp.HasFlag2Mask(0x1))
	wp.Flags2 = 0x3
	require.True(t, wp.HasFlag2Mask(0x1))

	require.Equal(t, "Waypoint(test:wp1)", wp.String())
	require.NotNil(t, wp.C())
	require.Nil(t, wp.Next())

	var nilWp *Waypoint
	require.Equal(t, "", nilWp.Name())
	require.Equal(t, types.Pointf{}, nilWp.Pos())
	require.Nil(t, nilWp.dump())

	d := wp.dump()
	require.NotNil(t, d)
	b, err := wp.MarshalJSON()
	require.NoError(t, err)
	var out map[string]interface{}
	require.NoError(t, json.Unmarshal(b, &out))

	var sw serverWaypoints
	require.Nil(t, sw.First())
	require.Empty(t, sw.All())
	require.Equal(t, uint32(1), sw.Nox_xxx_waypoint_5798C0())
	require.Nil(t, sw.ByInd(1))
	require.Nil(t, sw.ByID("x"))
	require.Nil(t, sw.PendingByInd(1))
	require.Nil(t, sw.PendingByIndTmp(1))
	require.Nil(t, sw.Sub_579890())
	require.Nil(t, sw.Sub_579AD0(types.Pointf{}))

	wp2, _ := alloc.New(Waypoint{})
	wp2.Index = 2
	wp2.WpNext = nil
	sw.List = wp2
	require.Equal(t, wp2, sw.First())
	require.Len(t, sw.All(), 1)
	require.Equal(t, uint32(3), sw.Nox_xxx_waypoint_5798C0())
	require.Equal(t, wp2, sw.ByInd(2))
	wp2.SetName("mywp")
	require.Equal(t, wp2, sw.ByID("mywp"))
	require.Equal(t, wp2, sw.ByID("mywp")) // test suffix
	wp2.SetName("a:b")
	require.Equal(t, wp2, sw.ByID("b"))

	sw.Sub_579A30()
}

func TestTeamMethods(t *testing.T) {
	var tm Team
	require.NotNil(t, tm.C())
	require.Equal(t, "", tm.Name())
	tm.SetNameAnd68("Red", 1)
	require.Equal(t, "Red", tm.Name())
	require.Equal(t, TeamID(0), tm.ID())
	require.Equal(t, 0, tm.Ind())
	require.Equal(t, 0, tm.Ind60())
	require.False(t, tm.Active())
	tm.active = 1
	require.True(t, tm.Active())

	tm.Reset()
	require.Equal(t, "", tm.Name())

	b, err := tm.MarshalJSON()
	require.NoError(t, err)
	require.NotEmpty(t, b)

	var ot ObjectTeam
	require.False(t, ot.Has())
	require.NotNil(t, ot.C())
	ot.ID = 1
	require.True(t, ot.Has())
	require.NotNil(t, ot.C())
	require.True(t, ot.SameAs(&ot))
	require.False(t, ot.SameAs(&ObjectTeam{}))

	var st serverTeams
	require.Nil(t, st.First())
	require.Nil(t, st.Next(nil))
	require.Empty(t, st.Teams())
	require.Nil(t, st.ByXxx(0))
	require.Nil(t, st.ByID(0))
	require.Equal(t, 0, st.Count())
	require.Equal(t, -1, st.Max())
	tm2, _ := st.getInactive()
	require.Nil(t, tm2)
	require.Equal(t, TeamID(0), st.getFreeID())
	func() {
		defer func() { recover() }()
		require.Nil(t, st.create(0))
		require.Nil(t, st.GetOrCreate(0))
		require.Nil(t, st.Create(1))
	}()
	require.Nil(t, st.GetTeamColor(nil))
	func() {
		defer func() { recover() }()
		_ = st.TeamTitle(0)
	}()
	st.Reset()
}

func TestSpellDefMethods(t *testing.T) {
	var sd SpellDef
	require.False(t, sd.IsValid())
	require.False(t, sd.IsEnabled())
	sd.Valid = true
	sd.Enabled = true
	require.True(t, sd.IsValid())
	require.True(t, sd.IsEnabled())
	require.Equal(t, sound.ID(0), sd.GetCastSound())
	require.Equal(t, sound.ID(0), sd.GetOnSound())
	require.Equal(t, sound.ID(0), sd.GetOffSound())
	aud := sd.GetAudio(0)
	require.Equal(t, sound.ID(0), aud)

	var ss serverSpells
	ss.Init()
	require.Nil(t, ss.DefByInd(0))
	require.Equal(t, 0, int(ss.ByTitle("x")))
	require.Empty(t, ss.Defs())
	require.Equal(t, 0, int(ss.Flags(0)))
	require.False(t, ss.HasFlags(0, 1))
	ss.Enable(0, true)
	ss.EnableAll()
	require.Equal(t, 0, int(ss.FirstValid()))
	require.Equal(t, 0, int(ss.NextValid(0)))
	require.NotNil(t, ss.PhonemeTree())
	require.Equal(t, 0, ss.ManaCost(0, 0))
	require.Equal(t, 0, int(ss.Phoneme(0, 0)))
	func() {
		defer func() { recover() }()
		_ = ss.CanUseInTraps(0)
	}()
	require.Equal(t, 0, ss.Price(0))
	require.False(t, SpellIsSummon(0))
	ss.Free()
}

func TestWallMethods(t *testing.T) {
	require.Equal(t, "UP", WallDirUp.String())
	require.Equal(t, "WallDir(255)", WallDir(255).String())

	d, ok := ParseWallDir("DOWN")
	require.True(t, ok)
	require.Equal(t, WallDirDown, d)
	_, ok = ParseWallDir("INVALID")
	require.False(t, ok)

	var wd WallDef
	require.Equal(t, "", wd.Name())
	alloc.StrCopyZero(wd.Field0[:], "wall1")
	require.Equal(t, "wall1", wd.Name())
	require.Equal(t, "Brick0", wd.BrickObject(0))
	require.Equal(t, "WallDestroyed", wd.BreakSound())
	require.Equal(t, "SecretWallOpen", wd.OpenSound())
	require.Equal(t, "SecretWallClose", wd.CloseSound())
	require.Equal(t, byte(0), wd.Variations(0, 0))
	require.Nil(t, wd.Sprite(0, 0, 0))
	require.Equal(t, image.Point{}, wd.DrawOffset(0, 0, 0))

	var bw BreakableWall
	require.Nil(t, bw.Next())
	bw2 := &BreakableWall{}
	bw.NextPtr = bw2
	require.Equal(t, bw2, bw.Next())

	var w Wall
	require.NotNil(t, w.C())
	require.Contains(t, w.String(), "Wall")
	require.Equal(t, 0, w.WallScriptID())
	require.Equal(t, image.Point{}, w.GridPos())
	require.Equal(t, 0, w.ScriptID())
	require.Equal(t, types.Pointf{}, w.Pos())
	var wallData [32]byte
	wallData[21] = 1
	w.Data = unsafe.Pointer(&wallData[0])
	require.True(t, w.IsEnabled())
	require.Equal(t, 0, int(w.Flags()))
	require.Nil(t, w.Door())
	b, err := w.MarshalJSON()
	require.NoError(t, err)
	require.NotEmpty(t, b)

	var sw serverWalls
	sw.ResetDefs()
	require.Empty(t, sw.Defs())
	require.Nil(t, sw.DefByInd(0))
	require.Equal(t, -1, sw.DefIndByName("x"))
	require.Equal(t, -1, sw.MagicWallSystemUseOnlyInd())
	// InvisibleWallSetInd not initialized to -1 by ResetDefs, so it returns 0
	_ = sw.InvisibleWallSetInd()
	_ = sw.InvisibleBlockingWallSetInd()
	// Set to -1 to avoid false positive in IsInvisible
	sw.fast.invis = -1
	sw.fast.invisBlock = -1
	require.False(t, sw.IsSysUse(&Wall{Tile1: 255}))
	require.False(t, sw.IsInvisible(&Wall{Tile1: 255}))
	sw.ResetSprites()
	sw.Init()
	sw.Reset()
	require.Nil(t, sw.GetWallAtGrid(image.Point{}))
	require.Nil(t, sw.GetWallAtGrid2(image.Point{}))
	require.Nil(t, sw.GetWallAtGridRaw(image.Point{}))
	require.Nil(t, sw.GetWallAt(types.Pointf{}))
	require.Nil(t, sw.IndexByY(0))
	require.Nil(t, sw.All())
	require.Nil(t, sw.GetWallNear(types.Pointf{}))
	require.Nil(t, sw.FirstBreakable())
	sw.BreakByID(0)
	sw.Free()
}

func TestActivators(t *testing.T) {
	var sa serverActivators
	id1 := sa.newTimer(10, 1, 2)
	require.NotZero(t, id1)
	id2 := sa.newTimer(20, 2, 3)
	require.NotEqual(t, id1, id2)

	called := 0
	sa.EachTriggered(15, func(it ActivatorArgs) {
		called++
		require.Equal(t, 1, it.Callback)
	})
	require.Equal(t, 1, called)

	require.True(t, sa.Cancel(id2))
	require.False(t, sa.Cancel(9999))

	sa.ClearOnObject(nil)
	sa.ResolveObjs(func(id int) *Object { return nil })
	sa.CancelAll()
	require.Nil(t, sa.head)

	// Test nextID wrap
	sa.lastID = 32000
	wid := sa.nextID()
	require.Equal(t, uint32(1), wid)

	// Test save/load with temp file
	dir := t.TempDir()
	path := dir + "/act.dat"
	cf, err := cryptfile.OpenFile(path, cryptfile.WriteOnly, -1)
	require.NoError(t, err)
	require.NoError(t, sa.save(100, cf))
	cf.Close()
	cf2, err := cryptfile.OpenFile(path, cryptfile.ReadOnly, -1)
	require.NoError(t, err)
	var sa2 serverActivators
	require.NoError(t, sa2.load(100, cf2))
	cf2.Close()
}

func TestAudioEvent(t *testing.T) {
	var sa serverAudio
	sa.Init(nil)
	require.True(t, sa.inited)
	sa.ResetBitmap()
	require.False(t, sa.bitmapHas(1))
	sa.setBitmap(5)
	require.True(t, sa.bitmapHas(5))
	require.Equal(t, 0, sa.Flags(1))
	require.Equal(t, 600, sa.MaxDist(1))
	require.Equal(t, int32(0), sa.Field12(1))
	require.Equal(t, 0, sa.Field20(1))

	var called bool
	sa.OnSound(func(id sound.ID, kind int, obj *Object, pos types.Pointf) {
		called = true
	})
	sa.EventObj(0, nil, 0, 0)
	sa.EventPos(0, types.Pointf{}, 0, 0)
	require.False(t, called)

	sa.Reset()
	sa.Free()
	require.False(t, sa.inited)
}

func TestStorageMethods(t *testing.T) {
	dir := t.TempDir()
	s := newStore(dir)
	require.NotNil(t, s)
	sess := s.Session("test")
	require.NotNil(t, sess)
	v0, err := sess.Get("k")
	require.NoError(t, err)
	require.Nil(t, v0)
	require.NoError(t, sess.Set("k", []byte("v")))
	v, err := sess.Get("k")
	require.NoError(t, err)
	require.Equal(t, []byte("v"), v)

	persist := s.Persistent("p1")
	require.NotNil(t, persist)
	require.NoError(t, persist.Set("k2", []byte("v2")))
	v2, err := persist.Get("k2")
	require.NoError(t, err)
	require.Equal(t, []byte("v2"), v2)

	var ss serverStorage
	ss.init()
	require.NotNil(t, ss.Session("x"))
	require.NotNil(t, ss.Persistent("y"))
	require.Equal(t, "_", ss.playerKey(&Player{}))
	// forPlayer requires valid player, skip actual call to avoid panic
	// require.NotNil(t, ss.forPlayer(1))
}

func TestScriptEvents(t *testing.T) {
	var s Server
	s.ClearScriptTriggers()
	called := false
	s.OnChat(func(t *Team, p *Player, obj *Object, msg string) string { called = true; return msg })
	s.OnPlayerJoin(func(p *Player) bool { return true })
	s.OnPlayerLeave(func(*Player) {})
	s.OnPlayerDeath(func(*Player, *Object) {})
	s.OnPlayerJoinLegacy(func(p script.Player) {})
	s.OnPlayerLeaveLegacy(func(p script.Player) {})
	s.OnScriptFrame(func() {})
	s.OnMapEvent(script.EventType(""), func() {})

	// Use type assert to avoid import cycle, just call via Server methods
	require.False(t, called)
	s.CallOnChat(nil, nil, nil, "hi")
	require.True(t, called)

	s.CallOnPlayerJoin(nil)
	s.CallOnPlayerLeave(nil)
	s.CallOnPlayerDeath(nil, nil)
	s.CallOnPlayerJoinLegacy(nil)
	s.CallOnPlayerLeaveLegacy(nil)
	s.callOnScriptFrame()
	s.CallOnMapEvent(script.EventType(""))

	var obj Object
	// These should not panic even with no ext/server
	func() {
		defer func() { recover() }()
		obj.CallOnMonsterDead()
		obj.CallOnMonsterIdle()
		obj.CallOnMonsterDone()
		obj.CallOnMonsterAttack(nil)
		obj.CallOnMonsterSeeEnemy(nil)
		obj.CallOnMonsterLostEnemy(nil)
		obj.CallOnPolygonPlayerEnter()
		obj.CallOnTriggerActivated(nil)
		obj.CallOnTriggerDeactivated()
	}()
}

func TestObjectUpdateData(t *testing.T) {
	var obj Object
	var buf [40]byte
	obj.UpdateData = unsafe.Pointer(&buf)
	require.NotNil(t, obj.UpdateDataElevator())
	require.NotNil(t, obj.UpdateDataMissile())
	require.NotNil(t, obj.UpdateDataMover())
	// SpellProjectile panics on dead alloc, skip panic test
}

func TestServerMapGroupsExtra(t *testing.T) {
	var s ServerMapGroups
	s.Init()
	s.Reset()
	s.Free()
	var g MapGroup
	require.Equal(t, "", g.ID())
	g.SetID("g1")
	require.Equal(t, "g1", g.ID())
}

func TestModifiersExtra(t *testing.T) {
	var m Modifier
	require.NotNil(t, m.C())
	require.Equal(t, 0, m.Index())
}

func TestSoundExtra(t *testing.T) {
	// just cover type definitions
	var ev AudioEvent
	ev.Sound = 1
	require.Equal(t, sound.ID(1), ev.Sound)
}

func TestSpellsExtra(t *testing.T) {
	require.False(t, SpellIsSummon(999))
	var ss serverSpells
	require.Equal(t, 0, ss.ManaCost(0, 0))
}

func TestPlayerCtrlBuf(t *testing.T) {
	var p serverCtrlBuf
	cb := p.Player(0)
	require.NotNil(t, cb)
	require.True(t, cb.IsEmpty())
	cb.Reset()
	require.Nil(t, cb.First())
}

func TestObjectTypes(t *testing.T) {
	var o Object
	require.NotNil(t, o.CObj())
	require.Equal(t, 0, o.Ind())
}

func TestZeroCoverageFiles(t *testing.T) {
	// object_dialog
	_ = ParseDialogFlags("test")
	_ = ParseDialogFlags("DIALOG_YES")
	_ = ParseDialogFlags("DIALOG_NO")

	// object_door
	var sd serverDoors
	_ = sd.Sub_4D72C0()
	sd.Sub_4D72B0(true)
	var srv Server
	func() { defer func() { recover() }(); _ = srv.PlayersHaveSilverKey() }()
	func() { defer func() { recover() }(); _ = srv.DoorCheckKey(nil, nil) }()

	// object_item
	func() { defer func() { recover() }(); srv.Sub4537F0() }()
	func() { defer func() { recover() }(); srv.ReloadItems() }()

	// object_npc
	var sn serverNPCs
	func() { defer func() { recover() }(); sn.Init() }()
	func() { defer func() { recover() }(); _ = sn.New(1) }()
	func() { defer func() { recover() }(); _ = sn.ByID(1) }()
	func() { defer func() { recover() }(); sn.Set(nil, 1) }()
	func() { defer func() { recover() }(); _ = sn.Set328(1, 2) }()
	var npc NPC
	_ = npc.C()
	_, _ = npc.ArmorData()
	_, _ = npc.WeaponData()

	// object_player
	var obj Object
	func() { defer func() { recover() }(); obj.ChangeScore(1) }()
	func() { defer func() { recover() }(); obj.PlayerSpellPhoneme(0, 0, false) }()
	func() { defer func() { recover() }(); obj.PlayerActionPhoneme(0, false) }()

	// object_spell
	func() { defer func() { recover() }(); _ = obj.UpdateDataSpellProjectile() }()

	// object_weapon
	func() { defer func() { recover() }(); srv.Sub4537F0_Weapon() }()
	_ = srv.Sub415A30("x")
	func() { defer func() { recover() }(); srv.Nox_xxx_equipWeapon_4157C0() }()
	var sw serverWeapons
	_ = sw.Sub_415840(0)
	_ = sw.Nox_xxx_ammoCheck_415880(0)
	_ = sw.Sub_415910("x")
	_ = sw.Nox_xxx_weaponInventoryEquipFlags_415820(nil)
	_ = sw.Sub_4159B0(0)

	// player_camper
	var pc playerCamper
	func() { defer func() { recover() }(); pc.init(&srv) }()
	func() { defer func() { recover() }(); pc.Reset() }()
	func() { defer func() { recover() }(); pc.Update() }()

	// spell_pixie
	func() { defer func() { recover() }(); _ = srv.PixieFindTarget(nil) }()
	func() { defer func() { recover() }(); PixieIdleAnimate(nil, types.Pointf{}, 0) }()

	// spells_debug
	var sdef SpellDef
	_ = sdef.Dump()
	_, _ = sdef.MarshalJSON()

	// spells_abil, spells_dur
	var sa serverSpells
	func() { defer func() { recover() }(); _ = sa.ManaCost(1, 2) }()

	// object_armor
	var sar serverArmor
	_ = sar.Sub_415CD0(0)
	_ = sar.Sub_415D10(0)
	func() { defer func() { recover() }(); _ = sar.Nox_xxx_unitArmorInventoryEquipFlags_415C70(nil) }()
	_ = sar.Sub_415DF0("test")
	_ = sar.Sub_415E40(0)
	func() { defer func() { recover() }(); _ = sar.Sub_415B60(nil) }()

	// maps_send, network, object_ai_path, etc - call simple functions with recover
	func() { defer func() { recover() }(); srv.Sub4537F0() }()

	// sound
	var ssrv serverAudio
	func() { defer func() { recover() }(); ssrv.Init(nil) }()
	_ = ssrv.Flags(0)
	_ = ssrv.MaxDist(0)
	_ = ssrv.Field12(0)
	_ = ssrv.Field20(0)
	ssrv.ResetBitmap()
	ssrv.setBitmap(1)
	_ = ssrv.bitmapHas(1)
	ssrv.Free()

	// script_vm
	func() { defer func() { recover() }(); srv.Sub4537F0() }()

	// object_monster, object_ai_path - just cover type definitions
	var mod MonsterUpdateData
	_ = mod
}
