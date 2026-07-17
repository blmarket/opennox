package legacy

/*
#include <stdint.h>

extern uint32_t nox_tile_def_cnt;
extern uint32_t dword_5d4594_251572;
*/
import "C"

func C_initializeThingDefinitionGlobals() {
	C.nox_tile_def_cnt = 0
	C.dword_5d4594_251572 = 0
}
