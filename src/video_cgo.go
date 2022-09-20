package opennox

/*
#include "GAME1_2.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME2_2.h"
#include "GAME2_3.h"
#include "GAME3.h"
#include "GAME3_1.h"
#include "common__magic__speltree.h"
extern unsigned int dword_5d4594_1193188;
extern unsigned int dword_5d4594_1305748;
extern unsigned int dword_5d4594_3799468;
extern int dword_5d4594_3799524;
extern int nox_win_width;
extern int nox_win_height;

extern uint8_t** nox_pixbuffer_rows_3798784;
extern uint32_t dword_5d4594_823776;

extern uint32_t nox_color_white_2523948;
extern uint32_t nox_color_red_2589776;
extern uint32_t nox_color_blue_2650684;
extern uint32_t nox_color_green_2614268;
extern uint32_t nox_color_cyan_2649820;
extern uint32_t nox_color_yellow_2589772;
extern uint32_t nox_color_violet_2598268;
extern uint32_t nox_color_black_2650656;
extern uint32_t nox_color_orange_2614256;

extern nox_render_data_t* nox_draw_curDrawData_3799572;
*/
import "C"
import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"unsafe"

	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/noximage"
	"github.com/noxworld-dev/opennox-lib/player"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

var (
	nox_win_width                          int
	nox_win_height                         int
	nox_pixbuffer_3798788_arr              []byte
	nox_client_spellDragnDrop_1097192      uint32
	nox_client_spellDragnDrop_type_1097196 int
	nox_client_itemDragnDrop_1097188       *Drawable
	dword_5d4594_1097204                   int
	dword_5d4594_805860                    bool
	nox_enable_threads                     = true
	noxPixBuffer                           struct {
		img      *noximage.Image16
		free     func()
		rows     []*uint16
		freeRows func()
		onResize []func(sz image.Point)
	}
	dword_5d4594_1311936 bool
	func_5d4594_1311924  func()
	dword_5d4594_1193672 bool

	nox_win_width_game  = noxDefaultWidth
	nox_win_height_game = noxDefaultHeight

	nox_color_black_2650656  = noxcolor.RGB5551Color(0, 0, 0)
	nox_color_white_2523948  = noxcolor.RGB5551Color(255, 255, 255)
	nox_color_violet_2598268 = noxcolor.RGB5551Color(100, 0, 0)
	nox_color_red_2589776    = noxcolor.RGB5551Color(255, 128, 128)
	nox_color_green_2614268  = noxcolor.RGB5551Color(128, 255, 128)
	nox_color_cyan_2649820   = noxcolor.RGB5551Color(0, 0, 255)
	nox_color_blue_2650684   = noxcolor.RGB5551Color(0, 160, 255)
	nox_color_orange_2614256 = noxcolor.RGB5551Color(240, 180, 42)
	nox_color_yellow_2589772 = noxcolor.RGB5551Color(255, 255, 0)
	nox_color_gray1          = noxcolor.RGB5551Color(8, 8, 8)
	nox_color_gray2          = noxcolor.RGB5551Color(115, 115, 115)
	nox_color_gray3          = noxcolor.RGB5551Color(212, 212, 212)
	nox_color_red            = noxcolor.RGB5551Color(255, 0, 0)
	nox_color_darkGreen      = noxcolor.RGB5551Color(0, 100, 0)
	nox_color_green          = noxcolor.RGB5551Color(0, 255, 0)
	nox_color_darkBlue       = noxcolor.RGB5551Color(0, 0, 140)
	nox_color_darkYellow     = noxcolor.RGB5551Color(255, 255, 128)
	drawColorPurple          = noxcolor.RGB5551Color(255, 0, 255)
	drawColorDarkPurple      = noxcolor.RGB5551Color(255, 180, 255)
)

func sub_48B3E0(v bool) bool {
	prev := dword_5d4594_1193672
	dword_5d4594_1193672 = v
	return prev
}

//export nox_video_getCutSize_4766D0
func nox_video_getCutSize_4766D0() C.int {
	return C.int(nox_video_getCutSize())
}

//export nox_video_setCutSize_4766A0
func nox_video_setCutSize_4766A0(v C.int) {
	nox_video_setCutSize(int(v))
}

func OnPixBufferResize(fnc func(sz image.Point)) {
	noxPixBuffer.onResize = append(noxPixBuffer.onResize, fnc)
}

