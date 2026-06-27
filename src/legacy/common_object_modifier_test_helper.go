package legacy

/*
#include "nox_wchar.h"

void nullsub_22();
void nullsub_36();
void nullsub_38();
void nullsub_39();
void nullsub_40();
void nullsub_41();
void nullsub_42();
void nullsub_43();
void nullsub_44();

wchar2_t* sub_413480(char a1);
*/
import "C"
import "unsafe"

func C_nullsubs() {
	C.nullsub_22()
	C.nullsub_36()
	C.nullsub_38()
	C.nullsub_39()
	C.nullsub_40()
	C.nullsub_41()
	C.nullsub_42()
	C.nullsub_43()
	C.nullsub_44()
}

func C_sub_413480(a1 byte) unsafe.Pointer {
	return unsafe.Pointer(C.sub_413480(C.char(a1)))
}
