#include "defs.h"
#include "operators.h"

// Declarations added by compile_audio.py
extern uint32_t dword_587000_126996;
extern uint32_t dword_5d4594_1045420;
extern uint32_t dword_5d4594_1045424;
extern uint32_t dword_5d4594_1045428;
extern uint32_t dword_5d4594_1045432;
extern uint32_t dword_5d4594_1045436;
extern void* dword_587000_127004;
char* nox_xxx_getSndName_40AF80(int a1);
int nox_common_randomIntMinMax_415FF0(int min, int max, const char* file, int line);
int sub_451920(struct200* a2_);
int sub_451CF0(uint32_t* a1);
int sub_451E80(int a1);
int sub_451FE0(struct576* a1_);
int sub_4523D0(struct576* a1_);
int sub_452410(struct576* a1_);
int sub_452580(struct576* a1_);
int sub_4526D0(int a1);
int sub_452770(uint32_t* a1);
int sub_452FA0(int a1);
int sub_4862E0(void* a3, int a4);
int sub_486350(void* a1, int a2);
int sub_4863B0(void* a2);
int sub_486520(void* a2);
int sub_4BD300(uint32_t* a1, int a2);
int sub_4BD650(int a1);
int sub_4BD660(int a1);
int sub_4BD710(int a1);
int sub_4BDA80(int a1);
int sub_4BDB20(int a1);
int sub_4BDB30(int a1);
int sub_4BDB40(int a2);
int* sub_452810(int a1, char a2);
long long sub_452690(struct576* a3_, long long a4, int a5);
nox_list_item_t* nox_common_list_getFirstSafe_425890(nox_list_item_t* list);
uint32_t* sub_4BD280(int a1, int a2);
uint32_t* sub_4BD2E0(uint32_t** a1);
uint32_t* sub_4BD340(int a1, int a2, int a3, int a4);
uint32_t* sub_4BD470(uint32_t** a1, int a2);
unsigned int sub_452F10(struct576* a1_, int a2);
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
void nox_common_list_clear_425760(nox_list_item_t* list);
void nox_common_list_remove_425920(void* a1);
void sub_4BD2D0(void* lpMem);
void sub_4BD3C0(void* lpMem);
void sub_4BDB90(uint32_t* a1, uint32_t* a2);
void* sub_425770(void* a1);
void* sub_486320(void* a1, int a2);
void* sub_4864A0(void* a3);

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

int sub_451920(struct200* a2_) {
	uint32_t* a2 = a2_;
	a2_->field_0 = 0;
	a2_->field_1 = 0;
	a2_->field_2 = 0;
	a2_->field_14 = 0;
	a2_->field_15 = 0;
	a2_->field_19 = 0;
	a2_->field_20 = 0;
	a2_->field_12 = 1;
	a2_->field_48 = 0;
	a2_->field_18 = 0;
	a2_->field_17 = 0;
	a2_->field_25 = 0;
	a2_->field_26 = 0;
	a2_->field_16 = 600;
	return sub_4862E0((int)(&a2_->field_4), 0x4000);
}

int sub_452010() {
	unsigned char* v0; // esi
	int v1;            // ebx
	int v2;            // edi

	v0 = getMemAt(0x5D4594, 839892);
	v1 = 6;
	do {
		v2 = 10;
		do {
			nox_common_list_clear_425760(v0);
			v0 += 12;
			--v2;
		} while (v2);
		--v1;
	} while (v1);
	return ++*getMemU32Ptr(0x5D4594, 1045444);
}

void sub_452190(struct200* a1_) {
	int a1 = a1_;
	nox_common_list_remove_425920((uint32_t**)(&a1_->field_28));
}

int* sub_4521A0(int a1) {
	int v1;            // ebp
	unsigned char* v2; // ebx
	int v3;            // edi
	int* v4;           // esi
	int* v5;           // eax

	v1 = 0;
	v2 = getMemAt(0x5D4594, 839892);
	if (a1 > 0) {
		while (1) {
			v3 = 0;
			v4 = (int*)v2;
			do {
				v5 = nox_common_list_getFirstSafe_425890(v4);
				if (v5) {
					return v5 - 28;
				}
				++v3;
				v4 += 3;
			} while (v3 < 10);
			++v1;
			v2 += 120;
			if (v1 < a1) {
				continue;
			}
			break;
		}
	}
	return 0;
}

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

