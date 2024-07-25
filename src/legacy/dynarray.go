package legacy

/*
#include "dynarray.h"
*/
import "C"
import (
	"unsafe"
)

type DynArrayElem[T any] struct {
	next  *DynArrayElem[T]
	value T
}

type DynArray[T any] struct {
	head     *DynArrayElem[T]
	elements []DynArrayElem[T]
}

func CreateRaw(a, b int) unsafe.Pointer {
	return unsafe.Pointer(C.sub_4BD280(C.int(a), C.int(b)))
}

// func NewDynarray[T any](a int) *DynArray[T] {
// 	return (*DynArray[T])(CreateRaw(a, unsafe.Sizeof(T)))
// }
