#include "audio.h"
#include "memmap.h"

// External variable declarations
extern struct28* dword_5d4594_1045424;
// It holds memory allocation holding 200 elements of struct576. The first 4 bytes points to next free element (or 0 if
// no remaining free space), and for each elements, first 4 bytes contain pointer to the next free element (or 0 if it's
// the last element).
extern uint32_t* dword_5d4594_1045436;
extern uint32_t dword_587000_126996;
extern uint32_t dword_5d4594_1045432;

// Common library using list access
void nox_common_list_append_4258E0(nox_list_item_t* list, nox_list_item_t* cur);
void nox_common_list_remove_425920(void* a1);
void* sub_425770(void* a1);

int* sub_452810(int a1, char a2);
int sub_4526D0(int a1);
int sub_4526F0(int a1);
int sub_4BDB20(int a1);
int sub_4BD710(int a1);
int sub_452FA0(int a1);
int sub_486350(/* timer* */ void* a1, int a2);
int sub_4863B0(/* timer* */ void* a2);
void* sub_486320(/* timer* */ void* a1, int a2);
void sub_4864A0(timerGroup* a3);
int nox_common_randomIntMinMax_415FF0(int min, int max, const char* file, int line);
int***** sub_452230();
int sub_4BDB40(int a2);
void sub_4BDB90(uint32_t* a1, uint32_t* a2);
int sub_4BDA80(int a1);
int sub_4BDB30(int a1);
void sub_452190(int a1);
struct200* sub_4521A0(int a1);
int sub_4BD300(uint32_t* a1, int a2);
int sub_4BD660(int a1);
int sub_4BD650(int a1);
uint32_t* sub_4BD470(uint32_t** a1, int a2);

//----- (00451CA0) --------------------------------------------------------
int sub_451CA0(struct576* a1p) {
	int v1;       // ecx
	int v3;       // eax
	uint32_t* v4; // ecx

	v1 = a1p->field_42;
	a1p->field_108 = v1;
	if (!v1) {
		return 0;
	}
	v3 = 0;
	if (v1 > 0) {
		v4 = a1p->field_76;
		do {
			*v4 = v3++;
			++v4;
		} while (v3 < a1p->field_108);
	}
	a1p->field_43 = -1;
	return sub_451CF0(a1p);
}

//----- (00451F30) --------------------------------------------------------
int sub_451F30(struct576* a1p, int a2) {
	int v2;     // edx
	int result; // eax

	a1p->field_10[a1p->field_42] = sub_4BD470((uint32_t**)&dword_5d4594_1045424->field_0, a1p->field_9->field_32[a2]);
	v2 = a1p->field_42;
	result = a1p->field_10[v2];
	if (result) {
		sub_4BD650(a1p->field_10[v2]);
		result = a1p->field_42 + 1;
		a1p->field_42 = result;
	}
	return result;
}

//----- (00451F90) --------------------------------------------------------
int sub_451F90(struct576* a1p) {
	int v1;     // edi
	int result; // eax
	int* v3;    // esi

	v1 = 0;
	result = a1p->field_42;
	if (result <= 0) {
		a1p->field_42 = 0;
	} else {
		v3 = a1p->field_10;
		do {
			sub_4BD660(v3[0]);
			v3[0] = 0;
			result = a1p->field_42;
			++v1;
			++v3;
		} while (v1 < result);
		a1p->field_42 = 0;
	}
	return result;
}

//----- (00451FE0) --------------------------------------------------------
int sub_451FE0(struct576* a1p) {
	nox_common_list_remove_425920(a1p);
	a1p->field_70 = 0;
	return sub_4BD300(*(uint32_t**)&dword_5d4594_1045436, a1p);
}

