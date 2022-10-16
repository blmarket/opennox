package opennox

/*
#include "GAME1_1.h"
#include "GAME1_3.h"
#include "GAME2.h"
#include "GAME2_1.h"
#include "GAME3_3.h"
#include "GAME4_2.h"
#include "server__xfer__savegame__savegame.h"
extern unsigned int dword_5d4594_825764;
extern uint32_t dword_5d4594_1049844;
void nox_xxx_unitsNewAddToList_4DAC00();
*/
import "C"
import (
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	crypt "github.com/noxworld-dev/noxcrypt"
	"github.com/noxworld-dev/opennox-lib/common"
	"github.com/noxworld-dev/opennox-lib/console"
	"github.com/noxworld-dev/opennox-lib/datapath"
	"github.com/noxworld-dev/opennox-lib/ifs"
	"github.com/noxworld-dev/opennox-lib/object"
	"github.com/noxworld-dev/opennox-lib/types"

	"github.com/noxworld-dev/opennox/v1/common/alloc"
	noxflags "github.com/noxworld-dev/opennox/v1/common/flags"
	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/server"
)

var (
	dword_5d4594_1559960 string
)

func nox_xxx_playerSaveToFile_41A140(path string, ind int) bool {
	return C.nox_xxx_playerSaveToFile_41A140(internCStr(path), C.int(ind)) != 0
}

//export sub_4DB790
func sub_4DB790(a1 *C.char) C.int {
	return C.int(bool2int(sub4DB790(GoString(a1))))
}

func sub4DB790(a1 string) bool {
	nox_xxx_mapLoadOrSaveMB_4DCC70(1)
	noxflags.SetGame(noxflags.GameFlag28)
	noxAudioServeT(500)
	res := nox_xxx_soloLoadGame_4DB7E0_savegame(a1)
	noxAudioServe()
	noxflags.UnsetGame(noxflags.GameFlag28)
	*memmap.PtrUint32(0x5D4594, 1563068) = gameFrame()
	return res
}

//export nox_xxx_saveMakeFolder_0_4DB1D0
func nox_xxx_saveMakeFolder_0_4DB1D0() C.bool {
	return ifs.Mkdir(datapath.Save()) == nil
}

//export nox_client_makeSaveDir_4DB540
func nox_client_makeSaveDir_4DB540(cname *C.char) C.bool {
	name := GoString(cname)
	_, err := nox_client_makeSaveDir(name)
	return err == nil
}

func nox_client_makeSaveDir(name string) (string, error) {
	dir := datapath.Save(name)
	err := ifs.Mkdir(dir)
	return dir, err
}

//export nox_client_makeSaveMapDir_4DB5A0
func nox_client_makeSaveMapDir_4DB5A0(a1c, a2c *C.char) C.bool {
	a1, a2 := GoString(a1c), GoString(a2c)
	return ifs.Mkdir(datapath.Save(a1, a2)) == nil
}

//export nox_savegame_rm_4DBE10
func nox_savegame_rm_4DBE10(cname *C.char, rmDir C.int) {
	if cname == nil {
		return
	}
	saveName := GoString(cname)
	_ = nox_savegame_rm(saveName, rmDir != 0)
}

func nox_savegame_rm(name string, rmdir bool) error {
	if name == "" {
		return nil
	}
	saveDir := datapath.Save(name)
	if err := ifs.RemoveAll(saveDir); err != nil {
		return err
	}
	// TODO: this should actually only remove the player file and map dirs; we just remove the whole dir instead
	//       implement it properly later, if it makes any difference
	if !rmdir {
		if err := ifs.Mkdir(saveDir); err != nil {
			return err
		}
	}
	return nil
}

//export nox_client_copySave_4DC100
func nox_client_copySave_4DC100(cfrom, cto *C.char) C.int {
	from, to := GoString(cfrom), GoString(cto)
	err := nox_client_copySave(from, to)
	if err != nil {
		return 0
	}
	return 1
}

func nox_client_copySave(from, to string) error {
	fromDir := datapath.Save(from)
	if err := nox_savegame_rm(to, true); err != nil {
		return err
	}
	toDir, err := nox_client_makeSaveDir(to)
	if err != nil {
		return err
	}
	// TODO: this should actually only copy the player file and map dirs; we just copy the whole dir instead
	//       implement it properly later, if it makes any difference
	return ifs.CopyDir(fromDir, toDir)
}

