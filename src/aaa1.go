package opennox

import (
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/common/memmap"
	"github.com/noxworld-dev/opennox/v1/legacy"
	"github.com/noxworld-dev/opennox/v1/legacy/common/ccall"
	"github.com/noxworld-dev/opennox/v1/legacy/timer"
)

// Some idea how to define linked lists in Golang
type ListItem[T any] struct {
	ptr0 *T
	ptr1 *T
	ptr2 *T
}

func (it *ListItem[T]) Self() *T {
	return (*T)(unsafe.Pointer(it))
}

func (it *ListItem[T]) Clear() {
	it.ptr0 = it.Self()
	it.ptr1 = it.Self()
	it.ptr2 = it.Self()
}

// See sub_4871C0 for where it created
type Struct264 struct {
	field_0   ListItem[Struct264]
	field_12  uint32
	field_16  uint32
	field_20  unsafe.Pointer
	field_24  unsafe.Pointer
	field_28  uint32
	field_32  uint32
	field_36  [6]uint32
	field_60  [7]uint32
	field_88  timer.TimerGroup
	field_184 uint32
	field_188 uint32
	field_192 uint32
	field_196 uint32
	field_200 ListItem[unsafe.Pointer]
	field_212 uint32
	field_216 unsafe.Pointer // function pointer
	field_220 [11]uint32
}

var _ = [1]struct{}{}[264-unsafe.Sizeof(Struct264{})]

type List155144 struct {
	field_0  ListItem[unsafe.Pointer] // TODO: clarify types
	field_12 ListItem[Struct264]
	field_24 int32
	field_28 uint32 // unknown
	field_32 timer.TimerGroup
}

func inst() *List155144 {
	return (*List155144)(legacy.Get_dword_587000_155144())
}

func sub_486F30() int {
	inst().field_0.Clear()
	inst().field_12.Clear()
	inst().field_24 = 0
	*memmap.PtrT[*timer.TimerGroup](0x5D4594, 1193340) = &(inst().field_32)
	inst().field_32.Init()
	dword_5d4594_1193336 = 1
	return 0
}

func sub_486EF0() {
	// fmt.Printf("sub_486EF0\n")
	if dword_5d4594_1193336 == 0 {
		return
	}
	if inst().field_24 != 0 {
		return
	}
	v1 := inst().field_12.ptr0
	for it := inst().field_12.Self(); v1 != it; v1 = v1.field_0.ptr0 {
		if v1.field_12&2 != 0 {
			continue
		}
		ccall.CallVoidPtr(v1.field_216, unsafe.Pointer(v1))
	}
}

func sub_487050(a1 unsafe.Pointer) {
	nox_common_list_append_4258E0(unsafe.Pointer(&inst().field_0), a1)
}

func sub_4870E0(a1 *unsafe.Pointer) unsafe.Pointer {
	result := nox_common_list_getFirstSafe_425890(unsafe.Pointer(&inst().field_0))
	*a1 = result
	return result
}

func sub_487310(a1 unsafe.Pointer) {
	inst().field_24 += 1
	nox_common_list_append_4258E0(unsafe.Pointer(&inst().field_12), a1)
	result := inst().field_24 - 1
	inst().field_24 = result
	if result < 0 {
		inst().field_24 = 0
	}
}

func sub_4875B0(a1 *unsafe.Pointer) unsafe.Pointer {
	result := nox_common_list_getFirstSafe_425890(unsafe.Pointer(&inst().field_12))
	*a1 = result
	return result
}

func sub_4875D0(a1 *unsafe.Pointer) unsafe.Pointer {
	if *a1 != nil {
		*a1 = nox_common_list_getNextSafe_4258A0(*a1)
	}
	return *a1
}

func sub_4875F0() int32 {
	var v3 unsafe.Pointer
	inst().field_24 += 1
	v0 := sub_4875B0(&v3)
	for v0 != nil {
		v1 := sub_4875D0(&v3)
		legacy.Sub_487680(v0)
		v0 = v1
	}
	result := inst().field_24 - 1
	inst().field_24 = result
	if result < 0 {
		inst().field_24 = 0
	}
	return result
}

func nox_common_list_remove_425920(a1 unsafe.Pointer) {
	(*listItem)(a1).Remove()
}

func sub_4876A0(a1 unsafe.Pointer) {
	inst().field_24 += 1
	nox_common_list_remove_425920(a1)
	result := inst().field_24 - 1
	inst().field_24 = result
	if result < 0 {
		inst().field_24 = 0
	}
}
