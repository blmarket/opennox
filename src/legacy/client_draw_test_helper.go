package legacy

/*
#include <stdlib.h>

#include "client__draw__armordraw.h"
#include "client__draw__basedraw.h"
#include "client__draw__staticdraw.h"
#include "client__draw__weapondraw.h"

static nox_drawable* test_skip_drawable(void) {
	nox_drawable* dr = (nox_drawable*)calloc(1, sizeof(nox_drawable));
	if (!dr) {
		return NULL;
	}
	dr->flags28 = 0x40000;
	return dr;
}

static int test_nox_thing_static_draw_skip(void) {
	nox_drawable* dr = test_skip_drawable();
	if (!dr) {
		return -1;
	}
	int ret = nox_thing_static_draw(NULL, dr);
	free(dr);
	return ret;
}

static int test_nox_thing_weapon_draw_skip(void) {
	nox_drawable* dr = test_skip_drawable();
	if (!dr) {
		return -1;
	}
	int ret = nox_thing_weapon_draw(NULL, dr);
	free(dr);
	return ret;
}

static int test_nox_thing_armor_draw_skip(void) {
	nox_drawable* dr = test_skip_drawable();
	if (!dr) {
		return -1;
	}
	int ret = nox_thing_armor_draw(NULL, dr);
	free(dr);
	return ret;
}

static int test_nox_thing_base_draw_skip(void) {
	nox_drawable* dr = test_skip_drawable();
	if (!dr) {
		return -1;
	}
	int ret = nox_thing_base_draw(NULL, dr);
	free(dr);
	return ret;
}
*/
import "C"

func C_nox_thing_static_draw_skip() int {
	return int(C.test_nox_thing_static_draw_skip())
}

func C_nox_thing_weapon_draw_skip() int {
	return int(C.test_nox_thing_weapon_draw_skip())
}

func C_nox_thing_armor_draw_skip() int {
	return int(C.test_nox_thing_armor_draw_skip())
}

func C_nox_thing_base_draw_skip() int {
	return int(C.test_nox_thing_base_draw_skip())
}
