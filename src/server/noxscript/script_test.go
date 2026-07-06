package noxscript

import (
	"image"
	"testing"

	"github.com/noxworld-dev/noxscript/ns/v4"
	"github.com/noxworld-dev/opennox-lib/types"
)

type mockVM struct {
	stack []interface{}
}

func (m *mockVM) NoxScript() ns.Implementation { return nil }
func (m *mockVM) DPos() image.Point { return image.Pt(1, 2) }
func (m *mockVM) DPosf() types.Pointf { return types.Pointf{X: 1, Y: 2} }
func (m *mockVM) NameSuff() string { return "test" }
func (m *mockVM) PushU32(v uint32) { m.stack = append(m.stack, v) }
func (m *mockVM) PushI32(v int32) { m.stack = append(m.stack, v) }
func (m *mockVM) PushF32(v float32) { m.stack = append(m.stack, v) }
func (m *mockVM) PushBool(v bool) { m.stack = append(m.stack, v) }
func (m *mockVM) PushString(s string) { m.stack = append(m.stack, s) }
func (m *mockVM) PushHandleNS(h ns.Handle) { m.stack = append(m.stack, h) }
func (m *mockVM) PopI32() int32 {
	if len(m.stack) == 0 {
		return 0
	}
	v := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if iv, ok := v.(int32); ok {
		return iv
	}
	if iv, ok := v.(int); ok {
		return int32(iv)
	}
	return 0
}
func (m *mockVM) PopU32() uint32 { return 0 }
func (m *mockVM) PopF32() float32 {
	if len(m.stack) == 0 {
		return 0
	}
	v := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if fv, ok := v.(float32); ok {
		return fv
	}
	return 0
}
func (m *mockVM) PopBool() bool { return false }
func (m *mockVM) PopString() string {
	if len(m.stack) == 0 {
		return ""
	}
	v := m.stack[len(m.stack)-1]
	m.stack = m.stack[:len(m.stack)-1]
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
func (m *mockVM) PopPoint() image.Point { return image.Pt(0, 0) }
func (m *mockVM) PopPointf() types.Pointf { return types.Pointf{} }
func (m *mockVM) PopWallGroupNS() ns.WallGroupObj { return nil }
func (m *mockVM) PopWaypointNS() ns.WaypointObj { return nil }
func (m *mockVM) PopWpGroupNS() ns.WaypointGroupObj { return nil }
func (m *mockVM) PopObjectNS() ns.Obj { return nil }
func (m *mockVM) PopObjGroupNS() ns.ObjGroup { return nil }

func TestNSFunctions(t *testing.T) {
	vm := &mockVM{}
	if nsAbort(vm) != 1 {
		t.Fatalf("expected 1")
	}
	vm.stack = append(vm.stack, int32(42))
	nsIntToString(vm)
	if len(vm.stack) == 0 {
		t.Fatalf("expected stack")
	}
	vm.stack = append(vm.stack, float32(3.14))
	nsFloatToString(vm)
	_ = nsRandomFloat
	_ = nsRandomInt
}
