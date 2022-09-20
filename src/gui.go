package opennox

/*
#include "GAME1_3.h"
#include "GAME2_1.h"
#include "GAME2_2.h"
#include "GAME2_3.h"
#include "client__gui__guicon.h"
#include "win.h"

extern nox_window* dword_5d4594_1064896;
extern nox_window_ref* nox_win_1064912;

extern unsigned int nox_client_renderGUI_80828;
extern unsigned int nox_xxx_xxxRenderGUI_587000_80832;
*/
import "C"
import (
	"context"
	"image"
	"image/color"
	"os"
	"unsafe"

	"golang.org/x/image/font"

	"github.com/noxworld-dev/opennox-lib/client/keybind"
	noxcolor "github.com/noxworld-dev/opennox-lib/color"
	"github.com/noxworld-dev/opennox-lib/log"
	"github.com/noxworld-dev/opennox-lib/strman"

	"github.com/noxworld-dev/opennox/v1/client/gui"
	"github.com/noxworld-dev/opennox/v1/client/input"
	"github.com/noxworld-dev/opennox/v1/common/alloc"
)

var (
	guiLog                       = log.New("gui")
	guiDebug                     = os.Getenv("NOX_DEBUG_GUI") == "true"
	nox_win_activeWindow_1064900 *Window
	nox_win_1064916              *Window
	nox_win_freeList             *Window // dword_5d4594_1064896
)

func enableGUIDrawing(enable bool) {
	if enable {
		// TODO: might be a bitfield
		C.nox_client_renderGUI_80828 = 1
		C.nox_xxx_xxxRenderGUI_587000_80832 = 1
	} else {
		C.nox_client_renderGUI_80828 = 0
		C.nox_xxx_xxxRenderGUI_587000_80832 = 0
	}
}

//export nox_xxx_wndGetFocus_46B4F0
func nox_xxx_wndGetFocus_46B4F0() *C.nox_window {
	return nox_win_cur_focused.C()
}

//export nox_xxx_windowFocus_46B500
func nox_xxx_windowFocus_46B500(win *C.nox_window) C.int {
	guiFocus(asWindow(win))
	return 0
}

//export nox_client_getWin1064916_46C720
func nox_client_getWin1064916_46C720() *C.nox_window {
	return nox_win_1064916.C()
}

//export nox_xxx_wndSetCaptureMain_46ADC0
func nox_xxx_wndSetCaptureMain_46ADC0(win *C.nox_window) C.int {
	return C.int(nox_xxx_wndSetCaptureMain(asWindow(win)))
}

//export nox_xxx_wndClearCaptureMain_46ADE0
func nox_xxx_wndClearCaptureMain_46ADE0(win *C.nox_window) C.int {
	nox_xxx_wndClearCaptureMain(asWindow(win))
	return 0
}

//export nox_xxx_wndGetCaptureMain_46AE00
func nox_xxx_wndGetCaptureMain_46AE00() *C.nox_window {
	return nox_xxx_wndGetCaptureMain().C()
}

func nox_xxx_wndSetCaptureMain(win *Window) int {
	if nox_win_cur_input != nil {
		return -4
	}
	nox_win_cur_input = win
	return 0
}

func nox_xxx_wndClearCaptureMain(win *Window) int {
	if win == nox_win_cur_input {
		nox_win_cur_input = nil
	}
	return 0
}

func nox_xxx_wndGetCaptureMain() *Window {
	return nox_win_cur_input
}

func asWindowData(data *C.nox_window_data) *WindowData {
	return asWindowDataP(unsafe.Pointer(data))
}

func asWindowDataP(data unsafe.Pointer) *WindowData {
	return (*WindowData)(data)
}

func newWindowData() (*WindowData, func()) {
	return alloc.New(WindowData{})
}

var _ = [1]struct{}{}[332-unsafe.Sizeof(WindowData{})]

