package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include "memmap.h"

extern uint32_t dword_5d4594_2487532;
extern uint32_t dword_5d4594_2487536;
extern uint32_t dword_5d4594_2487540;
extern uint32_t dword_5d4594_2487248;
extern uint32_t dword_5d4594_2487556;
extern uint32_t dword_5d4594_2487560;
extern uint32_t dword_5d4594_2487672;
extern uint32_t dword_5d4594_2487676;

int sub_521EB0(float* a1, float* a2);
int sub_521900(int a1, int a2, int a3);
int sub_520EA0(int a1);
void sub_520F80(void);
int sub_521200(int a1);
int sub_521290(int* a1);
void* nox_xxx_mapGenGetTopRoom_521710(void);
int nox_xxx_mapGenAddNewRoom_521730(uint32_t* a1);
int sub_521760(int a1);
int sub_521720(int a1);
int nox_xxx_mapGenUpdateRoomRect_521850(int a1);
int nox_xxx_mapGenSetRoomPos_521880(uint32_t* a1, float* a2);
float* nox_xxx_mapGenMakeRoomStruct_521940(int a1, int a2);
float* nox_xxx_mapGenPrepareRoom_521990(int a1);
void sub_521A10(void* p);
int sub_521A70(int a1, int a2, int a3);
int sub_521AA0(uint32_t* a1, int a2);
double sub_521B00(int a1, int a2);
double sub_521B30(int a1, int a2);
double sub_521B60(int a1, int a2);
double sub_521B90(int a1, int a2);
float* sub_521BC0(int a1, float* a2, float a3, float a4);
uint32_t* sub_521C10(int a1);
int sub_521F10(int a1, float* a2);
int sub_5226D0(int a1, float a2, int a3);
int sub_5227B0(int a1, float* a2);
char sub_522CA0(int a1, float* a2);
int sub_523920(int a1);
int sub_523970(int a1);
int sub_523C30(int a1, int a2);
int sub_523CB0(int a1, int a2);
float* sub_523D30(float* a1, float* a2);
float* sub_523E30(int a1, int a2, int a3);
float* nox_xxx_mapGenMakeHall_523EC0(int a1, int a2, int a3);
int sub_523960(int a1);
int sub_5239B0(int a1);
int sub_524070(int a1, int a2);
int sub_524090(int a1, int* a2);
char nox_xxx_mapGenDecorChkLimit_524220(int* a1, int a2);
int nox_xxx_mapgenAllocBuffer_5213E0(void);
void nox_xxx_mapgenFreeBuffer_521400(void);
uint32_t* nox_xxx_mapGenFreeTopRoom_521A40(void);
int sub_5217A0(int a1, int a2);
int sub_521820(int a1, int a2);
int sub_5218B0(int a1, int a2);
uint32_t* sub_522300(int a1, uint32_t* a2);
int sub_5244D0(int a1);
int sub_524500(float* a1, int a2);
int sub_524550(int* a1, int a2);
float* sub_5245A0(int a1, float* a2, int a3, int a4);
float* sub_524610(int a1, float* a2, int a3);
float* sub_525330(float* a1, int a2);
float* sub_525370(float* a1, int a2);
int sub_5253B0(float* a1);
void sub_5259F0(int a1, int a2, float a3);
int sub_5268F0(const char* a1);
int sub_526C40(int a1);
int sub_526C80(int a1);
int sub_526D50(int a1);
int sub_526DD0(float* a1, int* a2);
int sub_527380(float* a1);
int nox_xxx_mapGenReadLine_51E540(FILE* a1, uint8_t* a2);
FILE* nox_binfile_open_408CC0(char* path, int mode);
int nox_binfile_close_408D90(FILE* f);
int nox_xxx_genReadAlgData_51EBB0(int a1, FILE* a2);
int sub_51E800(int a1, uint32_t* a2);
int sub_51EAF0(int a1, uint32_t* a2);
int nox_xxx_genReadSpellSet_51EFB0(int a1, FILE* a2);
int nox_xxx_genReadWeaponSet_51F030(int a1, FILE* a2);
int nox_xxx_genReadArmorSet_51F640(int a1, FILE* a2);
int nox_xxx_genReadExit_51F800(int a1, FILE* a2);
int nox_xxx_genReadDecor_51F9F0(uint32_t* a1, FILE* a2);
int nox_xxx_genDecorReadOccurConstraint_520810(int a1, FILE* a2);
int nox_xxx_genDecorReadOccurLimit_5208D0(int a1, FILE* a2);
int nox_xxx_genDecorReadFrequency_520910(int a1, FILE* a2);
int nox_xxx_genDecorReadRoomSizeCon_5209F0(int a1, FILE* a2);
int nox_xxx_genDecorReadDoor_520A90(int a1, FILE* a2);
int nox_xxx_genDecorReadDoubleDoor_520AB0(int a1, FILE* a2);
int nox_xxx_mapgenCheckSettings_520AD0(int* a1);
char* nox_xxx_genDecorReadWallFloor_51FE00(int a1, FILE* a2);
int sub_51FEC0(int a1, int a2, FILE* a3);
int nox_xxx_genDecorReadDecorSet_51FFA0(int a1, FILE* a2);
uint32_t* nox_xxx_gen_520380(FILE* a1);
uint32_t* nox_xxx_gen_5205B0(FILE* a1);
void nox_xxx_mapGenFreeStr_51F1F0(void* p);
int nox_xxx_genReadPrefab_520BF0(int a1, FILE* a2);
char* sub_520CE0(int a1, FILE* a2);
uint32_t* sub_520D50(uint32_t* a1);
int nox_xxx_mapGenCheckRoomType_5238F0(int* a1);
int nox_xxx_mapGenDecorChkConstaint_5241C0(int a1, int a2);
int nox_xxx_mapGenChkDecorFillsRoom_5241F0(int a1, int a2);
int sub_51DE30(uint32_t* a1, uint32_t* a2, uint32_t* a3);
int nox_xxx_mapCountWallsMB_51DEA0(int a1);
int nox_xxx_AssignIfGreater_52A420(int* a1, int a2);

