package legacy

/*
#include "noxstring.h"
#include "memfile.h"
#include <stdarg.h>
#include <stdlib.h>

int test_nox_sprintf_d(void* str, void* format, int val) {
	return nox_sprintf((char*)str, (const char*)format, val);
}

int test_nox_sprintf_s(void* str, void* format, void* s) {
	return nox_sprintf((char*)str, (const char*)format, (const char*)s);
}

int test_nox_sprintf_S(void* str, void* format, void* s) {
	return nox_sprintf((char*)str, (const char*)format, (const wchar2_t*)s);
}

int test_nox_sprintf_f(void* str, void* format, double f) {
	return nox_sprintf((char*)str, (const char*)format, f);
}

int test_nox_sprintf_null_s(void* str, void* format) {
	return nox_sprintf((char*)str, (const char*)format, (const char*)NULL);
}

int test_nox_sprintf_null_S(void* str, void* format) {
	return nox_sprintf((char*)str, (const char*)format, (const wchar2_t*)NULL);
}

static int test_nox_vsnprintf_d_var(char* str, size_t count, const char* format, ...) {
	int ret;
	va_list ap;
	va_start(ap, format);
	ret = nox_vsnprintf(str, count, format, ap);
	va_end(ap);
	return ret;
}

int test_nox_vsnprintf_d(void* str, size_t count, void* format, int val) {
	return test_nox_vsnprintf_d_var((char*)str, count, (const char*)format, val);
}

int test_nox_swprintf_d(void* str, void* format, int val) {
	return nox_swprintf((wchar2_t*)str, (const wchar2_t*)format, val);
}

int test_nox_swprintf_s(void* str, void* format, void* s) {
	return nox_swprintf((wchar2_t*)str, (const wchar2_t*)format, (const wchar2_t*)s);
}

int test_nox_swprintf_S(void* str, void* format, void* s) {
	return nox_swprintf((wchar2_t*)str, (const wchar2_t*)format, (const char*)s);
}

int test_nox_swprintf_f(void* str, void* format, double f) {
	return nox_swprintf((wchar2_t*)str, (const wchar2_t*)format, f);
}

int test_nox_swprintf_null_s(void* str, void* format) {
	return nox_swprintf((wchar2_t*)str, (const wchar2_t*)format, (const wchar2_t*)NULL);
}

int test_nox_swprintf_null_S(void* str, void* format) {
	return nox_swprintf((wchar2_t*)str, (const wchar2_t*)format, (const char*)NULL);
}

static int test_nox_vsnwprintf_d_var(wchar2_t* str, size_t count, const wchar2_t* format, ...) {
	int ret;
	va_list ap;
	va_start(ap, format);
	ret = nox_vsnwprintf(str, count, format, ap);
	va_end(ap);
	return ret;
}

int test_nox_vsnwprintf_d(void* str, size_t count, void* format, int val) {
	return test_nox_vsnwprintf_d_var((wchar2_t*)str, count, (const wchar2_t*)format, val);
}

static int test_nox_vswprintf_d_var(wchar2_t* str, const wchar2_t* format, ...) {
	int ret;
	va_list ap;
	va_start(ap, format);
	ret = nox_vswprintf(str, format, ap);
	va_end(ap);
	return ret;
}

int test_nox_vswprintf_d(void* str, void* format, int val) {
	return test_nox_vswprintf_d_var((wchar2_t*)str, (const wchar2_t*)format, val);
}
*/
import "C"

import (
	"unsafe"
)

func NoxItow(val int, s *uint16, radix int) {
	C.nox_itow(C.int(val), (*C.wchar2_t)(unsafe.Pointer(s)), C.int(radix))
}

func NoxWcscat(dest *uint16, src *uint16) {
	C.nox_wcscat((*C.wchar2_t)(unsafe.Pointer(dest)), (*C.wchar2_t)(unsafe.Pointer(src)))
}

func NoxWcschr(nox_wcs *uint16, wc uint16) *uint16 {
	res := C.nox_wcschr((*C.wchar2_t)(unsafe.Pointer(nox_wcs)), C.wchar2_t(wc))
	return (*uint16)(unsafe.Pointer(res))
}

