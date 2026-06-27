package legacy

/*
#include <stdint.h>
#include <stdlib.h>

extern uint32_t dword_5d4594_815044;
extern uint32_t dword_5d4594_815748;
extern uint32_t dword_5d4594_816368;
extern uint32_t dword_5d4594_816372;
extern uint32_t dword_5d4594_816376;
extern uint32_t dword_587000_93156;

int sub_43B6D0();
int nox_sprintAddrPort_43BC80(const char* addr, unsigned short port, char* dst);
int sub_43BDB0();
int sub_43C650();
void sub_43CEB0();
int sub_43DA80();
void sub_43DAD0();
int sub_43DB20();
int sub_43DB30(int a1);
int sub_43DB60();
void sub_43DBA0();
int sub_43DC10();
int sub_43E8C0(int a1);
int sub_43F0E0(uint32_t* a1);
unsigned int nox_gui_xxx_check_446360();
*/
import "C"
import "unsafe"

type game13Globals struct {
	menuState  uint32
	tickCount  uint32
	musicSlot  uint32
	musicDepth uint32
	musicReady uint32
	musicFlag  uint32
}

func C_game13Globals() game13Globals {
	return game13Globals{
		menuState:  uint32(C.dword_5d4594_815044),
		tickCount:  uint32(C.dword_5d4594_815748),
		musicSlot:  uint32(C.dword_5d4594_816368),
		musicDepth: uint32(C.dword_5d4594_816372),
		musicReady: uint32(C.dword_5d4594_816376),
		musicFlag:  uint32(C.dword_587000_93156),
	}
}

func C_game13SetGlobals(v game13Globals) {
	C.dword_5d4594_815044 = C.uint32_t(v.menuState)
	C.dword_5d4594_815748 = C.uint32_t(v.tickCount)
	C.dword_5d4594_816368 = C.uint32_t(v.musicSlot)
	C.dword_5d4594_816372 = C.uint32_t(v.musicDepth)
	C.dword_5d4594_816376 = C.uint32_t(v.musicReady)
	C.dword_587000_93156 = C.uint32_t(v.musicFlag)
}

func C_sub_43B6D0() int {
	return int(C.sub_43B6D0())
}

func C_nox_sprintAddrPort_43BC80(addr string, port uint16) (int, string) {
	caddr := C.CString(addr)
	defer C.free(unsafe.Pointer(caddr))
	dst := C.malloc(128)
	defer C.free(dst)
	ret := int(C.nox_sprintAddrPort_43BC80(caddr, C.ushort(port), (*C.char)(dst)))
	return ret, C.GoString((*C.char)(dst))
}

func C_sub_43BDB0() int {
	return int(C.sub_43BDB0())
}

func C_sub_43C650() int {
	return int(C.sub_43C650())
}

func C_sub_43CEB0() {
	C.sub_43CEB0()
}

func C_sub_43DA80() int {
	return int(C.sub_43DA80())
}

func C_sub_43DAD0() {
	C.sub_43DAD0()
}

func C_sub_43DB20() int {
	return int(C.sub_43DB20())
}

func C_sub_43DB30(v int) int {
	return int(C.sub_43DB30(C.int(v)))
}

func C_sub_43DB60() int {
	return int(C.sub_43DB60())
}

func C_sub_43DBA0() {
	C.sub_43DBA0()
}

func C_sub_43DC10() int {
	return int(C.sub_43DC10())
}

func C_sub_43E8C0(v int) int {
	return int(C.sub_43E8C0(C.int(v)))
}

func C_sub_43F0E0(a1, a3, a4 uint32) int {
	buf := []uint32{0, a1, 0, a3, a4}
	return int(C.sub_43F0E0((*C.uint32_t)(unsafe.Pointer(&buf[0]))))
}

func C_nox_gui_xxx_check_446360() int {
	return int(C.nox_gui_xxx_check_446360())
}
