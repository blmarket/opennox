package legacy

/*
#include <stdint.h>
*/
import "C"

//export sub_48D4F0
func sub_48D4F0(a1, a2 C.ushort) C.int {
	v2 := uint16(10000)
	if int(a1)-10000 < 0 {
		if uint16(a2) >= 0xFFFF-uint16(10000-uint16(a1)) {
			return 1
		}
		v2 = uint16(a1)
	}
	if uint16(a2) < uint16(a1) && uint16(a2) >= uint16(a1)-v2 {
		return 1
	}
	return 0
}
