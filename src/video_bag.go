package opennox

/*
#include "defs.h"
extern void* dword_5d4594_1189592;
extern void* dword_5d4594_1189596;
*/
import "C"

import (
	"encoding/binary"
	"fmt"
	"image"
	"unsafe"

	"github.com/noxworld-dev/opennox-lib/bag"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/noximage/pcx"

	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"github.com/noxworld-dev/opennox/v1/common/alloc/handles"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

var (
	noxVideoBag *bag.File
	noxImages   struct {
		byHandle map[unsafe.Pointer]*Image
		byIndex  []*Image
	}
)

func init() {
	noxImages.byHandle = make(map[unsafe.Pointer]*Image)
}

type nox_video_bag_image_t = C.nox_video_bag_image_t

func asImage(p *nox_video_bag_image_t) *Image {
	return asImageP(unsafe.Pointer(p))
}

func asImageP(p unsafe.Pointer) *Image {
	if p == nil {
		return nil
	}
	img := noxImages.byHandle[p]
	if img == nil {
		err := fmt.Errorf("unexpected image handle: %x", p)
		videoLog.Printf("%v", err)
		if cgoSafe {
			panic(err)
		}
	}
	return img
}

func NewRawImage(typ int, data []byte) *Image {
	return &Image{typ: typ, raw: data, nocgo: true}
}

type Image struct {
	h         unsafe.Pointer
	typ       int
	bag       *bag.ImageRec
	raw       []byte
	override  []byte
	nocgo     bool
	cdata     []byte
	cfree     func()
	field_1_0 uint16
	field_1_1 uint16
}

func (img *Image) String() string {
	if img == nil {
		return "<nil>"
	}
	if img.override != nil {
		return fmt.Sprintf("{type=%d, override=[%d]}", img.Type(), len(img.override))
	}
	if img.bag != nil {
		return fmt.Sprintf("{type=%d, idx=%d, data=[%d]}", img.Type(), img.bag.Index, len(img.raw))
	}
	return fmt.Sprintf("{type=%d, raw=[%d]}", img.Type(), len(img.raw))
}

func (img *Image) C() *nox_video_bag_image_t {
	if img == nil {
		return nil
	}
	if img.nocgo {
		panic("image not allowed in cgo context")
	}
	if img.h == nil {
		img.h = handles.NewPtr()
		noxImages.byHandle[img.h] = img
	}
	return (*nox_video_bag_image_t)(img.h)
}

func (img *Image) Type() int {
	if img.bag != nil {
		return int(img.bag.Type)
	}
	return img.typ
}

func (img *Image) loadOverride() []byte {
	if img == nil || img.raw != nil {
		return nil
	}
	switch img.Type() {
	default:
		return nil
	case 3, 4, 5, 6:
	}
	if img.override != nil {
		return img.override
	}
	sect := int(img.bag.SegmInd)
	offs := int(img.bag.Offset)

	im, err := imageByBagSection(sect, offs)
	if err != nil {
		log.Println(err)
		return nil
	} else if im == nil {
		return nil
	}
	img.override = pcx.Encode(im)
	return img.override
}

func (img *Image) Pixdata() []byte {
	if img == nil {
		return nil
	}
	if img.cdata != nil {
		return img.cdata
	}
	data := img.loadOverride()
	if data == nil {
		data = img.bagPixdata()
	}
	if len(data) == 0 {
		panic("cannot load")
	}
	// TODO: remove interning when we get rid of C renderer
	img.cdata, img.cfree = alloc.CloneSlice(data)
	return img.cdata
}

func (img *Image) Meta() (off, sz image.Point, ok bool) {
	pix := img.Pixdata()
	if len(pix) < 8 {
		ok = false
		return
	}
	sz.X = int(binary.LittleEndian.Uint32(pix[0:]))
	sz.Y = int(binary.LittleEndian.Uint32(pix[4:]))
	if len(pix) < 16 {
		ok = false
		return
	}
	ok = true
	off.X = int(binary.LittleEndian.Uint32(pix[8:]))
	off.Y = int(binary.LittleEndian.Uint32(pix[12:]))
	return
}

func readVideobag(path string) error {
	f, err := bag.Open(path)
	if err != nil {
		return err
	}
	imgs, err := f.Images()
	if err != nil {
		_ = f.Close()
		return err
	}
	noxVideoBag = f
	noxImages.byIndex = make([]*Image, 0, len(imgs))
	for _, img := range imgs {
		noxImages.byIndex = append(noxImages.byIndex, &Image{bag: img})
	}
	return nil
}

func ReadVideoBag() error {
	return readVideobag("video.bag")
}

func (img *Image) bagPixdata() []byte { // nox_video_getImagePixdata_42FB30
	if img == nil {
		return nil
	}
	if img.Type()&0x3F == 7 {
		return nil
	}
	if img.Type()&0x80 != 0 {
		panic("unreachable")
	}
	if img.raw != nil {
		return img.raw
	}
	data, err := img.bag.Raw()
	if err != nil {
		panic(err)
	}
	img.raw = data
	return data
}

//export nox_video_bag_image_type
func nox_video_bag_image_type(img *nox_video_bag_image_t) C.int {
	return C.int(asImage(img).Type())
}

//export nox_xxx_readImgMB_42FAA0
func nox_xxx_readImgMB_42FAA0(known_idx C.int, typ C.char, cname2 *C.char) *nox_video_bag_image_t {
	if known_idx != -1 {
		return bagImageByIndex(int(known_idx)).C()
	}
	return nox_xxx_readImgMB42FAA0(int(known_idx), byte(typ), GoString(cname2)).C()
}

func nox_xxx_readImgMB42FAA0(ind int, typ byte, name2 string) *Image {
	if ind != -1 {
		return bagImageByIndex(ind)
	}
	log.Printf("nox_xxx_readImgMB42FAA0(%d, %d, %q)", ind, int(typ), name2)
	return nox_xxx_loadImage_47A8C0(typ, name2)
}

//export nox_xxx_gLoadImg_42F970
func nox_xxx_gLoadImg_42F970(name *C.char) *nox_video_bag_image_t {
	return nox_xxx_gLoadImg(GoString(name)).C()
}

func bagImageByIndex(ind int) *Image {
	return noxImages.byIndex[ind]
}

func nox_xxx_loadImage_47A8C0(typ byte, name string) *Image {
	// TODO: this one is supposed to load PCX images from FS
	panic("TODO: read PCX from FS")
}

//export nox_xxx_gLoadAnim_42FA20
func nox_xxx_gLoadAnim_42FA20(name *C.char) *C.nox_things_imageRef_t {
	return nox_xxx_gLoadAnim(GoString(name)).C()
}

func nox_xxx_gLoadImg(name string) *Image {
	if name == "" {
		return nil
	}
	for _, p := range nox_images_arr1_787156 {
		if name == p.Name() {
			ind := p.field24int()
			if ind == -1 {
				if p.field_25_0 == -1 {
					return nil
				}
				name2 := p.Name2()
				return nox_xxx_loadImage_47A8C0(byte(p.field_25_0), name2)
			}
			return noxImages.byIndex[ind]
		}
	}
	return nil
}

func asImageRefP(p unsafe.Pointer) *noxImageRef {
	return (*noxImageRef)(p)
}

func asImageRef(p *C.nox_things_imageRef_t) *noxImageRef {
	return asImageRefP(unsafe.Pointer(p))
}

type noxImageRef C.nox_things_imageRef_t

func (r *noxImageRef) C() *C.nox_things_imageRef_t {
	return (*C.nox_things_imageRef_t)(unsafe.Pointer(r))
}

func (r *noxImageRef) Name() string {
	return GoString(&r.name[0])
}

func (r *noxImageRef) Name2() string {
	return GoString(&r.name2[0])
}

func (r *noxImageRef) kind() int {
	return int(r.ref_kind)
}

func (r *noxImageRef) field24int() int {
	if r.kind() == 2 {
		panic("not an regular image")
	}
	return int(uintptr(r.field_24))
}

func (r *noxImageRef) field24ptr() *noxImageRefAnim {
	if r.kind() != 2 {
		panic("not an animation")
	}
	return (*noxImageRefAnim)(r.field_24)
}

type noxImageRefAnim C.nox_things_imageRef2_t

func (r *noxImageRefAnim) C() *C.nox_things_imageRef2_t {
	return (*C.nox_things_imageRef2_t)(unsafe.Pointer(r))
}

func (r *noxImageRefAnim) Images() []*nox_video_bag_image_t {
	return unsafe.Slice(r.images, r.images_sz)
}

func nox_xxx_gLoadAnim(name string) *noxImageRef {
	if name == "" {
		return nil
	}
	for _, p := range nox_images_arr1_787156 {
		if name == p.Name() {
			return p
		}
	}
	return nil
}

func nox_video_bagFree_42F4D0() {
	for _, p := range nox_images_arr1_787156 {
		if p.kind() == 2 {
			anim := p.field24ptr()
			alloc.Free(unsafe.Pointer(anim.images))
			alloc.Free(unsafe.Pointer(anim.C()))
		}
		alloc.Free(unsafe.Pointer(p.C()))
	}
	nox_images_arr1_787156 = nil
	sub_47D150()
	noxVideoBag.Close()
	noxVideoBag = nil
	for _, img := range noxImages.byIndex {
		if img.cdata != nil {
			img.cfree()
			img.cdata = nil
			img.cfree = nil
		}
	}
	noxImages.byIndex = nil
	noxImages.byHandle = make(map[unsafe.Pointer]*Image)
}

func sub_47D150() {
	if p := memmap.PtrPtr(0x5D4594, 1189588); *p != nil {
		alloc.Free(*p)
		*p = nil
	}
	if C.dword_5d4594_1189592 != nil {
		alloc.Free(C.dword_5d4594_1189592)
		C.dword_5d4594_1189592 = nil
	}
	if C.dword_5d4594_1189596 != nil {
		alloc.Free(C.dword_5d4594_1189596)
		C.dword_5d4594_1189596 = nil
	}
}
