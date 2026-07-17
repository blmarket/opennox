package legacy

/*
#include <stdint.h>
#include <stdlib.h>

extern uint32_t dword_5d4594_1556856;

int sub_4E14A0();
int sub_479D00();
int sub_4866D0(uint32_t* a1, int a2);
int nox_xxx_mapGetTypeMB_4CFFA0(void* a1);
int sub_4CFFC0(int a1);
int sub_4CE340(int a1, int a2);
int nox_xxx_updDrawCloud_4CE1D0(int a1, int a2);
int sub_4CE360(int a1, int a2);
float* nox_xxx_effectDamageMultiplier_4E04C0(int a1, int a2, int a3, int a4, float* a5);
int nox_xxx_effectProjectileSpeed_4E09B0(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_inversionEffect_4E03D0(int a1, int a2, int a3, int a4, int a5, int* a6);
int nox_xxx_gripEffect_4E0480(int a1, int a2, int a3, int a4, int a5, int* a6);
void sub_4D11A0(void);
void sub_4D11D0(void);
void sub_4D1210(int a1);
int* sub_4D1250(int a1);
int sub_4D12A0(int a1);
void nox_xxx_tileInitdataClear_4D3C50(const void* a1);
uint32_t* sub_4D3C80(uint32_t* a1);
int sub_4D3E30(float* a1, float* a2);
int sub_4D3FF0(int a1);
unsigned int sub_4D42E0(const char* a1);
void* nox_xxx_getRandMapName_4D4310(void);
void nox_xxx_effectSpeedEngage_4DFC30(int a1, int a2);
void nox_xxx_effectSpeedDisengage_4DFCA0(int a1, int a2);
void sub_4E0170(int a1, int a2);
void nox_xxx_effectRegeneration_4E01D0(int a1, int a2);
void nox_xxx_attribContinualReplen_4E02C0(int a1, uint32_t* a2);
int nox_xxx_unusedCheckGripEffect_4E03F0(int a1, int a2, int a3, int a4);
void nox_xxx_stunEffect_4E04D0(int a1, int a2, int a3, int a4);
void nox_xxx_recoilEffect_4E0640(int a1, int a2, int a3, int a4);
void nox_xxx_confuseEffect_4E0670(int a1, int a2, int a3, int a4);
void nox_xxx_lightngEffect_4E06F0(int a1, int a2, int a3, int a4);
void nox_xxx_drainMEffect_4E0740(int a1, int a2, int a3, int a4, int* a5);
void nox_xxx_vampirismEffect_4E07C0(int a1, int a2, int a3, int a4, int* a5);
void nox_xxx_poisonEffect_4E0850(int a1, int a2, int a3, int a4);
void nox_xxx_sympathyEffect_4E08E0(int a1, int a2, int a3, int a4, int* a5);
int sub_4E1470(int a1);
*/
import "C"
import (
	"unsafe"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

func C_resetShadowList() {
	C.dword_5d4594_1556856 = 0
}

func init() {
	// Ensure pure mapping is available in legacy tests without main package init.
	if Nox_mapToGameFlags == nil {
		Nox_mapToGameFlags = func(v int) noxflags.GameFlag {
			var out noxflags.GameFlag
			if v&1 != 0 {
				out |= noxflags.GameModeCoopTeam
			}
			if v&2 != 0 {
				out |= noxflags.GameModeQuest
			}
			if v&4 != 0 {
				out |= noxflags.GameModeArena
			}
			if v&0x20 != 0 {
				out |= noxflags.GameModeElimination
			}
			if v&8 != 0 {
				out |= noxflags.GameModeCTF
			}
			if v&0x10 != 0 {
				out |= noxflags.GameModeKOTR
			}
			if v&0x40 != 0 {
				out |= noxflags.GameModeFlagBall
			}
			if v < 0 {
				out |= noxflags.GameModeChat
			}
			return out
		}
	}
}

func C_sub_4E14A0() int {
	return int(C.sub_4E14A0())
}

func C_sub_479D00() int {
	return int(C.sub_479D00())
}

func C_sub_4866D0(v uint32, a2 int) int {
	buf := C.malloc(4)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(buf)) = v
	return int(C.sub_4866D0((*C.uint32_t)(buf), C.int(a2)))
}

