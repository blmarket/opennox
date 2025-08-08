package audio

import (
	"unsafe"
)

type Phase6Module struct {
	moduleName string
	externs    *AudioExterns

	// External variables

	// External functions
	nox_platform_get_ticks        func() uint64
	sub_486620                    func(unsafe.Pointer)
	sub_486550                    func(unsafe.Pointer) int
	sub_486570                    func(unsafe.Pointer, unsafe.Pointer)
	sub_4864A0                    func(unsafe.Pointer) unsafe.Pointer
	sub_4BD7A0                    func(unsafe.Pointer)
	sub_486520                    func(unsafe.Pointer) int
	nox_common_list_remove_425920 func(unsafe.Pointer)
}

func NewPhase6Module(
	externs *AudioExterns,
	moduleName string,
	nox_platform_get_ticks func() uint64,
	sub_486620 func(unsafe.Pointer),
	sub_486550 func(unsafe.Pointer) int,
	sub_486570 func(unsafe.Pointer, unsafe.Pointer),
	sub_4864A0 func(unsafe.Pointer) unsafe.Pointer,
	sub_4BD7A0 func(unsafe.Pointer),
	sub_486520 func(unsafe.Pointer) int,
	nox_common_list_remove_425920 func(unsafe.Pointer),
) *Phase6Module {
	return &Phase6Module{
		moduleName:                    moduleName,
		externs:                       externs,
		nox_platform_get_ticks:        nox_platform_get_ticks,
		sub_486620:                    sub_486620,
		sub_486550:                    sub_486550,
		sub_486570:                    sub_486570,
		sub_4864A0:                    sub_4864A0,
		sub_4BD7A0:                    sub_4BD7A0,
		sub_486520:                    sub_486520,
		nox_common_list_remove_425920: nox_common_list_remove_425920,
	}
}
