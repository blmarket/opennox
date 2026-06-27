package legacy

/*
#include <stdint.h>
#include <stdlib.h>

int sub_4F2570(int a1);
double nox_xxx_objectGetMass_4E4A70(int a1);
int nox_xxx_inventoryGetFirst_4E7980(int a1);
*/
import "C"
import "unsafe"

// C_sub_4F2570 is pure arithmetic: returns 1 iff 0 < a1 < 6.
func C_sub_4F2570(a1 int) int {
	return int(C.sub_4F2570(C.int(a1)))
}

// C_nox_xxx_objectGetMass_4E4A70 reads float at offset +120 and returns as double.
func C_nox_xxx_objectGetMass_4E4A70(val float32) float64 {
	buf := C.malloc(256)
	defer C.free(buf)
	*(*float32)(unsafe.Pointer(uintptr(buf) + 120)) = val
	return float64(C.nox_xxx_objectGetMass_4E4A70(C.int(uintptr(buf))))
}

// C_nox_xxx_inventoryGetFirst_4E7980 reads uint32 at offset +504.
func C_nox_xxx_inventoryGetFirst_4E7980(val uint32) int {
	buf := C.malloc(1024)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 504)) = val
	return int(C.nox_xxx_inventoryGetFirst_4E7980(C.int(uintptr(buf))))
}
