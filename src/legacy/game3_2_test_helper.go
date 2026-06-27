package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_4E14A0();
int sub_479D00();
int sub_4866D0(uint32_t* a1, int a2);
int nox_xxx_mapGetTypeMB_4CFFA0(void* a1);
int sub_4CFFC0(int a1);
int sub_4CE340(int a1, int a2);
float* nox_xxx_effectDamageMultiplier_4E04C0(int a1, int a2, int a3, int a4, float* a5);
int nox_xxx_effectProjectileSpeed_4E09B0(int a1, int a2, int a3, int a4, int a5);
int nox_xxx_inversionEffect_4E03D0(int a1, int a2, int a3, int a4, int a5, int* a6);
int nox_xxx_gripEffect_4E0480(int a1, int a2, int a3, int a4, int a5, int* a6);
*/
import "C"
import (
	"unsafe"

	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
)

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
