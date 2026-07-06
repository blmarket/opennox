package server

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestRegisterObjectFuncsExtra2(t *testing.T) {
	RegisterObjectCreate("TestCreateExtra2", unsafe.Pointer(nil))
	RegisterObjectInit("TestInitExtra2", unsafe.Pointer(nil), 0)
	RegisterObjectUpdate("TestUpdateExtra2", unsafe.Pointer(nil), 0)
	RegisterObjectCollide("TestCollideExtra2", unsafe.Pointer(nil), 0)
	RegisterObjectUse("TestUseExtra2", unsafe.Pointer(nil), 0)
	RegisterObjectUseParse("TestUseParseExtra2", func(objt *ObjectType, args []string) error { return nil })
	RegisterObjectDamage("TestDamageExtra2", unsafe.Pointer(nil))
	RegisterObjectDeath("TestDeathExtra2", unsafe.Pointer(nil), 0)
	RegisterObjectDrop("TestDropExtra2", unsafe.Pointer(nil))
	RegisterObjectPickup("TestPickupExtra2", unsafe.Pointer(nil))
	require.True(t, true)
}

func TestRegisterObjectPanics(t *testing.T) {
	name := "test_panic_" + t.Name()
	fn := unsafe.Pointer(&name)

	RegisterObjectCreate(name+"_create", fn)
	require.Panics(t, func() {
		RegisterObjectCreate(name+"_create", fn)
	})

	RegisterObjectInit(name+"_init", fn, 1)
	require.Panics(t, func() {
		RegisterObjectInit(name+"_init", fn, 1)
	})

	RegisterObjectUpdate(name+"_update", fn, 1)
	require.Panics(t, func() {
		RegisterObjectUpdate(name+"_update", fn, 1)
	})

	RegisterObjectCollide(name+"_collide", fn, 1)
	require.Panics(t, func() {
		RegisterObjectCollide(name+"_collide", fn, 1)
	})

	RegisterObjectUse(name+"_use", fn, 1)
	require.Panics(t, func() {
		RegisterObjectUse(name+"_use", fn, 1)
	})

	RegisterObjectDamage(name+"_damage", fn)
	require.Panics(t, func() {
		RegisterObjectDamage(name+"_damage", fn)
	})

	RegisterObjectDeath(name+"_death", fn, 1)
	require.Panics(t, func() {
		RegisterObjectDeath(name+"_death", fn, 1)
	})

	RegisterObjectDrop(name+"_drop", fn)
	require.Panics(t, func() {
		RegisterObjectDrop(name+"_drop", fn)
	})

	RegisterObjectPickup(name+"_pickup", fn)
	require.Panics(t, func() {
		RegisterObjectPickup(name+"_pickup", fn)
	})
}
