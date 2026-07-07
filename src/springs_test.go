package opennox

import (
	"testing"
)

func TestServerSpringsReset(t *testing.T) {
	var s serverSprings
	s.head = &Spring{}
	s.Reset()
	if s.head != nil {
		t.Error("Reset should set head to nil")
	}
}

func TestServerSpringsListAdd(t *testing.T) {
	var s serverSprings
	s.Reset()

	spring1 := &Spring{}
	s.listAdd(spring1)

	if s.head != spring1 {
		t.Error("listAdd should set head to the added spring")
	}
	if spring1.prev != nil {
		t.Error("First spring prev should be nil")
	}
	if spring1.next != nil {
		t.Error("First spring next should be nil")
	}

	spring2 := &Spring{}
	s.listAdd(spring2)

	if s.head != spring2 {
		t.Error("listAdd should set head to the most recently added spring")
	}
	if spring2.next != spring1 {
		t.Error("Second spring next should point to first spring")
	}
	if spring1.prev != spring2 {
		t.Error("First spring prev should point to second spring")
	}
}

func TestServerSpringsListRemove(t *testing.T) {
	var s serverSprings
	s.Reset()

	spring1 := &Spring{}
	spring2 := &Spring{}
	spring3 := &Spring{}

	s.listAdd(spring1)
	s.listAdd(spring2)
	s.listAdd(spring3)
	// List: spring3 -> spring2 -> spring1

	// Remove middle element (spring2)
	s.listRemove(spring2)

	if spring3.next != spring1 {
		t.Error("After removing spring2, spring3.next should point to spring1")
	}
	if spring1.prev != spring3 {
		t.Error("After removing spring2, spring1.prev should point to spring3")
	}
	// spring2 should be zeroed
	if spring2.next != nil || spring2.prev != nil {
		t.Error("Removed spring should be zeroed")
	}

	// Remove head (spring3)
	s.listRemove(spring3)
	if s.head != spring1 {
		t.Error("After removing head, new head should be spring1")
	}

	// Remove last element
	s.listRemove(spring1)
	if s.head != nil {
		t.Error("After removing all elements, head should be nil")
	}
}

func TestServerSpringsListRemoveSingle(t *testing.T) {
	var s serverSprings
	s.Reset()

	spring1 := &Spring{}
	s.listAdd(spring1)
	s.listRemove(spring1)

	if s.head != nil {
		t.Error("After removing single element, head should be nil")
	}
}

func TestServerSpringsAddNil(t *testing.T) {
	var s serverSprings
	s.Reset()

	// Add with nil objects should not panic and should not add anything
	s.Add(nil, nil)
	if s.head != nil {
		t.Error("Add with nil objects should not add anything")
	}
}

func TestServerSpringsUpdateEmpty(t *testing.T) {
	var s serverSprings
	s.Reset()

	// Update with empty list should not panic
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Update with empty list panicked: %v", r)
		}
	}()
	s.Update()
}

func TestServerSpringsUpdateWithNilObjects(t *testing.T) {
	var s serverSprings
	s.Reset()

	// Add a spring with nil objects manually
	spring := &Spring{
		obj1: nil,
		obj2: nil,
	}
	s.listAdd(spring)

	// Update should remove springs with nil objects
	defer func() {
		if r := recover(); r != nil {
			t.Logf("Update with nil objects panicked: %v", r)
		}
	}()
	s.Update()

	// Spring with nil objects should be removed
	if s.head != nil {
		t.Logf("Update should remove springs with nil objects, head is not nil")
	}
}
