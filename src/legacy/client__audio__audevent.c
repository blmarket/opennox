#include "client__audio__audevent.h"
#include "common__random.h"

#include "GAME2.h"
#include "GAME2_2.h"
#include "GAME3_1.h"

//----- (00451E80) --------------------------------------------------------
int sub_451E80(int a1) {
	int v1;        // eax
	int v2;        // ebx
	int v3;        // eax
	int v4;        // ecx
	int v5;        // edx
	int v6;        // eax
	int v7;        // edx
	int v8;        // eax
	int v9;        // edi
	int v10;       // ecx
	uint32_t* v11; // eax

	v1 = *(uint32_t*)(a1 + 36);
	v2 = *(uint32_t*)(v1 + 4);
	if (*(int*)(a1 + 568) <= 0) {
		v3 = *(uint32_t*)(v1 + 192);
		v4 = 0;
		*(uint32_t*)(a1 + 568) = v3;
		if (v3 > 0) {
			v5 = a1 + 440;
			do {
				v5 += 4;
				v6 = v3 - v4++ - 1;
				*(uint32_t*)(v5 - 4) = v6;
				v3 = *(uint32_t*)(a1 + 568);
			} while (v4 < v3);
		}
	}
	v7 = *(uint32_t*)(a1 + 568) - 1;
	*(uint32_t*)(a1 + 568) = v7;
	if (!(v2 & 2)) {
		return *(uint32_t*)(a1 + 4 * v7 + 440);
	}
	v8 = nox_common_randomIntMinMax_415FF0(0, v7, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 431);
	v9 = *(uint32_t*)(a1 + 4 * v8 + 440);
	v10 = v8;
	if (v8 < *(int*)(a1 + 568)) {
		v11 = (uint32_t*)(a1 + 4 * v8 + 440);
		do {
			++v10;
			*v11 = v11[1];
			++v11;
		} while (v10 < *(int*)(a1 + 568));
	}
	return v9;
}

//----- (00452770) --------------------------------------------------------
int sub_452770(uint32_t* a1) {
	uint32_t* v1;    // esi
	uint32_t* v2;    // ebx
	int v4;          // eax
	unsigned int v5; // eax

	v1 = (uint32_t*)a1[38];
	v2 = (uint32_t*)sub_451CF0((uint32_t*)a1[38]);
	if (*(uint32_t*)(v1[9] + 72) < 0x21u) {
		sub_4BDB90(a1, v2);
		return 0;
	}
	sub_4BDB90(a1, 0);
	v4 = v1[9];
	if (!(*(uint8_t*)(v4 + 4) & 8) || v2 || v1[142]) {
		v5 = nox_common_randomIntMinMax_415FF0(*(uint32_t*)(v4 + 68), *(uint32_t*)(v4 + 72),
											   "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 706);
		if (v5 < 0x21) {
			sub_4BDB90(a1, v2);
			return 0;
		}
		v1[71] = v5;
		v1[74] = v2;
	}
	return 0;
}
