package legacy

/*
#include "client__draw__basedraw.h"
#include "client__draw__weapondraw.h"
*/
import "C"

//export nox_thing_base_draw
func nox_thing_base_draw(a1 *C.int, dr *C.nox_drawable) C.int {
	C.nox_thing_weapon_draw(a1, dr)
	return 1
}