type WindowData struct {
	field0    uint32                 // 0, 0 (36)
	group     int32                  // 1, 4 (40)
	style     int32                  // 2, 8 (44)
	status    int32                  // 3, 12 (48)
	win       *C.nox_window          // 4, 16 (52)
	bgColor   uint32                 // 5, 20 (56)
	bgImage   *nox_video_bag_image_t // 6, 24 (60)
	enColor   uint32                 // 7, 28 (64)
	enImage   *nox_video_bag_image_t // 8, 32 (68)
	hlColor   uint32                 // 9, 36 (72)
	hlImage   *nox_video_bag_image_t // 10, 40 (76)
	disColor  uint32                 // 11, 44 (80)
	disImage  *nox_video_bag_image_t // 12, 48 (84)
	selColor  uint32                 // 13, 52 (88)
	selImage  *nox_video_bag_image_t // 14, 56 (92)
	imgPx     int32                  // 15, 60 (96)
	imgPy     int32                  // 16, 64 (100)
	textColor uint32                 // 17, 68 (104)
	text      [64]wchar_t            // 18, 72 (108)
	font      unsafe.Pointer         // 50, 200 (236)
	tooltip   [64]wchar_t            // 51, 204 (240)
}

func (d *WindowData) C() *C.nox_window_data {
	return (*C.nox_window_data)(unsafe.Pointer(d))
}

func (d *WindowData) Field0() uint32 {
	return d.field0
}

func (d *WindowData) Field0Set(flag uint32, v bool) {
	if v {
		d.field0 |= flag
	} else {
		d.field0 &^= flag
	}
}

func (d *WindowData) Group() int {
	return int(d.group)
}

func (d *WindowData) SetGroup(v int) {
	d.group = int32(v)
}

func (d *WindowData) StyleFlags() gui.StyleFlags {
	return gui.StyleFlags(d.style)
}

func (d *WindowData) SetStyleFlags(v gui.StyleFlags) {
	d.style = int32(v)
}

func (d *WindowData) Status() gui.StatusFlags {
	return gui.StatusFlags(d.status)
}

func (d *WindowData) SetStatus(v gui.StatusFlags) {
	d.status = int32(v)
}

func (d *WindowData) Window() *Window {
	return asWindow(d.win)
}

func (d *WindowData) Text() string {
	return GoWString(&d.text[0])
}

func (d *WindowData) SetText(s string) {
	n := len(d.text)
	WStrCopy(&d.text[0], n, s)
	d.text[n-1] = 0
}

func (d *WindowData) BackgroundImage() *Image {
	return asImage(d.bgImage)
}

func (d *WindowData) SetBackgroundImage(p *Image) {
	d.bgImage = p.C()
}

func (d *WindowData) BackgroundColor() noxcolor.Color16 {
	return noxcolor.RGBA5551(d.bgColor >> 16)
}

func (d *WindowData) SetBackgroundColor(cl color.Color) {
	d.bgColor = noxcolor.ToRGBA5551Color(cl).Color32()
}

func (d *WindowData) EnabledImage() *Image {
	return asImage(d.enImage)
}

func (d *WindowData) SetEnabledImage(p *Image) {
	d.enImage = p.C()
}

func (d *WindowData) EnabledColor() noxcolor.Color16 {
	return noxcolor.RGBA5551(d.enColor >> 16)
}

func (d *WindowData) SetEnabledColor(cl color.Color) {
	d.enColor = noxcolor.ToRGBA5551Color(cl).Color32()
}

func (d *WindowData) DisabledImage() *Image {
	return asImage(d.disImage)
}

func (d *WindowData) SetDisabledImage(p *Image) {
	d.disImage = p.C()
}

func (d *WindowData) DisabledColor() noxcolor.Color16 {
	return noxcolor.RGBA5551(d.disColor >> 16)
}

func (d *WindowData) SetDisabledColor(cl color.Color) {
	d.disColor = noxcolor.ToRGBA5551Color(cl).Color32()
}

