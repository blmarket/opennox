package legacy

/*
#include <stdint.h>
#include <stdlib.h>

struct obj_412ae0_t;
struct obj_412ae0_t* nox_xxx_modifNext_4133C0(struct obj_412ae0_t* a1);
int sub_413F60(const void* a1, const void* a2);
void* sub_4168F0(void* a1);
void* sub_416910(void* a1);
int nox_xxx_wallSecretNext_410790(int* a1);
void nox_server_gameSettingsUpdated_40A670();
int nox_server_gameDoSwitchMap_40A680();
void nox_server_gameUnsetMapLoad_40A690();
int sub_40A740();
int nox_xxx_checkGameFlagPause_413A50();
char* sub_4165B0();
char* sub_416640();
char* sub_4169F0();
int nox_xxx_getServerSubFlags_409E60();
int sub_40A6B0();
int nox_xxx_rateGet_40A6C0();
int sub_4139B0();
unsigned int sub_409B50(const char* a1);
char* sub_409B80();
int sub_40A6A0(int a1);
long long sub_40A310(int a1);
int sub_40AA30(int a1);
int sub_40AA60(int a1);
int sub_409E40(int a1);
int sub_409E70(int a1);
int sub_410550(short a1);
uint32_t* sub_410730();
uint32_t* nox_xxx_wallSecretBlock_410760(uint32_t* a1);
int* sub_4107A0(void* a1);
void sub_4164F0();
char* sub_416630();
int sub_417DC0();
char sub_417DE0();
int sub_42EBA0();
int nox_xxx_cursor_430B00();
int sub_431370();
char* sub_413890();
int sub_415960(uint16_t* a1);
double sub_415BD0(int a1);
int sub_415DA0(uint16_t* a1);
void sub_416820(int a1);
void* sub_416860(int a1);
char* nox_xxx_netUnmarkMinimapSpec_417470(int a1, int a2);
int sub_418390();
int nox_xxx_teamAssignFlags_418640();
char* nox_xxx_toggleAllTeamFlags_418690(int a1);

extern uint32_t dword_5d4594_3484;
extern uint32_t dword_5d4594_251744;
extern uint32_t dword_5d4594_371692;
extern uint32_t dword_5d4594_526276;
extern void* dword_5d4594_251560;
*/
import "C"
import "unsafe"

func C_nox_xxx_modifNext_4133C0() (ret uintptr, expected uintptr) {
	buf := C.malloc(144)
	defer C.free(buf)
	target := C.malloc(16)
	defer C.free(target)
	*(*uintptr)(unsafe.Pointer(uintptr(buf) + 136)) = uintptr(target)
	ret = uintptr(unsafe.Pointer(C.nox_xxx_modifNext_4133C0((*C.struct_obj_412ae0_t)(buf))))
	expected = uintptr(target)
	return
}

func C_sub_413F60(v1, v2 uint32) int {
	a1 := C.malloc(84)
	a2 := C.malloc(84)
	defer C.free(a1)
	defer C.free(a2)
	*(*uint32)(unsafe.Pointer(uintptr(a1) + 80)) = v1
	*(*uint32)(unsafe.Pointer(uintptr(a2) + 80)) = v2
	return int(C.sub_413F60(a1, a2))
}

func C_sub_4168F0(withNext bool) uintptr {
	list := C.malloc(12)
	defer C.free(list)
	if withNext {
		next := C.malloc(12)
		defer C.free(next)
		*(*uintptr)(unsafe.Pointer(uintptr(list) + 0)) = uintptr(next)
		*(*uintptr)(unsafe.Pointer(uintptr(next) + 0)) = 0
		*(*uintptr)(unsafe.Pointer(uintptr(next) + 8)) = 1
		return uintptr(unsafe.Pointer(C.sub_4168F0(list)))
	}
	*(*uintptr)(unsafe.Pointer(uintptr(list) + 0)) = 0
	return uintptr(unsafe.Pointer(C.sub_4168F0(list)))
}

