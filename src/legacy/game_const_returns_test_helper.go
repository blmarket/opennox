package legacy

/*
// Decompiled GAME*.c stubs that return a fixed constant regardless of input.
int nox_xxx_wndRetNULL_46A8A0();
int nox_xxx_wndRetNULL_0_46A8B0();
int nox_xxx_book_45BD30(int a1, int a2);
*/
import "C"

func C_nox_xxx_wndRetNULL_46A8A0() int   { return int(C.nox_xxx_wndRetNULL_46A8A0()) }
func C_nox_xxx_wndRetNULL_0_46A8B0() int { return int(C.nox_xxx_wndRetNULL_0_46A8B0()) }
func C_nox_xxx_book_45BD30(a1, a2 int) int {
	return int(C.nox_xxx_book_45BD30(C.int(a1), C.int(a2)))
}
