package legacy

/*
#include <stdint.h>
char* sub_4A0020();
*/
import "C"
import "unsafe"

func C_sub_4A0020() uintptr {
	return uintptr(unsafe.Pointer(C.sub_4A0020()))
}
