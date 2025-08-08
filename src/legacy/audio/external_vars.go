package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

// AudioExternalVars consolidates all external variables used across audio modules
// to avoid duplication and provide a single source of truth
type AudioExternalVars struct {
	// Common external variables shared across modules
	Dword_587000_126996       *uint32
	Dword_5d4594_1045420      *uint32
	Dword_5d4594_1045424      *uint32
	Dword_5d4594_1045428      **Struct264
	Dword_5d4594_1045432      *uint32
	Dword_5d4594_1045436      *uint32
	Dword_587000_127004       unsafe.Pointer
	Dword_587000_155144       *unsafe.Pointer
	Dword_5d4594_805984       *unsafe.Pointer
	Dword_587000_155144_Struct **Struct587000_155144

	// Timer-related variables
	TimerGroup_5d4594_1045228 *timer.TimerGroup

	// List-related variables
	ListHeads_5d4594_839892 *[6][10]ListHead[Struct200Field28, *Struct200Field28]
	ListHead_5d4594_840612  *ListHead[Struct576, *Struct576]

	// Function pointers
	Sub_4873C0_ptr                   unsafe.Pointer
	Sub_4BD8C0_ptr                   unsafe.Pointer
	Sub_4BD940_ptr                   unsafe.Pointer
	Sub_4BD9B0_ptr                   unsafe.Pointer
	Ptr_uint32_5d4594_1193340        *uint32

	// Platform and utility functions
	Nox_platform_get_ticks           func() uint64
	Sub_452770_ptr                   unsafe.Pointer
	Sub_4526F0_ptr                   unsafe.Pointer
	Sub_4526D0_ptr                   unsafe.Pointer
}

// NewAudioExternalVars creates a new instance of shared external variables
func NewAudioExternalVars(
	dword_587000_126996 *uint32,
	dword_5d4594_1045420 *uint32,
	dword_5d4594_1045424 *uint32,
	dword_5d4594_1045428 **Struct264,
	dword_5d4594_1045432 *uint32,
	dword_5d4594_1045436 *uint32,
	dword_587000_127004 unsafe.Pointer,
	dword_587000_155144 *unsafe.Pointer,
	dword_5d4594_805984 *unsafe.Pointer,
	dword_587000_155144_struct **Struct587000_155144,
	timerGroup_5d4594_1045228 *timer.TimerGroup,
	listHeads_5d4594_839892 *[6][10]ListHead[Struct200Field28, *Struct200Field28],
	listHead_5d4594_840612 *ListHead[Struct576, *Struct576],
	sub_4873C0_ptr unsafe.Pointer,
	sub_4BD8C0_ptr unsafe.Pointer,
	sub_4BD940_ptr unsafe.Pointer,
	sub_4BD9B0_ptr unsafe.Pointer,
	ptr_uint32_5d4594_1193340 *uint32,
	nox_platform_get_ticks func() uint64,
	sub_452770_ptr unsafe.Pointer,
	sub_4526F0_ptr unsafe.Pointer,
	sub_4526D0_ptr unsafe.Pointer,
) *AudioExternalVars {
	return &AudioExternalVars{
		Dword_587000_126996:              dword_587000_126996,
		Dword_5d4594_1045420:             dword_5d4594_1045420,
		Dword_5d4594_1045424:             dword_5d4594_1045424,
		Dword_5d4594_1045428:             dword_5d4594_1045428,
		Dword_5d4594_1045432:             dword_5d4594_1045432,
		Dword_5d4594_1045436:             dword_5d4594_1045436,
		Dword_587000_127004:              dword_587000_127004,
		Dword_587000_155144:              dword_587000_155144,
		Dword_5d4594_805984:              dword_5d4594_805984,
		Dword_587000_155144_Struct:       dword_587000_155144_struct,
		TimerGroup_5d4594_1045228:        timerGroup_5d4594_1045228,
		ListHeads_5d4594_839892:          listHeads_5d4594_839892,
		ListHead_5d4594_840612:           listHead_5d4594_840612,
		Sub_4873C0_ptr:                   sub_4873C0_ptr,
		Sub_4BD8C0_ptr:                   sub_4BD8C0_ptr,
		Sub_4BD940_ptr:                   sub_4BD940_ptr,
		Sub_4BD9B0_ptr:                   sub_4BD9B0_ptr,
		Ptr_uint32_5d4594_1193340:        ptr_uint32_5d4594_1193340,
		Nox_platform_get_ticks:           nox_platform_get_ticks,
		Sub_452770_ptr:                   sub_452770_ptr,
		Sub_4526F0_ptr:                   sub_4526F0_ptr,
		Sub_4526D0_ptr:                   sub_4526D0_ptr,
	}
}