package legacy

/*
#include <stdint.h>

// U32 memmap getters (return *getMemU32Ptr(base, off))
int nox_xxx_get3512_40A350();
int sub_40AA20();
int sub_416580();
int sub_416650();
int sub_4169C0();
int sub_41D1B0();
int sub_4200E0();
int sub_4207E0();
int nox_xxx_wallGet_426A30();
int nox_gui_getQuestStage_450B10();
int sub_453610();
int nox_xxx_guiSpell_460650();
int sub_461450();
int sub_469FA0();
int nox_xxx_guiCursor_477600();
int sub_4CFE00();
int nox_xxx_scavengerTreasureMax_4D1600();
int nox_xxx_isQuest_4D6F50();
int sub_4D6F70();
int sub_4D6FA0();
int sub_4D7300();
int sub_4D75E0();
int sub_51A950();

// Byte memmap getters (return getMemByte(base, off))
unsigned char sub_450750();
int sub_4604E0();
unsigned char sub_467430();
unsigned char sub_47DBC0();

// Constant memmap setters (*getMemU32Ptr(base, off) = const)
void nox_xxx_wallBreakableCounterClear_429520();
void nox_xxx_wallSecretCounterClear_4297B0();
void sub_4573B0();
void nox_xxx____setargv_11_473920();
void sub_4D15C0();
void sub_4D1610();
*/
import "C"

// --- U32 getters ---

func C_nox_xxx_get3512_40A350() int          { return int(C.nox_xxx_get3512_40A350()) }
func C_sub_40AA20() int                       { return int(C.sub_40AA20()) }
func C_sub_416580() int                       { return int(C.sub_416580()) }
func C_sub_416650() int                       { return int(C.sub_416650()) }
func C_sub_4169C0() int                       { return int(C.sub_4169C0()) }
func C_sub_41D1B0() int                       { return int(C.sub_41D1B0()) }
func C_sub_4200E0() int                       { return int(C.sub_4200E0()) }
func C_sub_4207E0() int                       { return int(C.sub_4207E0()) }
func C_nox_xxx_wallGet_426A30() int           { return int(C.nox_xxx_wallGet_426A30()) }
func C_nox_gui_getQuestStage_450B10() int     { return int(C.nox_gui_getQuestStage_450B10()) }
func C_sub_453610() int                       { return int(C.sub_453610()) }
func C_nox_xxx_guiSpell_460650() int          { return int(C.nox_xxx_guiSpell_460650()) }
func C_sub_461450() int                       { return int(C.sub_461450()) }
func C_sub_469FA0() int                       { return int(C.sub_469FA0()) }
func C_nox_xxx_guiCursor_477600() int         { return int(C.nox_xxx_guiCursor_477600()) }
func C_sub_4CFE00() int                       { return int(C.sub_4CFE00()) }
func C_nox_xxx_scavengerTreasureMax_4D1600() int { return int(C.nox_xxx_scavengerTreasureMax_4D1600()) }
func C_nox_xxx_isQuest_4D6F50() int           { return int(C.nox_xxx_isQuest_4D6F50()) }
func C_sub_4D6F70() int                       { return int(C.sub_4D6F70()) }
func C_sub_4D6FA0() int                       { return int(C.sub_4D6FA0()) }
func C_sub_4D7300() int                       { return int(C.sub_4D7300()) }
func C_sub_4D75E0() int                       { return int(C.sub_4D75E0()) }
func C_sub_51A950() int                       { return int(C.sub_51A950()) }

// --- Byte getters ---

func C_sub_450750() uint8 { return uint8(C.sub_450750()) }
func C_sub_4604E0() int   { return int(C.sub_4604E0()) }
func C_sub_467430() uint8 { return uint8(C.sub_467430()) }
func C_sub_47DBC0() uint8 { return uint8(C.sub_47DBC0()) }

// --- Constant setters ---

func C_nox_xxx_wallBreakableCounterClear_429520() { C.nox_xxx_wallBreakableCounterClear_429520() }
func C_nox_xxx_wallSecretCounterClear_4297B0()    { C.nox_xxx_wallSecretCounterClear_4297B0() }
func C_sub_4573B0()                               { C.sub_4573B0() }
func C_nox_xxx____setargv_11_473920()             { C.nox_xxx____setargv_11_473920() }
func C_sub_4D15C0()                               { C.sub_4D15C0() }
func C_sub_4D1610()                               { C.sub_4D1610() }
