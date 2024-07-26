#include "aaa1.h"
#include "dynarray.h"
#include "memmap.h"
#include "operators.h"

// Goal: Keep all references of 1045436 and 840612 in this file..
// 1045436: DONE

typedef struct s576 {
    nox_list_item_t list_item;
    uint32_t field_12;
    uint32_t field_16;
    uint32_t field_20;
    uint32_t field_24;
    uint32_t field_28;
    uint32_t field_32;
    uint32_t field_36;
    uint32_t field_40[36]; // 40~184
    uint32_t timer_184[3][8]; // 184~280
    uint32_t field_280; // 280~284
    uint32_t field_284[73]; // 284~576
} s576;

extern uint32_t dword_5d4594_1045420;
extern uint32_t dword_5d4594_1045424;
extern uint32_t dword_5d4594_1045428;
extern uint32_t dword_5d4594_1045432;
extern dynarray* dword_5d4594_1045436;
extern uint32_t dword_587000_126996;
extern void* dword_587000_127004;

int sub_4BDA80(int a1);
int sub_4BDB30(int a1);
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
int* sub_4521A0(int a1);
int sub_452490(uint32_t* a1);
int sub_451CA0(uint32_t* a1);
int sub_451DC0(int a1);
int sub_486350(void* a1, int a2);
void sub_452050(uint32_t* a1);
int sub_452010();
void sub_452510(int a3);
int sub_451BE0(int a1);
int sub_486520(void* a2);
void sub_4523D0(s576* a1);

