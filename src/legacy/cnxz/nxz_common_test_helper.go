package cnxz

/*
#include <stdint.h>
#include <stdlib.h>

uint32_t* sub57DD90(uint32_t* this);
void sub57DDC0(void** this);
*/
import "C"
import "unsafe"

// C_sub57DD90 allocates a 0x224 byte buffer, stores the pointer in *this, memset to 0, and returns this.
func C_sub57DD90() (thisPtr uintptr, bufPtr uintptr) {
	var this uint32
	C.sub57DD90((*C.uint32_t)(unsafe.Pointer(&this)))
	thisPtr = uintptr(this)
	bufPtr = uintptr(this)
	return
}

// C_sub57DDC0 frees the buffer pointed to by *this.
func C_sub57DDC0(bufPtr uintptr) {
	p := unsafe.Pointer(bufPtr)
	C.sub57DDC0((*unsafe.Pointer)(unsafe.Pointer(&p)))
}