func videoGetWindowSize() image.Point {
	return image.Point{
		X: nox_win_width,
		Y: nox_win_height,
	}
}

func videoSetWindowSize(sz image.Point) {
	C.nox_win_width = C.int(sz.X)
	C.nox_win_height = C.int(sz.Y)
	nox_win_width = sz.X
	nox_win_height = sz.Y
}

func cfgUpdateFullScreen() {
	g_fullscreen_cfg = noxClient.getWindowMode()
}

//export nox_video_setGammaSlider
func nox_video_setGammaSlider(v C.int) {
	setGammaSliderOpts(int(v))
}

//export sub_43BE50_get_video_mode_id
func sub_43BE50_get_video_mode_id() C.int {
	return C.int(videoModeID())
}

//export get_video_mode_string
func get_video_mode_string(cid C.int) *C.wchar_t {
	id := int(cid)
	if id < 0 || id >= len(noxVideoModes) {
		return internWStr("custom")
	}
	mode := noxVideoModes[id]
	return internWStr(fmt.Sprintf("%dx%d", mode.X, mode.Y))
}

//export nox_getBackbufWidth
func nox_getBackbufWidth() C.int {
	dx := noxClient.r.PixBufferRect().Dx()
	return C.int(dx)
}

//export nox_getBackbufHeight
func nox_getBackbufHeight() C.int {
	dy := noxClient.r.PixBufferRect().Dy()
	return C.int(dy)
}

//export nox_video_getFullScreen
func nox_video_getFullScreen() C.int {
	return C.int(noxClient.getWindowMode())
}

//export nox_video_setFullScreen
func nox_video_setFullScreen(v C.int) {
	noxClient.updateFullScreen(int(v))
}

//export sub_430C30_set_video_max
func sub_430C30_set_video_max(w, h C.int) {
	videoSetMaxSize(image.Point{X: int(w), Y: int(h)})
}

//export nox_xxx_screenGetSize_430C50_get_video_max
func nox_xxx_screenGetSize_430C50_get_video_max(pw, ph *C.int) {
	sz := videoGetMaxSize()
	*pw = C.int(sz.X)
	*ph = C.int(sz.Y)
}

func (c *Client) videoGetGameMode() image.Point {
	return image.Point{
		X: nox_win_width_game,
		Y: nox_win_height_game,
	}
}

func (c *Client) videoSetGameMode(mode image.Point) {
	nox_win_width_game = mode.X
	nox_win_height_game = mode.Y
	c.setScreenSize(mode)
}

func nox_video_setBackBufferCopyFunc_4AD100() error {
	if nox_video_renderTargetFlags&0x40 != 0 {
		return errors.New("nox_video_setBackBufferCopyFunc_4AD100: flag not implemented")
	} else {
		nox_video_setBackBufferCopyFunc2_4AD150()
	}
	*memmap.PtrUint32(0x973A20, 496) = 0
	return nil
}

func nox_video_setBackBufferCopyFunc2_4AD150() {
	if nox_video_renderTargetFlags&0x40 != 0 {
		panic("not implemented")
	}
}

//export nox_video_callCopyBackBuffer_4AD170
func nox_video_callCopyBackBuffer_4AD170() {
	noxClient.copyPixBuffer()
}

var (
	videoInitDone = false
	renderData1   *RenderData
	renderData2   *RenderData
)

func videoInit(sz image.Point, flags int) error {
	C.dword_5d4594_823776 = 0
	if renderData1 == nil {
		renderData1, _ = alloc.New(RenderData{})
		renderData2, _ = alloc.New(RenderData{})
	}
	noxClient.r.SetData(renderData1)
	C.nox_draw_curDrawData_3799572 = noxClient.r.Data().C()
	if err := drawInitAll(sz, flags); err != nil {
		videoLog.Println("init:", err)
		return err
	}
	noxClient.r.SetData(renderData2)
	C.nox_draw_curDrawData_3799572 = noxClient.r.Data().C()
	*renderData2 = *renderData1
	C.dword_5d4594_823776 = 1
	videoInitDone = true
	return nil
}

func videoInitStub() {
	noxClient.r.SetData(renderData2)
	C.nox_draw_curDrawData_3799572 = noxClient.r.Data().C()
	C.dword_5d4594_823776 = 1
	C.nox_win_width = noxDefaultWidth
	C.nox_win_height = noxDefaultHeight
}

