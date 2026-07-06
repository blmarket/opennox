#include "memfile.h"

int8_t nox_memfile_read_i8(nox_memfile* f) {
	if (!f->data)
		return 0;
	int8_t v = *(int8_t*)f->cur;
	f->cur++;
	return v;
}

uint8_t nox_memfile_read_u8(nox_memfile* f) {
	if (!f->data)
		return 0;
	uint8_t v = *(uint8_t*)f->cur;
	f->cur++;
	return v;
}

int16_t nox_memfile_read_i16(nox_memfile* f) {
	if (!f->data)
		return 0;
	int16_t v = *(int16_t*)f->cur;
	f->cur += 2;
	return v;
}

uint16_t nox_memfile_read_u16(nox_memfile* f) {
	if (!f->data)
		return 0;
	uint16_t v = *(uint16_t*)f->cur;
	f->cur += 2;
	return v;
}

void nox_memfile_skip(nox_memfile* f, int n) {
	if (!f->data)
		return;
	f->cur += n;
}
