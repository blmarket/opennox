package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_521EB0(float* a1, float* a2);
int sub_521900(int a1, int a2, int a3);
int nox_xxx_mapGenCheckRoomType_5238F0(int* a1);
int nox_xxx_mapGenDecorChkConstaint_5241C0(int a1, int a2);
int nox_xxx_mapGenChkDecorFillsRoom_5241F0(int a1, int a2);

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