func drawInitAll(sz image.Point, flags int) error {
	if err := nox_client_drawXxx_444AC0(sz.X, sz.Y, flags); err != nil {
		return err
	}
	sub_47D200()
	nox_video_initPixbuffer_486090(sz)
	sub_49F610(sz)
	if res := sub_4338D0(); res == 0 {
		return errors.New("sub_4338D0 failed")
	}
	if err := nox_video_setBackBufferCopyFunc_4AD100(); err != nil {
		return err
	}
	noxClient.r.initParticles()
	sub_4B02D0()
	noxClient.r.partfx.Init(noxClient.r)
	sub_4AE520()
	if err := loadGameFonts(); err != nil {
		return err
	}
	noxClient.r.ClearPoints()
	return nil
}

func sub_4B02D0() {
	dword_5d4594_1311936 = false
	func_5d4594_1311924 = nil
	movieFilesStackCur = 0
	*memmap.PtrUint32(0x5D4594, 1311932) = 0
}

func sub4B0640(fnc func()) {
	func_5d4594_1311924 = fnc
}

func sub_4B05D0() {
	if dword_5d4594_1311936 {
		dword_5d4594_1311936 = false
		movieFilesStackCur = 0
		if func_5d4594_1311924 != nil {
			nox_client_clearScreen_440900()
			func_5d4594_1311924()
		}
	}
}

func gameUpdateVideoMode(inMenu bool) error {
	return noxClient.gameResetVideoMode(inMenu, false)
}

func recreateBuffersAndTarget(sz image.Point) error {
	nox_video_freeFloorBuffer_430EC0()
	if err := recreateRenderTarget(sz); err != nil {
		videoLog.Println("recreate render target:", err)
		return err
	}
	videoLog.Println("recreate render target: ok")
	if err := nox_video_initFloorBuffer_430BA0(sz); err != nil {
		return err
	}
	return nil
}

func recreateRenderTarget(sz image.Point) error {
	flags := uint(0)
	flags |= 0x10
	if dword_5d4594_805860 {
		flags |= 0x18
	}
	if !nox_enable_threads {
		flags |= 0x100
	}
	if memmap.Uint32(0x5D4594, 805864) != 0 {
		flags |= 0x200
	}
	C.nox_xxx_setSomeFunc_48A210(C.int(uintptr(C.sub_47FCE0))) // TODO: another callback
	v1 := nox_client_getCursorType()
	nox_client_setCursorType(gui.CursorSelect)
	v2 := sub_48B3E0(false)
	if err := videoInit(videoGetWindowSize(), int(flags)); err != nil {
		v9 := strMan.GetStringInFile("result:ERROR", "C:\\NoxPost\\src\\Client\\Io\\Win95\\dxvideo.c")
		v4 := strMan.GetStringInFile("gfxDdraw.c:DXWarning", "C:\\NoxPost\\src\\Client\\Io\\Win95\\dxvideo.c")
		// TODO: show OS modal message
		_, _ = v4, v9
		return err
	}
	nox_xxx_cursorLoadAll_477710()
	nox_client_setCursorType(v1)
	sub_48B3E0(v2)
	noxClient.r.ClearScreen(noxClient.r.Data().BgColor())
	nox_xxx_setupSomeVideo_47FEF0()
	C.sub_49F6D0(1)
	noxClient.r.setRectFullScreen()
	*memmap.PtrUint32(0x973F18, 6060) = uint32(2 * sz.X * sz.Y)
	*memmap.PtrUint32(0x973F18, 7696) = 1
	C.sub_430B50(0, 0, noxDefaultWidth-1, noxDefaultHeight-1)
	return nil
}

//export nox_getBackbufferPitch
func nox_getBackbufferPitch() C.int {
	return C.int(2 * noxPixBuffer.img.Stride)
}

func nox_xxx_makeFillerColor_48BDE0() {
	*memmap.PtrUint32(0x5D4594, 1193592) = noxcolor.RGB5551Color(255, 0, 255).Color32()
}

func nox_client_drawGeneral_4B0340(a1 int) int {
	if err := drawGeneral_4B0340(a1); err != nil {
		videoLog.Println(err)
		return 0
	}
	return 1
}

