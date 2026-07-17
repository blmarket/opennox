package legacy

/*
#include <stdint.h>
#include <stdlib.h>

extern unsigned int dword_5d4594_1599628;
int sub_512E80(int);
int* nox_server_scriptMoveTo_5123C0(int, int);
void nox_xxx_scriptMonsterRoam_512930(void*);
void nox_server_gotoHome(void*);
int nox_script_SetRoamFlag_515C40(int);
int nox_script_GiveExp_516190(void);
int nox_script_MakeFriendly_516720(void);
int nox_script_MakeEnemy_516760(void);
int nox_script_builtin_516790(void*);
int nox_script_BecomePet_5167D0(void);
int nox_script_BecomeEnemy_516810(void);
int nox_script_builtin_516850(void*);
int nox_script_SetShopkeeperGreet_516BE0(void);
int nox_script_IsTalking_5166A0(void);
int nox_script_PlayerIsTrading_5166E0(void);

static int script_builtin_case(int which) {
	switch (which) {
	case 0: return nox_script_SetRoamFlag_515C40(0);
	case 1: return nox_script_GiveExp_516190();
	case 2: return nox_script_MakeFriendly_516720();
	case 3: return nox_script_MakeEnemy_516760();
	case 4: return nox_script_builtin_516790((void*)17);
	case 5: return nox_script_BecomePet_5167D0();
	case 6: return nox_script_BecomeEnemy_516810();
	case 7: return nox_script_builtin_516850((void*)23);
	case 8: return nox_script_SetShopkeeperGreet_516BE0();
	case 9: return nox_script_IsTalking_5166A0();
	case 10: return nox_script_PlayerIsTrading_5166E0();
	default: return -1;
	}
}

static void script_builtin_guarded_objects(void) {
	unsigned char* obj = calloc(1, 752);
	unsigned char* target = calloc(1, 480);
	*(uint32_t*)(obj + 16) = 0x8000;
	nox_server_scriptMoveTo_5123C0((int)obj, (int)target);
	*(uint32_t*)(obj + 16) = 0;
	nox_xxx_scriptMonsterRoam_512930(obj);
	nox_server_gotoHome(obj);
	free(target);
	free(obj);
}
*/
import "C"

import (
	"math"
	"testing"
	"unsafe"

	ns "github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/strman"
	"github.com/noxworld-dev/opennox/v1/server"
)

type scriptBuiltinVM struct {
	*server.NoxScriptVM
	stack   []uint32
	strings []string
}

func (vm *scriptBuiltinVM) PushI32(v int32)   { vm.stack = append(vm.stack, uint32(v)) }
func (vm *scriptBuiltinVM) PushF32(v float32) { vm.stack = append(vm.stack, math.Float32bits(v)) }
func (vm *scriptBuiltinVM) PopI32() int32 {
	if len(vm.stack) == 0 {
		return 0
	}
	v := vm.stack[len(vm.stack)-1]
	vm.stack = vm.stack[:len(vm.stack)-1]
	return int32(v)
}
func (vm *scriptBuiltinVM) PopF32() float32 { return math.Float32frombits(uint32(vm.PopI32())) }
func (vm *scriptBuiltinVM) LookupString(i uint32) (string, bool) {
	if i >= uint32(len(vm.strings)) {
		return "", false
	}
	return vm.strings[i], true
}
func (vm *scriptBuiltinVM) ScriptToObject(int) *server.Object { return nil }
func (vm *scriptBuiltinVM) ActResolveObjs()                   {}
func (vm *scriptBuiltinVM) NoxScript() ns.Implementation      { return nil }
func (vm *scriptBuiltinVM) PopObjectNS() ns.Obj               { return nil }
func (vm *scriptBuiltinVM) PopObjGroupNS() ns.ObjGroup        { return nil }
func (vm *scriptBuiltinVM) PopWallGroupNS() ns.WallGroupObj   { return nil }
func (vm *scriptBuiltinVM) ScriptCallback(*server.ScriptCallback, *server.Object, *server.Object, server.ScriptEventType) unsafe.Pointer {
	return nil
}

type scriptBuiltinServer struct {
	Server
	srv *server.Server
	vm  *scriptBuiltinVM
}

func (s *scriptBuiltinServer) S() *server.Server     { return s.srv }
func (s *scriptBuiltinServer) NoxScriptC() NoxScript { return s.vm }

func setScriptBuiltinServer(t testing.TB) *scriptBuiltinVM {
	t.Helper()
	oldGetServer := GetServer
	srv := server.New(nil, strman.New())
	vm := &scriptBuiltinVM{NoxScriptVM: &srv.NoxScriptVM}
	GetServer = func() Server { return &scriptBuiltinServer{srv: srv, vm: vm} }
	t.Cleanup(func() {
		GetServer = oldGetServer
		srv.Close()
	})
	return vm
}

func C_sub_512E80Sequence() (first, last int, count uint32) {
	C.dword_5d4594_1599628 = 0
	first = int(C.sub_512E80(0x12345678))
	for i := 1; i < 1025; i++ {
		last = int(C.sub_512E80(C.int(i)))
	}
	count = uint32(C.dword_5d4594_1599628)
	C.dword_5d4594_1599628 = 0
	return
}

func C_scriptBuiltinGuardedObjects() { C.script_builtin_guarded_objects() }

func C_scriptBuiltinCase(which int) int { return int(C.script_builtin_case(C.int(which))) }
