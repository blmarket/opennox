package legacy

import (
	"testing"
)

func TestSimple(t *testing.T) {
	// buf := (*[7]uintptr)(NewStack(3, 4))
	// must.Eq(t, buf[0], uintptr(unsafe.Pointer(&buf[1])))
	// must.Eq(t, buf[1], uintptr(unsafe.Pointer(&buf[3])))
	// must.Eq(t, buf[2], 0)
	// must.Eq(t, buf[3], uintptr(unsafe.Pointer(&buf[5])))
	// must.Eq(t, buf[4], 0)
	// must.Eq(t, buf[5], 0)
	// must.Eq(t, buf[6], 0)

	// ptr1 := StackPop(unsafe.Pointer(buf))
	// must.Eq(t, ptr1, unsafe.Pointer(&buf[2]))
	// must.Eq(t, buf[0], uintptr(unsafe.Pointer(&buf[3])))

	// ptr2 := StackPop(unsafe.Pointer(buf))
	// must.Eq(t, ptr2, unsafe.Pointer(&buf[4]))
	// must.Eq(t, buf[0], uintptr(unsafe.Pointer(&buf[5])))

	// ptr3 := StackPop(unsafe.Pointer(buf))
	// must.Eq(t, ptr3, unsafe.Pointer(&buf[6]))
	// must.Eq(t, buf[0], 0)

	// ptr4 := StackPop(unsafe.Pointer(buf))
	// must.Eq(t, ptr4, unsafe.Pointer(nil))
	// must.Eq(t, buf[0], 0)
}
