package legacy

/*
#include <stdint.h>
#include <stdlib.h>
char* sub_4A0020();
int sub_48D4F0(unsigned short a1, unsigned short a2);
void sub_48C580(void* a1, int num);
*/
import "C"
import "unsafe"

func C_sub_4A0020() uintptr {
	return uintptr(unsafe.Pointer(C.sub_4A0020()))
}

func C_sub_48D4F0(a1, a2 uint16) int {
	return int(C.sub_48D4F0(C.ushort(a1), C.ushort(a2)))
}

func C_sub_48C580(pixels []uint32) {
	if len(pixels) == 0 {
		return
	}
	C.sub_48C580(unsafe.Pointer(&pixels[0]), C.int(len(pixels)))
}