func drawGeneral_4B0340(a1 int) error {
	dword_5d4594_1311936 = true
	*memmap.PtrUint32(0x5D4594, 1311932) = uint32(a1)
	// FIXME
	v1 := false
	videoLog.Println("DrawGeneralStart")
	if /*noxflags.HasEngine(noxflags.EngineWindowed) ||*/ v1 /*|| nox_video_renderTargetFlags&0x10 != 0*/ {
		videoLog.Println("DrawGeneralSkip")
		sub_4B05D0()
		return nil
	}
	C.sub_431290()
	C.sub_43DBD0()
	sub_44D8F0()
	for C.sub_43DC40() != 0 || sub_44D930() {
		sub_4312C0()
	}
	sub_43E8E0(0)
	v12 := sub_48B3E0(false)
	//inpHandler.UnacquireMouse()

	playMovieFile(movieFilesStack[0])

	sub_43E910(0)
	C.sub_43DBE0()
	//inpHandler.AcquireMouse()
	sub_48B3E0(v12)
	sub_4B05D0()
	return nil
}

func nox_xxx_loadDefColor_4A94A0() {
	C.nox_color_black_2650656 = C.uint(nox_color_black_2650656.Color32())
	*memmap.PtrUint32(0x852978, 4) = nox_color_gray1.Color32()
	*memmap.PtrUint32(0x85B3FC, 956) = nox_color_gray2.Color32()
	*memmap.PtrUint32(0x5D4594, 2597996) = nox_color_gray3.Color32()
	C.nox_color_white_2523948 = C.uint(nox_color_white_2523948.Color32())
	C.nox_color_violet_2598268 = C.uint(nox_color_violet_2598268.Color32())
	*memmap.PtrUint32(0x85B3FC, 940) = nox_color_red.Color32()
	C.nox_color_red_2589776 = C.uint(nox_color_red_2589776.Color32())
	*memmap.PtrUint32(0x85B3FC, 984) = nox_color_darkGreen.Color32()
	*memmap.PtrUint32(0x8531A0, 2572) = nox_color_green.Color32()
	C.nox_color_green_2614268 = C.uint(nox_color_green_2614268.Color32())
	*memmap.PtrUint32(0x85B3FC, 944) = nox_color_darkBlue.Color32()
	C.nox_color_cyan_2649820 = C.uint(nox_color_cyan_2649820.Color32())
	C.nox_color_blue_2650684 = C.uint(nox_color_blue_2650684.Color32())
	C.nox_color_orange_2614256 = C.uint(nox_color_orange_2614256.Color32())
	C.nox_color_yellow_2589772 = C.uint(nox_color_yellow_2589772.Color32())
	*memmap.PtrUint32(0x852978, 0) = nox_color_darkYellow.Color32()

	*memmap.PtrPtr(0x85B3FC, 132) = unsafe.Pointer(&C.nox_color_black_2650656)
	*memmap.PtrPtr(0x85B3FC, 136) = memmap.PtrOff(0x852978, 4)
	*memmap.PtrPtr(0x85B3FC, 140) = memmap.PtrOff(0x85B3FC, 956)
	*memmap.PtrPtr(0x85B3FC, 144) = memmap.PtrOff(0x5D4594, 2597996)
	*memmap.PtrPtr(0x85B3FC, 148) = unsafe.Pointer(&C.nox_color_white_2523948)
	*memmap.PtrPtr(0x85B3FC, 152) = unsafe.Pointer(&C.nox_color_violet_2598268)
	*memmap.PtrPtr(0x85B3FC, 156) = memmap.PtrOff(0x85B3FC, 940)
	*memmap.PtrPtr(0x85B3FC, 160) = unsafe.Pointer(&C.nox_color_red_2589776)
	*memmap.PtrPtr(0x85B3FC, 164) = memmap.PtrOff(0x85B3FC, 984)
	*memmap.PtrPtr(0x85B3FC, 168) = memmap.PtrOff(0x8531A0, 2572)
	*memmap.PtrPtr(0x85B3FC, 172) = unsafe.Pointer(&C.nox_color_green_2614268)
	*memmap.PtrPtr(0x85B3FC, 176) = memmap.PtrOff(0x85B3FC, 944)
	*memmap.PtrPtr(0x85B3FC, 180) = unsafe.Pointer(&C.nox_color_cyan_2649820)
	*memmap.PtrPtr(0x85B3FC, 184) = unsafe.Pointer(&C.nox_color_blue_2650684)
	*memmap.PtrPtr(0x85B3FC, 188) = unsafe.Pointer(&C.nox_color_orange_2614256)
	*memmap.PtrPtr(0x85B3FC, 192) = unsafe.Pointer(&C.nox_color_yellow_2589772)
	*memmap.PtrPtr(0x85B3FC, 196) = memmap.PtrOff(0x852978, 0)
}

