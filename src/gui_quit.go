package opennox

/*
#include "client__gui__guiquit.h"
#include "GAME1_1.h"
#include "GAME1_3.h"
#include "GAME3_2.h"
#include "GAME4_1.h"
extern nox_window* nox_wnd_quitMenu_825760;
extern unsigned int dword_5d4594_2650652;
extern uint32_t dword_5d4594_830272;
int* nox_xxx_guiServerOptionsHide_4597E0(int a1);
*/
import "C"
import (
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
)

func nox_xxx_wndLoadQuitMenu_445790() int {
	win := newWindowFromFile("QuitMenu.wnd", wrapWindowFuncC(C.nox_xxx_menuGameOnButton_445840))
	if win == nil {
		return 0
	}
	C.nox_wnd_quitMenu_825760 = win.C()
	win.SetAllFuncs(func(win *Window, ev WindowEvent) WindowEventResp {
		return RawEventResp(1)
	}, func(win *Window, draw *WindowData) int {
		pos := win.GlobalPos()
		sz := win.Size()
		noxClient.r.DrawRectFilledAlpha(pos.X, pos.Y, sz.X, sz.Y)
		return 1
	}, nil)
	sz := win.Size()
	win.DrawData().SetBackgroundColor(nox_color_black_2650656)
	win.off_x = int32(nox_win_width-sz.X) / 2
	win.end_x = win.off_x + int32(sz.X)
	if nox_win_height > 768 {
		win.off_y = int32(nox_win_height-sz.Y-nox_win_height/3) / 2
		win.end_y = win.off_y + int32(sz.Y)
	}
	C.sub_445C40()
	dword_5d4594_825752 = 0
	dword_5d4594_825768 = 0
	str := strMan.GetStringInFile("Vote", "guiquit.c")
	StrCopy((*C.char)(memmap.PtrOff(0x5D4594, 825772)), 256, str)
	return 1
}

func sub_467440(a1 int) {
	*memmap.PtrUint32(0x5D4594, 1062544) = uint32(a1)
}

//export nox_xxx____setargv_4_44B000
func nox_xxx____setargv_4_44B000() {
	C.dword_5d4594_830272 = 1
}

func sub_446190() {
	nox_xxx_serverIsClosing_825764 = false
	if noxflags.HasGame(noxflags.GameHost) {
		str := strMan.GetStringInFile("ServerManualShutdown", "guiquit.c")
		netLog.Println(str)
	}
	if noxflags.HasGame(noxflags.GameFlag26) {
		nox_game_checkStateWol_43C260()
		sub_41E300(9)
		nox_xxx____setargv_4_44B000()
	} else if noxServer.nox_xxx_isQuest_4D6F50() {
		if sub4D6F30() {
			nox_game_checkStateMenu_43C2F0()
		}
		C.sub_4D70B0()
		sub_4D6F40(0)
		sub_4D6F90(0)
	}
	if dword_5d4594_825768 == 0 {
		dword_5d4594_825752 = 0
		dword_5d4594_825768 = gameFrame()
		if noxflags.HasGame(noxflags.GameHost) {
			if nox_xxx_check_flag_aaa_43AF70() == 1 && noxflags.HasGame(noxflags.GameFlag15|noxflags.GameFlag16) {
				sub_416150(15, 0)
			}
			C.sub_509CB0()
		}
		if noxflags.HasGame(noxflags.GameHost) {
			dword_5d4594_825752 = noxServer.cntPlayers() - 1
			sub_467440(0)
			if dword_5d4594_825752 > 0 {
				netSendServerQuit()
				v3 := strMan.GetStringInFile("ShuttingDown", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
				v2 := strMan.GetStringInFile("Wolchat.c:PleaseWait", "C:\\NoxPost\\src\\client\\Gui\\guiquit.c")
				NewDialogWindow(nil, v2, v3, 0, nil, nil)
				sub_44A360(1)
				goto LABEL_20
			}
		} else if !noxflags.HasGame(noxflags.GameOnline) {
			goto LABEL_20
		}
		sub_446380()
	}
LABEL_20:
	if C.dword_5d4594_2650652 != 0 {
		if sub_41E2F0() == 9 {
			C.sub_41F4B0()
			C.sub_41EC30()
			C.sub_446490(0)
			nox_xxx____setargv_4_44B000()
		}
	}
}

//export sub_446380
func sub_446380() {
	sub_44A400()
	if noxflags.HasGame(noxflags.GameOnline) && !noxflags.HasGame(noxflags.GameModeQuest) && !sub4D6F30() {
		C.sub_41CEE0(memmap.PtrOff(0x85B3FC, 10980), 1)
	}
	C.nox_xxx_guiServerOptionsHide_4597E0(1)
	nox_game_exit_xxx2()
	sub_446060()
}

//export sub_445B40
func sub_445B40() C.int {
	sub_413A00(0)
	if sub4DB790("AUTOSAVE") {
		return 1
	}
	v2 := strMan.GetStringInFile("GUISave.c:SaveErrorTitle", "guiquit.c")
	v1 := strMan.GetStringInFile("GUISave.c:SaveErrorTitle", "guiquit.c")
	NewDialogWindow(nil, v1, v2, 33, nil, nil)
	return 1
}