//----- (00452190) --------------------------------------------------------
void sub_452190(int a1) { nox_common_list_remove_425920((uint32_t**)(a1 + 112)); }

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
	v1 = sub_4BD2E0(dword_5d4594_1045436);
	if (!v1) {
		sub_452230();
		v1 = sub_4BD2E0(dword_5d4594_1045436);
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
int sub_451FE0(s576* a1) {
	nox_common_list_remove_425920(&a1->list_item);
    a1->field_280 = 0;
	return sub_4BD300(dword_5d4594_1045436, a1);
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

//----- (00452120) --------------------------------------------------------
int* sub_452120(int a1) {
	int v1;            // ebp
	int* result;       // eax
	int* v3;           // ebx
	unsigned char* v4; // esi
	unsigned char* v5; // edi

	v1 = 0;
	result = sub_4521A0(*(uint32_t*)(a1 + 300) + *(uint32_t*)(*(uint32_t*)(a1 + 36) + 48));
	v3 = result;
	if (result) {
		sub_452190((int)result);
		v4 = *(unsigned char**)getMemAt(0x5D4594, 840612);
		if (*(unsigned char**)getMemAt(0x5D4594, 840612) != getMemAt(0x5D4594, 840612)) {
			do {
				v5 = *(unsigned char**)v4;
				if (*((int**)v4 + 9) == v3) {
					sub_4523D0(v4);
					sub_451FE0((int)v4);
					v1 = 1;
				}
				v4 = v5;
			} while (v5 != getMemAt(0x5D4594, 840612));
		}
		result = (int*)v1;
	}
	return result;
}

//----- (004521F0) --------------------------------------------------------
int sub_4521F0() {
	int result;        // eax
	unsigned char* v1; // esi
	unsigned char* v2; // edi

	result = dword_5d4594_1045432;
	if (dword_5d4594_1045432) {
		v1 = *(unsigned char**)getMemAt(0x5D4594, 840612);
		if (*(unsigned char**)getMemAt(0x5D4594, 840612) != getMemAt(0x5D4594, 840612)) {
			do {
				v2 = *(unsigned char**)v1;
				sub_4523D0(v1);
				result = sub_451FE0((int)v1);
				v1 = v2;
			} while (v2 != getMemAt(0x5D4594, 840612));
		}
	}
	return result;
}

//----- (00452230) --------------------------------------------------------
int***** sub_452230() {
	int***** result; // eax
	int**** v1;      // esi

	result = *(int******)&dword_5d4594_1045432;
	if (dword_5d4594_1045432) {
		result = *(int******)getMemAt(0x5D4594, 840612);
		if (*(unsigned char**)getMemAt(0x5D4594, 840612) != getMemAt(0x5D4594, 840612)) {
			do {
				v1 = *result;
				if ((uint8_t)result[6] & 1) {
					sub_451FE0((int)result);
				}
				result = (int*****)v1;
			} while (v1 != (int****)getMemAt(0x5D4594, 840612));
		}
	}
	return result;
}

//----- (004519C0) --------------------------------------------------------
void sub_4519C0() {
	int result;        // eax
	int v1;            // esi
	int v2;            // eax
	int v3;            // ebp
	int v4;            // eax
	unsigned char* v5; // edi
	unsigned char* v6; // esi
	unsigned char* v7; // edi
	int v8;            // eax
	int v9;            // eax
	int v10;           // eax

	result = dword_5d4594_1045432;
	if (!dword_5d4594_1045432) {
		return;
	}
	result = *getMemU32Ptr(0x5D4594, 1045448);
	if (*getMemU32Ptr(0x5D4594, 1045448)) {
		return;
	}
	*getMemU32Ptr(0x5D4594, 1045448) = 1;
	sub_486520(*(unsigned int**)&dword_587000_127004);
	v1 = *getMemU32Ptr(0x5D4594, 840612);
	++*getMemU32Ptr(0x5D4594, 1045440);
	if (*(unsigned char**)getMemAt(0x5D4594, 840612) != getMemAt(0x5D4594, 840612)) {
		do {
			v2 = *(uint32_t*)(v1 + 36);
			if (*(uint32_t*)(v2 + 100) != *getMemU32Ptr(0x5D4594, 1045440)) {
				nox_common_list_clear_425760((uint32_t*)(v2 + 88));
				*(uint32_t*)(*(uint32_t*)(v1 + 36) + 52) = 0;
				*(uint32_t*)(*(uint32_t*)(v1 + 36) + 100) = *getMemU32Ptr(0x5D4594, 1045440);
			}
			sub_486520((unsigned int*)(v1 + 184));
			if (*(uint32_t*)(v1 + 28) != 4) {
				sub_451BE0(v1);
			}
			v1 = *(uint32_t*)v1;
		} while ((unsigned char*)v1 != getMemAt(0x5D4594, 840612));
		v1 = *getMemU32Ptr(0x5D4594, 840612);
		if (*(unsigned char**)getMemAt(0x5D4594, 840612) != getMemAt(0x5D4594, 840612)) {
			do {
				sub_452510(v1);
				v1 = *(uint32_t*)v1;
			} while ((unsigned char*)v1 != getMemAt(0x5D4594, 840612));
			v1 = *getMemU32Ptr(0x5D4594, 840612);
		}
	}
	v3 = 0;
	sub_452010();
	if ((unsigned char*)v1 != getMemAt(0x5D4594, 840612)) {
		do {
			v4 = *(uint32_t*)(v1 + 176);
			v5 = *(unsigned char**)v1;
			if (!v4 || v1 != *(uint32_t*)(v4 + 152)) {
				sub_4523D0((uint32_t*)v1);
			}
			if (*(uint8_t*)(v1 + 24) & 1) {
				sub_451FE0(v1);
			} else {
				v3 += (unsigned int)(33 * (*(uint32_t*)(*(uint32_t*)(v1 + 36) + 20) >> 16)) >> 14;
				sub_452050((uint32_t*)v1);
			}
			v1 = (int)v5;
		} while (v5 != getMemAt(0x5D4594, 840612));
	}
	if (v3 <= 100) {
		sub_486350((int)getMemAt(0x5D4594, 1045228), 0x4000);
	} else {
		sub_486350((int)getMemAt(0x5D4594, 1045228), 0x190000u / v3);
	}
	result = sub_486520(getMemUintPtr(0x5D4594, 1045228));
	v6 = *(unsigned char**)getMemAt(0x5D4594, 840612);
	if (*(unsigned char**)getMemAt(0x5D4594, 840612) != getMemAt(0x5D4594, 840612)) {
		do {
			v7 = *(unsigned char**)v6;
			result = *((uint32_t*)v6 + 7);
			if (result == 1) {
				sub_451DC0((int)v6);
				v8 = sub_451CA0(v6);
				*((uint32_t*)v6 + 74) = v8;
				if (!v8) {
					do {
						if (!sub_452120((int)v6)) {
							break;
						}
						v7 = *(unsigned char**)v6;
						sub_451DC0((int)v6);
						v9 = sub_451CA0(v6);
						*((uint32_t*)v6 + 74) = v9;
					} while (!v9);
				}
				v10 = sub_451CA0(v6);
				*((uint32_t*)v6 + 74) = v10;
				if (!v10 || (result = sub_452490(v6)) == 0) {
					sub_4523D0(v6);
					result = sub_451FE0((int)v6);
				}
			}
			v6 = v7;
		} while (v7 != getMemAt(0x5D4594, 840612));
	}
	*getMemU32Ptr(0x5D4594, 1045448) = 0;
}

//----- (004BD660) --------------------------------------------------------
int sub_4BD660(int a1) {
	int result; // eax

	result = *(uint32_t*)(a1 + 12) - 1;
	*(uint32_t*)(a1 + 12) = result;
	if (result < 0) {
		*(uint32_t*)(a1 + 12) = 0;
	}
	return result;
}

//----- (00451F90) --------------------------------------------------------
int sub_451F90(int a1) {
	int v1;     // edi
	int result; // eax
	int* v3;    // esi

	v1 = 0;
	result = *(uint32_t*)(a1 + 168);
	if (result <= 0) {
		*(uint32_t*)(a1 + 168) = 0;
	} else {
		v3 = (int*)(a1 + 40);
		do {
			sub_4BD660(*v3);
			*v3 = 0;
			result = *(uint32_t*)(a1 + 168);
			++v1;
			++v3;
		} while (v1 < result);
		*(uint32_t*)(a1 + 168) = 0;
	}
	return result;
}

//----- (00452410) --------------------------------------------------------
int sub_452410(int a1) {
	int result; // eax

	result = *(uint32_t*)(a1 + 176);
	if (result && a1 == *(uint32_t*)(result + 152)) {
		if (*(uint8_t*)(a1 + 24) & 2) {
			sub_4BDA80(*(uint32_t*)(a1 + 176));
		}
		sub_4BDB30(*(uint32_t*)(a1 + 176));
		*(uint32_t*)(*(uint32_t*)(a1 + 176) + 152) = 0;
		*(uint32_t*)(*(uint32_t*)(a1 + 176) + 148) = 0;
		result = *(uint32_t*)(a1 + 176);
		*(uint32_t*)(result + 140) = 0;
		*(uint32_t*)(*(uint32_t*)(a1 + 176) + 144) = 0;
		*(uint32_t*)(*(uint32_t*)(a1 + 176) + 112) = 0;
		*(uint32_t*)(a1 + 176) = 0;
	}
	return result;
}

//----- (004523D0) --------------------------------------------------------
void sub_4523D0(s576* a1p) {
	uint32_t* a1 = a1p;
	int result = 0; // eax

	if (!(a1p->field_24 & 1)) {
		sub_452410((int)a1);
		sub_451F90((int)a1);
        a1p->field_28 = 4;
        a1p->field_280 = 0;
        result = a1p->field_24;
		LOBYTE(result) = result | 1;
        a1p->field_24 = result;
	}
}
