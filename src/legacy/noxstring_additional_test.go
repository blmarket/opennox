package legacy

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/require"
)

func TestNoxStringAdditional(t *testing.T) {
	// Test NoxItow with different radix
	var buf [32]uint16
	NoxItow(255, &buf[0], 16)
	got := fromWcharArray(buf[:])
	require.Equal(t, "ff", got)

	NoxItow(10, &buf[0], 2)
	got = fromWcharArray(buf[:])
	require.Equal(t, "1010", got)

	// Test NoxItow with negative numbers
	NoxItow(-42, &buf[0], 10)
	got = fromWcharArray(buf[:])
	require.Equal(t, "-42", got)

	// Test NoxWcscmp with empty strings
	w1 := toWcharArray("")
	w2 := toWcharArray("")
	require.Equal(t, 0, NoxWcscmp(&w1[0], &w2[0]))

	w3 := toWcharArray("a")
	require.Less(t, NoxWcscmp(&w1[0], &w3[0]), 0)
	require.Greater(t, NoxWcscmp(&w3[0], &w1[0]), 0)

	// Test NoxWcslen with empty string
	wEmpty := toWcharArray("")
	require.Equal(t, 0, NoxWcslen(&wEmpty[0]))

	// Test NoxWcslen with long string
	wLong := toWcharArray("abcdefghijklmnopqrstuvwxyz")
	require.Equal(t, 26, NoxWcslen(&wLong[0]))

	// Test NoxWcschr with first character
	w := toWcharArray("hello")
	ptr := NoxWcschr(&w[0], uint16('h'))
	require.NotNil(t, ptr)
	offset := uintptr(unsafe.Pointer(ptr)) - uintptr(unsafe.Pointer(&w[0]))
	require.Equal(t, uintptr(0), offset)

	// Test NoxWcschr with last character
	ptr = NoxWcschr(&w[0], uint16('o'))
	require.NotNil(t, ptr)

	// Test NoxWcscat with empty src
	var dest [32]uint16
	destW := toWcharArray("hello")
	copy(dest[:], destW)
	emptyW := toWcharArray("")
	NoxWcscat(&dest[0], &emptyW[0])
	got = fromWcharArray(dest[:])
	require.Equal(t, "hello", got)

	// Test NoxWcscpy with empty string
	var dest2 [32]uint16
	emptyW2 := toWcharArray("")
	NoxWcscpy(&dest2[0], &emptyW2[0])
	got = fromWcharArray(dest2[:])
	require.Equal(t, "", got)

	// Test NoxWcsncpy with n larger than src
	var dest3 [32]uint16
	srcW := toWcharArray("hi")
	NoxWcsncpy(&dest3[0], &srcW[0], 10)
	got = fromWcharArray(dest3[:])
	require.Equal(t, "hi", got)

	// Test NoxWcsspn with no matching characters
	w = toWcharArray("abc")
	accept := toWcharArray("xyz")
	require.Equal(t, 0, NoxWcsspn(&w[0], &accept[0]))

	// Test NoxWcsspn with all matching characters
	w = toWcharArray("aaa")
	accept = toWcharArray("a")
	require.Equal(t, 3, NoxWcsspn(&w[0], &accept[0]))

	// Test NoxWcstok with single token (no delimiters)
	w = toWcharArray("single")
	delim := toWcharArray(",;")
	p1 := NoxWcstok(&w[0], &delim[0])
	require.NotNil(t, p1)
	t1 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p1))[:])
	require.Equal(t, "single", t1)
	p2 := NoxWcstok(nil, &delim[0])
	require.Nil(t, p2)

	// Test NoxWcsicmp with different cases
	w1 = toWcharArray("Hello")
	w2 = toWcharArray("HELLO")
	require.Equal(t, 0, NoxWcsicmp(&w1[0], &w2[0]))

	w3 = toWcharArray("World")
	require.Less(t, NoxWcsicmp(&w1[0], &w3[0]), 0)

	// Test NoxStrcmpi with empty strings
	require.Equal(t, 0, NoxStrcmpi("", ""))
	require.Less(t, NoxStrcmpi("", "a"), 0)
	require.Greater(t, NoxStrcmpi("a", ""), 0)

	// Test NoxStrnicmp with n larger than strings
	require.Equal(t, 0, NoxStrnicmp("abc", "ABC", 10))
	require.Equal(t, 0, NoxStrnicmp("abc", "abd", 2))

	// Test NoxWcstol with different bases
	w = toWcharArray("ff")
	var endptr *uint16
	res := NoxWcstol(&w[0], &endptr, 16)
	require.Equal(t, 255, res)

	w = toWcharArray("1010")
	res = NoxWcstol(&w[0], &endptr, 2)
	require.Equal(t, 10, res)

	w = toWcharArray("77")
	res = NoxWcstol(&w[0], &endptr, 8)
	require.Equal(t, 63, res)

	// Test NoxWcstol with negative number
	w = toWcharArray("-123")
	res = NoxWcstol(&w[0], &endptr, 10)
	require.Equal(t, -123, res)
}

func TestNoxStringFormattingAdditional(t *testing.T) {
	var cBuf [256]byte
	cString := func() string {
		for i, b := range cBuf {
			if b == 0 {
				return string(cBuf[:i])
			}
		}
		return string(cBuf[:])
	}

	// Test %s with empty string
	cBuf = [256]byte{}
	TestNoxSprintfS(&cBuf[0], "test %s end", "")
	require.Equal(t, "test  end", cString())

	// Test %d with zero
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%d", 0)
	require.Equal(t, "0", cString())

	// Test %d with negative
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%d", -999)
	require.Equal(t, "-999", cString())

	// Test %x with zero
	cBuf = [256]byte{}
	TestNoxSprintfD(&cBuf[0], "%x", 0)
	require.Equal(t, "0", cString())

	// Test multiple format specifiers
	cBuf = [256]byte{}
	TestNoxSprintfS(&cBuf[0], "%s %s", "hello")
	// This will only use first arg, second %s gets null
	// Just verify it doesn't crash

	// Test wide string formatting with empty
	var wBuf [256]uint16
	wFmt := toWcharArray("test %s end")
	wEmpty := toWcharArray("")
	TestNoxSwprintfS(&wBuf[0], &wFmt[0], &wEmpty[0])
	got := fromWcharArray(wBuf[:])
	require.Equal(t, "test  end", got)

	// Test wide string with %d zero
	wBuf = [256]uint16{}
	wFmt = toWcharArray("%d")
	TestNoxSwprintfD(&wBuf[0], &wFmt[0], 0)
	got = fromWcharArray(wBuf[:])
	require.Equal(t, "0", got)
}