func nox_video_initFloorBuffer_430BA0(sz image.Point) error {
	nox_xxx___cfltcvt_init_430CC0()
	nox_xxx_tileInitBuf_430DB0(sz.X, sz.Y)
	if lightsOutBuf == nil {
		lightsOutBuf, _ = alloc.Make([]uint32{}, 6)
		lightsOutBuf[0] = 255
		lightsOutBuf[1] = 255
		lightsOutBuf[2] = 255
	}
	return nil
}

func nox_xxx___cfltcvt_init_430CC0() {
	*memmap.PtrUint32(0x973F18, 7696) = 1
}

func sub_4AE520() {
	noxClient.r.circleSeg.Init(noxClient.r)
	C.sub_4AEE30()
}

func sub_49F610(sz image.Point) {
	p := noxClient.r.Data()
	p.useClip = 0
	p.SetClipRect(image.Rectangle{Max: sz})
	p.SetClipRect2(image.Rectangle{Max: image.Pt(sz.X-1, sz.Y-1)})
	p.SetRect3(image.Rectangle{Max: sz})
	C.dword_5d4594_1305748 = 0
}

//export nox_client_clearScreen_440900
func nox_client_clearScreen_440900() {
	r := noxClient.r
	r.ClearScreen(r.Data().BgColor())
}

func nox_free_pixbuffers_486110() {
	if memmap.Uint32(0x5D4594, 1193200) == 0 {
		noxPixBuffer.img = nil
		if noxPixBuffer.free != nil {
			noxPixBuffer.free()
			noxPixBuffer.free = nil
		}

		if nox_pixbuffer_3798788_arr != nil {
			alloc.FreeSlice(nox_pixbuffer_3798788_arr)
			nox_pixbuffer_3798788_arr = nil
		}
	}
	noxPixBuffer.rows = nil
	if noxPixBuffer.freeRows != nil {
		noxPixBuffer.freeRows()
		noxPixBuffer.freeRows = nil
		C.nox_pixbuffer_rows_3798784 = nil
	}
}

func nox_video_initPixbuffer_486090(sz image.Point) {
	videoLog.Printf("initializing pixbuffer: %v", sz)
	nox_video_initPixbufferData_4861D0(sz)
	nox_video_initPixbufferRows_486230()
	for _, fnc := range noxPixBuffer.onResize {
		fnc(sz)
	}
}

func nox_video_initPixbufferData_4861D0(sz image.Point) {
	if memmap.Uint32(0x5D4594, 1193200) != 0 {
		return
	}
	data, free := alloc.Make([]uint16{}, sz.X*sz.Y)
	noxPixBuffer.free = free
	noxPixBuffer.img = noximage.NewImage16WithData(data, sz)
	noxClient.r.SetPixBuffer(noxPixBuffer.img)
	if nox_video_renderTargetFlags&0x40 == 0 {
		return
	}

	nox_pixbuffer_3798788_arr, _ = alloc.Make([]byte{}, len(data))
}

func nox_video_initPixbufferRows_486230() {
	sz := noxPixBuffer.img.Size()
	ptrs, freeRows := alloc.Make([]*uint16{}, sz.Y)
	noxPixBuffer.rows = ptrs
	noxPixBuffer.freeRows = freeRows
	C.nox_pixbuffer_rows_3798784 = (**C.uchar)(unsafe.Pointer(&noxPixBuffer.rows[0]))
	for y := 0; y < sz.Y; y++ {
		noxPixBuffer.rows[y] = &noxPixBuffer.img.Row(y)[0]
	}
}

func (r *NoxRender) noxDrawCursor(a1 *Image, pos image.Point) int {
	if dword_5d4594_1193672 && a1 != nil {
		r.DrawImageAt(a1, pos)
	}
	return 1
}

