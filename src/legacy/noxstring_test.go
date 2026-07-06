package legacy

import (
	"testing"
	"unsafe"

	"github.com/noxworld-dev/opennox/v1/internal/binfile"
	"github.com/noxworld-dev/opennox/v1/legacy/common/alloc"
)

func toWcharArray(s string) []uint16 {
	res := make([]uint16, len(s)+1)
	for i, r := range s {
		res[i] = uint16(r)
	}
	res[len(s)] = 0
	return res
}

func fromWcharArray(w []uint16) string {
	var res []rune
	for _, val := range w {
		if val == 0 {
			break
		}
		res = append(res, rune(val))
	}
	return string(res)
}

func TestNoxStringItow(t *testing.T) {
	var buf [32]uint16
	NoxItow(12345, &buf[0], 10)
	got := fromWcharArray(buf[:])
	if got != "12345" {
		t.Errorf("expected 12345, got %q", got)
	}
}

func TestNoxStringWcscat(t *testing.T) {
	var dest [32]uint16
	destW := toWcharArray("hello")
	copy(dest[:], destW)

	srcW := toWcharArray(" world")
	NoxWcscat(&dest[0], &srcW[0])

	got := fromWcharArray(dest[:])
	if got != "hello world" {
		t.Errorf("expected 'hello world', got %q", got)
	}

	// Test nil src
	NoxWcscat(&dest[0], nil)
	got2 := fromWcharArray(dest[:])
	if got2 != "hello world" {
		t.Errorf("expected 'hello world', got %q", got2)
	}
}

func TestNoxStringWcschr(t *testing.T) {
	strW := toWcharArray("hello")
	ptr := NoxWcschr(&strW[0], uint16('e'))
	if ptr == nil {
		t.Fatalf("expected to find 'e'")
	}
	offset := uintptr(unsafe.Pointer(ptr)) - uintptr(unsafe.Pointer(&strW[0]))
	if offset != 2 { // 2 bytes for 1 uint16
		t.Errorf("expected offset 2, got %d", offset)
	}

	ptr2 := NoxWcschr(&strW[0], uint16('x'))
	if ptr2 != nil {
		t.Errorf("expected not to find 'x'")
	}
}

func TestNoxStringWcscmp(t *testing.T) {
	w1 := toWcharArray("abc")
	w2 := toWcharArray("abc")
	w3 := toWcharArray("abd")

	if res := NoxWcscmp(&w1[0], &w2[0]); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
	if res := NoxWcscmp(&w1[0], &w3[0]); res >= 0 {
		t.Errorf("expected negative, got %d", res)
	}
	if res := NoxWcscmp(&w3[0], &w1[0]); res <= 0 {
		t.Errorf("expected positive, got %d", res)
	}

	// Nil checks
	if res := NoxWcscmp(nil, nil); res != 0 {
		t.Errorf("expected 0 for both nil")
	}
	if res := NoxWcscmp(nil, &w1[0]); res >= 0 {
		t.Errorf("expected negative for first nil")
	}
	if res := NoxWcscmp(&w1[0], nil); res <= 0 {
		t.Errorf("expected positive for second nil")
	}
}

func TestNoxStringWcscpy(t *testing.T) {
	var dest [32]uint16
	src := toWcharArray("test")
	NoxWcscpy(&dest[0], &src[0])
	got := fromWcharArray(dest[:])
	if got != "test" {
		t.Errorf("expected 'test', got %q", got)
	}
}

func TestNoxStringWcslen(t *testing.T) {
	w := toWcharArray("hello")
	if l := NoxWcslen(&w[0]); l != 5 {
		t.Errorf("expected 5, got %d", l)
	}
}

func TestNoxStringWcsncpy(t *testing.T) {
	var dest [8]uint16
	src := toWcharArray("hello world")
	NoxWcsncpy(&dest[0], &src[0], 5)
	got := fromWcharArray(dest[:])
	if got != "hello" {
		t.Errorf("expected 'hello', got %q", got)
	}
}