char* nox_xxx_draw_452270(int a1) {
	char* result; // eax

	if (dword_5d4594_1045432 && a1 >= 0 && a1 < 1023) {
		result = (char*)getMemAt(0x5D4594, 840628 + 200 * a1);
	} else {
		result = 0;
	}
	return result;
}

int sub_4526F0(int a1) {
	uint32_t* v1; // esi
	int v2;       // eax

	v1 = *(uint32_t**)(a1 + 152);
	v1[6] &= 0xFFFFFFFD;
	v2 = 4;
	if (v1[7] != 4) {
		if (v1[74] || v1[142]) {
			v2 = 1;
		} else {
			v1[71] = 0;
		}
		if (v1[71]) {
			sub_452690((int)v1, (unsigned int)v1[71], v2);
			v1[71] = 0;
			return 0;
		}
		v1[7] = v2;
	}
	return 0;
}

int sub_451CA0(struct576* a1_) {
	uint32_t* a1 = a1_;
	int v1;       // ecx
	int v3;       // eax
	uint32_t* v4; // ecx

	v1 = a1_->field_42;
	a1_->field_108 = v1;
	if (!v1) {
		return 0;
	}
	v3 = 0;
	if (v1 > 0) {
		v4 = &a1_->field_76;
		do {
			*v4 = v3++;
			++v4;
		} while (v3 < a1_->field_108);
	}
	a1_->field_43 = -1;
	return sub_451CF0(a1_);
}

int sub_451F30(struct576* a1_, int a2) {
	int a1 = a1_;
	int v2;     // edx
	int result; // eax

	*(uint32_t*)(&a1_->field_10[*(uint32_t*)(&a1_->field_42)]) =
		sub_4BD470(*(uint32_t***)&dword_5d4594_1045424, *(short*)(*(uint32_t*)(&a1_->field_9) + 2 * a2 + 128));
	v2 = *(uint32_t*)(&a1_->field_42);
	result = *(uint32_t*)(&a1_->field_10[v2]);
	if (result) {
		sub_4BD650(*(uint32_t*)(&a1_->field_10[v2]));
		result = *(uint32_t*)(&a1_->field_42) + 1;
		*(uint32_t*)(&a1_->field_42) = result;
	}
	return result;
}

int sub_451F90(struct576* a1_) {
	int a1 = a1_;
	int v1;     // edi
	int result; // eax
	int* v3;    // esi

	v1 = 0;
	result = *(uint32_t*)(&a1_->field_42);
	if (result <= 0) {
		*(uint32_t*)(&a1_->field_42) = 0;
	} else {
		v3 = (int*)(&a1_->field_10);
		do {
			sub_4BD660(*v3);
			*v3 = 0;
			result = *(uint32_t*)(&a1_->field_42);
			++v1;
			++v3;
		} while (v1 < result);
		*(uint32_t*)(&a1_->field_42) = 0;
	}
	return result;
}

int sub_451FE0(struct576* a1_) {
	int a1 = a1_;
	nox_common_list_remove_425920(a1_);
	*(uint32_t*)(&a1_->field_70) = 0;
	return sub_4BD300(*(uint32_t**)&dword_5d4594_1045436, a1_);
}