//export nox_draw_setCutSize_476700
func nox_draw_setCutSize_476700(cutPerc C.int, a2 C.int) {
	noxClient.nox_draw_setCutSize(int(cutPerc), int(a2))
}
func (c *Client) nox_draw_setCutSize(perc int, a2 int) {
	vp := c.Viewport()
	bsz := noxPixBuffer.img.Size()
	v2 := a2
	v4 := vp.Size.X
	if a2 != 0 {
		v7 := 0
		for v7 < 4 {
			perc = v2 + 100*(bsz.X-2*vp.Screen.Min.X)/bsz.X
			v6 := perc * bsz.X / 100
			if v2 >= 0 {
				v2++
			} else {
				v2--
			}
			if v6-v4 < 0 {
				v7 = v4 - v6
			} else {
				v7 = v6 - v4
			}
		}
	}
	if perc < 40 {
		perc = 40
	}
	if perc > 100 {
		perc = 100
	}
	nox_video_cutSize = perc

	vp.Screen.Min.X = int(uint32((bsz.X-perc*bsz.X/100)/2) & 0xFFFFFFFC)
	if vp.Screen.Min.X < 0 {
		vp.Screen.Min.X = 0
	}

	vp.Screen.Min.Y = (bsz.Y - perc*bsz.Y/100) / 2
	if vp.Screen.Min.Y < 0 {
		vp.Screen.Min.Y = 0
	}

	vp.Screen.Max.X = int(uint32(bsz.X-vp.Screen.Min.X+2) & 0xFFFFFFFC)
	if vp.Screen.Max.X >= bsz.X {
		vp.Screen.Max.X = bsz.X - 1
	}

	vp.Screen.Max.Y = bsz.Y - vp.Screen.Min.Y - 1
	if vp.Screen.Max.Y >= bsz.Y {
		vp.Screen.Max.Y = bsz.Y - 1
	}

	vp.Size.X = vp.Screen.Dx() + 1
	vp.Size.Y = vp.Screen.Dy() + 1
	C.dword_5d4594_1193188 = 1
	C.dword_5d4594_3799524 = 1
}

func nox_client_drawXxx_444AC0(w, h int, flags int) error {
	//int v5;             // eax
	//bool v6;            // zf
	//unsigned char v7 = 0; // al
	//int v8;             // esi
	//int v9;             // eax
	//int v10;            // eax

	*memmap.PtrUint32(0x5D4594, 823780) = 1

	nox_video_renderTargetFlags = flags

	v7 := byte(nox_video_renderTargetFlags | 0x20)
	nox_video_renderTargetFlags |= 0x120
	v8 := int(uint32(w) & 0xFFFFFFE0)
	if v7&4 != 0 {
		panic("unreachable")
	}
	if err := noxClient.resetRenderer(image.Point{X: v8, Y: h}, true); err != nil {
		return err
	}
	return nil
}

func sub_48B680(a1 int) {
	p := noxClient.r.Data()
	if a1 != int(p.field_15) {
		p.multiply14 = uint32(a1)
	}
}

