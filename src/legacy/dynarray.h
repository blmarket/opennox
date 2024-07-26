#ifndef NOX_DYNARRAY_H
#define NOX_DYNARRAY_H

#include <stdint.h>

// Allocate a new dynarray
uint32_t* sub_4BD280(int a1, int a2);
// Push an element back to the dynarray
void* sub_4BD300(void* a1p, void* a2p);
// Destroy an array
void sub_4BD2D0(void* lpMem);
// Get an element to use from the dynarray
uint32_t* sub_4BD2E0(void* a1);

#endif // NOX_DYNARRAY_H