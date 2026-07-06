package legacy

/*
#include "common__gamemech__pausefx.h"
#include <stdint.h>

void test_sub_57AF30(int a1, int a2) {
	sub_57AF30(a1, a2);
}

extern uint32_t dword_5d4594_2523804;
extern uint32_t dword_5d4594_2523776;
extern uint32_t dword_5d4594_2523780;

uint32_t test_get_pausefx_2523804() {
	return dword_5d4594_2523804;
}

void test_set_pausefx_2523804(uint32_t v) {
	dword_5d4594_2523804 = v;
}

uint32_t test_get_pausefx_2523776() {
	return dword_5d4594_2523776;
}

void test_set_pausefx_2523776(uint32_t v) {
	dword_5d4594_2523776 = v;
}

uint32_t test_get_pausefx_2523780() {
	return dword_5d4594_2523780;
}

void test_set_pausefx_2523780(uint32_t v) {
	dword_5d4594_2523780 = v;
}
*/
import "C"

func C_sub_57AF30(a1 int, a2 int) {
	C.test_sub_57AF30(C.int(a1), C.int(a2))
}

func C_get_pausefx_2523804() uint32 {
	return uint32(C.test_get_pausefx_2523804())
}

func C_set_pausefx_2523804(v uint32) {
	C.test_set_pausefx_2523804(C.uint32_t(v))
}

func C_get_pausefx_2523776() uint32 {
	return uint32(C.test_get_pausefx_2523776())
}

func C_set_pausefx_2523776(v uint32) {
	C.test_set_pausefx_2523776(C.uint32_t(v))
}

func C_get_pausefx_2523780() uint32 {
	return uint32(C.test_get_pausefx_2523780())
}

func C_set_pausefx_2523780(v uint32) {
	C.test_set_pausefx_2523780(C.uint32_t(v))
}