int* sub_452120(struct576* a1_) {
	int a1 = a1_;
	int v1;            // ebp
	int* result;       // eax
	int* v3;           // ebx
	unsigned char* v4; // esi
	unsigned char* v5; // edi

	v1 = 0;
	result = sub_4521A0(*(uint32_t*)(&a1_->field_75) + *(uint32_t*)(*(uint32_t*)(&a1_->field_9) + 48));
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

int sub_4523D0(struct576* a1p_) {
	void* a1p = a1p_;
	uint32_t* a1 = a1p;
	int result = 0; // eax

	if (!(a1[6] & 1)) {
		sub_452410((int)a1);
		sub_451F90((int)a1);
		a1[7] = 4;
		a1[70] = 0;
		result = a1[6];
		LOBYTE(result) = result | 1;
		a1[6] = result;
	}
	return result;
}

int sub_452410(struct576* a1_) {
	int a1 = a1_;
	int result; // eax

	result = *(uint32_t*)(&a1_->field_44);
	if (result && a1_ == *(uint32_t*)(result + 152)) {
		if (*(uint8_t*)(&a1_->field_6) & 2) {
			sub_4BDA80(*(uint32_t*)(&a1_->field_44));
		}
		sub_4BDB30(*(uint32_t*)(&a1_->field_44));
		*(uint32_t*)(*(uint32_t*)(&a1_->field_44) + 152) = 0;
		*(uint32_t*)(*(uint32_t*)(&a1_->field_44) + 148) = 0;
		result = *(uint32_t*)(&a1_->field_44);
		*(uint32_t*)(result + 140) = 0;
		*(uint32_t*)(*(uint32_t*)(&a1_->field_44) + 144) = 0;
		*(uint32_t*)(*(uint32_t*)(&a1_->field_44) + 112) = 0;
		*(uint32_t*)(&a1_->field_44) = 0;
	}
	return result;
}

int sub_452490(struct576* a1_) {
	uint32_t* a1 = a1_;
	int v1; // eax
	int v3; // edi
	int v4; // eax

	v1 = a1_->field_44;
	if (a1_ != *(uint32_t**)(v1 + 152)) {
		return 0;
	}
	v3 = a1_->field_74;
	sub_4BDB90((uint32_t*)v1, (uint32_t*)a1_->field_74);
	a1_->field_7 = 3;
	v4 = a1_->field_6;
	LOBYTE(v4) = v4 | 2;
	a1_->field_6 = v4;
	a1_->field_74 = 0;
	if (!sub_4BDB40(a1_->field_44)) {
		return 1;
	}
	a1_->field_7 = 1;
	a1_->field_74 = v3;
	a1_->field_6 &= 0xFFFFFFFD;
	return 0;
}

void sub_452510(struct576* a3_) {
	int a3 = a3_;
	int v1; // eax
	int v2; // eax

	if (!dword_587000_126996) {
		*(uint32_t*)(&a3_->field_7) = 4;
	}
	while (1) {
		v1 = *(uint32_t*)(&a3_->field_7);
		if (!v1) {
			break;
		}
		v2 = v1 - 2;
		if (v2) {
			uint32_t v3 = v2 - 2;
			if (!v3) {
				sub_4523D0(a3_);
			}
			return;
		}
		if (nox_platform_get_ticks() <= *(uint64_t*)(&a3_->field_72)) {
			return;
		}
		*(uint32_t*)(&a3_->field_7) = *(uint32_t*)(&a3_->field_8);
	}
	if (!sub_452580(a3_)) {
		sub_4523D0(a3_);
	}
}

long long sub_452690(struct576* a3_, long long a4, int a5) {
	int a3 = a3_;
	long long result; // rax

	*(uint32_t*)(&a3_->field_8) = a5;
	result = a4 + nox_platform_get_ticks();
	*(uint64_t*)(&a3_->field_72) = result;
	*(uint32_t*)(&a3_->field_7) = 2;
	return result;
}

uint32_t* nox_xxx_draw_452300(struct200* a1_) {
	if (!dword_5d4594_1045432) {
		return 0;
	}
	if (!dword_587000_126996) {
		return 0;
	}
	if (!a1_->field_0) {
		return 0;
	}
	struct576* v1p = sub_4BD2E0(*(uint32_t***)&dword_5d4594_1045436);
	uint32_t* v1 = v1p;
	if (!v1) {
		sub_452230();
		v1 = sub_4BD2E0(*(uint32_t***)&dword_5d4594_1045436);
		if (!v1) {
			return 0;
		}
	}
	memset(v1p, 0, 0x240u);
	v1p->field_9 = a1_;
	sub_425770(v1p);
	v1p->field_7 = 0;
	v1p->field_75 = 0;
	v1p->field_142 = 0;
	v1p->field_108 = 0;
	v1p->field_42 = 0;
	sub_4864A0(&v1p->timerGroup_46);
	nox_common_list_append_4258E0((int)getMemAt(0x5D4594, 840612), v1p);
	v1p->field_70 = (*getMemU32Ptr(0x587000, 127000))++;
	return v1p;
}

int sub_452E90(uint32_t* a1, struct576* a2_) {
	int a2 = a2_;
	int result; // eax

	result = a2_;
	*a1 = a2_;
	if (a2_) {
		a1[1] = *(uint32_t*)(&a2_->field_70);
		result = *(uint32_t*)(&a2_->field_9);
		a1[2] = result;
	}
	return result;
}

int sub_452EE0(struct576* a1_, int a2) {
	int a1 = a1_;
	int v2; // eax

	v2 = sub_452F10(a1_, a2);
	sub_486320((uint32_t*)(&a1_->timerGroup_46), v2);
	return sub_4863B0((unsigned int*)(&a1_->timerGroup_46));
}

unsigned int sub_452F10(struct576* a1_, int a2) {
	int a1 = a1_;
	int v2; // ecx

	v2 = a2;
	if (a2 <= 100) {
		if (a2 < 0) {
			v2 = 0;
		}
	} else {
		v2 = 100;
	}
	return (unsigned int)(163 * v2 * (*(uint32_t*)(*(uint32_t*)(&a1_->field_9) + 20) >> 16)) >> 14;
}

int sub_452F50(struct576* a1_, int a2) {
	int a1 = a1_;
	int v2; // eax

	v2 = sub_452F10(a1_, a2);
	return sub_486350(&a1_->timerGroup_46, v2);
}

uint32_t* sub_452F80(struct576* a1_, int a2) {
	int a1 = a1_;
	int v2; // eax

	v2 = sub_452FA0(a2);
	return sub_486320((uint32_t*)(&a1_->timerGroup_46.field_16), v2);
}

int sub_451CF0(uint32_t* a1) {
	int v1;       // ecx
	int result;   // eax
	int v3;       // edx
	int v4;       // edi
	int v5;       // eax
	int v6;       // edi
	int v7;       // ecx
	uint32_t* v8; // eax
	int v9;       // eax

	v1 = a1[9];
	result = a1[108];
	v3 = *(uint32_t*)(v1 + 4);
	if (result) {
		if (v3 & 2) {
			v5 = nox_common_randomIntMinMax_415FF0(0, result - 1, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 376);
			v6 = a1[108] - 1;
			a1[43] = a1[v5 + 76];
			v7 = v5;
			if (v5 < v6) {
				v8 = &a1[v5 + 76];
				do {
					++v7;
					*v8 = v8[1];
					++v8;
				} while (v7 < a1[108] - 1);
			}
		} else {
			++a1[43];
		}
		v9 = a1[43];
		--a1[108];
		result = sub_4BD710(a1[v9 + 10]);
	} else if (v3 & 1) {
		if (*(uint32_t*)(v1 + 60) && (v4 = a1[109] + 1, a1[109] = v4, v4 >= *(int*)(v1 + 60))) {
			result = 0;
		} else {
			result = sub_451CA0(a1);
		}
	}
	return result;
}

int sub_451DC0(int a1) {
	uint32_t* v1; // esi
	int result;   // eax
	int v3;       // ebx
	int i;        // edi
	int v5;       // eax
	int v6;       // eax

	v1 = *(uint32_t**)(a1 + 36);
	result = *(uint32_t*)(a1 + 168);
	v3 = v1[1];
	if (result) {
		if (v1[17] < 0x21u) {
			return result;
		}
		sub_451F90(a1);
	}
	if (v3 & 4) {
		if (v1[17] >= 0x21u) {
			v5 = sub_451E80(a1);
			result = sub_451F30(a1, v5);
		} else {
			result = v1[48];
			for (i = 0; i < result; ++i) {
				sub_451F30(a1, i);
				result = v1[48];
			}
		}
	} else if (v3 & 2) {
		v6 = nox_common_randomIntMinMax_415FF0(0, v1[48] - 1, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 536);
		result = sub_451F30(a1, v6);
	} else {
		result = sub_451F30(a1, 0);
	}
	return result;
}

int sub_452580(struct576* a1_) {
	uint32_t* a1 = a1_;
	int v1;     // edi
	int result; // eax
	int v3;     // eax
	int v4;     // eax
	int v5;     // eax

	v1 = a1_->field_9;
	if (!*(uint32_t*)(v1 + 192)) {
		return 0;
	}
	v3 = a1_->field_75;
	a1_->field_109 = 0;
	result = sub_452810(*(uint32_t*)(v1 + 48) + v3, 0);
	a1_->field_44 = result;
	if (result) {
		v4 = nox_common_randomIntMinMax_415FF0(*(uint32_t*)(v1 + 76), *(uint32_t*)(v1 + 80),
											   "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 1482);
		sub_486320((uint32_t*)(a1_->field_44 + 48), v4 + 100);
		sub_4BDB20(a1_->field_44);
		*(uint32_t*)(a1_->field_44 + 152) = a1_;
		*(uint32_t*)(a1_->field_44 + 140) = sub_452770;
		*(uint32_t*)(a1_->field_44 + 144) = sub_4526F0;
		*(uint32_t*)(a1_->field_44 + 148) = sub_4526D0;
		a1_->field_7 = 1;
		*(uint32_t*)(a1_->field_44 + 112) = &a1_->timerGroup_46;
		if (*(uint8_t*)(v1 + 4) & 8) {
			v5 = nox_common_randomIntMinMax_415FF0(*(uint32_t*)(v1 + 68), *(uint32_t*)(v1 + 72),
												   "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 1497);
			if (v5 > 33) {
				sub_452690(a1_, v5, 1);
			}
		}
		result = 1;
	}
	return result;
}

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

void sub_452050(struct576* a1_) {
	uint32_t* a1 = a1_;
	uint32_t* v1;      // esi
	int v2;            // edi
	unsigned int v3;   // ebx
	unsigned char* v4; // ebp
	uint32_t* result;  // eax
	uint32_t** v6;     // esi
	uint32_t** v7;     // esi
	uint32_t* v8;      // esi

	v1 = (uint32_t*)a1_->field_9;
	v2 = v1[12] + a1_->field_75;
	v3 = (a1_->timerGroup_46.field_0.field_1 >> 16) / 0x666u;
	v4 = getMemAt(0x5D4594, 839892 + 120 * v2);
	if (v1[26] == *getMemU32Ptr(0x5D4594, 1045444)) {
		result = (uint32_t*)v1[27];
		if (v2 <= (int)result) {
			if ((uint32_t*)v2 == result && v3 > v1[31]) {
				v1[31] = v3;
				v7 = (uint32_t**)(v1 + 28);
				nox_common_list_remove_425920(v7);
				nox_common_list_append_4258E0((int)&v4[12 * v3], v7);
			}
		} else {
			v1[27] = v2;
			v1[31] = v3;
			v6 = (uint32_t**)(v1 + 28);
			nox_common_list_remove_425920(v6);
			nox_common_list_append_4258E0((int)&v4[12 * v3], v6);
		}
	} else {
		v1[26] = *getMemU32Ptr(0x5D4594, 1045444);
		v1[27] = v2;
		v1[31] = v3;
		v8 = v1 + 28;
		sub_425770(v8);
		nox_common_list_append_4258E0((int)&v4[12 * v3], v8);
	}
}

int sub_451BE0(struct576* a1_) {
	int a1 = a1_;
	int v1;          // eax
	int v2;          // edi
	unsigned int v3; // ebx
	uint32_t* v4;    // esi
	int v5;          // eax
	int v6;          // eax
	uint32_t* v7;    // ebx
	int result;      // eax
	int v9;          // esi
	uint32_t* v10;   // esi

	v1 = a1_;
	v2 = *(uint32_t*)(&a1_->field_9);
	v3 = *(uint32_t*)(&a1_->timerGroup_46.field_0.field_1) >> 16;
	v4 = *(uint32_t**)(v2 + 88);
	if (v4 != (uint32_t*)(v2 + 88)) {
		do {
			v5 = (v4[44] >> 16) - v3;
			if (v5 < 0) {
				v5 = v3 - (v4[44] >> 16);
			}
			if (v5 >= (*(uint32_t*)(v2 + 20) >> 16) / 10) {
				if (v4[44] >> 16 < v3) {
					break;
				}
			} else {
				v6 = v4[4];
				if (*(uint8_t*)(v2 + 4) & 0x10) {
					if (v6) {
						break;
					}
				} else if (!v6) {
					break;
				}
			}
			v4 = (uint32_t*)*v4;
		} while (v4 != (uint32_t*)(v2 + 88));
		v1 = a1_;
	}
	v7 = (uint32_t*)(v1 + 12);
	sub_425770((uint32_t*)(v1 + 12));
	nox_common_list_append_4258E0((int)v4, v7);
	result = *(uint32_t*)(v2 + 56);
	v9 = *(uint32_t*)(v2 + 52) + 1;
	*(uint32_t*)(v2 + 52) = v9;
	if (result) {
		if (v9 > result) {
			v10 = (uint32_t*)(*(uint32_t*)(v2 + 92) - 12);
			nox_common_list_remove_425920(*(uint32_t***)(v2 + 92));
			sub_4523D0(v10);
			result = *(uint32_t*)(v2 + 52) - 1;
			*(uint32_t*)(v2 + 52) = result;
		}
	}
	return result;
}

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

int sub_4BDB20(int a1) {
	int result; // eax

	result = a1;
	*(uint32_t*)(a1 + 124) |= 0x10u;
	return result;
}

int sub_4BD710(int a1) { return a1 + 24; }

int sub_4526D0(int a1) {
	*(uint32_t*)(*(uint32_t*)(a1 + 152) + 28) = 4;
	return 0;
}

int sub_452FE0(struct576* a1_, int a2) {
	int a1 = a1_;
	int v2; // eax

	v2 = sub_452FA0(a2);
	return sub_486350(&a1_->timerGroup_46.field_16, v2);
}

int sub_452FA0(int a1) {
	int v1; // eax

	v1 = a1;
	if (a1 <= 50) {
		if (a1 < -50) {
			v1 = -50;
		}
	} else {
		v1 = 50;
	}
	return (v1 * 8192) / 50 + 8192;
}

int sub_4BD650(int a1) {
	int result; // eax

	result = a1;
	++*(uint32_t*)(a1 + 12);
	return result;
}

int sub_4BD660(int a1) {
	int result; // eax

	result = *(uint32_t*)(a1 + 12) - 1;
	*(uint32_t*)(a1 + 12) = result;
	if (result < 0) {
		*(uint32_t*)(a1 + 12) = 0;
	}
	return result;
}

void nox_xxx_clientPlaySoundSpecial_452D80(int a1, int a2) {
	uint32_t* result; // eax
	uint32_t* v3;     // esi

	result = nox_xxx_draw_452270(a1);
	if (!result) {
		return;
	}
	result = nox_xxx_draw_452300(result);
	v3 = result;
	if (!result) {
		return;
	}
	sub_452EE0((int)result, a2);
	sub_452510((int)v3);
}

void sub_452DC0(int a1, int a2, int a3) {
	uint32_t* result; // eax
	uint32_t* v4;     // esi

	result = nox_xxx_draw_452270(a1);
	if (!result) {
		return;
	}
	result = nox_xxx_draw_452300(result);
	v4 = result;
	if (!result) {
		return;
	}
	sub_452EE0((int)result, a2);
	sub_452F80((int)v4, a3);
	sub_452510((int)v4);
}

void sub_452E10(int a1, int a2, int a3) {
	uint32_t* result; // eax
	uint32_t* v4;     // esi

	result = nox_xxx_draw_452270(a1);
	if (!result) {
		return;
	}
	result = nox_xxx_draw_452300(result);
	v4 = result;
	if (!result) {
		return;
	}
	sub_452EE0((int)result, a2);
	sub_452F80((int)v4, a3);
	v4[75] = 2;
	sub_452510((int)v4);
}
