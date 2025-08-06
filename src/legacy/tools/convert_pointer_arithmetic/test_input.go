package main

import "unsafe"

func testFunction(v1 *SomeStruct) {
	// Test case 1: dereferenced pointer arithmetic
	value1 := *&v1.field_12

	// Test case 2: non-dereferenced pointer arithmetic
	ptr1 := &v1.field_8

	// Test case 3: int32 type
	value2 := *&v1.field_5

	// Test case 4: simple integer offset
	value3 := *&v1.field_4

	_ = value1
	_ = ptr1
	_ = value2
	_ = value3
}

type SomeStruct struct {
	field_0  uint32
	field_1  uint32
	field_2  uint32
	field_3  uint32
	field_4  uint32
	field_5  int32
	field_6  uint32
	field_7  uint32
	field_8  uint32
	field_9  uint32
	field_10 uint32
	field_11 uint32
	field_12 uint32
}