void nox_xxx_mapGenSetRngSeed_526AB0(unsigned int a1);
int nox_xxx_mapGenRandFunc_526AC0(int a1, int a2);
int nox_xxx_mapGenRandFunc2_526B00(int a1, int a2);
double sub_526BC0(float a1, float a2);
unsigned int nox_xxx_isObjectMovable_52E020(int a1);
char* sub_526AA0(int a1);
*/
import "C"
import "unsafe"

func C_nox_xxx_mapGenSetRngSeed_526AB0(seed uint32) {
	C.nox_xxx_mapGenSetRngSeed_526AB0(C.uint(seed))
}

func C_nox_xxx_mapGenRandFunc_526AC0(a1, a2 int32) int {
	return int(C.nox_xxx_mapGenRandFunc_526AC0(C.int(a1), C.int(a2)))
}

func C_nox_xxx_mapGenRandFunc2_526B00(a1, a2 int32) int {
	return int(C.nox_xxx_mapGenRandFunc2_526B00(C.int(a1), C.int(a2)))
}

// C_nox_xxx_isObjectMovable_52E020 backs a small object buffer; field8 and
// field16 are written at the raw offsets the C function reads.
func C_nox_xxx_isObjectMovable_52E020(field8, field16 uint32) int {
	buf := C.malloc(32)
	defer C.free(buf)
	// zero the buffer so unread fields are deterministic
	for i := 0; i < 32; i += 4 {
		*(*uint32)(unsafe.Pointer(uintptr(buf) + uintptr(i))) = 0
	}
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 8)) = field8
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 16)) = field16
	return int(C.nox_xxx_isObjectMovable_52E020(C.int(uintptr(buf))))
}

// C_sub_526AA0 returns the computed pointer as a uintptr so callers can verify
// the base + (a1<<6) stride relationship.
func C_sub_526AA0(a1 int32) uintptr {
	return uintptr(unsafe.Pointer(C.sub_526AA0(C.int(a1))))
}

func C_sub_521EB0(a1 []float32, a2 []float32) int {
	// a1 need at least 13 floats, a2 at least 5
	buf1 := C.malloc(C.size_t(13 * 4))
	buf2 := C.malloc(C.size_t(5 * 4))
	defer C.free(buf1)
	defer C.free(buf2)
	for i := 0; i < 13 && i < len(a1); i++ {
		*(*float32)(unsafe.Pointer(uintptr(buf1) + uintptr(i*4))) = a1[i]
	}
	for i := 0; i < 5 && i < len(a2); i++ {
		*(*float32)(unsafe.Pointer(uintptr(buf2) + uintptr(i*4))) = a2[i]
	}
	return int(C.sub_521EB0((*C.float)(buf1), (*C.float)(buf2)))
}

func C_nox_xxx_mapGenCheckRoomType_5238F0(v int32) int {
	buf := C.malloc(4)
	defer C.free(buf)
	*(*int32)(unsafe.Pointer(buf)) = v
	return int(C.nox_xxx_mapGenCheckRoomType_5238F0((*C.int)(buf)))
}

func C_sub_521900(count uint8, direction int, neighbor uint32) (ret int, after uint8, stored uint32) {
	buf := C.calloc(1, 512)
	defer C.free(buf)
	countPtr := (*uint8)(unsafe.Pointer(uintptr(buf) + uintptr(216+direction)))
	*countPtr = count
	ret = int(C.sub_521900(C.int(uintptr(buf)), C.int(neighbor), C.int(direction)))
	after = *countPtr
	if count < 8 {
		stored = *(*uint32)(unsafe.Pointer(uintptr(buf) + uintptr(88+4*(int(count)+8*direction))))
	}
	return ret, after, stored
}

func C_nox_xxx_mapGenDecorChkConstaint_5241C0(mask, roomFlags uint8) int {
	decor := C.calloc(1, 128)
	room := C.calloc(1, 400)
	defer C.free(decor)
	defer C.free(room)
	*(*uint8)(unsafe.Pointer(uintptr(decor) + 64)) = mask
	*(*uint8)(unsafe.Pointer(uintptr(room) + 364)) = roomFlags
	return int(C.nox_xxx_mapGenDecorChkConstaint_5241C0(C.int(uintptr(decor)), C.int(uintptr(room))))
}

func C_nox_xxx_mapGenChkDecorFillsRoom_5241F0(min, max, width, height int32) int {
	decor := C.calloc(1, 128)
	room := C.calloc(1, 32)
	defer C.free(decor)
	defer C.free(room)
	*(*int32)(unsafe.Pointer(uintptr(decor) + 76)) = min
	*(*int32)(unsafe.Pointer(uintptr(decor) + 80)) = max
	*(*int32)(unsafe.Pointer(uintptr(room) + 12)) = width
	*(*int32)(unsafe.Pointer(uintptr(room) + 16)) = height
	return int(C.nox_xxx_mapGenChkDecorFillsRoom_5241F0(C.int(uintptr(decor)), C.int(uintptr(room))))
}

func C_sub_526BC0(min, max float32) float64 {
	return float64(C.sub_526BC0(C.float(min), C.float(max)))
}

type cMapGenGridResult struct {
	initialized, added, found, collided, removed int
	topMatches, emptyAfter                       bool
}

// C_mapGenGridLifecycle exercises the map generator's private occupancy grid
// as one transaction so no C allocation or global list state escapes the call.
func C_mapGenGridLifecycle() cMapGenGridResult {
	cfg := C.calloc(1, 80)
	defer C.free(cfg)
	*(*uint32)(unsafe.Pointer(uintptr(cfg) + 68)) = 1 // 3x3 grid, centered at zero
	var out cMapGenGridResult
	out.initialized = int(C.sub_520EA0(C.int(uintptr(cfg))))
	if out.initialized == 0 {
		return out
	}
	defer func() {
		C.sub_520F80()
		C.dword_5d4594_2487532 = 0
		C.dword_5d4594_2487536 = 0
		C.dword_5d4594_2487540 = 0
	}()

	room := C.nox_xxx_mapGenMakeRoomStruct_521940(1, 1)
	if room == nil {
		return out
	}
	defer C.sub_521A10(unsafe.Pointer(room))
	pos := [2]C.float{0, 0}
	C.nox_xxx_mapGenSetRoomPos_521880((*C.uint32_t)(unsafe.Pointer(room)), &pos[0])
	out.added = int(C.nox_xxx_mapGenAddNewRoom_521730((*C.uint32_t)(unsafe.Pointer(room))))
	out.topMatches = C.nox_xxx_mapGenGetTopRoom_521710() == unsafe.Pointer(room)
	cell := [2]C.int{0, 0}
	out.found = int(C.sub_521290(&cell[0]))
	out.collided = int(C.sub_521200(C.int(uintptr(unsafe.Pointer(room)))))
	out.removed = int(C.sub_521760(C.int(uintptr(unsafe.Pointer(room)))))
	out.emptyAfter = C.nox_xxx_mapGenGetTopRoom_521710() == nil && C.sub_521290(&cell[0]) == 0
	return out
}