func NoxWcscmp(s1 *uint16, s2 *uint16) int {
	return int(C.nox_wcscmp((*C.wchar2_t)(unsafe.Pointer(s1)), (*C.wchar2_t)(unsafe.Pointer(s2))))
}

func NoxWcscpy(dest *uint16, src *uint16) *uint16 {
	res := C.nox_wcscpy((*C.wchar2_t)(unsafe.Pointer(dest)), (*C.wchar2_t)(unsafe.Pointer(src)))
	return (*uint16)(unsafe.Pointer(res))
}

func NoxWcslen(nox_wcs *uint16) int {
	return int(C.nox_wcslen((*C.wchar2_t)(unsafe.Pointer(nox_wcs))))
}

func NoxWcsncpy(dest *uint16, src *uint16, n int) *uint16 {
	res := C.nox_wcsncpy((*C.wchar2_t)(unsafe.Pointer(dest)), (*C.wchar2_t)(unsafe.Pointer(src)), C.size_t(n))
	return (*uint16)(unsafe.Pointer(res))
}

func NoxWcsspn(nox_wcs *uint16, accept *uint16) int {
	return int(C.nox_wcsspn((*C.wchar2_t)(unsafe.Pointer(nox_wcs)), (*C.wchar2_t)(unsafe.Pointer(accept))))
}

func NoxWcstok(str *uint16, delim *uint16) *uint16 {
	res := C.nox_wcstok((*C.wchar2_t)(unsafe.Pointer(str)), (*C.wchar2_t)(unsafe.Pointer(delim)))
	return (*uint16)(unsafe.Pointer(res))
}

func NoxWcsicmp(string1 *uint16, string2 *uint16) int {
	return int(C._nox_wcsicmp((*C.wchar2_t)(unsafe.Pointer(string1)), (*C.wchar2_t)(unsafe.Pointer(string2))))
}

func NoxStrcmpi(string1 string, string2 string) int {
	c1 := C.CString(string1)
	c2 := C.CString(string2)
	defer C.free(unsafe.Pointer(c1))
	defer C.free(unsafe.Pointer(c2))
	return int(C.nox_strcmpi(c1, c2))
}

func NoxStrnicmp(string1 string, string2 string, sz int) int {
	c1 := C.CString(string1)
	c2 := C.CString(string2)
	defer C.free(unsafe.Pointer(c1))
	defer C.free(unsafe.Pointer(c2))
	return int(C.nox_strnicmp(c1, c2, C.int(sz)))
}

func NoxWcstol(nptr *uint16, endptr **uint16, base int) int {
	var cEndptr *C.wchar2_t
	res := C.nox_wcstol((*C.wchar2_t)(unsafe.Pointer(nptr)), &cEndptr, C.int(base))
	if endptr != nil {
		*endptr = (*uint16)(unsafe.Pointer(cEndptr))
	}
	return int(res)
}

func TestNoxSprintfD(str *byte, format string, val int) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_sprintf_d(unsafe.Pointer(str), unsafe.Pointer(cFormat), C.int(val)))
}

func TestNoxSprintfS(str *byte, format string, s string) int {
	cFormat := C.CString(format)
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cFormat))
	defer C.free(unsafe.Pointer(cStr))
	return int(C.test_nox_sprintf_s(unsafe.Pointer(str), unsafe.Pointer(cFormat), unsafe.Pointer(cStr)))
}

func TestNoxSprintfWS(str *byte, format string, ws *uint16) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_sprintf_S(unsafe.Pointer(str), unsafe.Pointer(cFormat), unsafe.Pointer(ws)))
}

func TestNoxSprintfF(str *byte, format string, f float64) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_sprintf_f(unsafe.Pointer(str), unsafe.Pointer(cFormat), C.double(f)))
}

func TestNoxSprintfNullS(str *byte, format string) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_sprintf_null_s(unsafe.Pointer(str), unsafe.Pointer(cFormat)))
}

