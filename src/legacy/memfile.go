package legacy

/*
#include "memfile.h"
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
)

var _ = [1]struct{}{}[16-unsafe.Sizeof(binfile.MemFile{})]

type nox_memfile = C.nox_memfile

func asMemfile(p *nox_memfile) *binfile.MemFile {
	return asMemfileP(unsafe.Pointer(p))
}

func asMemfileP(p unsafe.Pointer) *binfile.MemFile {
	return (*binfile.MemFile)(p)
}

//export nox_memfile_read
func nox_memfile_read(dst unsafe.Pointer, sz C.uint32_t, cnt C.int32_t, f *nox_memfile) C.uint32_t {
	n := uint32(sz) * uint32(cnt)
	read, _ := asMemfile(f).Read(unsafe.Slice((*byte)(dst), int(n)))
	return C.uint32_t(uint32(read) / uint32(sz))
}