func C_sub_521720(value uint32, useNil bool) int {
	if useNil {
		return int(C.sub_521720(0))
	}
	p := C.calloc(1, 64)
	defer C.free(p)
	*(*uint32)(unsafe.Pointer(uintptr(p) + 56)) = value
	return int(C.sub_521720(C.int(uintptr(p))))
}

func C_mapGenRoomRect(width, height int32, x, y float32, setPos bool) (gridX, gridY int32, rect [4]float32) {
	room := C.calloc(1, 64)
	defer C.free(room)
	*(*int32)(unsafe.Pointer(uintptr(room) + 12)) = width
	*(*int32)(unsafe.Pointer(uintptr(room) + 16)) = height
	*(*float32)(unsafe.Pointer(uintptr(room) + 20)) = x
	*(*float32)(unsafe.Pointer(uintptr(room) + 24)) = y
	*(*float32)(unsafe.Pointer(uintptr(room) + 28)) = float32(width) * 32.526913
	*(*float32)(unsafe.Pointer(uintptr(room) + 32)) = float32(height) * 32.526913
	if setPos {
		pos := [2]C.float{C.float(x), C.float(y)}
		C.nox_xxx_mapGenSetRoomPos_521880((*C.uint32_t)(room), &pos[0])
	} else {
		C.nox_xxx_mapGenUpdateRoomRect_521850(C.int(uintptr(room)))
	}
	gridX = *(*int32)(unsafe.Pointer(uintptr(room) + 4))
	gridY = *(*int32)(unsafe.Pointer(uintptr(room) + 8))
	for i := range rect {
		rect[i] = *(*float32)(unsafe.Pointer(uintptr(room) + uintptr(36+4*i)))
	}
	return
}

func C_mapGenMakeAndPrepareRooms() (made [5]float32, prepared [2]int32) {
	room := C.nox_xxx_mapGenMakeRoomStruct_521940(7, 9)
	if room != nil {
		made = [5]float32{
			float32(*(*uint32)(unsafe.Pointer(room))),
			float32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 12))),
			float32(*(*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 16))),
			*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 28)),
			*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 32)),
		}
		C.sub_521A10(unsafe.Pointer(room))
	}
	cfg := C.calloc(1, 64)
	defer C.free(cfg)
	*(*int32)(unsafe.Pointer(uintptr(cfg) + 32)) = 10
	*(*int32)(unsafe.Pointer(uintptr(cfg) + 36)) = 0
	room = C.nox_xxx_mapGenPrepareRoom_521990(C.int(uintptr(cfg)))
	if room != nil {
		prepared = [2]int32{int32(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 12))), int32(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 16)))}
		C.sub_521A10(unsafe.Pointer(room))
	}
	return
}

func C_sub_521AA0(roomType uint32, direction int32) int {
	v := C.uint32_t(roomType)
	return int(C.sub_521AA0(&v, C.int(direction)))
}

func C_mapGenRoomRandEdges() [4]float64 {
	a := C.calloc(1, 64)
	b := C.calloc(1, 64)
	defer C.free(a)
	defer C.free(b)
	*(*int32)(unsafe.Pointer(uintptr(a) + 12)) = 4
	*(*int32)(unsafe.Pointer(uintptr(a) + 16)) = 5
	*(*float32)(unsafe.Pointer(uintptr(a) + 20)) = 100
	*(*float32)(unsafe.Pointer(uintptr(a) + 24)) = 200
	*(*int32)(unsafe.Pointer(uintptr(b) + 12)) = 4
	*(*int32)(unsafe.Pointer(uintptr(b) + 16)) = 5
	*(*float32)(unsafe.Pointer(uintptr(b) + 20)) = 300
	*(*float32)(unsafe.Pointer(uintptr(b) + 24)) = 400
	C.nox_xxx_mapGenSetRngSeed_526AB0(17)
	return [4]float64{
		float64(C.sub_521B00(C.int(uintptr(a)), C.int(uintptr(b)))),
		float64(C.sub_521B30(C.int(uintptr(a)), C.int(uintptr(b)))),
		float64(C.sub_521B60(C.int(uintptr(a)), C.int(uintptr(b)))),
		float64(C.sub_521B90(C.int(uintptr(a)), C.int(uintptr(b)))),
	}
}

type cMapGenRectsResult struct {
	inside, outside, overlap, separate int
	countBefore, countAfter            int
	randomRet                          int
	random                             [2]float32
}