func nox_client_checkSaveMapExistsTmp(name string) (string, error) {
	name = strings.ToLower(name)
	if i := strings.IndexByte(name, ':'); i > 0 {
		name = name[:i]
	}
	dir := strings.TrimSuffix(name, ".map")
	fname := datapath.Save(common.SaveTmp, dir, name)
	if _, err := ifs.Stat(fname); err != nil {
		return "", err
	}
	return fname, nil
}

//export nox_client_countPlayerFiles04_4DC7D0
func nox_client_countPlayerFiles04_4DC7D0() C.int {
	cnt, err := nox_client_countPlayerFiles(0x4)
	if err != nil {
		gameLog.Println("count player files:", err)
		return 0
	}
	return C.int(cnt)
}

func nox_client_countPlayerFiles(flag byte) (int, error) {
	dir := datapath.Save()
	list, err := ifs.ReadDir(dir)
	if os.IsNotExist(err) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	cnt := 0
	for _, fi := range list {
		if fi.IsDir() {
			continue
		}
		if strings.ToLower(filepath.Ext(fi.Name())) != ".plr" {
			continue
		}
		fname := filepath.Join(dir, fi.Name())
		if sub_41A000_check0(fname)&flag != 0 {
			cnt++
		}
	}
	return cnt, nil
}

func sub_41A000_check0(path string) byte {
	path = ifs.Denormalize(path)
	cstr := CString(path)
	defer StrFree(cstr)
	save, freeSave := alloc.Make([]byte{}, 1280)
	defer freeSave()
	C.sub_41A000(cstr, (*C.nox_savegame_xxx)(unsafe.Pointer(&save[0])))
	return save[0]
}

//export nox_xxx_savePlayerMB_41C8F0
func nox_xxx_savePlayerMB_41C8F0(cdata *C.char, csz C.uint) C.int {
	data := unsafe.Slice((*byte)(unsafe.Pointer(cdata)), int(csz))
	path := memmap.String(0x85B3FC, 10984)
	const expSize = 216 // TODO: limit was *getMemIntPtr(0x587000, 55984) == 700, see #304
	if !noxflags.HasGame(noxflags.GameHost) && len(data) < expSize {
		str := strMan.GetStringInFile("Wol.c:WolApierror", "XferPlyr.c")
		noxConsole.Print(console.ColorBlue, str)
		str2 := strMan.GetStringInFile("Wol.c:Wolapierror", "XferPlyr.c")
		nox_xxx_printCentered_445490(str2)
		saveLog.Printf("SavePlayerOnClient: Error - character file too small %q (%d vs %d)\n", path, len(data), expSize)
		return 0
	}

	ifs.MkdirAll(datapath.Save())

	f, err := BinfileOpen(path, BinFileWO)
	if err != nil {
		saveLog.Printf("SavePlayerOnClient: Unable to open file %q\n", path)
		return 0
	}
	if err := f.SetKey(crypt.SaveKey); err != nil {
		_ = f.Close()
		saveLog.Printf("SavePlayerOnClient: Unable to key file %q: %v\n", path, err)
		return 0
	}
	if len(data) != 0 {
		f.Write(data)
	}
	f.Close()

	if noxflags.HasGame(noxflags.GameModeQuest) {
		*memmap.PtrUint8(0x85B3FC, 12257) = byte(C.dword_5d4594_1049844)
	} else {
		*memmap.PtrUint8(0x85B3FC, 12257) = 0
	}
	if C.nox_xxx_mapSavePlayerDataMB_41A230(internCStr(path)) == 0 {
		networkLogPrint("SavePlayerOnClient: Unable to save client data to file\n")
		return 0
	}
	str := strMan.GetStringInFile("CharacterSaved", "XferPlyr.c")
	noxConsole.Print(console.ColorBlue, str)
	str2 := strMan.GetStringInFile("CharacterSaved", "XferPlyr.c")
	nox_xxx_printCentered_445490(str2)
	return 1
}

//export sub_446140
func sub_446140() C.int {
	C.dword_5d4594_825764 = 1
	if noxflags.HasGame(noxflags.GameModeQuest) {
		if noxflags.HasGame(noxflags.GameHost) {
			sub_4D6B10(false)
		}
	}
	return 1
}

//export sub_419EB0
func sub_419EB0(i, val C.int) {
	if val == 1 {
		*memmap.PtrUint32(0x5D4594, 527716) |= 1 << i
	} else {
		*memmap.PtrUint32(0x5D4594, 527716) &^= 1 << i
	}
}

