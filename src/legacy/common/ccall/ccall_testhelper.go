//go:build cgo
// +build cgo

package ccall

/*
#include <stdint.h>

void test_void_void(void) {}
void test_void_uptr(uintptr_t a1) {}
void test_void_uptr2(uintptr_t a1, uintptr_t a2) {}
void test_void_uptr3(uintptr_t a1, uintptr_t a2, uintptr_t a3) {}
void test_void_uptr4(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) {}
void test_void_uptr5(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) {}
void test_void_uptr6(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) {}

void test_void_ptr(void* a1) {}
void test_void_ptr2(void* a1, void* a2) {}
void test_void_ptr3(void* a1, void* a2, void* a3) {}
void test_void_ptr4(void* a1, void* a2, void* a3, void* a4) {}
void test_void_ptr5(void* a1, void* a2, void* a3, void* a4, void* a5) {}
void test_void_ptr6(void* a1, void* a2, void* a3, void* a4, void* a5, void* a6) {}

void test_void_int(int a1) {}
void test_void_int2(int a1, int a2) {}
void test_void_int3(int a1, int a2, int a3) {}
void test_void_int4(int a1, int a2, int a3, int a4) {}
void test_void_int5(int a1, int a2, int a3, int a4, int a5) {}
void test_void_int6(int a1, int a2, int a3, int a4, int a5, int a6) {}

uintptr_t test_uptr_void(void) { return 1; }
uintptr_t test_uptr_uptr(uintptr_t a1) { return a1 + 1; }
uintptr_t test_uptr_uptr2(uintptr_t a1, uintptr_t a2) { return a1 + a2; }
uintptr_t test_uptr_uptr3(uintptr_t a1, uintptr_t a2, uintptr_t a3) { return a1 + a2 + a3; }
uintptr_t test_uptr_uptr4(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) { return a1 + a2 + a3 + a4; }
uintptr_t test_uptr_uptr5(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) { return a1 + a2 + a3 + a4 + a5; }
uintptr_t test_uptr_uptr6(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) { return a1 + a2 + a3 + a4 + a5 + a6; }

uintptr_t test_uptr_ptr(void* a1) { return 2; }
uintptr_t test_uptr_ptr2(void* a1, void* a2) { return 3; }
uintptr_t test_uptr_ptr3(void* a1, void* a2, void* a3) { return 4; }
uintptr_t test_uptr_ptr4(void* a1, void* a2, void* a3, void* a4) { return 5; }
uintptr_t test_uptr_ptr5(void* a1, void* a2, void* a3, void* a4, void* a5) { return 6; }
uintptr_t test_uptr_ptr6(void* a1, void* a2, void* a3, void* a4, void* a5, void* a6) { return 7; }

uintptr_t test_uptr_int(int a1) { return (uintptr_t)a1 + 10; }
uintptr_t test_uptr_int2(int a1, int a2) { return (uintptr_t)(a1 + a2); }
uintptr_t test_uptr_int3(int a1, int a2, int a3) { return (uintptr_t)(a1 + a2 + a3); }
uintptr_t test_uptr_int4(int a1, int a2, int a3, int a4) { return (uintptr_t)(a1 + a2 + a3 + a4); }
uintptr_t test_uptr_int5(int a1, int a2, int a3, int a4, int a5) { return (uintptr_t)(a1 + a2 + a3 + a4 + a5); }
uintptr_t test_uptr_int6(int a1, int a2, int a3, int a4, int a5, int a6) { return (uintptr_t)(a1 + a2 + a3 + a4 + a5 + a6); }

void* test_ptr_void(void) { return (void*)1; }
void* test_ptr_uptr(uintptr_t a1) { return (void*)(a1 + 1); }
void* test_ptr_uptr2(uintptr_t a1, uintptr_t a2) { return (void*)(a1 + a2); }
void* test_ptr_uptr3(uintptr_t a1, uintptr_t a2, uintptr_t a3) { return (void*)(a1 + a2 + a3); }
void* test_ptr_uptr4(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) { return (void*)(a1 + a2 + a3 + a4); }
void* test_ptr_uptr5(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) { return (void*)(a1 + a2 + a3 + a4 + a5); }
void* test_ptr_uptr6(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) { return (void*)(a1 + a2 + a3 + a4 + a5 + a6); }

void* test_ptr_ptr(void* a1) { return a1; }
void* test_ptr_ptr2(void* a1, void* a2) { return a1; }
void* test_ptr_ptr3(void* a1, void* a2, void* a3) { return a1; }
void* test_ptr_ptr4(void* a1, void* a2, void* a3, void* a4) { return a1; }
void* test_ptr_ptr5(void* a1, void* a2, void* a3, void* a4, void* a5) { return a1; }
void* test_ptr_ptr6(void* a1, void* a2, void* a3, void* a4, void* a5, void* a6) { return a1; }

void* test_ptr_int(int a1) { return (void*)(uintptr_t)a1; }
void* test_ptr_int2(int a1, int a2) { return (void*)(uintptr_t)(a1 + a2); }
void* test_ptr_int3(int a1, int a2, int a3) { return (void*)(uintptr_t)(a1 + a2 + a3); }
void* test_ptr_int4(int a1, int a2, int a3, int a4) { return (void*)(uintptr_t)(a1 + a2 + a3 + a4); }
void* test_ptr_int5(int a1, int a2, int a3, int a4, int a5) { return (void*)(uintptr_t)(a1 + a2 + a3 + a4 + a5); }
void* test_ptr_int6(int a1, int a2, int a3, int a4, int a5, int a6) { return (void*)(uintptr_t)(a1 + a2 + a3 + a4 + a5 + a6); }

int test_int_void(void) { return 42; }
int test_int_uptr(uintptr_t a1) { return (int)a1 + 1; }
int test_int_uptr2(uintptr_t a1, uintptr_t a2) { return (int)(a1 + a2); }
int test_int_uptr3(uintptr_t a1, uintptr_t a2, uintptr_t a3) { return (int)(a1 + a2 + a3); }
int test_int_uptr4(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4) { return (int)(a1 + a2 + a3 + a4); }
int test_int_uptr5(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5) { return (int)(a1 + a2 + a3 + a4 + a5); }
int test_int_uptr6(uintptr_t a1, uintptr_t a2, uintptr_t a3, uintptr_t a4, uintptr_t a5, uintptr_t a6) { return (int)(a1 + a2 + a3 + a4 + a5 + a6); }

int test_int_ptr(void* a1) { return 1; }
int test_int_ptr2(void* a1, void* a2) { return 2; }
int test_int_ptr3(void* a1, void* a2, void* a3) { return 3; }
int test_int_ptr4(void* a1, void* a2, void* a3, void* a4) { return 4; }
int test_int_ptr5(void* a1, void* a2, void* a3, void* a4, void* a5) { return 5; }
int test_int_ptr6(void* a1, void* a2, void* a3, void* a4, void* a5, void* a6) { return 6; }

int test_int_int(int a1) { return a1 + 1; }
int test_int_int2(int a1, int a2) { return a1 + a2; }
int test_int_int3(int a1, int a2, int a3) { return a1 + a2 + a3; }
int test_int_int4(int a1, int a2, int a3, int a4) { return a1 + a2 + a3 + a4; }
int test_int_int5(int a1, int a2, int a3, int a4, int a5) { return a1 + a2 + a3 + a4 + a5; }
int test_int_int6(int a1, int a2, int a3, int a4, int a5, int a6) { return a1 + a2 + a3 + a4 + a5 + a6; }
*/
import "C"
import "unsafe"