func TestNoxStringWcsspn(t *testing.T) {
	w := toWcharArray("12345abc")
	accept := toWcharArray("1234567890")
	if res := NoxWcsspn(&w[0], &accept[0]); res != 5 {
		t.Errorf("expected 5, got %d", res)
	}
}

func TestNoxStringWcstok(t *testing.T) {
	w := toWcharArray("abc,def;ghi")
	delim := toWcharArray(",;")

	p1 := NoxWcstok(&w[0], &delim[0])
	if p1 == nil {
		t.Fatalf("expected first token")
	}
	t1 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p1))[:])
	if t1 != "abc" {
		t.Errorf("expected 'abc', got %q", t1)
	}

	p2 := NoxWcstok(nil, &delim[0])
	if p2 == nil {
		t.Fatalf("expected second token")
	}
	t2 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p2))[:])
	if t2 != "def" {
		t.Errorf("expected 'def', got %q", t2)
	}

	p3 := NoxWcstok(nil, &delim[0])
	if p3 == nil {
		t.Fatalf("expected third token")
	}
	t3 := fromWcharArray((*[100]uint16)(unsafe.Pointer(p3))[:])
	if t3 != "ghi" {
		t.Errorf("expected 'ghi', got %q", t3)
	}

	p4 := NoxWcstok(nil, &delim[0])
	if p4 != nil {
		t.Errorf("expected no more tokens")
	}
}

func TestNoxStringWcsicmp(t *testing.T) {
	w1 := toWcharArray("abc")
	w2 := toWcharArray("ABC")
	w3 := toWcharArray("ABD")

	if res := NoxWcsicmp(&w1[0], &w2[0]); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
	if res := NoxWcsicmp(&w1[0], &w3[0]); res >= 0 {
		t.Errorf("expected negative, got %d", res)
	}
}

func TestNoxStrcmpi(t *testing.T) {
	if res := NoxStrcmpi("abc", "ABC"); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
	if res := NoxStrcmpi("abc", "ABD"); res >= 0 {
		t.Errorf("expected negative, got %d", res)
	}
}

func TestNoxStrnicmp(t *testing.T) {
	if res := NoxStrnicmp("abc", "XYZ", 0); res != 0 {
		t.Errorf("expected 0 for zero length, got %d", res)
	}
	if res := NoxStrnicmp("abcde", "ABCXX", 3); res != 0 {
		t.Errorf("expected 0, got %d", res)
	}
	if res := NoxStrnicmp("abcde", "ABCXX", 4); res >= 0 {
		t.Errorf("expected negative, got %d", res)
	}
}

func TestNoxWcstol(t *testing.T) {
	w := toWcharArray("12345abc")
	var endptr *uint16
	res := NoxWcstol(&w[0], &endptr, 10)
	if res != 12345 {
		t.Errorf("expected 12345, got %d", res)
	}
	if endptr == nil {
		t.Fatalf("expected non-nil endptr")
	}
	offset := uintptr(unsafe.Pointer(endptr)) - uintptr(unsafe.Pointer(&w[0]))
	if offset != 10 { // 5 characters * 2 bytes = 10 bytes
		t.Errorf("expected offset 10, got %d", offset)
	}
}

