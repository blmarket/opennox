package legacy

/*
#include "memmap.h"
*/
import "C"
import "unsafe"

func C_mem_getPtrSize(base uintptr, off uintptr, size uintptr) unsafe.Pointer {
	return C.mem_getPtrSize(C.uintptr_t(base), C.uintptr_t(off), C.uintptr_t(size))
}

func C_mem_getPtr(base uintptr, off uintptr) unsafe.Pointer {
	return C.mem_getPtr(C.uintptr_t(base), C.uintptr_t(off))
}

func C_mem_getU8Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getU8Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getI8Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getI8Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getU16Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getU16Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getI16Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getI16Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getU32Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getU32Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getI32Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getI32Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getU64Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getU64Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getI64Ptr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getI64Ptr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getFloatPtr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getFloatPtr(C.uintptr_t(base), C.uintptr_t(off)))
}

func C_mem_getDoublePtr(base uintptr, off uintptr) unsafe.Pointer {
	return unsafe.Pointer(C.mem_getDoublePtr(C.uintptr_t(base), C.uintptr_t(off)))
}
