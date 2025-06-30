#include "temp_subs.h"

void sub_43DD10(void* a1);
void sub_43D9E0(void* a1);

extern uint32_t dword_5d4594_816368;
extern uint32_t dword_5d4594_816372;

// Current understanding: int4[3][6]
#define MUSIC_STATE_ARRAY ((int4(*)[6])getMemAt(0x5D4594, 815772))

//----- (0043DA80) --------------------------------------------------------
int sub_43DA80() {
	int result; // eax

	if (*(int*)&dword_5d4594_816368 < 6) {
		sub_43DD10(&MUSIC_STATE_ARRAY[dword_5d4594_816372][dword_5d4594_816368]);
		++dword_5d4594_816368;
		result = 1;
	} else {
		dword_5d4594_816368 = 6;
		result = 0;
	}
	return result;
}

//----- (0043DAD0) --------------------------------------------------------
void sub_43DAD0() {
	if (dword_5d4594_816368 > 0) {
		sub_43D9E0(&MUSIC_STATE_ARRAY[dword_5d4594_816372][--dword_5d4594_816368]);
	}
	dword_5d4594_816368 = 0;
}

//----- (0043DB20) --------------------------------------------------------
int sub_43DB20() { return dword_5d4594_816368; }

//----- (0043DB30) --------------------------------------------------------
int sub_43DB30(int a1) {
	int result; // eax

	result = a1;
	dword_5d4594_816368 = a1;
	return result;
}

//----- (0043DB40) --------------------------------------------------------
int4* sub_43DB40(int a1) { return &MUSIC_STATE_ARRAY[dword_5d4594_816372][a1]; }
