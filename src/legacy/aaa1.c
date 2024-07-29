#include "defs.h"
#include "operators.h"
#include "static_assert.h"
#include <stdint.h>

typedef struct timerGroup {
	uint32_t field_0[24];
} timerGroup;

typedef struct s264 s264;

typedef struct s264 {
    nox_list_item_t field_0;
	uint32_t field_12;
	uint32_t field_16;
	void* field_20;
	void* field_24;
	uint32_t field_28;
	uint32_t field_32;
	uint32_t field_36[6];
	uint32_t field_60[7];
	timerGroup field_88;
	uint32_t field_184;
	uint32_t field_188;
    uint32_t field_192;
    uint32_t field_196;
	nox_list_item_t field_200;
	uint32_t field_212;
	int (*field_216)(s264*);
	uint32_t field_220;
	uint32_t field_224;
	uint32_t field_228;
	uint32_t field_232;
	uint32_t field_236;
	uint32_t field_240;
	uint32_t field_244;
	uint32_t field_248;
	uint32_t field_252;
	uint32_t field_256;
	uint32_t field_260;
} s264;
_Static_assert(sizeof(s264) == 264, "size of s264 is not 264");

void sub_486620(void* a1);  // timerGroup
void* sub_4864A0(void* a3); // timerGroup
int sub_486550(void* a1);   // timerGroup
int sub_486520(void* a2);   // timerGroup
int* sub_487360(int a1, int** a2, int* a3);
void* sub_425770(void* a1);
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
void nox_common_list_clear_425760(nox_list_item_t* list);
void sub_4872C0(s264* lpMem);
void sub_487590(s264* a1p, const void* a2);
void nullsub_10(uint32_t a1);
void sub_4BD840(int a3);
int sub_487910(s264* a1, int a2);
int* sub_4877D0(s264* a1, int* a2);
int* sub_4877F0(int** a1);
void sub_4BDA60(void* lpMem);
void sub_487310(s264* a1);
void sub_4876A0(void* a1);
int sub_4BDA80(int a1);
uint32_t* sub_4871C0(int a1, int a2, const void* a3);

s264* dword_5d4594_805984 = 0;

//----- (00487680) --------------------------------------------------------
void sub_487680(s264* lpMem) {
	sub_4876A0(lpMem);
	sub_4872C0(lpMem);
}

//----- (00431270) --------------------------------------------------------
void sub_431270() {
	if (dword_5d4594_805984) {
		sub_487680(dword_5d4594_805984);
		dword_5d4594_805984 = 0;
	}
}

//----- (004877D0) --------------------------------------------------------
int* sub_4877D0(s264 *a1p, int* a2) {
	// int a1 = a1p;
	int* result; // eax

	result = nox_common_list_getFirstSafe_425890(&a1p->field_200);
	*a2 = (int)result;
	return result;
}

//----- (00487970) --------------------------------------------------------
int* sub_487970(s264* a1, int a2) {
	int* result; // eax
	int* v3;     // edi
	int v4;      // ebx
	int* v5;     // esi
	int* aa1;

	result = sub_4877D0(a1, &aa1);
	v3 = result;
	if (result) {
		v4 = a2;
		do {
			result = sub_4877F0((int**)&aa1);
			v5 = result;
			if (v4 == -1 || v3[3] == v4) {
				result = (int*)sub_4BDA80((int)v3);
			}
			v3 = v5;
		} while (v5);
	}
	return result;
}

//----- (00431290) --------------------------------------------------------
void sub_431290() {
	if (dword_5d4594_805984) {
		sub_487970(dword_5d4594_805984, -1);
	}
}

#include <stdio.h>

//----- (00487150) --------------------------------------------------------
s264* sub_487150(int a1, const void* a2) {
	// a1 is always -1
	int v2;       // edi
	uint32_t* v3; // esi
	uint32_t* v4; // eax
	int v6;       // [esp+8h] [ebp-4h]

	v2 = a1;
	if (a1 == -1) {
		v2 = 0;
	}
	sub_487360(v2, (int**)&a1, &v6);
	printf("%x\n", a1);
	if (!a1) {
		return 0;
	}
	v3 = *(uint32_t**)(a1 + 4 * v6 + 24);
	if (!v3) {
		v4 = sub_4871C0(a1, v6, a2);
		v3 = v4;
		if (!v4) {
			return 0;
		}
		v4[47] = v2;
		sub_487310(v4);
	}
	++v3[4];
	return v3;
}

//----- (004872C0) --------------------------------------------------------
void sub_4872C0(s264* a1p) {
	void* lpMem = a1p;
	int v1; // eax
	int v2; // ecx

	sub_487910((int)lpMem, -1);
	(*(void (**)(void*))(*(uint32_t*)(*((uint32_t*)lpMem + 5) + 12) + 32))(lpMem);
	*(uint32_t*)(*((uint32_t*)lpMem + 5) + 4 * *((uint32_t*)lpMem + 6) + 24) = 0;
	v1 = *((uint32_t*)lpMem + 5);
	v2 = *(uint32_t*)(v1 + 16) - 1;
	*(uint32_t*)(v1 + 16) = v2;
	if (v2 < 0) {
		*(uint32_t*)(*((uint32_t*)lpMem + 5) + 16) = 0;
	}
	free(lpMem);
}