func C_mapGenRectList() cMapGenRectsResult {
	room := C.calloc(1, 400)
	defer C.free(room)
	// Room bounds used by the random-point helper.
	for off, v := range map[int]float32{36: 0, 40: 0, 44: 20, 48: 20} {
		*(*float32)(unsafe.Pointer(uintptr(room) + uintptr(off))) = v
	}
	p := [2]C.float{2, 3}
	n1 := C.sub_521BC0(C.int(uintptr(room)), &p[0], 4, 5)
	p = [2]C.float{12, 13}
	n2 := C.sub_521BC0(C.int(uintptr(room)), &p[0], 2, 2)
	var out cMapGenRectsResult
	for n := *(*uintptr)(unsafe.Pointer(uintptr(room) + 368)); n != 0; n = *(*uintptr)(unsafe.Pointer(n + 24)) {
		out.countBefore++
	}
	pt := [2]C.float{13, 14}
	out.inside = int(C.sub_5227B0(C.int(uintptr(room)), &pt[0]))
	pt = [2]C.float{30, 30}
	out.outside = int(C.sub_5227B0(C.int(uintptr(room)), &pt[0]))
	candidate := C.calloc(1, 28)
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 4)) = 12
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 8)) = 13
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 12)) = 14
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 16)) = 15
	out.overlap = int(C.sub_521F10(C.int(uintptr(room)), (*C.float)(candidate)))
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 4)) = 50
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 8)) = 50
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 12)) = 55
	*(*float32)(unsafe.Pointer(uintptr(candidate) + 16)) = 55
	out.separate = int(C.sub_521F10(C.int(uintptr(room)), (*C.float)(candidate)))
	C.free(candidate)
	if n1 != nil {
		*(*uint32)(unsafe.Pointer(n1)) = 1
	}
	C.sub_521C10(C.int(uintptr(room)))
	for n := *(*uintptr)(unsafe.Pointer(uintptr(room) + 368)); n != 0; n = *(*uintptr)(unsafe.Pointer(n + 24)) {
		out.countAfter++
	}
	var random [2]C.float
	out.randomRet = int(C.sub_5226D0(C.int(uintptr(room)), 0.9, C.int(uintptr(unsafe.Pointer(&random[0])))))
	out.random = [2]float32{float32(random[0]), float32(random[1])}
	if n2 != nil {
		*(*uint32)(unsafe.Pointer(n2)) = 1
	}
	C.sub_521C10(C.int(uintptr(room)))
	return out
}

func C_sub_522CA0(points [][2]float32) (counts []int) {
	room := C.calloc(1, 400)
	defer C.free(room)
	for _, point := range points {
		p := [2]C.float{C.float(point[0]), C.float(point[1])}
		counts = append(counts, int(C.sub_522CA0(C.int(uintptr(room)), &p[0])))
	}
	return counts
}

func C_mapGenDirections() (a, b []int) {
	for i := int32(0); i <= 6; i++ {
		a = append(a, int(C.sub_523920(C.int(i))))
		b = append(b, int(C.sub_523970(C.int(i))))
	}
	return a, b
}

func C_mapGenHallGeometry(roomType, width, height int32) (dims [2]int32, spans [2]float32, c30, cb0 [2]float32, center [2]float32) {
	room := C.sub_523E30(C.int(roomType), C.int(width), C.int(height))
	if room == nil {
		return
	}
	defer C.free(unsafe.Pointer(room))
	dims = [2]int32{int32(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 12))), int32(*(*C.int)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 16)))}
	spans = [2]float32{
		*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 28)),
		*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + 32)),
	}
	// Give the room a non-zero rectangle for the edge/center functions.
	for off, v := range map[int]float32{36: 10, 40: 20, 44: 50, 48: 80} {
		*(*float32)(unsafe.Pointer(uintptr(unsafe.Pointer(room)) + uintptr(off))) = v
	}
	C.sub_523C30(C.int(uintptr(unsafe.Pointer(room))), C.int(uintptr(unsafe.Pointer(&c30[0]))))
	C.sub_523CB0(C.int(uintptr(unsafe.Pointer(room))), C.int(uintptr(unsafe.Pointer(&cb0[0]))))
	C.sub_523D30(room, (*C.float)(unsafe.Pointer(&center[0])))
	return
}

type cMapGenMoreBasics struct {
	bufferAllocated, bufferReset       bool
	freeEmptyTop, freeNonEmptyTop      bool
	fitEmpty, fitSmall, fitTooLarge    int
	neighborEmpty, neighborRoom, other int
	linkedObject                       bool
	selectedListNode                   int
	horizontalZero, verticalZero       int
	fillZero, columnZero               bool
	clearHorizontal, clearVertical     bool
	nameMissing                        int
	toggleA, toggleB, direction        []int
	coordNoOutput, borderWalk          int
	recursionParent, recursionChild    float32
}