func C_nox_xxx_mapGetTypeMB_4CFFA0(v uint32) int {
	buf := C.malloc(2048)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 1392)) = v
	return int(C.nox_xxx_mapGetTypeMB_4CFFA0(buf))
}

func C_sub_4CFFC0(v uint32) int {
	buf := C.malloc(64)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 28)) = v
	return int(C.sub_4CFFC0(C.int(uintptr(buf))))
}

func C_sub_4CE340(val104 uint16, add432 uint8) (ret int, after uint16) {
	buf := C.malloc(512)
	defer C.free(buf)
	*(*uint16)(unsafe.Pointer(uintptr(buf) + 104)) = val104
	*(*uint8)(unsafe.Pointer(uintptr(buf) + 432)) = add432
	ret = int(C.sub_4CE340(0, C.int(uintptr(buf))))
	after = *(*uint16)(unsafe.Pointer(uintptr(buf) + 104))
	return
}

type game32SafePathsResult struct {
	listContains, listRemoved int
	sorted                    [8]uint32
	transformNil              int
	transform                 [2]float32
	directions                []int
	nameLen                   uint32
	name                      string
	cloud, cloudAlt           int
	gripNil                   int
	specialNil                int
	specialFalse              int
	specialTrue               int
}

func C_game32SafePaths() (out game32SafePathsResult) {
	C.sub_4D11A0()
	out.listContains = int(C.sub_4D12A0(7))
	C.sub_4D1210(7)
	out.listRemoved = int(uintptr(unsafe.Pointer(C.sub_4D1250(7))))
	C.sub_4D11D0()

	points := C.calloc(8, 4)
	defer C.free(points)
	values := []uint32{4, 1, 1, 4, 7, 4, 4, 7}
	for i, value := range values {
		*(*uint32)(unsafe.Pointer(uintptr(points) + uintptr(4*i))) = value
	}
	C.nox_xxx_tileInitdataClear_4D3C50(points)
	C.sub_4D3C80((*C.uint32_t)(points))
	for i := range out.sorted {
		out.sorted[i] = *(*uint32)(unsafe.Pointer(uintptr(points) + uintptr(4*i)))
	}

	dst := C.calloc(2, 4)
	defer C.free(dst)
	out.transformNil = int(C.sub_4D3E30(nil, (*C.float)(dst)))
	src := [2]C.float{0, 6000}
	C.sub_4D3E30(&src[0], (*C.float)(dst))
	out.transform = [2]float32{
		*(*float32)(dst),
		*(*float32)(unsafe.Pointer(uintptr(dst) + 4)),
	}
	for i := -1; i <= 9; i++ {
		out.directions = append(out.directions, int(C.sub_4D3FF0(C.int(i))))
	}
	name := C.CString("Synthetic")
	defer C.free(unsafe.Pointer(name))
	out.nameLen = uint32(C.sub_4D42E0(name))
	out.name = C.GoString((*C.char)(C.nox_xxx_getRandMapName_4D4310()))

	out.cloud = int(C.nox_xxx_updDrawCloud_4CE1D0(0, 0))
	out.cloudAlt = int(C.sub_4CE360(0, 0))
	C.nox_xxx_effectSpeedEngage_4DFC30(0, 0)
	C.nox_xxx_effectSpeedDisengage_4DFCA0(0, 0)
	C.sub_4E0170(0, 0)
	C.nox_xxx_effectRegeneration_4E01D0(0, 0)
	C.nox_xxx_attribContinualReplen_4E02C0(0, nil)
	out.gripNil = int(C.nox_xxx_unusedCheckGripEffect_4E03F0(0, 0, 0, 0))
	target := C.calloc(1, 16)
	defer C.free(target)
	C.nox_xxx_stunEffect_4E04D0(0, 0, 0, C.int(uintptr(target)))
	C.nox_xxx_recoilEffect_4E0640(0, 0, 0, 0)
	C.nox_xxx_confuseEffect_4E0670(0, 0, 0, C.int(uintptr(target)))
	C.nox_xxx_lightngEffect_4E06F0(0, 0, 0, 0)
	damage := C.int(10)
	C.nox_xxx_drainMEffect_4E0740(0, 0, 0, 0, &damage)
	C.nox_xxx_vampirismEffect_4E07C0(0, 0, 0, 0, &damage)
	C.nox_xxx_poisonEffect_4E0850(0, 0, 0, C.int(uintptr(target)))
	C.nox_xxx_sympathyEffect_4E08E0(0, 0, 0, 0, &damage)
	out.specialNil = int(C.sub_4E1470(0))
	unit := C.calloc(1, 16)
	defer C.free(unit)
	out.specialFalse = int(C.sub_4E1470(C.int(uintptr(unit))))
	*(*uint32)(unsafe.Pointer(uintptr(unit) + 8)) = 0x1000000
	*(*uint32)(unsafe.Pointer(uintptr(unit) + 12)) = 0x4000
	out.specialTrue = int(C.sub_4E1470(C.int(uintptr(unit))))
	return out
}