func C_sub_416910(withNext bool) uintptr {
	list := C.malloc(12)
	defer C.free(list)
	if withNext {
		next := C.malloc(12)
		defer C.free(next)
		*(*uintptr)(unsafe.Pointer(uintptr(list) + 0)) = uintptr(next)
		*(*uintptr)(unsafe.Pointer(uintptr(next) + 8)) = 1
		return uintptr(unsafe.Pointer(C.sub_416910(list)))
	}
	*(*uintptr)(unsafe.Pointer(uintptr(list) + 0)) = 0
	return uintptr(unsafe.Pointer(C.sub_416910(list)))
}

func C_nox_xxx_wallSecretNext_410790(withVal bool, val uint32) uint32 {
	if !withVal {
		return uint32(C.nox_xxx_wallSecretNext_410790(nil))
	}
	buf := C.malloc(4)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(buf)) = val
	return uint32(C.nox_xxx_wallSecretNext_410790((*C.int)(buf)))
}

func C_nox_server_gameSettingsUpdated_40A670() {
	C.nox_server_gameSettingsUpdated_40A670()
}

func C_nox_server_gameDoSwitchMap_40A680() int {
	return int(C.nox_server_gameDoSwitchMap_40A680())
}

func C_nox_server_gameUnsetMapLoad_40A690() {
	C.nox_server_gameUnsetMapLoad_40A690()
}

func C_sub_40A740_withSettingsByte(v byte) int {
	p := C.sub_4165B0()
	old := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 53))
	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 53)) = v
	ret := int(C.sub_40A740())
	*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 53)) = old
	return ret
}

func C_nox_xxx_checkGameFlagPause_413A50() int {
	return int(C.nox_xxx_checkGameFlagPause_413A50())
}

func C_sub_4169F0_resultByte(v byte) byte {
	p := C.sub_416640()
	bp := (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 100))
	old := *bp
	*bp = v
	ret := C.sub_4169F0()
	out := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ret)) + 100))
	*bp = old
	return out
}

func C_nox_xxx_getServerSubFlags_409E60() int {
	return int(C.nox_xxx_getServerSubFlags_409E60())
}

func C_sub_40A6B0() int {
	return int(C.sub_40A6B0())
}

func C_nox_xxx_rateGet_40A6C0() int {
	return int(C.nox_xxx_rateGet_40A6C0())
}

func C_sub_4139B0() int {
	return int(C.sub_4139B0())
}

func C_sub_409B50(s string) uint {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return uint(C.sub_409B50(cs))
}

func C_sub_409B80() string {
	return C.GoString(C.sub_409B80())
}

func C_sub_40A6A0(a1 int) int {
	return int(C.sub_40A6A0(C.int(a1)))
}

func C_sub_40A310(v int) int64 {
	return int64(C.sub_40A310(C.int(v)))
}

func C_sub_40AA30(v int) int {
	return int(C.sub_40AA30(C.int(v)))
}

func C_sub_40AA60(v int) int {
	return int(C.sub_40AA60(C.int(v)))
}

type game1WallSecretResult struct {
	firstFound, secondFound uintptr
	middleRemoved           bool
	headRemoved             bool
	missingReturnedNil      bool
	emptyAfter              bool
}

