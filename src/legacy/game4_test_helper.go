package legacy

/*
#include <stdint.h>
#include <stdlib.h>
#include "GAME4.h"

extern uint32_t dword_5d4594_1599596;
extern uint32_t dword_5d4594_1599656;
extern void* dword_5d4594_1599576;
extern void* dword_5d4594_1599588;
extern void* dword_5d4594_1599592;
extern void* dword_5d4594_1599540;

char sub_506870(int a1, int a2, wchar2_t* a3);
char sub_5068E0(int a1, int a2, wchar2_t* a3);
void sub_506C90(int a1, int a2, wchar2_t* a3);
void sub_506D00(int a1, wchar2_t* a2);
void sub_506DE0(int a1);
void sub_506E50(int a1, wchar2_t* a2);
int sub_507000(int a1);
int nox_xxx_spellGetPhoneme_4FE1C0(int a1, char a2);
*/
import "C"

import "unsafe"

type game4SafePathsResult struct {
	initResult        int
	missingName       int
	negativeIndex     uintptr
	zeroIndex         uintptr
	count             int
	setPrimary        int
	clearPrimary      int
	setSecondary      int
	clearSecondary    int
	invalidSave       int
	closeEmpty        uintptr
	seekEmpty         uintptr
	invalidFirstCoord float64
	invalidLastCoord  float64
	moveEmpty         int
	placeEmpty        int
	objectField       int
	nodeField         int
	removeNil         int
	removeMissing     int
	voteGuard         int8
	voteThreshold     int
	votesActive       bool
}

func C_game4SafePaths() (res game4SafePathsResult) {
	savedCount := C.dword_5d4594_1599596
	savedVotes := C.dword_5d4594_1599656
	savedEntries := C.dword_5d4594_1599576
	savedPrimary := C.dword_5d4594_1599588
	savedSecondary := C.dword_5d4594_1599592
	savedUnits := C.dword_5d4594_1599540
	C.dword_5d4594_1599596 = 0
	C.dword_5d4594_1599656 = 0
	C.dword_5d4594_1599576 = nil
	C.dword_5d4594_1599588 = nil
	C.dword_5d4594_1599592 = nil
	C.dword_5d4594_1599540 = nil
	defer func() {
		C.free(C.dword_5d4594_1599576)
		C.free(C.dword_5d4594_1599588)
		C.free(C.dword_5d4594_1599592)
		C.dword_5d4594_1599596 = savedCount
		C.dword_5d4594_1599656 = savedVotes
		C.dword_5d4594_1599576 = savedEntries
		C.dword_5d4594_1599588 = savedPrimary
		C.dword_5d4594_1599592 = savedSecondary
		C.dword_5d4594_1599540 = savedUnits
	}()

	res.initResult = int(C.sub_502B10())
	missing := C.CString("missing")
	primary := C.CString("primary.map")
	secondary := C.CString("secondary.map")
	defer C.free(unsafe.Pointer(missing))
	defer C.free(unsafe.Pointer(primary))
	defer C.free(unsafe.Pointer(secondary))
	res.missingName = int(C.sub_5029A0(missing))
	res.negativeIndex = uintptr(uint32(C.sub_5029F0(-1)))
	res.zeroIndex = uintptr(uint32(C.sub_5029F0(0)))
	res.count = int(C.sub_502A20())
	res.setPrimary = int(C.sub_502A50(primary))
	res.clearPrimary = int(C.sub_502A50(nil))
	res.setSecondary = int(C.sub_502AB0(secondary))
	res.clearSecondary = int(C.sub_502AB0(nil))
	res.invalidSave = int(C.sub_502D70(-1))
	res.closeEmpty = uintptr(unsafe.Pointer(C.sub_502DF0()))
	res.seekEmpty = uintptr(unsafe.Pointer(C.sub_502E10(-1)))
	res.invalidFirstCoord = float64(C.sub_502E70(-1))
	res.invalidLastCoord = float64(C.sub_502EA0(-1))

	res.moveEmpty = int(C.sub_504560(3, -4))
	res.placeEmpty = int(C.sub_504910(5, -6))
	object := C.calloc(1, 452)
	node := C.calloc(1, 8)
	defer C.free(object)
	defer C.free(node)
	*(*uint32)(unsafe.Pointer(uintptr(object) + 444)) = 77
	*(*uint32)(unsafe.Pointer(uintptr(node) + 4)) = 88
	res.objectField = int(C.sub_5049C0(C.int(uintptr(object))))
	_ = C.sub_5049D0()
	res.nodeField = int(C.sub_5049E0(C.int(uintptr(node))))
	res.removeNil = int(C.sub_504A10(0))
	res.removeMissing = int(C.sub_504A10(C.int(uintptr(object))))

	res.voteGuard = int8(C.sub_506870(0, 0, nil))
	_ = C.sub_5068E0(0, 0, nil)
	C.sub_506C90(0, 0, nil)
	C.sub_506D00(0, nil)
	C.sub_506DE0(0)
	C.sub_506E50(0, nil)
	vote := C.calloc(1, 32)
	defer C.free(vote)
	res.voteThreshold = int(C.sub_507000(C.int(uintptr(vote))))
	res.votesActive = C.sub_5071C0() != 0
	return res
}

func C_game4SpellPhoneme(netCode int, phoneme int8) int {
	return int(C.nox_xxx_spellGetPhoneme_4FE1C0(C.int(netCode), C.char(phoneme)))
}