//export sub_4DB9C0
func sub_4DB9C0() {
	var next *Object
	for it := noxServer.firstServerObject(); it != nil; it = next {
		next = it.Next()
		if C.nox_xxx_isUnit_4E5B50(it.CObj()) != 0 {
			it.Delete()
		}
	}
	next = nil
	for it := firstServerObjectUpdatable2(); it != nil; it = next {
		next = it.Next()
		if C.sub_4E5B80(it.CObj()) != 0 {
			it.Delete()
		}
	}
}

func sub_4738D0() int {
	C.nox_xxx_bookHideMB_45ACA0(1)
	return 1
}

func nox_xxx_soloLoadGame_4DB7E0_savegame(a1 string) bool {
	if a1 == "" {
		return false
	}
	if noxflags.HasGame(noxflags.GameClient) && sub_4738D0() == 0 {
		return false
	}
	sub_4DB9C0()
	if a1 != common.SaveTmp && nox_client_copySave(a1, common.SaveTmp) != nil {
		return false
	}
	path := datapath.Save(common.SaveTmp, common.PlayerFile)
	if _, err := ifs.Stat(path); os.IsNotExist(err) {
		str := strMan.GetStringInFile("AutoSaveNotFound", "SaveGame.c")
		PrintToPlayers(str)
		return false
	}
	v5, _ := sub41D090(path)
	noxServer.SetFirstObjectScriptID(server.ObjectScriptID(v5))
	nox_server_ResetObjectGIDs_4E3C70()
	nox_xxx_gameSetSwitchSolo_4DB220(1)
	nox_xxx_gameSetNoMPFlag_4DB230(1)
	if C.nox_xxx_cliPlrInfoLoadFromFile_41A2E0(internCStr(path), noxMaxPlayers-1) == nil {
		return false
	}
	C.nox_xxx_cliPrepareGameplay1_460E60()
	C.nox_xxx_cliPrepareGameplay2_4721D0()
	pl := noxServer.getPlayerByInd(noxMaxPlayers - 1)
	pl.Name()
	mname := pl.saveName()
	noxServer.nox_xxx_gameSetMapPath_409D70(mname + ".map")
	dword_5d4594_1559960 = datapath.Save(common.SaveTmp, mname, mname+".map")
	noxServer.switchMap(mname + ".map")
	nox_xxx_cliPlayMapIntro_44E0B0(0)
	if C.nox_xxx_plrLoad_41A480(internCStr(path)) == 0 {
		return false
	}
	guiCon.Clear()
	C.sub_445450()
	C.nox_xxx_destroyEveryChatMB_528D60()
	str := strMan.GetStringInFile("GameLoaded", "C:\\NoxPost\\src\\Server\\Xfer\\SaveGame\\SaveGame.c")
	PrintToPlayers(str)
	return true
}

func nox_xxx_gameSetSoloSavePath_4DB270(a1 string) {
	dword_5d4594_1559960 = a1
}

//export sub_4DD0B0
func sub_4DD0B0(a1p *nox_object_t) {
	u := asUnitC(a1p)
	if u == nil {
		return
	}
	pl := u.ControllingPlayer()
	if nox_xxx_player_4D7980(pl.Index()) {
		noxServer.getPlayerByInd(pl.Index()).Disconnect(4)
	} else {
		sub_419EB0(C.int(pl.Index()), 0)
		nox_xxx_sendGauntlet_4DCF80(pl.Index(), 0)
	}
}

func nox_xxx_player_4D7980(a1 int) bool {
	return (*memmap.PtrUint32(0x5D4594, 1556300) & (1 << a1)) != 0
}

//export nox_xxx_saveMakePlayerLocation_4DB600
func nox_xxx_saveMakePlayerLocation_4DB600(a1 unsafe.Pointer) bool {
	pl := noxServer.getPlayerByInd(noxMaxPlayers - 1)
	if pl == nil {
		return false
	}
	u := pl.UnitC()
	if u == nil {
		return false
	}
	obj := noxServer.newObjectByTypeID("SaveGameLocation")
	if obj == nil {
		return false
	}
	pos := u.Pos()
	if a1 != nil {
		ptr := *(*unsafe.Pointer)(unsafe.Add(a1, 700))
		pos = types.Pointf{
			X: *(*float32)(unsafe.Add(ptr, 80)),
			Y: *(*float32)(unsafe.Add(ptr, 84)),
		}
	}
	nox_xxx_createAt_4DAA50(obj, nil, pos)
	C.nox_xxx_unitsNewAddToList_4DAC00()
	obj.script_id = C.int(u.ScriptID())
	var next *Object
	for it := u.FirstOwned516(); it != nil; it = next {
		next = u.NextOwned512()
		if it.Flags().Has(object.FlagActive) {
			it.SetOwner(obj)
		}
	}
	return true
}
