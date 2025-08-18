package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type AudioExterns struct {
	Dword_5d4594_1193336 uint32

	Dword_587000_126996           *uint32
	Dword_5d4594_1045420          *uint32
	Dword_5d4594_1045424          *uint32
	Dword_5d4594_1045428          **Struct264
	Dword_5d4594_1045432          *uint32
	Dword_5d4594_1045436          *uint32
	Ptr_TimerGroup_587000_127004  **timer.TimerGroup
	Dword_587000_155144           **Struct587000_155144
	Sub_4873C0_ptr                unsafe.Pointer
	Sub_4BD8C0_ptr                unsafe.Pointer
	Sub_4BD940_ptr                unsafe.Pointer
	Sub_4BD9B0_ptr                unsafe.Pointer
	Dword_587000_127004           **timer.TimerGroup
	Dword_5d4594_805984           **Struct264
	Ptr_TimerGroup_5d4594_1193340 **timer.TimerGroup
	Sub_452770_ptr                unsafe.Pointer
	Sub_4526F0_ptr                unsafe.Pointer
	Sub_4526D0_ptr                unsafe.Pointer
	TimerGroup_5d4594_1045228     *timer.TimerGroup
	ListHeads_5d4594_839892       *[6][10]ListElement[Struct200Field28, *Struct200Field28]
	ListHead_5d4594_840612        *ListElement[Struct576, *Struct576]
	Struct200Arr_5d4594_840628    *[1023]Struct200
	Ptr_uint32_5d4594_1045444     *uint32
}

type AudioModule struct {
	moduleName string
	Externs    *AudioExterns

	nox_platform_get_ticks func() uint64

	// External functions
	sub_4BD300 func(unsafe.Pointer, int) int       // Release item back to free list
	sub_4BD280 func(int, int) unsafe.Pointer       // Create free list
	sub_4BD2E0 func(unsafe.Pointer) unsafe.Pointer // Get element from free list
	// Maybe move again?
	sub_4BD340 func(int, int, int, int) unsafe.Pointer
	sub_4BD3C0 func(unsafe.Pointer)
	// Better in separate module
	sub_4BD470                        func(unsafe.Pointer, int) unsafe.Pointer
	nox_xxx_getSndName_40AF80         func(int) unsafe.Pointer
	nox_common_randomIntMinMax_415FF0 func(int, int, unsafe.Pointer, int) int
	// Better inline? just free?
	sub_4BD2D0                    func(unsafe.Pointer)
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer)
}

func NewAudioModule(
	externs *AudioExterns,
	moduleName string,
	nox_platform_get_ticks func() uint64,
	sub_4BD300 func(unsafe.Pointer, int) int,
	sub_4BD280 func(int, int) unsafe.Pointer,
	sub_4BD340 func(int, int, int, int) unsafe.Pointer,
	sub_4BD2E0 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BD470 func(unsafe.Pointer, int) unsafe.Pointer,
	sub_4BD2D0 func(unsafe.Pointer),
	nox_xxx_getSndName_40AF80 func(int) unsafe.Pointer,
	nox_common_randomIntMinMax_415FF0 func(int, int, unsafe.Pointer, int) int,
	sub_4BD3C0 func(unsafe.Pointer),
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
) *AudioModule {
	return &AudioModule{
		moduleName:                        moduleName,
		Externs:                           externs,
		nox_platform_get_ticks:            nox_platform_get_ticks,
		sub_4BD300:                        sub_4BD300,
		sub_4BD280:                        sub_4BD280,
		sub_4BD340:                        sub_4BD340,
		sub_4BD2E0:                        sub_4BD2E0,
		sub_4BD470:                        sub_4BD470,
		sub_4BD2D0:                        sub_4BD2D0,
		nox_xxx_getSndName_40AF80:         nox_xxx_getSndName_40AF80,
		nox_common_randomIntMinMax_415FF0: nox_common_randomIntMinMax_415FF0,
		sub_4BD3C0:                        sub_4BD3C0,
		nox_common_list_append_4258E0:     nox_common_list_append_4258E0,
	}
}