func TestNoxSprintfAndSwprintf(t *testing.T) {
	var cBuf [128]byte

	helperGetStr := func() string {
		for i, b := range cBuf {
			if b == 0 {
				return string(cBuf[:i])
			}
		}
		return string(cBuf[:])
	}

	// Test integers with various flags, widths, precisions (asserting actual quirky C behavior)
	formats := []struct {
		fmt  string
		val  int
		want string
	}{
		{"%d", 42, "42"},
		{"%+d", 42, "42"}, // note: + flag is parsed but ignored by legacy nox_vsnprintf
		{"%05d", 42, "00042"},
		{"%.3d", 42, "042"},
		{"%-5d", 42, "   42"}, // note: - flag is parsed but behaves like right-align in legacy nox_vsnprintf
		{"%x", 255, "ff"},
		{"%X", 255, "FF"},
		{"%o", 8, "10"},
		{"%u", -1, "-1"}, // note: %u prints signed values in legacy nox_vsnprintf because nox_itoa takes C.int
	}

	for _, tc := range formats {
		cBuf = [128]byte{}
		TestNoxSprintfD(&cBuf[0], tc.fmt, tc.val)
		got := helperGetStr()
		if got != tc.want {
			t.Errorf("nox_sprintf(%q, %d): expected %q, got %q", tc.fmt, tc.val, tc.want, got)
		}
	}

	// Test strings
	cBuf = [128]byte{}
	TestNoxSprintfS(&cBuf[0], "hello %s", "world")
	if got := helperGetStr(); got != "hello world" {
		t.Errorf("expected 'hello world', got %q", got)
	}

	// Test wide string formatting to char
	cBuf = [128]byte{}
	wWorld := toWcharArray("world")
	TestNoxSprintfWS(&cBuf[0], "hello %S", &wWorld[0])
	if got := helperGetStr(); got != "hello world" {
		t.Errorf("expected 'hello world', got %q", got)
	}

	// Test percent escape and exclamation mark
	cBuf = [128]byte{}
	TestNoxSprintfD(&cBuf[0], "%% %!", 0)
	if got := helperGetStr(); got != "% !" {
		t.Errorf("expected '%% !', got %q", got)
	}

	// Test nox_swprintf
	var wBuf [128]uint16

	wFormats := []struct {
		fmt  string
		val  int
		want string
	}{
		{"%d", 42, "42"},
		{"%+d", 42, "42"}, // note: ignored
		{"%05d", 42, "00042"},
		{"%.3d", 42, "042"},
		{"%-5d", 42, "   42"}, // note: right-aligned
		{"%x", 255, "ff"},
		{"%X", 255, "FF"},
		{"%o", 8, "10"},
		{"%u", -1, "-1"}, // note: signed
	}

	for _, tc := range wFormats {
		wBuf = [128]uint16{}
		wFmt := toWcharArray(tc.fmt)
		TestNoxSwprintfD(&wBuf[0], &wFmt[0], tc.val)
		got := fromWcharArray(wBuf[:])
		if got != tc.want {
			t.Errorf("nox_swprintf(%q, %d): expected %q, got %q", tc.fmt, tc.val, tc.want, got)
		}
	}

	wBuf = [128]uint16{}
	wFormatS := toWcharArray("hello %s")
	wWorld2 := toWcharArray("world")
	TestNoxSwprintfS(&wBuf[0], &wFormatS[0], &wWorld2[0])
	if got := fromWcharArray(wBuf[:]); got != "hello world" {
		t.Errorf("expected 'hello world', got %q", got)
	}

	wBuf = [128]uint16{}
	wFormatS2 := toWcharArray("hello %S")
	TestNoxSwprintfWS(&wBuf[0], &wFormatS2[0], "world")
	if got := fromWcharArray(wBuf[:]); got != "hello world" {
		t.Errorf("expected 'hello world', got %q", got)
	}

	wBuf = [128]uint16{}
	wFormatF := toWcharArray("%.2f")
	TestNoxSwprintfF(&wBuf[0], &wFormatF[0], 3.14159)
	if got := fromWcharArray(wBuf[:]); got != "3.14" {
		t.Errorf("expected '3.14', got %q", got)
	}
}

func TestNoxStringWcsncpyPadsWithZeros(t *testing.T) {
	dest := [6]uint16{0xffff, 0xffff, 0xffff, 0xffff, 0xffff, 0xffff}
	src := toWcharArray("hi")

	NoxWcsncpy(&dest[0], &src[0], 5)

	want := [6]uint16{'h', 'i', 0, 0, 0, 0xffff}
	if dest != want {
		t.Errorf("expected %v, got %v", want, dest)
	}
}

func TestNoxStringWcstokDelimiterOnly(t *testing.T) {
	w := toWcharArray(",,;")
	delim := toWcharArray(",;")

	if got := NoxWcstok(&w[0], &delim[0]); got != nil {
		t.Fatalf("expected no token, got %p", got)
	}
}

