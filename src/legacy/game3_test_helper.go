package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include <string.h>

short* sub_4A3090(short* a1, int a2);
int sub_4A2560(uint32_t* a1, int a2);
uint32_t* sub_4A2830(int a1, int a2, uint32_t* a3);
int sub_4A28B0(void);
int sub_4A28C0(int a1);
int sub_4A7A60(int a1);
int sub_4A7A70(int a1);
int sub_4A7A80(const char* a1);
int sub_4A7AC0(const char* a1);
int sub_4A7B00(const char* a1);
int sub_4A7B40(char* a1);
int sub_4A7BA0(char* a1);
int sub_4A7BC0(const char* a1);
int sub_4A7C00(const char* a1);
int sub_4A7C40(char* a1);
int sub_4A7C60(char* a1);
int sub_4A7CE0(char* a1);
int sub_4A7D00(const char* a1);
int sub_4A7D50(char* a1);
long long sub_4AEE30(void);
int sub_4AB340(int a1, int a2, int a3, int a4);
int sub_4AB390(int a1, int a2, int* a3, int a4);
int sub_4B4860(int a1, int a2, int a3, int a4);
short sub_4B69F0(int a1);
extern uint32_t dword_5d4594_1307716;
extern uint32_t dword_5d4594_1307720;

static int test_game3_string_fn(int which, char* value) {
	switch (which) {
	case 0: return sub_4A7A80(value);
	case 1: return sub_4A7AC0(value);
	case 2: return sub_4A7B00(value);
	case 3: return sub_4A7BA0(value);
	case 4: return sub_4A7BC0(value);
	case 5: return sub_4A7C00(value);
	case 6: return sub_4A7C40(value);
	case 7: return sub_4A7C60(value);
	case 8: return sub_4A7CE0(value);
	case 9: return sub_4A7D00(value);
	case 10: return sub_4A7D50(value);
	case 11: return sub_4A7B40(value);
	default: return -1;
	}
}
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func C_sub_4A3090(count int16, arr []uint32, idx int) (ret uintptr, outArr []uint32) {
	// allocate struct buffer large enough: need offset 48 for pointer to array
	structBuf := C.malloc(128)
	defer C.free(structBuf)
	arrBuf := C.malloc(C.size_t(len(arr) * 4))
	defer C.free(arrBuf)
	// write count at offset 0 as short
	*(*int16)(unsafe.Pointer(structBuf)) = count
	// write pointer to arr at offset 48
	*(*uint32)(unsafe.Pointer(uintptr(structBuf) + 48)) = uint32(uintptr(arrBuf))
	// copy arr values into arrBuf
	for i, v := range arr {
		*(*uint32)(unsafe.Pointer(uintptr(arrBuf) + uintptr(i*4))) = v
	}
	retPtr := C.sub_4A3090((*C.short)(structBuf), C.int(idx))
	ret = uintptr(unsafe.Pointer(retPtr))
	// read back array
	outArr = make([]uint32, len(arr))
	for i := range outArr {
		outArr[i] = *(*uint32)(unsafe.Pointer(uintptr(arrBuf) + uintptr(i*4)))
	}
	return
}

func C_sub_4A2560(x, y, objX, objY int, radius float64) int {
	old := *memmap.PtrFloat64(0x581450, 9720)
	*memmap.PtrFloat64(0x581450, 9720) = radius
	defer func() { *memmap.PtrFloat64(0x581450, 9720) = old }()
	point := [2]C.uint32_t{C.uint32_t(x), C.uint32_t(y)}
	obj := C.calloc(1, 48)
	defer C.free(obj)
	*(*C.short)(unsafe.Pointer(uintptr(obj) + 44)) = C.short(objX)
	*(*C.short)(unsafe.Pointer(uintptr(obj) + 46)) = C.short(objY)
	return int(C.sub_4A2560(&point[0], C.int(uintptr(obj))))
}

