package legacy

/*
#include <stdint.h>
#include <stdlib.h>

#include "defs.h"

int nox_xxx_netAbilityReport_4D8060(int a1, int a2, int a3);
int nox_xxx_abilityRewardServ_4FB9C0_ability(int a1, int a2, int a3);
int nox_xxx_netReportGuideAward_4D8000(int a1, char a2, char a3, int a4);
int nox_xxx_awardBeastGuide_4FAE80_magic_plyrgide(int a1, int a2, int a3);
int nox_xxx_netSendSpellAward_4D7F90(int a1, int a2, char a3, int a4);
int nox_xxx_spellGrantToPlayer_4FB550(nox_object_t* a1, int a2, int a3, int a4, int a5);
uint32_t* sub_56F980(int a1, unsigned char a2);
*/
import "C"
import "unsafe"

func C_nox_xxx_netAbilityReport_4D8060_nonPlayer() int {
	obj := C.calloc(1, 1024)
	defer C.free(obj)
	return int(C.nox_xxx_netAbilityReport_4D8060(C.int(uintptr(obj)), 1, 1))
}

func C_nox_xxx_abilityRewardServ_4FB9C0_nonPlayer() int {
	obj := C.calloc(1, 1024)
	defer C.free(obj)
	return int(C.nox_xxx_abilityRewardServ_4FB9C0_ability(C.int(uintptr(obj)), 1, 1))
}

func C_nox_xxx_netReportGuideAward_4D8000_nonPlayer() int {
	obj := C.calloc(1, 1024)
	defer C.free(obj)
	return int(C.nox_xxx_netReportGuideAward_4D8000(C.int(uintptr(obj)), 1, 1, 1))
}

func C_nox_xxx_awardBeastGuide_4FAE80_nonPlayer() int {
	obj := C.calloc(1, 1024)
	defer C.free(obj)
	return int(C.nox_xxx_awardBeastGuide_4FAE80_magic_plyrgide(C.int(uintptr(obj)), 1, 1))
}

func C_nox_xxx_netSendSpellAward_4D7F90_nonPlayer() int {
	obj := C.calloc(1, 1024)
	defer C.free(obj)
	return int(C.nox_xxx_netSendSpellAward_4D7F90(C.int(uintptr(obj)), 1, 1, 1))
}

func C_nox_xxx_spellGrantToPlayer_4FB550_nonPlayer() int {
	obj := C.calloc(1, 1024)
	defer C.free(obj)
	return int(C.nox_xxx_spellGrantToPlayer_4FB550((*C.nox_object_t)(obj), 1, 1, 1, 1))
}

func C_sub_56F980_belowProtectedRange(v int, delta byte) uintptr {
	return uintptr(unsafe.Pointer(C.sub_56F980(C.int(v), C.uchar(delta))))
}