func TestNoxMemfile(t *testing.T) {
	// Prepare test data
	payload := []byte{
		0x01,       // i8
		0x02,       // u8
		0x03, 0x00, // i16: 3
		0x04, 0x00, // u16: 4
		0x05, 0x00, 0x00, 0x00, // i32: 5
		0x06, 0x00, 0x00, 0x00, // u32: 6
		0xaa, 0xbb, 0xcc, 0xdd, // read block
		0x00,                         // skip 1
		0x00, 0x00, 0x00, 0x00, 0x00, // alignment padding
		0x11, 0x22, 0x33, 0x44, 0x55, 0x66, 0x77, 0x88, // aligned block
	}

	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()

	cFile := f.C()

	// 1. Read i8
	if got := NoxMemfileReadI8(cFile); got != 1 {
		t.Errorf("expected 1, got %d", got)
	}

	// 2. Read u8
	if got := NoxMemfileReadU8(cFile); got != 2 {
		t.Errorf("expected 2, got %d", got)
	}

	// 3. Read i16
	if got := NoxMemfileReadI16(cFile); got != 3 {
		t.Errorf("expected 3, got %d", got)
	}

	// 4. Read u16
	if got := NoxMemfileReadU16(cFile); got != 4 {
		t.Errorf("expected 4, got %d", got)
	}

	// 5. Read i32
	if got := NoxMemfileReadI32(cFile); got != 5 {
		t.Errorf("expected 5, got %d", got)
	}

	// 6. Read u32
	if got := NoxMemfileReadU32(cFile); got != 6 {
		t.Errorf("expected 6, got %d", got)
	}

	// 7. Read block
	var buf [4]byte
	n := uint(NoxMemfileRead(unsafe.Pointer(&buf[0]), 1, 4, cFile))
	if n != 4 {
		t.Errorf("expected read 4, got %d", n)
	}
	if buf != [4]byte{0xaa, 0xbb, 0xcc, 0xdd} {
		t.Errorf("expected aa bb cc dd, got % x", buf)
	}

	// 8. Skip 1 byte
	NoxMemfileSkip(cFile, 1)

	// 9. Read 64align
	var alignBuf [4]byte
	n = NoxMemfileRead64Align(unsafe.Pointer(&alignBuf[0]), 1, 4, cFile)
	if n != 1 {
		t.Errorf("expected read64align 1, got %d", n)
	}
	if alignBuf != [4]byte{0x11, 0x22, 0x33, 0x44} {
		t.Errorf("expected 11 22 33 44, got % x", alignBuf)
	}
}

func TestNoxMemfileNull(t *testing.T) {
	// Allocate a zeroed block representing a nox_memfile struct with NULL data
	nullMemfile := make([]byte, 16)
	cFile := unsafe.Pointer(&nullMemfile[0])

	if got := NoxMemfileReadI8(cFile); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
	if got := NoxMemfileReadU8(cFile); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
	if got := NoxMemfileReadI16(cFile); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
	if got := NoxMemfileReadU16(cFile); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
	if got := NoxMemfileReadI32(cFile); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
	if got := NoxMemfileReadU32(cFile); got != 0 {
		t.Errorf("expected 0, got %d", got)
	}
	// should not panic/crash
	NoxMemfileSkip(cFile, 10)
}

func TestNoxMemfileShortReads(t *testing.T) {
	payload := []byte{0xaa, 0xbb, 0xcc, 0xdd}
	raw, _ := alloc.CloneSlice(payload)
	f := binfile.NewMemFile(unsafe.Pointer(&raw[0]), len(raw))
	defer f.Free()

	var aligned [4]byte
	if got := NoxMemfileRead64Align(unsafe.Pointer(&aligned[0]), 1, len(aligned), f.C()); got != 0 {
		t.Fatalf("short read64align result = %d, want 0", got)
	}
	if aligned != [4]byte{} {
		t.Fatalf("short read64align wrote % x, want zeros", aligned)
	}
}