func testPtrVoidVoid() unsafe.Pointer  { return unsafe.Pointer(C.test_void_void) }
func testPtrVoidUPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_void_uptr) }
func testPtrVoidUPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_void_uptr2) }
func testPtrVoidUPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_void_uptr3) }
func testPtrVoidUPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_void_uptr4) }
func testPtrVoidUPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_void_uptr5) }
func testPtrVoidUPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_void_uptr6) }

func testPtrVoidPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_void_ptr) }
func testPtrVoidPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_void_ptr2) }
func testPtrVoidPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_void_ptr3) }
func testPtrVoidPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_void_ptr4) }
func testPtrVoidPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_void_ptr5) }
func testPtrVoidPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_void_ptr6) }

func testPtrVoidInt() unsafe.Pointer  { return unsafe.Pointer(C.test_void_int) }
func testPtrVoidInt2() unsafe.Pointer { return unsafe.Pointer(C.test_void_int2) }
func testPtrVoidInt3() unsafe.Pointer { return unsafe.Pointer(C.test_void_int3) }
func testPtrVoidInt4() unsafe.Pointer { return unsafe.Pointer(C.test_void_int4) }
func testPtrVoidInt5() unsafe.Pointer { return unsafe.Pointer(C.test_void_int5) }
func testPtrVoidInt6() unsafe.Pointer { return unsafe.Pointer(C.test_void_int6) }

