package legacy

/*
#include "memfile.h"
*/
import "C"
import (
	"encoding/binary"
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

//export nox_memfile_read_u32
func nox_memfile_read_u32(f *nox_memfile) C.uint32_t {
	return C.uint32_t(asMemfile(f).ReadU32())
}

//export nox_memfile_read
func nox_memfile_read(dst unsafe.Pointer, sz C.uint32_t, cnt C.int32_t, f *nox_memfile) C.uint32_t {
	n := uint32(sz) * uint32(cnt)
	read, _ := asMemfile(f).Read(unsafe.Slice((*byte)(dst), int(n)))
	return C.uint32_t(uint32(read) / uint32(sz))
}

//export nox_memfile_read64align_40AD60
func nox_memfile_read64align_40AD60(dst *C.char, sz, cnt C.int32_t, f *nox_memfile) C.uint32_t {
	mf := asMemfile(f)
	off := uintptr(unsafe.Pointer(f.cur)) - uintptr(unsafe.Pointer(f.data))
	padding := (8 - off%8) % 8
	complete := len(mf.Data()) >= int(padding)+8
	v := mf.ReadU64Align()
	if !complete {
		return 0
	}

	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], v)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(dst)), int(sz)*int(cnt)), buf[:])
	return 1
}
