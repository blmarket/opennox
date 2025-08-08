package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type AudioExterns struct {
	Dword_587000_126996          *uint32
	Dword_5d4594_1045420         *uint32
	Dword_5d4594_1045424         *uint32
	Dword_5d4594_1045428         **Struct264
	Dword_5d4594_1045432         *uint32
	Dword_5d4594_1045436         *uint32
	Ptr_TimerGroup_587000_127004 **timer.TimerGroup
}

type AudioModule struct {
	moduleName string
	externs    *AudioExterns

	nox_platform_get_ticks func() uint64
	sub_452770_ptr         unsafe.Pointer
	sub_4526F0_ptr         unsafe.Pointer
	sub_4526D0_ptr         unsafe.Pointer

	// External variables
	dword_5d4594_1045420         *uint32
	dword_5d4594_1045424         *uint32
	dword_5d4594_1045428         **Struct264
	dword_5d4594_1045432         *uint32
	dword_5d4594_1045436         *uint32
	ptr_TimerGroup_587000_127004 **timer.TimerGroup
	timerGroup_5d4594_1045228    *timer.TimerGroup
	listHeads_5d4594_839892      *[6][10]ListHead[Struct200Field28, *Struct200Field28]
	listHead_5d4594_840612       *ListHead[Struct576, *Struct576]

	// External functions
	sub_425770                          func(unsafe.Pointer) unsafe.Pointer
	nox_common_list_append_4258E0       func(unsafe.Pointer, unsafe.Pointer)
	sub_4BDB30                          func(int) int
	sub_4BD300                          func(unsafe.Pointer, int) int
	sub_4BDB90                          func(unsafe.Pointer, unsafe.Pointer)
	sub_4BDB40                          func(int) int
	sub_4864A0                          func(unsafe.Pointer) unsafe.Pointer
	sub_486350                          func(unsafe.Pointer, int) int
	sub_486520                          func(unsafe.Pointer) int
	sub_4BD280                          func(int, int) unsafe.Pointer
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer
	sub_4BD340                          func(int, int, int, int) unsafe.Pointer
	sub_4BD2E0                          func(unsafe.Pointer) unsafe.Pointer
	sub_4BD470                          func(unsafe.Pointer, int) unsafe.Pointer
	sub_452810                          func(int, byte) unsafe.Pointer
	sub_4BD2D0                          func(unsafe.Pointer)
	sub_4BDA80                          func(int) int
	nox_common_list_clear_425760        func(unsafe.Pointer)
	sub_486320                          func(unsafe.Pointer, int) unsafe.Pointer
	nox_xxx_getSndName_40AF80           func(int) unsafe.Pointer
	nox_common_list_remove_425920       func(unsafe.Pointer)
	sub_4862E0                          func(unsafe.Pointer, int) int
	nox_common_randomIntMinMax_415FF0   func(int, int, unsafe.Pointer, int) int
	sub_4863B0                          func(unsafe.Pointer) int
	sub_4BD3C0                          func(unsafe.Pointer)
}

