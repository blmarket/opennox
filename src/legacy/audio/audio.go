package audio

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

type AudioExterns struct {
	Dword_5d4594_1193336 uint32
	Dword_5d4594_805980  *AudioStructXxx

	Dword_587000_126996           *uint32
	Dword_5d4594_1045420          **AudioStructXxx
	Dword_5d4594_1045424          **Struct28[[0x2000]byte]
	Dword_5d4594_1045428          **Struct264
	Dword_5d4594_1045432          *uint32
	Dword_5d4594_1045436          **FreeList[Struct576]
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
	ListHeads_5d4594_839892       *[6][10]ListElement[Struct200Field28]
	ListHead_5d4594_840612        *ListElement[Struct576]
	Struct200Arr_5d4594_840628    *[1023]Struct200
	Ptr_uint32_5d4594_1045444     *uint32
}

type AudioModule struct {
	moduleName string
	Externs    *AudioExterns

	nox_platform_get_ticks func() uint64

	// Better in separate module
	nox_xxx_getSndName_40AF80         func(int) unsafe.Pointer
	nox_common_randomIntMinMax_415FF0 func(int, int, unsafe.Pointer, int) int
	// Better inline? just free?
	nox_common_list_append_4258E0  func(unsafe.Pointer, unsafe.Pointer)
	nox_common_list_remove_425920  func(unsafe.Pointer)
	nox_common_list_getNext_425940 func(unsafe.Pointer) unsafe.Pointer
	nox_binfile_fread_raw_40ADD0   func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32
	nox_fs_close                   func(unsafe.Pointer)
	nox_fs_fseek                   func(unsafe.Pointer, int32, int32) int32
	nox_fs_open                    func(unsafe.Pointer) unsafe.Pointer
}

func NewAudioModule(
	externs *AudioExterns,
	moduleName string,
	nox_platform_get_ticks func() uint64,
	nox_xxx_getSndName_40AF80 func(int) unsafe.Pointer,
	nox_common_randomIntMinMax_415FF0 func(int, int, unsafe.Pointer, int) int,
	nox_common_list_append_4258E0 func(unsafe.Pointer, unsafe.Pointer),
	nox_common_list_remove_425920 func(unsafe.Pointer),
	nox_common_list_getNext_425940 func(unsafe.Pointer) unsafe.Pointer,
	sub_486B60 func(int, int) int,
	sub_486E00 func(int) unsafe.Pointer,
	sub_486DB0 func(int, unsafe.Pointer, int) int,
	nox_binfile_fread_raw_40ADD0 func(unsafe.Pointer, uint32, uint32, unsafe.Pointer) int32,
	nox_fs_close func(unsafe.Pointer),
	nox_fs_fseek func(unsafe.Pointer, int32, int32) int32,
	nox_fs_open func(unsafe.Pointer) unsafe.Pointer,
) *AudioModule {
	return &AudioModule{
		moduleName:                        moduleName,
		Externs:                           externs,
		nox_platform_get_ticks:            nox_platform_get_ticks,
		nox_xxx_getSndName_40AF80:         nox_xxx_getSndName_40AF80,
		nox_common_randomIntMinMax_415FF0: nox_common_randomIntMinMax_415FF0,
		nox_common_list_append_4258E0:     nox_common_list_append_4258E0,
		nox_common_list_remove_425920:     nox_common_list_remove_425920,
		nox_common_list_getNext_425940:    nox_common_list_getNext_425940,
		nox_binfile_fread_raw_40ADD0:      nox_binfile_fread_raw_40ADD0,
		nox_fs_close:                      nox_fs_close,
		nox_fs_fseek:                      nox_fs_fseek,
		nox_fs_open:                       nox_fs_open,
	}
}
