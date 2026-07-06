package legacy

/*
#include "common__object__armrlook.h"
#include "common__object__weaplook.h"

int nox_xxx_loadGuides_427070(void);

static void* test_load_look(void) {
	return nox_xxx_loadLook_415D50();
}

static void* test_load_modifyers(void) {
	return nox_xxx_loadModifyers_4158C0();
}

static int test_load_guides(void) {
	return nox_xxx_loadGuides_427070();
}
*/
import "C"

func C_testLoadLook()       { C.test_load_look() }
func C_testLoadModifyers()  { C.test_load_modifyers() }
func C_testLoadGuides() int { return int(C.test_load_guides()) }
