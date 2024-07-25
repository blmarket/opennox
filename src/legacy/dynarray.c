#include <stdlib.h>

#include "dynarray.h"

//----- (004BD300) --------------------------------------------------------
int sub_4BD300(uint32_t* a1, int a2) {
	int result; // eax

	result = a2 - 4;
	*(uint32_t*)(a2 - 4) = *a1;
	*a1 = a2 - 4;
	return result;
}

//----- (004BD2D0) --------------------------------------------------------
void sub_4BD2D0(void* lpMem) { free(lpMem); }

//----- (004BD2E0) --------------------------------------------------------
uint32_t* sub_4BD2E0(void* a1p) {
	uint32_t** a1 = a1p;
	uint32_t* result; // eax
	uint32_t* v2;     // edx

	result = *a1;
	if (*a1) {
		v2 = (uint32_t*)*result;
		++result;
		*a1 = v2;
	}
	return result;
}

//----- (004BD280) --------------------------------------------------------
uint32_t* sub_4BD280(int a1, int a2) {
	int v2;           // esi
	uint32_t* result; // eax
	uint32_t* v4;     // ecx
	int v5;           // edi

	v2 = a2 + 4;
	result = calloc(1, a1 * (a2 + 4) + 4);
	if (result) {
		v4 = result + 1;
		*result = result + 1;
		if (a1 != 1) {
			v5 = a1 - 1;
			do {
				--v5;
				*v4 = (char*)v4 + v2;
				v4 = (uint32_t*)((char*)v4 + v2);
			} while (v5);
		}
		*v4 = 0;
	}
	return result;
}
