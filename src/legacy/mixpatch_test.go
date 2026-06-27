package legacy

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/server"
)

type mockServer struct {
	Server
	srv *server.Server
}

func (m mockServer) S() *server.Server {
	return m.srv
}

func TestMixPatchSub9805EB(t *testing.T) {
	// 1. unit = nil
	res := C_sub_9805EB(nil)
	if res != nil {
		t.Errorf("Expected nil for nil unit, got %p", res)
	}

	// 2. unit with empty inventory
	unit := AllocNoxObject()
	defer FreeNoxObject(unit)
	res = C_sub_9805EB(unit)
	if res != nil {
		t.Errorf("Expected nil for empty inventory, got %p", res)
	}

	// 3. unit with items not matching class/flags
	item1 := AllocNoxObject()
	defer FreeNoxObject(item1)
	SetNoxObjectFields(item1, 0, 0, 0)
	SetNoxObjectInventory(unit, item1)

	res = C_sub_9805EB(unit)
	if res != nil {
		t.Errorf("Expected nil when items don't match, got %p", res)
	}

	// 4. unit with items matching class but not flags (flags must be exactly 16)
	item2 := AllocNoxObject()
	defer FreeNoxObject(item2)
	SetNoxObjectFields(item2, 0x2000000, 0, 0)
	SetNoxObjectNextItem(item1, item2)

	res = C_sub_9805EB(unit)
	if res != nil {
		t.Errorf("Expected nil when flags don't match, got %p", res)
	}

	// 5. unit with items matching both class and flags
	item3 := AllocNoxObject()
	defer FreeNoxObject(item3)
	SetNoxObjectFields(item3, 0x2000000, 16, 0)
	SetNoxObjectNextItem(item2, item3)

	res = C_sub_9805EB(unit)
	if res != item3 {
		t.Errorf("Expected %p, got %p", item3, res)
	}
}

func TestMixPatchSub980523(t *testing.T) {
	// 1. unit = nil (should return immediately without crash)
	C_sub_980523(nil)

	// 2. unit with empty inventory (should return immediately without crash)
	unit := AllocNoxObject()
	defer FreeNoxObject(unit)
	C_sub_980523(unit)

	// 3. Mock the server and initialize the armor table to test the full logic
	oldGetServer := GetServer
	defer func() { GetServer = oldGetServer }()

	srv := &server.Server{}
	srv.Types.ClientTypeByID = func(id string) int {
		if id == "SteelShield" {
			return 42
		}
		return 0
	}
	srv.Nox_xxx_equipArmor_415AB0()

	GetServer = func() Server {
		return mockServer{srv: srv}
	}

	// Prepare mock buffers for pointer assignment
	buf1 := AllocBuffer(300)
	defer FreeBuffer(buf1)
	buf2 := AllocBuffer(2600)
	defer FreeBuffer(buf2)

	// buf1 + 276 points to buf2
	*(*uintptr)(unsafe.Pointer(uintptr(buf1) + 276)) = uintptr(buf2)

	SetNoxObjectDataUpdate(unit, buf1)

	// Add item to inventory
	item := AllocNoxObject()
	defer FreeNoxObject(item)
	SetNoxObjectFields(item, 0x2000000, 0x100, 42) // Matches the ClientTypeByID lookup for SteelShield

	SetNoxObjectInventory(unit, item)

	C_sub_980523(unit)

	// Verify that item pointer was successfully written to buf2 + 2500
	assignedVal := *(*uintptr)(unsafe.Pointer(uintptr(buf2) + 2500))
	if assignedVal != uintptr(item) {
		t.Errorf("Expected assigned value to be %p, got %x", item, assignedVal)
	}
}