func C_mapGenMoreBasics() cMapGenMoreBasics {
	var out cMapGenMoreBasics
	oldTicks := PlatformTicks
	PlatformTicks = func() uint64 { return 0 }
	defer func() { PlatformTicks = oldTicks }()
	out.bufferAllocated = C.nox_xxx_mapgenAllocBuffer_5213E0() != 0 && C.dword_5d4594_2487556 != 0
	C.nox_xxx_mapgenFreeBuffer_521400()
	C.dword_5d4594_2487556 = 0
	out.bufferReset = C.dword_5d4594_2487556 == 0

	C.dword_5d4594_2487560 = 0
	out.freeEmptyTop = C.nox_xxx_mapGenFreeTopRoom_521A40() == nil && C.dword_5d4594_2487560 == 0
	r1 := C.calloc(1, 376)
	r2 := C.calloc(1, 376)
	*(*uint32)(unsafe.Pointer(uintptr(r1) + 56)) = uint32(uintptr(r2))
	C.dword_5d4594_2487560 = C.uint32_t(uintptr(r1))
	C.nox_xxx_mapGenFreeTopRoom_521A40()
	out.freeNonEmptyTop = C.dword_5d4594_2487560 == 0

	limit := C.calloc(1, 80)
	room := C.calloc(1, 400)
	defer C.free(limit)
	defer C.free(room)
	*(*float32)(unsafe.Pointer(uintptr(limit) + 64)) = 50
	*(*int32)(unsafe.Pointer(uintptr(room) + 12)) = 2
	*(*int32)(unsafe.Pointer(uintptr(room) + 16)) = 2
	out.fitSmall = int(C.sub_5217A0(C.int(uintptr(limit)), C.int(uintptr(room))))
	*(*float32)(unsafe.Pointer(uintptr(limit) + 64)) = 1
	out.fitTooLarge = int(C.sub_5217A0(C.int(uintptr(limit)), C.int(uintptr(room))))
	*(*int32)(unsafe.Pointer(uintptr(room) + 16)) = 0
	out.fitEmpty = int(C.sub_521820(C.int(uintptr(limit)), C.int(uintptr(room))))

	neighbor := C.calloc(1, 400)
	defer C.free(neighbor)
	out.neighborEmpty = int(C.sub_5218B0(C.int(uintptr(room)), 0))
	*(*uint8)(unsafe.Pointer(uintptr(room) + 216)) = 1
	*(*uint32)(unsafe.Pointer(uintptr(room) + 88)) = uint32(uintptr(neighbor))
	*(*uint32)(neighbor) = 1
	out.neighborRoom = int(C.sub_5218B0(C.int(uintptr(room)), 0))
	*(*uint32)(neighbor) = 2
	out.other = int(C.sub_5218B0(C.int(uintptr(room)), 0))

	container := C.calloc(1, 512)
	obj1 := C.calloc(1, 512)
	obj2 := C.calloc(1, 512)
	defer C.free(container)
	defer C.free(obj1)
	defer C.free(obj2)
	C.sub_522300(C.int(uintptr(container)), (*C.uint32_t)(obj1))
	C.sub_522300(C.int(uintptr(container)), (*C.uint32_t)(obj2))
	out.linkedObject = *(*uint32)(unsafe.Pointer(uintptr(container) + 504)) == uint32(uintptr(obj2)) &&
		*(*uint32)(unsafe.Pointer(uintptr(obj2) + 496)) == uint32(uintptr(obj1)) &&
		*(*uint32)(unsafe.Pointer(uintptr(obj1) + 500)) == uint32(uintptr(obj2))

	list := C.calloc(1, 256)
	nodes := [3]unsafe.Pointer{C.calloc(1, 128), C.calloc(1, 128), C.calloc(1, 128)}
	defer C.free(list)
	for _, n := range nodes {
		defer C.free(n)
	}
	*(*uint32)(unsafe.Pointer(uintptr(list) + 84)) = uint32(uintptr(nodes[0]))
	*(*uint32)(unsafe.Pointer(uintptr(list) + 88)) = 3
	*(*uint32)(unsafe.Pointer(uintptr(nodes[0]) + 124)) = uint32(uintptr(nodes[1]))
	*(*uint32)(unsafe.Pointer(uintptr(nodes[1]) + 124)) = uint32(uintptr(nodes[2]))
	for i, n := range nodes {
		*(*uint32)(n) = uint32(i + 1)
	}
	C.nox_xxx_mapGenSetRngSeed_526AB0(9)
	selected := C.sub_5244D0(C.int(uintptr(list)))
	if selected != 0 {
		out.selectedListNode = int(*(*uint32)(unsafe.Pointer(uintptr(selected))))
	}

	point := [2]C.float{10, 20}
	out.horizontalZero = int(C.sub_524500(&point[0], 0))
	out.verticalZero = int(C.sub_524550((*C.int)(unsafe.Pointer(&point[0])), 0))
	out.fillZero = C.sub_5245A0(0, &point[0], 0, 0) == nil
	out.columnZero = C.sub_524610(0, &point[0], 0) == &point[0]
	out.clearHorizontal = C.sub_525330(&point[0], 0) == &point[0]
	out.clearVertical = C.sub_525370(&point[0], 0) == &point[0]

	name := C.CString("missing")
	out.nameMissing = int(C.sub_5268F0(name))
	C.free(unsafe.Pointer(name))
	for _, v := range []int{-1, 0, 1, 2} {
		out.toggleA = append(out.toggleA, int(C.sub_526C40(C.int(v))))
		out.toggleB = append(out.toggleB, int(C.sub_526C80(C.int(v))))
	}
	for _, v := range []int{-1, 0, 1, 14, 15} {
		out.direction = append(out.direction, int(C.sub_526D50(C.int(v))))
	}
	out.coordNoOutput = int(C.sub_526DD0(&point[0], nil))
	border := [2]C.float{0, 0}
	out.borderWalk = int(C.sub_527380(&border[0]))

	parent := C.calloc(1, 400)
	child := C.calloc(1, 400)
	defer C.free(parent)
	defer C.free(child)
	*(*uint8)(unsafe.Pointer(uintptr(parent) + 216)) = 1
	*(*uint32)(unsafe.Pointer(uintptr(parent) + 88)) = uint32(uintptr(child))
	*(*uint32)(parent) = 1
	*(*float32)(unsafe.Pointer(uintptr(parent) + 28)) = 3
	*(*float32)(unsafe.Pointer(uintptr(parent) + 32)) = 4
	C.sub_5259F0(C.int(uintptr(parent)), 0, 2)
	out.recursionParent = *(*float32)(unsafe.Pointer(uintptr(parent) + 356))
	out.recursionChild = *(*float32)(unsafe.Pointer(uintptr(child) + 356))
	return out
}

func C_mapGenReadTokens(f unsafe.Pointer, conditional bool) []string {
	buf := C.calloc(1, 256)
	defer C.free(buf)
	if conditional {
		buf = C.mem_getPtr(0x5D4594, 2487264)
	}
	var out []string
	for len(out) < 64 && C.nox_xxx_mapGenReadLine_51E540((*C.FILE)(f), (*C.uint8_t)(buf)) != 0 {
		out = append(out, C.GoString((*C.char)(buf)))
	}
	return out
}

func C_mapGenOpenInput(path string) unsafe.Pointer {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))
	return unsafe.Pointer(C.nox_binfile_open_408CC0(cpath, 0))
}

func C_mapGenCloseInput(f unsafe.Pointer) {
	C.nox_binfile_close_408D90((*C.FILE)(f))
}

type cMapGenAlgorithmData struct {
	ints  [16]int32
	mapSz float32
}

func C_mapGenReadAlgorithmData(f unsafe.Pointer) (int, cMapGenAlgorithmData) {
	cfg := C.calloc(1, 128)
	defer C.free(cfg)
	ret := int(C.nox_xxx_genReadAlgData_51EBB0(C.int(uintptr(cfg)), (*C.FILE)(f)))
	var out cMapGenAlgorithmData
	for i, off := range []int{4, 8, 12, 16, 20, 24, 28, 32, 36, 40, 44, 48, 52, 56, 60, 72} {
		out.ints[i] = *(*int32)(unsafe.Pointer(uintptr(cfg) + uintptr(off)))
	}
	out.mapSz = *(*float32)(unsafe.Pointer(uintptr(cfg) + 64))
	return ret, out
}

func C_mapGenReadOperators(f unsafe.Pointer, n int) (rets, values []int) {
	for i := 0; i < n; i++ {
		value := C.uint32_t(99)
		rets = append(rets, int(C.sub_51EAF0(C.int(uintptr(f)), &value)))
		values = append(values, int(value))
	}
	return
}

