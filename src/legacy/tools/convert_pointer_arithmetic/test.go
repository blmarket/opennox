package main

import "unsafe"

func test() {
	var v1 *struct{}
	
	// Test case 1: Dereferenced pattern
	value1 := *(*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
	
	// Test case 2: Non-dereferenced pattern  
	ptr1 := (*uint32)(unsafe.Add(unsafe.Pointer(v1), 4*12))
	
	// Test case 3: Different offset
	value2 := *(*int32)(unsafe.Add(unsafe.Pointer(v1), 8))
	
	// Test case 4: Simple offset
	ptr2 := (*uint32)(unsafe.Add(unsafe.Pointer(v1), 16))
	
	_ = value1
	_ = ptr1
	_ = value2
	_ = ptr2
}