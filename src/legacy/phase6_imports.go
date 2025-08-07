package legacy

/*
#include "defs.h"

void sub_486620(void* a1);
int sub_486550(void* a1);
void sub_486570(void* a1, void* a2);
void* sub_4864A0(void* a3);
void sub_4BD7A0(void* lpMem);
int sub_486520(void* a2);
void nox_common_list_remove_425920(void* a1);
*/
import "C"
import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/audio"
)

var (
	Phase6Module *audio.Phase6Module
)

func initPhase6() {
	Phase6Module = audio.NewPhase6Module(
		"phase6",
		PlatformTicks,
		sub_486620,
		sub_486550,
		sub_486570,
		sub_4864A0,
		sub_4BD7A0,
		sub_486520,
		nox_common_list_remove_425920,
	)
}

//export sub_4873C0
func sub_4873C0(a3 int32) int32 {
	return Phase6Module.Sub_4873C0(a3)
}

//export sub_4BDA60
func sub_4BDA60(lpMem_ *C.struct312) {
	Phase6Module.Sub_4BDA60((*audio.Struct312)(unsafe.Pointer(lpMem_)))
}

//export sub_4BDA80
func sub_4BDA80(a1_ *C.struct312) int32 {
	return Phase6Module.Sub_4BDA80((*audio.Struct312)(unsafe.Pointer(a1_)))
}
