package legacy

/*
#include <stdint.h>
#include <stdlib.h>

extern uint32_t dword_5d4594_2516352;
extern uint32_t dword_5d4594_2516356;
extern uint32_t dword_5d4594_2516348;
extern uint32_t dword_5d4594_2516344;
extern uint32_t dword_5d4594_2516328;
extern uint32_t dword_5d4594_2523804;

unsigned int nox_xxx_netGetUnitCodeCli_578B00(int a1);
int nox_xxx_netClearHighBit_578B30(short a1);
unsigned int nox_xxx_netTestHighBit_578B70(unsigned int a1);
int nox_xxx_waypointNext_579870(int a1);
int sub_5798A0(int a1);
int sub_57BA10(int a1, short a2, short a3, int a4);
int nox_server_getNextMapGroup_57C090(int a1);
int sub_57CDB0(int* a1, float* a2, float* a3);
int nox_xxx_collideReflect_57B810(float* a1, int a2);
int nox_xxx_map_57B850(float* a1, float* a2, float* a3);
int sub_57B920(void* a1);
char nox_xxx_cliGenerateAlias_57B9A0(int a1, int a2, int a3, unsigned int a4);
void sub_57C790(float* a1, float* a2, float* a3, float a4);
int nox_xxx_mathPointOnTheLine_57C8A0(float* a1, float* a2, float* a3);
int nox_xxx_protectionStringCRC_56FAC0(int* a1, unsigned int a2);
int nox_xxx_protectionStringCRCLen_56FAE0(int* a1, unsigned int a2);
int sub_56FCB0(int a1, int a2);
void sub_56F720(int* a1, int* a2);
void sub_56FF00(int a1);
int sub_56FF80(int a1, int a2);
int nox_xxx_protectionCreateStructForInt_56F280(int a1, int a2);
int nox_xxx_protectionCreateStructForFloat_56F480(int a1, float a2);
int nox_xxx_protectionCreateInt_56F400(int a1);
uint32_t* sub_56F3B0(void);
uint32_t* sub_56F590(int a1);
uint32_t* sub_56F6F0(int a1);
int sub_56F510(int a1);
int sub_56F4F0(int* a1);
uint32_t* sub_56F780(int a1, int a2);
uint32_t* nox_xxx_playerResetProtectionCRC_56F7D0(int a1, int a2);
uint32_t* sub_56F820(int a1, unsigned char a2);
uint32_t* nox_xxx_protectPlayerHPMana_56F870(int a1, unsigned short a2);
uint32_t* sub_56F8C0(int a1, float a2);
uint32_t* sub_56F920(int a1, int a2);
uint32_t* nox_xxx_protectMana_56F9E0(int a1, short a2);
uint32_t* sub_56FA40(int a1, float a2);
int sub_56FB00(int* a1, unsigned int a2, int a3);
int nox_xxx_playerAwardSpellProtectionCRC_56FCE0(int a1, int a2, int a3);
int nox_xxx_playerApplyProtectionCRC_56FD50(int a1, void* a2, int a3);
int sub_56FB60(void* a1);
int nox_xxx_protect_56FBF0(int a1, void* a2);
int nox_xxx_protect_56FC50(int a1, void* a2);
int sub_56F1C0(void);
uint32_t* sub_579E70(void);
int nox_xxx_packetDynamicUnitCode_578B40(int a1);
int nox_xxx_playerCanTalkMB_57A160(int a1);
char* sub_57A1B0(short a1);
int sub_57AE30(const char* a1);
int nox_xxx_get_57AF20(void);
void sub_57B0A0(void);
long long nox_xxx___Getcvt_57B180(void);
int sub_57B190(unsigned short a1, unsigned short a2);
int* sub_57ADF0(int* a1);
char sub_57A1E0(int* a1, const char* a2, int* a3, char a4, short a5);
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func C_nox_xxx_netGetUnitCodeCli_578B00_nil() uint32 {
	return uint32(C.nox_xxx_netGetUnitCodeCli_578B00(0))
}

func C_nox_xxx_netGetUnitCodeCli_578B00(code, flags uint32) uint32 {
	buf := C.malloc(132)
	defer C.free(buf)
	base := uintptr(buf)
	*(*uint32)(unsafe.Pointer(base + 112)) = flags
	*(*uint32)(unsafe.Pointer(base + 128)) = code
	return uint32(C.nox_xxx_netGetUnitCodeCli_578B00(C.int(base)))
}

func C_nox_xxx_netClearHighBit_578B30(v int16) int {
	return int(C.nox_xxx_netClearHighBit_578B30(C.short(v)))
}

func C_nox_xxx_netTestHighBit_578B70(v uint32) uint32 {
	return uint32(C.nox_xxx_netTestHighBit_578B70(C.uint(v)))
}

func C_nox_xxx_waypointNext_579870_nil() int {
	return int(C.nox_xxx_waypointNext_579870(0))
}

func C_nox_xxx_waypointNext_579870(next uint32) int {
	buf := C.malloc(488)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 484)) = next
	return int(C.nox_xxx_waypointNext_579870(C.int(uintptr(buf))))
}

func C_sub_5798A0_nil() int {
	return int(C.sub_5798A0(0))
}

func C_sub_5798A0(next uint32) int {
	buf := C.malloc(488)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 484)) = next
	return int(C.sub_5798A0(C.int(uintptr(buf))))
}

func C_sub_57BA10(a2, a3 int16, a4 uint32) (ret uintptr, out0, out2 uint16, out4 uint32) {
	buf := C.malloc(8)
	defer C.free(buf)
	ret = uintptr(C.sub_57BA10(C.int(uintptr(buf)), C.short(a2), C.short(a3), C.int(a4)))
	out0 = *(*uint16)(unsafe.Pointer(uintptr(buf)))
	out2 = *(*uint16)(unsafe.Pointer(uintptr(buf) + 2))
	out4 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 4))
	return
}

func C_nox_server_getNextMapGroup_57C090_nil() int {
	return int(C.nox_server_getNextMapGroup_57C090(0))
}

func C_nox_server_getNextMapGroup_57C090(next uint32) int {
	buf := C.malloc(92)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 88)) = next
	return int(C.nox_server_getNextMapGroup_57C090(C.int(uintptr(buf))))
}

func C_sub_57CDB0(grid [2]int32, input [4]float32) (ret int, output [2]float32) {
	gridBuf := C.malloc(8)
	inputBuf := C.malloc(16)
	outputBuf := C.calloc(2, 4)
	defer C.free(gridBuf)
	defer C.free(inputBuf)
	defer C.free(outputBuf)
	*(*[2]int32)(gridBuf) = grid
	*(*[4]float32)(inputBuf) = input
	ret = int(C.sub_57CDB0((*C.int)(gridBuf), (*C.float)(inputBuf), (*C.float)(outputBuf)))
	output = *(*[2]float32)(outputBuf)
	return ret, output
}

func C_nox_xxx_collideReflect_57B810(normal, vector [2]float32) [2]float32 {
	normalBuf := C.malloc(8)
	vectorBuf := C.malloc(8)
	defer C.free(normalBuf)
	defer C.free(vectorBuf)
	*(*[2]float32)(normalBuf) = normal
	*(*[2]float32)(vectorBuf) = vector
	C.nox_xxx_collideReflect_57B810((*C.float)(normalBuf), C.int(uintptr(vectorBuf)))
	return *(*[2]float32)(vectorBuf)
}

func C_nox_xxx_map_57B850(origin [2]float32, data [11]float32, point [2]float32) int {
	originBuf := C.malloc(8)
	dataBuf := C.malloc(44)
	pointBuf := C.malloc(8)
	defer C.free(originBuf)
	defer C.free(dataBuf)
	defer C.free(pointBuf)
	*(*[2]float32)(originBuf) = origin
	*(*[11]float32)(dataBuf) = data
	*(*[2]float32)(pointBuf) = point
	return int(C.nox_xxx_map_57B850((*C.float)(originBuf), (*C.float)(dataBuf), (*C.float)(pointBuf)))
}

func C_sub_57B920() (ret int, allZero bool) {
	buf := C.malloc(0x7f8)
	defer C.free(buf)
	for i := uintptr(0); i < 0x7f8; i++ {
		*(*byte)(unsafe.Pointer(uintptr(buf) + i)) = 0xff
	}
	ret = int(C.sub_57B920(buf))
	allZero = true
	for i := uintptr(0); i < 0x7f8; i++ {
		if *(*byte)(unsafe.Pointer(uintptr(buf) + i)) != 0 {
			allZero = false
			break
		}
	}
	return ret, allZero
}

func C_nox_xxx_cliGenerateAlias_57B9A0(start, match int, exhaust bool) int8 {
	buf := C.calloc(256, 8)
	defer C.free(buf)
	if exhaust || match >= 0 {
		for i := 1; i < 256; i++ {
			p := uintptr(buf) + uintptr(8*i)
			*(*uint16)(unsafe.Pointer(p)) = 0x7ffe
			*(*uint16)(unsafe.Pointer(p + 2)) = 0x7fff
			*(*uint32)(unsafe.Pointer(p + 4)) = 100
		}
	}
	if match >= 0 {
		p := uintptr(buf) + uintptr(8*match)
		*(*uint16)(unsafe.Pointer(p)) = uint16(start)
		*(*uint16)(unsafe.Pointer(p + 2)) = 123
	}
	return int8(C.nox_xxx_cliGenerateAlias_57B9A0(C.int(uintptr(buf)), C.int(start), 123, 50))
}

func C_sub_57C790(line [4]float32, point [2]float32, scale float32) [2]float32 {
	lineBuf := C.malloc(16)
	pointBuf := C.malloc(8)
	outputBuf := C.calloc(2, 4)
	defer C.free(lineBuf)
	defer C.free(pointBuf)
	defer C.free(outputBuf)
	*(*[4]float32)(lineBuf) = line
	*(*[2]float32)(pointBuf) = point
	C.sub_57C790((*C.float)(lineBuf), (*C.float)(pointBuf), (*C.float)(outputBuf), C.float(scale))
	return *(*[2]float32)(outputBuf)
}

func C_nox_xxx_mathPointOnTheLine_57C8A0(line [4]float32, point [2]float32) (ret int, output [2]float32) {
	lineBuf := C.malloc(16)
	pointBuf := C.malloc(8)
	outputBuf := C.calloc(2, 4)
	defer C.free(lineBuf)
	defer C.free(pointBuf)
	defer C.free(outputBuf)
	*(*[4]float32)(lineBuf) = line
	*(*[2]float32)(pointBuf) = point
	ret = int(C.nox_xxx_mathPointOnTheLine_57C8A0((*C.float)(lineBuf), (*C.float)(pointBuf), (*C.float)(outputBuf)))
	return ret, *(*[2]float32)(outputBuf)
}

func C_nox_xxx_protectionStringCRC_56FAC0(words []uint32, byteLen uint32) int {
	if len(words) == 0 {
		return int(C.nox_xxx_protectionStringCRCLen_56FAE0(nil, C.uint(byteLen)))
	}
	buf := C.malloc(C.size_t(4 * len(words)))
	defer C.free(buf)
	for i, value := range words {
		*(*uint32)(unsafe.Pointer(uintptr(buf) + uintptr(4*i))) = value
	}
	return int(C.nox_xxx_protectionStringCRC_56FAC0((*C.int)(buf), C.uint(byteLen)))
}

func C_sub_56FCB0(bit int, enabled bool) int {
	var value C.int
	if enabled {
		value = 1
	}
	return int(C.sub_56FCB0(C.int(bit), value))
}

func C_sub_56F720(left, right [2]uint32) (afterLeft, afterRight [2]uint32) {
	leftBuf := C.malloc(8)
	rightBuf := C.malloc(8)
	defer C.free(leftBuf)
	defer C.free(rightBuf)
	*(*[2]uint32)(leftBuf) = left
	*(*[2]uint32)(rightBuf) = right
	C.sub_56F720((*C.int)(leftBuf), (*C.int)(rightBuf))
	return *(*[2]uint32)(leftBuf), *(*[2]uint32)(rightBuf)
}

func C_sub_56F720_nil() {
	C.sub_56F720(nil, nil)
}

func C_sub_56FF00(seed int) {
	C.sub_56FF00(C.int(seed))
}

func C_sub_56FF80(min, max int) int {
	return int(C.sub_56FF80(C.int(min), C.int(max)))
}

type protectionListResult struct {
	EmptyFind       bool
	FirstCreated    bool
	SecondCreated   bool
	CountAfterAdd   uint16
	DecodedFirst    uint32
	DecodedSecond   uint32
	IndexedBoth     bool
	MissingIndexNil bool
	RemovedFirst    bool
	RemovedAgain    bool
	ClearedHandle   bool
	CountAfterClear uint16
}

func resetProtectionList(key uint32) {
	C.sub_56F3B0()
	C.dword_5d4594_2516352 = 0
	C.dword_5d4594_2516344 = 0
	C.dword_5d4594_2516328 = 0
	C.dword_5d4594_2516348 = C.uint32_t(key)
	C.dword_5d4594_2516356 = 657757279
	*memmap.PtrUint16(0x587000, 311204) = 0
}

func C_protectionListLifecycle() (out protectionListResult) {
	const key = uint32(0x13579bdf)
	const firstID = 657757279
	const secondID = firstID + 1
	resetProtectionList(key)
	defer resetProtectionList(0)

	out.EmptyFind = C.sub_56F590(firstID) == nil
	out.FirstCreated = C.nox_xxx_protectionCreateStructForInt_56F280(firstID, 42) != 0
	first := C.sub_56F590(firstID)
	if first != nil {
		out.DecodedFirst = uint32((*[4]C.uint32_t)(unsafe.Pointer(first))[1]) ^ key
	}
	out.SecondCreated = C.nox_xxx_protectionCreateStructForFloat_56F480(secondID, C.float(3.5)) != 0
	second := C.sub_56F590(secondID)
	if second != nil {
		out.DecodedSecond = uint32((*[4]C.uint32_t)(unsafe.Pointer(second))[1]) ^ key
	}
	out.CountAfterAdd = *memmap.PtrUint16(0x587000, 311204)
	indexed0 := C.sub_56F6F0(0)
	indexed1 := C.sub_56F6F0(1)
	out.IndexedBoth = indexed0 != nil && indexed1 != nil && indexed0 != indexed1
	out.MissingIndexNil = C.sub_56F6F0(2) == nil
	if indexed0 != nil && indexed1 != nil {
		C.sub_56F720((*C.int)(unsafe.Pointer(indexed0)), (*C.int)(unsafe.Pointer(indexed1)))
	}
	out.RemovedFirst = C.sub_56F510(firstID) != 0
	out.RemovedAgain = C.sub_56F510(firstID) != 0
	handle := C.int(secondID)
	removed := C.sub_56F4F0(&handle)
	out.ClearedHandle = removed != 0 && handle == 0
	out.CountAfterClear = *memmap.PtrUint16(0x587000, 311204)
	return out
}

type protectionUpdateResult struct {
	UpdatesFound int
	ValidCRC     int
	InvalidCRC   int
	ValidFlags   int
	InvalidFlags int
	NilItemCRC   int
	MissingCalls int
	FinalCount   uint16
}

func C_protectionTypedUpdates() (out protectionUpdateResult) {
	const id = 657757279
	resetProtectionList(0x2468ace0)
	defer resetProtectionList(0)
	C.nox_xxx_protectionCreateStructForInt_56F280(id, 0)

	check := func() {
		if C.sub_56F590(id) != nil {
			out.UpdatesFound++
		}
	}
	C.sub_56F780(id, 10)
	check()
	C.nox_xxx_playerResetProtectionCRC_56F7D0(id, 20)
	check()
	C.sub_56F820(id, 30)
	check()
	C.nox_xxx_protectPlayerHPMana_56F870(id, 400)
	check()
	C.sub_56F8C0(id, C.float(12.75))
	check()
	C.sub_56F920(id, 5)
	check()
	C.nox_xxx_protectMana_56F9E0(id, 6)
	check()
	C.sub_56FA40(id, C.float(2.5))
	check()
	C.sub_56F780(id, 0)
	C.nox_xxx_playerAwardSpellProtectionCRC_56FCE0(id, 7, 1)
	check()
	flags := [9]C.int{}
	flags[7] = 1
	out.ValidFlags = int(C.nox_xxx_playerApplyProtectionCRC_56FD50(id, unsafe.Pointer(&flags[0]), 9))
	flags[7] = 0
	out.InvalidFlags = int(C.nox_xxx_playerApplyProtectionCRC_56FD50(id, unsafe.Pointer(&flags[0]), 9))
	out.NilItemCRC = int(C.sub_56FB60(nil))
	C.nox_xxx_protect_56FBF0(id, nil)
	check()
	C.nox_xxx_protect_56FC50(id, nil)
	check()

	words := [3]C.int{0x11111111, 0x22222222, 0x44444444}
	crc := int(words[0] ^ words[1] ^ words[2])
	C.sub_56F780(id, C.int(crc))
	out.ValidCRC = int(C.sub_56FB00(&words[0], 12, id))
	words[0]++
	out.InvalidCRC = int(C.sub_56FB00(&words[0], 12, id))

	missing := id + 1000
	cmissing := C.int(missing)
	if C.sub_56F780(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.nox_xxx_playerResetProtectionCRC_56F7D0(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.sub_56F820(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.nox_xxx_protectPlayerHPMana_56F870(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.sub_56F8C0(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.sub_56F920(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.nox_xxx_protectMana_56F9E0(cmissing, 1) == nil {
		out.MissingCalls++
	}
	if C.sub_56FA40(cmissing, 1) == nil {
		out.MissingCalls++
	}
	out.FinalCount = *memmap.PtrUint16(0x587000, 311204)
	return out
}

func C_protectionInitialize() (handle int, count uint16) {
	resetProtectionList(0)
	handle = int(C.sub_56F1C0())
	count = *memmap.PtrUint16(0x587000, 311204)
	resetProtectionList(0)
	return handle, count
}

func C_sub_579E70() (allocated, flagSet bool) {
	ptr := C.sub_579E70()
	if ptr == nil {
		return false, false
	}
	defer C.free(unsafe.Pointer(ptr))
	words := (*[129]C.uint32_t)(unsafe.Pointer(ptr))
	return true, uint32(words[120])&0x1000000 != 0
}

func C_nox_xxx_packetDynamicUnitCode_578B40(code int) int {
	return int(C.nox_xxx_packetDynamicUnitCode_578B40(C.int(code)))
}

func C_nox_xxx_playerCanTalkMB_57A160(flags uint32, nilPlayer bool) int {
	if nilPlayer {
		return int(C.nox_xxx_playerCanTalkMB_57A160(0))
	}
	buf := C.calloc(1, 3684)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 3680)) = flags
	return int(C.nox_xxx_playerCanTalkMB_57A160(C.int(uintptr(buf))))
}

type game52LookupResult struct {
	KnownRule   unsafe.Pointer
	KnownMask   int
	UnknownRule unsafe.Pointer
	UnknownMask int
}

func C_game52RuleLookups() (out game52LookupResult) {
	name := C.CString("fixture-rule")
	defer C.free(unsafe.Pointer(name))

	var saved [14]uint32
	for i := 0; i < 7; i++ {
		nameSlot := memmap.PtrUint32(0x587000, uintptr(312208+8*i))
		maskSlot := memmap.PtrUint32(0x587000, uintptr(312212+8*i))
		saved[2*i] = *nameSlot
		saved[2*i+1] = *maskSlot
		*nameSlot = uint32(uintptr(unsafe.Pointer(name)))
		*maskSlot = uint32(0x80 + 0x10*i)
	}
	defer func() {
		for i := 0; i < 7; i++ {
			*memmap.PtrUint32(0x587000, uintptr(312208+8*i)) = saved[2*i]
			*memmap.PtrUint32(0x587000, uintptr(312212+8*i)) = saved[2*i+1]
		}
	}()

	out.KnownRule = unsafe.Pointer(C.sub_57A1B0(0x80))
	out.KnownMask = int(C.sub_57AE30(name))
	out.UnknownRule = unsafe.Pointer(C.sub_57A1B0(0x7fff))
	unknown := C.CString("not-a-rule-type")
	defer C.free(unsafe.Pointer(unknown))
	out.UnknownMask = int(C.sub_57AE30(unknown))
	return out
}

func C_game52SimpleGlobals(value uint32) (get int, afterReset int, cvt int64) {
	C.dword_5d4594_2523804 = C.uint32_t(value)
	get = int(C.nox_xxx_get_57AF20())
	if value == 0 {
		C.sub_57B0A0()
	}
	afterReset = int(C.nox_xxx_get_57AF20())
	cvt = int64(C.nox_xxx___Getcvt_57B180())
	C.dword_5d4594_2523804 = 0
	return get, afterReset, cvt
}

func C_sub_57B190(left, right uint16) int {
	return int(C.sub_57B190(C.ushort(left), C.ushort(right)))
}

func C_game52RuleDefaults(flags int16) (ret int8, fields [7]int32) {
	buf := C.calloc(1, 64)
	defer C.free(buf)
	ret = int8(C.sub_57A1E0((*C.int)(buf), nil, nil, 0, C.short(flags)))
	for i := range fields {
		fields[i] = *(*int32)(unsafe.Pointer(uintptr(buf) + uintptr(24+4*i)))
	}
	emptyList := [2]C.int{}
	C.sub_57ADF0(&emptyList[0])
	return ret, fields
}