func (c *Client) nox_video_cursorDrawImpl_477A30(pos image.Point) {
	v18 := memmap.Uint32(0x973F18, 68)
	pos = pos.Sub(image.Point{X: 64, Y: 64})
	*memmap.PtrUint32(0x973F18, 68) = 0
	dword_5d4594_3798728 = true
	defer func() {
		dword_5d4594_3798728 = false
	}()
	dword_5d4594_1097212 = pos
	if gameFrame()&1 != 0 {
		*memmap.PtrUint32(0x5D4594, 1097288)++
	}
	c.r.Data().SetTextColor(nox_color_yellow_2589772)
	fh := c.r.FontHeight(nil)
	if C.nox_xxx_guiSpell_460650() != 0 || C.sub_4611A0() != 0 {
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Target, pos)
		nox_xxx_cursorTypePrev_587000_151528 = gui.CursorTarget
		*memmap.PtrUint32(0x973F18, 68) = v18
		return
	}

	if nox_client_mouseCursorType != nox_xxx_cursorTypePrev_587000_151528 && nox_client_mouseCursorType != 14 {
		sub_48B680(0)
	}
	switch typ := nox_client_getCursorType(); typ {
	case gui.CursorGrab:
		str := strMan.GetStringInFile("GRAB", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 54, Y: 64 - fh}))
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Grab, pos)
	case gui.CursorPickup:
		str := strMan.GetStringInFile("PICKUP", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 49, Y: 64 + fh}))
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Pickup, pos)
		dword_5d4594_1097208 = -2 * fh
	case gui.CursorShop:
		str := strMan.GetStringInFile("SHOPKEEPER", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 39, Y: 64 - fh}))
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Trade, pos)
	case gui.CursorTalk:
		str := strMan.GetStringInFile("TALK", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 49, Y: 64 - fh}))
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Talk, pos)
	case gui.CursorIdentify, gui.CursorCantIdentify:
		str := strMan.GetStringInFile("IDENTIFY", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 49, Y: +88}))
		if typ == gui.CursorIdentify {
			c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Identify, pos)
		} else {
			c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.IdentifyNo, pos)
		}
	case gui.CursorRepair:
		str := strMan.GetStringInFile("REPAIR", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 49, Y: 64 - fh}))
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Repair, pos)
		dword_5d4594_1097208 = 2*fh + 4
	case gui.CursorCreateGame:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.CreateGame, pos)
	case gui.CursorBusy:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Busy, pos)
	case gui.CursorBuy:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Buy, pos)
	case gui.CursorSell:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Sell, pos)
	case gui.CursorUse:
		str := strMan.GetStringInFile("USE", "C:\\NoxPost\\src\\Client\\Gui\\guicurs.c")
		c.r.DrawString(nil, str, pos.Add(image.Point{X: 54, Y: 64 + fh}))
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Use, pos)
		dword_5d4594_1097208 = -2 * fh
	case gui.CursorMoveArrow:
		mpos := c.inp.GetMousePos()
		v19 := types.Pointf{
			X: float32(mpos.X - nox_win_width/2),
			Y: float32(mpos.Y - nox_win_height/2),
		}
		v15 := nox_xxx_math_509ED0(v19) / 8
		if v19.X*v19.X+v19.Y*v19.Y > 100*100 || memmap.Uint32(0x852978, 8) != 0 && *(*uint32)(unsafe.Add(*memmap.PtrPtr(0x852978, 8), 276)) == 6 {
			v15 += 32
		}
		if v16 := nox_xxx_spriteGetMB_476F80(); v16 != nil {
			sub_48B680(1)
			if v16.Flags28()&6 == 0 || C.sub_495A80(C.int(v16.Field32())) != 0 {
				c.r.setColorMultAndIntensity(nox_color_blue_2650684)
			} else {
				c.r.setColorMultAndIntensity(nox_color_red)
			}
		} else {
			sub_48B680(0)
		}
		c.r.sub_4BE710(noxCursors.Move, pos, int(v15))
		c.r.Data().setMultiply14(0)
	case gui.CursorPickupFar:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.PickupFar, pos)
		dword_5d4594_1097208 = -2 * fh
	case gui.CursorCaution:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Caution, pos)
		dword_5d4594_1097208 = -fh
	default:
		c.nox_video_drawAnimatedImageOrCursorAt(noxCursors.Select, pos)
	}
	nox_xxx_cursorTypePrev_587000_151528 = nox_client_mouseCursorType
	*memmap.PtrUint32(0x973F18, 68) = v18
}

//export nox_xxx_bookSaveSpellForDragDrop_477640
func nox_xxx_bookSaveSpellForDragDrop_477640(a1, a2 C.int) {
	nox_client_spellDragnDrop_1097192 = uint32(a1)
	nox_client_spellDragnDrop_type_1097196 = int(a2)
}

//export nox_xxx_bookSpellDnDclear_477660
func nox_xxx_bookSpellDnDclear_477660() {
	nox_client_spellDragnDrop_1097192 = 0
	nox_client_spellDragnDrop_type_1097196 = 0
}

//export nox_xxx_bookGetSpellDnDType_477670
func nox_xxx_bookGetSpellDnDType_477670() C.int {
	return C.int(nox_client_spellDragnDrop_type_1097196)
}

func nox_xxx_cursorGetDraggedItem_477680() *Drawable {
	return nox_client_itemDragnDrop_1097188
}

//export nox_xxx_cursorSetDraggedItem_477690
func nox_xxx_cursorSetDraggedItem_477690(a1 *nox_drawable) {
	nox_client_itemDragnDrop_1097188 = asDrawable(a1)
}

//export nox_xxx_cursorResetDraggedItem_4776A0
func nox_xxx_cursorResetDraggedItem_4776A0() {
	nox_client_itemDragnDrop_1097188 = nil
}

