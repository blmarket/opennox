package legacy

/*
#include <stdint.h>

char* nox_xxx_getRandomName_4358A0(void);
*/
import "C"

import (
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func C_clientRandomNames(count int) (names []string, initialized uint32) {
	slot := memmap.PtrUint32(0x5D4594, 814516)
	saved := *slot
	*slot = 0
	defer func() { *slot = saved }()

	for i := 0; i < count; i++ {
		names = append(names, C.GoString(C.nox_xxx_getRandomName_4358A0()))
	}
	return names, *slot
}
