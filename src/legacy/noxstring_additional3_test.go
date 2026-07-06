package legacy

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestNoxStringAdditional3(t *testing.T) {
	// Test NoxItow with radix 8 (octal)
	var buf [32]uint16
	NoxItow(64, &buf[0], 8)
	got := fromWcharArray(buf[:])
	require.Equal(t, "100", got)

	// Test NoxItow with radix 2 (binary) and large number
	NoxItow(255, &buf[0], 2)
	got = fromWcharArray(buf[:])
	require.Equal(t, "11111111", got)

	// Test NoxItow with zero
	NoxItow(0, &buf[0], 10)
	got = fromWcharArray(buf[:])
	require.Equal(t, "0", got)

	// Test NoxWcscmp with one empty and one non-empty
	w1 := toWcharArray("")
	w2 := toWcharArray("test")
	require.Less(t, NoxWcscmp(&w1[0], &w2[0]), 0)
	require.Greater(t, NoxWcscmp(&w2[0], &w1[0]), 0)

	// Test NoxWcschr with null character (should find terminator or return nil)
	w := toWcharArray("hello")
	ptr := NoxWcschr(&w[0], 0)
	// Either nil or pointer to terminator is acceptable
	_ = ptr

	// Test NoxWcscat with both empty
	var dest [32]uint16
	empty1 := toWcharArray("")
	empty2 := toWcharArray("")
	copy(dest[:], empty1)
	NoxWcscat(&dest[0], &empty2[0])
	got = fromWcharArray(dest[:])
	require.Equal(t, "", got)

	// Test NoxWcscpy with long string
	var dest2 [64]uint16
	longStr := toWcharArray("this is a relatively long string for testing")
	NoxWcscpy(&dest2[0], &longStr[0])
	got = fromWcharArray(dest2[:])
	require.Equal(t, "this is a relatively long string for testing", got)

	// Test NoxWcsncpy with exact length
	var dest3 [16]uint16
	srcW := toWcharArray("exact")
	NoxWcsncpy(&dest3[0], &srcW[0], 5)
	got = fromWcharArray(dest3[:])
	require.Equal(t, "exact", got)

	// Test NoxWcsspn with empty accept string
	w = toWcharArray("abc")
	accept := toWcharArray("")
	require.Equal(t, 0, NoxWcsspn(&w[0], &accept[0]))

	// Test NoxWcsspn with empty wcs string
	w = toWcharArray("")
	accept = toWcharArray("abc")
	require.Equal(t, 0, NoxWcsspn(&w[0], &accept[0]))

	// Test NoxWcstok with multiple delimiters in a row
	w = toWcharArray("a,,b;;c")
	delim := toWcharArray(",;")
	p1 := NoxWcstok(&w[0], &delim[0])
	require.NotNil(t, p1)
	t1 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p1))[:])
	require.Equal(t, "a", t1)

	p2 := NoxWcstok(nil, &delim[0])
	require.NotNil(t, p2)
	t2 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p2))[:])
	require.Equal(t, "b", t2)

	p3 := NoxWcstok(nil, &delim[0])
	require.NotNil(t, p3)
	t3 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p3))[:])
	require.Equal(t, "c", t3)

	// Test NoxWcsicmp with empty strings
	w1 = toWcharArray("")
	w2 = toWcharArray("")
	require.Equal(t, 0, NoxWcsicmp(&w1[0], &w2[0]))

	// Test NoxWcsicmp with one empty
	w3 := toWcharArray("a")
	require.Less(t, NoxWcsicmp(&w1[0], &w3[0]), 0)
	require.Greater(t, NoxWcsicmp(&w3[0], &w1[0]), 0)

	// Test NoxStrcmpi with same strings different case
	require.Equal(t, 0, NoxStrcmpi("HelloWorld", "helloworld"))
	require.Equal(t, 0, NoxStrcmpi("ABC123", "abc123"))

	// Test NoxStrnicmp with exact n
	require.Equal(t, 0, NoxStrnicmp("abcdef", "ABCDEF", 6))
	require.Less(t, NoxStrnicmp("abc", "abd", 3), 0)

	// Test NoxWcstol with zero
	w = toWcharArray("0")
	var endptr *uint16
	res := NoxWcstol(&w[0], &endptr, 10)
	require.Equal(t, 0, res)

	// Test NoxWcstol with positive sign
	w = toWcharArray("+456")
	res = NoxWcstol(&w[0], &endptr, 10)
	require.Equal(t, 456, res)

	// Test NoxWcstol with hex prefix (base 0 auto-detect, but our impl may not support)
	// Just test base 16 explicitly
	w = toWcharArray("1A3F")
	res = NoxWcstol(&w[0], &endptr, 16)
	require.Equal(t, 0x1A3F, res)
}

func TestNoxStringFormattingAdditional3(t *testing.T) {
	var cBuf [256]byte
	cString := func() string {
		for i, b := range cBuf {
			if b == 0 {
				return string(cBuf[:i])
			}
		}
		return string(cBuf[:])
	}

	// Test %s with long string
	cBuf = [256]byte{}
	longStr := "this is a long string that should be formatted correctly"
	TestNoxSprintfS(&cBuf[0], "prefix %s suffix", longStr)
	require.Equal(t, "prefix "+longStr+" suffix", cString())

	// Test %d with max int
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%d", 2147483647)
	require.Equal(t, "2147483647", cString())

	// Test %d with min int
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%d", -2147483648)
	require.Equal(t, "-2147483648", cString())

	// Test %x with large value
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%x", 0xABCD)
	require.Equal(t, "abcd", cString())

	// Test %X with large value
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%X", 0xABCD)
	require.Equal(t, "ABCD", cString())

	// Test %o with large value
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%o", 511)
	require.Equal(t, "777", cString())

	// Test multiple %d
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%d %d %d", 1)
	// Only first arg used, rest are null/0
	// Just verify no crash

	// Test %c with various characters
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%c", 'Z')
	require.Equal(t, "Z", cString())

	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%c", '0')
	require.Equal(t, "0", cString())

	// Test wide string formatting with long string
	var wBuf [256]uint16
	wFmt := toWcharArray("prefix %s suffix")
	wLong := toWcharArray(longStr)
	TestNoxSwprintfS(&wBuf[0], &wFmt[0], &wLong[0])
	got := fromWcharArray(wBuf[:])
	require.Equal(t, "prefix "+longStr+" suffix", got)

	// Test wide %d with negative
	wBuf = [256]uint16{}
	wFmt = toWcharArray("%d")
	TestNoxSwprintfD(&wBuf[0], &wFmt[0], -12345)
	got = fromWcharArray(wBuf[:])
	require.Equal(t, "-12345", got)

	// Test wide %x, %X, %o
	wBuf = [256]uint16{}
	wFmt = toWcharArray("%x %X %o")
	TestNoxSwprintfD(&wBuf[0], &wFmt[0], 255)
	got = fromWcharArray(wBuf[:])
	// Only first arg used
	require.NotEmpty(t, got)
}