func (d *WindowData) HighlightImage() *Image {
	return asImage(d.hlImage)
}

func (d *WindowData) SetHighlightImage(p *Image) {
	d.hlImage = p.C()
}

func (d *WindowData) HighlightColor() noxcolor.Color16 {
	return noxcolor.RGBA5551(d.hlColor >> 16)
}

func (d *WindowData) SetHighlightColor(cl color.Color) {
	d.hlColor = noxcolor.ToRGBA5551Color(cl).Color32()
}

func (d *WindowData) SelectedImage() *Image {
	return asImage(d.selImage)
}

func (d *WindowData) SetSelectedImage(p *Image) {
	d.selImage = p.C()
}

func (d *WindowData) SelectedColor() noxcolor.Color16 {
	return noxcolor.RGBA5551(d.selColor >> 16)
}

func (d *WindowData) SetSelectedColor(cl color.Color) {
	d.selColor = noxcolor.ToRGBA5551Color(cl).Color32()
}

func (d *WindowData) TextColor() noxcolor.Color16 {
	return noxcolor.RGBA5551(d.textColor >> 16)
}

func (d *WindowData) SetTextColor(cl color.Color) {
	d.textColor = noxcolor.ToRGBA5551Color(cl).Color32()
}

func (d *WindowData) Font() font.Face {
	return asFont(d.font)
}

func (d *WindowData) FontC() unsafe.Pointer {
	return d.font
}

func (d *WindowData) SetFont(font unsafe.Pointer) {
	d.font = font
}

func (d *WindowData) ImagePoint() image.Point {
	return image.Point{
		X: int(d.imgPx),
		Y: int(d.imgPy),
	}
}

func (d *WindowData) SetImagePoint(p image.Point) {
	d.imgPx = int32(p.X)
	d.imgPy = int32(p.Y)
}

func (d *WindowData) Tooltip() string {
	return GoWString(&d.tooltip[0])
}

func (d *WindowData) SetTooltip(sm *strman.StringManager, s string) {
	n := len(d.tooltip)
	if CWLen(s) > n && sm != nil {
		s = sm.GetStringInFile("TooltipTooLong", "C:\\NoxPost\\src\\Client\\Gui\\GameWin\\gamewin.c")
	}
	WStrCopy(&d.tooltip[0], n, s)
	d.tooltip[n-1] = 0
}

func (d *WindowData) SetDefaults(def gui.StyleDefaults) {
	d.SetEnabledColor(def.EnabledColor)
	d.SetHighlightColor(def.HighlightColor)
	d.SetDisabledColor(def.DisabledColor)
	d.SetBackgroundColor(def.BackgroundColor)
	d.SetSelectedColor(def.SelectedColor)
	d.SetTextColor(def.TextColor)
}

func DrawGUI() {
	// back layer (background and some UI parts)
	for win := nox_win_xxx1_first; win != nil; win = win.Next() {
		if win.Flags().Has(gui.StatusBelow) {
			win.drawRecursive()
		}
	}
	// middle layer
	for win := nox_win_xxx1_first; win != nil; win = win.Next() {
		if win.Flags().HasNone(gui.StatusBelow | gui.StatusAbove) {
			win.drawRecursive()
		}
	}
	// front layer
	for win := nox_win_xxx1_first; win != nil; win = win.Next() {
		if win.Flags().Has(gui.StatusAbove) {
			win.drawRecursive()
		}
	}
}

//export nox_gui_draw
func nox_gui_draw() {
	DrawGUI()
}

//export nox_color_rgb_4344A0
func nox_color_rgb_4344A0(r, g, b C.int) C.uint32_t {
	return C.uint32_t(noxcolor.RGB5551Color(byte(r), byte(g), byte(b)).Color32())
}

