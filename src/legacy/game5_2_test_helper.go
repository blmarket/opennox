package legacy

/*
#include <stdint.h>
#include <stdlib.h>

unsigned int nox_xxx_netGetUnitCodeCli_578B00(int a1);
int nox_xxx_netClearHighBit_578B30(short a1);
unsigned int nox_xxx_netTestHighBit_578B70(unsigned int a1);
int nox_xxx_waypointNext_579870(int a1);
int sub_5798A0(int a1);
int sub_57BA10(int a1, short a2, short a3, int a4);
int nox_server_getNextMapGroup_57C090(int a1);
*/
import "C"
import "unsafe"

func C_nox_xxx_netGetUnitCodeCli_578B00_nil() uint32 {
	return uint32(C.nox_xxx_netGetUnitCodeCli_578B00(0))
}

func C_nox_xxx_netGetUnitCodeCli_578B00(code, flags uint32) uint32 {
	buf := C.malloc(132)
	defer C.free(buf)
	base := uintptr(buf)
	*(*uint32)(unsafe.Pointer(base + 112)) = flags
	*(*uint32)(unsafe.Pointer(base + 128)) = code
	return uint32(C.nox_xxx_netGetUnitCodeCli_578B00(C.int(base)))
}

func C_nox_xxx_netClearHighBit_578B30(v int16) int {
	return int(C.nox_xxx_netClearHighBit_578B30(C.short(v)))
}

func C_nox_xxx_netTestHighBit_578B70(v uint32) uint32 {
	return uint32(C.nox_xxx_netTestHighBit_578B70(C.uint(v)))
}

func C_nox_xxx_waypointNext_579870_nil() int {
	return int(C.nox_xxx_waypointNext_579870(0))
}

func C_nox_xxx_waypointNext_579870(next uint32) int {
	buf := C.malloc(488)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 484)) = next
	return int(C.nox_xxx_waypointNext_579870(C.int(uintptr(buf))))
}

func C_sub_5798A0_nil() int {
	return int(C.sub_5798A0(0))
}

func C_sub_5798A0(next uint32) int {
	buf := C.malloc(488)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 484)) = next
	return int(C.sub_5798A0(C.int(uintptr(buf))))
}

func C_sub_57BA10(a2, a3 int16, a4 uint32) (ret uintptr, out0, out2 uint16, out4 uint32) {
	buf := C.malloc(8)
	defer C.free(buf)
	ret = uintptr(C.sub_57BA10(C.int(uintptr(buf)), C.short(a2), C.short(a3), C.int(a4)))
	out0 = *(*uint16)(unsafe.Pointer(uintptr(buf)))
	out2 = *(*uint16)(unsafe.Pointer(uintptr(buf) + 2))
	out4 = *(*uint32)(unsafe.Pointer(uintptr(buf) + 4))
	return
}

func C_nox_server_getNextMapGroup_57C090_nil() int {
	return int(C.nox_server_getNextMapGroup_57C090(0))
}

func C_nox_server_getNextMapGroup_57C090(next uint32) int {
	buf := C.malloc(92)
	defer C.free(buf)
	*(*uint32)(unsafe.Pointer(uintptr(buf) + 88)) = next
	return int(C.nox_server_getNextMapGroup_57C090(C.int(uintptr(buf))))
}
