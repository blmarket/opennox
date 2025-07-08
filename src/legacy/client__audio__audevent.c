#include "client__audio__audevent.h"
#include "common__random.h"

#include "GAME2.h"
#include "GAME2_2.h"
#include "GAME3_1.h"

//----- (00452770) --------------------------------------------------------
int sub_452770(uint32_t* a1) {
	struct576* v1p;
	uint32_t* v1;    // esi
	uint32_t* v2;    // ebx
	int v4;          // eax
	unsigned int v5; // eax

	v1 = v1p = (uint32_t*)a1[38];
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
