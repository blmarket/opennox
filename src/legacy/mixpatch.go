package legacy

/*
#include "defs.h"
#include "MixPatch.h"
*/
import "C"
import "unsafe"

//export sub_980523
func sub_980523(unit *C.nox_object_t) {
	if unit == nil {
		return
	}
	for it := unit.inv_first_item; it != nil; it = it.inv_next_item {
		if (it.obj_class&0x2000000) != 0 && (it.obj_flags&0x100) != 0 {
			if nox_xxx_unitArmorInventoryEquipFlags_415C70(it)&0x3000000 != 0 {
				// *(uint32_t*)(*(uint32_t*)(((uint32_t)unit->data_update) + 276) + 2500) = it;
				dataUpdate := unit.data_update
				if dataUpdate == nil {
					continue
				}
				p1 := *(*uintptr)(unsafe.Pointer(uintptr(dataUpdate) + 276))
				if p1 == 0 {
					continue
				}
				*(*uintptr)(unsafe.Pointer(p1 + 2500)) = uintptr(unsafe.Pointer(it))
			}
		}
	}
}

//export sub_9805EB
func sub_9805EB(unit *C.nox_object_t) *C.nox_object_t {
	if unit == nil {
		return nil
	}
	for it := unit.inv_first_item; it != nil; it = it.inv_next_item {
		if (it.obj_class&0x2000000) != 0 && it.obj_flags == 16 {
			return it
		}
	}
	return nil
}