func TestNoxStringFormattingBranches(t *testing.T) {
	var cBuf [128]byte
	cString := func() string {
		for i, b := range cBuf {
			if b == 0 {
				return string(cBuf[:i])
			}
		}
		return string(cBuf[:])
	}

	tests := []struct {
		name string
		run  func()
		want string
	}{
		{
			name: "char",
			run:  func() { TestNoxSprintfD(&cBuf[0], "%c", 'A') },
			want: "A",
		},
		{
			name: "null string",
			run:  func() { TestNoxSprintfNullS(&cBuf[0], "%s") },
			want: "(null)",
		},
		{
			name: "null wide string",
			run:  func() { TestNoxSprintfNullWS(&cBuf[0], "%S") },
			want: "(null)",
		},
		{
			name: "hex space width",
			run:  func() { TestNoxSprintfD(&cBuf[0], "%5x", 15) },
			want: "    f",
		},
		{
			name: "hex zero width",
			run:  func() { TestNoxSprintfD(&cBuf[0], "%05x", 15) },
			want: "0000f",
		},
		{
			name: "hex precision",
			run:  func() { TestNoxSprintfD(&cBuf[0], "%.3x", 15) },
			want: "00f",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cBuf = [128]byte{}
			tc.run()
			if got := cString(); got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}

	cBuf = [128]byte{}
	if n := TestNoxVsnprintfD(&cBuf[0], 0, "%05x", 15); n != 5 {
		t.Fatalf("nox_vsnprintf count=0 returned %d, want 5", n)
	}
	if cBuf[0] != 0 {
		t.Fatalf("nox_vsnprintf count=0 wrote byte %d", cBuf[0])
	}

	var wBuf [128]uint16
	wideTests := []struct {
		name   string
		format string
		run    func(format *uint16)
		want   string
	}{
		{
			name:   "char",
			format: "%c",
			run:    func(format *uint16) { TestNoxSwprintfD(&wBuf[0], format, 'Z') },
			want:   "Z",
		},
		{
			name:   "null wide string",
			format: "%s",
			run:    func(format *uint16) { TestNoxSwprintfNullS(&wBuf[0], format) },
			want:   "(null)",
		},
		{
			name:   "null string",
			format: "%S",
			run:    func(format *uint16) { TestNoxSwprintfNullWS(&wBuf[0], format) },
			want:   "(null)",
		},
		{
			name:   "hex space width",
			format: "%5x",
			run:    func(format *uint16) { TestNoxSwprintfD(&wBuf[0], format, 15) },
			want:   "    f",
		},
		{
			name:   "hex zero width",
			format: "%05x",
			run:    func(format *uint16) { TestNoxSwprintfD(&wBuf[0], format, 15) },
			want:   "0000f",
		},
		{
			name:   "hex precision",
			format: "%.3x",
			run:    func(format *uint16) { TestNoxSwprintfD(&wBuf[0], format, 15) },
			want:   "00f",
		},
		{
			name:   "percent bang",
			format: "%% %!",
			run:    func(format *uint16) { TestNoxSwprintfD(&wBuf[0], format, 0) },
			want:   "% !",
		},
		{
			name:   "vswprintf",
			format: "%d",
			run:    func(format *uint16) { TestNoxVswprintfD(&wBuf[0], format, 77) },
			want:   "77",
		},
	}
	for _, tc := range wideTests {
		t.Run("wide "+tc.name, func(t *testing.T) {
			wBuf = [128]uint16{}
			format := toWcharArray(tc.format)
			tc.run(&format[0])
			if got := fromWcharArray(wBuf[:]); got != tc.want {
				t.Fatalf("expected %q, got %q", tc.want, got)
			}
		})
	}

	wBuf = [128]uint16{}
	format := toWcharArray("%05x")
	if n := TestNoxVsnwprintfD(&wBuf[0], 0, &format[0], 15); n != 5 {
		t.Fatalf("nox_vsnwprintf count=0 returned %d, want 5", n)
	}
	if wBuf[0] != 0 {
		t.Fatalf("nox_vsnwprintf count=0 wrote word %d", wBuf[0])
	}
}