func TestNoxSprintfNullWS(str *byte, format string) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_sprintf_null_S(unsafe.Pointer(str), unsafe.Pointer(cFormat)))
}

func TestNoxVsnprintfD(str *byte, count int, format string, val int) int {
	cFormat := C.CString(format)
	defer C.free(unsafe.Pointer(cFormat))
	return int(C.test_nox_vsnprintf_d(unsafe.Pointer(str), C.size_t(count), unsafe.Pointer(cFormat), C.int(val)))
}

func TestNoxSwprintfD(str *uint16, format *uint16, val int) int {
	return int(C.test_nox_swprintf_d(unsafe.Pointer(str), unsafe.Pointer(format), C.int(val)))
}

func TestNoxSwprintfS(str *uint16, format *uint16, s *uint16) int {
	return int(C.test_nox_swprintf_s(unsafe.Pointer(str), unsafe.Pointer(format), unsafe.Pointer(s)))
}

func TestNoxSwprintfWS(str *uint16, format *uint16, s string) int {
	cS := C.CString(s)
	defer C.free(unsafe.Pointer(cS))
	return int(C.test_nox_swprintf_S(unsafe.Pointer(str), unsafe.Pointer(format), unsafe.Pointer(cS)))
}

func TestNoxSwprintfF(str *uint16, format *uint16, f float64) int {
	return int(C.test_nox_swprintf_f(unsafe.Pointer(str), unsafe.Pointer(format), C.double(f)))
}

func TestNoxSwprintfNullS(str *uint16, format *uint16) int {
	return int(C.test_nox_swprintf_null_s(unsafe.Pointer(str), unsafe.Pointer(format)))
}

func TestNoxSwprintfNullWS(str *uint16, format *uint16) int {
	return int(C.test_nox_swprintf_null_S(unsafe.Pointer(str), unsafe.Pointer(format)))
}

func TestNoxVsnwprintfD(str *uint16, count int, format *uint16, val int) int {
	return int(C.test_nox_vsnwprintf_d(unsafe.Pointer(str), C.size_t(count), unsafe.Pointer(format), C.int(val)))
}

func TestNoxVswprintfD(str *uint16, format *uint16, val int) int {
	return int(C.test_nox_vswprintf_d(unsafe.Pointer(str), unsafe.Pointer(format), C.int(val)))
}

func NoxMemfileReadI8(f unsafe.Pointer) int8 {
	return int8(C.nox_memfile_read_i8((*C.nox_memfile)(f)))
}
func NoxMemfileReadU8(f unsafe.Pointer) uint8 {
	return uint8(C.nox_memfile_read_u8((*C.nox_memfile)(f)))
}
func NoxMemfileReadI16(f unsafe.Pointer) int16 {
	return int16(C.nox_memfile_read_i16((*C.nox_memfile)(f)))
}
func NoxMemfileReadU16(f unsafe.Pointer) uint16 {
	return uint16(C.nox_memfile_read_u16((*C.nox_memfile)(f)))
}
func NoxMemfileReadI32(f unsafe.Pointer) int32 {
	return int32(C.nox_memfile_read_i32((*C.nox_memfile)(f)))
}
func NoxMemfileReadU32(f unsafe.Pointer) uint32 {
	return uint32(C.nox_memfile_read_u32((*C.nox_memfile)(f)))
}
func NoxMemfileSkip(f unsafe.Pointer, n int) {
	C.nox_memfile_skip((*C.nox_memfile)(f), C.int(n))
}
func NoxMemfileRead(dst unsafe.Pointer, sz uint, cnt int, f unsafe.Pointer) uint {
	return uint(C.nox_memfile_read(dst, C.uint(sz), C.int(cnt), (*C.nox_memfile)(f)))
}
func NoxMemfileRead64Align(dest unsafe.Pointer, sz int, cnt int, f unsafe.Pointer) uint {
	return uint(C.nox_memfile_read64align_40AD60((*C.char)(dest), C.int(sz), C.int(cnt), (*C.nox_memfile)(f)))
}