func C_game1WallSecretLifecycle() game1WallSecretResult {
	C.dword_5d4594_251560 = nil
	defer func() {
		C.sub_410730()
		C.dword_5d4594_251560 = nil
	}()

	wall1 := C.calloc(1, 16)
	wall2 := C.calloc(1, 16)
	defer C.free(wall1)
	defer C.free(wall2)
	*(*uint16)(unsafe.Pointer(uintptr(wall1) + 10)) = 0x1234
	*(*uint16)(unsafe.Pointer(uintptr(wall2) + 10)) = 0x5678

	node1 := C.calloc(1, 16)
	node2 := C.calloc(1, 16)
	*(*uintptr)(unsafe.Pointer(uintptr(node1) + 12)) = uintptr(wall1)
	*(*uintptr)(unsafe.Pointer(uintptr(node2) + 12)) = uintptr(wall2)
	C.nox_xxx_wallSecretBlock_410760((*C.uint32_t)(node1))
	C.nox_xxx_wallSecretBlock_410760((*C.uint32_t)(node2))

	var out game1WallSecretResult
	out.firstFound = uintptr(C.sub_410550(0x1234))
	out.secondFound = uintptr(C.sub_410550(0x5678))
	out.middleRemoved = C.sub_4107A0(node1) != nil
	out.headRemoved = C.sub_4107A0(node2) != nil

	missing := C.calloc(1, 16)
	out.missingReturnedNil = C.sub_4107A0(missing) == nil
	C.free(missing)
	out.emptyAfter = C.sub_410550(0x1234) == 0 && C.sub_410550(0x5678) == 0
	return out
}

func C_game1RemainingStateHelpers(marker uint32) (beforeReset, afterReset uint32, shared uintptr, teamCount int, mapInfo uint32) {
	oldReset := C.dword_5d4594_371692
	oldMapInfo := C.dword_5d4594_526276
	defer func() {
		C.dword_5d4594_371692 = oldReset
		C.dword_5d4594_526276 = oldMapInfo
	}()
	C.dword_5d4594_371692 = C.uint32_t(marker)
	beforeReset = uint32(C.dword_5d4594_371692)
	C.sub_4164F0()
	afterReset = uint32(C.dword_5d4594_371692)
	shared = uintptr(unsafe.Pointer(C.sub_416630()))
	C.dword_5d4594_526276 = C.uint32_t(marker)
	mapInfo = uint32(C.sub_417DC0())
	teamCount = int(C.sub_417DE0())
	return
}

type game1SafeEmptyResult struct {
	emptyGeneratedName uintptr
	weaponName         int
	armorDefense       float64
	armorName          int
	unmarkResult       uintptr
	teamStart          int
	assignResult       int
	toggleOff          uintptr
	toggleOn           uintptr
}

func C_game1SafeEmptyPaths() (out game1SafeEmptyResult) {
	out.emptyGeneratedName = uintptr(unsafe.Pointer(C.sub_413890()))
	out.weaponName = int(C.sub_415960(nil))
	obj := C.calloc(1, 16)
	defer C.free(obj)
	out.armorDefense = float64(C.sub_415BD0(C.int(uintptr(obj))))
	out.armorName = int(C.sub_415DA0(nil))
	C.sub_416820(0)
	C.sub_416860(0)
	out.unmarkResult = uintptr(unsafe.Pointer(C.nox_xxx_netUnmarkMinimapSpec_417470(0, 0)))
	out.teamStart = int(C.sub_418390())
	out.assignResult = int(C.nox_xxx_teamAssignFlags_418640())
	out.toggleOff = uintptr(unsafe.Pointer(C.nox_xxx_toggleAllTeamFlags_418690(0)))
	out.toggleOn = uintptr(unsafe.Pointer(C.nox_xxx_toggleAllTeamFlags_418690(1)))
	return out
}

func C_sub_409E40(a1 int) int {
	return int(C.sub_409E40(C.int(a1)))
}

func C_sub_409E70(a1 int) int {
	return int(C.sub_409E70(C.int(a1)))
}

func C_sub_42EBA0() int {
	return int(C.sub_42EBA0())
}

func C_nox_xxx_cursor_430B00() int {
	return int(C.nox_xxx_cursor_430B00())
}

func C_sub_431370() int {
	return int(C.sub_431370())
}

type Game1Globals struct {
	v3484   uint32
	v251744 uint32
}

func C_game1Globals() Game1Globals {
	return Game1Globals{
		v3484:   uint32(C.dword_5d4594_3484),
		v251744: uint32(C.dword_5d4594_251744),
	}
}

func C_game1SetGlobals(v Game1Globals) {
	C.dword_5d4594_3484 = C.uint32_t(v.v3484)
	C.dword_5d4594_251744 = C.uint32_t(v.v251744)
}