func C_nox_xxx_effectDamageMultiplier_4E04C0(factor float32, initial float32) float32 {
	buf1 := C.malloc(64)
	defer C.free(buf1)
	*(*float32)(unsafe.Pointer(uintptr(buf1) + 44)) = factor
	buf5 := C.malloc(4)
	defer C.free(buf5)
	*(*float32)(unsafe.Pointer(buf5)) = initial
	C.nox_xxx_effectDamageMultiplier_4E04C0(C.int(uintptr(buf1)), 0, 0, 0, (*C.float)(buf5))
	return *(*float32)(unsafe.Pointer(buf5))
}

func C_nox_xxx_effectProjectileSpeed_4E09B0(factor float32, initial float32) float32 {
	buf1 := C.malloc(64)
	defer C.free(buf1)
	*(*float32)(unsafe.Pointer(uintptr(buf1) + 44)) = factor
	buf5 := C.malloc(1024)
	defer C.free(buf5)
	*(*float32)(unsafe.Pointer(uintptr(buf5) + 544)) = initial
	C.nox_xxx_effectProjectileSpeed_4E09B0(C.int(uintptr(buf1)), 0, 0, 0, C.int(uintptr(buf5)))
	return *(*float32)(unsafe.Pointer(uintptr(buf5) + 544))
}

func C_nox_xxx_inversionEffect_4E03D0(val96 uint32) (ret int, out int) {
	buf1 := C.malloc(128)
	defer C.free(buf1)
	*(*uint32)(unsafe.Pointer(uintptr(buf1) + 96)) = val96
	outBuf := C.malloc(4)
	defer C.free(outBuf)
	ret = int(C.nox_xxx_inversionEffect_4E03D0(C.int(uintptr(buf1)), 0, 0, 0, 0, (*C.int)(outBuf)))
	out = int(*(*C.int)(outBuf))
	return
}

func C_nox_xxx_gripEffect_4E0480(val96 uint32) (ret int, out int) {
	buf1 := C.malloc(128)
	defer C.free(buf1)
	*(*uint32)(unsafe.Pointer(uintptr(buf1) + 96)) = val96
	outBuf := C.malloc(4)
	defer C.free(outBuf)
	ret = int(C.nox_xxx_gripEffect_4E0480(C.int(uintptr(buf1)), 0, 0, 0, 0, (*C.int)(outBuf)))
	out = int(*(*C.int)(outBuf))
	return
}
