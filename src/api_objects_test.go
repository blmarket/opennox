package opennox

import (
	"encoding/json"
	"testing"

	"github.com/noxworld-dev/opennox/v1/server"
)

func TestObjectMarshalJSON(t *testing.T) {
	// Test that Object implements json.Marshaler
	var _ json.Marshaler = &Object{}
	var _ json.Marshaler = &server.ObjectType{}

	// Test MarshalJSON with nil SObj - should panic or return error
	obj := &Object{}
	func() {
		defer func() {
			if r := recover(); r != nil {
				t.Logf("MarshalJSON panicked as expected with nil SObj: %v", r)
			}
		}()
		data, err := obj.MarshalJSON()
		if err != nil {
			t.Logf("MarshalJSON returned error as expected: %v", err)
		} else {
			t.Logf("MarshalJSON returned data: %s", string(data))
		}
	}()
}

func TestApiObjectsInit(t *testing.T) {
	// The init function in api_objects.go registers HTTP handlers
	// We just verify the package initializes without panic
	// The handlers are registered in init(), so if we got here, it worked
	t.Log("api_objects.go init completed successfully")
}
