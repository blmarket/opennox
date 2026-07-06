package legacy

/*
#include <stdlib.h>
#include "client__draw__arrowdraw.h"
#include "client__draw__boulderdraw.h"
#include "client__draw__bubbledraw.h"
#include "client__draw__doordraw.h"
#include "client__draw__flagdraw.h"
#include "client__draw__glowdraw.h"
#include "client__draw__glyphdraw.h"
#include "client__draw__lightning.h"
#include "client__draw__plasma.h"

static nox_drawable* test_skip_drawable3(void) {
	nox_drawable* dr = (nox_drawable*)calloc(1, sizeof(nox_drawable));
	if (!dr) return NULL;
	dr->flags28 = 0x40000;
	return dr;
}

static void test_draw_arrow(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_arrow_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_boulder(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_boulder_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_bubble(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_bubble_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_door(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_door_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_flag(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_flag_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_glow(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_glow_orb_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_glyph(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_glyph_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_lightning(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_lightning_draw(NULL, dr);
		free(dr);
	}
}

static void test_draw_plasma(void) {
	nox_drawable* dr = test_skip_drawable3();
	if (dr) {
		nox_thing_plasma_draw(NULL, dr);
		free(dr);
	}
}
*/
import "C"

func C_testDrawArrow()     { C.test_draw_arrow() }
func C_testDrawBoulder()   { C.test_draw_boulder() }
func C_testDrawBubble()    { C.test_draw_bubble() }
func C_testDrawDoor()      { C.test_draw_door() }
func C_testDrawFlag()      { C.test_draw_flag() }
func C_testDrawGlow()      { C.test_draw_glow() }
func C_testDrawGlyph()     { C.test_draw_glyph() }
func C_testDrawLightning() { C.test_draw_lightning() }
func C_testDrawPlasma()    { C.test_draw_plasma() }
