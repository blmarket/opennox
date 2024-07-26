#include "aaa1.h"
#include "dynarray.h"
#include "memmap.h"

typedef struct s576 {
    nox_list_item_t list_item;
    uint32_t field_3[43]; // 12~184
    uint32_t timer_46[3][8]; // 184~280
    uint32_t field_70[74]; // 280~576
} s576;

extern uint32_t dword_5d4594_1045420;
extern uint32_t dword_5d4594_1045424;
extern uint32_t dword_5d4594_1045428;
extern uint32_t dword_5d4594_1045432;
extern uint32_t dword_5d4594_1045436;
extern uint32_t dword_587000_126996;

void nox_common_list_remove_425920(void* a1);
int sub_4521F0();
void nox_common_list_clear_425760(nox_list_item_t* list);
uint32_t* sub_4BD340(int a1, int a2, int a3, int a4);
int sub_451920(uint32_t* a2);
char* nox_xxx_getSndName_40AF80(int a1);
void sub_4BD3C0(void* lpMem);
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
void* sub_4864A0(void* a3);
void* sub_425770(void* a1);
int***** sub_452230();

//----- (00452300) --------------------------------------------------------
uint32_t* nox_xxx_draw_452300(uint32_t* a1) {
    s576* v1p;
	uint32_t* v1; // esi

	if (!dword_5d4594_1045432) {
		return 0;
	}
	if (!dword_587000_126996) {
		return 0;
	}
	if (!*a1) {
		return 0;
	}
	v1 = sub_4BD2E0(*(uint32_t***)&dword_5d4594_1045436);
	if (!v1) {
		sub_452230();
		v1 = sub_4BD2E0(*(uint32_t***)&dword_5d4594_1045436);
		if (!v1) {
			return 0;
		}
	}
	memset(v1, 0, 0x240u);
	v1[9] = a1;
	sub_425770(v1); // seems multiple structs are sharing the logic
	v1[7] = 0;
	v1[75] = 0;
	v1[142] = 0;
	v1[108] = 0;
	v1[42] = 0;
	sub_4864A0(v1 + 46);
	nox_common_list_append_4258E0((int)getMemAt(0x5D4594, 840612), v1);
	v1[70] = (*getMemU32Ptr(0x587000, 127000))++;
	return v1;
}

//----- (00451FE0) --------------------------------------------------------
int sub_451FE0(int a1) {
	nox_common_list_remove_425920((uint32_t**)a1);
	*(uint32_t*)(a1 + 280) = 0;
	return sub_4BD300(*(uint32_t**)&dword_5d4594_1045436, a1);
}

//----- (00451970) --------------------------------------------------------
void sub_451970() {
	sub_4521F0();
	sub_452230();
	if (dword_5d4594_1045424) {
		sub_4BD3C0(*(void**)&dword_5d4594_1045424);
		dword_5d4594_1045424 = 0;
	}
	if (dword_5d4594_1045436) {
		sub_4BD2D0(*(void**)&dword_5d4594_1045436);
		dword_5d4594_1045436 = 0;
	}
	dword_5d4594_1045432 = 0;
}

//----- (00451850) --------------------------------------------------------
int sub_451850(int a2, void* a3p) {
	int a3 = a3p;
	int v2;            // edi
	unsigned char* v3; // esi
	int result;        // eax

	v2 = 0;
	v3 = getMemAt(0x5D4594, 840712);
	do {
		sub_451920((uint32_t*)v3 - 21);
		*(uint32_t*)v3 = nox_xxx_getSndName_40AF80(v2);
		v3 += 200;
		++v2;
	} while ((int)v3 < (int)getMemAt(0x5D4594, 1045312));
	dword_5d4594_1045420 = a3;
	dword_5d4594_1045428 = a2;
	if (a3) {
		dword_5d4594_1045424 = sub_4BD340(a3, 0x100000, 200, 0x2000);
		dword_5d4594_1045436 = sub_4BD280(200, 576);
	}
	if (!dword_5d4594_1045424 || !dword_5d4594_1045420 || !dword_5d4594_1045428 || !dword_5d4594_1045436) {
		return 0;
	}
	nox_common_list_clear_425760(getMemAt(0x5D4594, 840612));
	sub_4864A0(getMemAt(0x5D4594, 1045228));
	result = 1;
	*(uint32_t*)(dword_5d4594_1045428 + 184) = getMemAt(0x5D4594, 1045228);
	dword_5d4594_1045432 = 1;
	return result;
}
