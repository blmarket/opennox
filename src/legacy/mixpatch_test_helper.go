package legacy

/*
#include "MixPatch.h"
#include <stdlib.h>
#include <string.h>
*/
import "C"
import "unsafe"

func C_sub_980523(unit unsafe.Pointer) {
	C.sub_980523((*C.nox_object_t)(unit))
}

func C_sub_9805EB(unit unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.sub_9805EB((*C.nox_object_t)(unit)))
}

func AllocNoxObject() unsafe.Pointer {
	ptr := C.malloc(C.sizeof_nox_object_t)
	C.memset(ptr, 0, C.sizeof_nox_object_t)
	return ptr
}

func FreeNoxObject(p unsafe.Pointer) {
	C.free(p)
}

func AllocBuffer(sz int) unsafe.Pointer {
	ptr := C.malloc(C.size_t(sz))
	C.memset(ptr, 0, C.size_t(sz))
	return ptr
}

func FreeBuffer(p unsafe.Pointer) {
	C.free(p)
}

func SetNoxObjectFields(p unsafe.Pointer, objClass uint32, objFlags uint32, typInd uint16) {
	obj := (*C.nox_object_t)(p)
	obj.obj_class = C.uint(objClass)
	obj.obj_flags = C.uint(objFlags)
	obj.typ_ind = C.ushort(typInd)
}

func SetNoxObjectInventory(p unsafe.Pointer, firstItem unsafe.Pointer) {
	obj := (*C.nox_object_t)(p)
	obj.inv_first_item = (*C.nox_object_t)(firstItem)
}

func SetNoxObjectNextItem(p unsafe.Pointer, nextItem unsafe.Pointer) {
	obj := (*C.nox_object_t)(p)
	obj.inv_next_item = (*C.nox_object_t)(nextItem)
}

func SetNoxObjectDataUpdate(p unsafe.Pointer, dataUpdate unsafe.Pointer) {
	obj := (*C.nox_object_t)(p)
	obj.data_update = dataUpdate
}