//export nox_set_color_rgb_434430
func nox_set_color_rgb_434430(r, g, b C.int) {
	noxClient.r.Data().SetColor2(noxcolor.RGB5551Color(byte(r), byte(g), byte(b)))
}

func unsafePtrToInt(p unsafe.Pointer) C.int {
	return C.int(uintptr(p))
}

func dataPtrToInt(p *WindowData) C.int {
	return C.int(uintptr(unsafe.Pointer(p)))
}

//export nox_xxx_wndWddSetTooltip_46B000
func nox_xxx_wndWddSetTooltip_46B000(draw *C.nox_window_data, str *C.wchar_t) {
	d := asWindowData(draw)
	if str == nil {
		d.SetTooltip(strMan, "")
		return
	}
	d.SetTooltip(strMan, GoWString(str))
}

func nox_xxx_windowUpdateKeysMB_46B6B0(inp *input.Handler, key keybind.Key) {
	root := nox_win_cur_focused
	if root == nil {
		return
	}
	if key == 0 {
		return
	}
	if inp.GetKeyFlag(key) {
		return
	}
	ok := false
	for win := root; win != nil; win = win.Parent() {
		if eventRespBool(win.Func93(WindowKeyPress{Key: key, Pressed: inp.IsPressed(key)})) {
			ok = true
			break
		}
	}
	inp.SetKeyFlag(key, ok)
}

//var dword_5d4594_2618912 *noxKeyEventInt

func (c *Client) keyBindingsCheckActive() int {
	if C.sub_46A4A0() != 0 {
		return 1
	}
	for _, key := range c.inp.KeyboardKeys() {
		//dword_5d4594_2618912 = p
		if !c.inp.GetKeyFlag(key) && !c.inp.IsPressed(key) && nox_xxx_guiCursor_477600() == 0 {
			if noxConsole.ExecMacros(context.Background(), key) {
				c.inp.SetKeyFlag(key, true)
			}
		}
	}
	return 1
}

func nox_xxx_pointInRect_4281F0(p image.Point, r image.Rectangle) bool {
	return p.In(r)
}

func freeAllWindowsInList() {
	win := nox_win_freeList
	nox_win_freeList = nil
	for win != nil {
		prev := win.Prev()
		win.prev = nil
		if nox_win_cur_input == win {
			nox_win_cur_input = nil
		}
		if nox_win_cur_focused == win {
			guiFocus(nil)
		}
		if nox_win_1064912 != nil && win == nox_win_1064912.Win {
			nox_xxx_wnd46C6E0(nox_win_1064912.Win)
		}
		if nox_win_activeWindow_1064900 == win {
			nox_win_activeWindow_1064900 = nil
		}
		if nox_win_1064916 == win {
			nox_win_1064916 = nil
		}
		win.Func94(WindowDestroy{})
		nox_alloc_window.FreeObjectFirst(win)
		win = prev
	}
}

//export sub_46B120
func sub_46B120(a1, a2 *C.nox_window) C.int {
	return C.int(sub46B120(asWindow(a1), asWindow(a2)))
}

func sub46B120(win, par *Window) int {
	if win == nil {
		return -2
	}
	if win.Parent() != nil {
		sub_46B180(win)
	} else {
		nox_client_wndListXxxRemove_46A960(win)
	}
	if par != nil {
		win.setParent(par)
	} else {
		nox_client_wndListXxxAdd_46A920(win)
		win.parent = nil
	}
	return 0
}

func sub_46AEE0(a1 *Window, a2 string) {
	a1.Func94(asWindowEvent(guiEventStaticTextSetText, uintptr(unsafe.Pointer(internWStr(a2))), 0))
}

func nox_xxx_wndEditProc_487D70(a1 *Window, ev WindowEvent) RawEventResp {
	a2 := ev.EventCode()
	a3, a4 := ev.EventArgsC()
	return RawEventResp(C.nox_xxx_wndEditProc_487D70(a1.C(), C.int(a2), C.int(a3), C.int(a4)))
}
