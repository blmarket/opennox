package legacy

import (
	"testing"
	"unsafe"

	"github.com/shoenig/test/must"
)

func TestSimple(t *testing.T) {
	buf := (*[7]uintptr)(CreateRaw(3, 4))
	must.Eq(t, buf[0], uintptr(unsafe.Pointer(&buf[1])))
	must.Eq(t, buf[1], uintptr(unsafe.Pointer(&buf[3])))
	must.Eq(t, buf[2], 0)
	must.Eq(t, buf[3], uintptr(unsafe.Pointer(&buf[5])))
	must.Eq(t, buf[4], 0)
	must.Eq(t, buf[5], 0)
	must.Eq(t, buf[6], 0)
}
