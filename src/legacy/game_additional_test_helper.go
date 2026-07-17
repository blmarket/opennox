package legacy

/*
#include <stdint.h>
#include <stdlib.h>

// GAME1.c functions
int sub_409A70(short a1);
char* nox_server_currentMapGetFilename_409B30();
char* nox_xxx_mapGetMapName_409B40();
int sub_409EF0(int a1);
int sub_409F40(int a1);
int nox_xxx_servGetPlrLimit_409FA0();
int sub_40A220();
int sub_40A740();
int sub_40AA00();
int sub_40AA40();
void sub_40E090();
int sub_413920();

// GAME2.c functions
int sub_452010();
int sub_4521F0();
extern uint32_t dword_5d4594_1045432;
*/
import "C"

func C_sub_409A70(a1 int16) int {
	return int(C.sub_409A70(C.short(a1)))
}

func C_nox_server_currentMapGetFilename_409B30() string {
	return C.GoString(C.nox_server_currentMapGetFilename_409B30())
}

func C_nox_xxx_mapGetMapName_409B40() string {
	return C.GoString(C.nox_xxx_mapGetMapName_409B40())
}

func C_sub_409EF0(a1 int) int {
	return int(C.sub_409EF0(C.int(a1)))
}

func C_sub_409F40(a1 int) int {
	return int(C.sub_409F40(C.int(a1)))
}

func C_nox_xxx_servGetPlrLimit_409FA0() int {
	return int(C.nox_xxx_servGetPlrLimit_409FA0())
}

func C_sub_40A220() int {
	return int(C.sub_40A220())
}

func C_sub_40A740() int {
	return int(C.sub_40A740())
}

func C_sub_40AA00() int {
	return int(C.sub_40AA00())
}

func C_sub_40AA40() int {
	return int(C.sub_40AA40())
}

func C_sub_40E090() {
	C.sub_40E090()
}

func C_sub_413920() int {
	return int(C.sub_413920())
}

// GAME2.c
func C_sub_452010() int {
	return int(C.sub_452010())
}

func C_sub_4521F0() int {
	C.dword_5d4594_1045432 = 0
	return int(C.sub_4521F0())
}