//----- (00452120) --------------------------------------------------------
bool sub_452120(struct576* a1p) {
	// int* result; // eax
	// int* v3; // ebx

	bool v1 = false;
	struct200* res;
	struct200* v3p;
	res = sub_4521A0(a1p->field_75 + a1p->field_9->field_12);
	v3p = res;
	if (res == 0) {
		return false;
	}
	sub_452190(res);
	struct576* v4p;
	struct576* v5p;
	v4p = get_list_at_840612()->field_0;
	if (get_list_at_840612()->field_0 != get_list_at_840612()) {
		do {
			v5p = v4p->next;
			if (v4p->field_9 == v3p) {
				sub_4523D0(v4p);
				sub_451FE0(v4p);
				v1 = true;
			}
			v4p = v5p;
		} while (v5p != get_list_at_840612());
	}
	return v1;
}

//----- (004523D0) --------------------------------------------------------
int sub_4523D0(struct576* a1p) {
	int result = 0; // eax

	if (!(a1p->field_6 & 1)) {
		sub_452410(a1p);
		sub_451F90(a1p);
		a1p->field_7 = 4;
		a1p->field_70 = 0;
		a1p->field_6 |= 1;
	}
	return result;
}

//----- (00452410) --------------------------------------------------------
int sub_452410(struct576* a1p) {
	int result; // eax

	result = a1p->field_44;
	if (result && a1p == *(uint32_t*)(result + 152)) {
		if (a1p->field_6 & 2) {
			sub_4BDA80(a1p->field_44);
		}
		sub_4BDB30(a1p->field_44);
		*(uint32_t*)(a1p->field_44 + 152) = 0;
		*(uint32_t*)(a1p->field_44 + 148) = 0;
		result = a1p->field_44;
		*(uint32_t*)(result + 140) = 0;
		*(uint32_t*)(a1p->field_44 + 144) = 0;
		*(uint32_t*)(a1p->field_44 + 112) = 0;
		a1p->field_44 = 0;
	}
	return result;
}

//----- (00452490) --------------------------------------------------------
int sub_452490(struct576* a1p) {
	int v1; // eax
	int v3; // edi

	v1 = a1p->field_44;
	if (a1p != *(uint32_t**)(v1 + 152)) {
		return 0;
	}
	v3 = a1p->field_74;
	sub_4BDB90((uint32_t*)v1, (uint32_t*)a1p->field_74);
	a1p->field_7 = 3;
	a1p->field_6 |= 2;
	a1p->field_74 = 0;
	if (!sub_4BDB40(a1p->field_44)) {
		return 1;
	}
	a1p->field_7 = 1;
	a1p->field_74 = v3;
	a1p->field_6 &= 0xFFFFFFFD;
	return 0;
}

//----- (00452510) --------------------------------------------------------
void sub_452510(struct576* a3) {
	int v1; // eax
	int v2; // eax

	if (!dword_587000_126996) {
		a3->field_7 = 4;
	}
	while (1) {
		v1 = a3->field_7;
		if (!v1) {
			break;
		}
		v2 = v1 - 2;
		if (v2) {
			uint32_t v3 = v2 - 2;
			if (!v3) {
				sub_4523D0((uint32_t*)a3);
			}
			return;
		}
		if (nox_platform_get_ticks() <= a3->field_72) {
			return;
		}
		a3->field_7 = a3->field_8;
	}
	if (!sub_452580(a3)) {
		sub_4523D0(a3);
	}
}

//----- (00452690) --------------------------------------------------------
long long sub_452690(struct576* a3p, long long a4, int a5) {
	long long result; // rax

	a3p->field_8 = a5;
	result = a4 + nox_platform_get_ticks();
	a3p->field_72 = result;
	a3p->field_7 = 2;
	return result;
}

