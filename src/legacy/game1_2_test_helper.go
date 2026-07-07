package legacy

/*
#include <stdint.h>

int sub_43AF40();
extern uint32_t nox_game_createOrJoin_815048;
*/
import "C"

func C_sub_43AF40() int {
	return int(C.sub_43AF40())
}

func C_game1_2_setCreateOrJoin(v uint32) {
	C.nox_game_createOrJoin_815048 = C.uint32_t(v)
}

func C_game1_2_getCreateOrJoin() uint32 {
	return uint32(C.nox_game_createOrJoin_815048)
}
