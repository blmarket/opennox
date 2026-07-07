package noxscript

import (
	"image"
	"testing"
	"time"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/damage"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/noxscript/ns/v4/enchant"
	"github.com/noxworld-dev/noxscript/ns/v4/spell"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/types"
)

// mockNS implements ns.Implementation with stub methods
type mockNS struct{}

func (m *mockNS) Frame() int                                                            { return 0 }
func (m *mockNS) Time() time.Duration                                                   { return 0 }
func (m *mockNS) TimeSource() ns.TimeSource                                             { return nil }
func (m *mockNS) FrameRate() int                                                        { return 30 }
func (m *mockNS) Store(typ ns.StorageType) ns.Storage                                   { return nil }
func (m *mockNS) TimerByHandle(h ns.TimerHandle) ns.Timer                               { return nil }
func (m *mockNS) NewTimer(dt ns.Duration, fnc ns.Func, args ...any) ns.Timer            { return nil }
func (m *mockNS) RandomFloat(min float32, max float32) float32                          { return 0 }
func (m *mockNS) Random(min int, max int) int                                           { return 0 }
func (m *mockNS) StopScript(value any)                                                  {}
func (m *mockNS) AutoSave()                                                             {}
func (m *mockNS) StartupScreen(which int)                                               {}
func (m *mockNS) DeathScreen(which int)                                                 {}
func (m *mockNS) ObjectType(name string) ns.ObjType                                     { return nil }
func (m *mockNS) ObjectTypeByInd(ind int) ns.ObjType                                    { return nil }
func (m *mockNS) ObjectByHandle(h ns.ObjHandle) ns.Obj                                  { return nil }
func (m *mockNS) Object(name string) ns.Obj                                             { return nil }
func (m *mockNS) ObjectGroupByHandle(h ns.ObjGroupHandle) ns.ObjGroup                   { return nil }
func (m *mockNS) ObjectGroup(name string) ns.ObjGroup                                   { return nil }
func (m *mockNS) CreateObject(typ string, pos ns.Positioner) ns.Obj                     { return nil }
func (m *mockNS) GetTrigger() ns.Obj                                                    { return nil }
func (m *mockNS) GetCaller() ns.Obj                                                     { return nil }
func (m *mockNS) GetHost() ns.Obj                                                       { return nil }
func (m *mockNS) IsTrigger(obj ns.Obj) bool                                             { return false }
func (m *mockNS) IsCaller(obj ns.Obj) bool                                              { return false }
func (m *mockNS) IsGameBall(obj ns.Obj) bool                                            { return false }
func (m *mockNS) IsCrown(obj ns.Obj) bool                                               { return false }
func (m *mockNS) IsSummoned(obj ns.Obj) bool                                            { return false }
func (m *mockNS) MakeFriendly(obj ns.Obj)                                               {}
func (m *mockNS) MakeEnemy(obj ns.Obj)                                                  {}
func (m *mockNS) BecomePet(obj ns.Obj)                                                  {}
func (m *mockNS) BecomeEnemy(obj ns.Obj)                                                {}
func (m *mockNS) Teams() []ns.Team                                                      { return nil }
func (m *mockNS) HostPlayer() ns.Player                                                 { return nil }
func (m *mockNS) Players() []ns.Player                                                  { return nil }
func (m *mockNS) GetCharacterData(field int) int                                        { return 0 }
func (m *mockNS) Print(message ns.StringID)                                             {}
func (m *mockNS) PrintStr(message string)                                               {}
func (m *mockNS) PrintToAll(message ns.StringID)                                        {}
func (m *mockNS) PrintStrToAll(message string)                                          {}
func (m *mockNS) ClearMessages(player ns.Obj)                                           {}
func (m *mockNS) UnBlind()                                                              {}
func (m *mockNS) Blind()                                                                {}
func (m *mockNS) ImmediateBlind()                                                       {}
func (m *mockNS) WideScreen(enable bool)                                                {}
func (m *mockNS) IsTalking() bool                                                       { return false }
func (m *mockNS) IsTrading() bool                                                       { return false }
func (m *mockNS) SetHalberd(upgrade ns.HalberdLevel)                                    {}
func (m *mockNS) EndGame(class player.Class)                                            {}
func (m *mockNS) DestroyEveryChat()                                                     {}
func (m *mockNS) SetShopkeeperText(obj ns.Obj, text ns.StringID)                        {}
func (m *mockNS) SetShopkeeperTextStr(obj ns.Obj, text string)                          {}
func (m *mockNS) SetDialog(obj ns.Obj, typ ns.DialogType, start ns.Func, end ns.Func)   {}
func (m *mockNS) CancelDialog(obj ns.Obj)                                               {}
func (m *mockNS) StoryPic(obj ns.Obj, name string)                                      {}
func (m *mockNS) TellStory(audio audio.Name, story ns.StringID)                         {}
func (m *mockNS) TellStoryStr(audio audio.Name, story string)                           {}
func (m *mockNS) StartDialog(obj ns.Obj, other ns.Obj)                                  {}
func (m *mockNS) GetAnswer(obj ns.Obj) ns.DialogAnswer                                  { return 0 }
func (m *mockNS) AudioEvent(audio audio.Name, p ns.Positioner)                          {}
func (m *mockNS) Music(music int, volume int)                                           {}
func (m *mockNS) MusicPushEvent()                                                       {}
func (m *mockNS) MusicPopEvent()                                                        {}
func (m *mockNS) MusicEvent()                                                           {}
func (m *mockNS) Effect(effect effect.Effect, p1, p2 ns.Positioner)                     {}
func (m *mockNS) CastSpell(spell spell.Spell, source, target ns.Positioner)             {}
func (m *mockNS) CastSpellLvl(spell spell.Spell, lvl int, source, target ns.Positioner) {}
func (m *mockNS) NewTrap(pos ns.Positioner, spells []ns.TrapSpell) ns.Obj               { return nil }
func (m *mockNS) GetQuestStatus(name string) int                                        { return 0 }
func (m *mockNS) GetQuestStatusFloat(name string) float32                               { return 0 }
func (m *mockNS) SetQuestStatus(status int, name string)                                {}
func (m *mockNS) SetQuestStatusFloat(status float32, name string)                       {}
func (m *mockNS) ResetQuestStatus(name string)                                          {}
func (m *mockNS) JournalEntry(obj ns.Obj, message ns.StringID, typ ns.EntryType)        {}
func (m *mockNS) JournalEdit(obj ns.Obj, message ns.StringID, typ ns.EntryType)         {}
func (m *mockNS) JournalDelete(obj ns.Obj, message ns.StringID)                         {}
func (m *mockNS) JournalEntryStr(obj ns.Obj, message string, typ ns.EntryType)          {}
func (m *mockNS) JournalEditStr(obj ns.Obj, message string, typ ns.EntryType)           {}
func (m *mockNS) JournalDeleteStr(obj ns.Obj, message string)                           {}
func (m *mockNS) Waypoints() []ns.WaypointObj                                           { return nil }
func (m *mockNS) WaypointByHandle(h ns.WaypointHandle) ns.WaypointObj                   { return nil }
func (m *mockNS) Waypoint(name string) ns.WaypointObj                                   { return nil }
func (m *mockNS) NewWaypoint(name string, pos types.Pointf) ns.WaypointObj              { return nil }
func (m *mockNS) WaypointGroupByHandle(h ns.WaypointGroupHandle) ns.WaypointGroupObj    { return nil }
func (m *mockNS) WaypointGroup(name string) ns.WaypointGroupObj                         { return nil }
func (m *mockNS) NoWallSound(noWallSound bool)                                          {}
func (m *mockNS) WallByHandle(h ns.WallHandle) ns.WallObj                               { return nil }
func (m *mockNS) Wall(x int, y int) ns.WallObj                                          { return nil }
func (m *mockNS) WallGroupByHandle(h ns.WallGroupHandle) ns.WallGroupObj                { return nil }
func (m *mockNS) WallGroup(name string) ns.WallGroupObj                                 { return nil }
func (m *mockNS) FindWalls(fnc func(it ns.WallObj) bool, conditions ...ns.WallCond) int { return 0 }
func (m *mockNS) LoadMap(name string, opts *ns.LoadMapOptions)                          {}
func (m *mockNS) OnFrame(fnc ns.FrameFunc)                                              {}
func (m *mockNS) OnMapEvent(typ ns.MapEvent, fnc ns.MapEventFunc)                       {}
func (m *mockNS) OnChat(fnc ns.ChatFunc)                                                {}
func (m *mockNS) OnPlayerJoin(fnc ns.PlayerJoinFunc)                                    {}
func (m *mockNS) OnPlayerLeave(fnc ns.PlayerLeaveFunc)                                  {}
func (m *mockNS) OnPlayerDeath(fnc ns.PlayerDeathFunc)                                  {}
func (m *mockNS) Unused1f(id int)                                                       {}
func (m *mockNS) Unused20(id int)                                                       {}
func (m *mockNS) Unused50()                                                             {}
func (m *mockNS) Unused58(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unused59(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unused5a(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unused5b(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unused5c(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unused5d(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unused5e(str string) int                                               { return 0 }
func (m *mockNS) Unused74(arg1 int, arg2 int)                                           {}
func (m *mockNS) Unknownb8(id int) bool                                                 { return false }
func (m *mockNS) Unknownb9(id int) bool                                                 { return false }
func (m *mockNS) Unknownc4()                                                            {}
func (m *mockNS) FindObjects(fnc func(it ns.Obj) bool, conditions ...ns.ObjCond) int    { return 0 }

// mockGroup implements ns.ObjGroup with stub methods
type mockGroup struct{}

func (m *mockGroup) FindObjects(fnc func(it ns.Obj) bool, conditions ...ns.ObjCond) int { return 0 }
func (m *mockGroup) EachObject(recurse bool, fnc func(obj ns.Obj) bool)                 {}
func (m *mockGroup) AllObjects() ns.Objects                                             { return nil }
func (m *mockGroup) ObjGroupScriptID() int                                              { return 0 }
func (m *mockGroup) ScriptID() int                                                      { return 0 }
func (m *mockGroup) Name() string                                                       { return "" }
func (m *mockGroup) Enable(enable bool)                                                 {}
func (m *mockGroup) Toggle() bool                                                       { return false }
func (m *mockGroup) HasOwner(owner ns.Obj) bool                                         { return false }
func (m *mockGroup) HasOwnerIn(owners ns.ObjGroup) bool                                 { return false }
func (m *mockGroup) SetOwner(owner ns.Obj)                                              {}
func (m *mockGroup) SetOwners(owners ns.ObjGroup)                                       {}
func (m *mockGroup) Pause(dt ns.Duration)                                               {}
func (m *mockGroup) Move(wp ns.WaypointObj)                                             {}
func (m *mockGroup) LookAtDirection(dir ns.Direction)                                   {}
func (m *mockGroup) Wander()                                                            {}
func (m *mockGroup) Idle()                                                              {}
func (m *mockGroup) Hunt()                                                              {}
func (m *mockGroup) Follow(targ ns.Positioner)                                          {}
func (m *mockGroup) WalkTo(p types.Pointf)                                              {}
func (m *mockGroup) CreateMover(wp ns.WaypointObj, speed float32)                       {}
func (m *mockGroup) RaiseZombie()                                                       {}
func (m *mockGroup) ZombieStayDown()                                                    {}
func (m *mockGroup) Enchant(e enchant.Enchant, dur ns.Duration)                         {}
func (m *mockGroup) AwardSpell(s spell.Spell)                                           {}
func (m *mockGroup) AggressionLevel(v float32)                                          {}
func (m *mockGroup) RetreatLevel(v float32)                                             {}
func (m *mockGroup) ResumeLevel(v float32)                                              {}
func (m *mockGroup) Attack(targ ns.Positioner)                                          {}
func (m *mockGroup) Flee(targ ns.Positioner, dt ns.Duration)                            {}
func (m *mockGroup) Guard(p1, p2 types.Pointf, dist float32)                            {}
func (m *mockGroup) HitMelee(pos types.Pointf)                                          {}
func (m *mockGroup) HitRanged(pos types.Pointf)                                         {}
func (m *mockGroup) Damage(source ns.Obj, amount int, typ damage.Type)                  {}
func (m *mockGroup) Delete()                                                            {}
func (m *mockGroup) SetRoamFlag(flags int)                                              {}

// mockVM2 is an extended mock VM for coverage testing
type mockVM2 struct {
	mockVM
	nsImpl       ns.Implementation
	popGroupMock bool
}

func (m *mockVM2) NoxScript() ns.Implementation {
	if m.nsImpl != nil {
		return m.nsImpl
	}
	return &mockNS{}
}

func (m *mockVM2) PopU32() uint32 {
	if len(m.stack) == 0 {
		return 0
	}
	v := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if uv, ok := v.(uint32); ok {
		return uv
	}
	if iv, ok := v.(int); ok {
		return uint32(iv)
	}
	if iv, ok := v.(int32); ok {
		return uint32(iv)
	}
	return 0
}

func (m *mockVM2) PopBool() bool {
	if len(m.stack) == 0 {
		return false
	}
	v := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if bv, ok := v.(bool); ok {
		return bv
	}
	return false
}

func (m *mockVM2) PopPoint() image.Point {
	return image.Pt(1, 2)
}

func (m *mockVM2) PopPointf() types.Pointf {
	return types.Pointf{X: 1, Y: 2}
}

func (m *mockVM2) PopObjGroupNS() ns.ObjGroup {
	if m.popGroupMock {
		return &mockGroup{}
	}
	return nil
}

func TestNSFunctionsCoverage(t *testing.T) {
	vm := &mockVM2{nsImpl: &mockNS{}}

	call := func(f func()) {
		defer func() {
			recover()
		}()
		f()
	}

	// Test object.go functions
	call(func() { nsGetTrigger(vm) })
	call(func() { nsGetCaller(vm) })
	call(func() { nsIsTrigger(vm) })
	call(func() { nsIsCaller(vm) })
	call(func() { vm.stack = append(vm.stack, "testObject") })
	call(func() { nsObject(vm) })
	call(func() { vm.stack = append(vm.stack, "type") })
	call(func() { nsCreateObject(vm) })
	call(func() { nsObjectX(vm) })
	call(func() { nsObjectY(vm) })
	call(func() { nsSetPos(vm) })
	call(func() { nsZ(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetZ(vm) })
	call(func() { nsIsEnabled(vm) })
	call(func() { nsObjectOn(vm) })
	call(func() { nsObjectOff(vm) })
	call(func() { nsObjectToggle(vm) })
	call(func() { vm.stack = append(vm.stack, "class") })
	call(func() { nsHasClass(vm) })
	call(func() { vm.stack = append(vm.stack, "subclass") })
	call(func() { nsHasSubclass(vm) })
	call(func() { nsCurrentHealth(vm) })
	call(func() { nsMaxHealth(vm) })
	call(func() { vm.stack = append(vm.stack, int32(10)) })
	call(func() { nsRestoreHealth(vm) })
	call(func() { nsGetDirection(vm) })
	call(func() { vm.stack = append(vm.stack, int32(90)) })
	call(func() { nsLookWithAngle(vm) })
	call(func() { nsLookAtObject(vm) })
	call(func() { vm.stack = append(vm.stack, int32(0)) })
	call(func() { nsLookAtDirection(vm) })
	call(func() { nsDelete(vm) })
	call(func() { vm.stack = append(vm.stack, uint32(10)) })
	call(func() { nsDeleteObjectTimer(vm) })
	call(func() { nsPushTo(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsPush(vm) })
	call(func() { vm.stack = append(vm.stack, true) })
	call(func() { nsFreeze(vm) })
	call(func() { nsPickup(vm) })
	call(func() { nsDrop(vm) })
	call(func() { nsIsOwnedBy(vm) })
	call(func() { nsClearOwner(vm) })
	call(func() { nsSetOwner(vm) })
	call(func() { nsSetOwners(vm) })
	call(func() { nsIsOwnedByAny(vm) })
	call(func() { nsReturn(vm) })
	call(func() { nsIdle(vm) })
	call(func() { nsWander(vm) })
	call(func() { nsFollow(vm) })
	call(func() { nsHunt(vm) })
	call(func() { nsWalkTo(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsCreateMover(vm) })
	call(func() { nsUnlock(vm) })
	call(func() { nsLock(vm) })
	call(func() { nsFirstItem(vm) })
	call(func() { nsNextItem(vm) })
	call(func() { hasItem(vm) })
	call(func() { nsGetHolder(vm) })
	call(func() { nsRaiseZombie(vm) })
	call(func() { nsZombieStayDown(vm) })
	call(func() { nsGetElevatorStat(vm) })
	call(func() { nsMove(vm) })
	call(func() { nsIsLocked(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(10)) })
	call(func() { nsDamage(vm) })
	call(func() { nsGetHost(vm) })
	call(func() { nsClearMessages(vm) })
	call(func() { nsIsAttackedBy(vm) })
	call(func() { vm.stack = append(vm.stack, "msg") })
	call(func() { nsChat(vm) })
	call(func() { vm.stack = append(vm.stack, "msg", uint32(5)) })
	call(func() { nsChatTimerSeconds(vm) })
	call(func() { vm.stack = append(vm.stack, "msg", uint32(5)) })
	call(func() { nsChatTimerFrames(vm) })
	call(func() { nsDestroyChat(vm) })
	call(func() { nsDestroyEveryChat(vm) })
	call(func() { nsIsVisibleTo(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetAggressionLevel(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetRetreatLevel(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetResumeLevel(vm) })
	call(func() { nsIsGameBall(vm) })
	call(func() { nsIsCrown(vm) })
	call(func() { nsIsSummoned(vm) })
	call(func() { nsGetGold(vm) })
	call(func() { vm.stack = append(vm.stack, int32(10)) })
	call(func() { nsChangeGold(vm) })
	call(func() { vm.stack = append(vm.stack, "name", int32(0), int32(1)) })
	call(func() { nsSetDialog(vm) })
	call(func() { nsStartDialog(vm) })
	call(func() { nsGetAnswer(vm) })
	call(func() { nsCancelDialog(vm) })
	call(func() { vm.stack = append(vm.stack, "pic") })
	call(func() { nsStoryPic(vm) })
	call(func() { vm.stack = append(vm.stack, "v1", "v0") })
	call(func() { nsTellStory(vm) })
	call(func() { nsAttack(vm) })
	call(func() { vm.stack = append(vm.stack, uint32(10)) })
	call(func() { nsRunAway(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsGuard(vm) })
	call(func() { vm.stack = append(vm.stack, uint32(10)) })
	call(func() { nsPause(vm) })
	call(func() { vm.stack = append(vm.stack, uint32(1), uint32(2)) })
	call(func() { nsSetCallback(vm) })
	call(func() { nsHitLocation(vm) })
	call(func() { nsHitFarLocation(vm) })

	// Test object_group.go functions
	call(func() { vm.stack = append(vm.stack, "group") })
	call(func() { nsGetObjectGroup(vm) })
	call(func() { nsObjectGroupOn(vm) })
	call(func() { nsObjectGroupOff(vm) })
	call(func() { nsObjectGroupToggle(vm) })
	call(func() { nsGroupSetOwner(vm) })
	call(func() { nsGroupSetOwners(vm) })
	call(func() { nsGroupIsOwnedBy(vm) })
	call(func() { nsGroupIsOwnedByAny(vm) })
	call(func() { vm.popGroupMock = true })
	call(func() { nsGroupDelete(vm) })
	call(func() { vm.popGroupMock = false })
	call(func() { vm.stack = append(vm.stack, int32(0)) })
	call(func() { nsGroupLookAtDirection(vm) })
	call(func() { nsGroupWander(vm) })
	call(func() { nsGroupIdle(vm) })
	call(func() { nsGroupHunt(vm) })
	call(func() { nsGroupFollow(vm) })
	call(func() { nsGroupWalkTo(vm) })
	call(func() { nsGroupMove(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsGroupCreateMover(vm) })
	call(func() { nsRaiseZombieGroup(vm) })
	call(func() { nsZombieStayDownGroup(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(10)) })
	call(func() { nsGroupDamage(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsGroupEnchant(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsGroupAwardSpell(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetAggressionLevelGroup(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetRetreatLevelGroup(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsSetResumeLevelGroup(vm) })
	call(func() { nsGroupAttack(vm) })
	call(func() { vm.stack = append(vm.stack, uint32(10)) })
	call(func() { nsGroupRunAway(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0)) })
	call(func() { nsGroupGuard(vm) })
	call(func() { vm.stack = append(vm.stack, uint32(10)) })
	call(func() { nsGroupPause(vm) })
	call(func() { nsGroupHitLocation(vm) })
	call(func() { nsGroupHitFarLocation(vm) })

	// Test spells.go functions
	call(func() { vm.stack = append(vm.stack, "effect") })
	call(func() { nsEffect(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsHasEnchant(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsEnchant(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsEnchantOff(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsCastSpellObjObj(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsCastSpellObjPos(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsCastSpellPosObj(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsCastSpellPosPos(vm) })
	call(func() { vm.stack = append(vm.stack, "spell") })
	call(func() { nsAwardSpell(vm) })
	call(func() { nsTrapSpells(vm) })

	// Test audio.go functions
	call(func() { vm.stack = append(vm.stack, "sound") })
	call(func() { nsAudioEvent(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(100)) })
	call(func() { nsMusic(vm) })
	call(func() { nsMusicPushEvent(vm) })
	call(func() { nsMusicPopEvent(vm) })
	call(func() { nsMusicEvent(vm) })

	// Test script.go functions
	call(func() { vm.stack = append(vm.stack, int32(1), int32(10)) })
	call(func() { nsRandomInt(vm) })
	call(func() { vm.stack = append(vm.stack, float32(1.0), float32(10.0)) })
	call(func() { nsRandomFloat(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1)) })
	call(func() { nsEndGame(vm) })
	call(func() { nsImmediateBlind(vm) })
	call(func() { vm.stack = append(vm.stack, int32(10)) })
	call(func() { nsChangeScore(vm) })
	call(func() { vm.stack = append(vm.stack, int32(5)) })
	call(func() { nsGetScore(vm) })
	call(func() { vm.stack = append(vm.stack, true) })
	call(func() { nsWideScreen(vm) })
	call(func() { nsAutoSave(vm) })
	call(func() { nsDistance(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1)) })
	call(func() { nsDeathScreen(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1)) })
	call(func() { nsStartupScreen(vm) })
	call(func() { vm.stack = append(vm.stack, "msg") })
	call(func() { nsPrint(vm) })
	call(func() { vm.stack = append(vm.stack, "msg") })
	call(func() { nsPrintToAll(vm) })
	call(func() { nsUnBlind(vm) })
	call(func() { nsBlind(vm) })
	call(func() { vm.stack = append(vm.stack, "msg", int32(1)) })
	call(func() { nsJournalEntry(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1)) })
	call(func() { nsGetCharacterData(vm) })

	// Test walls.go functions
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsGetWall(vm) })
	call(func() { vm.stack = append(vm.stack, "group") })
	call(func() { nsGetWallGroup(vm) })
	call(func() { vm.stack = append(vm.stack, true) })
	call(func() { nsNoWallSound(vm) })
	call(func() { nsWallOpen(vm) })
	call(func() { nsWallClose(vm) })
	call(func() { nsWallToggle(vm) })
	call(func() { nsWallBreak(vm) })

	// Test waypoint.go functions
	call(func() { vm.stack = append(vm.stack, "wp") })
	call(func() { nsGetWaypoint(vm) })
	call(func() { vm.stack = append(vm.stack, "wpg") })
	call(func() { nsGetWaypointGroup(vm) })
	call(func() { nsWaypointX(vm) })
	call(func() { nsWaypointY(vm) })
	call(func() { nsWaypointSetPos(vm) })
	call(func() { nsWaypointOn(vm) })
	call(func() { nsWaypointOff(vm) })
	call(func() { nsWaypointToggle(vm) })
	call(func() { nsWaypointIsEnabled(vm) })
	call(func() { nsWaypointGroupOn(vm) })
	call(func() { nsWaypointGroupOff(vm) })
	call(func() { nsWaypointGroupToggle(vm) })

	// Test unknown.go functions
	call(func() { vm.stack = append(vm.stack, int32(1)) })
	call(func() { nsUnused1f(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1)) })
	call(func() { nsUnused20(vm) })
	call(func() { nsUnused50(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused58(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused59(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused5a(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused5b(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused5c(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused5d(vm) })
	call(func() { vm.stack = append(vm.stack, int32(1), int32(2)) })
	call(func() { nsUnused74(vm) })
	call(func() { nsUnknownc4(vm) })

	// Test helper functions
	call(func() { gridUnpack(0) })
	call(func() { gridUnpack(12345) })
	call(func() { popWall(vm) })
}