func testPtrUPtrVoid() unsafe.Pointer  { return unsafe.Pointer(C.test_uptr_void) }
func testPtrUPtrUPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_uptr_uptr) }
func testPtrUPtrUPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_uptr2) }
func testPtrUPtrUPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_uptr3) }
func testPtrUPtrUPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_uptr4) }
func testPtrUPtrUPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_uptr5) }
func testPtrUPtrUPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_uptr6) }

func testPtrUPtrPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_uptr_ptr) }
func testPtrUPtrPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_ptr2) }
func testPtrUPtrPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_ptr3) }
func testPtrUPtrPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_ptr4) }
func testPtrUPtrPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_ptr5) }
func testPtrUPtrPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_ptr6) }

func testPtrUPtrInt() unsafe.Pointer  { return unsafe.Pointer(C.test_uptr_int) }
func testPtrUPtrInt2() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_int2) }
func testPtrUPtrInt3() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_int3) }
func testPtrUPtrInt4() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_int4) }
func testPtrUPtrInt5() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_int5) }
func testPtrUPtrInt6() unsafe.Pointer { return unsafe.Pointer(C.test_uptr_int6) }

func testPtrPtrVoid() unsafe.Pointer  { return unsafe.Pointer(C.test_ptr_void) }
func testPtrPtrUPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_ptr_uptr) }
func testPtrPtrUPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_uptr2) }
func testPtrPtrUPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_uptr3) }
func testPtrPtrUPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_uptr4) }
func testPtrPtrUPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_uptr5) }
func testPtrPtrUPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_uptr6) }

func testPtrPtrPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_ptr_ptr) }
func testPtrPtrPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_ptr2) }
func testPtrPtrPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_ptr3) }
func testPtrPtrPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_ptr4) }
func testPtrPtrPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_ptr5) }
func testPtrPtrPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_ptr6) }

func testPtrPtrInt() unsafe.Pointer  { return unsafe.Pointer(C.test_ptr_int) }
func testPtrPtrInt2() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_int2) }
func testPtrPtrInt3() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_int3) }
func testPtrPtrInt4() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_int4) }
func testPtrPtrInt5() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_int5) }
func testPtrPtrInt6() unsafe.Pointer { return unsafe.Pointer(C.test_ptr_int6) }

func testPtrIntVoid() unsafe.Pointer  { return unsafe.Pointer(C.test_int_void) }
func testPtrIntUPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_int_uptr) }
func testPtrIntUPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_int_uptr2) }
func testPtrIntUPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_int_uptr3) }
func testPtrIntUPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_int_uptr4) }
func testPtrIntUPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_int_uptr5) }
func testPtrIntUPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_int_uptr6) }

func testPtrIntPtr() unsafe.Pointer  { return unsafe.Pointer(C.test_int_ptr) }
func testPtrIntPtr2() unsafe.Pointer { return unsafe.Pointer(C.test_int_ptr2) }
func testPtrIntPtr3() unsafe.Pointer { return unsafe.Pointer(C.test_int_ptr3) }
func testPtrIntPtr4() unsafe.Pointer { return unsafe.Pointer(C.test_int_ptr4) }
func testPtrIntPtr5() unsafe.Pointer { return unsafe.Pointer(C.test_int_ptr5) }
func testPtrIntPtr6() unsafe.Pointer { return unsafe.Pointer(C.test_int_ptr6) }

func testPtrIntInt() unsafe.Pointer  { return unsafe.Pointer(C.test_int_int) }
func testPtrIntInt2() unsafe.Pointer { return unsafe.Pointer(C.test_int_int2) }
func testPtrIntInt3() unsafe.Pointer { return unsafe.Pointer(C.test_int_int3) }
func testPtrIntInt4() unsafe.Pointer { return unsafe.Pointer(C.test_int_int4) }
func testPtrIntInt5() unsafe.Pointer { return unsafe.Pointer(C.test_int_int5) }
func testPtrIntInt6() unsafe.Pointer { return unsafe.Pointer(C.test_int_int6) }