func (c *Client) nox_client_drawCursorAndTooltips_477830() {
	if noxCursors.Select == nil {
		nox_xxx_cursorLoadAll_477710()
	}
	mpos := c.inp.GetMousePos()
	vp, freeVp := alloc.New(Viewport{})
	defer freeVp()
	vp.Screen = image.Rect(0, 0, nox_win_width, nox_win_height)
	vp.World.Min = image.Pt(0, 0)
	vp.Size = image.Pt(nox_win_width, nox_win_height)
	dword_5d4594_1097204 = 0
	dword_5d4594_1097208 = noxClient.r.FontHeight(nil) + 4
	if nox_client_itemDragnDrop_1097188 != nil { // Dragging item
		nox_client_itemDragnDrop_1097188.SetPos(mpos)
		nox_client_itemDragnDrop_1097188.DrawFunc(vp)
	}
	if nox_client_spellDragnDrop_1097192 != 0 { // Player is dragging spell or ability
		pl := noxServer.getPlayerByID(clientPlayerNetCode())
		if pl == nil || pl.PlayerClass() != player.Warrior {
			v2 := nox_xxx_spellIcon_424A90(C.int(nox_client_spellDragnDrop_1097192)) // Spell icon
			if v2 != nil {
				c.r.DrawImageAt(asImageP(unsafe.Pointer(v2)), mpos.Sub(image.Point{X: 15, Y: 15}))
			}
		} else {
			v2 := nox_xxx_spellGetAbilityIcon_425310(C.int(nox_client_spellDragnDrop_1097192), 0) // Ability icon
			if v2 != nil {
				c.r.DrawImageAt(asImageP(unsafe.Pointer(v2)), mpos.Sub(image.Point{X: 15, Y: 15}))
			}
		}
	}
	c.nox_video_cursorDrawImpl_477A30(mpos)
	if str := GoWStringP(memmap.PtrOff(0x5D4594, 1096676)); str != "" && nox_client_showTooltips_80840 {
		sz := c.r.GetStringSizeWrapped(nil, str, 0)
		px := mpos.X - dword_5d4594_1097204
		py := mpos.Y - dword_5d4594_1097208
		sz.X += 4
		sz.Y += 4
		if px+sz.X >= nox_win_width {
			px -= sz.X
		}
		if px < 0 {
			px = 0
		}
		if py+sz.Y >= nox_win_height {
			py = nox_win_height - sz.Y
		}
		if py < 0 {
			py = 0
		}
		c.r.DrawRectFilledAlpha(px, py, sz.X, sz.Y)
		c.r.Data().SetTextColor(nox_color_yellow_2589772)
		c.r.DrawStringWrapped(nil, str, image.Rect(px+2, py+2, px+2, py+2))
		if C.dword_5d4594_3799468 != 0 {
			vp := noxClient.Viewport()
			if px < vp.Screen.Min.X || px+sz.X > vp.Screen.Max.X || py < vp.Screen.Min.Y || py+sz.Y > vp.Screen.Max.Y {
				C.dword_5d4594_3799524 = 1
			}
		}
	}
}

func (c *Client) sub_477F80() {
	if C.dword_5d4594_3799468 != 0 {
		vp := c.Viewport()
		if dword_5d4594_1097212.X < vp.Screen.Min.X || dword_5d4594_1097212.X+cursorSize >= vp.Screen.Max.X ||
			dword_5d4594_1097212.Y < vp.Screen.Min.Y || dword_5d4594_1097212.Y+cursorSize >= vp.Screen.Max.Y {
			c.r.DrawRectFilledOpaque(dword_5d4594_1097212.X+cursorSize/2, dword_5d4594_1097212.Y+cursorSize/2, cursorSize, cursorSize, color.Black)
		}
	}
}

func (c *Client) sub_444C50() {
	if C.dword_5d4594_823776 != 0 {
		nox_free_pixbuffers_486110()
		nox_draw_freeColorTables_433C20()
		c.r.FadeReset()
		c.r.freeParticles()
		c.r.partfx.Free()
		c.r.circleSeg.Free()
		nox_xxx_FontDestroy_43F2E0()
		C.dword_5d4594_823776 = 0
		if memmap.Uint32(0x5D4594, 823780) != 0 {
			*memmap.PtrUint32(0x5D4594, 823780) = 0
		}
	}
}

//export sub_478000
func sub_478000() C.int {
	C.sub_467CD0()
	if nox_client_spellDragnDrop_type_1097196 != 0 {
		v1 := nox_xxx_wndGetCaptureMain()
		nox_xxx_wndClearCaptureMain(v1)
		nox_xxx_bookSpellDnDclear_477660()
	}
	return 0
}
