package legacy

/*
#include "GameEx.h"

int nox_CharToOemW(const wchar2_t* pSrc, char* pDst);
char playerInfoStructParser_0(void* a1p);
char playerInfoStructParser_1(void* a1p, int* a3);
char mix_MouseKeyboardWeaponRoll(nox_object_t* playerObjP, char a2);
*/
import "C"
import "unsafe"

func C_nox_CharToOemW(pSrc *uint16, pDst *byte) int {
	return int(C.nox_CharToOemW((*C.wchar2_t)(unsafe.Pointer(pSrc)), (*C.char)(unsafe.Pointer(pDst))))
}

func C_getPlayerClassFromObjPtr(a1 int) byte {
	return byte(C.getPlayerClassFromObjPtr(C.int(a1)))
}

func C_getFlagValueFromFlagIndex(a1 int) int {
	return int(C.getFlagValueFromFlagIndex(C.int(a1)))
}

func C_playerDropATrap(playerObj int) byte {
	return byte(C.playerDropATrap(C.int(playerObj)))
}

func C_playerInfoStructParser0Sentinel() byte {
	return byte(C.playerInfoStructParser_0(unsafe.Pointer(^uintptr(1))))
}

func C_playerInfoStructParser1Sentinel() byte {
	return byte(C.playerInfoStructParser_1(unsafe.Pointer(^uintptr(1)), nil))
}

func C_mix_MouseKeyboardWeaponRoll(playerObj unsafe.Pointer, forward bool) byte {
	a2 := C.char(0)
	if forward {
		a2 = 1
	}
	return byte(C.mix_MouseKeyboardWeaponRoll((*C.nox_object_t)(playerObj), a2))
}