func C_sub_4A2830(x, y int) [2]int32 {
	var out [2]C.uint32_t
	C.sub_4A2830(C.int(x), C.int(y), &out[0])
	return [2]int32{int32(out[0]), int32(out[1])}
}

func C_game3SelectionGlobals() (empty bool, indexed uint32, missing uint32) {
	C.dword_5d4594_1307716 = 0
	empty = C.sub_4A28B0() == 0
	C.dword_5d4594_1307720 = 1
	*memmap.PtrUint32(0x5D4594, 1307316) = 0x12345678
	indexed = uint32(C.sub_4A28C0(0))
	missing = uint32(C.sub_4A28C0(1))
	C.dword_5d4594_1307720 = 0
	return
}

func C_game3SetState(value int) (mapState, mode int) {
	mapState = int(C.sub_4A7A60(C.int(value)))
	mode = int(C.sub_4A7A70(C.int(value + 1)))
	return
}

func C_game3StringFunction(which int, value *string) int {
	if value == nil {
		return int(C.test_game3_string_fn(C.int(which), nil))
	}
	str := C.CString(*value)
	defer C.free(unsafe.Pointer(str))
	return int(C.test_game3_string_fn(C.int(which), str))
}

func C_game3BlobString(off uintptr) string {
	return C.GoString((*C.char)(memmap.PtrOff(0x5D4594, off)))
}

func C_game3BlobInt(off uintptr) int32 {
	return *memmap.PtrInt32(0x5D4594, off)
}

func C_game3NameLookup(value string) (known, unknown int, stored uint32) {
	name := C.CString(value)
	defer C.free(unsafe.Pointer(name))
	saved := [3]uint32{
		*memmap.PtrUint32(0x587000, 171856),
		*memmap.PtrUint32(0x587000, 171860),
		*memmap.PtrUint32(0x587000, 171864),
	}
	defer func() {
		*memmap.PtrUint32(0x587000, 171856) = saved[0]
		*memmap.PtrUint32(0x587000, 171860) = saved[1]
		*memmap.PtrUint32(0x587000, 171864) = saved[2]
	}()
	*memmap.PtrUint32(0x587000, 171856) = uint32(uintptr(unsafe.Pointer(name)))
	*memmap.PtrUint32(0x587000, 171860) = 0xaabbccdd
	*memmap.PtrUint32(0x587000, 171864) = 0
	knownValue := value
	known = C_game3StringFunction(11, &knownValue)
	stored = *memmap.PtrUint32(0x5D4594, 1308184)
	unknownValue := "missing-name"
	unknown = C_game3StringFunction(11, &unknownValue)
	return
}

func C_game3SimpleEventHandlers() (results [7]int) {
	results[0] = int(C.sub_4AB340(0, 20, 0, 0))
	results[1] = int(C.sub_4AB340(0, 21, 1, 0))
	results[2] = int(C.sub_4AB340(0, 21, 0, 0))
	results[3] = int(C.sub_4AB390(0, 23, nil, 0))
	results[4] = int(C.sub_4AB390(0, 22, nil, 0))
	buf := C.calloc(1, 64)
	defer C.free(buf)
	results[5] = int(C.sub_4B4860(C.int(uintptr(buf)), 5, 0, 0))
	results[6] = int(C.sub_4B4860(C.int(uintptr(buf)), 0, 0, 0))
	return
}

func C_sub_4AEE30() int64 {
	return int64(C.sub_4AEE30())
}

func C_sub_4B69F0(position uint16, velocity uint8) (result int16, nextPosition uint16, nextVelocity uint8) {
	buf := C.calloc(1, 300)
	defer C.free(buf)
	*(*uint16)(unsafe.Pointer(uintptr(buf) + 104)) = position
	*(*uint8)(unsafe.Pointer(uintptr(buf) + 296)) = velocity
	result = int16(C.sub_4B69F0(C.int(uintptr(buf))))
	nextPosition = *(*uint16)(unsafe.Pointer(uintptr(buf) + 104))
	nextVelocity = *(*uint8)(unsafe.Pointer(uintptr(buf) + 296))
	return
}
