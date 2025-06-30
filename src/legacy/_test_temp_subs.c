#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "temp_subs.h"

// Mock types and globals
uint32_t dword_5d4594_816368 = 0;
uint32_t dword_5d4594_816372 = 0;

// Mock memory area
static char mock_memory[0x100000];

// Simple hash function for memory footprint
uint32_t hash_memory(const void* data, size_t len) {
	const uint8_t* bytes = (const uint8_t*)data;
	uint32_t hash = 0x811c9dc5; // FNV-1a 32-bit offset basis
	for (size_t i = 0; i < len; i++) {
		hash ^= bytes[i];
		hash *= 0x01000193; // FNV-1a 32-bit prime
	}
	return hash;
}

void* mem_getPtrSize(uintptr_t base, uintptr_t off, uintptr_t size) { return &mock_memory[off]; }
void* mem_getPtr(uintptr_t base, uintptr_t off) { return mem_getPtrSize(base, off, 4); }
int32_t* mem_getI32Ptr(uintptr_t base, uintptr_t off) { return (int32_t*)mem_getPtrSize(base, off, sizeof(int32_t)); }

void sub_43DD10(void* a1) { printf("sub_43DD10 called with %p\n", (uintptr_t)a1 - (uintptr_t)mock_memory); }

void sub_43D9E0(void* a1) { printf("sub_43D9E0 called with %p\n", (uintptr_t)a1 - (uintptr_t)mock_memory); }

int main() {
	printf("Testing temp_subs functions:\n");

	// Test initial state
	uint32_t initial_hash = hash_memory(mock_memory, sizeof(mock_memory));
	printf("Initial state: %d, memory hash: 0x%08x\n", sub_43DB20(), initial_hash);

	// Test sub_43DA80 - should succeed first 6 times
	for (int i = 0; i < 8; i++) {
		int result = sub_43DA80();
		uint32_t hash = hash_memory(mock_memory, sizeof(mock_memory));
		printf("sub_43DA80() call %d: result=%d, state=%d, memory hash: 0x%08x\n", i + 1, result, sub_43DB20(), hash);
	}

	// Test sub_43DAD0 - should reset
	printf("Calling sub_43DAD0()\n");
	sub_43DAD0();
	uint32_t reset_hash = hash_memory(mock_memory, sizeof(mock_memory));
	printf("After reset: %d, memory hash: 0x%08x\n", sub_43DB20(), reset_hash);

	// Test sub_43DB30 - set state
	printf("Setting state to 3\n");
	sub_43DB30(3);
	uint32_t state_hash = hash_memory(mock_memory, sizeof(mock_memory));
	printf("Current state: %d, memory hash: 0x%08x\n", sub_43DB20(), state_hash);

	// Test sub_43DB40 - get memory pointer
	void* ptr = sub_43DB40(2);
	uint32_t ptr_hash = hash_memory(mock_memory, sizeof(mock_memory));
	printf("sub_43DB40(2) returned: %p, memory hash: 0x%08x\n", ptr, ptr_hash);

	return 0;
}