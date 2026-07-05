#include <stdlib.h>
#include <string.h>

#include "client__gui__window.h"
#include "common/alloc/classes/alloc_class.h"

#include "defs.h" // nox_video_bag_section_t, nox_video_bag_image_t
#include "memmap.h"

// TODO: part of the GUI toolkit
void nox_client_drawImageAt_47D2C0(void* a1, int x, int y);
int nox_xxx_drawGetStringSize_43F840(int a1, unsigned short* a2, int* a3, int* a4, int a5);
void nox_client_drawSetColor_434460(int a1);
void nox_client_drawRectFilledOpaque_49CE30(int xLeft, int yTop, int a3, int a4);
nox_video_bag_image_t* nox_xxx_gLoadImg_42F970(const char* a1);

//----- (0046ACE0) --------------------------------------------------------
void sub_46ACE0(unsigned int* a1, int a2, int a3, int a4) {
	for (int i = a2; i <= a3; i++) {
		unsigned int* v5 = nox_xxx_wndGetChildByID_46B0C0(a1, i);
		nox_window_set_hidden((int)v5, a4);
	}
}

//----- (0046AD20) --------------------------------------------------------
void sub_46AD20(unsigned int* a1, int a2, int a3, int a4) {
	int i;            // esi
	unsigned int* v5; // eax

	for (i = a2; i <= a3; ++i) {
		v5 = nox_xxx_wndGetChildByID_46B0C0(a1, i);
		nox_xxx_wnd_46ABB0((int)v5, a4);
	}
}

//----- (0046AB20) --------------------------------------------------------
int sub_46AB20(unsigned int* a1, int a2, int a3) {
	int v4; // esi

	if (!a1) {
		return -2;
	}
	a1[6] = a2 + a1[4];
	v4 = a3 + a1[5];
	a1[2] = a2;
	a1[3] = a3;
	a1[7] = v4;
	nox_window_call_field_94((int)a1, 16388, a2, a3);
	return 0;
}

//----- (0046ABB0) --------------------------------------------------------
int nox_xxx_wnd_46ABB0(nox_window* win, int a2) {
	int a1 = win;
	int v3;          // ecx
	unsigned int v4; // ecx
	int v5;          // esi

	if (!a1) {
		return -2;
	}
	v3 = *(unsigned int*)(a1 + 4);
	if (a2) {
		v4 = v3 | 8;
	} else {
		v4 = v3 & 0xFFFFFFF7;
	}
	v5 = *(unsigned int*)(a1 + 400);
	for (*(unsigned int*)(a1 + 4) = v4; v5; v5 = *(unsigned int*)(v5 + 388)) {
		nox_xxx_wnd_46ABB0(v5, a2);
	}
	return 0;
}

//----- (0046ADA0) --------------------------------------------------------
int nox_xxx_wndGetFlags_46ADA0(int a1) {
	int result; // eax

	if (a1) {
		result = *(unsigned int*)(a1 + 4);
	} else {
		result = -2;
	}
	return result;
}

//----- (0046B280) --------------------------------------------------------
int nox_xxx_wnd_46B280(int a1, int a2) {
	if (!a1) {
		return -2;
	}
	if (a2) {
		*(unsigned int*)(a1 + 52) = a2;
	} else {
		*(unsigned int*)(a1 + 52) = a1;
	}
	return 0;
}