func NewAudioModule(
	externs *AudioExterns,
	moduleName string,
	nox_platform_get_ticks func() uint64,
	sub_452770_ptr unsafe.Pointer,
	sub_4526F0_ptr unsafe.Pointer,
	sub_4526D0_ptr unsafe.Pointer,
	dword_587000_126996 *uint32,
	dword_5d4594_1045420 *uint32,
	dword_5d4594_1045424 *uint32,
	dword_5d4594_1045428 **Struct264,
	dword_5d4594_1045432 *uint32,
	dword_5d4594_1045436 *uint32,
	dword_587000_127004 **timer.TimerGroup,
	timerGroup_5d4594_1045228 *timer.TimerGroup,
	listHeads_5d4594_839892 *[6][10]ListHead[Struct200Field28, *Struct200Field28],
	listHead_5d4594_840612 *ListHead[Struct576, *Struct576],
	sub_425770 func(unsafe.Pointer) unsafe.Pointer,
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
	sub_4BDB30 func(int) int,
	sub_4BD300 func(unsafe.Pointer, int) int,
	sub_4BDB90 func(unsafe.Pointer, unsafe.Pointer),
	sub_4BDB40 func(int) int,
	sub_4864A0 func(unsafe.Pointer) unsafe.Pointer,
	sub_486350 func(unsafe.Pointer, int) int,
	sub_486520 func(unsafe.Pointer) int,
	sub_4BD280 func(int, int) unsafe.Pointer,
	nox_common_list_getFirstSafe_425890 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BD340 func(int, int, int, int) unsafe.Pointer,
	sub_4BD2E0 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BD470 func(unsafe.Pointer, int) unsafe.Pointer,
	sub_452810 func(int, byte) unsafe.Pointer,
	sub_4BD2D0 func(unsafe.Pointer),
	sub_4BDA80 func(int) int,
	nox_common_list_clear_425760 func(unsafe.Pointer),
	sub_486320 func(unsafe.Pointer, int) unsafe.Pointer,
	nox_xxx_getSndName_40AF80 func(int) unsafe.Pointer,
	nox_common_list_remove_425920 func(unsafe.Pointer),
	sub_4862E0 func(unsafe.Pointer, int) int,
	nox_common_randomIntMinMax_415FF0 func(int, int, unsafe.Pointer, int) int,
	sub_4863B0 func(unsafe.Pointer) int,
	sub_4BD3C0 func(unsafe.Pointer),
) *AudioModule {
	return &AudioModule{
		moduleName:                          moduleName,
		externs:                             externs,
		nox_platform_get_ticks:              nox_platform_get_ticks,
		sub_452770_ptr:                      sub_452770_ptr,
		sub_4526F0_ptr:                      sub_4526F0_ptr,
		sub_4526D0_ptr:                      sub_4526D0_ptr,
		dword_5d4594_1045420:                dword_5d4594_1045420,
		dword_5d4594_1045424:                dword_5d4594_1045424,
		dword_5d4594_1045428:                dword_5d4594_1045428,
		dword_5d4594_1045432:                dword_5d4594_1045432,
		dword_5d4594_1045436:                dword_5d4594_1045436,
		ptr_TimerGroup_587000_127004:        dword_587000_127004,
		timerGroup_5d4594_1045228:           timerGroup_5d4594_1045228,
		listHeads_5d4594_839892:             listHeads_5d4594_839892,
		listHead_5d4594_840612:              listHead_5d4594_840612,
		sub_425770:                          sub_425770,
		nox_common_list_append_4258E0:       nox_common_list_append_4258E0,
		sub_4BDB30:                          sub_4BDB30,
		sub_4BD300:                          sub_4BD300,
		sub_4BDB90:                          sub_4BDB90,
		sub_4BDB40:                          sub_4BDB40,
		sub_4864A0:                          sub_4864A0,
		sub_486350:                          sub_486350,
		sub_486520:                          sub_486520,
		sub_4BD280:                          sub_4BD280,
		nox_common_list_getFirstSafe_425890: nox_common_list_getFirstSafe_425890,
		sub_4BD340:                          sub_4BD340,
		sub_4BD2E0:                          sub_4BD2E0,
		sub_4BD470:                          sub_4BD470,
		sub_452810:                          sub_452810,
		sub_4BD2D0:                          sub_4BD2D0,
		sub_4BDA80:                          sub_4BDA80,
		nox_common_list_clear_425760:        nox_common_list_clear_425760,
		sub_486320:                          sub_486320,
		nox_xxx_getSndName_40AF80:           nox_xxx_getSndName_40AF80,
		nox_common_list_remove_425920:       nox_common_list_remove_425920,
		sub_4862E0:                          sub_4862E0,
		nox_common_randomIntMinMax_415FF0:   nox_common_randomIntMinMax_415FF0,
		sub_4863B0:                          sub_4863B0,
		sub_4BD3C0:                          sub_4BD3C0,
	}
}
