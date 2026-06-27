package legacy

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// These characterization tests exercise the many tiny decompiled GAME*.c
// accessor functions that read or write a single fixed location in the
// simulated legacy memory map. We poke a sentinel value through the memmap
// helpers, call the legacy accessor, and assert it observes / mutates the
// same word. Old values are restored so we don't disturb integration tests.

func TestGameMemmapU32Getters(t *testing.T) {
	const sentinel = 0x12345678

	cases := []struct {
		name string
		base uintptr
		off  uintptr
		get  func() int
	}{
		{"get3512_40A350", 0x5D4594, 3512, C_nox_xxx_get3512_40A350},
		{"sub_40AA20", 0x5D4594, 3536, C_sub_40AA20},
		{"sub_416580", 0x5D4594, 371688, C_sub_416580},
		{"sub_416650", 0x5D4594, 371700, C_sub_416650},
		{"sub_4169C0", 0x5D4594, 371704, C_sub_4169C0},
		{"sub_41D1B0", 0x5D4594, 527720, C_sub_41D1B0},
		{"sub_4200E0", 0x587000, 60072, C_sub_4200E0},
		{"sub_4207E0", 0x5D4594, 534812, C_sub_4207E0},
		{"wallGet_426A30", 0x5D4594, 739992, C_nox_xxx_wallGet_426A30},
		{"getQuestStage_450B10", 0x5D4594, 832468, C_nox_gui_getQuestStage_450B10},
		{"sub_453610", 0x5D4594, 1045456, C_sub_453610},
		{"guiSpell_460650", 0x5D4594, 1047928, C_nox_xxx_guiSpell_460650},
		{"sub_461450", 0x5D4594, 1049688, C_sub_461450},
		{"sub_469FA0", 0x5D4594, 1064848, C_sub_469FA0},
		{"guiCursor_477600", 0x5D4594, 1096672, C_nox_xxx_guiCursor_477600},
		{"sub_4CFE00", 0x5D4594, 1523072, C_sub_4CFE00},
		{"scavengerTreasureMax_4D1600", 0x5D4594, 1548528, C_nox_xxx_scavengerTreasureMax_4D1600},
		{"isQuest_4D6F50", 0x5D4594, 1556160, C_nox_xxx_isQuest_4D6F50},
		{"sub_4D6F70", 0x5D4594, 1556164, C_sub_4D6F70},
		{"sub_4D6FA0", 0x5D4594, 1556104, C_sub_4D6FA0},
		{"sub_4D7300", 0x5D4594, 1556132, C_sub_4D7300},
		{"sub_4D75E0", 0x5D4594, 1556120, C_sub_4D75E0},
		{"sub_51A950", 0x5D4594, 2388656, C_sub_51A950},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := (*uint32)(C_mem_getU32Ptr(tc.base, tc.off))
			require.NotNil(t, p)
			old := *p
			t.Cleanup(func() { *p = old })

			*p = sentinel
			require.Equal(t, sentinel, tc.get())
		})
	}
}

func TestGameMemmapByteGetters(t *testing.T) {
	const sentinel = 0xAB

	cases := []struct {
		name string
		base uintptr
		off  uintptr
		get  func() int // normalized to int for comparison
	}{
		{"sub_450750", 0x5D4594, 831252, func() int { return int(C_sub_450750()) }},
		{"sub_4604E0", 0x5D4594, 1048140, C_sub_4604E0},
		{"sub_467430", 0x5D4594, 1062536, func() int { return int(C_sub_467430()) }},
		{"sub_47DBC0", 0x5D4594, 1193128, func() int { return int(C_sub_47DBC0()) }},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := (*uint8)(C_mem_getU8Ptr(tc.base, tc.off))
			require.NotNil(t, p)
			old := *p
			t.Cleanup(func() { *p = old })

			*p = sentinel
			require.Equal(t, sentinel, tc.get())
		})
	}
}

func TestGameMemmapConstSetters(t *testing.T) {
	cases := []struct {
		name string
		base uintptr
		off  uintptr
		set  func()
		want uint32
	}{
		{"wallBreakableCounterClear_429520", 0x5D4594, 741344, C_nox_xxx_wallBreakableCounterClear_429520, 0},
		{"wallSecretCounterClear_4297B0", 0x5D4594, 741352, C_nox_xxx_wallSecretCounterClear_4297B0, 0},
		{"sub_4573B0", 0x5D4594, 1045696, C_sub_4573B0, 0},
		{"setargv_11_473920", 0x5D4594, 1096520, C_nox_xxx____setargv_11_473920, 1},
		{"sub_4D15C0", 0x5D4594, 1548508, C_sub_4D15C0, 0},
		{"sub_4D1610", 0x5D4594, 1548528, C_sub_4D1610, 0},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			p := (*uint32)(C_mem_getU32Ptr(tc.base, tc.off))
			require.NotNil(t, p)
			old := *p
			t.Cleanup(func() { *p = old })

			// Seed with a value distinct from the constant the setter writes.
			*p = 0xCAFEF00D
			tc.set()
			require.Equal(t, tc.want, *p)
		})
	}
}