func C_mapGenReadConditions(f unsafe.Pointer, n int) (rets, values []int) {
	for i := 0; i < n; i++ {
		value := C.uint32_t(99)
		rets = append(rets, int(C.sub_51E800(C.int(uintptr(f)), &value)))
		values = append(values, int(value))
	}
	return
}

type cMapGenThemeSections struct {
	spellRet, weaponRet, armorRet, exitRet int
	spellCount, weaponCount, armorCount    uint32
	weaponName, weaponBase                 string
	armorName, armorBase                   string
	exitCount                              uint32
	exitNames                              [4]string
	exitDirections                         [4]uint32
	linkData                               string
}

func C_mapGenReadThemeSections(spellFile, weaponFile, armorFile, exitFile unsafe.Pointer) cMapGenThemeSections {
	var out cMapGenThemeSections
	theme := C.calloc(1, 1116)
	defer C.free(theme)
	out.spellRet = int(C.nox_xxx_genReadSpellSet_51EFB0(C.int(uintptr(theme)), (*C.FILE)(spellFile)))
	out.weaponRet = int(C.nox_xxx_genReadWeaponSet_51F030(C.int(uintptr(theme)), (*C.FILE)(weaponFile)))
	out.armorRet = int(C.nox_xxx_genReadArmorSet_51F640(C.int(uintptr(theme)), (*C.FILE)(armorFile)))
	out.exitRet = int(C.nox_xxx_genReadExit_51F800(C.int(uintptr(theme)), (*C.FILE)(exitFile)))
	out.spellCount = *(*uint32)(unsafe.Pointer(uintptr(theme) + 1096))
	out.weaponCount = *(*uint32)(unsafe.Pointer(uintptr(theme) + 1104))
	out.armorCount = *(*uint32)(unsafe.Pointer(uintptr(theme) + 1112))

	if p := uintptr(*(*uint32)(unsafe.Pointer(uintptr(theme) + 1100))); p != 0 {
		out.weaponBase = C.GoString((*C.char)(unsafe.Pointer(p)))
		out.weaponName = C.GoString((*C.char)(unsafe.Pointer(p + 60)))
	}
	if p := uintptr(*(*uint32)(unsafe.Pointer(uintptr(theme) + 1108))); p != 0 {
		out.armorBase = C.GoString((*C.char)(unsafe.Pointer(p)))
		out.armorName = C.GoString((*C.char)(unsafe.Pointer(p + 60)))
	}
	out.exitCount = *(*uint32)(unsafe.Pointer(uintptr(theme) + 472))
	for i := range out.exitNames {
		base := uintptr(theme) + 216 + uintptr(64*i)
		out.exitNames[i] = C.GoString((*C.char)(unsafe.Pointer(base)))
		out.exitDirections[i] = *(*uint32)(unsafe.Pointer(base + 60))
	}
	out.linkData = C.GoString((*C.char)(unsafe.Pointer(uintptr(theme) + 476)))

	for _, off := range []uintptr{1100, 1108} {
		p := uintptr(*(*uint32)(unsafe.Pointer(uintptr(theme) + off)))
		for p != 0 {
			next := uintptr(*(*uint32)(unsafe.Pointer(p + 152)))
			C.nox_xxx_mapGenFreeStr_51F1F0(unsafe.Pointer(p))
			p = next
		}
		*(*uint32)(unsafe.Pointer(uintptr(theme) + off)) = 0
	}
	return out
}

type cMapGenFullDecor struct {
	ret, kind, count        int
	name                    string
	constraint, limit, must uint8
	frequency               uint32
	min, max                int32
	door, doubleDoor        string
}

func C_mapGenReadFullDecor(f unsafe.Pointer) cMapGenFullDecor {
	var out cMapGenFullDecor
	theme := C.calloc(1, 256)
	defer C.free(theme)
	out.ret = int(C.nox_xxx_genReadDecor_51F9F0((*C.uint32_t)(theme), (*C.FILE)(f)))
	for kind, off := range []uintptr{88, 120, 152, 184} {
		p := uintptr(*(*uint32)(unsafe.Pointer(uintptr(theme) + off)))
		if p == 0 {
			continue
		}
		out.kind = kind
		out.count = int(*(*uint32)(unsafe.Pointer(uintptr(theme) + off + 4)))
		out.name = C.GoString((*C.char)(unsafe.Pointer(p)))
		out.constraint = *(*uint8)(unsafe.Pointer(p + 64))
		out.limit = *(*uint8)(unsafe.Pointer(p + 65))
		out.must = *(*uint8)(unsafe.Pointer(p + 67))
		out.frequency = *(*uint32)(unsafe.Pointer(p + 72))
		out.min = *(*int32)(unsafe.Pointer(p + 76))
		out.max = *(*int32)(unsafe.Pointer(p + 80))
		out.door = C.GoString((*C.char)(unsafe.Pointer(p + 100)))
		out.doubleDoor = C.GoString((*C.char)(unsafe.Pointer(p + 160)))
		C.free(unsafe.Pointer(p))
		*(*uint32)(unsafe.Pointer(uintptr(theme) + off)) = 0
		break
	}
	return out
}

type cMapGenHeadlessAlgorithms struct {
	connectionReturns  []int
	connectionCounts   [][2]int
	oppositeDirections []int
	hallDims           [][2]int32
	selectionReturns   []bool
	selectionLimits    []uint8
	selectionDisabled  []uint8
	selectionWeights   []int32
	stackEmptyRet      int
	stackRet           int
	stackValues        [3]uint32
	wallCountRet       int
	wallMins           [2]uint32
	assignReturns      [2]int
	assignValues       [2]int32
}

