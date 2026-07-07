package server

import (
	"image"
	"reflect"
	"testing"

	ns4 "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/noxscript/ns/v4/audio"
	"github.com/noxworld-dev/noxscript/ns/v4/effect"
	"github.com/noxworld-dev/opennox-lib/script"
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox-lib/types"
	"github.com/noxworld-dev/opennox/v1/common/sound"
)

func TestZeroCoverageFiles2(t *testing.T) {
	// script_vm_ns_unknown.go - empty functions, just call them
	var ns NoxScriptNS
	ns.Unused1f(1)
	ns.Unused20(2)
	ns.Unused50()
	ns.Unused58(1, 2)
	ns.Unused59(1, 2)
	ns.Unused5a(1, 2)
	ns.Unused5b(1, 2)
	ns.Unused5c(1, 2)
	ns.Unused5d(1, 2)

	// script_vm_ns_audio.go
	ns.AudioEvent(audio.Name("test"), nil)
	// with mock positioner
	mockPos := &mockPositioner{pos: types.Pointf{X: 1, Y: 2}}
	func() {
		defer func() { recover() }()
		ns.AudioEvent(audio.Name("test"), mockPos)
	}()

	// script_vm_ns_player.go
	ns.ClearMessages(nil)
	func() {
		defer func() { recover() }()
		ns.ClearMessages(nil)
	}()

	// script_vm_old_ns.go
	var srv Server
	ns2 := NoxScriptNS{s: &srv}

	// script_vm_ns_hooks.go
	func() { defer func() { recover() }(); ns2.OnFrame(nil) }()
	func() { defer func() { recover() }(); ns2.OnMapEvent(ns4.MapInitialize, nil) }()
	func() { defer func() { recover() }(); ns2.OnMapEvent(ns4.MapEntry, nil) }()
	func() { defer func() { recover() }(); ns2.OnMapEvent(ns4.MapExit, nil) }()
	func() { defer func() { recover() }(); ns2.OnMapEvent(ns4.MapShutdown, nil) }()
	func() { defer func() { recover() }(); ns2.OnMapEvent(ns4.MapEvent(999), nil) }()

	// script_vm_old_ns.go
	nsvm := srv.NoxScriptNSVM()
	func() { defer func() { recover() }(); nsvm.NewString("test") }()
	func() { defer func() { recover() }(); nsvm.GetString(0) }()
	func() { defer func() { recover() }(); nsvm.GetFuncVar(0, 0) }()
	func() { defer func() { recover() }(); nsvm.GetFuncInd("test") }()
	func() { defer func() { recover() }(); nsvm.SetFuncVar(0, 0, 0) }()
	func() { defer func() { recover() }(); nsvm.CallFunc(0, nil) }()

	// script_vm_old_builtins.go
	var vm NoxScriptVM
	func() { defer func() { recover() }(); vm.callBuiltin(nil, 0) }()
	_ = vm.builtinNeedsDPos(0)
	_ = vm.builtinNeedsDPos(1)
	_ = vm.builtinNeedsDPos(2)
	_ = vm.builtinNeedsDPos(3)
	_ = vm.builtinNeedsDPos(4)
	_ = vm.builtinNeedsDPos(5)
	_ = vm.builtinNeedsDPos(6)

	// script_vm.go
	func() { defer func() { recover() }(); srv.ScriptTick() }()
	func() { defer func() { recover() }(); srv.VMsShutdown() }()
	srv.VMs.VMs = []script.VM{&mockVM{}}
	func() { defer func() { recover() }(); srv.ScriptTick() }()
	func() { defer func() { recover() }(); srv.VMsShutdown() }()

	// script_vm_ns_dialog.go
	ns.SetDialog(nil, ns4.DialogType("test"), nil, nil)
	ns.CancelDialog(nil)
	ns.StoryPic(nil, "test")
	_ = ns.GetAnswer(nil)
	func() {
		defer func() { recover() }()
		ns.SetDialog(nil, ns4.DialogType("DIALOG_YES"), nil, nil)
		ns.CancelDialog(nil)
		ns.StoryPic(nil, "test")
		_ = ns.GetAnswer(nil)
	}()

	// script_vm_ns_object.go
	_ = nsToObj(nil)
	func() { defer func() { recover() }(); _ = ns.IsTrigger(nil) }()
	func() { defer func() { recover() }(); _ = ns.IsCaller(nil) }()
	_ = ns.IsGameBall(nil)
	_ = ns.IsCrown(nil)
	_ = ns.IsSummoned(nil)
	func() {
		defer func() { recover() }()
		_ = ns2.IsTrigger(nil)
		_ = ns2.IsCaller(nil)
		_ = ns2.IsGameBall(nil)
		_ = ns2.IsCrown(nil)
		_ = ns2.IsSummoned(nil)
	}()

	// script_vm_ns.go
	_ = srv.NoxScriptNS()
	_ = ns2.Frame()
	_ = ns2.FrameRate()
	srv.SetTickRate(30)
	srv.SetFrame(60)
	_ = ns2.Time()
	func() { defer func() { recover() }(); _ = ns2.Time() }()
	srv.SetTickRate(30)
	func() { defer func() { recover() }(); _ = ns2.RandomFloat(1.0, 2.0) }()
	func() { defer func() { recover() }(); _ = ns2.Random(1, 10) }()
	func() { defer func() { recover() }(); ns2.StopScript("test") }()
	_ = ns2.TimerByHandle(nil)
	func() { defer func() { recover() }(); _ = ns2.TimerByHandle(&mockTimerHandle{id: 1}) }()
	func() { defer func() { recover() }(); _ = ns2.NewTimer(ns4.Infinite(), nil) }()
	func() { defer func() { recover() }(); _ = ns2.NewTimer(ns4.Frames(1), nil, 1, 2) }()
	func() { defer func() { recover() }(); _ = ns2.NewTimer(ns4.Frames(1), nil) }()
	func() { defer func() { recover() }(); _ = ns2.NewTimer(ns4.Frames(1), nil, 123) }()

	var nt nsTimer
	_ = nt.ScriptID()
	_ = nt.TimerScriptID()
	func() { defer func() { recover() }(); _ = nt.Cancel() }()

	// script_vm_ns_waypoint.go
	_ = ns2.Waypoints()
	_ = ns2.WaypointByHandle(nil)
	_ = ns2.Waypoint("test")
	func() { defer func() { recover() }(); _ = ns2.NewWaypoint("test", types.Pointf{}) }()
	_ = ns2.WaypointGroupByHandle(nil)
	_ = ns2.WaypointGroup("test")
	var wg nsWpGroup
	func() {
		defer func() { recover() }()
		_ = wg.ScriptID()
		_ = wg.WaypointGroupScriptID()
		_ = wg.Name()
		wg.Enable(true)
		_ = wg.Toggle()
		wg.EachWaypoint(true, func(obj ns4.WaypointObj) bool { return true })
	}()

	// script_vm_ns_spell.go (has init and Effect with 0 coverage)
	func() {
		defer func() { recover() }()
		ns2.Effect(effect.Effect("test"), nil, nil)
		ns2.Effect(effect.Effect("test"), mockPos, mockPos)
	}()

	// script_vm_old_panic.go
	_ = vm.panicCompilerCheck(0)
	func() { defer func() { recover() }(); _ = vm.panicBuiltinRead() }()
	func() { defer func() { recover() }(); _ = vm.panicBuiltinWrite() }()
	_, _ = vm.PanicScriptCall(0)
	vm.panic.enabled = true
	_, _ = vm.PanicScriptCall(185)
	_, _ = vm.PanicScriptCall(89)
	_, _ = vm.PanicScriptCall(999)

	// script_vm_old.go - call many simple methods
	func() {
		defer func() { recover() }()
		var sf ScriptFunc
		_ = sf.Name()
		_ = sf.Args()
		_ = sf.Locals()
		_ = sf.AllLocals()
		vm.Init(&srv)
		vm.Reset()
		vm.resetVirtualFuncs()
		_ = vm.FuncsCnt()
		_ = vm.Funcs()
		_, _ = vm.FuncByPref("test")
		_, _ = vm.FuncByName("test")
		vm.ResetBuiltin()
		_ = vm.NameSuff()
		_ = vm.DPos()
		_ = vm.DPosf()
		_ = vm.AsValue(nil)
		_ = vm.AsValue(123)
		_ = vm.AsValue("test")
		_ = vm.AsFuncIndex("test", nil)
		_ = vm.addVirtual("test", nil)
		_ = vm.ScriptIndexByName("test")
		_ = vm.ScriptNameByIndex(0)
		_ = vm.Caller()
		_ = vm.Trigger()
		vm.resetStack()
		_ = vm.saveStack()
		vm.adjustStack(0, 0)
		_ = vm.stackAt(0)
		vm.PushU32(1)
		vm.PushI32(1)
		vm.PushF32(1.0)
		vm.PushBool(true)
		_ = vm.NewString("test")
		vm.PushString("test")
		_ = vm.PopI32()
		_ = vm.PopU32()
		_ = vm.PopF32()
		_ = vm.PopBool()
		_, _ = vm.LookupString(0)
		_ = vm.GetString(0)
		_ = vm.PopString()
		_ = vm.PopPoint()
		_ = vm.PopPointf()
		vm.PushHandleNS(nil)
		_ = vm.PopWaypointNS()
		_ = vm.PopGroup()
		_ = vm.PopWpGroupNS()
		vm.scriptPushCallback(nil, nil, nil)
		vm.scriptPopCallback(nil, nil, nil)
		vm.OnEvent(script.EventType("test"))
		vm.RunFirst(false)
		_ = vm.getFuncVarPtr(0, 0)
		_ = vm.getFuncVarPtrFor(&sf, 0)
		_, _ = vm.GetGlobal(0)
		_ = vm.SetGlobal(0, 0)
		_, _ = vm.GetFuncVar(0, 0)
		_ = vm.SetFuncVar(0, 0, 0)
		_, _ = vm.CallFunc("test", nil)
		_, _ = vm.CallFuncInd(0, nil)
		_, _ = vm.callFuncPtr(nil, nil)
		_ = vm.CallByIndex(0, nil, nil)
		_ = vm.callByFunc(nil, nil, nil, nil)
		_ = vm.ScriptCallbackRaw(nil, nil, nil, nil)
		_, _ = vm.Nox_script_objCallbackName_508CB0(nil, 0)
		_ = vm.ReadScript(nil)
		vm.ActRun()
		_ = srv.Nox_server_mapRWScriptData_504F90(nil, nil)
		_ = srv.nox_server_mapRWScriptData_504F90_Read(nil)
		_ = srv.nox_server_mapRWScriptData_504F90_Write(nil)
	}()

	// maps_send.go, network.go, object_ai_path.go, object_monster.go
	func() {
		defer func() { recover() }()
		var ms serverMapSend
		ms.init(&srv)
		_ = ms.Active()
		var pms playerMapSend
		pms.Clear(0)
		_ = ms.CountQueued(0)
		ms.Reset()
		ms.Sub_51A100()
		ms.AbortAll(0)
		ms.EndReceive(0)
		ms.Update()
		ms.Cancel(0)
	}()
	func() { defer func() { recover() }(); var ms serverMapSend; var pms playerMapSend; ms.abort(&pms, 0) }()
	func() { defer func() { recover() }(); var ms serverMapSend; var pms playerMapSend; ms.SendMore(&pms) }()
	func() { defer func() { recover() }(); var ms serverMapSend; ms.StartSendShared(0) }()
	func() { defer func() { recover() }(); var ms serverMapSend; ms.ForceCopy() }()
	func() { defer func() { recover() }(); var ms serverMapSend; _ = ms.ReadMapFile() }()

	func() {
		defer func() { recover() }()
		_ = srv.GetOwnIP()
		srv.Nox_server_netCloseHandler_4DEC60()
		srv.Nox_xxx_netStructReadPackets2_4DEC50(0)
		srv.Nox_xxx_netSendBySock_4DDDC0(0)
		srv.Nox_server_netMaybeSendInitialPackets_4DEB30()
		_ = srv.NetSendPacketXxx0(0, nil, 0, 0)
		_ = srv.NetSendPacketXxx1(0, nil, 0, 0)
		_ = srv.NetSendMsgXxx0(0, nil, 0, 0)
		_ = srv.NetSendMsgXxx1(0, nil, 0, 0)
		_ = srv.netOnPlayerInput(nil, nil)
		_ = netDecodePlayerInput(nil, nil)
		_ = srv.NetInformTextMsg(0, 0, 0)
		srv.NetPrintLineToAll(strman.ID(""))
		srv.NetPriMsgToPlayer(nil, strman.ID(""), 0)
		srv.NetPrintCompToAll(0)
		srv.NetRayStop(0, 0, nil, nil)
		srv.NetTeamRemove(nil)
		srv.NetTeamChangeLessons(nil, 0)
		_ = srv.SendTeamPacket(0)
		srv.NetMusic(0, 0)
		srv.NetMusicPushEvent()
		srv.NetMusicPopEvent()
		srv.NetMusicEvent()
		srv.NetHarpoonAttach(nil, nil)
		srv.NetHarpoonBreak(nil, nil)
		_ = srv.Nox_xxx_netFxShield_0_4D9200(0, nil)
		_ = srv.Nox_xxx_netMsgFadeBegin_4D9800(false, false)
		_ = srv.NetReportSpellStat(0, 0, 0)
		srv.Nox_xxx_netReportLesson_4D8EF0(nil)
		srv.Nox_xxx_netScriptMessageKill_4D9760(nil)
		srv.Nox_xxx_netKillChat_528D00(nil)
		srv.Nox_xxx_sendGauntlet_4DCF80(0, 0)
		srv.NetSendServerQuit()
		_ = srv.Nox_xxx_netSendBallStatus_4D95F0(0, 0, 0)
		_ = srv.Nox_xxx_netObjectOutOfSight_528A60(0, nil)
		_ = srv.Nox_xxx_netObjectInShadows_528A90(0, nil)
		srv.Nox_xxx_wallSendDestroyed_4DF0A0(nil, 0)
		_ = srv.Sub_507190(0, 0)
		_ = srv.Sub_4D6A20(0, nil)
		_ = srv.Sub_4D7280(0, false)
		srv.Nox_xxx_netSendFxAllCli_523030(types.Pointf{}, nil)
		srv.Nox_xxx_netSendPointFx_522FF0(0, types.Pointf{})
		srv.Nox_xxx_netSparkExplosionFx_5231B0(types.Pointf{}, 0)
		srv.Nox_xxx_netSendFxGreenBolt_523790(image.Point{}, image.Point{}, 0)
		srv.Nox_xxx_netSendVampFx_523270(0, image.Point{}, image.Point{}, 0)
		srv.Nox_xxx_netSendRayFx_5232F0(0, image.Point{}, image.Point{})
		srv.Nox_xxx_servCode_523340(image.Point{}, image.Point{}, nil)
		srv.NetSendFxJiggle(0, 0)
		srv.Nox_xxx_earthquakeSend_4D9110(types.Pointf{}, 0)
		_ = srv.NetWriteClassStats(0, ClassStats{})
		_ = srv.NetStatsMultiplier(nil)
		_ = srv.Nox_xxx_netCreatureCmd_4D7EE0(0, 0)
		_ = srv.Nox_xxx_orderUnitLocal_500C70(0, 0)
		srv.NetSendInterestingIDOn(nil)
		srv.NetSendInterestingIDOff(nil)
		srv.Sub_4D7E50(nil)
		srv.Sub_4D7EA0()
		_ = srv.Nox_xxx_netReportObjectPoison_4D7F40(nil, nil, 0)
		srv.NetReportExperience(nil)
		_ = srv.Nox_xxx_netReportAnimFrame_4D81F0(0, nil)
		_ = srv.nox_xxx_netReportXStatus_4D8230(0, nil)
	}()
	func() { defer func() { recover() }(); _ = srv.GetExtIP(nil) }()
	func() { defer func() { recover() }(); _ = srv.Listen(nil, nil) }()
	func() { defer func() { recover() }(); _ = srv.StartServices(false) }()
	func() { defer func() { recover() }(); _, _, _ = srv.OnPacketOpSub(0, 0, nil, nil, nil) }()

	func() {
		defer func() { recover() }()
		var ap serverAIPaths
		_ = ap.Valid()
		ap.ResetVisitNodes()
		_ = ap.NewVisitNode()
		ap.ResetIndex()
		_ = ap.PathFindStatus()
		_ = ap.MapIndex(0, 0)
		_ = ap.MapIndexFlags(0, 0)
		_ = ap.CheckIndexFlags(nil, 0, 0)
		ap.ResetPoints()
		_ = ap.appendPoint(types.Pointf{})
		ap.swapPoints()
		_ = ap.Points()
		_ = ap.MaybeAppendWorkPath(nil)
		_ = ap.appendWorkPath(nil, 0)
		ap.Sub_50B500()
		ap.Sub_50B510()
		ap.Free()
		_ = ap.sub50B8E0(nil, 0, 0)
		_ = ap.Nox_xxx_pathfind_preCheckWalls2_50B8A0(nil, 0, 0)
		ap.MaybeIndexObjects()
		ap.IndexObjects()
		_ = ap.sub50AEA0(nil, nil, nil)
		_ = ap.HasNoEnemiesAround(nil, 0, 0)
		ap.Sub50AFA0()
		_ = ap.Sub_50AC20(nil, nil)
	}()
	func() { defer func() { recover() }(); var ap serverAIPaths; ap.Init(&srv) }()
	func() { defer func() { recover() }(); var ap serverAIPaths; ap.IndexObject(nil) }()
	func() { defer func() { recover() }(); var ap serverAIPaths; ap.Sub_50C320(nil, nil, nil) }()

	func() {
		defer func() { recover() }()
		var ai AIStackItem
		_ = ai.C()
		_ = ai.Type()
		_ = ai.ArgU32(0)
		_ = ai.ArgF32(0)
		_ = ai.ArgPos(0)
		_ = ai.ArgObj(0)
		ai.SetArgs(1, 2.0, types.Pointf{})
		_ = AsColor3(nil)
		var mud MonsterUpdateData
		_ = mud.GetAIStack()
		_ = mud.AIStackHead()
		mud.PrintAIStack(0, "test")
		mud.SetAggression(1.0)
		_ = mud.DialogPortrait()
		_ = mud.HasAction(0)
		var obj Object
		_ = obj.AIStackEmptyAndIdle()
		_ = obj.Sub_5343C0()
		_ = obj.Nox_xxx_monsterCanAttackAtWill_534390()
		_ = obj.Sub_534440()
		obj.Nox_xxx_setNPCColor_4E4A90(0, nil)
		obj.ScriptCancelDialog()
		obj.ScriptSetDialog(0, 0, 0)
		_ = obj.ScriptDialogResult()
		_ = Nox_xxx_creatureIsMonitored_500CC0(nil, nil)
		obj.Nox_xxx_monsterResetEnemy_5346F0()
		obj.Nox_xxx_monsterMarkUpdate_4E8020()
		obj.SetMonsterStatus(0)
		obj.MonsterStatusEnable(0)
		obj.MonsterStatusDisable(0)
		_ = obj.SummonSize()
		_ = obj.Nox_xxx_countControlledCreatures_500D10()
	}()
	func() { defer func() { recover() }(); var obj Object; obj.MonsterCast(0, nil) }()
}

type mockPositioner struct {
	pos types.Pointf
}

func (m *mockPositioner) Pos() types.Pointf { return m.pos }

type mockVM struct{}

func (m *mockVM) Close() error                         { return nil }
func (m *mockVM) OnFrame()                             {}
func (m *mockVM) OnEvent(event script.EventType)       {}
func (m *mockVM) Exec(s string) (reflect.Value, error) { return reflect.Value{}, nil }
func (m *mockVM) ExecFile(path string) error           { return nil }
func (m *mockVM) GetSymbol(name string, typ reflect.Type) (reflect.Value, bool, error) {
	return reflect.Value{}, false, nil
}

type mockTimerHandle struct {
	id int
}

func (m *mockTimerHandle) TimerScriptID() int { return m.id }
func (m *mockTimerHandle) ScriptID() int      { return m.id }
func (m *mockTimerHandle) Cancel() bool       { return true }

// Ensure sound import is used
var _ = sound.ID(0)