//----- (00452300) --------------------------------------------------------
struct576* nox_xxx_draw_452300(struct200* a1p) {
	// uint32_t* a1 = a1p;
	struct576* v1; // esi

	if (!dword_5d4594_1045432) {
		return 0;
	}
	if (!dword_587000_126996) {
		return 0;
	}
	if (!a1p->field_0) {
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
	memset(v1, 0, sizeof(struct576));
	v1->field_9 = a1p;
	sub_425770(v1);
	v1->field_7 = 0;
	v1->field_75 = 0;
	v1->field_142 = 0;
	v1->field_108 = 0;
	v1->field_42 = 0;
	sub_4864A0(&v1->timerGroup_46.field_0);
	nox_common_list_append_4258E0(get_list_at_840612(), v1);
	v1->field_70 = (*getMemU32Ptr(0x587000, 127000))++;
	return v1;
}

//----- (00452E90) --------------------------------------------------------
void sub_452E90(nox_drawable_inner3* a1, struct576* a2) {
	a1->field_0 = a2;
	if (a2) {
		a1->field_1 = a2->field_70;
		a1->field_2 = a2->field_9;
	}
}

//----- (00452EB0) --------------------------------------------------------
struct576* sub_452EB0(nox_drawable_inner3* a1) {
	struct576* res;

	res = a1->field_0;
	if (a1->field_0 && (a1->field_2 != res->field_9 || a1->field_1 != res->field_70)) {
		res = 0;
		a1->field_0 = 0;
	}
	return res;
}

//----- (00452EE0) --------------------------------------------------------
int sub_452EE0(struct576* a1, int a2) {
	int v2; // eax

	v2 = sub_452F10(a1, a2);
	sub_486320(&a1->timerGroup_46.field_0, v2);
	return sub_4863B0(&a1->timerGroup_46.field_0);
}

//----- (00452F10) --------------------------------------------------------
unsigned int sub_452F10(struct576* a1p, int a2) {
	int v2; // ecx

	v2 = a2;
	if (a2 <= 100) {
		if (a2 < 0) {
			v2 = 0;
		}
	} else {
		v2 = 100;
	}
	return (unsigned int)(163 * v2 * (a1p->field_9->field_4.field_1 >> 16)) >> 14;
}

//----- (00452F50) --------------------------------------------------------
int sub_452F50(struct576* a1p, int a2) {
	int v2; // eax

	v2 = sub_452F10(a1p, a2);
	return sub_486350(&a1p->timerGroup_46.field_0, v2);
}

//----- (00452F80) --------------------------------------------------------
uint32_t* sub_452F80(struct576* a1, int a2) {
	int v2; // eax

	v2 = sub_452FA0(a2);
	return sub_486320(&a1->timerGroup_46.field_16, v2);
}

//----- (00451CF0) --------------------------------------------------------
int sub_451CF0(struct576* a1p) {
	int v1;       // ecx
	int result;   // eax
	int v3;       // edx
	int v4;       // edi
	int v5;       // eax
	int v6;       // edi
	int v7;       // ecx
	uint32_t* v8; // eax
	int v9;       // eax

	struct200* v1p;
	v1 = v1p = a1p->field_9;
	result = a1p->field_108;
	v3 = *(uint32_t*)(v1 + 4);
	if (result) {
		if (v3 & 2) {
			v5 = nox_common_randomIntMinMax_415FF0(0, result - 1, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 376);
			v6 = a1p->field_108 - 1;
			a1p->field_43 = a1p->field_76[v5];
			v7 = v5;
			if (v5 < v6) {
				v8 = &a1p->field_76[v5];
				do {
					++v7;
					*v8 = v8[1];
					++v8;
				} while (v7 < a1p->field_108 - 1);
			}
		} else {
			++a1p->field_43;
		}
		v9 = a1p->field_43;
		--a1p->field_108;
		result = sub_4BD710(a1p->field_10[v9]);
	} else if (v3 & 1) {
		if (*(uint32_t*)(v1 + 60) && (v4 = a1p->field_109 + 1, a1p->field_109 = v4, v4 >= *(int*)(v1 + 60))) {
			result = 0;
		} else {
			result = sub_451CA0(a1p);
		}
	}
	return result;
}

//----- (00451DC0) --------------------------------------------------------
int sub_451DC0(struct576* a1p) {
	uint32_t* v1; // esi
	int result;   // eax
	int v3;       // ebx
	int i;        // edi
	int v5;       // eax
	int v6;       // eax

	struct200* v1p;
	v1 = v1p = a1p->field_9;
	result = a1p->field_42;
	v3 = v1[1];
	if (result) {
		if (v1[17] < 0x21u) {
			return result;
		}
		sub_451F90(a1p);
	}
	if (v3 & 4) {
		if (v1[17] >= 0x21u) {
			v5 = sub_451E80(a1p);
			result = sub_451F30(a1p, v5);
		} else {
			result = v1[48];
			for (i = 0; i < result; ++i) {
				sub_451F30(a1p, i);
				result = v1[48];
			}
		}
	} else if (v3 & 2) {
		v6 = nox_common_randomIntMinMax_415FF0(0, v1[48] - 1, "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 536);
		result = sub_451F30(a1p, v6);
	} else {
		result = sub_451F30(a1p, 0);
	}
	return result;
}

//----- (00452580) --------------------------------------------------------
int sub_452580(struct576* a1p) {
	int v1;     // edi
	int result; // eax
	int v3;     // eax
	int v4;     // eax
	int v5;     // eax

	struct200* v1p;
	v1 = v1p = a1p->field_9;
	if (!*(uint32_t*)(v1 + 192)) {
		return 0;
	}
	v3 = a1p->field_75;
	a1p->field_109 = 0;
	result = sub_452810(*(uint32_t*)(v1 + 48) + v3, 0);
	a1p->field_44 = result;
	if (result) {
		v4 = nox_common_randomIntMinMax_415FF0(*(uint32_t*)(v1 + 76), *(uint32_t*)(v1 + 80),
											   "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 1482);
		sub_486320((uint32_t*)(a1p->field_44 + 48), v4 + 100);
		sub_4BDB20(a1p->field_44);
		*(uint32_t*)(a1p->field_44 + 152) = (uint32_t)a1p;
		*(uint32_t*)(a1p->field_44 + 140) = (uint32_t)sub_452770;
		*(uint32_t*)(a1p->field_44 + 144) = (uint32_t)sub_4526F0;
		*(uint32_t*)(a1p->field_44 + 148) = (uint32_t)sub_4526D0;
		a1p->field_7 = 1;
		*(uint32_t*)(a1p->field_44 + 112) =
			(uint32_t)&a1p->timerGroup_46; // Not sure it's timerGroup_46 or timerGroup_46.field_0. Their pointer
										   // addresses are the same.
		if (*(uint8_t*)(v1 + 4) & 8) {
			v5 = nox_common_randomIntMinMax_415FF0(*(uint32_t*)(v1 + 68), *(uint32_t*)(v1 + 72),
												   "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 1497);
			if (v5 > 33) {
				sub_452690((int)a1p, v5, 1);
			}
		}
		result = 1;
	}
	return result;
}

//----- (004BD2E0) --------------------------------------------------------
uint32_t* sub_4BD2E0(uint32_t** a1) {
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

//----- (00451E80) --------------------------------------------------------
int sub_451E80(struct576* a1p) {
	int a1 = a1p;
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
	struct576* v1p;
	uint32_t* v2;    // ebx
	int v4;          // eax
	unsigned int v5; // eax

	v1p = (uint32_t*)a1[38];
	v2 = (uint32_t*)sub_451CF0((uint32_t*)a1[38]);
	if (v1p->field_9->field_18 < 0x21u) {
		sub_4BDB90(a1, v2);
		return 0;
	}
	sub_4BDB90(a1, 0);
	struct200* v4p;
	v4 = v4p = v1p->field_9;
	if (!(*(uint8_t*)(v4 + 4) & 8) || v2 || v1p->field_142) {
		v5 = nox_common_randomIntMinMax_415FF0(*(uint32_t*)(v4 + 68), *(uint32_t*)(v4 + 72),
											   "C:\\NoxPost\\src\\client\\Audio\\AudEvent.c", 706);
		if (v5 < 0x21) {
			sub_4BDB90(a1, v2);
			return 0;
		}
		v1p->field_71 = v5;
		v1p->field_74 = v2;
	}
	return 0;
}

//----- (00452050) --------------------------------------------------------
void sub_452050(struct576* a1p) {
	// uint32_t* v1;      // esi
	int v2;               // edi
	unsigned int v3;      // ebx
	nox_list_item_t* v4p; // ebp
	uint32_t* result;     // eax
	uint32_t** v6;        // esi
	uint32_t** v7;        // esi
	uint32_t* v8;         // esi

	struct200* v1p;
	v1p = a1p->field_9;
	v2 = v1p->field_12 + a1p->field_75;
	v3 = (a1p->timerGroup_46.field_0.field_1 >> 16) / 0x666u;
	nox_list_item_t(*v4s)[10] = getMemAt(0x5D4594, 839892);
	v4p = &v4s[v2];
	if (v1p->field_26 == *getMemU32Ptr(0x5D4594, 1045444)) {
		result = v1p->field_27;
		if (v2 <= (int)result) {
			if ((uint32_t*)v2 == result && v3 > v1p->field_31) {
				v1p->field_31 = v3;
				v7 = &v1p->field_28;
				nox_common_list_remove_425920(v7);
				nox_common_list_append_4258E0(&v4p[v3], v7);
			}
		} else {
			v1p->field_27 = v2;
			v1p->field_31 = v3;
			v6 = &v1p->field_28;
			nox_common_list_remove_425920(v6);
			nox_common_list_append_4258E0(&v4p[v3], v6);
		}
	} else {
		v1p->field_26 = *getMemU32Ptr(0x5D4594, 1045444);
		v1p->field_27 = v2;
		v1p->field_31 = v3;
		v8 = &v1p->field_28;
		sub_425770(v8);
		nox_common_list_append_4258E0(&v4p[v3], v8);
	}
}

//----- (00451BE0) --------------------------------------------------------
int sub_451BE0(struct576* a1p) {
	// int v2;          // edi
	unsigned int v3; // ebx
	uint32_t* v4;    // esi
	int v5;          // eax
	int v6;          // eax
	uint32_t* v7;    // ebx
	int result;      // eax
	int v9;          // esi
	uint32_t* v10;   // esi

	struct200* v2p;
	v2p = a1p->field_9;
	v3 = (uint32_t)(a1p->timerGroup_46.field_0.field_1) >> 16;
	v4 = v2p->field_22.field_0;
	if (v4 != v2p->field_22.field_0) {
		do {
			v5 = (v4[44] >> 16) - v3;
			if (v5 < 0) {
				v5 = v3 - (v4[44] >> 16);
			}
			if (v5 >= (v2p->field_4.field_1 >> 16) / 10) {
				if (v4[44] >> 16 < v3) {
					break;
				}
			} else {
				v6 = v4[4];
				if (v2p->field_1 & 0x10) {
					if (v6) {
						break;
					}
				} else if (!v6) {
					break;
				}
			}
			v4 = (uint32_t*)*v4;
		} while (v4 != &v2p->field_22);
	}
	v7 = (uint32_t*)&a1p->field_3;
	sub_425770((uint32_t*)&a1p->field_3);
	nox_common_list_append_4258E0((int)v4, v7);
	result = v2p->field_14;
	v9 = v2p->field_13 + 1;
	v2p->field_13 = v9;
	if (result) {
		if (v9 > result) {
			v10 = (uint32_t*)((uint32_t)v2p->field_22.field_1 - 12);
			nox_common_list_remove_425920(v2p->field_22.field_1);
			sub_4523D0(v10);
			result = v2p->field_13 - 1;
			v2p->field_13 = result;
		}
	}
	return result;
}