func C_mapGenHeadlessAlgorithmBatch() cMapGenHeadlessAlgorithms {
	var out cMapGenHeadlessAlgorithms
	for direction := 0; direction < 4; direction++ {
		a := C.calloc(1, 400)
		b := C.calloc(1, 400)
		ret := int(C.sub_521A70(C.int(uintptr(a)), C.int(uintptr(b)), C.int(direction)))
		out.connectionReturns = append(out.connectionReturns, ret)
		out.connectionCounts = append(out.connectionCounts, [2]int{
			int(*(*uint8)(unsafe.Pointer(uintptr(a) + 216 + uintptr(direction)))),
			int(*(*uint8)(unsafe.Pointer(uintptr(b) + 216))) +
				int(*(*uint8)(unsafe.Pointer(uintptr(b) + 217))) +
				int(*(*uint8)(unsafe.Pointer(uintptr(b) + 218))) +
				int(*(*uint8)(unsafe.Pointer(uintptr(b) + 219))),
		})
		out.oppositeDirections = append(out.oppositeDirections,
			int(C.sub_523960(C.int(direction))),
			int(C.sub_5239B0(C.int(direction+2))),
		)
		C.free(a)
		C.free(b)
	}

	cfg := C.calloc(1, 16)
	*(*int32)(unsafe.Pointer(uintptr(cfg) + 4)) = 7
	for _, kind := range []int{2, 3, 4, 5, 99} {
		hall := C.nox_xxx_mapGenMakeHall_523EC0(C.int(uintptr(cfg)), C.int(kind), 3)
		if hall == nil {
			out.hallDims = append(out.hallDims, [2]int32{})
			continue
		}
		out.hallDims = append(out.hallDims, [2]int32{
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(hall)) + 12)),
			*(*int32)(unsafe.Pointer(uintptr(unsafe.Pointer(hall)) + 16)),
		})
		C.free(unsafe.Pointer(hall))
	}
	C.free(cfg)

	weightIndex := map[uint8]uintptr{1: 2, 2: 3, 4: 4, 8: 5, 16: 6, 32: 7}
	for _, mask := range []uint8{1, 2, 4, 8, 16, 32} {
		theme := C.calloc(1, 256)
		room := C.calloc(1, 400)
		decor := C.calloc(1, 224)
		settings := unsafe.Pointer(uintptr(theme) + 184)
		*(*uint32)(settings) = uint32(uintptr(decor))
		for i := 2; i < 8; i++ {
			*(*int32)(unsafe.Pointer(uintptr(settings) + uintptr(4*i))) = 10
		}
		*(*uint8)(unsafe.Pointer(uintptr(room) + 364)) = mask
		*(*int32)(unsafe.Pointer(uintptr(room) + 12)) = 5
		*(*int32)(unsafe.Pointer(uintptr(room) + 16)) = 4
		*(*uint8)(unsafe.Pointer(uintptr(decor) + 64)) = mask
		*(*uint8)(unsafe.Pointer(uintptr(decor) + 65)) = 1
		*(*uint32)(unsafe.Pointer(uintptr(decor) + 72)) = 10
		*(*int32)(unsafe.Pointer(uintptr(decor) + 76)) = 1
		*(*int32)(unsafe.Pointer(uintptr(decor) + 80)) = 10
		selected := uint32(C.sub_524070(C.int(uintptr(theme)), C.int(uintptr(room))))
		out.selectionReturns = append(out.selectionReturns, selected == uint32(uintptr(decor)))
		out.selectionLimits = append(out.selectionLimits, *(*uint8)(unsafe.Pointer(uintptr(decor) + 65)))
		out.selectionDisabled = append(out.selectionDisabled, *(*uint8)(unsafe.Pointer(uintptr(decor) + 66)))
		out.selectionWeights = append(out.selectionWeights,
			*(*int32)(unsafe.Pointer(uintptr(settings) + 4*weightIndex[mask])))
		C.free(decor)
		C.free(room)
		C.free(theme)
	}

	var a, b, c C.uint32_t
	C.dword_5d4594_2487248 = 0
	out.stackEmptyRet = int(C.sub_51DE30(&a, &b, &c))
	*(*uint32)(C.mem_getPtr(0x973F18, 16200)) = 11
	*(*uint32)(C.mem_getPtr(0x973F18, 16204)) = 22
	*(*uint32)(C.mem_getPtr(0x973F18, 16208)) = 33
	C.dword_5d4594_2487248 = 1
	out.stackRet = int(C.sub_51DE30(&a, &b, &c))
	out.stackValues = [3]uint32{uint32(a), uint32(b), uint32(c)}
	C.dword_5d4594_2487248 = 0

	wall := C.calloc(1, 8)
	*(*uint8)(unsafe.Pointer(uintptr(wall) + 5)) = 12
	*(*uint8)(unsafe.Pointer(uintptr(wall) + 6)) = 34
	*(*uint32)(C.mem_getPtr(0x5D4594, 2487252)) = 256
	*(*uint32)(C.mem_getPtr(0x5D4594, 2487256)) = 256
	out.wallCountRet = int(C.nox_xxx_mapCountWallsMB_51DEA0(C.int(uintptr(wall))))
	out.wallMins = [2]uint32{
		*(*uint32)(C.mem_getPtr(0x5D4594, 2487252)),
		*(*uint32)(C.mem_getPtr(0x5D4594, 2487256)),
	}
	C.free(wall)

	value := C.int(5)
	out.assignReturns[0] = int(C.nox_xxx_AssignIfGreater_52A420(&value, -10))
	out.assignValues[0] = int32(value)
	out.assignReturns[1] = int(C.nox_xxx_AssignIfGreater_52A420(&value, 4))
	out.assignValues[1] = int32(value)
	return out
}

type cMapGenDecorParsed struct {
	constraint, limit uint8
	frequency         uint32
	min, max          int32
	door, doubleDoor  string
}

func C_mapGenReadDecorPrimitives(f unsafe.Pointer) (rets [6]int, out cMapGenDecorParsed) {
	decor := C.calloc(1, 256)
	defer C.free(decor)
	rets[0] = int(C.nox_xxx_genDecorReadOccurConstraint_520810(C.int(uintptr(decor)), (*C.FILE)(f)))
	rets[1] = int(C.nox_xxx_genDecorReadOccurLimit_5208D0(C.int(uintptr(decor)), (*C.FILE)(f)))
	rets[2] = int(C.nox_xxx_genDecorReadFrequency_520910(C.int(uintptr(decor)), (*C.FILE)(f)))
	rets[3] = int(C.nox_xxx_genDecorReadRoomSizeCon_5209F0(C.int(uintptr(decor)), (*C.FILE)(f)))
	rets[4] = int(C.nox_xxx_genDecorReadDoor_520A90(C.int(uintptr(decor)), (*C.FILE)(f)))
	rets[5] = int(C.nox_xxx_genDecorReadDoubleDoor_520AB0(C.int(uintptr(decor)), (*C.FILE)(f)))
	out.constraint = *(*uint8)(unsafe.Pointer(uintptr(decor) + 64))
	out.limit = *(*uint8)(unsafe.Pointer(uintptr(decor) + 65))
	out.frequency = *(*uint32)(unsafe.Pointer(uintptr(decor) + 72))
	out.min = *(*int32)(unsafe.Pointer(uintptr(decor) + 76))
	out.max = *(*int32)(unsafe.Pointer(uintptr(decor) + 80))
	out.door = C.GoString((*C.char)(unsafe.Pointer(uintptr(decor) + 100)))
	out.doubleDoor = C.GoString((*C.char)(unsafe.Pointer(uintptr(decor) + 160)))
	return
}

