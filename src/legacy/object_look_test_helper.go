package legacy

/*
#include "common__object__armrlook.h"
#include "common__object__weaplook.h"
*/
import "C"
import "unsafe"

func C_nox_xxx_loadLook_415D50() unsafe.Pointer {
	return unsafe.Pointer(C.nox_xxx_loadLook_415D50())
}

func C_nox_xxx_loadModifyers_4158C0() unsafe.Pointer {
	return unsafe.Pointer(C.nox_xxx_loadModifyers_4158C0())
}
