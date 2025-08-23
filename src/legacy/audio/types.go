package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

var _ = [1]struct{}{}[36-unsafe.Sizeof(AudioStructYyy{})]

type AudioStructYyy struct {
	Field0  [32]byte
	Field32 uint32
}

func (p *AudioStructYyy) C() unsafe.Pointer {
	return unsafe.Pointer(p)
}

type AudioStruct1 struct {
	Field0  uint16
	Field2  uint16
	Field4  uint32
	Field8  uint32
	Field12 uint16
	Field14 uint16
}

// FILE is defined as void* to avoid circular imports
type FILE unsafe.Pointer

var _ = [1]struct{}{}[288-unsafe.Sizeof(AudioStructXxx{})]

type AudioStructXxx struct {
	Arr0       *AudioStructYyy
	Size4      uint
	Path2_8    [260]byte
	Bagfile268 FILE   // 67, 268
	Field272   FILE   // 68, 272
	Field276   uint32 // 69, 276
	Field280   uint32 // 70, 280
	Field284   int32  // 71, 284
}

func (p *AudioStructXxx) C() unsafe.Pointer {
	return unsafe.Pointer(p)
}

func (p *AudioStructXxx) Free(m *AudioModule) {
	if p == nil {
		return
	}
	if p.Bagfile268 != nil {
		m.nox_fs_close(unsafe.Pointer(p.Bagfile268))
		p.Bagfile268 = nil
	}
	if p.Field272 != nil {
		m.nox_fs_close(unsafe.Pointer(p.Field272))
		p.Field272 = nil
	}
	if p.Arr0 != nil {
		alloc.Free(p.Arr0)
		p.Arr0 = nil
	}
	alloc.Free(p)
}