func C_mapGenReadFrequencies(f unsafe.Pointer, n int) (rets, values []int) {
	decor := C.calloc(1, 128)
	defer C.free(decor)
	for i := 0; i < n; i++ {
		rets = append(rets, int(C.nox_xxx_genDecorReadFrequency_520910(C.int(uintptr(decor)), (*C.FILE)(f))))
		values = append(values, int(*(*uint32)(unsafe.Pointer(uintptr(decor) + 72))))
	}
	return
}

func C_mapGenCheckSyntheticSettings() (success int, sums [6]int32, failures []int) {
	settings := C.calloc(1, 128)
	defer C.free(settings)
	nodes := make([]unsafe.Pointer, 7)
	for i := range nodes {
		nodes[i] = C.calloc(1, 224)
		defer C.free(nodes[i])
		*(*uint32)(unsafe.Pointer(uintptr(nodes[i]) + 72)) = uint32(i + 1)
	}
	for i := 0; i+1 < len(nodes); i++ {
		*(*uint32)(unsafe.Pointer(uintptr(nodes[i]) + 220)) = uint32(uintptr(nodes[i+1]))
	}
	// One unrestricted node contributes to every bucket, followed by each
	// individual constraint bit to cover all accumulation paths.
	for i, mask := range []uint8{0, 1, 2, 4, 8, 16, 32} {
		*(*uint8)(unsafe.Pointer(uintptr(nodes[i]) + 64)) = mask
	}
	*(*uint32)(settings) = uint32(uintptr(nodes[0]))
	success = int(C.nox_xxx_mapgenCheckSettings_520AD0((*C.int)(settings)))
	for i := range sums {
		sums[i] = *(*int32)(unsafe.Pointer(uintptr(settings) + uintptr(8+4*i)))
	}
	// A single constrained node leaves five buckets empty. Rotate its bit to
	// exercise each early-return condition and the final false return.
	for _, mask := range []uint8{1, 2, 4, 8, 16, 32} {
		*(*uint32)(settings) = uint32(uintptr(nodes[1]))
		*(*uint32)(unsafe.Pointer(uintptr(nodes[1]) + 220)) = 0
		*(*uint8)(unsafe.Pointer(uintptr(nodes[1]) + 64)) = mask
		failures = append(failures, int(C.nox_xxx_mapgenCheckSettings_520AD0((*C.int)(settings))))
	}
	return
}

type cMapGenCompositeParsers struct {
	wallFloorRet, appendRet, setRet int
	wallName, floorName             string
	appendType                      uint32
	appendFirst, appendSecond       string
	setMin, setMax, setCount        int32
	containsWeight                  uint32
	prefabRet, areaRet              int
	prefabName                      string
}

func C_mapGenCompositeParsers(wallFile, appendFile, setFile, containsFile, prefabFile unsafe.Pointer) cMapGenCompositeParsers {
	var out cMapGenCompositeParsers
	decor := C.calloc(1, 256)
	defer C.free(decor)
	wall := C.nox_xxx_genDecorReadWallFloor_51FE00(C.int(uintptr(decor)), (*C.FILE)(wallFile))
	if wall != nil {
		out.wallFloorRet = 1
		out.wallName = C.GoString(wall)
		out.floorName = C.GoString((*C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(wall)) + 60)))
		C.free(unsafe.Pointer(wall))
	}
	out.appendRet = int(C.sub_51FEC0(C.int(uintptr(decor)), 4, (*C.FILE)(appendFile)))
	if p := *(*uintptr)(unsafe.Pointer(uintptr(decor) + 120)); p != 0 {
		out.appendType = *(*uint32)(unsafe.Pointer(p))
		out.appendFirst = C.GoString((*C.char)(unsafe.Pointer(p + 4)))
		out.appendSecond = C.GoString((*C.char)(unsafe.Pointer(p + 64)))
		C.free(unsafe.Pointer(p))
		*(*uint32)(unsafe.Pointer(uintptr(decor) + 120)) = 0
	}
	out.setRet = int(C.nox_xxx_genDecorReadDecorSet_51FFA0(C.int(uintptr(decor)), (*C.FILE)(setFile)))
	if p := *(*uintptr)(unsafe.Pointer(uintptr(decor) + 92)); p != 0 {
		out.setMin = *(*int32)(unsafe.Pointer(p))
		out.setMax = *(*int32)(unsafe.Pointer(p + 4))
		out.setCount = *(*int32)(unsafe.Pointer(p + 12))
		C.free(unsafe.Pointer(p))
		*(*uint32)(unsafe.Pointer(uintptr(decor) + 92)) = 0
	}
	contains := C.nox_xxx_gen_520380((*C.FILE)(containsFile))
	if contains != nil {
		out.containsWeight = uint32(*contains)
		C.nox_xxx_mapGenFreeStr_51F1F0(unsafe.Pointer(contains))
	}
	theme := C.calloc(1, 1116)
	defer C.free(theme)
	out.prefabRet = int(C.nox_xxx_genReadPrefab_520BF0(C.int(uintptr(theme)), (*C.FILE)(prefabFile)))
	if p := *(*uintptr)(unsafe.Pointer(uintptr(theme) + 80)); p != 0 {
		out.prefabName = C.GoString((*C.char)(unsafe.Pointer(p)))
		out.areaRet = 1
	}
	C.sub_520D50((*C.uint32_t)(theme))
	return out
}
